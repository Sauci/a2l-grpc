package a2l

// INSTANCE, TYPEDEF_CHARACTERISTIC, TYPEDEF_MEASUREMENT, TYPEDEF_STRUCTURE and
// STRUCTURE_COMPONENT.
//
// None of these keywords belongs to ASAP2 1.51, they were introduced by later versions of the
// standard. The grammar accepts them, but the protobuf definition has no message for them, so
// they never reach the tree.

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrammar_INSTANCE(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		deviation(t, "INSTANCE has no node in the tree", func(t assert.TestingT) {
			assertPreserved(t, moduleScope(
				"/begin INSTANCE instance \"long identifier\" type_name 0x1000\n/end INSTANCE"), "INSTANCE")
		})
	})

	t.Run("optional/IF_DATA and ECU_ADDRESS_EXTENSION", func(t *testing.T) {
		deviation(t, "INSTANCE has no node in the tree", func(t assert.TestingT) {
			assertPreserved(t, moduleScope("/begin INSTANCE instance \"\" type_name 0x1000\n"+
				"ECU_ADDRESS_EXTENSION 1\n/begin IF_DATA XCP\n/end IF_DATA\n/end INSTANCE"), "INSTANCE")
		})
	})

	t.Run("reject/missing address", func(t *testing.T) {
		parseFails(t, moduleScope("/begin INSTANCE instance \"\" type_name\n/end INSTANCE"))
	})

	t.Run("reject/missing type name", func(t *testing.T) {
		parseFails(t, moduleScope("/begin INSTANCE instance \"\" 0x1000\n/end INSTANCE"))
	})
}

func TestGrammar_TYPEDEF_CHARACTERISTIC(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		deviation(t, "TYPEDEF_CHARACTERISTIC has no node in the tree", func(t assert.TestingT) {
			assertPreserved(t, moduleScope("/begin TYPEDEF_CHARACTERISTIC typedef_characteristic "+
				"\"long identifier\" VALUE record_layout 0 compu_method 0 100\n/end TYPEDEF_CHARACTERISTIC"),
				"TYPEDEF_CHARACTERISTIC")
		})
	})

	for _, characteristicType := range []string{
		"ASCII", "CURVE", "MAP", "CUBOID", "CUBE_4", "CUBE_5", "VAL_BLK", "VALUE",
	} {
		t.Run("enum/"+characteristicType, func(t *testing.T) {
			parse(t, moduleScope("/begin TYPEDEF_CHARACTERISTIC typedef_characteristic \"\" "+
				characteristicType+" record_layout 0 compu_method 0 100\n/end TYPEDEF_CHARACTERISTIC"))
		})
	}

	// A TYPEDEF_CHARACTERISTIC describes a type, it has no address.
	t.Run("reject/address", func(t *testing.T) {
		parseFails(t, moduleScope("/begin TYPEDEF_CHARACTERISTIC typedef_characteristic \"\" VALUE 0x0 "+
			"record_layout 0 compu_method 0 100\n/end TYPEDEF_CHARACTERISTIC"))
	})

	t.Run("reject/missing upper limit", func(t *testing.T) {
		parseFails(t, moduleScope("/begin TYPEDEF_CHARACTERISTIC typedef_characteristic \"\" VALUE "+
			"record_layout 0 compu_method 0\n/end TYPEDEF_CHARACTERISTIC"))
	})
}

func TestGrammar_TYPEDEF_MEASUREMENT(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		deviation(t, "TYPEDEF_MEASUREMENT has no node in the tree", func(t assert.TestingT) {
			assertPreserved(t, moduleScope("/begin TYPEDEF_MEASUREMENT typedef_measurement "+
				"\"long identifier\" UBYTE compu_method 8 0.5 0 255\n/end TYPEDEF_MEASUREMENT"),
				"TYPEDEF_MEASUREMENT")
		})
	})

	t.Run("reject/unknown data type", func(t *testing.T) {
		parseFails(t, moduleScope("/begin TYPEDEF_MEASUREMENT typedef_measurement \"\" UINT24 "+
			"compu_method 8 0.5 0 255\n/end TYPEDEF_MEASUREMENT"))
	})

	t.Run("reject/missing upper limit", func(t *testing.T) {
		parseFails(t, moduleScope("/begin TYPEDEF_MEASUREMENT typedef_measurement \"\" UBYTE "+
			"compu_method 8 0.5 0\n/end TYPEDEF_MEASUREMENT"))
	})
}

func TestGrammar_TYPEDEF_STRUCTURE(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		deviation(t, "TYPEDEF_STRUCTURE has no node in the tree", func(t assert.TestingT) {
			assertPreserved(t, moduleScope("/begin TYPEDEF_STRUCTURE typedef_structure \"long identifier\" "+
				"4 SYMBOL_TYPE_LINK \"symbol\"\n/end TYPEDEF_STRUCTURE"), "TYPEDEF_STRUCTURE")
		})
	})

	t.Run("optional/STRUCTURE_COMPONENT", func(t *testing.T) {
		deviation(t, "TYPEDEF_STRUCTURE has no node in the tree", func(t assert.TestingT) {
			assertPreserved(t, moduleScope("/begin TYPEDEF_STRUCTURE typedef_structure \"\" 4 "+
				"SYMBOL_TYPE_LINK \"symbol\"\n"+
				"/begin STRUCTURE_COMPONENT component type_name 0 SYMBOL_TYPE_LINK \"symbol\"\n"+
				"/end STRUCTURE_COMPONENT\n/end TYPEDEF_STRUCTURE"), "TYPEDEF_STRUCTURE")
		})
	})

	// SYMBOL_TYPE_LINK is an optional parameter of TYPEDEF_STRUCTURE.
	t.Run("accept/without SYMBOL_TYPE_LINK", func(t *testing.T) {
		deviation(t, "TYPEDEF_STRUCTURE requires SYMBOL_TYPE_LINK", func(t assert.TestingT) {
			parse(t, moduleScope("/begin TYPEDEF_STRUCTURE typedef_structure \"\" 4\n/end TYPEDEF_STRUCTURE"))
		})
	})

	t.Run("reject/missing size", func(t *testing.T) {
		parseFails(t, moduleScope("/begin TYPEDEF_STRUCTURE typedef_structure \"\" "+
			"SYMBOL_TYPE_LINK \"symbol\"\n/end TYPEDEF_STRUCTURE"))
	})
}

func TestGrammar_STRUCTURE_COMPONENT(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		deviation(t, "STRUCTURE_COMPONENT has no node in the tree", func(t assert.TestingT) {
			assertPreserved(t, moduleScope("/begin TYPEDEF_STRUCTURE typedef_structure \"\" 4 "+
				"SYMBOL_TYPE_LINK \"symbol\"\n"+
				"/begin STRUCTURE_COMPONENT component type_name 0 SYMBOL_TYPE_LINK \"symbol\"\n"+
				"/end STRUCTURE_COMPONENT\n/end TYPEDEF_STRUCTURE"), "STRUCTURE_COMPONENT")
		})
	})

	// SYMBOL_TYPE_LINK is an optional parameter of STRUCTURE_COMPONENT.
	t.Run("accept/without SYMBOL_TYPE_LINK", func(t *testing.T) {
		deviation(t, "STRUCTURE_COMPONENT requires SYMBOL_TYPE_LINK", func(t assert.TestingT) {
			parse(t, moduleScope("/begin TYPEDEF_STRUCTURE typedef_structure \"\" 4 "+
				"SYMBOL_TYPE_LINK \"symbol\"\n"+
				"/begin STRUCTURE_COMPONENT component type_name 0\n"+
				"/end STRUCTURE_COMPONENT\n/end TYPEDEF_STRUCTURE"))
		})
	})

	t.Run("reject/missing address offset", func(t *testing.T) {
		parseFails(t, moduleScope("/begin TYPEDEF_STRUCTURE typedef_structure \"\" 4 "+
			"SYMBOL_TYPE_LINK \"symbol\"\n"+
			"/begin STRUCTURE_COMPONENT component type_name SYMBOL_TYPE_LINK \"symbol\"\n"+
			"/end STRUCTURE_COMPONENT\n/end TYPEDEF_STRUCTURE"))
	})
}
