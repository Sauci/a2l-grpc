package a2l

// INSTANCE, TYPEDEF_CHARACTERISTIC, TYPEDEF_MEASUREMENT, TYPEDEF_STRUCTURE and
// STRUCTURE_COMPONENT belong to ASAP2 1.7, which is not covered by the supported
// specifications (ASAP2 1.51 and ASAM MCD-2 MC 1.6.1). The grammar does not support them,
// files using these keywords are rejected.

import (
	"testing"
)

func TestGrammar_INSTANCE(t *testing.T) {
	t.Run("reject/not part of the supported specifications", func(t *testing.T) {
		parseFails(t, moduleScope(
			"/begin INSTANCE instance \"long identifier\" type_name 0x1000\n/end INSTANCE"))
	})
}

func TestGrammar_TYPEDEF_CHARACTERISTIC(t *testing.T) {
	t.Run("reject/not part of the supported specifications", func(t *testing.T) {
		parseFails(t, moduleScope("/begin TYPEDEF_CHARACTERISTIC typedef_characteristic "+
			"\"long identifier\" VALUE record_layout 0 compu_method 0 100\n/end TYPEDEF_CHARACTERISTIC"))
	})
}

func TestGrammar_TYPEDEF_MEASUREMENT(t *testing.T) {
	t.Run("reject/not part of the supported specifications", func(t *testing.T) {
		parseFails(t, moduleScope("/begin TYPEDEF_MEASUREMENT typedef_measurement "+
			"\"long identifier\" UBYTE compu_method 8 0.5 0 255\n/end TYPEDEF_MEASUREMENT"))
	})
}

func TestGrammar_TYPEDEF_STRUCTURE(t *testing.T) {
	t.Run("reject/not part of the supported specifications", func(t *testing.T) {
		parseFails(t, moduleScope("/begin TYPEDEF_STRUCTURE typedef_structure \"long identifier\" "+
			"4 SYMBOL_TYPE_LINK \"symbol\"\n/end TYPEDEF_STRUCTURE"))
	})
}

func TestGrammar_STRUCTURE_COMPONENT(t *testing.T) {
	t.Run("reject/not part of the supported specifications", func(t *testing.T) {
		parseFails(t, moduleScope("/begin TYPEDEF_STRUCTURE typedef_structure \"\" 4\n" +
			"/begin STRUCTURE_COMPONENT component type_name 0\n" +
			"/end STRUCTURE_COMPONENT\n/end TYPEDEF_STRUCTURE"))
	})
}
