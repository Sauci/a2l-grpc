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

	// The dimension of a member of a named type must stay a rule of the parser: the A2L lexer
	// consumes the point of an identifier, and swallowing the brackets as well would turn
	// "my_type[4]" into a single identifier and drop the dimension.
	t.Run("struct/array of a named type", func(t *testing.T) {
		a2ml, ok := parseA2ML(t, "struct { struct my_type[4]; };")
		if !ok || !assert.Len(t, a2ml.Declaration, 1) {
			return
		}

		structTypeName := a2ml.Declaration[0].GetTypeDefinition().GetTypeName().GetStructTypeName()
		if !assert.NotNil(t, structTypeName) || !assert.Len(t, structTypeName.StructMemberList, 1) {
			return
		}

		equalNode(t, &ArraySpecifier{Constant: []*LongType{longVal("4")}},
			structTypeName.StructMemberList[0].Member.ArraySpecifier)
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

	// Chapter 5.2 declares two forms of taggedstruct_definition: "tag [ member ]" and
	// "tag "(" member ")*;"". The second one is the sequence with a base type, and it carries its
	// tag outside of the parentheses.
	t.Run("taggedstruct/repeated tagged member", func(t *testing.T) {
		a2ml, ok := parseA2ML(t, "taggedstruct { \"EVENT\" (uint)*; };")
		if !ok || !assert.Len(t, a2ml.Declaration, 1) {
			return
		}

		taggedstruct := a2ml.Declaration[0].GetTypeDefinition().GetTypeName().GetTaggedstructTypeName()
		if !assert.NotNil(t, taggedstruct) || !assert.Len(t, taggedstruct.TaggedstructMemberList, 1) {
			return
		}

		definition := taggedstruct.TaggedstructMemberList[0].GetTaggedstructDefinition()
		if assert.NotNil(t, definition) {
			equalNode(t, &TagType{Value: "EVENT"}, definition.Tag)
			assert.True(t, definition.Star, "the repeated form should be marked with a star")
		}
	})

	// Chapter 5.2: "All elements are optional and each element is identified by its tag", and both
	// forms of taggedstruct_definition start with the tag.
	t.Run("reject/taggedstruct member without a tag", func(t *testing.T) {
		parseFails(t, a2mlScope("taggedstruct { uint; };"))
	})

	// Chapter 5.2 declares the array specifier as "[" <constant> "]". The homonymous rule of the
	// A2L grammar also accepts an alphabetic string, because chapter 3.2 allows one as the index
	// of a partial identifier, which is a different thing.
	t.Run("reject/symbolic array size", func(t *testing.T) {
		parseFailsWithSyntaxError(t, a2mlScope("struct { uchar[MAX_LEN]; };"))
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

// ASAM MCD-2 MC 1.6.1 chapter 5.2: "Within the AML own name spaces are used. In this case it is
// allowed to reuse ASAM MCD-2 MC keyword names. The definitions from the AML are exclusively valid
// in IF_DATA. Outside IF_DATA only the keywords according to chapter 3.5 are valid."
func TestGrammar_A2ML_ReusesA2LKeywordNames(t *testing.T) {
	for _, keyword := range []string{"IF_DATA", "MEASUREMENT", "UNIT", "VERSION", "RESERVED"} {
		t.Run("taggedunion named "+keyword, func(t *testing.T) {
			a2ml, ok := parseA2ML(t, "taggedunion "+keyword+" { \"FIRST\" uint; };")
			if !ok || !assert.Len(t, a2ml.Declaration, 1) {
				return
			}

			taggedunion := a2ml.Declaration[0].GetTypeDefinition().GetTypeName().GetTaggedunionTypeName()
			if assert.NotNil(t, taggedunion) {
				equalNode(t, identVal(keyword), taggedunion.Identifier)
			}
		})

		t.Run("struct named "+keyword, func(t *testing.T) {
			parse(t, a2mlScope("struct "+keyword+" { uint; };"))
		})
	}

	// The keywords of the metalanguage itself stay reserved inside the AML.
	t.Run("reject/a metalanguage keyword as an identifier", func(t *testing.T) {
		parseFails(t, a2mlScope("struct struct { uint; };"))
	})

	// Appendix B.1, SUPP1_IF.AML: the AML the specification itself ships.
	t.Run("the reference AML of the specification", func(t *testing.T) {
		parse(t, a2mlScope(`enum mem_typ { "INTERN" = 0, "EXTERN" = 1 };
enum addr_typ { "BYTE" = 1, "WORD" = 2, "LONG" = 4 };
enum addr_mode { "DIRECT" = 0, "INDIRECT" = 1 };

taggedunion IF_DATA {
  "DIM" taggedstruct {
    (block "SOURCE" struct {
      struct {
        char [101];
        int;
        long;
      };
      taggedstruct {
        block "QP_BLOB" struct {
          ulong;
          int;
          ulong;
          long;
        };
      };
    }
    )*;

    block "TP_BLOB" struct {
      int;
    };

    block "KP_BLOB" struct {
      ulong;
      enum addr_typ;
    };

    block "DP_BLOB" struct {
      enum mem_typ;
    };
    block "PA_BLOB" struct {
      enum addr_mode;
    };
  };
};`))
	})
}
