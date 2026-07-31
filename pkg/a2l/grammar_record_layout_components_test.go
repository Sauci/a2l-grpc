package a2l

import (
	"testing"
)

func TestGrammar_AXIS_PTS_X(t *testing.T) {
	testAxisPtsKeyword(t, "AXIS_PTS_X",
		func(position *IntType, dataType *DataTypeType, indexIncr *IndexOrderType, addressing *AddrTypeType) *AxisPtsXType {
			return &AxisPtsXType{Position: position, DataType: dataType, IndexIncr: indexIncr, Addressing: addressing}
		},
		func(recordLayout *RecordLayoutType) *AxisPtsXType { return recordLayout.AXIS_PTS_X })
}

func TestGrammar_AXIS_PTS_Y(t *testing.T) {
	testAxisPtsKeyword(t, "AXIS_PTS_Y",
		func(position *IntType, dataType *DataTypeType, indexIncr *IndexOrderType, addressing *AddrTypeType) *AxisPtsYType {
			return &AxisPtsYType{Position: position, DataType: dataType, IndexIncr: indexIncr, Addressing: addressing}
		},
		func(recordLayout *RecordLayoutType) *AxisPtsYType { return recordLayout.AXIS_PTS_Y })
}

func TestGrammar_AXIS_PTS_Z(t *testing.T) {
	testAxisPtsKeyword(t, "AXIS_PTS_Z",
		func(position *IntType, dataType *DataTypeType, indexIncr *IndexOrderType, addressing *AddrTypeType) *AxisPtsZType {
			return &AxisPtsZType{Position: position, DataType: dataType, IndexIncr: indexIncr, Addressing: addressing}
		},
		func(recordLayout *RecordLayoutType) *AxisPtsZType { return recordLayout.AXIS_PTS_Z })
}

func TestGrammar_AXIS_PTS_5(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		assertPreserved(t, recordLayoutScope("AXIS_PTS_5 1 UBYTE INDEX_INCR DIRECT"), "AXIS_PTS_5")
	})
}

func TestGrammar_AXIS_RESCALE_Y(t *testing.T) {
	testAxisRescaleKeyword(t, "AXIS_RESCALE_Y",
		func(position *IntType, dataType *DataTypeType, maxNumberOfRescalePairs *IntType,
			indexIncr *IndexOrderType, addressing *AddrTypeType) *AxisRescaleYType {
			return &AxisRescaleYType{
				Position:                position,
				DataType:                dataType,
				MaxNumberOfRescalePairs: maxNumberOfRescalePairs,
				IndexIncr:               indexIncr,
				Addressing:              addressing,
			}
		},
		func(recordLayout *RecordLayoutType) *AxisRescaleYType { return recordLayout.AXIS_RESCALE_Y })
}

func TestGrammar_AXIS_RESCALE_Z(t *testing.T) {
	testAxisRescaleKeyword(t, "AXIS_RESCALE_Z",
		func(position *IntType, dataType *DataTypeType, maxNumberOfRescalePairs *IntType,
			indexIncr *IndexOrderType, addressing *AddrTypeType) *AxisRescaleZType {
			return &AxisRescaleZType{
				Position:                position,
				DataType:                dataType,
				MaxNumberOfRescalePairs: maxNumberOfRescalePairs,
				IndexIncr:               indexIncr,
				Addressing:              addressing,
			}
		},
		func(recordLayout *RecordLayoutType) *AxisRescaleZType { return recordLayout.AXIS_RESCALE_Z })
}

func TestGrammar_AXIS_RESCALE_4(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		assertPreserved(t, recordLayoutScope("AXIS_RESCALE_4 1 UBYTE 4 INDEX_INCR DIRECT"), "AXIS_RESCALE_4")
	})
}

func TestGrammar_AXIS_RESCALE_5(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		assertPreserved(t, recordLayoutScope("AXIS_RESCALE_5 1 UBYTE 4 INDEX_INCR DIRECT"), "AXIS_RESCALE_5")
	})
}

func TestGrammar_NO_AXIS_PTS_X(t *testing.T) {
	testPositionDataTypeKeyword(t, "NO_AXIS_PTS_X",
		func(position *IntType, dataType *DataTypeType) *NoAxisPtsXType {
			return &NoAxisPtsXType{Position: position, DataType: dataType}
		},
		func(recordLayout *RecordLayoutType) *NoAxisPtsXType { return recordLayout.NO_AXIS_PTS_X })
}

func TestGrammar_NO_AXIS_PTS_Y(t *testing.T) {
	testPositionDataTypeKeyword(t, "NO_AXIS_PTS_Y",
		func(position *IntType, dataType *DataTypeType) *NoAxisPtsYType {
			return &NoAxisPtsYType{Position: position, DataType: dataType}
		},
		func(recordLayout *RecordLayoutType) *NoAxisPtsYType { return recordLayout.NO_AXIS_PTS_Y })
}

func TestGrammar_NO_AXIS_PTS_Z(t *testing.T) {
	testPositionDataTypeKeyword(t, "NO_AXIS_PTS_Z",
		func(position *IntType, dataType *DataTypeType) *NoAxisPtsZType {
			return &NoAxisPtsZType{Position: position, DataType: dataType}
		},
		func(recordLayout *RecordLayoutType) *NoAxisPtsZType { return recordLayout.NO_AXIS_PTS_Z })
}

func TestGrammar_NO_AXIS_PTS_5(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		assertPreserved(t, recordLayoutScope("NO_AXIS_PTS_5 1 UBYTE"), "NO_AXIS_PTS_5")
	})
}

func TestGrammar_NO_RESCALE_X(t *testing.T) {
	testPositionDataTypeKeyword(t, "NO_RESCALE_X",
		func(position *IntType, dataType *DataTypeType) *NoRescaleXType {
			return &NoRescaleXType{Position: position, DataType: dataType}
		},
		func(recordLayout *RecordLayoutType) *NoRescaleXType { return recordLayout.NO_RESCALE_X })
}

func TestGrammar_NO_RESCALE_Y(t *testing.T) {
	testPositionDataTypeKeyword(t, "NO_RESCALE_Y",
		func(position *IntType, dataType *DataTypeType) *NoRescaleYType {
			return &NoRescaleYType{Position: position, DataType: dataType}
		},
		func(recordLayout *RecordLayoutType) *NoRescaleYType { return recordLayout.NO_RESCALE_Y })
}

func TestGrammar_NO_RESCALE_Z(t *testing.T) {
	testPositionDataTypeKeyword(t, "NO_RESCALE_Z",
		func(position *IntType, dataType *DataTypeType) *NoRescaleZType {
			return &NoRescaleZType{Position: position, DataType: dataType}
		},
		func(recordLayout *RecordLayoutType) *NoRescaleZType { return recordLayout.NO_RESCALE_Z })
}

func TestGrammar_NO_RESCALE_4(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		assertPreserved(t, recordLayoutScope("NO_RESCALE_4 1 UBYTE"), "NO_RESCALE_4")
	})
}

func TestGrammar_NO_RESCALE_5(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		assertPreserved(t, recordLayoutScope("NO_RESCALE_5 1 UBYTE"), "NO_RESCALE_5")
	})
}

func TestGrammar_FIX_NO_AXIS_PTS_X(t *testing.T) {
	testFixNoAxisPtsKeyword(t, "FIX_NO_AXIS_PTS_X",
		func(numberOfAxisPoints *IntType) *FixNoAxisPtsXType {
			return &FixNoAxisPtsXType{NumberOfAxisPoints: numberOfAxisPoints}
		},
		func(recordLayout *RecordLayoutType) *FixNoAxisPtsXType { return recordLayout.FIX_NO_AXIS_PTS_X })
}

func TestGrammar_FIX_NO_AXIS_PTS_Y(t *testing.T) {
	testFixNoAxisPtsKeyword(t, "FIX_NO_AXIS_PTS_Y",
		func(numberOfAxisPoints *IntType) *FixNoAxisPtsYType {
			return &FixNoAxisPtsYType{NumberOfAxisPoints: numberOfAxisPoints}
		},
		func(recordLayout *RecordLayoutType) *FixNoAxisPtsYType { return recordLayout.FIX_NO_AXIS_PTS_Y })
}

func TestGrammar_FIX_NO_AXIS_PTS_Z(t *testing.T) {
	testFixNoAxisPtsKeyword(t, "FIX_NO_AXIS_PTS_Z",
		func(numberOfAxisPoints *IntType) *FixNoAxisPtsZType {
			return &FixNoAxisPtsZType{NumberOfAxisPoints: numberOfAxisPoints}
		},
		func(recordLayout *RecordLayoutType) *FixNoAxisPtsZType { return recordLayout.FIX_NO_AXIS_PTS_Z })
}

func TestGrammar_FIX_NO_AXIS_PTS_4(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		assertPreserved(t, recordLayoutScope("FIX_NO_AXIS_PTS_4 17"), "FIX_NO_AXIS_PTS_4")
	})
}

func TestGrammar_FIX_NO_AXIS_PTS_5(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		assertPreserved(t, recordLayoutScope("FIX_NO_AXIS_PTS_5 17"), "FIX_NO_AXIS_PTS_5")
	})
}

func TestGrammar_SRC_ADDR_X(t *testing.T) {
	testPositionDataTypeKeyword(t, "SRC_ADDR_X",
		func(position *IntType, dataType *DataTypeType) *SrcAddrXType {
			return &SrcAddrXType{Position: position, DataType: dataType}
		},
		func(recordLayout *RecordLayoutType) *SrcAddrXType { return recordLayout.SRC_ADDR_X })
}

func TestGrammar_SRC_ADDR_Y(t *testing.T) {
	testPositionDataTypeKeyword(t, "SRC_ADDR_Y",
		func(position *IntType, dataType *DataTypeType) *SrcAddrYType {
			return &SrcAddrYType{Position: position, DataType: dataType}
		},
		func(recordLayout *RecordLayoutType) *SrcAddrYType { return recordLayout.SRC_ADDR_Y })
}

func TestGrammar_SRC_ADDR_Z(t *testing.T) {
	testPositionDataTypeKeyword(t, "SRC_ADDR_Z",
		func(position *IntType, dataType *DataTypeType) *SrcAddrZType {
			return &SrcAddrZType{Position: position, DataType: dataType}
		},
		func(recordLayout *RecordLayoutType) *SrcAddrZType { return recordLayout.SRC_ADDR_Z })
}

func TestGrammar_SRC_ADDR_4(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		assertPreserved(t, recordLayoutScope("SRC_ADDR_4 1 UBYTE"), "SRC_ADDR_4")
	})
}

func TestGrammar_SRC_ADDR_5(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		assertPreserved(t, recordLayoutScope("SRC_ADDR_5 1 UBYTE"), "SRC_ADDR_5")
	})
}

func TestGrammar_RIP_ADDR_W(t *testing.T) {
	testPositionDataTypeKeyword(t, "RIP_ADDR_W",
		func(position *IntType, dataType *DataTypeType) *RipAddrWType {
			return &RipAddrWType{Position: position, DataType: dataType}
		},
		func(recordLayout *RecordLayoutType) *RipAddrWType { return recordLayout.RIP_ADDR_W })
}

func TestGrammar_RIP_ADDR_X(t *testing.T) {
	testPositionDataTypeKeyword(t, "RIP_ADDR_X",
		func(position *IntType, dataType *DataTypeType) *RipAddrXType {
			return &RipAddrXType{Position: position, DataType: dataType}
		},
		func(recordLayout *RecordLayoutType) *RipAddrXType { return recordLayout.RIP_ADDR_X })
}

func TestGrammar_RIP_ADDR_Y(t *testing.T) {
	testPositionDataTypeKeyword(t, "RIP_ADDR_Y",
		func(position *IntType, dataType *DataTypeType) *RipAddrYType {
			return &RipAddrYType{Position: position, DataType: dataType}
		},
		func(recordLayout *RecordLayoutType) *RipAddrYType { return recordLayout.RIP_ADDR_Y })
}

func TestGrammar_RIP_ADDR_Z(t *testing.T) {
	testPositionDataTypeKeyword(t, "RIP_ADDR_Z",
		func(position *IntType, dataType *DataTypeType) *RipAddrZType {
			return &RipAddrZType{Position: position, DataType: dataType}
		},
		func(recordLayout *RecordLayoutType) *RipAddrZType { return recordLayout.RIP_ADDR_Z })
}

func TestGrammar_RIP_ADDR_4(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		assertPreserved(t, recordLayoutScope("RIP_ADDR_4 1 UBYTE"), "RIP_ADDR_4")
	})
}

func TestGrammar_RIP_ADDR_5(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		assertPreserved(t, recordLayoutScope("RIP_ADDR_5 1 UBYTE"), "RIP_ADDR_5")
	})
}

func TestGrammar_SHIFT_OP_X(t *testing.T) {
	testPositionDataTypeKeyword(t, "SHIFT_OP_X",
		func(position *IntType, dataType *DataTypeType) *ShiftOpXType {
			return &ShiftOpXType{Position: position, DataType: dataType}
		},
		func(recordLayout *RecordLayoutType) *ShiftOpXType { return recordLayout.SHIFT_OP_X })
}

func TestGrammar_SHIFT_OP_Y(t *testing.T) {
	testPositionDataTypeKeyword(t, "SHIFT_OP_Y",
		func(position *IntType, dataType *DataTypeType) *ShiftOpYType {
			return &ShiftOpYType{Position: position, DataType: dataType}
		},
		func(recordLayout *RecordLayoutType) *ShiftOpYType { return recordLayout.SHIFT_OP_Y })
}

func TestGrammar_SHIFT_OP_Z(t *testing.T) {
	testPositionDataTypeKeyword(t, "SHIFT_OP_Z",
		func(position *IntType, dataType *DataTypeType) *ShiftOpZType {
			return &ShiftOpZType{Position: position, DataType: dataType}
		},
		func(recordLayout *RecordLayoutType) *ShiftOpZType { return recordLayout.SHIFT_OP_Z })
}

func TestGrammar_SHIFT_OP_4(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		assertPreserved(t, recordLayoutScope("SHIFT_OP_4 1 UBYTE"), "SHIFT_OP_4")
	})
}

func TestGrammar_SHIFT_OP_5(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		assertPreserved(t, recordLayoutScope("SHIFT_OP_5 1 UBYTE"), "SHIFT_OP_5")
	})
}

func TestGrammar_OFFSET_X(t *testing.T) {
	testPositionDataTypeKeyword(t, "OFFSET_X",
		func(position *IntType, dataType *DataTypeType) *OffsetXType {
			return &OffsetXType{Position: position, DataType: dataType}
		},
		func(recordLayout *RecordLayoutType) *OffsetXType { return recordLayout.OFFSET_X })
}

func TestGrammar_OFFSET_Y(t *testing.T) {
	testPositionDataTypeKeyword(t, "OFFSET_Y",
		func(position *IntType, dataType *DataTypeType) *OffsetYType {
			return &OffsetYType{Position: position, DataType: dataType}
		},
		func(recordLayout *RecordLayoutType) *OffsetYType { return recordLayout.OFFSET_Y })
}

func TestGrammar_OFFSET_Z(t *testing.T) {
	testPositionDataTypeKeyword(t, "OFFSET_Z",
		func(position *IntType, dataType *DataTypeType) *OffsetZType {
			return &OffsetZType{Position: position, DataType: dataType}
		},
		func(recordLayout *RecordLayoutType) *OffsetZType { return recordLayout.OFFSET_Z })
}

func TestGrammar_OFFSET_4(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		assertPreserved(t, recordLayoutScope("OFFSET_4 1 UBYTE"), "OFFSET_4")
	})
}

func TestGrammar_OFFSET_5(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		assertPreserved(t, recordLayoutScope("OFFSET_5 1 UBYTE"), "OFFSET_5")
	})
}

func TestGrammar_DIST_OP_X(t *testing.T) {
	testPositionDataTypeKeyword(t, "DIST_OP_X",
		func(position *IntType, dataType *DataTypeType) *DistOpXType {
			return &DistOpXType{Position: position, DataType: dataType}
		},
		func(recordLayout *RecordLayoutType) *DistOpXType { return recordLayout.DIST_OP_X })
}

func TestGrammar_DIST_OP_Y(t *testing.T) {
	testPositionDataTypeKeyword(t, "DIST_OP_Y",
		func(position *IntType, dataType *DataTypeType) *DistOpYType {
			return &DistOpYType{Position: position, DataType: dataType}
		},
		func(recordLayout *RecordLayoutType) *DistOpYType { return recordLayout.DIST_OP_Y })
}

func TestGrammar_DIST_OP_Z(t *testing.T) {
	testPositionDataTypeKeyword(t, "DIST_OP_Z",
		func(position *IntType, dataType *DataTypeType) *DistOpZType {
			return &DistOpZType{Position: position, DataType: dataType}
		},
		func(recordLayout *RecordLayoutType) *DistOpZType { return recordLayout.DIST_OP_Z })
}

func TestGrammar_DIST_OP_4(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		assertPreserved(t, recordLayoutScope("DIST_OP_4 1 UBYTE"), "DIST_OP_4")
	})
}

func TestGrammar_DIST_OP_5(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		assertPreserved(t, recordLayoutScope("DIST_OP_5 1 UBYTE"), "DIST_OP_5")
	})
}

func TestGrammar_IDENTIFICATION(t *testing.T) {
	testPositionDataTypeKeyword(t, "IDENTIFICATION",
		func(position *IntType, dataType *DataTypeType) *IdentificationType {
			return &IdentificationType{Position: position, DataType: dataType}
		},
		func(recordLayout *RecordLayoutType) *IdentificationType { return recordLayout.IDENTIFICATION })
}
