package a2l

// IF_DATA. The content of an IF_DATA block is described by the A2ML of the module, so the
// grammar stores it as a generic blob instead of validating it.

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrammar_IF_DATA(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		module, ok := parseModule(t, "/begin IF_DATA XCP\n/end IF_DATA")
		if !ok {
			return
		}

		equalNodes(t, []*IfDataType{{Name: identVal("XCP")}}, module.IF_DATA)
	})

	t.Run("list/several IF_DATA", func(t *testing.T) {
		module, ok := parseModule(t, "/begin IF_DATA XCP\n/end IF_DATA\n/begin IF_DATA CANAPE\n/end IF_DATA")
		if !ok {
			return
		}

		equalNodes(t, []*IfDataType{
			{Name: identVal("XCP")},
			{Name: identVal("CANAPE")},
		}, module.IF_DATA)
	})

	t.Run("blob/scalar parameters", func(t *testing.T) {
		module, ok := parseModule(t, "/begin IF_DATA XCP identifier \"string\" 12 0x0C 1.5\n/end IF_DATA")
		if !ok {
			return
		}

		equalNodes(t, []*IfDataType{{
			Name: identVal("XCP"),
			Blob: []*GenericParameterType{
				{Oneof: &GenericParameterType_Identifier{Identifier: identVal("identifier")}},
				{Oneof: &GenericParameterType_String_{String_: strVal("string")}},
				{Oneof: &GenericParameterType_Long{Long: longVal("12")}},
				{Oneof: &GenericParameterType_Long{Long: longVal("0x0C")}},
				{Oneof: &GenericParameterType_Float{Float: floatVal("1.5")}},
			},
		}}, module.IF_DATA)
	})

	t.Run("blob/nested block", func(t *testing.T) {
		module, ok := parseModule(t,
			"/begin IF_DATA XCP\n/begin PROTOCOL_LAYER 0x100 1000\n/end PROTOCOL_LAYER\n/end IF_DATA")
		if !ok {
			return
		}

		equalNodes(t, []*IfDataType{{
			Name: identVal("XCP"),
			Blob: []*GenericParameterType{{
				Oneof: &GenericParameterType_Generic{Generic: &GenericNodeType{
					Name: identVal("PROTOCOL_LAYER"),
					Element: []*GenericParameterType{
						{Oneof: &GenericParameterType_Long{Long: longVal("0x100")}},
						{Oneof: &GenericParameterType_Long{Long: longVal("1000")}},
					},
				}},
			}},
		}}, module.IF_DATA)
	})

	t.Run("blob/interface specific element", func(t *testing.T) {
		module, ok := parseModule(t, "/begin IF_DATA XCP\n/begin DAQ EVENT 1\n/end DAQ\n/end IF_DATA")
		if !ok || !assert.Len(t, module.IF_DATA, 1) {
			return
		}

		assert.Len(t, module.IF_DATA[0].Blob, 1)
	})

	// Chapter 5.2: "Within the AML own name spaces are used. In this case it is allowed to reuse
	// ASAM MCD-2 MC keyword names. The definitions from the AML are exclusively valid in IF_DATA."
	// The lexer is shared by the A2L and IF_DATA grammars, so a blob element spelled like an A2L
	// keyword arrives here as that keyword and must be accepted as an identifier again.
	t.Run("blob/A2L keyword used as a blob element", func(t *testing.T) {
		for _, keyword := range []string{"VERSION", "RESERVED", "UNIT", "ECU", "IDENTIFICATION"} {
			module, ok := parseModule(t, "/begin IF_DATA XCP "+keyword+" 1\n/end IF_DATA")
			if !ok || !assert.Len(t, module.IF_DATA, 1, "keyword %s", keyword) {
				continue
			}

			if assert.Len(t, module.IF_DATA[0].Blob, 2, "keyword %s", keyword) {
				equalNode(t, identVal(keyword), module.IF_DATA[0].Blob[0].GetIdentifier())
			}
		}
	})

	t.Run("blob/A2L keyword used as the name of a nested block", func(t *testing.T) {
		module, ok := parseModule(t,
			"/begin IF_DATA ASAP1B_CAN\n/begin CHECKSUM \"checksum.dll\"\n/end CHECKSUM\n/end IF_DATA")
		if !ok || !assert.Len(t, module.IF_DATA, 1) {
			return
		}

		if assert.Len(t, module.IF_DATA[0].Blob, 1) {
			equalNode(t, identVal("CHECKSUM"), module.IF_DATA[0].Blob[0].GetGeneric().GetName())
		}
	})

	// The name of the IF_DATA block itself is an identifier of the same namespace.
	t.Run("A2L keyword used as the name of the IF_DATA block", func(t *testing.T) {
		module, ok := parseModule(t, "/begin IF_DATA UNIT\n/end IF_DATA")
		if !ok || !assert.Len(t, module.IF_DATA, 1) {
			return
		}

		equalNode(t, identVal("UNIT"), module.IF_DATA[0].Name)
	})

	// The following cases reproduce the canonical IF_DATA forms defined by ASAP2 1.51: the
	// interface specific keywords are not part of the A2L grammar itself, they are described by
	// the A2ML of the module, so they must all be representable by the generic blob.

	// Chapter 6.3.74, IF_DATA ASAP1B_ADDRESS (MEASUREMENT): a tagged scalar.
	t.Run("blob/ASAP1B_ADDRESS with KP_BLOB address", func(t *testing.T) {
		module, ok := parseModule(t,
			"/begin MEASUREMENT measurement \"\" UBYTE compu_method 0 0 0 0\n"+
				"/begin IF_DATA ASAP1B_ADDRESS KP_BLOB 0x1234\n/end IF_DATA\n/end MEASUREMENT")
		if !ok || !assert.Len(t, module.MEASUREMENT, 1) {
			return
		}

		equalNodes(t, []*IfDataType{{
			Name: identVal("ASAP1B_ADDRESS"),
			Blob: []*GenericParameterType{
				{Oneof: &GenericParameterType_Identifier{Identifier: identVal("KP_BLOB")}},
				{Oneof: &GenericParameterType_Long{Long: longVal("0x1234")}},
			},
		}}, module.MEASUREMENT[0].IF_DATA)
	})

	// Chapter 6.3.75, IF_DATA ASAP1B_CAN (MEASUREMENT): a nested KP_BLOB block with the optional
	// MULTIPLEX keyword (chapter 6.3.93).
	t.Run("blob/ASAP1B_CAN with nested KP_BLOB and MULTIPLEX", func(t *testing.T) {
		parse(t, measurementScope(`/begin IF_DATA ASAP1B_CAN
/begin KP_BLOB message_name 0x100 8 "sender" 0 16
MULTIPLEX 16 8 1
/end KP_BLOB
/end IF_DATA`))
	})

	// Chapter 6.3.77, IF_DATA (MEMORY_SEGMENT): ADDRESS_MAPPING (chapter 6.3.4).
	t.Run("blob/ADDRESS_MAPPING", func(t *testing.T) {
		parse(t, modParScope("/begin MEMORY_SEGMENT memory_segment \"\" DATA RAM INTERN 0x0 0x100 "+
			"-1 -1 -1 -1 -1\n/begin IF_DATA ASAP1B_EXAMPLE\n"+
			"ADDRESS_MAPPING 0x10000 0x20000 0x100\n/end IF_DATA\n/end MEMORY_SEGMENT"))
	})

	// Chapter 6.3.78, IF_DATA (MODULE): SOURCE (6.3.120), RASTER (6.3.103), EVENT_GROUP (6.3.53),
	// SEED_KEY (6.3.116) and CHECKSUM (6.3.28).
	t.Run("blob/module interface data", func(t *testing.T) {
		parse(t, moduleScope(`/begin IF_DATA ASAP1B_EXAMPLE
/begin SOURCE time_synchronous_10ms 3 10
/end SOURCE
/begin RASTER "time synchronous 10ms" "10ms" 1 3 10
/end RASTER
/begin EVENT_GROUP "time synchronous" "sync" 1
/end EVENT_GROUP
/begin SEED_KEY "cal.dll" "daq.dll" "pgm.dll"
/end SEED_KEY
/begin CHECKSUM "checksum.dll"
/end CHECKSUM
/end IF_DATA`))
	})

	t.Run("in CHARACTERISTIC", func(t *testing.T) {
		characteristic, ok := parseCharacteristic(t, "/begin IF_DATA XCP\n/end IF_DATA")
		if !ok {
			return
		}

		equalNodes(t, []*IfDataType{{Name: identVal("XCP")}}, characteristic.IF_DATA)
	})

	t.Run("reject/missing name", func(t *testing.T) {
		parseFails(t, moduleScope("/begin IF_DATA\n/end IF_DATA"))
	})

	t.Run("reject/unbalanced nested block", func(t *testing.T) {
		parseFails(t, moduleScope("/begin IF_DATA XCP\n/begin PROTOCOL_LAYER\n/end IF_DATA"))
	})
}
