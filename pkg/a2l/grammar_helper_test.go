package a2l

// Helpers shared by the grammar test suite.
//
// The suite contains one test function per A2L keyword. Each of them checks, in a minimal
// but valid enclosing context:
//   - the mandatory parameters,
//   - every optional parameter,
//   - list parameters with zero, one and several elements,
//   - inputs which the specification requires to be rejected.
//
// Every successful parse is additionally checked for round-trip stability
// (tree -> MarshalA2L -> tree), so that a grammar change which is not reflected in the
// protobuf tree or in the serializer is caught here as well.
//
// Reference: ASAM MCD-2MC (ASAP2) interface specification, chapter 6.3.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

/*
** Scope builders: wrap an A2L fragment into the smallest valid enclosing blocks.
 */

func fileScope(body string) string {
	return body + "\n"
}

func projectScope(body string) string {
	return fileScope("/begin PROJECT project \"\"\n" + body + "\n/end PROJECT")
}

func moduleScope(body string) string {
	return projectScope("/begin MODULE module \"\"\n" + body + "\n/end MODULE")
}

func modCommonScope(body string) string {
	return moduleScope("/begin MOD_COMMON \"\"\n" + body + "\n/end MOD_COMMON")
}

func modParScope(body string) string {
	return moduleScope("/begin MOD_PAR \"\"\n" + body + "\n/end MOD_PAR")
}

func characteristicScope(body string) string {
	return moduleScope("/begin CHARACTERISTIC characteristic \"\" VALUE 0x0 record_layout 0 compu_method 0 0\n" +
		body + "\n/end CHARACTERISTIC")
}

func axisDescrScope(body string) string {
	return characteristicScope("/begin AXIS_DESCR STD_AXIS input_quantity compu_method 1 0 0\n" +
		body + "\n/end AXIS_DESCR")
}

func measurementScope(body string) string {
	return moduleScope("/begin MEASUREMENT measurement \"\" UBYTE compu_method 0 0 0 0\n" +
		body + "\n/end MEASUREMENT")
}

func axisPtsScope(body string) string {
	return moduleScope("/begin AXIS_PTS axis_pts \"\" 0x0 input_quantity record_layout 0 compu_method 8 0 0\n" +
		body + "\n/end AXIS_PTS")
}

func recordLayoutScope(body string) string {
	return moduleScope("/begin RECORD_LAYOUT record_layout\n" + body + "\n/end RECORD_LAYOUT")
}

func compuMethodScope(body string) string {
	return moduleScope("/begin COMPU_METHOD compu_method \"\" RAT_FUNC \"%4.2\" \"unit\"\n" +
		body + "\n/end COMPU_METHOD")
}

func functionScope(body string) string {
	return moduleScope("/begin FUNCTION function \"\"\n" + body + "\n/end FUNCTION")
}

func groupScope(body string) string {
	return moduleScope("/begin GROUP group \"\"\n" + body + "\n/end GROUP")
}

func unitScope(body string) string {
	return moduleScope("/begin UNIT unit \"\" \"display\" DERIVED\n" + body + "\n/end UNIT")
}

func userRightsScope(body string) string {
	return moduleScope("/begin USER_RIGHTS user_level\n" + body + "\n/end USER_RIGHTS")
}

func frameScope(body string) string {
	return moduleScope("/begin FRAME frame \"\" 1 10\n" + body + "\n/end FRAME")
}

func variantCodingScope(body string) string {
	return moduleScope("/begin VARIANT_CODING\n" + body + "\n/end VARIANT_CODING")
}

func varCriterionScope(body string) string {
	return variantCodingScope("/begin VAR_CRITERION criterion \"\" first second\n" + body + "\n/end VAR_CRITERION")
}

/*
** Parse helpers. Every helper returns ok == false when the fragment could not be parsed or
** when the expected node is missing, so that the caller can stop without dereferencing nil.
 */

// parse parses src, asserts that no syntax error is reported and that the resulting tree is
// round-trip stable.
func parse(t assert.TestingT, src string) (*RootNodeType, bool) {
	markHelper(t)

	tree, ok := parseOnly(t, src)
	if !ok {
		return nil, false
	}

	assertRoundTrip(t, tree)

	return tree, true
}

// parseOnly parses src and asserts that no syntax error is reported, without checking the
// round-trip stability of the tree. It is meant for the cases where the serializer is known not to
// reproduce the layout of the original content.
func parseOnly(t assert.TestingT, src string) (*RootNodeType, bool) {
	markHelper(t)

	tree, err := GetTreeFromString(src)
	if !assert.NoError(t, err, "unexpected syntax error while parsing:\n%s", src) {
		return nil, false
	}

	return tree, true
}

/*
** Navigation inside a parsed tree.
 */

func moduleOf(t assert.TestingT, tree *RootNodeType) (*ModuleType, bool) {
	markHelper(t)

	if tree == nil || tree.PROJECT == nil || len(tree.PROJECT.MODULE) != 1 {
		return nil, assert.Fail(t, "MODULE is missing from the tree")
	}

	return tree.PROJECT.MODULE[0], true
}

func characteristicOf(t assert.TestingT, tree *RootNodeType) (*CharacteristicType, bool) {
	markHelper(t)

	module, ok := moduleOf(t, tree)
	if !ok {
		return nil, false
	}

	if len(module.CHARACTERISTIC) != 1 {
		return nil, assert.Fail(t, "CHARACTERISTIC is missing from the tree")
	}

	return module.CHARACTERISTIC[0], true
}

func axisDescrOf(t assert.TestingT, tree *RootNodeType) (*AxisDescrType, bool) {
	markHelper(t)

	characteristic, ok := characteristicOf(t, tree)
	if !ok {
		return nil, false
	}

	if len(characteristic.AXIS_DESCR) != 1 {
		return nil, assert.Fail(t, "AXIS_DESCR is missing from the tree")
	}

	return characteristic.AXIS_DESCR[0], true
}

// parseFails asserts that src is rejected by the parser.
func parseFails(t assert.TestingT, src string) bool {
	markHelper(t)

	_, err := GetTreeFromString(src)

	return assert.Error(t, err, "the parser should have rejected:\n%s", src)
}

// parseFailsWithSyntaxError asserts that src is rejected by the parser, and that the rejection is
// reported as a syntax error instead of an internal error raised by a panic.
func parseFailsWithSyntaxError(t assert.TestingT, src string) bool {
	markHelper(t)

	_, err := GetTreeFromString(src)
	if !assert.Error(t, err, "the parser should have rejected:\n%s", src) {
		return false
	}

	return assert.NotContains(t, err.Error(), "error while building A2L tree",
		"the parser should report a syntax error instead of an internal error, for:\n%s", src)
}

func parseProject(t assert.TestingT, body string) (*ProjectType, bool) {
	markHelper(t)

	tree, ok := parse(t, projectScope(body))
	if !ok {
		return nil, false
	}

	return tree.PROJECT, assert.NotNil(t, tree.PROJECT, "PROJECT is missing from the tree")
}

func parseModule(t assert.TestingT, body string) (*ModuleType, bool) {
	markHelper(t)

	tree, ok := parse(t, moduleScope(body))
	if !ok {
		return nil, false
	}

	return moduleOf(t, tree)
}

func parseModCommon(t assert.TestingT, body string) (*ModCommonType, bool) {
	markHelper(t)

	module, ok := parseModule(t, "/begin MOD_COMMON \"\"\n"+body+"\n/end MOD_COMMON")
	if !ok {
		return nil, false
	}

	return module.MOD_COMMON, assert.NotNil(t, module.MOD_COMMON, "MOD_COMMON is missing from the tree")
}

func parseCharacteristic(t assert.TestingT, body string) (*CharacteristicType, bool) {
	markHelper(t)

	tree, ok := parse(t, characteristicScope(body))
	if !ok {
		return nil, false
	}

	return characteristicOf(t, tree)
}

func parseAxisDescr(t assert.TestingT, body string) (*AxisDescrType, bool) {
	markHelper(t)

	tree, ok := parse(t, axisDescrScope(body))
	if !ok {
		return nil, false
	}

	return axisDescrOf(t, tree)
}

func parseRecordLayout(t assert.TestingT, body string) (*RecordLayoutType, bool) {
	markHelper(t)

	module, ok := parseModule(t, "/begin RECORD_LAYOUT record_layout\n"+body+"\n/end RECORD_LAYOUT")
	if !ok {
		return nil, false
	}

	if len(module.RECORD_LAYOUT) != 1 {
		return nil, assert.Fail(t, "RECORD_LAYOUT is missing from the tree")
	}

	return module.RECORD_LAYOUT[0], true
}

func parseModPar(t assert.TestingT, body string) (*ModParType, bool) {
	markHelper(t)

	module, ok := parseModule(t, "/begin MOD_PAR \"\"\n"+body+"\n/end MOD_PAR")
	if !ok {
		return nil, false
	}

	return module.MOD_PAR, assert.NotNil(t, module.MOD_PAR, "MOD_PAR is missing from the tree")
}

func parseAxisPts(t assert.TestingT, body string) (*AxisPtsType, bool) {
	markHelper(t)

	tree, ok := parse(t, axisPtsScope(body))
	if !ok {
		return nil, false
	}

	module, ok := moduleOf(t, tree)
	if !ok {
		return nil, false
	}

	if len(module.AXIS_PTS) != 1 {
		return nil, assert.Fail(t, "AXIS_PTS is missing from the tree")
	}

	return module.AXIS_PTS[0], true
}

func parseMeasurement(t assert.TestingT, body string) (*MeasurementType, bool) {
	markHelper(t)

	module, ok := parseModule(t,
		"/begin MEASUREMENT measurement \"\" UBYTE compu_method 0 0 0 0\n"+body+"\n/end MEASUREMENT")
	if !ok {
		return nil, false
	}

	if len(module.MEASUREMENT) != 1 {
		return nil, assert.Fail(t, "MEASUREMENT is missing from the tree")
	}

	return module.MEASUREMENT[0], true
}

func parseCompuMethod(t assert.TestingT, body string) (*CompuMethodType, bool) {
	markHelper(t)

	tree, ok := parse(t, compuMethodScope(body))
	if !ok {
		return nil, false
	}

	module, ok := moduleOf(t, tree)
	if !ok {
		return nil, false
	}

	if len(module.COMPU_METHOD) != 1 {
		return nil, assert.Fail(t, "COMPU_METHOD is missing from the tree")
	}

	return module.COMPU_METHOD[0], true
}

func parseFunction(t assert.TestingT, body string) (*FunctionType, bool) {
	markHelper(t)

	module, ok := parseModule(t, "/begin FUNCTION function \"\"\n"+body+"\n/end FUNCTION")
	if !ok {
		return nil, false
	}

	if len(module.FUNCTION) != 1 {
		return nil, assert.Fail(t, "FUNCTION is missing from the tree")
	}

	return module.FUNCTION[0], true
}

func parseGroup(t assert.TestingT, body string) (*GroupType, bool) {
	markHelper(t)

	module, ok := parseModule(t, "/begin GROUP group \"\"\n"+body+"\n/end GROUP")
	if !ok {
		return nil, false
	}

	if len(module.GROUP) != 1 {
		return nil, assert.Fail(t, "GROUP is missing from the tree")
	}

	return module.GROUP[0], true
}

func parseUnit(t assert.TestingT, body string) (*UnitType, bool) {
	markHelper(t)

	module, ok := parseModule(t, "/begin UNIT unit \"\" \"display\" DERIVED\n"+body+"\n/end UNIT")
	if !ok {
		return nil, false
	}

	if len(module.UNIT) != 1 {
		return nil, assert.Fail(t, "UNIT is missing from the tree")
	}

	return module.UNIT[0], true
}

func parseUserRights(t assert.TestingT, body string) (*UserRightsType, bool) {
	markHelper(t)

	module, ok := parseModule(t, "/begin USER_RIGHTS user_level\n"+body+"\n/end USER_RIGHTS")
	if !ok {
		return nil, false
	}

	if len(module.USER_RIGHTS) != 1 {
		return nil, assert.Fail(t, "USER_RIGHTS is missing from the tree")
	}

	return module.USER_RIGHTS[0], true
}

func parseFrame(t assert.TestingT, body string) (*FrameType, bool) {
	markHelper(t)

	module, ok := parseModule(t, "/begin FRAME frame \"\" 1 10\n"+body+"\n/end FRAME")
	if !ok {
		return nil, false
	}

	return module.FRAME, assert.NotNil(t, module.FRAME, "FRAME is missing from the tree")
}

func parseVarCriterion(t assert.TestingT, body string) (*VarCriterionType, bool) {
	markHelper(t)

	variantCoding, ok := parseVariantCoding(t,
		"/begin VAR_CRITERION criterion \"\" first second\n"+body+"\n/end VAR_CRITERION")
	if !ok {
		return nil, false
	}

	if len(variantCoding.VAR_CRITERION) != 1 {
		return nil, assert.Fail(t, "VAR_CRITERION is missing from the tree")
	}

	return variantCoding.VAR_CRITERION[0], true
}

func parseVariantCoding(t assert.TestingT, body string) (*VariantCodingType, bool) {
	markHelper(t)

	module, ok := parseModule(t, "/begin VARIANT_CODING\n"+body+"\n/end VARIANT_CODING")
	if !ok {
		return nil, false
	}

	return module.VARIANT_CODING, assert.NotNil(t, module.VARIANT_CODING, "VARIANT_CODING is missing from the tree")
}

/*
** Assertions.
 */

// equalNode compares a single node of the tree with its expected content.
func equalNode(t assert.TestingT, expected, actual proto.Message) bool {
	markHelper(t)

	if isNilNode(actual) {
		return assert.Fail(t, "node is missing from the tree", "expected:\n%s", protoJSON(expected))
	}

	if proto.Equal(expected, actual) {
		return true
	}

	if assert.Equal(t, protoJSON(expected), protoJSON(actual)) {
		// Same rendering but not equal: report it rather than silently passing.
		return assert.Fail(t, "nodes are not equal although they render identically")
	}

	return false
}

// equalNodes compares a repeated field of the tree with its expected content.
func equalNodes[T proto.Message](t assert.TestingT, expected, actual []T) bool {
	markHelper(t)

	if !assert.Len(t, actual, len(expected), "unexpected number of elements") {
		return false
	}

	result := true

	for i := range expected {
		if !equalNode(t, expected[i], actual[i]) {
			result = false
		}
	}

	return result
}

// assertPreserved checks that a keyword contained in src survives into the tree, by looking it up
// in the serialized tree. It detects keywords which the grammar accepts but which no node of the
// tree is able to hold, and which are therefore silently dropped.
func assertPreserved(t assert.TestingT, src string, keyword string) bool {
	markHelper(t)

	tree, err := GetTreeFromString(src)
	if !assert.NoError(t, err, "unexpected syntax error while parsing:\n%s", src) {
		return false
	}

	content, ok := marshalTree(t, tree)
	if !ok {
		return false
	}

	return assert.Contains(t, content, keyword,
		"%s is accepted by the grammar but does not reach the tree", keyword)
}

// marshalTree serializes the tree, reporting a panic of the serializer as a failed assertion
// instead of letting it abort the test binary.
func marshalTree(t assert.TestingT, tree *RootNodeType) (content string, ok bool) {
	markHelper(t)

	defer func() {
		if recovered := recover(); recovered != nil {
			content, ok = "", assert.Fail(t, "serializing the tree panics", "%v", recovered)
		}
	}()

	return tree.MarshalA2L(0, "  ", false), true
}

// assertRoundTrip checks that serializing the tree and parsing the result yields the same tree.
func assertRoundTrip(t assert.TestingT, tree *RootNodeType) bool {
	markHelper(t)

	content, ok := marshalTree(t, tree)
	if !ok {
		return false
	}

	reparsed, err := GetTreeFromString(content)
	if !assert.NoError(t, err, "the serialized tree cannot be parsed again:\n%s", content) {
		return false
	}

	if proto.Equal(tree, reparsed) {
		return true
	}

	return assert.Equal(t, protoJSON(tree), protoJSON(reparsed),
		"tree -> A2L -> tree is not stable, serialized content:\n%s", content)
}

/*
** Shared test bodies for the keyword families whose prototype is identical.
 */

// testPositionDataTypeKeyword covers the RECORD_LAYOUT components declared as
// "<keyword> Position Datatype" (chapters 6.3.46, 6.3.71, 6.3.94, 6.3.96, 6.3.98, 6.3.114,
// 6.3.118 and 6.3.121).
func testPositionDataTypeKeyword[T proto.Message](
	t *testing.T,
	keyword string,
	newNode func(position *IntType, dataType *DataTypeType) T,
	get func(recordLayout *RecordLayoutType) T,
) {
	t.Helper()

	t.Run("mandatory parameters", func(t *testing.T) {
		recordLayout, ok := parseRecordLayout(t, keyword+" 1 SWORD")
		if !ok {
			return
		}

		equalNode(t, newNode(intVal("1"), dataTypeVal("SWORD")), get(recordLayout))
	})

	t.Run("reject/missing data type", func(t *testing.T) {
		parseFails(t, recordLayoutScope(keyword+" 1"))
	})

	t.Run("reject/unknown data type", func(t *testing.T) {
		parseFails(t, recordLayoutScope(keyword+" 1 UINT24"))
	})

	t.Run("reject/data size instead of data type", func(t *testing.T) {
		parseFails(t, recordLayoutScope(keyword+" 1 WORD"))
	})
}

// testAxisPtsKeyword covers the RECORD_LAYOUT components declared as
// "<keyword> Position Datatype IndexIncr Addressing" (chapter 6.3.19).
func testAxisPtsKeyword[T proto.Message](
	t *testing.T,
	keyword string,
	newNode func(position *IntType, dataType *DataTypeType, indexIncr *IndexOrderType, addressing *AddrTypeType) T,
	get func(recordLayout *RecordLayoutType) T,
) {
	t.Helper()

	t.Run("mandatory parameters", func(t *testing.T) {
		recordLayout, ok := parseRecordLayout(t, keyword+" 1 UWORD INDEX_INCR DIRECT")
		if !ok {
			return
		}

		equalNode(t, newNode(intVal("1"), dataTypeVal("UWORD"), indexOrderVal("INDEX_INCR"), addrTypeVal("DIRECT")),
			get(recordLayout))
	})

	for _, indexOrder := range []string{"INDEX_INCR", "INDEX_DECR"} {
		t.Run("enum/index order "+indexOrder, func(t *testing.T) {
			recordLayout, ok := parseRecordLayout(t, keyword+" 1 UWORD "+indexOrder+" DIRECT")
			if !ok {
				return
			}

			equalNode(t, newNode(intVal("1"), dataTypeVal("UWORD"), indexOrderVal(indexOrder), addrTypeVal("DIRECT")),
				get(recordLayout))
		})
	}

	for _, addrType := range []string{"PBYTE", "PWORD", "PLONG", "DIRECT"} {
		t.Run("enum/addressing "+addrType, func(t *testing.T) {
			recordLayout, ok := parseRecordLayout(t, keyword+" 1 UWORD INDEX_INCR "+addrType)
			if !ok {
				return
			}

			equalNode(t, newNode(intVal("1"), dataTypeVal("UWORD"), indexOrderVal("INDEX_INCR"), addrTypeVal(addrType)),
				get(recordLayout))
		})
	}

	t.Run("reject/missing addressing", func(t *testing.T) {
		parseFails(t, recordLayoutScope(keyword+" 1 UWORD INDEX_INCR"))
	})

	t.Run("reject/unknown index order", func(t *testing.T) {
		parseFails(t, recordLayoutScope(keyword+" 1 UWORD INDEX_CONST DIRECT"))
	})
}

// testAxisRescaleKeyword covers the RECORD_LAYOUT components declared as
// "<keyword> Position Datatype MaxNumberOfRescalePairs IndexIncr Addressing" (chapter 6.3.20).
func testAxisRescaleKeyword[T proto.Message](
	t *testing.T,
	keyword string,
	newNode func(position *IntType, dataType *DataTypeType, maxNumberOfRescalePairs *IntType,
		indexIncr *IndexOrderType, addressing *AddrTypeType) T,
	get func(recordLayout *RecordLayoutType) T,
) {
	t.Helper()

	t.Run("mandatory parameters", func(t *testing.T) {
		recordLayout, ok := parseRecordLayout(t, keyword+" 1 UWORD 4 INDEX_INCR DIRECT")
		if !ok {
			return
		}

		equalNode(t, newNode(intVal("1"), dataTypeVal("UWORD"), intVal("4"),
			indexOrderVal("INDEX_INCR"), addrTypeVal("DIRECT")), get(recordLayout))
	})

	t.Run("reject/missing number of rescale pairs", func(t *testing.T) {
		parseFails(t, recordLayoutScope(keyword+" 1 UWORD INDEX_INCR DIRECT"))
	})

	t.Run("reject/float number of rescale pairs", func(t *testing.T) {
		parseFails(t, recordLayoutScope(keyword+" 1 UWORD 4.0 INDEX_INCR DIRECT"))
	})
}

// testFixNoAxisPtsKeyword covers the RECORD_LAYOUT components declared as
// "<keyword> NumberOfAxisPoints" (chapter 6.3.58).
func testFixNoAxisPtsKeyword[T proto.Message](
	t *testing.T,
	keyword string,
	newNode func(numberOfAxisPoints *IntType) T,
	get func(recordLayout *RecordLayoutType) T,
) {
	t.Helper()

	t.Run("mandatory parameters", func(t *testing.T) {
		recordLayout, ok := parseRecordLayout(t, keyword+" 17")
		if !ok {
			return
		}

		equalNode(t, newNode(intVal("17")), get(recordLayout))
	})

	t.Run("reject/missing number of axis points", func(t *testing.T) {
		parseFails(t, recordLayoutScope(keyword))
	})

	t.Run("reject/float number of axis points", func(t *testing.T) {
		parseFails(t, recordLayoutScope(keyword+" 17.0"))
	})
}

// testIdentifierListKeyword covers the keywords declared as a block containing a list of
// identifiers, e.g. "/begin <keyword> { Identifier }* /end <keyword>".
func testIdentifierListKeyword[T proto.Message](
	t *testing.T,
	keyword string,
	scope func(body string) string,
	newNode func(identifiers []*IdentType) T,
	parseIn func(t assert.TestingT, body string) (T, bool),
) {
	t.Helper()

	type testCaseType struct {
		name        string
		content     string
		identifiers []*IdentType
	}

	for _, testCase := range []testCaseType{
		{name: "list/no identifier", content: "", identifiers: nil},
		{name: "list/single identifier", content: "first", identifiers: []*IdentType{identVal("first")}},
		{
			name:        "list/several identifiers",
			content:     "first second third",
			identifiers: []*IdentType{identVal("first"), identVal("second"), identVal("third")},
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			node, ok := parseIn(t, "/begin "+keyword+" "+testCase.content+"\n/end "+keyword)
			if !ok {
				return
			}

			equalNode(t, newNode(testCase.identifiers), node)
		})
	}

	t.Run("reject/string instead of identifier", func(t *testing.T) {
		parseFails(t, scope("/begin "+keyword+" \"identifier\"\n/end "+keyword))
	})

	t.Run("reject/missing /end", func(t *testing.T) {
		parseFails(t, scope("/begin "+keyword+" identifier"))
	})
}

/*
** Known deviations from the specification.
 */

// deviation documents a known deviation of the current implementation from the specification. The
// assertions of body describe the behaviour required by the specification:
//
//   - as long as they fail, the test fails as well and reports the deviation, so that a run of the
//     suite lists everything which still has to be corrected;
//   - as soon as they all pass, the deviation has been corrected and the test fails with a
//     different message, so that the wrapper gets removed and the assertions become binding.
func deviation(t *testing.T, reference string, body func(t assert.TestingT)) {
	t.Helper()

	recorder := &deviationRecorder{}

	body(recorder)

	if len(recorder.failures) == 0 {
		t.Fatalf("known deviation %q does not reproduce anymore: "+
			"remove the deviation() wrapper so that the assertions become binding", reference)
	}

	t.Errorf("deviation from the specification: %s\n%s", reference, strings.Join(recorder.failures, "\n"))
}

type deviationRecorder struct {
	failures []string
}

func (r *deviationRecorder) Errorf(format string, args ...interface{}) {
	r.failures = append(r.failures, fmt.Sprintf(format, args...))
}

/*
** Value builders: they return the node the parser is expected to produce for an A2L literal.
 */

func intVal(literal string) *IntType {
	base, digits := numberBase(literal)

	value, err := strconv.ParseInt(digits, base, 64)
	if err != nil {
		panic(err)
	}

	return &IntType{Value: int32(value), Base: uint32(base), Size: uint32(len(digits))}
}

func longVal(literal string) *LongType {
	base, digits := numberBase(literal)

	value, err := strconv.ParseInt(digits, base, 64)
	if err != nil {
		panic(err)
	}

	return &LongType{Value: value, Base: uint32(base), Size: uint32(len(digits))}
}

func floatVal(literal string) *FloatType {
	value, err := strconv.ParseFloat(literal, 64)
	if err != nil {
		panic(err)
	}

	result := &FloatType{Value: value}
	mantissa := literal

	if sign := mantissa[0:1]; sign == "+" || sign == "-" {
		result.IntegralSign = proto.String(sign)
		mantissa = mantissa[1:]
	}

	if integral, decimal, found := strings.Cut(mantissa, "."); found {
		result.IntegralSize = uint32(len(integral))

		if significand, exponent, hasExponent := strings.Cut(decimal, "e"); hasExponent {
			result.DecimalSize = uint32(len(significand))

			if sign := exponent[0:1]; sign == "+" || sign == "-" {
				result.ExponentSign = proto.String(sign)
				exponent = exponent[1:]
			}

			result.ExponentSize = uint32(len(exponent))
		} else {
			result.DecimalSize = uint32(len(decimal))
		}
	}

	return result
}

func identVal(value string) *IdentType {
	return &IdentType{Value: value}
}

func strVal(value string) *StringType {
	return &StringType{Value: value}
}

func dataTypeVal(value string) *DataTypeType {
	return &DataTypeType{Value: value}
}

func addrTypeVal(value string) *AddrTypeType {
	return &AddrTypeType{Value: value}
}

func indexOrderVal(value string) *IndexOrderType {
	return &IndexOrderType{Value: value}
}

func numberBase(literal string) (base int, digits string) {
	if strings.HasPrefix(literal, "0x") || strings.HasPrefix(literal, "0X") {
		return 16, literal[2:]
	}

	return 10, literal
}

/*
** Miscellaneous.
 */

type tHelper interface {
	Helper()
}

func markHelper(t assert.TestingT) {
	if helper, ok := t.(tHelper); ok {
		helper.Helper()
	}
}

func isNilNode(node proto.Message) bool {
	if node == nil {
		return true
	}

	value := reflect.ValueOf(node)

	return value.Kind() == reflect.Ptr && value.IsNil()
}

func protoJSON(message proto.Message) string {
	if isNilNode(message) {
		return "<nil>"
	}

	raw, err := protojson.Marshal(message)
	if err != nil {
		return fmt.Sprintf("<%v>", err)
	}

	// Note: protojson deliberately randomizes its whitespace, json.Indent normalizes it back.
	var indented bytes.Buffer
	if err = json.Indent(&indented, raw, "", "  "); err != nil {
		return string(raw)
	}

	return indented.String()
}
