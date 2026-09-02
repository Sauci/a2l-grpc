package a2l

import (
	"fmt"
	"github.com/sauci/a2l-grpc/pkg/a2l/parser"
	"strconv"
	"strings"
)

func a2lIntToIntType(integerValue parser.IIntegerValueContext) (result *IntType) {
	var err error
	var tmpResult int64
	var base = uint32(0)
	var rawString string

	if integerValue != nil {
		if integerValue.HEX() != nil {
			base = 16
			rawString = integerValue.HEX().GetText()
			rawString = strings.Replace(strings.Replace(rawString, "0X", "", -1), "0x", "", -1)
		} else if integerValue.INT() != nil {
			base = 10
			rawString = integerValue.INT().GetText()
		} else {
			panic(fmt.Errorf("unimplemented int conversion"))
		}

		if tmpResult, err = strconv.ParseInt(rawString, int(base), 64); err == nil {
			result = &IntType{
				Value: int32(tmpResult),
				Base:  base,
				Size:  uint32(len(rawString)),
			}
		} else {
			panic(err)
		}
	}

	return result
}

// parseUint64Literal splits an integer literal into its base and digits and parses it as an
// unsigned 64 bit value. It is shared by the tree builder and by the check which reports an out of
// range literal as a syntax error.
func parseUint64Literal(literal string) (value uint64, base int, digits string, err error) {
	base, digits = 10, literal

	if strings.HasPrefix(digits, "0x") || strings.HasPrefix(digits, "0X") {
		base, digits = 16, digits[2:]
	}

	value, err = strconv.ParseUint(digits, base, 64)

	return value, base, digits, err
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

func a2lLongToLongType(integerValue parser.IIntegerValueContext) (result *LongType) {
	var err error
	var tmpResult int64
	var base = uint32(0)
	var rawString string

	if integerValue != nil {
		if integerValue.HEX() != nil {
			base = 16
			rawString = integerValue.HEX().GetText()
			rawString = strings.Replace(strings.Replace(rawString, "0X", "", -1), "0x", "", -1)
		} else if integerValue.INT() != nil {
			base = 10
			rawString = integerValue.INT().GetText()
		} else {
			panic(fmt.Errorf("unimplemented int conversion"))
		}

		if tmpResult, err = strconv.ParseInt(rawString, int(base), 64); err == nil {
			result = &LongType{
				Value: tmpResult,
				Base:  base,
				Size:  uint32(len(rawString)),
			}
		} else {
			panic(err)
		}
	}

	return result
}

func numericToLongType(integerValue parser.INumericValueContext) (result *LongType) {
	var err error
	var tmpResult int64
	var base = uint32(0)
	var rawString string

	if integerValue != nil {
		if integerValue.HEX() != nil {
			base = 16
			rawString = integerValue.HEX().GetText()
			rawString = strings.Replace(strings.Replace(rawString, "0X", "", -1), "0x", "", -1)
		} else if integerValue.INT() != nil {
			base = 10
			rawString = integerValue.INT().GetText()
		} else {
			panic(fmt.Errorf("unimplemented int conversion"))
		}

		if tmpResult, err = strconv.ParseInt(rawString, int(base), 64); err == nil {
			result = &LongType{
				Value: tmpResult,
				Base:  base,
				Size:  uint32(len(rawString)),
			}
		} else {
			panic(err)
		}
	}

	return result
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
