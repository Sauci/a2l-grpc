package a2l

import (
	"fmt"
	"reflect"
	"strings"
)

var indentLevel = 0

type IIndentedNode interface {
	MapChildNodes(node any)
	MarshalA2L() (result string)
	A2LTag() (result *string)
}

func indentContent(content string, level int) (result string) {
	for i := 0; i < level; i++ {
		result += "  "
	}

	return result + content
}

func marshalA2L[T any](t T) (result string) {
	value := reflect.ValueOf(t).Elem()
	var tag *string

	if value.IsValid() {
		numOfFields := value.NumField()

		if _, ok := reflect.ValueOf(t).Interface().(IIndentedNode); ok {
			tag = reflect.ValueOf(t).Interface().(IIndentedNode).A2LTag()

			if tag != nil {
				result += indentContent(fmt.Sprintf("/begin"), indentLevel)
				indentLevel++
			}
		}

		for i := 0; i < numOfFields; i++ {
			jsonTag := reflect.TypeOf(t).Elem().Field(i).Tag.Get("json")

			if jsonTag != "" {
				jsonTag = strings.Split(jsonTag, ",")[0]

				if _, ok := value.Field(i).Interface().(IIndentedNode); ok {
					//tag := value.Field(i).Interface().(IIndentedNode).A2LTag()

					//if tag != nil {
					//	result += indentContent(fmt.Sprintf("/begin"), indentLevel)
					//	indentLevel++
					//}

					result += indentContent(fmt.Sprintf("%s %s", jsonTag, value.Field(i).Interface().(IIndentedNode).MarshalA2L()), indentLevel)

					//if tag != nil {
					//	indentLevel--
					//	result += indentContent(fmt.Sprintf("/end %s", *tag), indentLevel)
					//}
				} else if reflect.TypeOf(t).Elem().Field(i).Type.Kind() == reflect.Slice {
					s := value.Field(i)

					for i := 0; i < s.Len(); i++ {
						e := s.Index(i).Interface()
						result += indentContent(marshalA2L(e), indentLevel)
					}
				} else if _, ok := value.Field(i).Interface().(A2LStringer); ok {
					result += indentContent(fmt.Sprintf("%s 2\n", value.Field(i).Interface().(A2LStringer).A2LString()), indentLevel)
				} else {
					result += fmt.Sprintf(" %v 3\n", value.Field(i).Interface())
				}
			}
		}

		if tag != nil {
			indentLevel--
			result += indentContent(fmt.Sprintf("/end %s", *tag), indentLevel)
		}
	}

	return result
}
