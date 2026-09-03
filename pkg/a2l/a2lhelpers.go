package a2l

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/sauci/a2l-grpc/pkg/a2l/parser"
	"strconv"
	"strings"
)

// literalContext is the smallest surface of a parser context needed to convert a numeric literal;
// it is satisfied by both parser.IIntegerValueContext and parser.INumericValueContext.
type literalContext interface {
	GetText() string
	GetStart() antlr.Token
}

// splitIntegerLiteral separates the base of an integer literal from its digits. The notation of a
// hexadecimal value is fixed by the specification (ASAM MCD-2 MC 1.6.1, chapter 3.2), so only the
// leading 0x or 0X is a base marker.
func splitIntegerLiteral(literal string) (base int, digits string) {
	if strings.HasPrefix(literal, "0x") || strings.HasPrefix(literal, "0X") {
		return 16, literal[2:]
	}

	return 10, literal
}

// parseUint64Literal parses an integer literal as an unsigned 64 bit value. It is shared by the
// tree builder and by the check which reports an out of range literal as a syntax error.
func parseUint64Literal(literal string) (value uint64, base int, digits string, err error) {
	base, digits = splitIntegerLiteral(literal)

	value, err = strconv.ParseUint(digits, base, 64)

	return value, base, digits, err
}

// parseSignedLiteral converts an integer literal to a signed value of the passed width. A literal
// which does not fit that width is rejected: the widths are the ones the specification declares
// for the parameters (chapter 3.2, "Predefined data types"), and a value which was silently
// wrapped around could not be serialized back to A2L either.
func parseSignedLiteral(literal literalContext, bitSize int) (value int64, base int, digits string) {
	var err error

	base, digits = splitIntegerLiteral(literal.GetText())

	if value, err = strconv.ParseInt(digits, base, bitSize); err != nil {
		token := literal.GetStart()

		panic(fmt.Errorf("%v:%v %v: %v", token.GetLine(), token.GetColumn(), literal.GetText(), err))
	}

	return value, base, digits
}

// a2lIntToIntType converts the parameters which the specification declares as int or uint. Both
// are 16 bit wide (chapter 3.2), so the 32 bit width of the node is already permissive; a wider
// literal is reported rather than truncated.
func a2lIntToIntType(integerValue parser.IIntegerValueContext) (result *IntType) {
	if integerValue == nil {
		return nil
	}

	value, base, digits := parseSignedLiteral(integerValue, 32)

	return &IntType{Value: int32(value), Base: uint32(base), Size: uint32(len(digits))}
}

// a2lULongToULongType converts the parameters which ASAM MCD-2 MC 1.6.1 declares as uint64. An
// out of range literal yields no node: it is already reported as a syntax error, so the tree is
// discarded anyway, and returning nil keeps the conversion free of panics.
func a2lULongToULongType(integerValue parser.IIntegerValueContext) (result *ULongType) {
	if integerValue == nil {
		return nil
	}

	value, base, digits, err := parseUint64Literal(integerValue.GetText())
	if err != nil {
		return nil
	}

	return &ULongType{Value: value, Base: uint32(base), Size: uint32(len(digits))}
}

// a2lLongToLongType converts the parameters which the specification declares as long or ulong.
// Both fit the 64 bit width of the node, an unsigned 32 bit address such as 0xFFFFFFFF included.
func a2lLongToLongType(integerValue parser.IIntegerValueContext) (result *LongType) {
	if integerValue == nil {
		return nil
	}

	value, base, digits := parseSignedLiteral(integerValue, 64)

	return &LongType{Value: value, Base: uint32(base), Size: uint32(len(digits))}
}

func numericToLongType(integerValue parser.INumericValueContext) (result *LongType) {
	if integerValue == nil {
		return nil
	}

	value, base, digits := parseSignedLiteral(integerValue, 64)

	return &LongType{Value: value, Base: uint32(base), Size: uint32(len(digits))}
}

// floatTextContext is the smallest surface of a parser context needed to convert a float value; it
// is satisfied by both parser.IFloatValueContext and parser.INumericValueContext.
type floatTextContext interface {
	GetText() string
}

func floatToFloatType(integerValue floatTextContext) (result *FloatType) {
	var err error
	var tmpResult float64

	if integerValue != nil {
		rawString := integerValue.GetText()

		if tmpResult, err = strconv.ParseFloat(rawString, 64); err == nil {
			result = &FloatType{Value: tmpResult, Source: rawString}
		} else {
			panic(err)
		}
	}

	return result
}

func identifierToIdentType(identifierValue parser.IIdentifierValueContext) (result *IdentType) {
	if identifierValue != nil {
		result = &IdentType{Value: identifierValue.GetText()}
	}

	return result
}

func a2lStringToStringType(stringValue parser.IStringValueContext) (result *StringType) {
	if stringValue != nil {
		rawString := stringValue.GetText()
		result = &StringType{Value: rawString[1 : len(rawString)-1]}
	}

	return result
}
