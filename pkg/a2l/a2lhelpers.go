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
