package a2l

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

func (t *IfDataType) MarshalA2L() (result string) {
	return marshalA2L[*IfDataType](t)
}

func (t *IfDataType) A2LTag() *string {
	tag := "IF_DATA"
	return &tag
}

func (t *GenericParameterType) MapChildNodes(node any) {
	switch node.(type) {
	case *GenericNodeType:
		t.Oneof = &GenericParameterType_Generic{Generic: node.(*GenericNodeType)}
	default:
		panic("not implemented yet...")
	}
}

func (t *GenericParameterType) MarshalA2L() (result string) {
	return marshalA2L[*GenericParameterType](t)
}

func (t *GenericParameterType) A2LTag() (result *string) {
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

func (t *GenericNodeType) MarshalA2L() (result string) {
	return marshalA2L[*GenericNodeType](t)
}

func (t *GenericNodeType) A2LTag() *string {
	tag := t.Name.A2LString()
	return &tag
}
