package a2l

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func a2mlScope(body string) string {
	return moduleScope("/begin A2ML\n" + body + "\n/end A2ML")
}

func parseA2ML(t assert.TestingT, body string) (*A2MLType, bool) {
	markHelper(t)

	module, ok := parseModule(t, "/begin A2ML\n"+body+"\n/end A2ML")
	if !ok {
		return nil, false
	}

	return module.A2ML, assert.NotNil(t, module.A2ML, "A2ML is missing from the tree")
}

func TestGrammar_A2ML(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		a2ml, ok := parseA2ML(t, "")
		if !ok {
			return
		}

		equalNode(t, &A2MLType{}, a2ml)
	})

	for _, predefinedType := range []string{
		"char", "int", "long", "uchar", "uint", "ulong", "double", "float",
	} {
		t.Run("predefined type/"+predefinedType, func(t *testing.T) {
			a2ml, ok := parseA2ML(t, "struct { "+predefinedType+"; };")
			if !ok {
				return
			}

			equalNode(t, &A2MLType{Declaration: []*Declaration{{
				Oneof: &Declaration_TypeDefinition{TypeDefinition: &TypeDefinition{
					TypeName: &TypeName{Oneof: &TypeName_StructTypeName{StructTypeName: &StructTypeName{
						StructMemberList: []*StructMember{{
							Member: &Member{TypeName: &TypeName{
								Oneof: &TypeName_PredefinedTypeName{
									PredefinedTypeName: &PredefinedTypeName{Name: predefinedType},
								},
							}},
						}},
					}}},
				}},
			}}}, a2ml)
		})
	}

	t.Run("struct/named", func(t *testing.T) {
		a2ml, ok := parseA2ML(t, "struct identifier { uint; };")
		if !ok || !assert.Len(t, a2ml.Declaration, 1) {
			return
		}

		structTypeName := a2ml.Declaration[0].GetTypeDefinition().GetTypeName().GetStructTypeName()
		if assert.NotNil(t, structTypeName) {
			equalNode(t, identVal("identifier"), structTypeName.Identifier)
		}
	})

	t.Run("struct/array member", func(t *testing.T) {
		a2ml, ok := parseA2ML(t, "struct { uchar[8]; };")
		if !ok || !assert.Len(t, a2ml.Declaration, 1) {
			return
		}

		structTypeName := a2ml.Declaration[0].GetTypeDefinition().GetTypeName().GetStructTypeName()
		if !assert.NotNil(t, structTypeName) || !assert.Len(t, structTypeName.StructMemberList, 1) {
			return
		}

		equalNode(t, &ArraySpecifier{Constant: []*LongType{longVal("8")}},
			structTypeName.StructMemberList[0].Member.ArraySpecifier)
	})

	// Chapter 8.1 declares the array specifier as "[" <constant> "]", a constant may be written in
	// hexadecimal notation.
	t.Run("struct/hexadecimal array size", func(t *testing.T) {
		parse(t, a2mlScope("struct { uchar[0x08]; };"))
	})

	t.Run("enum/with values", func(t *testing.T) {
		a2ml, ok := parseA2ML(t, "enum identifier { \"FIRST\" = 0, \"SECOND\" = 1 };")
		if !ok || !assert.Len(t, a2ml.Declaration, 1) {
			return
		}

		enumTypeName := a2ml.Declaration[0].GetTypeDefinition().GetTypeName().GetEnumTypeName()
		if !assert.NotNil(t, enumTypeName) {
			return
		}

		equalNode(t, identVal("identifier"), enumTypeName.Identifier)
		equalNodes(t, []*Enumerator{
			{Keyword: strVal("FIRST"), Constant: longVal("0")},
			{Keyword: strVal("SECOND"), Constant: longVal("1")},
		}, enumTypeName.EnumeratorList)
	})

	t.Run("enum/without values", func(t *testing.T) {
		a2ml, ok := parseA2ML(t, "enum { \"FIRST\", \"SECOND\" };")
		if !ok || !assert.Len(t, a2ml.Declaration, 1) {
			return
		}

		enumTypeName := a2ml.Declaration[0].GetTypeDefinition().GetTypeName().GetEnumTypeName()
		if assert.NotNil(t, enumTypeName) {
			assert.Len(t, enumTypeName.EnumeratorList, 2)
		}
	})

	t.Run("taggedstruct/tagged member", func(t *testing.T) {
		a2ml, ok := parseA2ML(t, "taggedstruct { \"ADDRESS\" ulong; };")
		if !ok || !assert.Len(t, a2ml.Declaration, 1) {
			return
		}

		taggedstruct := a2ml.Declaration[0].GetTypeDefinition().GetTypeName().GetTaggedstructTypeName()
		if !assert.NotNil(t, taggedstruct) || !assert.Len(t, taggedstruct.TaggedstructMemberList, 1) {
			return
		}

		definition := taggedstruct.TaggedstructMemberList[0].GetTaggedstructDefinition()
		if assert.NotNil(t, definition) {
			equalNode(t, &TagType{Value: "ADDRESS"}, definition.Tag)
		}
	})

	t.Run("taggedstruct/repeated member", func(t *testing.T) {
		a2ml, ok := parseA2ML(t, "taggedstruct { (\"EVENT\" uint)*; };")
		if !ok || !assert.Len(t, a2ml.Declaration, 1) {
			return
		}

		taggedstruct := a2ml.Declaration[0].GetTypeDefinition().GetTypeName().GetTaggedstructTypeName()
		if assert.NotNil(t, taggedstruct) {
			assert.Len(t, taggedstruct.TaggedstructMemberList, 1)
		}
	})

	t.Run("taggedunion/tagged members", func(t *testing.T) {
		a2ml, ok := parseA2ML(t, "taggedunion { \"FIRST\" uint; \"SECOND\" ulong; };")
		if !ok || !assert.Len(t, a2ml.Declaration, 1) {
			return
		}

		taggedunion := a2ml.Declaration[0].GetTypeDefinition().GetTypeName().GetTaggedunionTypeName()
		if assert.NotNil(t, taggedunion) {
			assert.Len(t, taggedunion.TaggedunionMemberList, 2)
		}
	})

	t.Run("block definition", func(t *testing.T) {
		a2ml, ok := parseA2ML(t, "block \"IF_DATA\" struct { uint; };")
		if !ok || !assert.Len(t, a2ml.Declaration, 1) {
			return
		}

		blockDefinition := a2ml.Declaration[0].GetBlockDefinition()
		if assert.NotNil(t, blockDefinition) {
			equalNode(t, &TagType{Value: "IF_DATA"}, blockDefinition.Tag)
		}
	})

	t.Run("list/several declarations", func(t *testing.T) {
		a2ml, ok := parseA2ML(t, "struct { uint; };\nenum { \"FIRST\" };\nblock \"IF_DATA\" struct { uint; };")
		if !ok {
			return
		}

		assert.Len(t, a2ml.Declaration, 3)
	})

	// Chapter 6.3.89: "Attention: The interface-specific parameters must be specified directly
	// after the last mandatory parameter 'long identifier'."
	t.Run("reject/A2ML not directly after the module header", func(t *testing.T) {
		parseFails(t, moduleScope(
			"/begin MOD_PAR \"\"\n/end MOD_PAR\n/begin A2ML\nstruct { uint; };\n/end A2ML"))
	})

	t.Run("reject/missing semicolon", func(t *testing.T) {
		parseFails(t, a2mlScope("struct { uint; }"))
	})

	t.Run("reject/unknown predefined type", func(t *testing.T) {
		parseFails(t, a2mlScope("struct { byte; };"))
	})

	t.Run("reject/64 bit predefined types", func(t *testing.T) {
		parseFails(t, a2mlScope("struct { int64; };"))
		parseFails(t, a2mlScope("struct { uint64; };"))
	})

	t.Run("reject/unterminated struct", func(t *testing.T) {
		parseFails(t, a2mlScope("struct { uint; ;"))
	})

	t.Run("reject/block without tag", func(t *testing.T) {
		parseFails(t, a2mlScope("block struct { uint; };"))
	})
}
