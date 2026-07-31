package a2l

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrammar_RECORD_LAYOUT(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		module, ok := parseModule(t, "/begin RECORD_LAYOUT record_layout\n/end RECORD_LAYOUT")
		if !ok {
			return
		}

		equalNodes(t, []*RecordLayoutType{{Name: identVal("record_layout")}}, module.RECORD_LAYOUT)
	})

	t.Run("optional/axis component keywords", func(t *testing.T) {
		recordLayout, ok := parseRecordLayout(t, `FNC_VALUES 1 UBYTE COLUMN_DIR DIRECT
IDENTIFICATION 2 UWORD
AXIS_PTS_X 3 UBYTE INDEX_INCR DIRECT
AXIS_PTS_Y 4 UBYTE INDEX_DECR PBYTE
NO_AXIS_PTS_X 5 UBYTE
NO_AXIS_PTS_Y 6 UWORD
FIX_NO_AXIS_PTS_X 8
SRC_ADDR_X 9 ULONG
RIP_ADDR_W 10 ULONG
SHIFT_OP_X 11 UBYTE
OFFSET_X 12 UBYTE
DIST_OP_X 13 UBYTE`)
		if !ok {
			return
		}

		equalNode(t, &RecordLayoutType{
			Name:              identVal("record_layout"),
			FNC_VALUES:        &FncValuesType{Position: intVal("1"), DataType: dataTypeVal("UBYTE"), IndexMode: "COLUMN_DIR", AddressType: addrTypeVal("DIRECT")},
			IDENTIFICATION:    &IdentificationType{Position: intVal("2"), DataType: dataTypeVal("UWORD")},
			AXIS_PTS_X:        &AxisPtsXType{Position: intVal("3"), DataType: dataTypeVal("UBYTE"), IndexIncr: indexOrderVal("INDEX_INCR"), Addressing: addrTypeVal("DIRECT")},
			AXIS_PTS_Y:        &AxisPtsYType{Position: intVal("4"), DataType: dataTypeVal("UBYTE"), IndexIncr: indexOrderVal("INDEX_DECR"), Addressing: addrTypeVal("PBYTE")},
			NO_AXIS_PTS_X:     &NoAxisPtsXType{Position: intVal("5"), DataType: dataTypeVal("UBYTE")},
			NO_AXIS_PTS_Y:     &NoAxisPtsYType{Position: intVal("6"), DataType: dataTypeVal("UWORD")},
			FIX_NO_AXIS_PTS_X: &FixNoAxisPtsXType{NumberOfAxisPoints: intVal("8")},
			SRC_ADDR_X:        &SrcAddrXType{Position: intVal("9"), DataType: dataTypeVal("ULONG")},
			RIP_ADDR_W:        &RipAddrWType{Position: intVal("10"), DataType: dataTypeVal("ULONG")},
			SHIFT_OP_X:        &ShiftOpXType{Position: intVal("11"), DataType: dataTypeVal("UBYTE")},
			OFFSET_X:          &OffsetXType{Position: intVal("12"), DataType: dataTypeVal("UBYTE")},
			DIST_OP_X:         &DistOpXType{Position: intVal("13"), DataType: dataTypeVal("UBYTE")},
		}, recordLayout)
	})

	t.Run("list/several RESERVED", func(t *testing.T) {
		recordLayout, ok := parseRecordLayout(t, "RESERVED 1 BYTE\nRESERVED 2 WORD\nRESERVED 3 LONG")
		if !ok {
			return
		}

		equalNodes(t, []*ReservedType{
			{Position: intVal("1"), DataSize: "BYTE"},
			{Position: intVal("2"), DataSize: "WORD"},
			{Position: intVal("3"), DataSize: "LONG"},
		}, recordLayout.RESERVED)
	})

	t.Run("reject/missing name", func(t *testing.T) {
		parseFails(t, moduleScope("/begin RECORD_LAYOUT\n/end RECORD_LAYOUT"))
	})
}

func TestGrammar_FNC_VALUES(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		recordLayout, ok := parseRecordLayout(t, "FNC_VALUES 1 FLOAT32_IEEE ROW_DIR DIRECT")
		if !ok {
			return
		}

		equalNode(t, &FncValuesType{
			Position:    intVal("1"),
			DataType:    dataTypeVal("FLOAT32_IEEE"),
			IndexMode:   "ROW_DIR",
			AddressType: addrTypeVal("DIRECT"),
		}, recordLayout.FNC_VALUES)
	})

	for _, indexMode := range []string{
		"ALTERNATE_CURVES", "ALTERNATE_WITH_X", "ALTERNATE_WITH_Y", "COLUMN_DIR", "ROW_DIR",
	} {
		t.Run("enum/index mode "+indexMode, func(t *testing.T) {
			recordLayout, ok := parseRecordLayout(t, "FNC_VALUES 1 UBYTE "+indexMode+" DIRECT")
			if !ok || !assert.NotNil(t, recordLayout.FNC_VALUES) {
				return
			}

			assert.Equal(t, indexMode, recordLayout.FNC_VALUES.IndexMode)
		})
	}

	for _, dataType := range []string{
		"UBYTE", "SBYTE", "UWORD", "SWORD", "ULONG", "SLONG",
		"A_UINT64", "A_INT64", "FLOAT32_IEEE", "FLOAT64_IEEE",
	} {
		t.Run("enum/data type "+dataType, func(t *testing.T) {
			recordLayout, ok := parseRecordLayout(t, "FNC_VALUES 1 "+dataType+" ROW_DIR DIRECT")
			if !ok || !assert.NotNil(t, recordLayout.FNC_VALUES) {
				return
			}

			equalNode(t, dataTypeVal(dataType), recordLayout.FNC_VALUES.DataType)
		})
	}

	for _, addressType := range []string{"PBYTE", "PWORD", "PLONG", "DIRECT"} {
		t.Run("enum/address type "+addressType, func(t *testing.T) {
			recordLayout, ok := parseRecordLayout(t, "FNC_VALUES 1 UBYTE ROW_DIR "+addressType)
			if !ok || !assert.NotNil(t, recordLayout.FNC_VALUES) {
				return
			}

			equalNode(t, addrTypeVal(addressType), recordLayout.FNC_VALUES.AddressType)
		})
	}

	t.Run("reject/missing address type", func(t *testing.T) {
		parseFails(t, recordLayoutScope("FNC_VALUES 1 UBYTE ROW_DIR"))
	})

	t.Run("reject/unknown data type", func(t *testing.T) {
		parseFails(t, recordLayoutScope("FNC_VALUES 1 UINT24 ROW_DIR DIRECT"))
	})

	t.Run("reject/FLOAT16_IEEE data type", func(t *testing.T) {
		parseFails(t, recordLayoutScope("FNC_VALUES 1 FLOAT16_IEEE ROW_DIR DIRECT"))
	})
}

func TestGrammar_AXIS_RESCALE_X(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		recordLayout, ok := parseRecordLayout(t, "AXIS_RESCALE_X 1 UBYTE 4 INDEX_INCR DIRECT")
		if !ok {
			return
		}

		equalNode(t, &AxisRescaleXType{
			Position:                intVal("1"),
			DataType:                dataTypeVal("UBYTE"),
			MaxNumberOfRescalePairs: intVal("4"),
			IndexIncr:               indexOrderVal("INDEX_INCR"),
			Addressing:              addrTypeVal("DIRECT"),
		}, recordLayout.AXIS_RESCALE_X)
	})

	t.Run("reject/missing number of rescale pairs", func(t *testing.T) {
		parseFails(t, recordLayoutScope("AXIS_RESCALE_X 1 UBYTE INDEX_INCR DIRECT"))
	})
}

func TestGrammar_RESERVED(t *testing.T) {
	for _, dataSize := range []string{"BYTE", "WORD", "LONG"} {
		t.Run("enum/"+dataSize, func(t *testing.T) {
			recordLayout, ok := parseRecordLayout(t, "RESERVED 1 "+dataSize)
			if !ok {
				return
			}

			equalNodes(t, []*ReservedType{{Position: intVal("1"), DataSize: dataSize}}, recordLayout.RESERVED)
		})
	}

	t.Run("reject/data type instead of data size", func(t *testing.T) {
		parseFails(t, recordLayoutScope("RESERVED 1 UBYTE"))
	})
}

func TestGrammar_STATIC_RECORD_LAYOUT(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		assertPreserved(t, recordLayoutScope("STATIC_RECORD_LAYOUT"), "STATIC_RECORD_LAYOUT")
	})
}

func TestGrammar_AXIS_PTS_4(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		assertPreserved(t, recordLayoutScope("AXIS_PTS_4 1 UBYTE INDEX_INCR DIRECT"), "AXIS_PTS_4")
	})

	t.Run("reject/missing addressing", func(t *testing.T) {
		parseFails(t, recordLayoutScope("AXIS_PTS_4 1 UBYTE INDEX_INCR"))
	})
}

func TestGrammar_NO_AXIS_PTS_4(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		assertPreserved(t, recordLayoutScope("NO_AXIS_PTS_4 1 UBYTE"), "NO_AXIS_PTS_4")
	})
}
