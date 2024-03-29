package a2l

import (
	"fmt"
	"strings"
)

func (t *IfDataType) MapChildNodes(node any) {
	switch node.(type) {
	case *GenericParameterType:
		if t.Blob == nil {
			t.Blob = make([]*GenericParameterType, 0)
		}

		t.Blob = append(t.Blob, node.(*GenericParameterType))
	default:
		panic("not implemented yet...")
	}
}

func (t *IfDataType) MarshalA2L(indentLevel int, indentString string, sort bool) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin IF_DATA %s", t.Name.A2LString()), indentLevel, indentString)}

	if t.Blob != nil {
		for _, blob := range t.Blob {
			tmpResult = append(tmpResult, blob.MarshalA2L(indentLevel+1, indentString, sort))
		}
	}

	tmpResult = append(tmpResult, indentContent("/end IF_DATA", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *GenericParameterType) MapChildNodes(node any) {
	switch node.(type) {
	case *GenericNodeType:
		t.Oneof = &GenericParameterType_Generic{Generic: node.(*GenericNodeType)}
	default:
		panic("not implemented yet...")
	}
}

func (t *GenericParameterType) MarshalA2L(indentLevel int, indentString string, sort bool) (result string) {
	switch t.Oneof.(type) {
	case *GenericParameterType_Tag:
		result = indentContent(t.Oneof.(*GenericParameterType_Tag).Tag.A2LString(), indentLevel, indentString)
	case *GenericParameterType_String_:
		result = indentContent(t.Oneof.(*GenericParameterType_String_).String_.A2LString(), indentLevel, indentString)
	case *GenericParameterType_Long:
		result = indentContent(t.Oneof.(*GenericParameterType_Long).Long.A2LString(), indentLevel, indentString)
	case *GenericParameterType_Float:
		result = indentContent(t.Oneof.(*GenericParameterType_Float).Float.A2LString(), indentLevel, indentString)
	case *GenericParameterType_Generic:
		result = t.Oneof.(*GenericParameterType_Generic).Generic.MarshalA2L(indentLevel, indentString, sort)
	case *GenericParameterType_Identifier:
		result = indentContent(t.Oneof.(*GenericParameterType_Identifier).Identifier.A2LString(), indentLevel, indentString)
	default:
		panic("not implemented yet...")
	}

	return result
}

func (t *GenericNodeType) MapChildNodes(node any) {
	switch node.(type) {
	case *GenericParameterType:
		if t.Element == nil {
			t.Element = make([]*GenericParameterType, 0)
		}

		t.Element = append(t.Element, node.(*GenericParameterType))
	default:
		panic("not implemented yet...")
	}
}

func (t *GenericNodeType) MarshalA2L(indentLevel int, indentString string, sort bool) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin %s", t.Name.A2LString()), indentLevel, indentString)}

	if t.Element != nil {
		for _, element := range t.Element {
			tmpResult = append(tmpResult, element.MarshalA2L(indentLevel+1, indentString, sort))
		}
	}

	tmpResult = append(tmpResult, indentContent(fmt.Sprintf("/end %s", t.Name.A2LString()), indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}
