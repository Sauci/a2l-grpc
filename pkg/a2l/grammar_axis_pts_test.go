package a2l

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrammar_AXIS_PTS(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		module, ok := parseModule(t, "/begin AXIS_PTS axis_pts \"long identifier\" 0x1000 input_quantity "+
			"record_layout 1.5 compu_method 17 -10 10\n/end AXIS_PTS")
		if !ok {
			return
		}

		equalNodes(t, []*AxisPtsType{{
			Name:           identVal("axis_pts"),
			LongIdentifier: strVal("long identifier"),
			Address:        longVal("0x1000"),
			InputQuantity:  identVal("input_quantity"),
			DepositR:       identVal("record_layout"),
			MaxDiff:        floatVal("1.5"),
			Conversion:     identVal("compu_method"),
			MaxAxisPoints:  intVal("17"),
			LowerLimit:     floatVal("-10"),
			UpperLimit:     floatVal("10"),
		}}, module.AXIS_PTS)
	})

	t.Run("optional/single occurrence keywords", func(t *testing.T) {
		axisPts, ok := parseAxisPts(t, `DISPLAY_IDENTIFIER display_identifier
READ_ONLY
FORMAT "%4.2"
DEPOSIT DIFFERENCE
BYTE_ORDER MSB_FIRST
REF_MEMORY_SEGMENT memory_segment
GUARD_RAILS
EXTENDED_LIMITS -20 20
CALIBRATION_ACCESS CALIBRATION
ECU_ADDRESS_EXTENSION 1`)
		if !ok {
			return
		}

		equalNode(t, &AxisPtsType{
			Name:                  identVal("axis_pts"),
			LongIdentifier:        strVal(""),
			Address:               longVal("0x0"),
			InputQuantity:         identVal("input_quantity"),
			DepositR:              identVal("record_layout"),
			MaxDiff:               floatVal("0"),
			Conversion:            identVal("compu_method"),
			MaxAxisPoints:         intVal("8"),
			LowerLimit:            floatVal("0"),
			UpperLimit:            floatVal("0"),
			DISPLAY_IDENTIFIER:    &DisplayIdentifierType{DisplayName: identVal("display_identifier")},
			READ_ONLY:             &ReadOnlyType{Present: true},
			FORMAT:                &FormatType{FormatString: strVal("%4.2")},
			DEPOSIT:               &DepositType{Mode: "DIFFERENCE"},
			BYTE_ORDER:            &ByteOrderType{ByteOrder: "MSB_FIRST"},
			REF_MEMORY_SEGMENT:    &RefMemorySegmentType{Name: identVal("memory_segment")},
			GUARD_RAILS:           &GuardRailsType{Present: true},
			EXTENDED_LIMITS:       &ExtendedLimitsType{LowerLimit: floatVal("-20"), UpperLimit: floatVal("20")},
			CALIBRATION_ACCESS:    &CalibrationAccessType{Type: "CALIBRATION"},
			ECU_ADDRESS_EXTENSION: &EcuAddressExtensionType{Extension: intVal("1")},
		}, axisPts)
	})

	t.Run("optional/FUNCTION_LIST", func(t *testing.T) {
		axisPts, ok := parseAxisPts(t, "/begin FUNCTION_LIST function\n/end FUNCTION_LIST")
		if !ok {
			return
		}

		equalNode(t, &FunctionListType{Name: []*IdentType{identVal("function")}}, axisPts.FUNCTION_LIST)
	})

	t.Run("list/several ANNOTATION", func(t *testing.T) {
		axisPts, ok := parseAxisPts(t,
			"/begin ANNOTATION ANNOTATION_LABEL \"first\" /end ANNOTATION\n"+
				"/begin ANNOTATION ANNOTATION_LABEL \"second\" /end ANNOTATION")
		if !ok {
			return
		}

		assert.Len(t, axisPts.ANNOTATION, 2)
	})

	t.Run("list/several IF_DATA", func(t *testing.T) {
		axisPts, ok := parseAxisPts(t,
			"/begin IF_DATA XCP\n/end IF_DATA\n/begin IF_DATA CANAPE\n/end IF_DATA")
		if !ok {
			return
		}

		equalNodes(t, []*IfDataType{{Name: identVal("XCP")}, {Name: identVal("CANAPE")}}, axisPts.IF_DATA)
	})

	t.Run("reject/missing upper limit", func(t *testing.T) {
		parseFails(t, moduleScope("/begin AXIS_PTS axis_pts \"\" 0x0 input_quantity record_layout 0 "+
			"compu_method 8 0\n/end AXIS_PTS"))
	})

	t.Run("reject/float number of axis points", func(t *testing.T) {
		parseFails(t, moduleScope("/begin AXIS_PTS axis_pts \"\" 0x0 input_quantity record_layout 0 "+
			"compu_method 8.0 0 0\n/end AXIS_PTS"))
	})

	// ASAP2 1.51 declares MONOTONY for AXIS_DESCR only, ASAM MCD-2 MC 1.6.1 (chapter 3.5.18)
	// added it to AXIS_PTS.
	t.Run("optional/MONOTONY", func(t *testing.T) {
		axisPts, ok := parseAxisPts(t, "MONOTONY MON_INCREASE")
		if !ok {
			return
		}

		equalNode(t, &MonotonyType{Monotony: "MON_INCREASE"}, axisPts.MONOTONY)
	})
}

func TestGrammar_AXIS_PTS_REF(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		axisDescr, ok := parseAxisDescr(t, "AXIS_PTS_REF axis_pts")
		if !ok {
			return
		}

		equalNode(t, &AxisPtsRefType{AxisPoints: identVal("axis_pts")}, axisDescr.AXIS_PTS_REF)
	})

	t.Run("reject/string parameter", func(t *testing.T) {
		parseFails(t, axisDescrScope("AXIS_PTS_REF \"axis_pts\""))
	})

	t.Run("reject/missing reference", func(t *testing.T) {
		parseFails(t, axisDescrScope("AXIS_PTS_REF"))
	})
}

func TestGrammar_CURVE_AXIS_REF(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		axisDescr, ok := parseAxisDescr(t, "CURVE_AXIS_REF curve_axis")
		if !ok {
			return
		}

		equalNode(t, &CurveAxisRefType{CurveAxis: identVal("curve_axis")}, axisDescr.CURVE_AXIS_REF)
	})

	t.Run("reject/missing reference", func(t *testing.T) {
		parseFails(t, axisDescrScope("CURVE_AXIS_REF"))
	})
}

func TestGrammar_FIX_AXIS_PAR(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		axisDescr, ok := parseAxisDescr(t, "FIX_AXIS_PAR 0 4 6")
		if !ok {
			return
		}

		equalNode(t, &FixAxisParType{
			Offset:    floatVal("0"),
			Shift:     floatVal("4"),
			Numberapo: intVal("6"),
		}, axisDescr.FIX_AXIS_PAR)
	})

	// Offset and Shift are of type float since ASAM MCD-2 MC 1.6.1 (chapter 1.4.4).
	t.Run("accept/float offset and shift", func(t *testing.T) {
		axisDescr, ok := parseAxisDescr(t, "FIX_AXIS_PAR 0.5 4.25 6")
		if !ok {
			return
		}

		equalNode(t, &FixAxisParType{
			Offset:    floatVal("0.5"),
			Shift:     floatVal("4.25"),
			Numberapo: intVal("6"),
		}, axisDescr.FIX_AXIS_PAR)
	})

	t.Run("reject/missing number of axis points", func(t *testing.T) {
		parseFails(t, axisDescrScope("FIX_AXIS_PAR 0 4"))
	})

	t.Run("reject/float number of axis points", func(t *testing.T) {
		parseFails(t, axisDescrScope("FIX_AXIS_PAR 0 4 6.0"))
	})
}

func TestGrammar_FIX_AXIS_PAR_DIST(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		axisDescr, ok := parseAxisDescr(t, "FIX_AXIS_PAR_DIST 0 100 8")
		if !ok {
			return
		}

		equalNode(t, &FixAxisParDistType{
			Offset:    floatVal("0"),
			Distance:  floatVal("100"),
			Numberapo: intVal("8"),
		}, axisDescr.FIX_AXIS_PAR_DIST)
	})

	t.Run("reject/missing number of axis points", func(t *testing.T) {
		parseFails(t, axisDescrScope("FIX_AXIS_PAR_DIST 0 100"))
	})
}

func TestGrammar_FIX_AXIS_PAR_LIST(t *testing.T) {
	type testCaseType struct {
		name    string
		content string
		values  []*FloatType
	}

	for _, testCase := range []testCaseType{
		{name: "list/no value", content: "", values: nil},
		{name: "list/single value", content: "1.5", values: []*FloatType{floatVal("1.5")}},
		{
			name:    "list/several values",
			content: "0 1.5 3",
			values:  []*FloatType{floatVal("0"), floatVal("1.5"), floatVal("3")},
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			axisDescr, ok := parseAxisDescr(t,
				"/begin FIX_AXIS_PAR_LIST "+testCase.content+"\n/end FIX_AXIS_PAR_LIST")
			if !ok {
				return
			}

			equalNode(t, &FixAxisParListType{AxisPtsValue: testCase.values}, axisDescr.FIX_AXIS_PAR_LIST)
		})
	}

	t.Run("reject/string value", func(t *testing.T) {
		parseFails(t, axisDescrScope("/begin FIX_AXIS_PAR_LIST \"1.5\"\n/end FIX_AXIS_PAR_LIST"))
	})
}

func TestGrammar_MAX_GRAD(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		axisDescr, ok := parseAxisDescr(t, "MAX_GRAD 20.5")
		if !ok {
			return
		}

		equalNode(t, &MaxGradType{MaxGradient: floatVal("20.5")}, axisDescr.MAX_GRAD)
	})

	t.Run("reject/missing gradient", func(t *testing.T) {
		parseFails(t, axisDescrScope("MAX_GRAD"))
	})
}

func TestGrammar_MONOTONY(t *testing.T) {
	for _, monotony := range []string{
		"MON_INCREASE", "MON_DECREASE", "STRICT_INCREASE", "STRICT_DECREASE",
		"MONOTONOUS", "STRICT_MON", "NOT_MON",
	} {
		t.Run("enum/"+monotony, func(t *testing.T) {
			axisDescr, ok := parseAxisDescr(t, "MONOTONY "+monotony)
			if !ok {
				return
			}

			equalNode(t, &MonotonyType{Monotony: monotony}, axisDescr.MONOTONY)
		})
	}

	t.Run("reject/unknown monotony", func(t *testing.T) {
		parseFails(t, axisDescrScope("MONOTONY INCREASING"))
	})

	t.Run("reject/missing monotony", func(t *testing.T) {
		parseFails(t, axisDescrScope("MONOTONY"))
	})
}
