package a2l

import (
	"fmt"
	"strings"
)

func (t *A2MLType) MapChildNodes(node any) {
	switch node.(type) {
	case *Declaration:
		if t.Declaration == nil {
			t.Declaration = make([]*Declaration, 0)
		}

		t.Declaration = append(t.Declaration, node.(*Declaration))
	case *Declaration_TypeDefinition:
		if t.Declaration == nil {
			t.Declaration = make([]*Declaration, 0)
		}

		t.Declaration = append(t.Declaration, &Declaration{Oneof: node.(*Declaration_TypeDefinition)})
	default:
		panic("not implemented yet...")
	}
}

func (t *A2MLType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent("/begin A2ML", indentLevel, indentString)}

	if t.Declaration != nil {
		for _, declaration := range t.Declaration {
			tmpResult = append(tmpResult, declaration.MarshalA2L(indentLevel+1, indentString))
		}
	}

	tmpResult = append(tmpResult, indentContent("/end A2ML", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *Declaration) MapChildNodes(node any) {
	switch node.(type) {
	case *TypeDefinition:
		t.Oneof = &Declaration_TypeDefinition{TypeDefinition: node.(*TypeDefinition)}
	case *BlockDefinition:
		t.Oneof = &Declaration_BlockDefinition{BlockDefinition: node.(*BlockDefinition)}
	default:
		panic("not implemented yet...")
	}
}

func (t *Declaration) MarshalA2L(indentLevel int, indentString string) (result string) {
	switch t.Oneof.(type) {
	case *Declaration_TypeDefinition:
		result = t.Oneof.(*Declaration_TypeDefinition).TypeDefinition.MarshalA2L(indentLevel, indentString)
	case *Declaration_BlockDefinition:
		result = t.Oneof.(*Declaration_BlockDefinition).BlockDefinition.MarshalA2L(indentLevel, indentString)
	default:
		panic("not implemented yet...")
	}

	return result + ";"
}

func (t *TypeDefinition) MapChildNodes(node any) {
	switch node.(type) {
	case *TypeName:
		t.TypeName = node.(*TypeName)
	default:
		panic("not implemented yet...")
	}
}

func (t *TypeDefinition) MarshalA2L(indentLevel int, indentString string) (result string) {
	return t.TypeName.MarshalA2L(indentLevel, indentString)
}

func (t *TypeName) MapChildNodes(node any) {
	switch node.(type) {
	case *PredefinedTypeName:
		panic("not implemented yet...")
	case *TypeName_PredefinedTypeName:
		t.Oneof = node.(*TypeName_PredefinedTypeName)
	case *EnumTypeName:
		panic("not implemented yet...")
	case *TypeName_EnumTypeName:
		t.Oneof = node.(*TypeName_EnumTypeName)
	case *StructTypeName:
		panic("not implemented yet...")
	case *TypeName_StructTypeName:
		t.Oneof = node.(*TypeName_StructTypeName)
	case *TaggedstructTypeName:
		panic("not implemented yet...")
	case *TypeName_TaggedstructTypeName:
		t.Oneof = node.(*TypeName_TaggedstructTypeName)
	case *TypeName_TaggedunionTypeName:
		t.Oneof = node.(*TypeName_TaggedunionTypeName)
	default:
		panic("not implemented yet...")
	}
}

func (t *TypeName) MarshalA2L(indentLevel int, indentString string) (result string) {
	switch t.Oneof.(type) {
	case *TypeName_PredefinedTypeName:
		result = t.Oneof.(*TypeName_PredefinedTypeName).PredefinedTypeName.MarshalA2L(indentLevel, indentString)
	case *TypeName_StructTypeName:
		result = t.Oneof.(*TypeName_StructTypeName).StructTypeName.MarshalA2L(indentLevel, indentString)
	case *TypeName_TaggedstructTypeName:
		result = t.Oneof.(*TypeName_TaggedstructTypeName).TaggedstructTypeName.MarshalA2L(indentLevel, indentString)
	case *TypeName_TaggedunionTypeName:
		result = t.Oneof.(*TypeName_TaggedunionTypeName).TaggedunionTypeName.MarshalA2L(indentLevel, indentString)
	case *TypeName_EnumTypeName:
		result = t.Oneof.(*TypeName_EnumTypeName).EnumTypeName.MarshalA2L(indentLevel, indentString)
	default:
		panic("not implemented yet...")
	}

	return result
}

func (t *PredefinedTypeName) MapChildNodes(_ any) {
	panic("not implemented yet...")
}

func (t *PredefinedTypeName) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(t.Name, indentLevel, indentString)
}

func (t *TypeName_PredefinedTypeName) MapChildNodes(_ any) {
	panic("not implemented yet...")
}

func (t *TypeName_PredefinedTypeName) MarshalA2L(_ int, _ string) (result string) {
	return result
}

func (t *TypeName_EnumTypeName) MapChildNodes(node any) {
	switch node.(type) {
	case *Enumerator:
		if t.EnumTypeName.EnumeratorList == nil {
			t.EnumTypeName.EnumeratorList = make([]*Enumerator, 0)
		}

		t.EnumTypeName.EnumeratorList = append(t.EnumTypeName.EnumeratorList, node.(*Enumerator))
	default:
		panic("not implemented yet...")
	}
}

func (t *TypeName_EnumTypeName) MarshalA2L(_ int, _ string) (result string) {
	return result
}

func (t *TypeName_StructTypeName) MapChildNodes(node any) {
	switch node.(type) {
	case *StructMember:
		if t.StructTypeName.StructMemberList == nil {
			t.StructTypeName.StructMemberList = make([]*StructMember, 0)
		}

		t.StructTypeName.StructMemberList = append(t.StructTypeName.StructMemberList, node.(*StructMember))
	default:
		panic("not implemented yet...")
	}
}

func (t *TypeName_StructTypeName) MarshalA2L(_ int, _ string) (result string) {
	return result
}

func (t *TypeName_TaggedstructTypeName) MapChildNodes(node any) {
	switch node.(type) {
	case *TaggedstructMember_TaggedstructDefinition:
		if t.TaggedstructTypeName.TaggedstructMemberList == nil {
			t.TaggedstructTypeName.TaggedstructMemberList = make([]*TaggedstructMember, 0)
		}

		// TODO: handle this case properly...
		t.TaggedstructTypeName.TaggedstructMemberList = append(t.TaggedstructTypeName.TaggedstructMemberList, &TaggedstructMember{
			Oneof: node.(*TaggedstructMember_TaggedstructDefinition),
		})
	case *TaggedstructMember:
		if t.TaggedstructTypeName.TaggedstructMemberList == nil {
			t.TaggedstructTypeName.TaggedstructMemberList = make([]*TaggedstructMember, 0)
		}

		t.TaggedstructTypeName.TaggedstructMemberList = append(t.TaggedstructTypeName.TaggedstructMemberList, node.(*TaggedstructMember))
	default:
		panic("not implemented yet...")
	}
}

func (t *TypeName_TaggedstructTypeName) MarshalA2L(_ int, _ string) (result string) {
	return result
}

func (t *TypeName_TaggedunionTypeName) MapChildNodes(node any) {
	switch node.(type) {
	case *TaggedunionMember:
		if t.TaggedunionTypeName.TaggedunionMemberList == nil {
			t.TaggedunionTypeName.TaggedunionMemberList = make([]*TaggedunionMember, 0)
		}

		t.TaggedunionTypeName.TaggedunionMemberList = append(t.TaggedunionTypeName.TaggedunionMemberList, node.(*TaggedunionMember))
	case *TaggedunionMember_TagMember:
		if t.TaggedunionTypeName.TaggedunionMemberList == nil {
			t.TaggedunionTypeName.TaggedunionMemberList = make([]*TaggedunionMember, 0)
		}

		t.TaggedunionTypeName.TaggedunionMemberList = append(t.TaggedunionTypeName.TaggedunionMemberList, &TaggedunionMember{
			Oneof: node.(*TaggedunionMember_TagMember),
		})
	default:
		panic("not implemented yet...")
	}
}

func (t *TypeName_TaggedunionTypeName) MarshalA2L(_ int, _ string) (result string) {
	return result
}

func (t *BlockDefinition) MapChildNodes(node any) {
	switch node.(type) {
	case *TypeName:
		t.Oneof = &BlockDefinition_TypeName{TypeName: node.(*TypeName)}
	case *Member:
		t.Oneof = &BlockDefinition_Member{Member: node.(*Member)}
	default:
		panic("not implemented yet...")
	}
}

func (t *BlockDefinition) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{fmt.Sprintf(indentContent("block ", indentLevel, indentString))}

	if t.Tag != nil {
		tmpResult[0] += fmt.Sprintf("%s ", t.Tag.A2LString())
	}

	switch t.Oneof.(type) {
	case *BlockDefinition_TypeName:
		tmpResult = append(tmpResult, t.Oneof.(*BlockDefinition_TypeName).TypeName.MarshalA2L(indentLevel+1, indentString))
	case *BlockDefinition_Member:
		tmpResult = append(tmpResult, t.Oneof.(*BlockDefinition_Member).Member.MarshalA2L(indentLevel+1, indentString))
	default:
		panic("not implemented yet...")
	}

	return strings.Join(tmpResult, "\n")
}

func (t *EnumTypeName) MapChildNodes(node any) {
	switch node.(type) {
	case *Enumerator:
		if t.EnumeratorList == nil {
			t.EnumeratorList = make([]*Enumerator, 0)
		}

		t.EnumeratorList = append(t.EnumeratorList, node.(*Enumerator))
	default:
		panic("not implemented yet...")
	}
}

func (t *EnumTypeName) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent("enum ", indentLevel, indentString)}

	if t.Identifier != nil {
		tmpResult[0] += fmt.Sprintf("%s ", t.Identifier.A2LString())
	}

	if t.EnumeratorList != nil {
		tmpResult[0] += "{"

		for i, enumerator := range t.EnumeratorList {
			tmpResult = append(tmpResult, indentContent(enumerator.MarshalA2L(indentLevel, indentString), indentLevel+1, indentString))

			if i < len(t.EnumeratorList)-1 {
				tmpResult[len(tmpResult)-1] += ","
			}
		}

		tmpResult = append(tmpResult, indentContent("}", indentLevel, indentString))
	}

	return strings.Join(tmpResult, "\n")
}

func (t *Enumerator) MapChildNodes(_ any) {
	panic("not implemented yet...")
}

func (t *Enumerator) MarshalA2L(_ int, _ string) (result string) {
	result = t.Keyword.A2LString()

	if t.Constant != nil {
		result += fmt.Sprintf(" = %s", t.Constant.A2LString())
	}

	return result
}

func (t *StructTypeName) MapChildNodes(node any) {
	switch node.(type) {
	case *StructMember:
		if t.StructMemberList == nil {
			t.StructMemberList = make([]*StructMember, 0)
		}

		t.StructMemberList = append(t.StructMemberList, node.(*StructMember))
	default:
		panic("not implemented yet...")
	}
}

func (t *StructTypeName) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent("struct ", indentLevel, indentString)}

	if t.Identifier != nil {
		tmpResult[0] += fmt.Sprintf("%s ", t.Identifier.A2LString())
	}

	if t.StructMemberList != nil {
		tmpResult[0] += "{"

		for _, structMember := range t.StructMemberList {
			tmpResult = append(tmpResult, structMember.MarshalA2L(indentLevel+1, indentString))
		}

		tmpResult = append(tmpResult, indentContent("}", indentLevel, indentString))
	}

	return strings.Join(tmpResult, "\n")
}

func (t *StructMember) MapChildNodes(node any) {
	switch node.(type) {
	case *Member:
		t.Member = node.(*Member)
	default:
		panic("not implemented yet...")
	}
}

func (t *StructMember) MarshalA2L(indentLevel int, indentString string) (result string) {
	return fmt.Sprintf("%s;", t.Member.MarshalA2L(indentLevel, indentString))
}

func (t *Member) MapChildNodes(node any) {
	switch node.(type) {
	case *TypeName:
		t.TypeName = node.(*TypeName)
	default:
		panic("not implemented yet...")
	}
}

func (t *Member) MarshalA2L(indentLevel int, indentString string) (result string) {
	result = t.TypeName.MarshalA2L(indentLevel, indentString)

	if t.ArraySpecifier != nil {
		result += "["

		for _, constant := range t.ArraySpecifier.Constant {
			result += constant.A2LString()
		}

		result += "]"
	}

	return result
}

func (t *TaggedstructTypeName) MapChildNodes(node any) {
	switch node.(type) {
	case *TaggedstructMember_TaggedstructDefinition:
		if t.TaggedstructMemberList == nil {
			t.TaggedstructMemberList = make([]*TaggedstructMember, 0)
		}

		t.TaggedstructMemberList = append(t.TaggedstructMemberList, &TaggedstructMember{
			Oneof: node.(*TaggedstructMember_TaggedstructDefinition),
		})
	default:
		panic("not implemented yet...")
	}
}

func (t *TaggedstructTypeName) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent("taggedstruct ", indentLevel, indentString)}

	if t.Identifier != nil {
		tmpResult[0] += fmt.Sprintf("%s ", t.Identifier.A2LString())
	}

	if t.TaggedstructMemberList != nil {
		tmpResult[0] += "{"

		for _, taggedStructMember := range t.TaggedstructMemberList {
			tmpResult = append(tmpResult, taggedStructMember.MarshalA2L(indentLevel+1, indentString))
		}

		tmpResult = append(tmpResult, indentContent("}", indentLevel, indentString))
	}

	return strings.Join(tmpResult, "\n")
}

func (t *TaggedstructMember) MapChildNodes(node any) {
	switch node.(type) {
	case *TaggedstructDefinition:
		// TODO: handle this case properly...
		t.Oneof = &TaggedstructMember_TaggedstructDefinition{
			TaggedstructDefinition: node.(*TaggedstructDefinition),
		}
	case *TaggedstructMember_TaggedstructDefinition:
		t.Oneof = node.(*TaggedstructMember_TaggedstructDefinition)
	case *BlockDefinition:
		t.Oneof = &TaggedstructMember_BlockDefinition{
			BlockDefinition: node.(*BlockDefinition),
		}
	default:
		panic("not implemented yet...")
	}
}

func (t *TaggedstructMember) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := make([]string, 0)

	if t.Star {
		tmpResult = append(tmpResult, indentContent("(", indentLevel+1, indentString))
	}

	switch t.Oneof.(type) {
	case *TaggedstructMember_TaggedstructDefinition:
		tmpResult = append(tmpResult, t.Oneof.(*TaggedstructMember_TaggedstructDefinition).TaggedstructDefinition.MarshalA2L(indentLevel+1, indentString))
	case *TaggedstructMember_BlockDefinition:
		tmpResult = append(tmpResult, t.Oneof.(*TaggedstructMember_BlockDefinition).BlockDefinition.MarshalA2L(indentLevel+1, indentString))
	default:
		panic("not implemented yet...")
	}

	if t.Star {
		tmpResult = append(tmpResult, indentContent(")*;", indentLevel+1, indentString))
	} else {
		tmpResult[len(tmpResult)-1] += ";"
	}

	return strings.Join(tmpResult, "\n")
}

func (t *TaggedstructMember_TaggedstructDefinition) MapChildNodes(_ any) {
	panic("not implemented yet...")
}

func (t *TaggedstructMember_TaggedstructDefinition) MarshalA2L(_ int, _ string) (result string) {
	return result
}

func (t *TaggedstructMember_BlockDefinition) MapChildNodes(_ any) {
	panic("not implemented yet...")
}

func (t *TaggedstructMember_BlockDefinition) MarshalA2L(_ int, _ string) (result string) {
	return result
}

func (t *TaggedstructDefinition) MapChildNodes(node any) {
	switch node.(type) {
	case *Member:
		t.Member = node.(*Member)
	default:
		panic("not implemented yet...")
	}
}

func (t *TaggedstructDefinition) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("%s", t.Tag.A2LString()), indentLevel, indentString)}

	if t.Star {
		tmpResult = append(tmpResult, indentContent("(", indentLevel, indentString))
	}

	if t.Member != nil { // not part of the spec...
		tmpResult = append(tmpResult, t.Member.MarshalA2L(indentLevel, indentString))
	}

	if t.Star {
		tmpResult = append(tmpResult, indentContent(")*", indentLevel, indentString))
	}

	return strings.Join(tmpResult, "\n")
}

func (t *TaggedunionTypeName) MapChildNodes(node any) {
	switch node.(type) {
	case *TaggedunionMember:
		if t.TaggedunionMemberList == nil {
			t.TaggedunionMemberList = make([]*TaggedunionMember, 0)
		}

		t.TaggedunionMemberList = append(t.TaggedunionMemberList, node.(*TaggedunionMember))
	case *TagMember:
		if t.TaggedunionMemberList == nil {
			t.TaggedunionMemberList = make([]*TaggedunionMember, 0)
		}

		t.TaggedunionMemberList = append(t.TaggedunionMemberList, &TaggedunionMember{Oneof: node.(*TaggedunionMember_TagMember)})
	case *TaggedunionMember_TagMember:
		if t.TaggedunionMemberList == nil {
			t.TaggedunionMemberList = make([]*TaggedunionMember, 0)
		}

		t.TaggedunionMemberList = append(t.TaggedunionMemberList, &TaggedunionMember{Oneof: node.(*TaggedunionMember_TagMember)})
	default:
		panic("not implemented yet...")
	}
}

func (t *TaggedunionTypeName) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent("taggedunion ", indentLevel, indentString)}

	if t.Identifier != nil {
		tmpResult[0] += fmt.Sprintf("%s ", t.Identifier.A2LString())
	}

	if t.TaggedunionMemberList != nil {
		tmpResult[0] += "{"

		for _, taggedunionMember := range t.TaggedunionMemberList {
			tmpResult = append(tmpResult, taggedunionMember.MarshalA2L(indentLevel, indentString))
		}

		tmpResult = append(tmpResult, indentContent("}", indentLevel, indentString))
	}

	return strings.Join(tmpResult, "\n")
}

func (t *TaggedunionMember) MapChildNodes(node any) {
	switch node.(type) {
	case *BlockDefinition:
		t.Oneof = &TaggedunionMember_BlockDefinition{BlockDefinition: node.(*BlockDefinition)}
	default:
		panic("not implemented yet...")
	}
}

func (t *TaggedunionMember) MarshalA2L(indentLevel int, indentString string) (result string) {
	switch t.Oneof.(type) {
	case *TaggedunionMember_TagMember:
		result = t.Oneof.(*TaggedunionMember_TagMember).TagMember.MarshalA2L(indentLevel+1, indentString)
	case *TaggedunionMember_BlockDefinition:
		result = t.Oneof.(*TaggedunionMember_BlockDefinition).BlockDefinition.MarshalA2L(indentLevel+1, indentString)
	default:
		panic("not implemented yet...")
	}

	return result + ";"
}

func (t *TagMember) MapChildNodes(node any) {
	switch node.(type) {
	case *Member:
		t.Member = node.(*Member)
	default:
		panic("not implemented yet...")
	}
}

func (t *TagMember) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := make([]string, 0)
	if t.Tag != nil {
		tmpResult = append(tmpResult, indentContent(t.Tag.A2LString(), indentLevel, indentString))
	}

	tmpResult = append(tmpResult, t.Member.MarshalA2L(indentLevel+1, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *TaggedunionMember_TagMember) MapChildNodes(node any) {
	switch node.(type) {
	case *Member:
		t.TagMember.Member = node.(*Member)
	default:
		panic("not implemented yet...")
	}
}

func (t *TaggedunionMember_TagMember) MarshalA2L(_ int, _ string) (result string) {
	return result
}
