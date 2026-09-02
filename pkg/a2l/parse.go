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

// ParseOptions controls the behaviour of the parser.
type ParseOptions struct {
	// EnforceVersionCheck turns the version warnings into errors: a file declaring a version with
	// ASAP2_VERSION and containing keywords which require a newer version of the standard is then
	// rejected instead of being parsed with warnings.
	EnforceVersionCheck bool
}

// GetTreeFromString parses the passed A2L content and returns the corresponding tree. Keywords
// requiring a newer version of the standard than the one declared with ASAP2_VERSION are
// tolerated; use GetTreeFromStringWithOptions to retrieve the corresponding warnings or to turn
// them into errors.
func GetTreeFromString(a2lString string) (result *RootNodeType, err error) {
	result, _, err = GetTreeFromStringWithOptions(a2lString, ParseOptions{})

	return result, err
}

// GetTreeFromStringWithOptions parses the passed A2L content with the passed options and returns
// the corresponding tree. Every construct requiring a newer version of the standard than the one
// the file declares with ASAP2_VERSION is reported in warnings; with
// ParseOptions.EnforceVersionCheck the warnings are returned as errors instead.
func GetTreeFromStringWithOptions(a2lString string, options ParseOptions) (result *RootNodeType, warnings []SyntaxError, err error) {
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

	// The byte-order mark declares the encoding of the file (ASAM MCD-2 MC 1.6.1, chapter 1.5.2).
	// It is not part of the content and the lexer has no rule for it, so it is dropped here.
	a2lString = strings.TrimPrefix(a2lString, "\uFEFF")

	lexer := parser.NewA2LLexer(antlr.NewInputStream(a2lString))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(&errorListener)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewA2LParser(tokenStream)
	p.RemoveErrorListeners()
	p.AddErrorListener(&errorListener)
	p.BuildParseTrees = true

	listener := NewListener()

	fileTree := p.A2lFile()

	versionCheck := newVersionCheckListener()
	antlr.ParseTreeWalkerDefault.Walk(versionCheck, fileTree)

	if options.EnforceVersionCheck {
		errorListener.Errors = append(errorListener.Errors, versionCheck.warnings...)
	} else {
		warnings = versionCheck.warnings
	}

	antlr.ParseTreeWalkerDefault.Walk(listener, fileTree)

	if len(errorListener.Errors) != 0 {
		err = fmt.Errorf(strings.Join(errorListener.GetErrors(), "\n"))
	}

	return listener.Tree(), warnings, err
}
