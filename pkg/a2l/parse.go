package a2l

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/sauci/a2l-grpc/pkg/a2l/parser"
	"strings"
)

type SyntaxError struct {
	Line, Column int
	Msg          string
}

func (e SyntaxError) String() string {
	return fmt.Sprintf("%v:%v %v", e.Line, e.Column, e.Msg)
}

type ErrorListener struct {
	*antlr.DefaultErrorListener // Embed default which ensures we fit the interface
	Errors                      []SyntaxError
}

func (l *ErrorListener) GetErrors() (result []string) {
	result = make([]string, 0)

	for _, e := range l.Errors {
		result = append(result, e.String())
	}

	return result
}

func (l *ErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	l.Errors = append(l.Errors, SyntaxError{Line: line, Column: column, Msg: msg})
}

// GetTreeFromString parses the passed A2L content and returns the corresponding tree.
func GetTreeFromString(a2lString string) (result *RootNodeType, err error) {
	errorListener := ErrorListener{Errors: make([]SyntaxError, 0)}

	defer func() {
		if r := recover(); r != nil {
			result = nil
			if len(errorListener.Errors) != 0 {
				err = fmt.Errorf("%v\n%v", strings.Join(errorListener.GetErrors(), "\n"), r)
			} else {
				err = fmt.Errorf("error while building A2L tree: %v", r)
			}
		}
	}()

	lexer := parser.NewA2LLexer(antlr.NewInputStream(a2lString))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(&errorListener)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewA2LParser(tokenStream)
	p.RemoveErrorListeners()
	p.AddErrorListener(&errorListener)
	p.BuildParseTrees = true

	listener := NewListener()

	antlr.ParseTreeWalkerDefault.Walk(listener, p.A2lFile())

	if len(errorListener.Errors) != 0 {
		err = fmt.Errorf(strings.Join(errorListener.GetErrors(), "\n"))
	}

	return listener.Tree(), err
}
