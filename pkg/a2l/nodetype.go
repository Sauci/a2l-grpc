package a2l

type IIndentedNode interface {
	MapChildNodes(node any)
	MarshalA2L(indentLevel int, indentString string, sort bool) (result string)
}

type SortableNode interface {
	UniqueKey() string
}

func indentContent(content string, indentLevel int, indentString string) (result string) {
	indent := ""

	for i := 0; i < indentLevel; i++ {
		indent += indentString
	}

	return indent + content
}
