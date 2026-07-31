package a2l

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrammar_UNIT(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		module, ok := parseModule(t,
			"/begin UNIT unit \"long identifier\" \"km/h\" EXTENDED_SI\n/end UNIT")
		if !ok {
			return
		}

		equalNodes(t, []*UnitType{{
			Name:           identVal("unit"),
			LongIdentifier: strVal("long identifier"),
			Display:        strVal("km/h"),
			Type:           "EXTENDED_SI",
		}}, module.UNIT)
	})

	for _, unitType := range []string{"DERIVED", "EXTENDED_SI"} {
		t.Run("enum/"+unitType, func(t *testing.T) {
			module, ok := parseModule(t, "/begin UNIT unit \"\" \"display\" "+unitType+"\n/end UNIT")
			if !ok || !assert.Len(t, module.UNIT, 1) {
				return
			}

			assert.Equal(t, unitType, module.UNIT[0].Type)
		})
	}

	t.Run("optional/all parameters", func(t *testing.T) {
		unit, ok := parseUnit(t, `SI_EXPONENTS 1 0 -1 0 0 0 0
REF_UNIT other_unit
UNIT_CONVERSION 3.6 0`)
		if !ok {
			return
		}

		equalNode(t, &UnitType{
			Name:           identVal("unit"),
			LongIdentifier: strVal(""),
			Display:        strVal("display"),
			Type:           "DERIVED",
			SI_EXPONENTS: &SiExponentsType{
				Length:            intVal("1"),
				Mass:              intVal("0"),
				Time:              intVal("-1"),
				ElectricCurrent:   intVal("0"),
				Temperature:       intVal("0"),
				AmountOfSubstance: intVal("0"),
				LuminousIntensity: intVal("0"),
			},
			REF_UNIT:        &RefUnitType{Unit: identVal("other_unit")},
			UNIT_CONVERSION: &UnitConversionType{Gradient: floatVal("3.6"), Offset: floatVal("0")},
		}, unit)
	})

	t.Run("reject/unknown type", func(t *testing.T) {
		parseFails(t, moduleScope("/begin UNIT unit \"\" \"display\" SI\n/end UNIT"))
	})

	t.Run("reject/missing type", func(t *testing.T) {
		parseFails(t, moduleScope("/begin UNIT unit \"\" \"display\"\n/end UNIT"))
	})

	t.Run("reject/identifier as display string", func(t *testing.T) {
		parseFails(t, moduleScope("/begin UNIT unit \"\" display DERIVED\n/end UNIT"))
	})
}

func TestGrammar_SI_EXPONENTS(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		unit, ok := parseUnit(t, "SI_EXPONENTS 1 2 3 4 5 6 7")
		if !ok {
			return
		}

		equalNode(t, &SiExponentsType{
			Length:            intVal("1"),
			Mass:              intVal("2"),
			Time:              intVal("3"),
			ElectricCurrent:   intVal("4"),
			Temperature:       intVal("5"),
			AmountOfSubstance: intVal("6"),
			LuminousIntensity: intVal("7"),
		}, unit.SI_EXPONENTS)
	})

	t.Run("reject/six exponents", func(t *testing.T) {
		parseFails(t, unitScope("SI_EXPONENTS 1 2 3 4 5 6"))
	})

	t.Run("reject/eight exponents", func(t *testing.T) {
		parseFails(t, unitScope("SI_EXPONENTS 1 2 3 4 5 6 7 8"))
	})

	t.Run("reject/float exponent", func(t *testing.T) {
		parseFails(t, unitScope("SI_EXPONENTS 1.0 2 3 4 5 6 7"))
	})
}

func TestGrammar_UNIT_CONVERSION(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		unit, ok := parseUnit(t, "UNIT_CONVERSION 3.6 -273.15")
		if !ok {
			return
		}

		equalNode(t, &UnitConversionType{
			Gradient: floatVal("3.6"),
			Offset:   floatVal("-273.15"),
		}, unit.UNIT_CONVERSION)
	})

	t.Run("reject/missing offset", func(t *testing.T) {
		parseFails(t, unitScope("UNIT_CONVERSION 3.6"))
	})
}

func TestGrammar_FRAME(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		module, ok := parseModule(t, "/begin FRAME frame \"long identifier\" 3 100\n/end FRAME")
		if !ok {
			return
		}

		equalNode(t, &FrameType{
			Name:           identVal("frame"),
			LongIdentifier: strVal("long identifier"),
			ScalingUnit:    intVal("3"),
			Rate:           longVal("100"),
		}, module.FRAME)
	})

	t.Run("optional/FRAME_MEASUREMENT", func(t *testing.T) {
		frame, ok := parseFrame(t, "FRAME_MEASUREMENT first second")
		if !ok {
			return
		}

		equalNode(t, &FrameMeasurementType{
			Identifier: []*IdentType{identVal("first"), identVal("second")},
		}, frame.FRAME_MEASUREMENT)
	})

	t.Run("list/several IF_DATA", func(t *testing.T) {
		frame, ok := parseFrame(t, "/begin IF_DATA XCP\n/end IF_DATA\n/begin IF_DATA CANAPE\n/end IF_DATA")
		if !ok {
			return
		}

		equalNodes(t, []*IfDataType{{Name: identVal("XCP")}, {Name: identVal("CANAPE")}}, frame.IF_DATA)
	})

	// Chapter 6.3.89 declares [-> FRAME]: a module contains at most one frame.
	t.Run("reject/repeated FRAME", func(t *testing.T) {
		parseFails(t, moduleScope("/begin FRAME first \"\" 3 100\n/end FRAME\n"+
			"/begin FRAME second \"\" 3 100\n/end FRAME"))
	})

	t.Run("reject/missing rate", func(t *testing.T) {
		parseFails(t, moduleScope("/begin FRAME frame \"\" 3\n/end FRAME"))
	})

	t.Run("reject/float scaling unit", func(t *testing.T) {
		parseFails(t, moduleScope("/begin FRAME frame \"\" 3.0 100\n/end FRAME"))
	})
}

func TestGrammar_FRAME_MEASUREMENT(t *testing.T) {
	// FRAME_MEASUREMENT is not delimited by /begin and /end.
	type testCaseType struct {
		name        string
		content     string
		identifiers []*IdentType
	}

	for _, testCase := range []testCaseType{
		{name: "list/no identifier", content: "", identifiers: nil},
		{name: "list/single identifier", content: "first", identifiers: []*IdentType{identVal("first")}},
		{
			name:        "list/several identifiers",
			content:     "first second third",
			identifiers: []*IdentType{identVal("first"), identVal("second"), identVal("third")},
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			frame, ok := parseFrame(t, "FRAME_MEASUREMENT "+testCase.content)
			if !ok {
				return
			}

			equalNode(t, &FrameMeasurementType{Identifier: testCase.identifiers}, frame.FRAME_MEASUREMENT)
		})
	}

	t.Run("reject/string instead of identifier", func(t *testing.T) {
		parseFails(t, frameScope("FRAME_MEASUREMENT \"measurement\""))
	})

	t.Run("reject/block delimiters", func(t *testing.T) {
		parseFails(t, frameScope("/begin FRAME_MEASUREMENT measurement\n/end FRAME_MEASUREMENT"))
	})
}
