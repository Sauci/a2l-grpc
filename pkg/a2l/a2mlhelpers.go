package a2l

import (
	"fmt"
	"github.com/sauci/a2l-grpc/pkg/a2l/parser"
	"strconv"
)

func stringToTagType(stringValue parser.IStringValueContext) (result *TagType) {
	rawString := stringValue.GetText()

	return &TagType{Value: rawString[1 : len(rawString)-1]}
}

func tagToTagType(tagValue parser.ITagValueContext) (result *TagType) {
	rawString := tagValue.GetText()

	return &TagType{Value: rawString[1 : len(rawString)-1]}
}

func arraySpecifierToLongType(integerValue parser.IArraySpecifierContext) (result *LongType) {
	var err error
	var tmpResult int64
	var base = uint32(0)
	var rawString string

	if integerValue.IDENT() != nil {
		panic("not implemented yet...")
	} else if integerValue.INT() != nil {
		base = 10
		rawString = integerValue.INT().GetText()
	} else {
		panic(fmt.Errorf("unimplemented array specifier conversion"))
	}

	if tmpResult, err = strconv.ParseInt(rawString, int(base), 64); err == nil {
		result = &LongType{
			Value: int32(tmpResult),
			Base:  base,
			Size:  uint32(len(rawString)),
		}
	} else {
		panic(err)
	}

	return result
}
