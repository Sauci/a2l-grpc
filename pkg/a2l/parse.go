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

// maxQuotedTokens is the number of tokens a syntax error quotes of the input which precedes it.
// Eight reaches the beginning of the element which failed together with the parameters it did
// accept, which is what names the mistake; more than that only adds the elements before it, which
// are already closed and have nothing to do with the failure.
const maxQuotedTokens = 8

// boundedTokenStream quotes a bounded part of the input in a syntax error.
//
// ANTLR reports a dead end with "no viable alternative at input" followed by the whole text between
// the token where the failed decision started and the offending one. The decisions of this grammar
// are the loops of optional elements, which start at the beginning of a block, so for an error in
// the last MODULE of a file that text is nearly the entire file: a single message of a 10 MB file
// measured 9.3 million characters. Building it is quadratic in the size of the file, since the
// parser reports again as it recovers, which made an error in a large file take longer than the
// parse itself by two orders of magnitude, and made an error in a 100 MB file unbounded in
// practice.
//
// The parser reaches the stream through the TokenStream interface, so bounding the text here is
// enough, and it needs neither a fork of the runtime nor an error strategy of our own: the fields
// the strategy would have to reach are not exported. Quoting the last few tokens before the error
// is also the more useful message.
type boundedTokenStream struct {
	*antlr.CommonTokenStream
}

func (s *boundedTokenStream) GetTextFromTokens(start, end antlr.Token) string {
	if start == nil || end == nil {
		return s.CommonTokenStream.GetTextFromTokens(start, end)
	}

	if end.GetTokenIndex()-start.GetTokenIndex() <= maxQuotedTokens {
		return s.CommonTokenStream.GetTextFromTokens(start, end)
	}

	return "..." + s.CommonTokenStream.GetTextFromTokens(s.Get(end.GetTokenIndex()-maxQuotedTokens), end)
}

// sllBailOut abandons the first stage. The BailErrorStrategy of the Go runtime does not stop the
// parse the way the one of the Java runtime does: it marks the state of the parser, which the
// generated errorExit reports and then clears, so the parse carries on. A file the first stage
// cannot parse would thus be walked to its end on the error path, which is far slower than the
// parse itself, and only then be parsed again. Leaving the parse by a panic is what actually
// abandons it; parseFast recovers it at once and throws the parser away, so nothing ever observes
// the state it was interrupted in.
type sllBailOut struct{}

// sllErrorListener abandons the first stage at the first error it is told about.
type sllErrorListener struct {
	*antlr.DefaultErrorListener
}

func (l *sllErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, _, _ int, _ string, _ antlr.RecognitionException) {
	panic(sllBailOut{})
}

// parseFast is the first stage: SLL prediction, no error recovery, abandoned at the first error.
func parseFast(tokenStream *antlr.CommonTokenStream) (fileTree parser.IA2lFileContext, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			if _, bail := r.(sllBailOut); !bail {
				panic(r)
			}

			fileTree, ok = nil, false
		}
	}()

	p := parser.NewA2LParser(&boundedTokenStream{CommonTokenStream: tokenStream})
	p.RemoveErrorListeners()
	p.AddErrorListener(&sllErrorListener{DefaultErrorListener: antlr.NewDefaultErrorListener()})
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	p.SetErrorHandler(antlr.NewBailErrorStrategy())
	p.BuildParseTrees = true

	return p.A2lFile(), true
}

// parseA2lFile parses the token stream in the two stages ANTLR recommends. The first one predicts
// with SLL, which is many times faster than full LL but may report an error on a valid file, for
// lack of the full context; it bails out at the first error instead of recovering. Only when it
// reported anything is the file parsed again, with full LL prediction and error recovery, and that
// second parse is the one whose errors reach the caller. A valid file is thus parsed once, in the
// fast mode, and an invalid one twice, with exactly the errors the slow mode reports on its own;
// the errors of the lexer are recorded once either way, since the tokens are read once.
func parseA2lFile(tokenStream *antlr.CommonTokenStream, errorListener antlr.ErrorListener) parser.IA2lFileContext {
	if fileTree, ok := parseFast(tokenStream); ok {
		return fileTree
	}

	tokenStream.Seek(0)

	second := parser.NewA2LParser(&boundedTokenStream{CommonTokenStream: tokenStream})
	second.RemoveErrorListeners()
	second.AddErrorListener(errorListener)
	second.BuildParseTrees = true

	return second.A2lFile()
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

	listener := NewListener()

	fileTree := parseA2lFile(tokenStream, &errorListener)

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
