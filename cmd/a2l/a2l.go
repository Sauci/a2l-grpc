package main

import (
	"C"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/sauci/a2l-grpc/pkg/a2l"
	"github.com/sauci/a2l-grpc/pkg/a2l/parser"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io"
	"net"
	"strings"
	"sync"
)

// It's the size that we set for one chunk
var protocolSizeMargin = 256
var sizeOfStreamMsg = 4*1024*1024 - protocolSizeMargin

func chunkifyBySize(data []byte, chunkSize int) [][]byte {
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

type A2LSyntaxError struct {
	line, column int
	msg          string
}

func (e A2LSyntaxError) String() string {
	return fmt.Sprintf("%v:%v %v", e.line, e.column, e.msg)
}

type A2LErrorListener struct {
	*antlr.DefaultErrorListener // Embed default which ensures we fit the interface
	Errors                      []A2LSyntaxError
}

func (l *A2LErrorListener) GetErrors() (result []string) {
	result = make([]string, 0)

	for _, e := range l.Errors {
		result = append(result, e.String())
	}

	return result
}

func (l *A2LErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	l.Errors = append(l.Errors, A2LSyntaxError{line: line, column: column, msg: msg})
}

func getTreeFromString(a2lString string) (result *a2l.RootNodeType, err error) {
	errorListener := A2LErrorListener{Errors: make([]A2LSyntaxError, 0)}
	lexer := parser.NewA2LLexer(antlr.NewInputStream(a2lString))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(&errorListener)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewA2LParser(tokenStream)
	p.RemoveErrorListeners()
	p.AddErrorListener(&errorListener)
	p.BuildParseTrees = true

	listener := a2l.NewListener()

	antlr.ParseTreeWalkerDefault.Walk(listener, p.A2lFile())

	if len(errorListener.Errors) != 0 {
		err = fmt.Errorf(strings.Join(errorListener.GetErrors(), "\n"))
	}

	return listener.Tree(), err
}

type grpcA2LImplType struct {
	a2l.UnimplementedA2LServer
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

	for {
		request, err = stream.Recv()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
		buffer.Write(request.A2L)
	}

	if err == nil {
		if tree, parseError = getTreeFromString(buffer.String()); parseError == nil {
			if serializedTree, err = proto.Marshal(tree); err == nil {
				for _, chunk = range chunkifyBySize(serializedTree, sizeOfStreamMsg) {
					response.SerializedTreeChunk = chunk
					err = stream.Send(response)
				}
			} else {
				response.Error = proto.String(fmt.Sprintf("An error occure during serialization of Tree: %v", err))
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
					for _, chunk = range chunkifyBySize(rawData, sizeOfStreamMsg) {
						response.Json = chunk
						err = stream.Send(response)
					}
				} else {
					response.Error = proto.String(fmt.Sprintf("An error occure during json indent: %v", err))
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
				for _, chunk = range chunkifyBySize(serializedTree, sizeOfStreamMsg) {
					response.SerializedTreeChunk = chunk
					err = stream.Send(response)
				}
			} else {
				response.Error = proto.String(fmt.Sprintf("An error occure during serialization of Tree: %v", err))
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
		if err = proto.Unmarshal(buffer.Bytes(), tree); err == nil {
			a2lDataBytes = []byte(tree.MarshalA2L(0, indent, sorted))
			for _, chunk = range chunkifyBySize(a2lDataBytes, sizeOfStreamMsg) {
				response.A2L = chunk
				err = stream.Send(response)
			}
		}
	}

	return err
}

var serverMutex sync.Mutex

var server *grpc.Server

//export Create
func Create(port C.int) (result C.int) {
	var err error
	var listener net.Listener

	serverMutex.Lock()
	defer serverMutex.Unlock()

	if server != nil {
		result = 1
	} else {
		result = 0
	}

	if result == 0 {
		if listener, err = net.Listen("tcp", fmt.Sprintf(":%v", port)); err == nil {
			server = grpc.NewServer(grpc.MaxRecvMsgSize(sizeOfStreamMsg+protocolSizeMargin), grpc.MaxSendMsgSize(sizeOfStreamMsg+protocolSizeMargin))

			a2l.RegisterA2LServer(server, &grpcA2LImplType{})

			go func() {
				err = server.Serve(listener)
			}()
		}
	}

	return result
}

//export Close
func Close() (result C.int) {
	serverMutex.Lock()
	defer serverMutex.Unlock()

	if server == nil {
		result = 1
	} else {
		server.Stop()

		server = nil

		result = 0
	}

	return result
}

func main() {
	Create(3333)

	for {
		select {}
	}
}
