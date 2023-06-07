package main

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/sauci/a2l-grpc/pkg/a2l"
	"github.com/sauci/a2l-grpc/pkg/a2l/parser"
	"log"
	"os"
	"testing"
)

func Test_Main(t *testing.T) {
	var err error
	var input *antlr.FileStream

	if input, err = antlr.NewFileStream("C:\\git\\gomodparser\\ECU3D06_D1MC_241700_339747_6E0060_RTM.a2l"); err == nil {
		lexer := parser.NewA2LLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewA2LParser(stream)
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
		p.BuildParseTrees = true

		listener := a2l.NewListener()

		antlr.ParseTreeWalkerDefault.Walk(listener, p.A2lFile())

		tree := listener.Tree()

		//if result, err = tree.MarshalA2L(); err == nil {
		//	if err = os.WriteFile("C:\\git\\gomodparser\\output.a2l", result, 0666); err == nil {
		if err = os.WriteFile("C:\\git\\gomodparser\\output.a2l", []byte(tree.MarshalA2L()), 0666); err != nil {
			log.Print(err)
		}
	} else {
		log.Print(err)
	}
	//} else {
	//	log.Print(err)
	//}
	//} else {
	//  log.Print(err)
}
