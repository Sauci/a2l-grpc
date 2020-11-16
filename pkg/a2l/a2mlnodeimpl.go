package a2l

import (
	"stash.lmb.liebherr.i/tel/gomodparser/pkg/a2l/parser"
)

func (n *Listener) EnterA2ml(_ *parser.A2mlContext) {
	a2ml := &A2MLType{}
	n.Push(a2ml)
}

func (n *Listener) ExitA2ml(_ *parser.A2mlContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterDeclaration(_ *parser.DeclarationContext) {
	declaration := &Declaration{}

	n.Push(declaration)
}

func (n *Listener) ExitDeclaration(_ *parser.DeclarationContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterTypeDefinition(_ *parser.TypeDefinitionContext) {
	typeDefinition := &TypeDefinition{}
	n.Push(typeDefinition)
}

func (n *Listener) ExitTypeDefinition(_ *parser.TypeDefinitionContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterA2mlTypeName(_ *parser.A2mlTypeNameContext) {
	typeName := &TypeName{}
	n.Push(typeName)
}

func (n *Listener) ExitA2mlTypeName(_ *parser.A2mlTypeNameContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterPredefinedTypeName(ctx *parser.PredefinedTypeNameContext) {
	predefinedTypeName := &TypeName_PredefinedTypeName{
		PredefinedTypeName: &PredefinedTypeName{Name: ctx.GetName().GetText()},
	}
	n.Push(predefinedTypeName)
}

func (n *Listener) ExitPredefinedTypeName(_ *parser.PredefinedTypeNameContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterBlockDefinition(ctx *parser.BlockDefinitionContext) {
	blockDefinition := &BlockDefinition{}

	if ctx.GetTag0() != nil {
		blockDefinition.Tag = tagToTagType(ctx.GetTag0())
	} else {
		blockDefinition.Tag = stringToTagType(ctx.GetTag1())
	}

	n.Push(blockDefinition)
}

func (n *Listener) ExitBlockDefinition(_ *parser.BlockDefinitionContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterEnumTypeName(ctx *parser.EnumTypeNameContext) {
	enumTypeName := &TypeName_EnumTypeName{EnumTypeName: &EnumTypeName{}}

	if ctx.GetIdentifier() != nil {
		enumTypeName.EnumTypeName.Identifier = identifierToIdentType(ctx.GetIdentifier())
	}

	n.Push(enumTypeName)
}

func (n *Listener) ExitEnumTypeName(_ *parser.EnumTypeNameContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterEnumerator(ctx *parser.EnumeratorContext) {
	enumerator := &Enumerator{}

	if ctx.GetTag0() != nil {
		enumerator.Keyword = ctx.GetTag0().GetText()
	} else {
		enumerator.Keyword = ctx.GetTag1().GetText()
	}

	if ctx.GetConstant() != nil {
		enumerator.Constant = numericToLongType(ctx.GetConstant())
	}

	n.Push(enumerator)
}

func (n *Listener) ExitEnumerator(_ *parser.EnumeratorContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterStructTypeName(ctx *parser.StructTypeNameContext) {
	structTypeName := &TypeName_StructTypeName{StructTypeName: &StructTypeName{}}

	if ctx.GetIdentifier() != nil {
		structTypeName.StructTypeName.Identifier = identifierToIdentType(ctx.GetIdentifier())
	}

	n.Push(structTypeName)
}

func (n *Listener) ExitStructTypeName(_ *parser.StructTypeNameContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterStructMember(ctx *parser.StructMemberContext) {
	structMember := &StructMember{}

	if ctx.GetM0() != nil {
		structMember.Star = true
	}

	n.Push(structMember)
}

func (n *Listener) ExitStructMember(_ *parser.StructMemberContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterMember(ctx *parser.MemberContext) {
	member := &Member{}

	if ctx.GetDimension() != nil {
		member.ArraySpecifier = &ArraySpecifier{Constant: make([]*LongType, 0)}
		for _, dimension := range ctx.GetDimension() {
			member.ArraySpecifier.Constant = append(member.ArraySpecifier.Constant, arraySpecifierToLongType(dimension))
		}
	}

	n.Push(member)
}

func (n *Listener) ExitMember(_ *parser.MemberContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterTaggedStructTypeName(ctx *parser.TaggedStructTypeNameContext) {
	taggedStructTypeName := &TypeName_TaggedstructTypeName{TaggedstructTypeName: &TaggedstructTypeName{}}

	if ctx.GetIdentifier() != nil {
		taggedStructTypeName.TaggedstructTypeName.Identifier = identifierToIdentType(ctx.GetIdentifier())
	}

	n.Push(taggedStructTypeName)
}

func (n *Listener) ExitTaggedStructTypeName(_ *parser.TaggedStructTypeNameContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterTaggedStructMember(ctx *parser.TaggedStructMemberContext) {
	taggedstructMember := &TaggedstructMember{}

	if ctx.GetTs0() != nil || ctx.GetBl0() != nil {
		taggedstructMember.Star = true
	}

	if ctx.GetTs0() != nil || ctx.GetTs1() != nil {
		taggedstructMember.Oneof = &TaggedstructMember_TaggedstructDefinition{}
	} else {
		taggedstructMember.Oneof = &TaggedstructMember_BlockDefinition{}
	}

	n.Push(taggedstructMember)
}

func (n *Listener) ExitTaggedStructMember(_ *parser.TaggedStructMemberContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterTaggedStructDefinition(ctx *parser.TaggedStructDefinitionContext) {
	taggedStructDefinition := &TaggedstructDefinition{}

	if ctx.GetTag0() != nil {
		taggedStructDefinition.Tag = tagToTagType(ctx.GetTag0())
	} else if ctx.GetTag1() != nil {
		taggedStructDefinition.Tag = tagToTagType(ctx.GetTag1())
		taggedStructDefinition.Star = true
	} else if ctx.GetTag2() != nil {
		taggedStructDefinition.Tag = stringToTagType(ctx.GetTag2())
	}

	n.Push(taggedStructDefinition)
}

func (n *Listener) ExitTaggedStructDefinition(_ *parser.TaggedStructDefinitionContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterTaggedUnionTypeName(ctx *parser.TaggedUnionTypeNameContext) {
	taggedUnionTypeName := &TypeName_TaggedunionTypeName{
		TaggedunionTypeName: &TaggedunionTypeName{},
	}

	if ctx.GetIdentifier() != nil {
		taggedUnionTypeName.TaggedunionTypeName.Identifier = identifierToIdentType(ctx.GetIdentifier())
	}

	n.Push(taggedUnionTypeName)
}

func (n *Listener) ExitTaggedUnionTypeName(_ *parser.TaggedUnionTypeNameContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterTaggedUnionMember(ctx *parser.TaggedUnionMemberContext) {

	if ctx.GetTag0() != nil {
		tagMember := &TaggedunionMember_TagMember{TagMember: &TagMember{}}
		tagMember.TagMember.Tag = tagToTagType(ctx.GetTag0())
		n.Push(tagMember)
	} else if ctx.GetTag1() != nil {
		tagMember := &TaggedunionMember_TagMember{TagMember: &TagMember{}}
		tagMember.TagMember.Tag = stringToTagType(ctx.GetTag1())
		n.Push(tagMember)
	} else {
		taggedUnionMember := &TaggedunionMember{}
		n.Push(taggedUnionMember)
		// Let the parser call the proper block_definition callback here...
	}
}

func (n *Listener) ExitTaggedUnionMember(_ *parser.TaggedUnionMemberContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}
