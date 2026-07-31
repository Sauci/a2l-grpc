package a2l

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrammar_VARIANT_CODING(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		variantCoding, ok := parseVariantCoding(t, "")
		if !ok {
			return
		}

		equalNode(t, &VariantCodingType{}, variantCoding)
	})

	t.Run("optional/all parameters", func(t *testing.T) {
		variantCoding, ok := parseVariantCoding(t, `VAR_SEPARATOR "."
VAR_NAMING NUMERIC
/begin VAR_CRITERION criterion "long identifier" first second
/end VAR_CRITERION
/begin VAR_FORBIDDEN_COMB criterion first
/end VAR_FORBIDDEN_COMB
/begin VAR_CHARACTERISTIC characteristic criterion
/end VAR_CHARACTERISTIC`)
		if !ok {
			return
		}

		equalNode(t, &VariantCodingType{
			VAR_SEPARATOR: &VarSeparatorType{Separator: strVal(".")},
			VAR_NAMING:    &VarNamingType{Tag: "NUMERIC"},
			VAR_CRITERION: []*VarCriterionType{{
				Name:           identVal("criterion"),
				LongIdentifier: strVal("long identifier"),
				Value:          []*IdentType{identVal("first"), identVal("second")},
			}},
			VAR_FORBIDDEN_COMB: []*VarForbiddenCombType{{
				CriterionNameCriterionValue: []*VarForbiddenCombType_CriterionType{
					{CriterionName: identVal("criterion"), CriterionValue: identVal("first")},
				},
			}},
			VAR_CHARACTERISTIC: []*VarCharacteristicType{{
				Name:          identVal("characteristic"),
				CriterionName: []*IdentType{identVal("criterion")},
			}},
		}, variantCoding)
	})
}

func TestGrammar_VAR_NAMING(t *testing.T) {
	t.Run("enum/NUMERIC", func(t *testing.T) {
		variantCoding, ok := parseVariantCoding(t, "VAR_NAMING NUMERIC")
		if !ok {
			return
		}

		equalNode(t, &VarNamingType{Tag: "NUMERIC"}, variantCoding.VAR_NAMING)
	})

	// NUMERIC is the only value defined by ASAP2 1.51 and ASAM MCD-2 MC 1.6.1 (chapter 3.5.134);
	// ALPHA is named by the specification but reserved for a future extension, so it is accepted
	// by the grammar and reported as a warning. The misspelled value APLHA exists in no version of
	// the standard.
	t.Run("enum/ALPHA", func(t *testing.T) {
		variantCoding, ok := parseVariantCoding(t, "VAR_NAMING ALPHA")
		if !ok {
			return
		}

		equalNode(t, &VarNamingType{Tag: "ALPHA"}, variantCoding.VAR_NAMING)
	})

	t.Run("enum/ALPHA is warned about", func(t *testing.T) {
		_, warnings, err := GetTreeFromStringWithOptions(variantCodingScope("VAR_NAMING ALPHA"),
			ParseOptions{})
		assert.NoError(t, err)

		if assert.NotEmpty(t, warnings) {
			assert.Contains(t, warnings[0].String(),
				"VAR_NAMING ALPHA is reserved for a future extension of the standard")
		}
	})

	t.Run("reject/misspelled ALPHA", func(t *testing.T) {
		parseFails(t, variantCodingScope("VAR_NAMING APLHA"))
	})

	t.Run("reject/unknown tag", func(t *testing.T) {
		parseFails(t, variantCodingScope("VAR_NAMING HEXADECIMAL"))
	})
}

func TestGrammar_VAR_SEPARATOR(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		variantCoding, ok := parseVariantCoding(t, "VAR_SEPARATOR \".\"")
		if !ok {
			return
		}

		equalNode(t, &VarSeparatorType{Separator: strVal(".")}, variantCoding.VAR_SEPARATOR)
	})

	t.Run("reject/identifier parameter", func(t *testing.T) {
		parseFails(t, variantCodingScope("VAR_SEPARATOR separator"))
	})
}

func TestGrammar_VAR_CRITERION(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		variantCoding, ok := parseVariantCoding(t,
			"/begin VAR_CRITERION criterion \"long identifier\"\n/end VAR_CRITERION")
		if !ok {
			return
		}

		equalNodes(t, []*VarCriterionType{{
			Name:           identVal("criterion"),
			LongIdentifier: strVal("long identifier"),
		}}, variantCoding.VAR_CRITERION)
	})

	t.Run("list/several values", func(t *testing.T) {
		variantCoding, ok := parseVariantCoding(t,
			"/begin VAR_CRITERION criterion \"\" first second third\n/end VAR_CRITERION")
		if !ok {
			return
		}

		equalNodes(t, []*VarCriterionType{{
			Name:           identVal("criterion"),
			LongIdentifier: strVal(""),
			Value:          []*IdentType{identVal("first"), identVal("second"), identVal("third")},
		}}, variantCoding.VAR_CRITERION)
	})

	t.Run("optional/all parameters", func(t *testing.T) {
		varCriterion, ok := parseVarCriterion(t,
			"VAR_MEASUREMENT measurement\nVAR_SELECTION_CHARACTERISTIC characteristic")
		if !ok {
			return
		}

		equalNode(t, &VarCriterionType{
			Name:                         identVal("criterion"),
			LongIdentifier:               strVal(""),
			Value:                        []*IdentType{identVal("first"), identVal("second")},
			VAR_MEASUREMENT:              &VarMeasurementType{Name: identVal("measurement")},
			VAR_SELECTION_CHARACTERISTIC: &VarSelectionCharacteristicType{Name: identVal("characteristic")},
		}, varCriterion)
	})

	t.Run("list/several VAR_CRITERION", func(t *testing.T) {
		variantCoding, ok := parseVariantCoding(t,
			"/begin VAR_CRITERION first \"\"\n/end VAR_CRITERION\n"+
				"/begin VAR_CRITERION second \"\"\n/end VAR_CRITERION")
		if !ok {
			return
		}

		assert.Len(t, variantCoding.VAR_CRITERION, 2)
	})

	t.Run("reject/missing long identifier", func(t *testing.T) {
		parseFails(t, variantCodingScope("/begin VAR_CRITERION criterion\n/end VAR_CRITERION"))
	})

	t.Run("reject/string as criterion value", func(t *testing.T) {
		parseFails(t, variantCodingScope("/begin VAR_CRITERION criterion \"\" \"first\"\n/end VAR_CRITERION"))
	})
}

func TestGrammar_VAR_MEASUREMENT(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		varCriterion, ok := parseVarCriterion(t, "VAR_MEASUREMENT measurement")
		if !ok {
			return
		}

		equalNode(t, &VarMeasurementType{Name: identVal("measurement")}, varCriterion.VAR_MEASUREMENT)
	})

	t.Run("reject/string parameter", func(t *testing.T) {
		parseFails(t, varCriterionScope("VAR_MEASUREMENT \"measurement\""))
	})

	t.Run("reject/missing name", func(t *testing.T) {
		parseFails(t, varCriterionScope("VAR_MEASUREMENT"))
	})
}

func TestGrammar_VAR_SELECTION_CHARACTERISTIC(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		varCriterion, ok := parseVarCriterion(t, "VAR_SELECTION_CHARACTERISTIC characteristic")
		if !ok {
			return
		}

		equalNode(t, &VarSelectionCharacteristicType{Name: identVal("characteristic")},
			varCriterion.VAR_SELECTION_CHARACTERISTIC)
	})

	t.Run("reject/missing name", func(t *testing.T) {
		parseFails(t, varCriterionScope("VAR_SELECTION_CHARACTERISTIC"))
	})
}

func TestGrammar_VAR_FORBIDDEN_COMB(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		variantCoding, ok := parseVariantCoding(t,
			"/begin VAR_FORBIDDEN_COMB\n/end VAR_FORBIDDEN_COMB")
		if !ok {
			return
		}

		equalNodes(t, []*VarForbiddenCombType{{}}, variantCoding.VAR_FORBIDDEN_COMB)
	})

	t.Run("list/several criterion pairs", func(t *testing.T) {
		variantCoding, ok := parseVariantCoding(t,
			"/begin VAR_FORBIDDEN_COMB first one second two\n/end VAR_FORBIDDEN_COMB")
		if !ok {
			return
		}

		equalNodes(t, []*VarForbiddenCombType{{
			CriterionNameCriterionValue: []*VarForbiddenCombType_CriterionType{
				{CriterionName: identVal("first"), CriterionValue: identVal("one")},
				{CriterionName: identVal("second"), CriterionValue: identVal("two")},
			},
		}}, variantCoding.VAR_FORBIDDEN_COMB)
	})

	t.Run("reject/incomplete criterion pair", func(t *testing.T) {
		parseFails(t, variantCodingScope("/begin VAR_FORBIDDEN_COMB first one second\n/end VAR_FORBIDDEN_COMB"))
	})
}

func TestGrammar_VAR_ADDRESS(t *testing.T) {
	type testCaseType struct {
		name      string
		content   string
		addresses []*LongType
	}

	for _, testCase := range []testCaseType{
		{name: "list/no address", content: "", addresses: nil},
		{name: "list/single address", content: "0x1000", addresses: []*LongType{longVal("0x1000")}},
		{
			name:      "list/several addresses",
			content:   "0x1000 0x2000 0x3000",
			addresses: []*LongType{longVal("0x1000"), longVal("0x2000"), longVal("0x3000")},
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			variantCoding, ok := parseVariantCoding(t,
				"/begin VAR_CHARACTERISTIC characteristic criterion\n"+
					"/begin VAR_ADDRESS "+testCase.content+"\n/end VAR_ADDRESS\n/end VAR_CHARACTERISTIC")
			if !ok || !assert.Len(t, variantCoding.VAR_CHARACTERISTIC, 1) {
				return
			}

			equalNode(t, &VarAddressType{Address: testCase.addresses},
				variantCoding.VAR_CHARACTERISTIC[0].VAR_ADDRESS)
		})
	}

	t.Run("reject/string address", func(t *testing.T) {
		parseFails(t, variantCodingScope("/begin VAR_CHARACTERISTIC characteristic criterion\n"+
			"/begin VAR_ADDRESS \"0x1000\"\n/end VAR_ADDRESS\n/end VAR_CHARACTERISTIC"))
	})
}

func TestGrammar_VAR_CHARACTERISTIC(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		variantCoding, ok := parseVariantCoding(t,
			"/begin VAR_CHARACTERISTIC characteristic\n/end VAR_CHARACTERISTIC")
		if !ok {
			return
		}

		equalNodes(t, []*VarCharacteristicType{{Name: identVal("characteristic")}},
			variantCoding.VAR_CHARACTERISTIC)
	})

	t.Run("list/several criterion names", func(t *testing.T) {
		variantCoding, ok := parseVariantCoding(t,
			"/begin VAR_CHARACTERISTIC characteristic first second third\n/end VAR_CHARACTERISTIC")
		if !ok {
			return
		}

		equalNodes(t, []*VarCharacteristicType{{
			Name:          identVal("characteristic"),
			CriterionName: []*IdentType{identVal("first"), identVal("second"), identVal("third")},
		}}, variantCoding.VAR_CHARACTERISTIC)
	})

	t.Run("optional/VAR_ADDRESS", func(t *testing.T) {
		variantCoding, ok := parseVariantCoding(t,
			"/begin VAR_CHARACTERISTIC characteristic criterion\n"+
				"/begin VAR_ADDRESS 0x1000 0x2000\n/end VAR_ADDRESS\n/end VAR_CHARACTERISTIC")
		if !ok || !assert.Len(t, variantCoding.VAR_CHARACTERISTIC, 1) {
			return
		}

		equalNode(t, &VarAddressType{Address: []*LongType{longVal("0x1000"), longVal("0x2000")}},
			variantCoding.VAR_CHARACTERISTIC[0].VAR_ADDRESS)
	})

	t.Run("reject/missing name", func(t *testing.T) {
		parseFails(t, variantCodingScope("/begin VAR_CHARACTERISTIC\n/end VAR_CHARACTERISTIC"))
	})
}
