package a2l

// Version gating of the parsed content.
//
// The grammar accepts a superset of ASAP2 1.51: a number of keywords and enumeration values were
// introduced by later versions of the standard (ASAM MCD-2MC 1.60 and 1.70). When the parsed file
// declares its version with ASAP2_VERSION, the constructs which require a newer version than the
// declared one are reported as warnings, or as errors when the check is enforced with
// ParseOptions.EnforceVersionCheck. A file which does not declare a version is not gated, since
// there is nothing to check against; when the check is enforced, the missing ASAP2_VERSION is
// reported instead, so that an empty result is not mistaken for a verified file.

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/sauci/a2l-grpc/pkg/a2l/parser"
	"reflect"
	"strings"
)

type asap2Version struct {
	versionNo, upgradeNo int32
}

func (v asap2Version) atLeast(other asap2Version) bool {
	if v.versionNo != other.versionNo {
		return v.versionNo > other.versionNo
	}

	return v.upgradeNo >= other.upgradeNo
}

func (v asap2Version) String() string {
	return fmt.Sprintf("%d.%d", v.versionNo, v.upgradeNo)
}

var asap2Version160 = asap2Version{versionNo: 1, upgradeNo: 60}
var asap2Version161 = asap2Version{versionNo: 1, upgradeNo: 61}

type versionCheckListener struct {
	*parser.BaseA2LListener
	version  *asap2Version
	warnings []SyntaxError
	// errors holds the violations which do not depend on the declared version and which the
	// grammar cannot express on its own, e.g. an enumeration value the parser has to accept as an
	// identifier. They are reported unconditionally, not only when the check is enforced.
	errors []SyntaxError
	// cardinalities caches the analysis of the context structs, so that it is done once per kind
	// of node instead of once per node.
	cardinalities map[reflect.Type]*contextCardinality
}

func newVersionCheckListener() *versionCheckListener {
	return &versionCheckListener{cardinalities: map[reflect.Type]*contextCardinality{}}
}

func (l *versionCheckListener) warn(message string, at antlr.Token) {
	l.warnings = append(l.warnings, SyntaxError{
		Line:   at.GetLine(),
		Column: at.GetColumn(),
		Msg:    message,
	})
}

// contextCardinality describes how often each kind of child may appear below one kind of parse
// tree context. It is derived from the labels of the grammar: ANTLR turns "x = rule" into a single
// "IXContext" field of the context struct and "x += rule" into an additional "[]IXContext" field,
// so the labels are the declaration of the cardinality the specification prescribes.
type contextCardinality struct {
	allowed    map[reflect.Type]int
	repeatable map[reflect.Type]bool
	resolved   map[reflect.Type]reflect.Type
}

// isRuleContextInterface keeps the scan to the IXContext interfaces which ANTLR generates for the
// rules of the grammar, leaving out the antlr.Parser and antlr.Token fields of a context.
func isRuleContextInterface(candidate reflect.Type) bool {
	return candidate.Kind() == reflect.Interface &&
		strings.HasPrefix(candidate.Name(), "I") && strings.HasSuffix(candidate.Name(), "Context")
}

func cardinalityOf(contextType reflect.Type) *contextCardinality {
	result := &contextCardinality{
		allowed:    map[reflect.Type]int{},
		repeatable: map[reflect.Type]bool{},
		resolved:   map[reflect.Type]reflect.Type{},
	}

	structType := contextType.Elem()

	for i := 0; i < structType.NumField(); i++ {
		fieldType := structType.Field(i).Type

		if isRuleContextInterface(fieldType) {
			result.allowed[fieldType]++

			continue
		}

		if fieldType.Kind() == reflect.Slice && isRuleContextInterface(fieldType.Elem()) {
			result.repeatable[fieldType.Elem()] = true
		}
	}

	return result
}

// counted returns the interface under which a child is counted, or nil when the child is not
// labelled or is collected by a repeatable label.
func (c *contextCardinality) counted(childType reflect.Type) reflect.Type {
	if known, ok := c.resolved[childType]; ok {
		return known
	}

	var match reflect.Type

	for candidate := range c.allowed {
		if !c.repeatable[candidate] && childType.Implements(candidate) {
			match = candidate

			break
		}
	}

	c.resolved[childType] = match

	return match
}

// keywordOf names the element which introduces a context and the token to report it at: for a
// block that is the name behind /begin, for every other element the first token.
func keywordOf(ctx antlr.ParserRuleContext) (keyword string, at antlr.Token) {
	for _, child := range ctx.GetChildren() {
		terminal, ok := child.(antlr.TerminalNode)
		if !ok {
			continue
		}

		if text := terminal.GetText(); text != "/begin" {
			return text, terminal.GetSymbol()
		}
	}

	at = ctx.GetStart()

	return at.GetText(), at
}

// EnterEveryRule reports an element which the specification allows at most once but which appears
// several times below the same parent. Chapter 3.5.1: "If the keyword can be used multiple times
// this is shown with help of asterisk after the closing bracket, e.g. [-> keyword]*." The grammar
// cannot enforce it, every optional element sits in a "( a | b | c )*" loop, and the tree keeps
// only the last occurrence, so without this the earlier ones are discarded silently.
func (l *versionCheckListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	contextType := reflect.TypeOf(ctx)
	if contextType == nil || contextType.Kind() != reflect.Pointer ||
		contextType.Elem().Kind() != reflect.Struct {
		return
	}

	cardinality, known := l.cardinalities[contextType]
	if !known {
		cardinality = cardinalityOf(contextType)
		l.cardinalities[contextType] = cardinality
	}

	if len(cardinality.allowed) == 0 {
		return
	}

	seen := map[reflect.Type]int{}

	for _, child := range ctx.GetChildren() {
		rule, ok := child.(antlr.ParserRuleContext)
		if !ok {
			continue
		}

		counted := cardinality.counted(reflect.TypeOf(rule))
		if counted == nil {
			continue
		}

		seen[counted]++

		if seen[counted] > cardinality.allowed[counted] {
			keyword, at := keywordOf(rule)
			l.warn(fmt.Sprintf(
				"%s may be used at most once here, only the last occurrence is kept", keyword), at)
		}
	}
}

func (l *versionCheckListener) fail(message string, at antlr.Token) {
	l.errors = append(l.errors, SyntaxError{
		Line:   at.GetLine(),
		Column: at.GetColumn(),
		Msg:    message,
	})
}

func (l *versionCheckListener) require(min asap2Version, construct string, at antlr.Token) {
	if l.version == nil || l.version.atLeast(min) {
		return
	}

	l.warn(fmt.Sprintf("%s requires ASAP2 version %s, but the file declares ASAP2_VERSION %d %d",
		construct, min, l.version.versionNo, l.version.upgradeNo), at)
}

// removedIn reports a construct which the standard withdrew with version removed. The grammar
// keeps accepting it, because it covers the older versions as well, but a file which declares
// removed or a newer version must not use it any more.
func (l *versionCheckListener) removedIn(removed asap2Version, construct string, at antlr.Token) {
	if l.version == nil || !l.version.atLeast(removed) {
		return
	}

	l.warn(fmt.Sprintf("%s was removed in ASAP2 version %s, but the file declares ASAP2_VERSION %d %d",
		construct, removed, l.version.versionNo, l.version.upgradeNo), at)
}

// reducedIn reports a construct which the standard reduced to a single occurrence with version
// reduced. The grammar keeps a repeatable label, because it covers the older versions as well, so
// the generic check of EnterEveryRule, which reads the cardinality from the labels, cannot see it.
func (l *versionCheckListener) reducedIn(reduced asap2Version, construct string, at antlr.Token) {
	if l.version == nil || !l.version.atLeast(reduced) {
		return
	}

	l.warn(fmt.Sprintf(
		"%s may be used at most once since ASAP2 version %s, but the file declares ASAP2_VERSION %d %d",
		construct, reduced, l.version.versionNo, l.version.upgradeNo), at)
}

// ASAP2 1.51 (chapter 6.3.26) declares CALIBRATION_HANDLE as "{-> CALIBRATION_HANDLE}*", ASAM
// MCD-2 MC 1.6.1 (chapter 3.5.28) reduced it to "[-> CALIBRATION_HANDLE]".
func (l *versionCheckListener) EnterCalibrationMethod(ctx *parser.CalibrationMethodContext) {
	handles := ctx.GetV_calibrationHandle()
	if len(handles) < 2 {
		return
	}

	for _, handle := range handles[1:] {
		l.reducedIn(asap2Version161, "CALIBRATION_HANDLE", handle.GetStart())
	}
}

func (l *versionCheckListener) requireValue(min asap2Version, keyword string, value antlr.Token, gated ...string) {
	if value == nil {
		return
	}

	for _, g := range gated {
		if value.GetText() == g {
			l.require(min, fmt.Sprintf("%s %s", keyword, value.GetText()), value)
			return
		}
	}
}

func (l *versionCheckListener) EnterAsap2Version(ctx *parser.Asap2VersionContext) {
	version := &asap2Version{}

	if v := a2lIntToIntType(ctx.GetVersionNo()); v != nil {
		version.versionNo = v.Value
	}

	if v := a2lIntToIntType(ctx.GetUpgradeNo()); v != nil {
		version.upgradeNo = v.Value
	}

	l.version = version
}

func (l *versionCheckListener) EnterStepSize(ctx *parser.StepSizeContext) {
	l.require(asap2Version160, "STEP_SIZE", ctx.GetStart())
}

func (l *versionCheckListener) EnterPhysUnit(ctx *parser.PhysUnitContext) {
	l.require(asap2Version160, "PHYS_UNIT", ctx.GetStart())
}

func (l *versionCheckListener) EnterDiscrete(ctx *parser.DiscreteContext) {
	l.require(asap2Version160, "DISCRETE", ctx.GetStart())
}

func (l *versionCheckListener) EnterSymbolLink(ctx *parser.SymbolLinkContext) {
	l.require(asap2Version160, "SYMBOL_LINK", ctx.GetStart())
}

func (l *versionCheckListener) EnterLayout(ctx *parser.LayoutContext) {
	l.require(asap2Version160, "LAYOUT", ctx.GetStart())
}

func (l *versionCheckListener) EnterCoeffsLinear(ctx *parser.CoeffsLinearContext) {
	l.require(asap2Version160, "COEFFS_LINEAR", ctx.GetStart())
}

func (l *versionCheckListener) EnterStatusStringRef(ctx *parser.StatusStringRefContext) {
	l.require(asap2Version160, "STATUS_STRING_REF", ctx.GetStart())
}

func (l *versionCheckListener) EnterDefaultValueNumeric(ctx *parser.DefaultValueNumericContext) {
	l.require(asap2Version160, "DEFAULT_VALUE_NUMERIC", ctx.GetStart())
}

func (l *versionCheckListener) EnterAlignmentInt64(ctx *parser.AlignmentInt64Context) {
	l.require(asap2Version160, "ALIGNMENT_INT64", ctx.GetStart())
}

func (l *versionCheckListener) EnterStaticRecordLayout(ctx *parser.StaticRecordLayoutContext) {
	l.require(asap2Version160, "STATIC_RECORD_LAYOUT", ctx.GetStart())
}

func (l *versionCheckListener) EnterAxisPts4(ctx *parser.AxisPts4Context) {
	l.require(asap2Version160, "AXIS_PTS_4", ctx.GetStart())
}

func (l *versionCheckListener) EnterAxisPts5(ctx *parser.AxisPts5Context) {
	l.require(asap2Version160, "AXIS_PTS_5", ctx.GetStart())
}

// ASAM MCD-2 MC 1.6.1 (chapter 1.4.4) reduced the AXIS_RESCALE family to AXIS_RESCALE_X and the
// NO_RESCALE family to NO_RESCALE_X. The _4 and _5 members therefore exist in version 1.60 only.
func (l *versionCheckListener) EnterAxisRescaleY(ctx *parser.AxisRescaleYContext) {
	l.removedIn(asap2Version161, "AXIS_RESCALE_Y", ctx.GetStart())
}

func (l *versionCheckListener) EnterAxisRescaleZ(ctx *parser.AxisRescaleZContext) {
	l.removedIn(asap2Version161, "AXIS_RESCALE_Z", ctx.GetStart())
}

func (l *versionCheckListener) EnterAxisRescale4(ctx *parser.AxisRescale4Context) {
	l.require(asap2Version160, "AXIS_RESCALE_4", ctx.GetStart())
	l.removedIn(asap2Version161, "AXIS_RESCALE_4", ctx.GetStart())
}

func (l *versionCheckListener) EnterAxisRescale5(ctx *parser.AxisRescale5Context) {
	l.require(asap2Version160, "AXIS_RESCALE_5", ctx.GetStart())
	l.removedIn(asap2Version161, "AXIS_RESCALE_5", ctx.GetStart())
}

func (l *versionCheckListener) EnterNoAxisPts4(ctx *parser.NoAxisPts4Context) {
	l.require(asap2Version160, "NO_AXIS_PTS_4", ctx.GetStart())
}

func (l *versionCheckListener) EnterNoAxisPts5(ctx *parser.NoAxisPts5Context) {
	l.require(asap2Version160, "NO_AXIS_PTS_5", ctx.GetStart())
}

func (l *versionCheckListener) EnterNoRescaleY(ctx *parser.NoRescaleYContext) {
	l.removedIn(asap2Version161, "NO_RESCALE_Y", ctx.GetStart())
}

func (l *versionCheckListener) EnterNoRescaleZ(ctx *parser.NoRescaleZContext) {
	l.removedIn(asap2Version161, "NO_RESCALE_Z", ctx.GetStart())
}

func (l *versionCheckListener) EnterNoRescale4(ctx *parser.NoRescale4Context) {
	l.require(asap2Version160, "NO_RESCALE_4", ctx.GetStart())
	l.removedIn(asap2Version161, "NO_RESCALE_4", ctx.GetStart())
}

func (l *versionCheckListener) EnterNoRescale5(ctx *parser.NoRescale5Context) {
	l.require(asap2Version160, "NO_RESCALE_5", ctx.GetStart())
	l.removedIn(asap2Version161, "NO_RESCALE_5", ctx.GetStart())
}

// S_REC_LAYOUT is declared by ASAP2 1.51 (chapter 6.3.126) for MOD_COMMON and was withdrawn by
// ASAM MCD-2 MC 1.6.1 (chapter 1.4.4).
func (l *versionCheckListener) EnterSRecLayout(ctx *parser.SRecLayoutContext) {
	l.removedIn(asap2Version161, "S_REC_LAYOUT", ctx.GetStart())
}

func (l *versionCheckListener) EnterFixNoAxisPts4(ctx *parser.FixNoAxisPts4Context) {
	l.require(asap2Version160, "FIX_NO_AXIS_PTS_4", ctx.GetStart())
}

func (l *versionCheckListener) EnterFixNoAxisPts5(ctx *parser.FixNoAxisPts5Context) {
	l.require(asap2Version160, "FIX_NO_AXIS_PTS_5", ctx.GetStart())
}

func (l *versionCheckListener) EnterSrcAddr4(ctx *parser.SrcAddr4Context) {
	l.require(asap2Version160, "SRC_ADDR_4", ctx.GetStart())
}

func (l *versionCheckListener) EnterSrcAddr5(ctx *parser.SrcAddr5Context) {
	l.require(asap2Version160, "SRC_ADDR_5", ctx.GetStart())
}

func (l *versionCheckListener) EnterRipAddr4(ctx *parser.RipAddr4Context) {
	l.require(asap2Version160, "RIP_ADDR_4", ctx.GetStart())
}

func (l *versionCheckListener) EnterRipAddr5(ctx *parser.RipAddr5Context) {
	l.require(asap2Version160, "RIP_ADDR_5", ctx.GetStart())
}

func (l *versionCheckListener) EnterShiftOp4(ctx *parser.ShiftOp4Context) {
	l.require(asap2Version160, "SHIFT_OP_4", ctx.GetStart())
}

func (l *versionCheckListener) EnterShiftOp5(ctx *parser.ShiftOp5Context) {
	l.require(asap2Version160, "SHIFT_OP_5", ctx.GetStart())
}

func (l *versionCheckListener) EnterOffset4(ctx *parser.Offset4Context) {
	l.require(asap2Version160, "OFFSET_4", ctx.GetStart())
}

func (l *versionCheckListener) EnterOffset5(ctx *parser.Offset5Context) {
	l.require(asap2Version160, "OFFSET_5", ctx.GetStart())
}

func (l *versionCheckListener) EnterDistOp4(ctx *parser.DistOp4Context) {
	l.require(asap2Version160, "DIST_OP_4", ctx.GetStart())
}

func (l *versionCheckListener) EnterDistOp5(ctx *parser.DistOp5Context) {
	l.require(asap2Version160, "DIST_OP_5", ctx.GetStart())
}

func (l *versionCheckListener) EnterCalibrationHandleText(ctx *parser.CalibrationHandleTextContext) {
	l.require(asap2Version160, "CALIBRATION_HANDLE_TEXT", ctx.GetStart())
}

// ASAP2 1.51 does not declare IF_DATA for FUNCTION and GROUP, ASAM MCD-2 MC 1.6.0 added it.
func (l *versionCheckListener) EnterIfData(ctx *parser.IfDataContext) {
	switch ctx.GetParent().(type) {
	case *parser.FunctionContext:
		l.require(asap2Version160, "IF_DATA in FUNCTION", ctx.GetStart())
	case *parser.GroupContext:
		l.require(asap2Version160, "IF_DATA in GROUP", ctx.GetStart())
	}
}

func (l *versionCheckListener) EnterCharacteristic(ctx *parser.CharacteristicContext) {
	l.requireValue(asap2Version160, "CHARACTERISTIC type", ctx.GetType_(), "CUBE_4", "CUBE_5")
}

func (l *versionCheckListener) EnterMonotony(ctx *parser.MonotonyContext) {
	l.requireValue(asap2Version160, "MONOTONY", ctx.GetMonotony_(), "MONOTONOUS", "STRICT_MON", "NOT_MON")

	// ASAP2 1.51 declares MONOTONY for AXIS_DESCR only, ASAM MCD-2 MC 1.6.0 added it to AXIS_PTS
	if _, ok := ctx.GetParent().(*parser.AxisPtsContext); ok {
		l.require(asap2Version160, "MONOTONY in AXIS_PTS", ctx.GetStart())
	}
}

func (l *versionCheckListener) EnterCompuMethod(ctx *parser.CompuMethodContext) {
	l.requireValue(asap2Version160, "conversion type", ctx.GetConversionType(), "IDENTICAL", "LINEAR")
}

func (l *versionCheckListener) EnterDataType(ctx *parser.DataTypeContext) {
	l.requireValue(asap2Version160, "data type", ctx.GetV(), "A_UINT64", "A_INT64")
}

// allowedStringEscapes are the characters which may follow a backslash inside a string: ASAP2
// 1.51 (chapter 6.2) and ASAM MCD-2 MC 1.6.1 (chapter 3.2) allow \", \', \\, \n, \r and \t.
const allowedStringEscapes = "rnt\"'\\"

// EnterStringValue rejects a backslash which does not start one of the allowed escape sequences.
// The lexer deliberately consumes any backslash, so that a malformed sequence does not
// desynchronize it; the restriction of the specification is enforced here instead.
func (l *versionCheckListener) EnterStringValue(ctx *parser.StringValueContext) {
	value := ctx.GetS()
	if value == nil {
		return
	}

	text := value.GetText()

	for i := 0; i < len(text); i++ {
		if text[i] != '\\' {
			continue
		}

		// the token always ends with the closing double inverted comma, so a backslash is never
		// its last character; the guard keeps a malformed token from panicking here
		if i+1 >= len(text) {
			break
		}

		if strings.IndexByte(allowedStringEscapes, text[i+1]) >= 0 {
			i++

			continue
		}

		l.fail(fmt.Sprintf(
			"\\%c is not an escape sequence allowed inside a string, expected one of \\\" \\' \\\\ \\n \\r \\t",
			text[i+1]), value)

		return
	}
}

// checkUint64 reports a literal which does not fit the uint64 parameters that ASAM MCD-2 MC 1.6.1
// introduced for BIT_MASK and ERROR_MASK (chapter 1.4.4), so that a negative or oversized value is
// reported at its own position instead of surfacing as an internal error while the tree is built.
func (l *versionCheckListener) checkUint64(keyword string, value parser.IIntegerValueContext) {
	if value == nil {
		return
	}

	if _, _, _, err := parseUint64Literal(value.GetText()); err != nil {
		l.fail(fmt.Sprintf("%s is not a valid %s, it must fit an unsigned 64 bit integer",
			value.GetText(), keyword), value.GetStart())
	}
}

func (l *versionCheckListener) EnterBitMask(ctx *parser.BitMaskContext) {
	l.checkUint64("BIT_MASK", ctx.GetMask())
}

func (l *versionCheckListener) EnterErrorMask(ctx *parser.ErrorMaskContext) {
	l.checkUint64("ERROR_MASK", ctx.GetMask())
}

// EnterPartialIdentifier rejects an empty partial identifier. Chapter 3.2 describes an identifier
// as a "hierarchical concatenation of partial strings separated by points", so "a." and "a..b"
// name a partial string which is not there. A partial string beginning with a digit is valid: the
// chapter requires a letter or an underscore for the first character of the identifier only, and
// names the partial string explicitly where it means one. The point is part of the IDENT token,
// for the reason the grammar gives, so the grammar cannot express the restriction on its own.
func (l *versionCheckListener) EnterPartialIdentifier(ctx *parser.PartialIdentifierContext) {
	identifier := ctx.GetI()
	if identifier == nil {
		return
	}

	text := identifier.GetText()

	for _, partial := range strings.Split(text, ".") {
		if partial == "" {
			l.fail(fmt.Sprintf(
				"%s is not a valid identifier, a partial identifier must not be empty", text), identifier)

			return
		}
	}
}

// EnterArraySpecifier rejects a malformed array index. Chapter 3.2 requires the brackets at the
// end of a partial identifier to "contain a number or an alphabetic string (description of the
// index of an array element)", the latter being "a symbolic string which is defined as an
// enumerator of an ENUM definition of the C program". A signed literal is not such a number and an
// enumerator of C contains no point; the sign belongs to the INT token and the point to the IDENT
// token, so the grammar cannot express either restriction on its own.
func (l *versionCheckListener) EnterArraySpecifier(ctx *parser.ArraySpecifierContext) {
	if index := ctx.GetI(); index != nil {
		if text := index.GetText(); strings.HasPrefix(text, "-") || strings.HasPrefix(text, "+") {
			l.fail(fmt.Sprintf("%s is not a valid array index, it must be unsigned", text), index)
		}

		return
	}

	if name := ctx.GetN(); name != nil {
		if text := name.GetText(); strings.Contains(text, ".") {
			l.fail(fmt.Sprintf(
				"%s is not a valid array index, a symbolic index is a single enumerator", text), name)
		}
	}
}

// EnterVarNaming validates the tag. The grammar matches it as an identifier so that ALPHA stays
// usable as an ordinary identifier, which means the set of accepted values is checked here
// instead. NUMERIC is the only value the standard defines; ALPHA is reserved for a future
// extension (ASAM MCD-2 MC 1.6.1, chapter 3.5.134) and is therefore reported as a warning, which
// unlike the version gates above does not depend on the declared version; anything else is
// rejected.
func (l *versionCheckListener) EnterVarNaming(ctx *parser.VarNamingContext) {
	tag := ctx.GetTag()
	if tag == nil {
		return
	}

	switch tag.GetText() {
	case "NUMERIC":
	case "ALPHA":
		l.warn("VAR_NAMING ALPHA is reserved for a future extension of the standard", tag)
	default:
		l.fail(fmt.Sprintf("%s is not a valid VAR_NAMING tag, expected NUMERIC", tag.GetText()), tag)
	}
}
