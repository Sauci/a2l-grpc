package main

import (
	"C"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/sauci/a2l-grpc/pkg/a2l"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"io"
	"net"
	"strings"
	"sync"
	"unicode/utf8"
)

// protocolSizeMargin is the part of the maximum message size a chunk of payload leaves to the
// framing of its message: the tag and the length prefix of the field which carries it.
const protocolSizeMargin = 256

// tagSize is the size of the tag of a field of the responses; they are all numbered below 16.
const tagSize = 1

// chunkifyBySize splits data into chunks of at most chunkSize bytes. A chunkSize which is not
// positive would make the loop below spin forever or slice backwards, so it is treated as "do not
// split": Create rejects such a configuration up front, this guard only keeps a direct caller from
// hanging the process.
func chunkifyBySize(data []byte, chunkSize int) [][]byte {
	if chunkSize <= 0 {
		if len(data) == 0 {
			return nil
		}

		return [][]byte{data}
	}

	var chunks [][]byte
	for start := 0; start < len(data); start += chunkSize {
		end := start + chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunks = append(chunks, data[start:end])
	}
	return chunks
}

func getTreeFromString(a2lString string) (result *a2l.RootNodeType, err error) {
	return a2l.GetTreeFromString(a2lString)
}

// grpcA2LImplType serves the API. Every response it sends fits maxMsgSize, the largest message
// either side of the connection accepts: gRPC rejects a larger one on the sending as well as on
// the receiving side, with a transport error which reaches the client without a word about the
// file. The payloads are chunked to that end, and the fields which are not payload, the warnings
// and the error, are bounded as well, since a large file may produce any number of either.
type grpcA2LImplType struct {
	a2l.UnimplementedA2LServer
	// maxMsgSize is the size limit of a message on the wire, in bytes
	maxMsgSize int
	// chunkSize is the number of bytes of payload a response carries; the rest of the limit is
	// left to the framing of the message
	chunkSize int
}

func newGrpcA2LImpl(maxMsgSize int) *grpcA2LImplType {
	return &grpcA2LImplType{maxMsgSize: maxMsgSize, chunkSize: maxMsgSize - protocolSizeMargin}
}

// stringFieldSize is the number of bytes a string takes on the wire as a field of a response.
func stringFieldSize(s string) int {
	return tagSize + protowire.SizeBytes(len(s))
}

// fitString shortens s so that a response carrying it alone fits the budget, and marks the cut.
// It only ever cuts when the server was created with a maximum message size of a few hundred
// bytes, which no single message of the parser fills otherwise.
func fitString(s string, budget int) string {
	const mark = "..."

	if stringFieldSize(s) <= budget {
		return s
	}

	// the length prefix of the shortened string is at most as long as the one of the budget
	n := budget - tagSize - protowire.SizeVarint(uint64(budget)) - len(mark)
	if n < 0 {
		n = 0
	}

	for n > 0 && !utf8.RuneStart(s[n]) {
		n--
	}

	return s[:n] + mark
}

// boundedError shortens a multi-line error so that a response carrying it fits the budget. The
// lines are the messages the parser reported, in order of position: the first ones name the cause
// and the following ones tend to be its consequences, so the tail is dropped and counted.
func boundedError(message string, budget int) string {
	if stringFieldSize(message) <= budget {
		return message
	}

	lines := strings.Split(message, "\n")
	tail := func(omitted int) string {
		return fmt.Sprintf("\n... %d further errors omitted", omitted)
	}

	// the tail is longest when it counts every line, and the length prefix of the result is at
	// most as long as the one of the budget
	room := budget - tagSize - protowire.SizeVarint(uint64(budget)) - len(tail(len(lines)))

	kept, size := 0, 0
	for kept < len(lines) && size+len(lines[kept])+1 <= room {
		size += len(lines[kept]) + 1
		kept++
	}

	if kept == 0 {
		return fitString(message, budget)
	}

	return strings.Join(lines[:kept], "\n") + tail(len(lines)-kept)
}

// warningResponses spreads the warnings over as many responses as their number requires for each
// response to fit the budget. They precede the chunks of the tree, which fill a message on their
// own and leave no room for them.
func warningResponses(warnings []a2l.SyntaxError, budget int) (result []*a2l.TreeResponse) {
	response, size := &a2l.TreeResponse{}, 0

	for _, warning := range warnings {
		text := fitString(warning.String(), budget)

		if size+stringFieldSize(text) > budget && len(response.Warnings) > 0 {
			result = append(result, response)
			response, size = &a2l.TreeResponse{}, 0
		}

		response.Warnings = append(response.Warnings, text)
		size += stringFieldSize(text)
	}

	if len(response.Warnings) > 0 {
		result = append(result, response)
	}

	return result
}

func treeErrorResponse(message string) *a2l.TreeResponse { return &a2l.TreeResponse{Error: &message} }
func jsonErrorResponse(message string) *a2l.JSONResponse { return &a2l.JSONResponse{Error: &message} }
func a2lErrorResponse(message string) *a2l.A2LResponse   { return &a2l.A2LResponse{Error: &message} }

// sendWithin hands a response to its stream once it is known to fit the maximum message size. The
// responses are built to fit, so a larger one is a bug of this server; it is reported through the
// error field of a replacement response, instead of being left to gRPC, whose transport error
// would reach the client without a word about the cause.
func sendWithin[M proto.Message](stream interface{ Send(M) error }, response M, maxMsgSize int,
	withError func(string) M) error {
	if size := proto.Size(response); size > maxMsgSize {
		response = withError(fmt.Sprintf(
			"internal error: a response of %d bytes exceeds the maximum message size of %d bytes",
			size, maxMsgSize))
	}

	return stream.Send(response)
}

func (s *grpcA2LImplType) GetTreeFromA2L(stream a2l.A2L_GetTreeFromA2LServer) error {
	var buffer bytes.Buffer
	options := a2l.ParseOptions{}
	// the options are read from the first request of the stream only
	optionsParsed := false

	for {
		request, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if !optionsParsed {
			if request.EnforceVersionCheck != nil {
				options.EnforceVersionCheck = *request.EnforceVersionCheck
			}
			optionsParsed = true
		}
		buffer.Write(request.A2L)
	}

	tree, warnings, parseError := a2l.GetTreeFromStringWithOptions(buffer.String(), options)

	// the warnings come first, in responses of their own, whether the file parsed or not
	for _, response := range warningResponses(warnings, s.maxMsgSize) {
		if err := sendWithin(stream, response, s.maxMsgSize, treeErrorResponse); err != nil {
			return err
		}
	}

	if parseError != nil {
		return sendWithin(stream, treeErrorResponse(boundedError(parseError.Error(), s.maxMsgSize)),
			s.maxMsgSize, treeErrorResponse)
	}

	serializedTree, err := proto.Marshal(tree)
	if err != nil {
		return sendWithin(stream, treeErrorResponse(boundedError(
			fmt.Sprintf("an error occurred during serialization of the tree: %v", err), s.maxMsgSize)),
			s.maxMsgSize, treeErrorResponse)
	}

	for _, chunk := range chunkifyBySize(serializedTree, s.chunkSize) {
		response := &a2l.TreeResponse{SerializedTreeChunk: chunk}
		if err := sendWithin(stream, response, s.maxMsgSize, treeErrorResponse); err != nil {
			return err
		}
	}

	return nil
}

func (s *grpcA2LImplType) GetJSONFromTree(stream a2l.A2L_GetJSONFromTreeServer) (err error) {
	var rawData []byte
	var buffer bytes.Buffer
	var chunk []byte
	var parseError error
	var request *a2l.JSONFromTreeRequest
	tree := &a2l.RootNodeType{}
	response := &a2l.JSONResponse{}
	indent := ""
	allowPartial := false
	emitUnpopulated := false
	// Note: optionsParsed := false avoid to parse option for each chunk
	optionsParsed := false

	for {
		request, err = stream.Recv()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}

		if !optionsParsed {
			if request.Indent != nil {
				for i := uint32(0); i < *request.Indent; i++ {
					indent += " "
				}
			}
			if request.AllowPartial != nil {
				allowPartial = *request.AllowPartial
			}
			if request.EmitUnpopulated != nil {
				emitUnpopulated = *request.EmitUnpopulated
			}
			optionsParsed = true
		}
		// Use a buffer to receive all chunks
		buffer.Write(request.Tree)
	}
	if err == nil {
		if parseError = proto.Unmarshal(buffer.Bytes(), tree); parseError == nil {
			opt := protojson.MarshalOptions{
				AllowPartial:    allowPartial,
				EmitUnpopulated: emitUnpopulated}

			if rawData, parseError = opt.Marshal(tree); parseError == nil {
				// Note: see https://github.com/golang/protobuf/issues/1121
				var indentedBuffer bytes.Buffer
				if err = json.Indent(&indentedBuffer, rawData, "", indent); err == nil {
					rawData = indentedBuffer.Bytes()
					for _, chunk = range chunkifyBySize(rawData, s.chunkSize) {
						response.Json = chunk
						if err = sendWithin(stream, response, s.maxMsgSize, jsonErrorResponse); err != nil {
							break
						}
					}
				} else {
					response.Error = proto.String(boundedError(
						fmt.Sprintf("an error occurred during json indent: %v", err), s.maxMsgSize))
					err = sendWithin(stream, response, s.maxMsgSize, jsonErrorResponse)
				}
			} else {
				response.Error = proto.String(boundedError(parseError.Error(), s.maxMsgSize))
				err = sendWithin(stream, response, s.maxMsgSize, jsonErrorResponse)
			}
		} else {
			response.Error = proto.String(boundedError(parseError.Error(), s.maxMsgSize))
			err = sendWithin(stream, response, s.maxMsgSize, jsonErrorResponse)
		}
	}

	return err
}

func (s *grpcA2LImplType) GetTreeFromJSON(stream a2l.A2L_GetTreeFromJSONServer) (err error) {
	var parseError error
	var buffer bytes.Buffer
	var serializedTree []byte
	var chunk []byte
	var request *a2l.TreeFromJSONRequest
	tree := &a2l.RootNodeType{}
	response := &a2l.TreeResponse{}
	allowPartial := false
	// Note: optionsParsed := false avoid to parse option for each chunk
	optionsParsed := false

	for {
		request, err = stream.Recv()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
		if !optionsParsed {
			if request.AllowPartial != nil {
				allowPartial = *request.AllowPartial
			}
			optionsParsed = true
		}
		// Use a buffer to receive all chunks
		buffer.Write(request.Json)
	}
	if err == nil {
		opt := protojson.UnmarshalOptions{
			AllowPartial: allowPartial,
		}

		if parseError = opt.Unmarshal(buffer.Bytes(), tree); parseError == nil {
			if serializedTree, err = proto.Marshal(tree); err == nil {
				for _, chunk = range chunkifyBySize(serializedTree, s.chunkSize) {
					response.SerializedTreeChunk = chunk
					if err = sendWithin(stream, response, s.maxMsgSize, treeErrorResponse); err != nil {
						break
					}
				}
			} else {
				response.Error = proto.String(boundedError(
					fmt.Sprintf("an error occurred during serialization of the tree: %v", err), s.maxMsgSize))
				err = sendWithin(stream, response, s.maxMsgSize, treeErrorResponse)
			}
		} else {
			response.Error = proto.String(boundedError(parseError.Error(), s.maxMsgSize))
			err = sendWithin(stream, response, s.maxMsgSize, treeErrorResponse)
		}
	}

	return err
}

func (s *grpcA2LImplType) GetA2LFromTree(stream a2l.A2L_GetA2LFromTreeServer) (err error) {
	var buffer bytes.Buffer
	var request *a2l.A2LFromTreeRequest
	var chunk []byte
	var a2lDataBytes []byte
	tree := &a2l.RootNodeType{}
	response := &a2l.A2LResponse{}
	indent := ""
	sorted := false
	// Note: optionsParsed := false avoid to parse option for each chunk
	optionsParsed := false

	for {
		request, err = stream.Recv()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
		if !optionsParsed {
			if request.Indent != nil {
				for i := uint32(0); i < *request.Indent; i++ {
					indent += " "
				}
			}
			if request.Sorted != nil {
				sorted = *request.Sorted
			}
			optionsParsed = true
		}
		// Use a buffer to receive all chunks
		buffer.Write(request.Tree)
	}
	if err == nil {
		var marshalError error

		if marshalError = proto.Unmarshal(buffer.Bytes(), tree); marshalError == nil {
			a2lDataBytes, marshalError = marshalA2L(tree, indent, sorted)
		}

		if marshalError == nil {
			for _, chunk = range chunkifyBySize(a2lDataBytes, s.chunkSize) {
				response.A2L = chunk
				if err = sendWithin(stream, response, s.maxMsgSize, a2lErrorResponse); err != nil {
					break
				}
			}
		} else {
			response.Error = proto.String(boundedError(marshalError.Error(), s.maxMsgSize))
			err = sendWithin(stream, response, s.maxMsgSize, a2lErrorResponse)
		}
	}

	return err
}

// marshalA2L serializes the tree back to A2L. A tree which lacks an element the parser always
// fills, but which a client building a tree by hand may leave out, makes the marshaller
// dereference a nil node. The panic is turned into an error here so that the caller is told what
// happened through the error field of the response, like every other failure of the API.
func marshalA2L(tree *a2l.RootNodeType, indent string, sorted bool) (result []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			result, err = nil, fmt.Errorf("an error occurred during serialization of the tree to A2L: %v", r)
		}
	}()

	return []byte(tree.MarshalA2L(0, indent, sorted)), nil
}

//export GetJSONByteArrayFromA2LByteArray
func GetJSONByteArrayFromA2LByteArray(a2lByteArray []byte) {
	_ = a2lByteArray
}

func recoverUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = status.Errorf(codes.Internal, "%v: %v", info.FullMethod, r)
		}
	}()

	return handler(ctx, req)
}

func recoverStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = status.Errorf(codes.Internal, "%v: %v", info.FullMethod, r)
		}
	}()

	return handler(srv, ss)
}

var serverMutex sync.Mutex

var server *grpc.Server

// serverPort is the port the running server listens on, 0 while no server is running.
var serverPort int

// createServer starts the gRPC server on the passed port and returns 0 on success, 1 on failure.
// It fails when a server is already running, when maxMsgSize leaves no room for a chunk of tree,
// and when the port cannot be bound. A port of 0 lets the operating system choose a free one,
// which getPort reports; that is how several processes run a server each without colliding.
//
// The server is bound to the loopback interface: it is an unauthenticated endpoint which serves
// the process which loaded this library, not the network it sits on.
//
// The cgo types stay in Create, so that this function can be exercised by the tests, which cannot
// import "C".
func createServer(port int, maxMsgSize int) int {
	serverMutex.Lock()
	defer serverMutex.Unlock()

	if server != nil {
		return 1
	}

	// the responses carry the serialized tree in chunks of maxMsgSize - protocolSizeMargin bytes,
	// which must leave at least one byte of payload
	if maxMsgSize <= protocolSizeMargin {
		return 1
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%v", port))
	if err != nil {
		return 1
	}

	newServer := grpc.NewServer(
		grpc.MaxRecvMsgSize(maxMsgSize),
		grpc.MaxSendMsgSize(maxMsgSize),
		grpc.ChainUnaryInterceptor(recoverUnaryInterceptor),
		grpc.ChainStreamInterceptor(recoverStreamInterceptor))

	a2l.RegisterA2LServer(newServer, newGrpcA2LImpl(maxMsgSize))

	server = newServer
	serverPort = listener.Addr().(*net.TCPAddr).Port

	go func() {
		// Serve returns when closeServer stops the server; there is nobody left to report to
		_ = newServer.Serve(listener)
	}()

	return 0
}

// closeServer stops the running server and returns 0, or 1 when no server is running.
func closeServer() int {
	serverMutex.Lock()
	defer serverMutex.Unlock()

	if server == nil {
		return 1
	}

	server.Stop()
	server = nil
	serverPort = 0

	return 0
}

// getPort returns the port the running server listens on, or 0 when no server is running.
func getPort() int {
	serverMutex.Lock()
	defer serverMutex.Unlock()

	return serverPort
}

//export Create
func Create(port C.int, maxMsgSize C.int) (result C.int) {
	return C.int(createServer(int(port), int(maxMsgSize)))
}

//export Close
func Close() (result C.int) {
	return C.int(closeServer())
}

// GetPort returns the TCP port the server listens on, or 0 when no server is running. It is the
// way to learn the port when Create was given 0 and the operating system chose one.
//
//export GetPort
func GetPort() (result C.int) {
	return C.int(getPort())
}

func main() {
	if createServer(3333, 4*1024*1024) != 0 {
		fmt.Println("unable to start the gRPC server")

		return
	}

	select {}
}
