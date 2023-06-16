package main

import (
	"C"
	"context"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/sauci/a2l-grpc/pkg/a2l"
	"github.com/sauci/a2l-grpc/pkg/a2l/parser"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"io"
	"net"
	"os"
	"sync"
)

func getTreeFromString(a2lString string) (result *a2l.RootNodeType, err error) {
	stdErr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	lexer := parser.NewA2LLexer(antlr.NewInputStream(a2lString))
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewA2LParser(tokenStream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	listener := a2l.NewListener()

	antlr.ParseTreeWalkerDefault.Walk(listener, p.A2lFile())

	_ = w.Close()
	out, _ := io.ReadAll(r)

	os.Stderr = stdErr

	if len(out) != 0 {
		err = fmt.Errorf("%s", string(out))
	}

	return listener.Tree(), err
}

type grpcA2LImplType struct {
	a2l.UnimplementedA2LServer
}

func (s *grpcA2LImplType) GetTreeFromA2L(_ context.Context, request *a2l.A2LRequest) (result *a2l.TreeResponse, err error) {
	var tree *a2l.RootNodeType
	var parseError error

	result = &a2l.TreeResponse{}

	if tree, parseError = getTreeFromString(request.GetA2L()); parseError == nil {
		result.Tree = tree
	} else {
		errString := parseError.Error()
		result.Error = &errString
	}

	return result, err
}

func (s *grpcA2LImplType) GetJSONFromTree(_ context.Context, request *a2l.TreeRequest) (result *a2l.JSONResponse, err error) {
	var rawData []byte
	var parseError error

	result = &a2l.JSONResponse{}

	if rawData, parseError = protojson.Marshal(request.Tree); parseError == nil {
		result.Json = string(rawData)
	} else {
		errString := parseError.Error()
		result.Error = &errString
	}

	return result, err
}

func (s *grpcA2LImplType) GetTreeFromJSON(_ context.Context, request *a2l.JSONRequest) (result *a2l.TreeResponse, err error) {
	var parseError error

	result = &a2l.TreeResponse{Tree: &a2l.RootNodeType{}}

	if parseError = protojson.Unmarshal([]byte(request.Json), result.Tree); parseError != nil {
		errString := parseError.Error()
		result.Error = &errString
	}

	return result, err
}

func (s *grpcA2LImplType) GetA2LFromTree(_ context.Context, request *a2l.TreeRequest) (result *a2l.A2LResponse, err error) {
	panic("not implemented yet...")
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
		server.GracefulStop()

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
