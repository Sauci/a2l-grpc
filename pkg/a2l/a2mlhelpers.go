package a2l

import (
	"github.com/sauci/a2l-grpc/pkg/a2l/parser"
)

func stringToTagType(stringValue parser.IStringValueContext) (result *TagType) {
	if stringValue == nil {
		return nil
	}

	rawString := stringValue.GetText()

	return &TagType{Value: rawString[1 : len(rawString)-1]}
}

// a2mlIdentifierToIdentType converts an identifier of the metalanguage. It may be an ordinary
// identifier or a keyword of the A2L grammar, which chapter 5.2 allows to be reused inside A2ML
// and IF_DATA; the text of the node is the identifier in both cases.
func a2mlIdentifierToIdentType(identifier parser.IA2mlIdentifierContext) (result *IdentType) {
	if identifier == nil {
		return nil
	}

	return &IdentType{Value: identifier.GetText()}
}

// ifDataIdentifierToIdentType converts an identifier used inside an IF_DATA block.
func ifDataIdentifierToIdentType(identifier parser.IIfDataIdentifierContext) (result *IdentType) {
	if identifier == nil {
		return nil
	}

	return &IdentType{Value: identifier.GetText()}
}

// arraySpecifierToLongType converts the constant of an A2ML array specifier (chapter 5.2).
func arraySpecifierToLongType(arraySpecifier parser.IA2mlArraySpecifierContext) (result *LongType) {
	if arraySpecifier == nil {
		return nil
	}

	return a2lLongToLongType(arraySpecifier.GetValue())
}
