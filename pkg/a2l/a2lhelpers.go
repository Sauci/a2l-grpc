package a2l

import (
	"fmt"
	"stash.lmb.liebherr.i/tel/gomodparser/pkg/a2l/parser"
	"strconv"
	"strings"
)

func a2lIntToIntType(integerValue parser.IIntegerValueContext) (result *IntType) {
	var err error
	var tmpResult int64
	var base = uint32(0)
	var rawString string

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
			Value: tmpResult,
			Base:  base,
			Size:  uint32(len(rawString)),
		}
	} else {
		panic(err)
	}

	return result
}

func a2lLongToLongType(integerValue parser.IIntegerValueContext) (result *LongType) {
	var err error
	var tmpResult int64
	var base = uint32(0)
	var rawString string

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

	return result
}

func numericToLongType(integerValue parser.INumericValueContext) (result *LongType) {
	var err error
	var tmpResult int64
	var base = uint32(0)
	var rawString string

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

	return result
}

func floatToFloatType(integerValue parser.INumericValueContext) (result *FloatType) {
	var err error
	var tmpResult float64
	var split []string
	var integralSign string
	var exponentSign string

	result = &FloatType{}

	processedString := integerValue.GetText()
	originalString := integerValue.GetText()

	if strings.HasPrefix(processedString, "+") {
		integralSign = "+"
		result.IntegralSign = &integralSign
		processedString = strings.TrimLeft(processedString, "+")
	} else if strings.HasPrefix(processedString, "-") {
		integralSign = "-"
		result.IntegralSign = &integralSign
		processedString = strings.TrimLeft(processedString, "-")
	}

	if strings.Contains(processedString, ".") {
		split = strings.Split(processedString, ".")
		result.IntegralSize = uint32(len(split[0]))

		if len(split) > 1 {
			decimalString := split[1]
			if strings.Contains(decimalString, "e") {
				split = strings.Split(decimalString, "e")
				result.DecimalSize = uint32(len(split[0]))
				exponentString := split[1]

				if strings.HasPrefix(exponentString, "+") {
					exponentSign = "+"
					result.ExponentSign = &exponentSign
					exponentString = strings.TrimLeft(exponentString, "+")
				} else if strings.HasPrefix(exponentString, "-") {
					exponentSign = "-"
					result.ExponentSign = &exponentSign
					exponentString = strings.TrimLeft(exponentString, "-")
				}

				result.ExponentSize = uint32(len(exponentString))
			} else {
				result.DecimalSize = uint32(len(split[1]))
			}
		}
	}

	if tmpResult, err = strconv.ParseFloat(originalString, 64); err == nil {
		result.Value = tmpResult
	} else {
		panic(err)
	}

	return result
}

func identifierToIdentType(identifierValue parser.IIdentifierValueContext) (result *IdentType) {
	return &IdentType{Value: identifierValue.GetText()}
}

func a2lStringToStringType(stringValue parser.IStringValueContext) (result *StringType) {
	rawString := stringValue.GetText()

	return &StringType{Value: rawString[1 : len(rawString)-1]}
}

//func a2mlTagToString(tagValue parser.ITagValueContext) (result tagType) {
//	return tagType(tagValue.GetText())
//}
//
//func a2mlIdentifierToString(identifierValue parser.IIdentifierValueContext) (result IdentType) {
//	return IdentType{Value: identifierValue.GetText()}
//}
//
//func a2mlStringToString(stringValue parser.IStringValueContext) (result StringType) {
//	rawString := stringValue.GetText()
//	return StringType{Value: rawString[1 : len(rawString)-1]}
//}
