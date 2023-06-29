package a2l

type IIndentedNode interface {
	MapChildNodes(node any)
	MarshalA2L(indent int) (result string)
}

func indentContent(content string, level int) (result string) {
	indent := ""
	//lines := make([]string, 0)
	for i := 0; i < level; i++ {
		indent += "  "
	}
	//
	//for _, line := range strings.Split(content, "\n") {
	//	lines = append(lines, indent+line)
	//}
	//
	//return strings.Join(lines, "\n")
	return indent + content
}
