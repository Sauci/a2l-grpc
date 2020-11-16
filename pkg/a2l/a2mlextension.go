package a2l

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

func (t *A2MLType) MarshalA2L() (result string) {
	return marshalA2L[*A2MLType](t)
}

func (t *A2MLType) A2LTag() *string {
	tag := "A2ML"
	return &tag
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

func (t *Declaration) MarshalA2L() (result string) {
	return marshalA2L[*Declaration](t)
}

func (t *Declaration) A2LTag() (result *string) {
	return result
}

func (t *TypeDefinition) MapChildNodes(node any) {
	switch node.(type) {
	case *TypeName:
		t.TypeName = node.(*TypeName)
	default:
		panic("not implemented yet...")
	}
}

func (t *TypeDefinition) MarshalA2L() (result string) {
	return marshalA2L[*TypeDefinition](t)
}

func (t *TypeDefinition) A2LTag() (result *string) {
	return result
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

func (t *TypeName) MarshalA2L() (result string) {
	return marshalA2L[*TypeName](t)
}

func (t *TypeName) A2LTag() (result *string) {
	return result
}

func (t *PredefinedTypeName) MapChildNodes(_ any) {
	panic("not implemented yet...")
}

func (t *PredefinedTypeName) MarshalA2L() (result string) {
	return marshalA2L[*PredefinedTypeName](t)
}

func (t *PredefinedTypeName) A2LTag() (result *string) {
	return result
}

func (t *TypeName_PredefinedTypeName) MapChildNodes(_ any) {
	panic("not implemented yet...")
}

func (t *TypeName_PredefinedTypeName) MarshalA2L() (result string) {
	return result
}

func (t *TypeName_PredefinedTypeName) A2LTag() (result *string) {
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

func (t *TypeName_EnumTypeName) MarshalA2L() (result string) {
	return result
}

func (t *TypeName_EnumTypeName) A2LTag() (result *string) {
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

func (t *TypeName_StructTypeName) MarshalA2L() (result string) {
	return result
}

func (t *TypeName_StructTypeName) A2LTag() (result *string) {
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

func (t *TypeName_TaggedstructTypeName) MarshalA2L() (result string) {
	return result
}

func (t *TypeName_TaggedstructTypeName) A2LTag() (result *string) {
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

func (t *TypeName_TaggedunionTypeName) MarshalA2L() (result string) {
	return result
}

func (t *TypeName_TaggedunionTypeName) A2LTag() (result *string) {
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

func (t *BlockDefinition) MarshalA2L() (result string) {
	return marshalA2L[*BlockDefinition](t)
}

func (t *BlockDefinition) A2LTag() (result *string) {
	return result
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

func (t *EnumTypeName) MarshalA2L() (result string) {
	return marshalA2L[*EnumTypeName](t)
}

func (t *EnumTypeName) A2LTag() (result *string) {
	return result
}

func (t *Enumerator) MapChildNodes(_ any) {
	panic("not implemented yet...")
}

func (t *Enumerator) MarshalA2L() (result string) {
	return marshalA2L[*Enumerator](t)
}

func (t *Enumerator) A2LTag() (result *string) {
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

func (t *StructTypeName) MarshalA2L() (result string) {
	return marshalA2L[*StructTypeName](t)
}

func (t *StructTypeName) A2LTag() (result *string) {
	return result
}

func (t *StructMember) MapChildNodes(node any) {
	switch node.(type) {
	case *Member:
		t.Member = node.(*Member)
	default:
		panic("not implemented yet...")
	}
}

func (t *StructMember) MarshalA2L() (result string) {
	return marshalA2L[*StructMember](t)
}

func (t *StructMember) A2LTag() (result *string) {
	return result
}

func (t *Member) MapChildNodes(node any) {
	switch node.(type) {
	case *TypeName:
		t.TypeName = node.(*TypeName)
	default:
		panic("not implemented yet...")
	}
}

func (t *Member) MarshalA2L() (result string) {
	return marshalA2L[*Member](t)
}

func (t *Member) A2LTag() (result *string) {
	return result
}

func (t *TaggedstructTypeName) MapChildNodes(node any) {
	switch node.(type) {
	case *TaggedstructMember:
		panic("should not happen...")
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

func (t *TaggedstructTypeName) MarshalA2L() (result string) {
	return marshalA2L[*TaggedstructTypeName](t)
}

func (t *TaggedstructTypeName) A2LTag() (result *string) {
	return result
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

func (t *TaggedstructMember) MarshalA2L() (result string) {
	return marshalA2L[*TaggedstructMember](t)
}

func (t *TaggedstructMember) A2LTag() (result *string) {
	return result
}

func (t *TaggedstructMember_TaggedstructDefinition) MapChildNodes(_ any) {
	panic("not implemented yet...")
}

func (t *TaggedstructMember_TaggedstructDefinition) MarshalA2L() (result string) {
	return result
}

func (t *TaggedstructMember_TaggedstructDefinition) A2LTag() (result *string) {
	return result
}

func (t *TaggedstructMember_BlockDefinition) MapChildNodes(_ any) {
	panic("not implemented yet...")
}

func (t *TaggedstructMember_BlockDefinition) MarshalA2L() (result string) {
	return result
}

func (t *TaggedstructMember_BlockDefinition) A2LTag() (result *string) {
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

func (t *TaggedstructDefinition) MarshalA2L() (result string) {
	return marshalA2L[*TaggedstructDefinition](t)
}

func (t *TaggedstructDefinition) A2LTag() (result *string) {
	return result
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

func (t *TaggedunionTypeName) MarshalA2L() (result string) {
	return marshalA2L[*TaggedunionTypeName](t)
}

func (t *TaggedunionTypeName) A2LTag() (result *string) {
	return result
}

func (t *TaggedunionMember) MapChildNodes(node any) {
	switch node.(type) {
	case *BlockDefinition:
		t.Oneof = &TaggedunionMember_BlockDefinition{BlockDefinition: node.(*BlockDefinition)}
	default:
		panic("not implemented yet...")
	}
}

func (t *TaggedunionMember) MarshalA2L() (result string) {
	return marshalA2L[*TaggedunionMember](t)
}

func (t *TaggedunionMember) A2LTag() (result *string) {
	return result
}

func (t *TagMember) MapChildNodes(node any) {
	switch node.(type) {
	case *Member:
		t.Member = node.(*Member)
	default:
		panic("not implemented yet...")
	}
}

func (t *TagMember) MarshalA2L() (result string) {
	return marshalA2L[*TagMember](t)
}

func (t *TagMember) A2LTag() (result *string) {
	return result
}

func (t *TaggedunionMember_TagMember) MapChildNodes(node any) {
	switch node.(type) {
	case *Member:
		t.TagMember.Member = node.(*Member)
	default:
		panic("not implemented yet...")
	}
}

func (t *TaggedunionMember_TagMember) MarshalA2L() (result string) {
	return result
}

func (t *TaggedunionMember_TagMember) A2LTag() (result *string) {
	return result
}
