package a2l

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrammar_MEASUREMENT(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		module, ok := parseModule(t, "/begin MEASUREMENT measurement \"long identifier\" SWORD "+
			"compu_method 16 1.5 -100 100\n/end MEASUREMENT")
		if !ok {
			return
		}

		equalNodes(t, []*MeasurementType{{
			Name:           identVal("measurement"),
			LongIdentifier: strVal("long identifier"),
			DataType:       dataTypeVal("SWORD"),
			Conversion:     identVal("compu_method"),
			Resolution:     intVal("16"),
			Accuracy:       floatVal("1.5"),
			LowerLimit:     floatVal("-100"),
			UpperLimit:     floatVal("100"),
		}}, module.MEASUREMENT)
	})

	for _, dataType := range []string{
		"UBYTE", "SBYTE", "UWORD", "SWORD", "ULONG", "SLONG", "FLOAT32_IEEE", "FLOAT64_IEEE",
	} {
		t.Run("enum/data type "+dataType, func(t *testing.T) {
			module, ok := parseModule(t, "/begin MEASUREMENT measurement \"\" "+dataType+
				" compu_method 0 0 0 0\n/end MEASUREMENT")
			if !ok || !assert.Len(t, module.MEASUREMENT, 1) {
				return
			}

			equalNode(t, dataTypeVal(dataType), module.MEASUREMENT[0].DataType)
		})
	}

	t.Run("optional/single occurrence keywords", func(t *testing.T) {
		measurement, ok := parseMeasurement(t, `DISPLAY_IDENTIFIER display_identifier
READ_WRITE
FORMAT "%4.2"
ARRAY_SIZE 8
BIT_MASK 0xFF
BYTE_ORDER MSB_FIRST
MAX_REFRESH 3 10
ECU_ADDRESS 0x1000
ERROR_MASK 0x0F
REF_MEMORY_SEGMENT memory_segment
MATRIX_DIM 2 3 4
ECU_ADDRESS_EXTENSION 1`)
		if !ok {
			return
		}

		equalNode(t, &MeasurementType{
			Name:                  identVal("measurement"),
			LongIdentifier:        strVal(""),
			DataType:              dataTypeVal("UBYTE"),
			Conversion:            identVal("compu_method"),
			Resolution:            intVal("0"),
			Accuracy:              floatVal("0"),
			LowerLimit:            floatVal("0"),
			UpperLimit:            floatVal("0"),
			DISPLAY_IDENTIFIER:    &DisplayIdentifierType{DisplayName: identVal("display_identifier")},
			READ_WRITE:            &ReadWriteType{Present: true},
			FORMAT:                &FormatType{FormatString: strVal("%4.2")},
			ARRAY_SIZE:            &ArraySizeType{Number: intVal("8")},
			BIT_MASK:              &BitMaskType{Mask: uLongVal("0xFF")},
			BYTE_ORDER:            &ByteOrderType{ByteOrder: "MSB_FIRST"},
			MAX_REFRESH:           &MaxRefreshType{ScalingUnit: intVal("3"), Rate: longVal("10")},
			ECU_ADDRESS:           &EcuAddressType{Address: longVal("0x1000")},
			ERROR_MASK:            &ErrorMaskType{Mask: uLongVal("0x0F")},
			REF_MEMORY_SEGMENT:    &RefMemorySegmentType{Name: identVal("memory_segment")},
			MATRIX_DIM:            &MatrixDimType{XDim: intVal("2"), YDim: intVal("3"), ZDim: intVal("4")},
			ECU_ADDRESS_EXTENSION: &EcuAddressExtensionType{Extension: intVal("1")},
		}, measurement)
	})

	t.Run("optional/FUNCTION_LIST", func(t *testing.T) {
		measurement, ok := parseMeasurement(t, "/begin FUNCTION_LIST first second\n/end FUNCTION_LIST")
		if !ok {
			return
		}

		equalNode(t, &FunctionListType{Name: []*IdentType{identVal("first"), identVal("second")}},
			measurement.FUNCTION_LIST)
	})

	t.Run("list/several ANNOTATION", func(t *testing.T) {
		measurement, ok := parseMeasurement(t,
			"/begin ANNOTATION ANNOTATION_LABEL \"first\" /end ANNOTATION\n"+
				"/begin ANNOTATION ANNOTATION_LABEL \"second\" /end ANNOTATION")
		if !ok {
			return
		}

		assert.Len(t, measurement.ANNOTATION, 2)
	})

	t.Run("list/several IF_DATA", func(t *testing.T) {
		measurement, ok := parseMeasurement(t,
			"/begin IF_DATA XCP\n/end IF_DATA\n/begin IF_DATA CANAPE\n/end IF_DATA")
		if !ok {
			return
		}

		equalNodes(t, []*IfDataType{{Name: identVal("XCP")}, {Name: identVal("CANAPE")}}, measurement.IF_DATA)
	})

	t.Run("reject/missing upper limit", func(t *testing.T) {
		parseFails(t, moduleScope(
			"/begin MEASUREMENT measurement \"\" UBYTE compu_method 0 0 0\n/end MEASUREMENT"))
	})

	t.Run("reject/unknown data type", func(t *testing.T) {
		parseFails(t, moduleScope(
			"/begin MEASUREMENT measurement \"\" UINT24 compu_method 0 0 0 0\n/end MEASUREMENT"))
	})

	t.Run("reject/data size instead of data type", func(t *testing.T) {
		parseFails(t, moduleScope(
			"/begin MEASUREMENT measurement \"\" WORD compu_method 0 0 0 0\n/end MEASUREMENT"))
	})

	// A CHARACTERISTIC keyword must not be accepted inside a MEASUREMENT.
	t.Run("reject/GUARD_RAILS", func(t *testing.T) {
		parseFails(t, measurementScope("GUARD_RAILS"))
	})
}

func TestGrammar_ARRAY_SIZE(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		measurement, ok := parseMeasurement(t, "ARRAY_SIZE 8")
		if !ok {
			return
		}

		equalNode(t, &ArraySizeType{Number: intVal("8")}, measurement.ARRAY_SIZE)
	})

	t.Run("reject/float number", func(t *testing.T) {
		parseFails(t, measurementScope("ARRAY_SIZE 8.0"))
	})

	t.Run("reject/missing number", func(t *testing.T) {
		parseFails(t, measurementScope("ARRAY_SIZE"))
	})
}

func TestGrammar_BIT_OPERATION(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		measurement, ok := parseMeasurement(t, "/begin BIT_OPERATION\n/end BIT_OPERATION")
		if !ok {
			return
		}

		equalNode(t, &BitOperationType{}, measurement.BIT_OPERATION)
	})

	t.Run("optional/all parameters", func(t *testing.T) {
		measurement, ok := parseMeasurement(t, `/begin BIT_OPERATION
LEFT_SHIFT 4
RIGHT_SHIFT 2
SIGN_EXTEND
/end BIT_OPERATION`)
		if !ok {
			return
		}

		equalNode(t, &BitOperationType{
			LEFT_SHIFT:  &LeftShiftType{BitCount: longVal("4")},
			RIGHT_SHIFT: &RightShiftType{BitCount: longVal("2")},
			SIGN_EXTEND: &SignExtendType{Present: true},
		}, measurement.BIT_OPERATION)
	})

	t.Run("reject/missing /end", func(t *testing.T) {
		parseFails(t, measurementScope("/begin BIT_OPERATION LEFT_SHIFT 4"))
	})
}

func TestGrammar_LEFT_SHIFT(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		measurement, ok := parseMeasurement(t, "/begin BIT_OPERATION LEFT_SHIFT 4 /end BIT_OPERATION")
		if !ok || !assert.NotNil(t, measurement.BIT_OPERATION) {
			return
		}

		equalNode(t, &LeftShiftType{BitCount: longVal("4")}, measurement.BIT_OPERATION.LEFT_SHIFT)
	})

	t.Run("reject/missing bit count", func(t *testing.T) {
		parseFails(t, measurementScope("/begin BIT_OPERATION LEFT_SHIFT /end BIT_OPERATION"))
	})
}

func TestGrammar_RIGHT_SHIFT(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		measurement, ok := parseMeasurement(t, "/begin BIT_OPERATION RIGHT_SHIFT 4 /end BIT_OPERATION")
		if !ok || !assert.NotNil(t, measurement.BIT_OPERATION) {
			return
		}

		equalNode(t, &RightShiftType{BitCount: longVal("4")}, measurement.BIT_OPERATION.RIGHT_SHIFT)
	})

	t.Run("reject/missing bit count", func(t *testing.T) {
		parseFails(t, measurementScope("/begin BIT_OPERATION RIGHT_SHIFT /end BIT_OPERATION"))
	})
}

func TestGrammar_SIGN_EXTEND(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		measurement, ok := parseMeasurement(t, "/begin BIT_OPERATION SIGN_EXTEND /end BIT_OPERATION")
		if !ok || !assert.NotNil(t, measurement.BIT_OPERATION) {
			return
		}

		equalNode(t, &SignExtendType{Present: true}, measurement.BIT_OPERATION.SIGN_EXTEND)
	})

	t.Run("reject/with parameter", func(t *testing.T) {
		parseFails(t, measurementScope("/begin BIT_OPERATION SIGN_EXTEND 1 /end BIT_OPERATION"))
	})
}

func TestGrammar_ECU_ADDRESS(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		measurement, ok := parseMeasurement(t, "ECU_ADDRESS 0x12345678")
		if !ok {
			return
		}

		equalNode(t, &EcuAddressType{Address: longVal("0x12345678")}, measurement.ECU_ADDRESS)
	})

	t.Run("reject/missing address", func(t *testing.T) {
		parseFails(t, measurementScope("ECU_ADDRESS"))
	})
}

func TestGrammar_ERROR_MASK(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		measurement, ok := parseMeasurement(t, "ERROR_MASK 0x0F")
		if !ok {
			return
		}

		equalNode(t, &ErrorMaskType{Mask: uLongVal("0x0F")}, measurement.ERROR_MASK)
	})

	t.Run("reject/missing mask", func(t *testing.T) {
		parseFails(t, measurementScope("ERROR_MASK"))
	})
}

func TestGrammar_READ_WRITE(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		measurement, ok := parseMeasurement(t, "READ_WRITE")
		if !ok {
			return
		}

		equalNode(t, &ReadWriteType{Present: true}, measurement.READ_WRITE)
	})

	t.Run("reject/with parameter", func(t *testing.T) {
		parseFails(t, measurementScope("READ_WRITE 1"))
	})

	// READ_WRITE is defined for MEASUREMENT only, an adjustable object uses READ_ONLY.
	t.Run("reject/in CHARACTERISTIC", func(t *testing.T) {
		parseFails(t, characteristicScope("READ_WRITE"))
	})
}

func TestGrammar_VIRTUAL(t *testing.T) {
	type testCaseType struct {
		name     string
		content  string
		channels []*IdentType
	}

	for _, testCase := range []testCaseType{
		{name: "list/no measuring channel", content: "", channels: nil},
		{name: "list/single measuring channel", content: "first", channels: []*IdentType{identVal("first")}},
		{
			name:     "list/several measuring channels",
			content:  "first second",
			channels: []*IdentType{identVal("first"), identVal("second")},
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			measurement, ok := parseMeasurement(t, "/begin VIRTUAL "+testCase.content+"\n/end VIRTUAL")
			if !ok {
				return
			}

			equalNode(t, &VirtualType{MeasuringChannel: testCase.channels}, measurement.VIRTUAL)
		})
	}

	t.Run("reject/string parameter", func(t *testing.T) {
		parseFails(t, measurementScope("/begin VIRTUAL \"measurement\"\n/end VIRTUAL"))
	})
}

func TestGrammar_LAYOUT(t *testing.T) {
	for _, indexMode := range []string{"ROW_DIR", "COLUMN_DIR"} {
		t.Run("enum/"+indexMode, func(t *testing.T) {
			measurement, ok := parseMeasurement(t, "LAYOUT "+indexMode)
			if !ok {
				return
			}

			equalNode(t, &LayoutType{IndexMode: indexMode}, measurement.LAYOUT)
		})
	}

	t.Run("reject/unknown index mode", func(t *testing.T) {
		parseFails(t, measurementScope("LAYOUT ALTERNATE_WITH_X"))
	})
}
