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
	"net"
	"sync"
)

func getTreeFromString(a2lString string) (result *a2l.RootNodeType, err error) {
	lexer := parser.NewA2LLexer(antlr.NewInputStream(a2lString))
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewA2LParser(tokenStream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	listener := a2l.NewListener()

	antlr.ParseTreeWalkerDefault.Walk(listener, p.A2lFile())

	return listener.Tree(), err
}

type grpcA2LImplType struct {
	a2l.UnimplementedA2LServer
}

func (s *grpcA2LImplType) GetTree(_ context.Context, request *a2l.GetTreeRequest) (result *a2l.GetTreeResponse, err error) {
	var tree *a2l.RootNodeType
	var internalErr error

	result = &a2l.GetTreeResponse{}

	if tree, internalErr = getTreeFromString(request.GetA2L()); internalErr == nil {
		result.Tree = tree
	} else {
		errString := internalErr.Error()
		result.Error = &errString
	}

	return result, err
}

func (s *grpcA2LImplType) GetJSONTree(_ context.Context, request *a2l.RootNodeType) (result *a2l.GetJSONTreeResponse, err error) {
	var rawData []byte
	var internalErr error

	result = &a2l.GetJSONTreeResponse{}

	if rawData, internalErr = protojson.Marshal(request); internalErr == nil {
		result.Json = string(rawData)
	} else {
		errString := internalErr.Error()
		result.Error = &errString
	}

	return result, err
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
			server = grpc.NewServer(grpc.MaxRecvMsgSize(100*1024*1024), grpc.MaxSendMsgSize(100*1024*1024))

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

func main() {}
