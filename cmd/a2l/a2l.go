package main

import (
	"C"
	"context"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"google.golang.org/grpc"
	"os"

	"log"
	"net"
	"stash.lmb.liebherr.i/tel/gomodparser/pkg/a2l"
	"stash.lmb.liebherr.i/tel/gomodparser/pkg/a2l/parser"
)

func NewTreeFromString(a2lContent string) (tree *a2l.RootNodeType, err error) {
	input := antlr.NewInputStream(a2lContent)
	lexer := parser.NewA2LLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewA2LParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	listener := a2l.NewListener()

	antlr.ParseTreeWalkerDefault.Walk(listener, p.A2lFile())

	tree = listener.Tree()

	return tree, err
}

type grpcA2LImplType struct {
	a2l.UnimplementedA2LServer
}

func (s *grpcA2LImplType) GetTree(_ context.Context, request *a2l.GetTreeRequest) (result *a2l.RootNodeType, err error) {
	var tree *a2l.RootNodeType

	var file *os.File
	file, err = os.OpenFile("C:\\git\\gomodparser\\a2l_parser.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer func() { _ = file.Close() }()

	logger := log.New(file, "C:\\git\\gomodparser\\a2l_parser.log", log.LstdFlags)

	logger.Printf("GetTree request: %s", request.String())

	if tree, err = NewTreeFromString(request.GetA2L()); err == nil {
		result = tree
	} else {
		logger.Print(err)
	}

	return result, err
}

//export NewServer
func NewServer(port C.int) {
	var err error
	var lis net.Listener
	var file *os.File

	file, err = os.OpenFile("C:\\git\\gomodparser\\a2l_parser.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer func() { _ = file.Close() }()

	logger := log.New(file, "C:\\git\\gomodparser\\a2l_parser.log", log.LstdFlags)

	logger.Printf("start listening on port %v", port)

	if lis, err = net.Listen("tcp", fmt.Sprintf(":%v", port)); err == nil {

		s := grpc.NewServer(grpc.MaxRecvMsgSize(100*1024*1024), grpc.MaxSendMsgSize(100*1024*1024))

		a2l.RegisterA2LServer(s, &grpcA2LImplType{})

		go func() {
			if err = s.Serve(lis); err != nil {
				logger.Print(err)
			}
		}()
	} else {
		logger.Print(err)
	}
}

func main() {}
