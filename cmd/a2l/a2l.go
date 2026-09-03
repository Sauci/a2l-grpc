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
	"google.golang.org/protobuf/proto"
	"io"
	"net"
	"sync"
)

const protocolSizeMargin = 256

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

type grpcA2LImplType struct {
	a2l.UnimplementedA2LServer
	chunkSize int
}

func (s *grpcA2LImplType) GetTreeFromA2L(stream a2l.A2L_GetTreeFromA2LServer) error {
	var buffer bytes.Buffer
	var parseError error
	var err error
	var serializedTree []byte
	var chunk []byte
	var request *a2l.TreeFromA2LRequest
	tree := &a2l.RootNodeType{}
	response := &a2l.TreeResponse{}
	options := a2l.ParseOptions{}
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
			if request.EnforceVersionCheck != nil {
				options.EnforceVersionCheck = *request.EnforceVersionCheck
			}
			optionsParsed = true
		}
		buffer.Write(request.A2L)
	}

	if err == nil {
		var warnings []a2l.SyntaxError

		tree, warnings, parseError = a2l.GetTreeFromStringWithOptions(buffer.String(), options)

		// warnings are attached to the first response of the stream
		for _, warning := range warnings {
			response.Warnings = append(response.Warnings, warning.String())
		}

		if parseError == nil {
			if serializedTree, err = proto.Marshal(tree); err == nil {
				for _, chunk = range chunkifyBySize(serializedTree, s.chunkSize) {
					response.SerializedTreeChunk = chunk
					if err = stream.Send(response); err != nil {
						break
					}
					response.Warnings = nil
				}
			} else {
				response.Error = proto.String(fmt.Sprintf("An error occured during serialization of Tree: %v", err))
				err = stream.Send(response)
			}
		} else {
			errString := parseError.Error()
			response.Error = &errString
			err = stream.Send(response)
		}
	}

	return err
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
						if err = stream.Send(response); err != nil {
							break
						}
					}
				} else {
					response.Error = proto.String(fmt.Sprintf("An error occured during json indent: %v", err))
					err = stream.Send(response)
				}
			} else {
				errString := parseError.Error()
				response.Error = &errString
				err = stream.Send(response)
			}
		} else {
			errString := parseError.Error()
			response.Error = &errString
			err = stream.Send(response)
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
					if err = stream.Send(response); err != nil {
						break
					}
				}
			} else {
				response.Error = proto.String(fmt.Sprintf("An error occured during serialization of Tree: %v", err))
				err = stream.Send(response)
			}
		} else {
			errString := parseError.Error()
			response.Error = &errString
			err = stream.Send(response)
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
				if err = stream.Send(response); err != nil {
					break
				}
			}
		} else {
			response.Error = proto.String(marshalError.Error())
			err = stream.Send(response)
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

// createServer starts the gRPC server on the passed port and returns 0 on success, 1 on failure.
// It fails when a server is already running, when maxMsgSize leaves no room for a chunk of tree,
// and when the port cannot be bound.
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

	a2l.RegisterA2LServer(newServer, &grpcA2LImplType{chunkSize: maxMsgSize - protocolSizeMargin})

	server = newServer

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

	return 0
}

//export Create
func Create(port C.int, maxMsgSize C.int) (result C.int) {
	return C.int(createServer(int(port), int(maxMsgSize)))
}

//export Close
func Close() (result C.int) {
	return C.int(closeServer())
}

func main() {
	if createServer(3333, 4*1024*1024) != 0 {
		fmt.Println("unable to start the gRPC server")

		return
	}

	select {}
}
