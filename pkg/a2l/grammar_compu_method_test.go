package a2l

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrammar_COMPU_METHOD(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		module, ok := parseModule(t,
			"/begin COMPU_METHOD compu_method \"long identifier\" RAT_FUNC \"%4.2\" \"km/h\"\n/end COMPU_METHOD")
		if !ok {
			return
		}

		equalNodes(t, []*CompuMethodType{{
			Name:           identVal("compu_method"),
			LongIdentifier: strVal("long identifier"),
			ConversionType: "RAT_FUNC",
			Format:         strVal("%4.2"),
			Unit:           strVal("km/h"),
		}}, module.COMPU_METHOD)
	})

	for _, conversionType := range []string{
		"TAB_INTP", "TAB_NOINTP", "TAB_VERB", "RAT_FUNC", "FORM", "IDENTICAL", "LINEAR",
	} {
		t.Run("enum/"+conversionType, func(t *testing.T) {
			module, ok := parseModule(t, "/begin COMPU_METHOD compu_method \"\" "+conversionType+
				" \"%4.2\" \"unit\"\n/end COMPU_METHOD")
			if !ok || !assert.Len(t, module.COMPU_METHOD, 1) {
				return
			}

			assert.Equal(t, conversionType, module.COMPU_METHOD[0].ConversionType)
		})
	}

	t.Run("optional/COEFFS and COMPU_TAB_REF and REF_UNIT", func(t *testing.T) {
		compuMethod, ok := parseCompuMethod(t, `COEFFS 0 1 2 3 4 5
COMPU_TAB_REF compu_tab
REF_UNIT unit`)
		if !ok {
			return
		}

		equalNode(t, &CompuMethodType{
			Name:           identVal("compu_method"),
			LongIdentifier: strVal(""),
			ConversionType: "RAT_FUNC",
			Format:         strVal("%4.2"),
			Unit:           strVal("unit"),
			COEFFS: &CoeffsType{
				A: floatVal("0"),
				B: floatVal("1"),
				C: floatVal("2"),
				D: floatVal("3"),
				E: floatVal("4"),
				F: floatVal("5"),
			},
			COMPU_TAB_REF: &CompuTabRefType{ConversionTable: identVal("compu_tab")},
			REF_UNIT:      &RefUnitType{Unit: identVal("unit")},
		}, compuMethod)
	})

	t.Run("optional/FORMULA", func(t *testing.T) {
		compuMethod, ok := parseCompuMethod(t, "/begin FORMULA \"X1*2\"\n/end FORMULA")
		if !ok {
			return
		}

		equalNode(t, &FormulaType{FX: strVal("X1*2")}, compuMethod.FORMULA)
	})

	t.Run("reject/missing unit", func(t *testing.T) {
		parseFails(t, moduleScope("/begin COMPU_METHOD compu_method \"\" RAT_FUNC \"%4.2\"\n/end COMPU_METHOD"))
	})

	t.Run("reject/unknown conversion type", func(t *testing.T) {
		parseFails(t, moduleScope(
			"/begin COMPU_METHOD compu_method \"\" TAB_LIN \"%4.2\" \"unit\"\n/end COMPU_METHOD"))
	})

	t.Run("reject/identifier as format", func(t *testing.T) {
		parseFails(t, moduleScope(
			"/begin COMPU_METHOD compu_method \"\" RAT_FUNC format \"unit\"\n/end COMPU_METHOD"))
	})
}

func TestGrammar_COEFFS(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		compuMethod, ok := parseCompuMethod(t, "COEFFS 0 1.5 -2 3 4 5")
		if !ok {
			return
		}

		equalNode(t, &CoeffsType{
			A: floatVal("0"),
			B: floatVal("1.5"),
			C: floatVal("-2"),
			D: floatVal("3"),
			E: floatVal("4"),
			F: floatVal("5"),
		}, compuMethod.COEFFS)
	})

	t.Run("reject/five coefficients", func(t *testing.T) {
		parseFails(t, compuMethodScope("COEFFS 0 1 2 3 4"))
	})

	t.Run("reject/seven coefficients", func(t *testing.T) {
		parseFails(t, compuMethodScope("COEFFS 0 1 2 3 4 5 6"))
	})
}

func TestGrammar_COEFFS_LINEAR(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		assertPreserved(t, compuMethodScope("COEFFS_LINEAR 1.5 0"), "COEFFS_LINEAR")
	})

	t.Run("reject/single coefficient", func(t *testing.T) {
		parseFails(t, compuMethodScope("COEFFS_LINEAR 1.5"))
	})
}

func TestGrammar_COMPU_TAB_REF(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		compuMethod, ok := parseCompuMethod(t, "COMPU_TAB_REF compu_tab")
		if !ok {
			return
		}

		equalNode(t, &CompuTabRefType{ConversionTable: identVal("compu_tab")}, compuMethod.COMPU_TAB_REF)
	})

	t.Run("reject/string parameter", func(t *testing.T) {
		parseFails(t, compuMethodScope("COMPU_TAB_REF \"compu_tab\""))
	})
}

func TestGrammar_REF_UNIT(t *testing.T) {
	t.Run("mandatory parameters/in COMPU_METHOD", func(t *testing.T) {
		compuMethod, ok := parseCompuMethod(t, "REF_UNIT unit")
		if !ok {
			return
		}

		equalNode(t, &RefUnitType{Unit: identVal("unit")}, compuMethod.REF_UNIT)
	})

	t.Run("mandatory parameters/in UNIT", func(t *testing.T) {
		unit, ok := parseUnit(t, "REF_UNIT other_unit")
		if !ok {
			return
		}

		equalNode(t, &RefUnitType{Unit: identVal("other_unit")}, unit.REF_UNIT)
	})

	t.Run("reject/missing unit", func(t *testing.T) {
		parseFails(t, compuMethodScope("REF_UNIT"))
	})
}

func TestGrammar_FORMULA(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		compuMethod, ok := parseCompuMethod(t, "/begin FORMULA \"sqrt(X1)\"\n/end FORMULA")
		if !ok {
			return
		}

		equalNode(t, &FormulaType{FX: strVal("sqrt(X1)")}, compuMethod.FORMULA)
	})

	t.Run("optional/FORMULA_INV", func(t *testing.T) {
		compuMethod, ok := parseCompuMethod(t,
			"/begin FORMULA \"sqrt(X1)\"\nFORMULA_INV \"X1*X1\"\n/end FORMULA")
		if !ok {
			return
		}

		equalNode(t, &FormulaType{
			FX:          strVal("sqrt(X1)"),
			FORMULA_INV: &FormulaInvType{GX: strVal("X1*X1")},
		}, compuMethod.FORMULA)
	})

	t.Run("reject/missing formula", func(t *testing.T) {
		parseFails(t, compuMethodScope("/begin FORMULA\n/end FORMULA"))
	})

	t.Run("reject/identifier as formula", func(t *testing.T) {
		parseFails(t, compuMethodScope("/begin FORMULA X1\n/end FORMULA"))
	})
}

func TestGrammar_FORMULA_INV(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		compuMethod, ok := parseCompuMethod(t, "/begin FORMULA \"X1\"\nFORMULA_INV \"X1\"\n/end FORMULA")
		if !ok || !assert.NotNil(t, compuMethod.FORMULA) {
			return
		}

		equalNode(t, &FormulaInvType{GX: strVal("X1")}, compuMethod.FORMULA.FORMULA_INV)
	})

	t.Run("reject/outside of FORMULA", func(t *testing.T) {
		parseFails(t, compuMethodScope("FORMULA_INV \"X1\""))
	})
}

func TestGrammar_STATUS_STRING_REF(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		assertPreserved(t, compuMethodScope("STATUS_STRING_REF compu_vtab"), "STATUS_STRING_REF")
	})

	t.Run("reject/missing conversion table", func(t *testing.T) {
		parseFails(t, compuMethodScope("STATUS_STRING_REF"))
	})
}

func TestGrammar_COMPU_VTAB_RANGE(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		module, ok := parseModule(t, "/begin COMPU_VTAB_RANGE compu_vtab_range \"long identifier\" 2 "+
			"0 9 \"low\" 10 19 \"high\"\n/end COMPU_VTAB_RANGE")
		if !ok {
			return
		}

		equalNodes(t, []*CompuVTabRangeType{{
			Name:                  identVal("compu_vtab_range"),
			LongIdentifier:        strVal("long identifier"),
			NumberOfValuesTriples: intVal("2"),
			InValMinInValMaxOutVal: []*CompuVTabRangeType_InValMinInValMaxOutValType{
				{InValMin: floatVal("0"), InValMax: floatVal("9"), OutVal: strVal("low")},
				{InValMin: floatVal("10"), InValMax: floatVal("19"), OutVal: strVal("high")},
			},
		}}, module.COMPU_VTAB_RANGE)
	})

	t.Run("list/no value triple", func(t *testing.T) {
		module, ok := parseModule(t,
			"/begin COMPU_VTAB_RANGE compu_vtab_range \"\" 0\n/end COMPU_VTAB_RANGE")
		if !ok || !assert.Len(t, module.COMPU_VTAB_RANGE, 1) {
			return
		}

		assert.Empty(t, module.COMPU_VTAB_RANGE[0].InValMinInValMaxOutVal)
	})

	t.Run("optional/DEFAULT_VALUE", func(t *testing.T) {
		module, ok := parseModule(t, "/begin COMPU_VTAB_RANGE compu_vtab_range \"\" 1 0 9 \"low\" "+
			"DEFAULT_VALUE \"out of range\"\n/end COMPU_VTAB_RANGE")
		if !ok || !assert.Len(t, module.COMPU_VTAB_RANGE, 1) {
			return
		}

		equalNode(t, &DefaultValueType{DisplayString: strVal("out of range")},
			module.COMPU_VTAB_RANGE[0].DEFAULT_VALUE)
	})

	t.Run("reject/incomplete value triple", func(t *testing.T) {
		parseFails(t, moduleScope(
			"/begin COMPU_VTAB_RANGE compu_vtab_range \"\" 1 0 9\n/end COMPU_VTAB_RANGE"))
	})

	// COMPU_VTAB_RANGE has no conversion type parameter.
	t.Run("reject/conversion type", func(t *testing.T) {
		parseFails(t, moduleScope(
			"/begin COMPU_VTAB_RANGE compu_vtab_range \"\" TAB_VERB 0\n/end COMPU_VTAB_RANGE"))
	})
}
