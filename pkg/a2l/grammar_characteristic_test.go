package a2l

// CHARACTERISTIC, AXIS_DESCR and MATRIX_DIM.

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrammar_CHARACTERISTIC(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		module, ok := parseModule(t,
			"/begin CHARACTERISTIC characteristic \"long identifier\" VALUE 0x1234 record_layout 1.5 "+
				"compu_method -10 10\n/end CHARACTERISTIC")
		if !ok {
			return
		}

		equalNodes(t, []*CharacteristicType{{
			Name:           identVal("characteristic"),
			LongIdentifier: strVal("long identifier"),
			Type:           "VALUE",
			Address:        longVal("0x1234"),
			Deposit:        identVal("record_layout"),
			MaxDiff:        floatVal("1.5"),
			Conversion:     identVal("compu_method"),
			LowerLimit:     floatVal("-10"),
			UpperLimit:     floatVal("10"),
		}}, module.CHARACTERISTIC)
	})

	// VALUE, CURVE, MAP, CUBOID, VAL_BLK and ASCII are defined by ASAP2 1.51, CUBE_4 and CUBE_5
	// were added later.
	for _, characteristicType := range []string{
		"ASCII", "CURVE", "MAP", "CUBOID", "CUBE_4", "CUBE_5", "VAL_BLK", "VALUE",
	} {
		t.Run("enum/"+characteristicType, func(t *testing.T) {
			module, ok := parseModule(t,
				"/begin CHARACTERISTIC characteristic \"\" "+characteristicType+
					" 0x0 record_layout 0 compu_method 0 0\n/end CHARACTERISTIC")
			if !ok || !assert.Len(t, module.CHARACTERISTIC, 1) {
				return
			}

			assert.Equal(t, characteristicType, module.CHARACTERISTIC[0].Type)
		})
	}

	t.Run("optional/single occurrence keywords", func(t *testing.T) {
		characteristic, ok := parseCharacteristic(t, `DISPLAY_IDENTIFIER display_identifier
FORMAT "%4.2"
BYTE_ORDER MSB_FIRST
BIT_MASK 0xFF
NUMBER 8
EXTENDED_LIMITS -20 20
READ_ONLY
GUARD_RAILS
MAX_REFRESH 3 10
REF_MEMORY_SEGMENT memory_segment
COMPARISON_QUANTITY measurement
CALIBRATION_ACCESS CALIBRATION
MATRIX_DIM 2 3 4
ECU_ADDRESS_EXTENSION 1`)
		if !ok {
			return
		}

		equalNode(t, &CharacteristicType{
			Name:                  identVal("characteristic"),
			LongIdentifier:        strVal(""),
			Type:                  "VALUE",
			Address:               longVal("0x0"),
			Deposit:               identVal("record_layout"),
			MaxDiff:               floatVal("0"),
			Conversion:            identVal("compu_method"),
			LowerLimit:            floatVal("0"),
			UpperLimit:            floatVal("0"),
			DISPLAY_IDENTIFIER:    &DisplayIdentifierType{DisplayName: identVal("display_identifier")},
			FORMAT:                &FormatType{FormatString: strVal("%4.2")},
			BYTE_ORDER:            &ByteOrderType{ByteOrder: "MSB_FIRST"},
			BIT_MASK:              &BitMaskType{Mask: longVal("0xFF")},
			NUMBER:                &NumberType{Number: intVal("8")},
			EXTENDED_LIMITS:       &ExtendedLimitsType{LowerLimit: floatVal("-20"), UpperLimit: floatVal("20")},
			READ_ONLY:             &ReadOnlyType{Present: true},
			GUARD_RAILS:           &GuardRailsType{Present: true},
			MAX_REFRESH:           &MaxRefreshType{ScalingUnit: intVal("3"), Rate: longVal("10")},
			REF_MEMORY_SEGMENT:    &RefMemorySegmentType{Name: identVal("memory_segment")},
			COMPARISON_QUANTITY:   &ComparisonQuantityType{Name: identVal("measurement")},
			CALIBRATION_ACCESS:    &CalibrationAccessType{Type: "CALIBRATION"},
			MATRIX_DIM:            &MatrixDimType{XDim: intVal("2"), YDim: intVal("3"), ZDim: intVal("4")},
			ECU_ADDRESS_EXTENSION: &EcuAddressExtensionType{Extension: intVal("1")},
		}, characteristic)
	})

	// Deviation: DISCRETE, PHYS_UNIT and STEP_SIZE are parsed and stored in the tree, but
	// CharacteristicType.MarshalA2L does not write them back, so they are lost by a
	// parse/serialize cycle.
	t.Run("optional/keywords dropped by the serializer", func(t *testing.T) {
		deviation(t, "CHARACTERISTIC serializer ignores DISCRETE, PHYS_UNIT and STEP_SIZE",
			func(t assert.TestingT) {
				characteristic, ok := parseCharacteristic(t, "DISCRETE\nPHYS_UNIT \"km/h\"\nSTEP_SIZE 0.5")
				if !ok {
					return
				}

				equalNode(t, &DiscreteType{Present: true}, characteristic.DISCRETE)
				equalNode(t, &PhysUnitType{Unit: strVal("km/h")}, characteristic.PHYS_UNIT)
				equalNode(t, &StepSizeType{StepSize: floatVal("0.5")}, characteristic.STEP_SIZE)
			})
	})

	t.Run("list/several AXIS_DESCR", func(t *testing.T) {
		characteristic, ok := parseCharacteristic(t,
			"/begin AXIS_DESCR STD_AXIS input_x compu_method 1 0 0 /end AXIS_DESCR\n"+
				"/begin AXIS_DESCR STD_AXIS input_y compu_method 1 0 0 /end AXIS_DESCR")
		if !ok {
			return
		}

		if assert.Len(t, characteristic.AXIS_DESCR, 2) {
			assert.Equal(t, "input_x", characteristic.AXIS_DESCR[0].InputQuantity.Value)
			assert.Equal(t, "input_y", characteristic.AXIS_DESCR[1].InputQuantity.Value)
		}
	})

	t.Run("reject/missing upper limit", func(t *testing.T) {
		parseFails(t, moduleScope(
			"/begin CHARACTERISTIC characteristic \"\" VALUE 0x0 record_layout 0 compu_method 0\n/end CHARACTERISTIC"))
	})

	t.Run("reject/unknown type", func(t *testing.T) {
		parseFails(t, moduleScope(
			"/begin CHARACTERISTIC characteristic \"\" CUBE_6 0x0 record_layout 0 compu_method 0 0\n/end CHARACTERISTIC"))
	})

	t.Run("reject/string as address", func(t *testing.T) {
		parseFails(t, moduleScope(
			"/begin CHARACTERISTIC characteristic \"\" VALUE \"0x0\" record_layout 0 compu_method 0 0\n/end CHARACTERISTIC"))
	})
}

func TestGrammar_AXIS_DESCR(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		characteristic, ok := parseCharacteristic(t,
			"/begin AXIS_DESCR STD_AXIS input_quantity compu_method 8 -1.5 1.5\n/end AXIS_DESCR")
		if !ok {
			return
		}

		equalNodes(t, []*AxisDescrType{{
			Attribute:     "STD_AXIS",
			InputQuantity: identVal("input_quantity"),
			Conversion:    identVal("compu_method"),
			MaxAxisPoints: intVal("8"),
			LowerLimit:    floatVal("-1.5"),
			UpperLimit:    floatVal("1.5"),
		}}, characteristic.AXIS_DESCR)
	})

	for _, attribute := range []string{"CURVE_AXIS", "COM_AXIS", "FIX_AXIS", "RES_AXIS", "STD_AXIS"} {
		t.Run("enum/"+attribute, func(t *testing.T) {
			characteristic, ok := parseCharacteristic(t,
				"/begin AXIS_DESCR "+attribute+" input_quantity compu_method 1 0 0\n/end AXIS_DESCR")
			if !ok || !assert.Len(t, characteristic.AXIS_DESCR, 1) {
				return
			}

			assert.Equal(t, attribute, characteristic.AXIS_DESCR[0].Attribute)
		})
	}

	t.Run("optional/single occurrence keywords", func(t *testing.T) {
		axisDescr, ok := parseAxisDescr(t, `READ_ONLY
FORMAT "%4.2"
AXIS_PTS_REF axis_pts
MAX_GRAD 10.5
MONOTONY MON_INCREASE
BYTE_ORDER MSB_FIRST
EXTENDED_LIMITS -20 20
DEPOSIT ABSOLUTE
CURVE_AXIS_REF curve_axis`)
		if !ok {
			return
		}

		equalNode(t, &AxisDescrType{
			Attribute:       "STD_AXIS",
			InputQuantity:   identVal("input_quantity"),
			Conversion:      identVal("compu_method"),
			MaxAxisPoints:   intVal("1"),
			LowerLimit:      floatVal("0"),
			UpperLimit:      floatVal("0"),
			READ_ONLY:       &ReadOnlyType{Present: true},
			FORMAT:          &FormatType{FormatString: strVal("%4.2")},
			AXIS_PTS_REF:    &AxisPtsRefType{AxisPoints: identVal("axis_pts")},
			MAX_GRAD:        &MaxGradType{MaxGradient: floatVal("10.5")},
			MONOTONY:        &MonotonyType{Monotony: "MON_INCREASE"},
			BYTE_ORDER:      &ByteOrderType{ByteOrder: "MSB_FIRST"},
			EXTENDED_LIMITS: &ExtendedLimitsType{LowerLimit: floatVal("-20"), UpperLimit: floatVal("20")},
			DEPOSIT:         &DepositType{Mode: "ABSOLUTE"},
			CURVE_AXIS_REF:  &CurveAxisRefType{CurveAxis: identVal("curve_axis")},
		}, axisDescr)
	})

	// Deviation: as for CHARACTERISTIC, the serializer of AXIS_DESCR ignores these two keywords.
	t.Run("optional/keywords dropped by the serializer", func(t *testing.T) {
		deviation(t, "AXIS_DESCR serializer ignores PHYS_UNIT and STEP_SIZE", func(t assert.TestingT) {
			axisDescr, ok := parseAxisDescr(t, "STEP_SIZE 0.5\nPHYS_UNIT \"km/h\"")
			if !ok {
				return
			}

			equalNode(t, &StepSizeType{StepSize: floatVal("0.5")}, axisDescr.STEP_SIZE)
			equalNode(t, &PhysUnitType{Unit: strVal("km/h")}, axisDescr.PHYS_UNIT)
		})
	})

	t.Run("optional/FIX_AXIS_PAR", func(t *testing.T) {
		axisDescr, ok := parseAxisDescr(t, "FIX_AXIS_PAR 0 4 6")
		if !ok {
			return
		}

		equalNode(t, &FixAxisParType{
			Offset:    intVal("0"),
			Shift:     intVal("4"),
			Numberapo: intVal("6"),
		}, axisDescr.FIX_AXIS_PAR)
	})

	t.Run("optional/FIX_AXIS_PAR_DIST", func(t *testing.T) {
		axisDescr, ok := parseAxisDescr(t, "FIX_AXIS_PAR_DIST 0 100 8")
		if !ok {
			return
		}

		equalNode(t, &FixAxisParDistType{
			Offset:    intVal("0"),
			Distance:  intVal("100"),
			Numberapo: intVal("8"),
		}, axisDescr.FIX_AXIS_PAR_DIST)
	})

	t.Run("optional/FIX_AXIS_PAR_LIST", func(t *testing.T) {
		axisDescr, ok := parseAxisDescr(t, "/begin FIX_AXIS_PAR_LIST 0 1.5 3 /end FIX_AXIS_PAR_LIST")
		if !ok {
			return
		}

		equalNode(t, &FixAxisParListType{
			AxisPtsValue: []*FloatType{floatVal("0"), floatVal("1.5"), floatVal("3")},
		}, axisDescr.FIX_AXIS_PAR_LIST)
	})

	t.Run("reject/missing upper limit", func(t *testing.T) {
		parseFails(t, characteristicScope(
			"/begin AXIS_DESCR STD_AXIS input_quantity compu_method 1 0\n/end AXIS_DESCR"))
	})

	t.Run("reject/unknown attribute", func(t *testing.T) {
		parseFails(t, characteristicScope(
			"/begin AXIS_DESCR ANY_AXIS input_quantity compu_method 1 0 0\n/end AXIS_DESCR"))
	})
}

func TestGrammar_MATRIX_DIM(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		characteristic, ok := parseCharacteristic(t, "MATRIX_DIM 2 3 4")
		if !ok {
			return
		}

		equalNode(t, &MatrixDimType{
			XDim: intVal("2"),
			YDim: intVal("3"),
			ZDim: intVal("4"),
		}, characteristic.MATRIX_DIM)
	})

	t.Run("in MEASUREMENT", func(t *testing.T) {
		module, ok := parseModule(t,
			"/begin MEASUREMENT measurement \"\" UBYTE compu_method 0 0 0 0\nMATRIX_DIM 2 3 4\n/end MEASUREMENT")
		if !ok || !assert.Len(t, module.MEASUREMENT, 1) {
			return
		}

		equalNode(t, &MatrixDimType{
			XDim: intVal("2"),
			YDim: intVal("3"),
			ZDim: intVal("4"),
		}, module.MEASUREMENT[0].MATRIX_DIM)
	})

	// Deviation: ASAP2 1.51 requires the three dimensions. The grammar makes yDim and zDim
	// optional, and the resulting tree cannot be serialized again.
	t.Run("reject/single dimension", func(t *testing.T) {
		deviation(t, "MATRIX_DIM accepts less than three dimensions", func(t assert.TestingT) {
			parseFails(t, characteristicScope("MATRIX_DIM 2"))
		})
	})

	t.Run("reject/four dimensions", func(t *testing.T) {
		parseFails(t, characteristicScope("MATRIX_DIM 2 3 4 5"))
	})

	t.Run("reject/float dimension", func(t *testing.T) {
		parseFails(t, characteristicScope("MATRIX_DIM 2.0 3 4"))
	})
}
