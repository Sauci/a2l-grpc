package main

import (
	"C"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/sauci/a2l-grpc/pkg/a2l"
	"github.com/sauci/a2l-grpc/pkg/a2l/parser"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"net"
	"strings"
	"sync"
)

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

//export GetTreeFromA2LLocal
func GetTreeFromA2LLocal(request []byte) (result *a2l.RootNodeType, errString *string) {
	if tree, parseError := getTreeFromString(string(request)); parseError == nil {
		result = tree
	} else {
		errStringValue := parseError.Error()
		errString = &errStringValue
	}

	return result, errString
}

func (s *grpcA2LImplType) GetTreeFromA2L(_ context.Context, request *a2l.TreeFromA2LRequest) (result *a2l.TreeResponse, err error) {
	var treeResult *a2l.RootNodeType
	var errString *string

	treeResult, errString = GetTreeFromA2LLocal(request.A2L)

	return &a2l.TreeResponse{Tree: treeResult, Error: errString}, err
}

//export GetJSONFromTreeLocal
func GetJSONFromTreeLocal(request *a2l.RootNodeType, indent *uint32, allowPartial *bool, emitUnpopulated *bool) (result []byte, errString *string) {
	var rawData []byte
	var indentedData []byte
	var parseError error
	indentValue := ""
	allowPartialValue := false
	emitUnpopulatedValue := false

	if indent != nil {
		for i := uint32(0); i < *indent; i++ {
			indentValue += " "
		}
	}

	if allowPartial != nil {
		allowPartialValue = *allowPartial
	}

	if emitUnpopulated != nil {
		emitUnpopulatedValue = *emitUnpopulated
	}

	opt := protojson.MarshalOptions{
		AllowPartial:    allowPartialValue,
		EmitUnpopulated: emitUnpopulatedValue}

	if rawData, parseError = opt.Marshal(request); parseError == nil {
		// Note: see https://github.com/golang/protobuf/issues/1121
		buffer := bytes.NewBuffer(indentedData)
		if err := json.Indent(buffer, rawData, "", indentValue); err == nil {
			result = buffer.Bytes()
		} else {
			errStringValue := err.Error()
			errString = &errStringValue
		}
	} else {
		errStringValue := parseError.Error()
		errString = &errStringValue
	}

	return result, errString
}

func (s *grpcA2LImplType) GetJSONFromTree(_ context.Context, request *a2l.JSONFromTreeRequest) (result *a2l.JSONResponse, err error) {
	var byteResult []byte
	var errString *string

	byteResult, errString = GetJSONFromTreeLocal(request.Tree, request.Indent, request.AllowPartial, request.EmitUnpopulated)

	return &a2l.JSONResponse{Json: byteResult, Error: errString}, err
}

//export GetTreeFromJSONLocal
func GetTreeFromJSONLocal(request []byte, allowPartial *bool) (result *a2l.RootNodeType, errString *string) {
	var parseError error
	allowPartialValue := false

	if allowPartial != nil {
		allowPartialValue = *allowPartial
	}

	opt := protojson.UnmarshalOptions{
		AllowPartial: allowPartialValue,
	}

	if parseError = opt.Unmarshal(request, result); parseError != nil {
		errStringValue := parseError.Error()
		errString = &errStringValue
	}

	return result, errString
}

func (s *grpcA2LImplType) GetTreeFromJSON(_ context.Context, request *a2l.TreeFromJSONRequest) (result *a2l.TreeResponse, err error) {
	var treeResult *a2l.RootNodeType
	var errString *string

	treeResult, errString = GetTreeFromJSONLocal(request.Json, request.AllowPartial)

	return &a2l.TreeResponse{Tree: treeResult, Error: errString}, err
}

//export GetA2LFromTreeLocal
func GetA2LFromTreeLocal(request *a2l.RootNodeType, indent *uint32, sorted *bool) (result []byte, err error) {
	indentValue := ""
	sortedValue := false
	if indent != nil {
		for i := uint32(0); i < *indent; i++ {
			indentValue += " "
		}
	}

	if sorted != nil {
		sortedValue = *sorted
	}

	return []byte(request.MarshalA2L(0, indentValue, sortedValue)), nil
}

func (s *grpcA2LImplType) GetA2LFromTree(_ context.Context, request *a2l.A2LFromTreeRequest) (result *a2l.A2LResponse, err error) {
	var byteResult []byte

	byteResult, err = GetA2LFromTreeLocal(request.Tree, request.Indent, request.Sorted)

	return &a2l.A2LResponse{A2L: byteResult}, err
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
			server = grpc.NewServer(grpc.MaxRecvMsgSize(200*1024*1024), grpc.MaxSendMsgSize(200*1024*1024))

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
