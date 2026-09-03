package a2l

import (
	"errors"
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

// unresolvedIncludes reports every /include statement of the file. The include mechanism of ASAM
// MCD-2 MC 1.6.1 chapter 4 is a text replacement, which needs the file system the file was read
// from; this parser is given the content of a single file, so it can only name the references it
// cannot resolve. Reporting them keeps a tree which silently lacks the content of the included
// files from being taken for a complete one.
func unresolvedIncludes(tokenStream *antlr.CommonTokenStream) (result []SyntaxError) {
	tokenStream.Fill()

	for _, token := range tokenStream.GetAllTokens() {
		if token.GetTokenType() != parser.A2LLexerINCLUDE {
			continue
		}

		result = append(result, SyntaxError{
			Line:   token.GetLine(),
			Column: token.GetColumn(),
			Msg: fmt.Sprintf("%s is not resolved by this parser, the content of the included "+
				"file must be substituted by the caller", strings.TrimSpace(token.GetText())),
		})
	}

	return result
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

	errorListener.Errors = append(errorListener.Errors, unresolvedIncludes(tokenStream)...)

	versionCheck := newVersionCheckListener()
	antlr.ParseTreeWalkerDefault.Walk(versionCheck, fileTree)

	// Violations which do not depend on the declared version are always reported, the enforcement
	// option only controls the version warnings.
	errorListener.Errors = append(errorListener.Errors, versionCheck.errors...)

	// ASAP2_VERSION is mandatory since ASAM MCD-2 MC 1.6.1 (chapters 1.4.4 and 3.5.16). Without it
	// no construct can be gated, so a caller which asked for the check to be enforced is told that
	// nothing could be verified instead of taking the empty result for a clean bill of health.
	if options.EnforceVersionCheck && versionCheck.version == nil {
		versionCheck.warnings = append(versionCheck.warnings, SyntaxError{
			Line:   1,
			Column: 0,
			Msg:    "ASAP2_VERSION is missing, the version check cannot be performed",
		})
	}

	if options.EnforceVersionCheck {
		errorListener.Errors = append(errorListener.Errors, versionCheck.warnings...)
	} else {
		warnings = versionCheck.warnings
	}

	antlr.ParseTreeWalkerDefault.Walk(listener, fileTree)

	if len(errorListener.Errors) != 0 {
		// the messages quote the offending part of the file, so they must never be used as a
		// format string: a per cent sign in the parsed content would be taken for a verb
		err = errors.New(strings.Join(errorListener.GetErrors(), "\n"))
	}

	return listener.Tree(), warnings, err
}
