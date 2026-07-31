package a2l

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrammar_COMPU_TAB(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		module, ok := parseModule(t,
			"/begin COMPU_TAB compu_tab \"long identifier\" TAB_INTP 2 0 0.5 1 1.5\n/end COMPU_TAB")
		if !ok {
			return
		}

		equalNodes(t, []*CompuTabType{{
			Name:             identVal("compu_tab"),
			LongIdentifier:   strVal("long identifier"),
			ConversionType:   "TAB_INTP",
			NumberValuePairs: intVal("2"),
			InValOutVal: []*CompuTabType_InValOutValType{
				{InVal: floatVal("0"), OutVal: floatVal("0.5")},
				{InVal: floatVal("1"), OutVal: floatVal("1.5")},
			},
		}}, module.COMPU_TAB)
	})

	for _, conversionType := range []string{"TAB_INTP", "TAB_NOINTP"} {
		t.Run("enum/"+conversionType, func(t *testing.T) {
			module, ok := parseModule(t,
				"/begin COMPU_TAB compu_tab \"\" "+conversionType+" 0\n/end COMPU_TAB")
			if !ok || !assert.Len(t, module.COMPU_TAB, 1) {
				return
			}

			assert.Equal(t, conversionType, module.COMPU_TAB[0].ConversionType)
		})
	}

	t.Run("list/no value pair", func(t *testing.T) {
		module, ok := parseModule(t, "/begin COMPU_TAB compu_tab \"\" TAB_INTP 0\n/end COMPU_TAB")
		if !ok || !assert.Len(t, module.COMPU_TAB, 1) {
			return
		}

		assert.Empty(t, module.COMPU_TAB[0].InValOutVal)
	})

	t.Run("optional/DEFAULT_VALUE", func(t *testing.T) {
		module, ok := parseModule(t,
			"/begin COMPU_TAB compu_tab \"\" TAB_INTP 1 0 1 DEFAULT_VALUE \"out of range\"\n/end COMPU_TAB")
		if !ok || !assert.Len(t, module.COMPU_TAB, 1) {
			return
		}

		equalNode(t, &DefaultValueType{DisplayString: strVal("out of range")}, module.COMPU_TAB[0].DEFAULT_VALUE)
	})

	t.Run("reject/unknown conversion type", func(t *testing.T) {
		parseFails(t, moduleScope("/begin COMPU_TAB compu_tab \"\" TAB_VERB 0\n/end COMPU_TAB"))
	})

	t.Run("reject/incomplete value pair", func(t *testing.T) {
		parseFails(t, moduleScope("/begin COMPU_TAB compu_tab \"\" TAB_INTP 1 0\n/end COMPU_TAB"))
	})

	t.Run("reject/DEFAULT_VALUE before the value pairs", func(t *testing.T) {
		parseFails(t, moduleScope(
			"/begin COMPU_TAB compu_tab \"\" TAB_INTP 1 DEFAULT_VALUE \"default\" 0 1\n/end COMPU_TAB"))
	})
}

func TestGrammar_COMPU_VTAB(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		module, ok := parseModule(t,
			"/begin COMPU_VTAB compu_vtab \"long identifier\" TAB_VERB 2 0 \"off\" 1 \"on\"\n/end COMPU_VTAB")
		if !ok {
			return
		}

		equalNodes(t, []*CompuVTabType{{
			Name:             identVal("compu_vtab"),
			LongIdentifier:   strVal("long identifier"),
			ConversionType:   "TAB_VERB",
			NumberValuePairs: intVal("2"),
			InValOutVal: []*CompuVTabType_InValOutValType{
				{InVal: floatVal("0"), OutVal: strVal("off")},
				{InVal: floatVal("1"), OutVal: strVal("on")},
			},
		}}, module.COMPU_VTAB)
	})

	t.Run("optional/DEFAULT_VALUE", func(t *testing.T) {
		module, ok := parseModule(t,
			"/begin COMPU_VTAB compu_vtab \"\" TAB_VERB 1 0 \"off\" DEFAULT_VALUE \"unknown\"\n/end COMPU_VTAB")
		if !ok || !assert.Len(t, module.COMPU_VTAB, 1) {
			return
		}

		equalNode(t, &DefaultValueType{DisplayString: strVal("unknown")}, module.COMPU_VTAB[0].DEFAULT_VALUE)
	})

	t.Run("reject/conversion type other than TAB_VERB", func(t *testing.T) {
		parseFails(t, moduleScope("/begin COMPU_VTAB compu_vtab \"\" TAB_INTP 0\n/end COMPU_VTAB"))
	})

	t.Run("reject/numeric output value", func(t *testing.T) {
		parseFails(t, moduleScope("/begin COMPU_VTAB compu_vtab \"\" TAB_VERB 1 0 1\n/end COMPU_VTAB"))
	})
}

func TestGrammar_DEFAULT_VALUE(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		module, ok := parseModule(t,
			"/begin COMPU_VTAB compu_vtab \"\" TAB_VERB 0 DEFAULT_VALUE \"default\"\n/end COMPU_VTAB")
		if !ok || !assert.Len(t, module.COMPU_VTAB, 1) {
			return
		}

		equalNode(t, &DefaultValueType{DisplayString: strVal("default")}, module.COMPU_VTAB[0].DEFAULT_VALUE)
	})

	t.Run("reject/numeric parameter", func(t *testing.T) {
		parseFails(t, moduleScope("/begin COMPU_VTAB compu_vtab \"\" TAB_VERB 0 DEFAULT_VALUE 1\n/end COMPU_VTAB"))
	})
}

func TestGrammar_DEFAULT_VALUE_NUMERIC(t *testing.T) {
	t.Run("mandatory parameters/in COMPU_TAB", func(t *testing.T) {
		parse(t, moduleScope(
			"/begin COMPU_TAB compu_tab \"\" TAB_INTP 1 0 1 DEFAULT_VALUE_NUMERIC 1.5\n/end COMPU_TAB"))
	})

	// Neither ASAP2 1.51 nor ASAM MCD-2 MC 1.6.1 (chapter 3.5.37) declare DEFAULT_VALUE_NUMERIC
	// for COMPU_VTAB_RANGE, the keyword belongs to COMPU_TAB only.
	t.Run("reject/in COMPU_VTAB_RANGE", func(t *testing.T) {
		parseFails(t, moduleScope(
			"/begin COMPU_VTAB_RANGE compu_vtab_range \"\" 1 0 1 \"on\" DEFAULT_VALUE_NUMERIC 1.5\n"+
				"/end COMPU_VTAB_RANGE"))
	})
}
