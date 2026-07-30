package a2l

// MOD_COMMON and the keywords it contains: ALIGNMENT_*, BYTE_ORDER, DATA_SIZE, DEPOSIT and
// S_REC_LAYOUT.

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrammar_MOD_COMMON(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		modCommon, ok := parseModCommon(t, "")
		if !ok {
			return
		}

		equalNode(t, &ModCommonType{Comment: strVal("")}, modCommon)
	})

	t.Run("optional/all parameters", func(t *testing.T) {
		modCommon, ok := parseModCommon(t, `S_REC_LAYOUT record_layout
DEPOSIT ABSOLUTE
BYTE_ORDER MSB_LAST
DATA_SIZE 16
ALIGNMENT_BYTE 1
ALIGNMENT_WORD 2
ALIGNMENT_LONG 4
ALIGNMENT_FLOAT32_IEEE 4
ALIGNMENT_FLOAT64_IEEE 8`)
		if !ok {
			return
		}

		equalNode(t, &ModCommonType{
			Comment:                strVal(""),
			S_REC_LAYOUT:           &SRecLayoutType{Name: identVal("record_layout")},
			DEPOSIT:                &DepositType{Mode: "ABSOLUTE"},
			BYTE_ORDER:             &ByteOrderType{ByteOrder: "MSB_LAST"},
			DATA_SIZE:              &DataSizeType{Size: intVal("16")},
			ALIGNMENT_BYTE:         &AlignmentByteType{AlignmentBorder: intVal("1")},
			ALIGNMENT_WORD:         &AlignmentWordType{AlignmentBorder: intVal("2")},
			ALIGNMENT_LONG:         &AlignmentLongType{AlignmentBorder: intVal("4")},
			ALIGNMENT_FLOAT32_IEEE: &AlignmentFloat32IeeeType{AlignmentBorder: intVal("4")},
			ALIGNMENT_FLOAT64_IEEE: &AlignmentFloat64IeeeType{AlignmentBorder: intVal("8")},
		}, modCommon)
	})

	t.Run("reject/missing comment", func(t *testing.T) {
		parseFails(t, moduleScope("/begin MOD_COMMON\n/end MOD_COMMON"))
	})

	// Deviation: the optional part of every block is a repetition, so keywords which the
	// specification declares once may occur several times. Only the last one reaches the tree.
	t.Run("reject/repeated optional keyword", func(t *testing.T) {
		deviation(t, "optional keywords may be repeated", func(t assert.TestingT) {
			parseFails(t, modCommonScope("BYTE_ORDER MSB_LAST\nBYTE_ORDER MSB_FIRST"))
		})
	})
}

func TestGrammar_ALIGNMENT_BYTE(t *testing.T) {
	t.Run("mandatory parameters/in MOD_COMMON", func(t *testing.T) {
		modCommon, ok := parseModCommon(t, "ALIGNMENT_BYTE 1")
		if !ok {
			return
		}

		equalNode(t, &AlignmentByteType{AlignmentBorder: intVal("1")}, modCommon.ALIGNMENT_BYTE)
	})

	t.Run("mandatory parameters/in RECORD_LAYOUT", func(t *testing.T) {
		recordLayout, ok := parseRecordLayout(t, "ALIGNMENT_BYTE 1")
		if !ok {
			return
		}

		equalNode(t, &AlignmentByteType{AlignmentBorder: intVal("1")}, recordLayout.ALIGNMENT_BYTE)
	})

	t.Run("reject/missing alignment border", func(t *testing.T) {
		parseFails(t, modCommonScope("ALIGNMENT_BYTE"))
	})

	t.Run("reject/float alignment border", func(t *testing.T) {
		parseFails(t, modCommonScope("ALIGNMENT_BYTE 1.0"))
	})
}

func TestGrammar_ALIGNMENT_WORD(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		modCommon, ok := parseModCommon(t, "ALIGNMENT_WORD 2")
		if !ok {
			return
		}

		equalNode(t, &AlignmentWordType{AlignmentBorder: intVal("2")}, modCommon.ALIGNMENT_WORD)
	})

	t.Run("reject/missing alignment border", func(t *testing.T) {
		parseFails(t, modCommonScope("ALIGNMENT_WORD"))
	})
}

func TestGrammar_ALIGNMENT_LONG(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		modCommon, ok := parseModCommon(t, "ALIGNMENT_LONG 4")
		if !ok {
			return
		}

		equalNode(t, &AlignmentLongType{AlignmentBorder: intVal("4")}, modCommon.ALIGNMENT_LONG)
	})

	t.Run("reject/missing alignment border", func(t *testing.T) {
		parseFails(t, modCommonScope("ALIGNMENT_LONG"))
	})
}

func TestGrammar_ALIGNMENT_FLOAT32_IEEE(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		modCommon, ok := parseModCommon(t, "ALIGNMENT_FLOAT32_IEEE 4")
		if !ok {
			return
		}

		equalNode(t, &AlignmentFloat32IeeeType{AlignmentBorder: intVal("4")}, modCommon.ALIGNMENT_FLOAT32_IEEE)
	})

	t.Run("reject/missing alignment border", func(t *testing.T) {
		parseFails(t, modCommonScope("ALIGNMENT_FLOAT32_IEEE"))
	})
}

func TestGrammar_ALIGNMENT_FLOAT64_IEEE(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		modCommon, ok := parseModCommon(t, "ALIGNMENT_FLOAT64_IEEE 8")
		if !ok {
			return
		}

		equalNode(t, &AlignmentFloat64IeeeType{AlignmentBorder: intVal("8")}, modCommon.ALIGNMENT_FLOAT64_IEEE)
	})

	t.Run("reject/missing alignment border", func(t *testing.T) {
		parseFails(t, modCommonScope("ALIGNMENT_FLOAT64_IEEE"))
	})
}

// ALIGNMENT_INT64 and ALIGNMENT_FLOAT16_IEEE exist in the grammar but in neither the protobuf
// definition nor the listener, so they are parsed and then silently dropped.
func TestGrammar_ALIGNMENT_INT64(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		deviation(t, "ALIGNMENT_INT64 has no node in the tree", func(t assert.TestingT) {
			assertPreserved(t, modCommonScope("ALIGNMENT_INT64 8"), "ALIGNMENT_INT64")
		})
	})

	t.Run("reject/missing alignment border", func(t *testing.T) {
		parseFails(t, modCommonScope("ALIGNMENT_INT64"))
	})
}

func TestGrammar_ALIGNMENT_FLOAT16_IEEE(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		deviation(t, "ALIGNMENT_FLOAT16_IEEE has no node in the tree", func(t assert.TestingT) {
			assertPreserved(t, modCommonScope("ALIGNMENT_FLOAT16_IEEE 2"), "ALIGNMENT_FLOAT16_IEEE")
		})
	})

	t.Run("reject/missing alignment border", func(t *testing.T) {
		parseFails(t, modCommonScope("ALIGNMENT_FLOAT16_IEEE"))
	})
}

func TestGrammar_BYTE_ORDER(t *testing.T) {
	for _, byteOrder := range []string{"LITTLE_ENDIAN", "BIG_ENDIAN", "MSB_LAST", "MSB_FIRST"} {
		t.Run("enum/"+byteOrder, func(t *testing.T) {
			modCommon, ok := parseModCommon(t, "BYTE_ORDER "+byteOrder)
			if !ok {
				return
			}

			equalNode(t, &ByteOrderType{ByteOrder: byteOrder}, modCommon.BYTE_ORDER)
		})
	}

	t.Run("mandatory parameters/in CHARACTERISTIC", func(t *testing.T) {
		characteristic, ok := parseCharacteristic(t, "BYTE_ORDER MSB_FIRST")
		if !ok {
			return
		}

		equalNode(t, &ByteOrderType{ByteOrder: "MSB_FIRST"}, characteristic.BYTE_ORDER)
	})

	t.Run("reject/unknown byte order", func(t *testing.T) {
		parseFails(t, modCommonScope("BYTE_ORDER MSB_MIDDLE"))
	})

	t.Run("reject/missing byte order", func(t *testing.T) {
		parseFails(t, modCommonScope("BYTE_ORDER"))
	})
}

func TestGrammar_DATA_SIZE(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		modCommon, ok := parseModCommon(t, "DATA_SIZE 16")
		if !ok {
			return
		}

		equalNode(t, &DataSizeType{Size: intVal("16")}, modCommon.DATA_SIZE)
	})

	t.Run("reject/missing size", func(t *testing.T) {
		parseFails(t, modCommonScope("DATA_SIZE"))
	})
}

func TestGrammar_DEPOSIT(t *testing.T) {
	for _, mode := range []string{"ABSOLUTE", "DIFFERENCE"} {
		t.Run("enum/"+mode, func(t *testing.T) {
			modCommon, ok := parseModCommon(t, "DEPOSIT "+mode)
			if !ok {
				return
			}

			equalNode(t, &DepositType{Mode: mode}, modCommon.DEPOSIT)
		})
	}

	t.Run("mandatory parameters/in AXIS_DESCR", func(t *testing.T) {
		axisDescr, ok := parseAxisDescr(t, "DEPOSIT DIFFERENCE")
		if !ok {
			return
		}

		equalNode(t, &DepositType{Mode: "DIFFERENCE"}, axisDescr.DEPOSIT)
	})

	t.Run("reject/unknown mode", func(t *testing.T) {
		parseFails(t, modCommonScope("DEPOSIT RELATIVE"))
	})
}

func TestGrammar_S_REC_LAYOUT(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		modCommon, ok := parseModCommon(t, "S_REC_LAYOUT record_layout")
		if !ok {
			return
		}

		equalNode(t, &SRecLayoutType{Name: identVal("record_layout")}, modCommon.S_REC_LAYOUT)
	})

	t.Run("reject/string parameter", func(t *testing.T) {
		parseFails(t, modCommonScope("S_REC_LAYOUT \"record_layout\""))
	})
}
