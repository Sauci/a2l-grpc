package a2l

import (
	"github.com/ngamber/a2l-grpc/pkg/a2l/parser"
)

func (n *Listener) EnterIfData(ctx *parser.IfDataContext) {
	ifData := &IfDataType{Name: identifierToIdentType(ctx.GetName())}

	n.Push(ifData)
}

func (n *Listener) ExitIfData(_ *parser.IfDataContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterGenericParameter(ctx *parser.GenericParameterContext) {
	genericParameter := &GenericParameterType{}

	if ctx.GetTag() != nil {
		genericParameter.Oneof = &GenericParameterType_Tag{Tag: tagToTagType(ctx.GetTag())}
	} else if ctx.GetSting() != nil {
		genericParameter.Oneof = &GenericParameterType_String_{String_: a2lStringToStringType(ctx.GetSting())}
	} else if ctx.GetNumeric() != nil {
		if ctx.GetNumeric().GetH() != nil {
			genericParameter.Oneof = &GenericParameterType_Long{Long: numericToLongType(ctx.GetNumeric())}
		} else if ctx.GetNumeric().GetI() != nil {
			genericParameter.Oneof = &GenericParameterType_Long{Long: numericToLongType(ctx.GetNumeric())}
		} else {
			genericParameter.Oneof = &GenericParameterType_Float{Float: floatToFloatType(ctx.GetNumeric())}
		}
	} else if ctx.GetGeneric() != nil {
		// let the recursion happening here...
	} else if ctx.GetIdentifier() != nil {
		genericParameter.Oneof = &GenericParameterType_Identifier{Identifier: identifierToIdentType(ctx.GetIdentifier())}
	} else {
		panic("runtime...")
	}

	n.Push(genericParameter)
}

func (n *Listener) ExitGenericParameter(_ *parser.GenericParameterContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterGenericNode(ctx *parser.GenericNodeContext) {
	genericNode := &GenericNodeType{Name: identifierToIdentType(ctx.GetName())}

	n.Push(genericNode)
}

func (n *Listener) ExitGenericNode(_ *parser.GenericNodeContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}
