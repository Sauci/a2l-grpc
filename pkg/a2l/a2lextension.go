package a2l

import "fmt"

//func (t *AddrTypeType) MapChildNodes(_ any) {
//	panic("leaf node")
//}

//func (t *DataTypeType) MapChildNodes(_ any) {
//	panic("leaf node")
//}

//func (t *IndexOrderType) MapChildNodes(_ any) {
//	panic("leaf node")
//}

func (t *A2MLVersionType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *A2MLVersionType) MarshalA2L() (result string) {
	return fmt.Sprintf("A2ML_VERSION %s %s", t.VersionNo.A2LString(), t.UpgradeNo.A2LString())
}

func (t *A2MLVersionType) A2LTag() (result *string) {
	return result
}

func (t *AddrEpkType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AddrEpkType) MarshalA2L() (result string) {
	return marshalA2L[*AddrEpkType](t)
}

func (t *AddrEpkType) A2LTag() (result *string) {
	return result
}

func (t *AlignmentByteType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AlignmentByteType) MarshalA2L() (result string) {
	return marshalA2L[*AlignmentByteType](t)
}

func (t *AlignmentByteType) A2LTag() (result *string) {
	return result
}

func (t *AlignmentFloat32IeeeType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AlignmentFloat32IeeeType) MarshalA2L() (result string) {
	return marshalA2L[*AlignmentFloat32IeeeType](t)
}

func (t *AlignmentFloat32IeeeType) A2LTag() (result *string) {
	return result
}

func (t *AlignmentFloat64IeeeType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AlignmentFloat64IeeeType) MarshalA2L() (result string) {
	return marshalA2L[*AlignmentFloat64IeeeType](t)
}

func (t *AlignmentFloat64IeeeType) A2LTag() (result *string) {
	return result
}

func (t *AlignmentLongType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AlignmentLongType) MarshalA2L() (result string) {
	return marshalA2L[*AlignmentLongType](t)
}

func (t *AlignmentLongType) A2LTag() (result *string) {
	return result
}

func (t *AlignmentWordType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AlignmentWordType) MarshalA2L() (result string) {
	return marshalA2L[*AlignmentWordType](t)
}

func (t *AlignmentWordType) A2LTag() (result *string) {
	return result
}

func (t *AnnotationLabelType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AnnotationLabelType) MarshalA2L() (result string) {
	return marshalA2L[*AnnotationLabelType](t)
}

func (t *AnnotationLabelType) A2LTag() (result *string) {
	return result
}

func (t *AnnotationOriginType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AnnotationOriginType) MarshalA2L() (result string) {
	return marshalA2L[*AnnotationOriginType](t)
}

func (t *AnnotationOriginType) A2LTag() (result *string) {
	return result
}

func (t *AnnotationTextType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AnnotationTextType) MarshalA2L() (result string) {
	return marshalA2L[*AnnotationTextType](t)
}

func (t *AnnotationTextType) A2LTag() (result *string) {
	return result
}

func (t *AnnotationType) MapChildNodes(node any) {
	switch node.(type) {
	case *AnnotationLabelType:
		t.ANNOTATION_LABEL = node.(*AnnotationLabelType)
	case *AnnotationOriginType:
		t.ANNOTATION_ORIGIN = node.(*AnnotationOriginType)
	case *AnnotationTextType:
		t.ANNOTATION_TEXT = node.(*AnnotationTextType)
	default:
		panic("not implemented yet...")
	}
}

func (t *AnnotationType) MarshalA2L() (result string) {
	return marshalA2L[*AnnotationType](t)
}

func (t *AnnotationType) A2LTag() (result *string) {
	return result
}

func (t *ArraySizeType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *ArraySizeType) MarshalA2L() (result string) {
	return marshalA2L[*ArraySizeType](t)
}

func (t *ArraySizeType) A2LTag() (result *string) {
	return result
}

func (t *ASAP2VersionType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *ASAP2VersionType) MarshalA2L() (result string) {
	return fmt.Sprintf("ASAP2_VERSION %s %s", t.VersionNo.A2LString(), t.UpgradeNo.A2LString())
}

func (t *ASAP2VersionType) A2LTag() (result *string) {
	return result
}

func (t *AxisDescrType) MapChildNodes(node any) {
	switch node.(type) {
	case *ReadOnlyType:
		t.READ_ONLY = node.(*ReadOnlyType)
	case *FormatType:
		t.FORMAT = node.(*FormatType)
	case *AnnotationType:
		if t.ANNOTATION == nil {
			t.ANNOTATION = make([]*AnnotationType, 0)
		}

		t.ANNOTATION = append(t.ANNOTATION, node.(*AnnotationType))
	case *AxisPtsRefType:
		t.AXIS_PTS_REF = node.(*AxisPtsRefType)
	case *MaxGradType:
		t.MAX_GRAD = node.(*MaxGradType)
	case *MonotonyType:
		t.MONOTONY = node.(*MonotonyType)
	case *ByteOrderType:
		t.BYTE_ORDER = node.(*ByteOrderType)
	case *ExtendedLimitsType:
		t.EXTENDED_LIMITS = node.(*ExtendedLimitsType)
	case *FixAxisParType:
		t.FIX_AXIS_PAR = node.(*FixAxisParType)
	case *FixAxisParDistType:
		t.FIX_AXIS_PAR_DIST = node.(*FixAxisParDistType)
	case *FixAxisParListType:
		t.FIX_AXIS_PAR_LIST = node.(*FixAxisParListType)
	case *DepositType:
		t.DEPOSIT = node.(*DepositType)
	case *CurveAxisRefType:
		t.CURVE_AXIS_REF = node.(*CurveAxisRefType)
	default:
		panic("not implemented yet...")
	}
}

func (t *AxisDescrType) MarshalA2L() (result string) {
	return marshalA2L[*AxisDescrType](t)
}

func (t *AxisDescrType) A2LTag() (result *string) {
	return result
}

func (t *AxisPtsRefType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AxisPtsRefType) MarshalA2L() (result string) {
	return marshalA2L[*AxisPtsRefType](t)
}

func (t *AxisPtsRefType) A2LTag() (result *string) {
	return result
}

func (t *AxisPtsType) MapChildNodes(node any) {
	switch node.(type) {
	case *DisplayIdentifierType:
		t.DISPLAY_IDENTIFIER = node.(*DisplayIdentifierType)
	case *ReadOnlyType:
		t.READ_ONLY = node.(*ReadOnlyType)
	case *FormatType:
		t.FORMAT = node.(*FormatType)
	case *DepositType:
		t.DEPOSIT = node.(*DepositType)
	case *ByteOrderType:
		t.BYTE_ORDER = node.(*ByteOrderType)
	case *FunctionListType:
		t.FUNCTION_LIST = node.(*FunctionListType)
	case *RefMemorySegmentType:
		t.REF_MEMORY_SEGMENT = node.(*RefMemorySegmentType)
	case *GuardRailsType:
		t.GUARD_RAILS = node.(*GuardRailsType)
	case *ExtendedLimitsType:
		t.EXTENDED_LIMITS = node.(*ExtendedLimitsType)
	case *AnnotationType:
		if t.ANNOTATION == nil {
			t.ANNOTATION = make([]*AnnotationType, 0)
		}

		t.ANNOTATION = append(t.ANNOTATION, node.(*AnnotationType))
	case *IfDataType:
		if t.IF_DATA == nil {
			t.IF_DATA = make([]*IfDataType, 0)
		}

		t.IF_DATA = append(t.IF_DATA, node.(*IfDataType))
	case *CalibrationAccessType:
		t.CALIBRATION_ACCESS = node.(*CalibrationAccessType)
	case *EcuAddressExtensionType:
		t.ECU_ADDRESS_EXTENSION = node.(*EcuAddressExtensionType)
	default:
		panic("not implemented yet...")
	}
}

func (t *AxisPtsType) MarshalA2L() (result string) {
	return marshalA2L[*AxisPtsType](t)
}

func (t *AxisPtsType) A2LTag() (result *string) {
	return result
}

func (t *AxisPtsXType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AxisPtsXType) MarshalA2L() (result string) {
	return marshalA2L[*AxisPtsXType](t)
}

func (t *AxisPtsXType) A2LTag() (result *string) {
	return result
}

func (t *AxisPtsYType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AxisPtsYType) MarshalA2L() (result string) {
	return marshalA2L[*AxisPtsYType](t)
}

func (t *AxisPtsYType) A2LTag() (result *string) {
	return result
}

func (t *AxisPtsZType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AxisPtsZType) MarshalA2L() (result string) {
	return marshalA2L[*AxisPtsZType](t)
}

func (t *AxisPtsZType) A2LTag() (result *string) {
	return result
}

func (t *AxisRescaleXType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AxisRescaleXType) MarshalA2L() (result string) {
	return marshalA2L[*AxisRescaleXType](t)
}

func (t *AxisRescaleXType) A2LTag() (result *string) {
	return result
}

func (t *AxisRescaleYType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AxisRescaleYType) MarshalA2L() (result string) {
	return marshalA2L[*AxisRescaleYType](t)
}

func (t *AxisRescaleYType) A2LTag() (result *string) {
	return result
}

func (t *AxisRescaleZType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AxisRescaleZType) MarshalA2L() (result string) {
	return marshalA2L[*AxisRescaleZType](t)
}

func (t *AxisRescaleZType) A2LTag() (result *string) {
	return result
}

func (t *BitMaskType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *BitMaskType) MarshalA2L() (result string) {
	return marshalA2L[*BitMaskType](t)
}

func (t *BitMaskType) A2LTag() (result *string) {
	return result
}

func (t *BitOperationType) MapChildNodes(node any) {
	switch node.(type) {
	case *LeftShiftType:
		t.LEFT_SHIFT = node.(*LeftShiftType)
	case *RightShiftType:
		t.RIGHT_SHIFT = node.(*RightShiftType)
	case *SignExtendType:
		t.SIGN_EXTEND = node.(*SignExtendType)
	default:
		panic("not implemented yet...")
	}
}

func (t *BitOperationType) MarshalA2L() (result string) {
	return marshalA2L[*BitOperationType](t)
}

func (t *BitOperationType) A2LTag() (result *string) {
	return result
}

func (t *ByteOrderType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *ByteOrderType) MarshalA2L() (result string) {
	return marshalA2L[*ByteOrderType](t)
}

func (t *ByteOrderType) A2LTag() (result *string) {
	return result
}

func (t *CalibrationAccessType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *CalibrationAccessType) MarshalA2L() (result string) {
	return marshalA2L[*CalibrationAccessType](t)
}

func (t *CalibrationAccessType) A2LTag() (result *string) {
	return result
}

func (t *CalibrationHandleType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *CalibrationHandleType) MarshalA2L() (result string) {
	return marshalA2L[*CalibrationHandleType](t)
}

func (t *CalibrationHandleType) A2LTag() (result *string) {
	return result
}

func (t *CalibrationMethodType) MapChildNodes(node any) {
	switch node.(type) {
	case *CalibrationHandleType:
		if t.CALIBRATION_HANDLE == nil {
			t.CALIBRATION_HANDLE = make([]*CalibrationHandleType, 0)
		}

		t.CALIBRATION_HANDLE = append(t.CALIBRATION_HANDLE, node.(*CalibrationHandleType))
	default:
		panic("not implemented yet...")
	}
}

func (t *CalibrationMethodType) MarshalA2L() (result string) {
	return marshalA2L[*CalibrationMethodType](t)
}

func (t *CalibrationMethodType) A2LTag() (result *string) {
	return result
}

func (t *CharacteristicType) MapChildNodes(node any) {
	switch node.(type) {
	case *DisplayIdentifierType:
		t.DISPLAY_IDENTIFIER = node.(*DisplayIdentifierType)
	case *FormatType:
		t.FORMAT = node.(*FormatType)
	case *ByteOrderType:
		t.BYTE_ORDER = node.(*ByteOrderType)
	case *BitMaskType:
		t.BIT_MASK = node.(*BitMaskType)
	case *FunctionListType:
		t.FUNCTION_LIST = node.(*FunctionListType)
	case *NumberType:
		t.NUMBER = node.(*NumberType)
	case *ExtendedLimitsType:
		t.EXTENDED_LIMITS = node.(*ExtendedLimitsType)
	case *ReadOnlyType:
		t.READ_ONLY = node.(*ReadOnlyType)
	case *GuardRailsType:
		t.GUARD_RAILS = node.(*GuardRailsType)
	case *MapListType:
		t.MAP_LIST = node.(*MapListType)
	case *MaxRefreshType:
		t.MAX_REFRESH = node.(*MaxRefreshType)
	case *DependentCharacteristicType:
		t.DEPENDENT_CHARACTERISTIC = node.(*DependentCharacteristicType)
	case *VirtualCharacteristicType:
		t.VIRTUAL_CHARACTERISTIC = node.(*VirtualCharacteristicType)
	case *RefMemorySegmentType:
		t.REF_MEMORY_SEGMENT = node.(*RefMemorySegmentType)
	case *AnnotationType:
		if t.ANNOTATION == nil {
			t.ANNOTATION = make([]*AnnotationType, 0)
		}

		t.ANNOTATION = append(t.ANNOTATION, node.(*AnnotationType))
	case *ComparisonQuantityType:
		t.COMPARISON_QUANTITY = node.(*ComparisonQuantityType)
	case *IfDataType:
		if t.IF_DATA == nil {
			t.IF_DATA = make([]*IfDataType, 0)
		}

		t.IF_DATA = append(t.IF_DATA, node.(*IfDataType))
	case *AxisDescrType:
		if t.AXIS_DESCR == nil {
			t.AXIS_DESCR = make([]*AxisDescrType, 0)
		}

		t.AXIS_DESCR = append(t.AXIS_DESCR, node.(*AxisDescrType))
	case *CalibrationAccessType:
		t.CALIBRATION_ACCESS = node.(*CalibrationAccessType)
	case *MatrixDimType:
		t.MATRIX_DIM = node.(*MatrixDimType)
	case *EcuAddressExtensionType:
		t.ECU_ADDRESS_EXTENSION = node.(*EcuAddressExtensionType)
	default:
		panic("not implemented yet...")
	}
}

func (t *CharacteristicType) MarshalA2L() (result string) {
	return marshalA2L[*CharacteristicType](t)
}

func (t *CharacteristicType) A2LTag() (result *string) {
	return result
}

func (t *CoeffsType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *CoeffsType) MarshalA2L() (result string) {
	return marshalA2L[*CoeffsType](t)
}

func (t *CoeffsType) A2LTag() (result *string) {
	return result
}

func (t *ComparisonQuantityType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *ComparisonQuantityType) MarshalA2L() (result string) {
	return marshalA2L[*ComparisonQuantityType](t)
}

func (t *ComparisonQuantityType) A2LTag() (result *string) {
	return result
}

func (t *CompuMethodType) MapChildNodes(node any) {
	switch node.(type) {
	case *FormulaType:
		t.FORMULA = node.(*FormulaType)
	case *CoeffsType:
		t.COEFFS = node.(*CoeffsType)
	case *CompuTabRefType:
		t.COMPU_TAB_REF = node.(*CompuTabRefType)
	case *RefUnitType:
		t.REF_UNIT = node.(*RefUnitType)
	default:
		panic("not implemented yet...")
	}
}

func (t *CompuMethodType) MarshalA2L() (result string) {
	return marshalA2L[*CompuMethodType](t)
}

func (t *CompuMethodType) A2LTag() (result *string) {
	return result
}

func (t *CompuTabRefType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *CompuTabRefType) MarshalA2L() (result string) {
	return marshalA2L[*CompuTabRefType](t)
}

func (t *CompuTabRefType) A2LTag() (result *string) {
	return result
}

func (t *CompuTabType) MapChildNodes(node any) {
	switch node.(type) {
	case *DefaultValueType:
		t.DEFAULT_VALUE = node.(*DefaultValueType)
	default:
		panic("not implemented yet...")
	}
}

func (t *CompuTabType) MarshalA2L() (result string) {
	return marshalA2L[*CompuTabType](t)
}

func (t *CompuTabType) A2LTag() (result *string) {
	return result
}

func (t *CompuVTabRangeType) MapChildNodes(node any) {
	switch node.(type) {
	case *DefaultValueType:
		t.DEFAULT_VALUE = node.(*DefaultValueType)
	default:
		panic("not implemented yet...")
	}
}

func (t *CompuVTabRangeType) MarshalA2L() (result string) {
	return marshalA2L[*CompuVTabRangeType](t)
}

func (t *CompuVTabRangeType) A2LTag() (result *string) {
	return result
}

func (t *CompuVTabType) MapChildNodes(node any) {
	switch node.(type) {
	case *DefaultValueType:
		t.DEFAULT_VALUE = node.(*DefaultValueType)
	default:
		panic("not implemented yet...")
	}
}

func (t *CompuVTabType) MarshalA2L() (result string) {
	return marshalA2L[*CompuVTabType](t)
}

func (t *CompuVTabType) A2LTag() (result *string) {
	return result
}

func (t *CpuTypeType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *CpuTypeType) MarshalA2L() (result string) {
	return marshalA2L[*CpuTypeType](t)
}

func (t *CpuTypeType) A2LTag() (result *string) {
	return result
}

func (t *CurveAxisRefType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *CurveAxisRefType) MarshalA2L() (result string) {
	return marshalA2L[*CurveAxisRefType](t)
}

func (t *CurveAxisRefType) A2LTag() (result *string) {
	return result
}

func (t *CustomerNoType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *CustomerNoType) MarshalA2L() (result string) {
	return marshalA2L[*CustomerNoType](t)
}

func (t *CustomerNoType) A2LTag() (result *string) {
	return result
}

func (t *CustomerType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *CustomerType) MarshalA2L() (result string) {
	return marshalA2L[*CustomerType](t)
}

func (t *CustomerType) A2LTag() (result *string) {
	return result
}

func (t *DataSizeType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *DataSizeType) MarshalA2L() (result string) {
	return marshalA2L[*DataSizeType](t)
}

func (t *DataSizeType) A2LTag() (result *string) {
	return result
}

func (t *DefaultValueType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *DefaultValueType) MarshalA2L() (result string) {
	return marshalA2L[*DefaultValueType](t)
}

func (t *DefaultValueType) A2LTag() (result *string) {
	return result
}

func (t *DefCharacteristicType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *DefCharacteristicType) MarshalA2L() (result string) {
	return marshalA2L[*DefCharacteristicType](t)
}

func (t *DefCharacteristicType) A2LTag() (result *string) {
	return result
}

func (t *DependentCharacteristicType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *DependentCharacteristicType) MarshalA2L() (result string) {
	return marshalA2L[*DependentCharacteristicType](t)
}

func (t *DependentCharacteristicType) A2LTag() (result *string) {
	return result
}

func (t *DepositType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *DepositType) MarshalA2L() (result string) {
	return marshalA2L[*DepositType](t)
}

func (t *DepositType) A2LTag() (result *string) {
	return result
}

func (t *DisplayIdentifierType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *DisplayIdentifierType) MarshalA2L() (result string) {
	return marshalA2L[*DisplayIdentifierType](t)
}

func (t *DisplayIdentifierType) A2LTag() (result *string) {
	return result
}

func (t *DistOpXType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *DistOpXType) MarshalA2L() (result string) {
	return marshalA2L[*DistOpXType](t)
}

func (t *DistOpXType) A2LTag() (result *string) {
	return result
}

func (t *DistOpYType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *DistOpYType) MarshalA2L() (result string) {
	return marshalA2L[*DistOpYType](t)
}

func (t *DistOpYType) A2LTag() (result *string) {
	return result
}

func (t *DistOpZType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *DistOpZType) MarshalA2L() (result string) {
	return marshalA2L[*DistOpZType](t)
}

func (t *DistOpZType) A2LTag() (result *string) {
	return result
}

func (t *EcuAddressExtensionType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *EcuAddressExtensionType) MarshalA2L() (result string) {
	return marshalA2L[*EcuAddressExtensionType](t)
}

func (t *EcuAddressExtensionType) A2LTag() (result *string) {
	return result
}

func (t *EcuAddressType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *EcuAddressType) MarshalA2L() (result string) {
	return marshalA2L[*EcuAddressType](t)
}

func (t *EcuAddressType) A2LTag() (result *string) {
	return result
}

func (t *EcuCalibrationOffsetType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *EcuCalibrationOffsetType) MarshalA2L() (result string) {
	return marshalA2L[*EcuCalibrationOffsetType](t)
}

func (t *EcuCalibrationOffsetType) A2LTag() (result *string) {
	return result
}

func (t *EcuType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *EcuType) MarshalA2L() (result string) {
	return marshalA2L[*EcuType](t)
}

func (t *EcuType) A2LTag() (result *string) {
	return result
}

func (t *EpkType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *EpkType) MarshalA2L() (result string) {
	return marshalA2L[*EpkType](t)
}

func (t *EpkType) A2LTag() (result *string) {
	return result
}

func (t *ErrorMaskType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *ErrorMaskType) MarshalA2L() (result string) {
	return marshalA2L[*ErrorMaskType](t)
}

func (t *ErrorMaskType) A2LTag() (result *string) {
	return result
}

func (t *ExtendedLimitsType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *ExtendedLimitsType) MarshalA2L() (result string) {
	return marshalA2L[*ExtendedLimitsType](t)
}

func (t *ExtendedLimitsType) A2LTag() (result *string) {
	return result
}

func (t *FixAxisParDistType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *FixAxisParDistType) MarshalA2L() (result string) {
	return marshalA2L[*FixAxisParDistType](t)
}

func (t *FixAxisParDistType) A2LTag() (result *string) {
	return result
}

func (t *FixAxisParListType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *FixAxisParListType) MarshalA2L() (result string) {
	return marshalA2L[*FixAxisParListType](t)
}

func (t *FixAxisParListType) A2LTag() (result *string) {
	return result
}

func (t *FixAxisParType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *FixAxisParType) MarshalA2L() (result string) {
	return marshalA2L[*FixAxisParType](t)
}

func (t *FixAxisParType) A2LTag() (result *string) {
	return result
}

func (t *FixNoAxisPtsXType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *FixNoAxisPtsXType) MarshalA2L() (result string) {
	return marshalA2L[*FixNoAxisPtsXType](t)
}

func (t *FixNoAxisPtsXType) A2LTag() (result *string) {
	return result
}

func (t *FixNoAxisPtsYType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *FixNoAxisPtsYType) MarshalA2L() (result string) {
	return marshalA2L[*FixNoAxisPtsYType](t)
}

func (t *FixNoAxisPtsYType) A2LTag() (result *string) {
	return result
}

func (t *FixNoAxisPtsZType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *FixNoAxisPtsZType) MarshalA2L() (result string) {
	return marshalA2L[*FixNoAxisPtsZType](t)
}

func (t *FixNoAxisPtsZType) A2LTag() (result *string) {
	return result
}

func (t *FncValuesType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *FncValuesType) MarshalA2L() (result string) {
	return marshalA2L[*FncValuesType](t)
}

func (t *FncValuesType) A2LTag() (result *string) {
	return result
}

func (t *FormatType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *FormatType) MarshalA2L() (result string) {
	return marshalA2L[*FormatType](t)
}

func (t *FormatType) A2LTag() (result *string) {
	return result
}

func (t *FormulaInvType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *FormulaInvType) MarshalA2L() (result string) {
	return marshalA2L[*FormulaInvType](t)
}

func (t *FormulaInvType) A2LTag() (result *string) {
	return result
}

func (t *FormulaType) MapChildNodes(node any) {
	switch node.(type) {
	case *FormulaInvType:
		t.FORMULA_INV = node.(*FormulaInvType)
	default:
		panic("not implemented yet...")
	}
}

func (t *FormulaType) MarshalA2L() (result string) {
	return marshalA2L[*FormulaType](t)
}

func (t *FormulaType) A2LTag() (result *string) {
	return result
}

func (t *FrameMeasurementType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *FrameMeasurementType) MarshalA2L() (result string) {
	return marshalA2L[*FrameMeasurementType](t)
}

func (t *FrameMeasurementType) A2LTag() (result *string) {
	return result
}

func (t *FrameType) MapChildNodes(node any) {
	switch node.(type) {
	case *FrameMeasurementType:
		t.FRAME_MEASUREMENT = node.(*FrameMeasurementType)
	case *IfDataType:
		if t.IF_DATA == nil {
			t.IF_DATA = make([]*IfDataType, 0)
		}

		t.IF_DATA = append(t.IF_DATA, node.(*IfDataType))
	default:
		panic("not implemented yet...")
	}
}

func (t *FrameType) MarshalA2L() (result string) {
	return marshalA2L[*FrameType](t)
}

func (t *FrameType) A2LTag() (result *string) {
	return result
}

func (t *FunctionListType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *FunctionListType) MarshalA2L() (result string) {
	return marshalA2L[*FunctionListType](t)
}

func (t *FunctionListType) A2LTag() (result *string) {
	return result
}

func (t *FunctionType) MapChildNodes(node any) {
	switch node.(type) {
	case *AnnotationType:
		if t.ANNOTATION == nil {
			t.ANNOTATION = make([]*AnnotationType, 0)
		}

		t.ANNOTATION = append(t.ANNOTATION, node.(*AnnotationType))
	case *DefCharacteristicType:
		t.DEF_CHARACTERISTIC = node.(*DefCharacteristicType)
	case *RefCharacteristicType:
		t.REF_CHARACTERISTIC = node.(*RefCharacteristicType)
	case *InMeasurementType:
		t.IN_MEASUREMENT = node.(*InMeasurementType)
	case *OutMeasurementType:
		t.OUT_MEASUREMENT = node.(*OutMeasurementType)
	case *LocMeasurementType:
		t.LOC_MEASUREMENT = node.(*LocMeasurementType)
	case *SubFunctionType:
		t.SUB_FUNCTION = node.(*SubFunctionType)
	case *FunctionVersionType:
		t.FUNCTION_VERSION = node.(*FunctionVersionType)
	default:
		panic("not implemented yet...")
	}
}

func (t *FunctionType) MarshalA2L() (result string) {
	return marshalA2L[*FunctionType](t)
}

func (t *FunctionType) A2LTag() (result *string) {
	return result
}

func (t *FunctionVersionType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *FunctionVersionType) MarshalA2L() (result string) {
	return marshalA2L[*FunctionVersionType](t)
}

func (t *FunctionVersionType) A2LTag() (result *string) {
	return result
}

func (t *GroupType) MapChildNodes(node any) {
	switch node.(type) {
	case *AnnotationType:
		if t.ANNOTATION == nil {
			t.ANNOTATION = make([]*AnnotationType, 0)
		}

		t.ANNOTATION = append(t.ANNOTATION, node.(*AnnotationType))
	case RootType:
		t.ROOT = node.(*RootType)
	case *RefCharacteristicType:
		t.REF_CHARACTERISTIC = node.(*RefCharacteristicType)
	case *RefMeasurementType:
		t.REF_MEASUREMENT = node.(*RefMeasurementType)
	case *FunctionListType:
		t.FUNCTION_LIST = node.(*FunctionListType)
	case *SubGroupType:
		t.SUB_GROUP = node.(*SubGroupType)
	default:
		panic("not implemented yet...")
	}
}

func (t *GroupType) MarshalA2L() (result string) {
	return marshalA2L[*GroupType](t)
}

func (t *GroupType) A2LTag() (result *string) {
	return result
}

func (t *GuardRailsType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *GuardRailsType) MarshalA2L() (result string) {
	return marshalA2L[*GuardRailsType](t)
}

func (t *GuardRailsType) A2LTag() (result *string) {
	return result
}

func (t *HeaderType) MapChildNodes(node any) {
	switch node.(type) {
	case *VersionType:
		t.VERSION = node.(*VersionType)
	case *ProjectNoType:
		t.PROJECT_NO = node.(*ProjectNoType)
	default:
		panic("not implemented yet...")
	}
}

func (t *HeaderType) MarshalA2L() (result string) {
	return fmt.Sprintf(`/begin HEADER
	%s
	%s
	%s
/end HEADER`, t.Comment.A2LString(), t.VERSION.MarshalA2L(), t.PROJECT_NO.MarshalA2L())
}

func (t *HeaderType) A2LTag() *string {
	tag := "HEADER"
	return &tag
}

func (t *IdentificationType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *IdentificationType) MarshalA2L() (result string) {
	return marshalA2L[*IdentificationType](t)
}

func (t *IdentificationType) A2LTag() (result *string) {
	return result
}

func (t *InMeasurementType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *InMeasurementType) MarshalA2L() (result string) {
	return marshalA2L[*InMeasurementType](t)
}

func (t *InMeasurementType) A2LTag() (result *string) {
	return result
}

func (t *LeftShiftType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *LeftShiftType) MarshalA2L() (result string) {
	return marshalA2L[*LeftShiftType](t)
}

func (t *LeftShiftType) A2LTag() (result *string) {
	return result
}

func (t *LocMeasurementType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *LocMeasurementType) MarshalA2L() (result string) {
	return marshalA2L[*LocMeasurementType](t)
}

func (t *LocMeasurementType) A2LTag() (result *string) {
	return result
}

func (t *MapListType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *MapListType) MarshalA2L() (result string) {
	return marshalA2L[*MapListType](t)
}

func (t *MapListType) A2LTag() (result *string) {
	return result
}

func (t *MatrixDimType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *MatrixDimType) MarshalA2L() (result string) {
	return marshalA2L[*MatrixDimType](t)
}

func (t *MatrixDimType) A2LTag() (result *string) {
	return result
}

func (t *MaxGradType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *MaxGradType) MarshalA2L() (result string) {
	return marshalA2L[*MaxGradType](t)
}

func (t *MaxGradType) A2LTag() (result *string) {
	return result
}

func (t *MaxRefreshType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *MaxRefreshType) MarshalA2L() (result string) {
	return marshalA2L[*MaxRefreshType](t)
}

func (t *MaxRefreshType) A2LTag() (result *string) {
	return result
}

func (t *MeasurementType) MapChildNodes(node any) {
	switch node.(type) {
	case *DataTypeType:
		// TODO: we might be able to remove dynamic assignation for this type....
	case *DisplayIdentifierType:
		t.DISPLAY_IDENTIFIER = node.(*DisplayIdentifierType)
	case *ReadWriteType:
		t.READ_WRITE = node.(*ReadWriteType)
	case *FormatType:
		t.FORMAT = node.(*FormatType)
	case *ArraySizeType:
		t.ARRAY_SIZE = node.(*ArraySizeType)
	case *BitMaskType:
		t.BIT_MASK = node.(*BitMaskType)
	case *BitOperationType:
		t.BIT_OPERATION = node.(*BitOperationType)
	case *ByteOrderType:
		t.BYTE_ORDER = node.(*ByteOrderType)
	case *MaxRefreshType:
		t.MAX_REFRESH = node.(*MaxRefreshType)
	case *VirtualType:
		t.VIRTUAL = node.(*VirtualType)
	case *FunctionListType:
		t.FUNCTION_LIST = node.(*FunctionListType)
	case *EcuAddressType:
		t.ECU_ADDRESS = node.(*EcuAddressType)
	case *ErrorMaskType:
		t.ERROR_MASK = node.(*ErrorMaskType)
	case *RefMemorySegmentType:
		t.REF_MEMORY_SEGMENT = node.(*RefMemorySegmentType)
	case *AnnotationType:
		if t.ANNOTATION == nil {
			t.ANNOTATION = make([]*AnnotationType, 0)
		}

		t.ANNOTATION = append(t.ANNOTATION, node.(*AnnotationType))
	case *IfDataType:
		if t.IF_DATA == nil {
			t.IF_DATA = make([]*IfDataType, 0)
		}

		t.IF_DATA = append(t.IF_DATA, node.(*IfDataType))
	case *MatrixDimType:
		t.MATRIX_DIM = node.(*MatrixDimType)
	case *EcuAddressExtensionType:
		t.ECU_ADDRESS_EXTENSION = node.(*EcuAddressExtensionType)
	default:
		panic("not implemented yet...")
	}
}

func (t *MeasurementType) MarshalA2L() (result string) {
	return marshalA2L[*MeasurementType](t)
}

func (t *MeasurementType) A2LTag() (result *string) {
	return result
}

func (t *MemoryLayoutType) MapChildNodes(node any) {
	switch node.(type) {
	case *IfDataType:
		if t.IF_DATA == nil {
			t.IF_DATA = make([]*IfDataType, 0)
		}

		t.IF_DATA = append(t.IF_DATA, node.(*IfDataType))
	default:
		panic("not implemented yet...")
	}
}

func (t *MemoryLayoutType) MarshalA2L() (result string) {
	return marshalA2L[*MemoryLayoutType](t)
}

func (t *MemoryLayoutType) A2LTag() (result *string) {
	return result
}

func (t *MemorySegmentType) MapChildNodes(node any) {
	switch node.(type) {
	case *IfDataType:
		if t.IF_DATA == nil {
			t.IF_DATA = make([]*IfDataType, 0)
		}

		t.IF_DATA = append(t.IF_DATA, node.(*IfDataType))
	default:
		panic("not implemented yet...")
	}
}

func (t *MemorySegmentType) MarshalA2L() (result string) {
	return marshalA2L[*MemorySegmentType](t)
}

func (t *MemorySegmentType) A2LTag() *string {
	tag := "MEMORY_SEGMENT"
	return &tag
}

func (t *ModCommonType) MapChildNodes(node any) {
	switch node.(type) {
	case *SRecLayoutType:
		t.S_REC_LAYOUT = node.(*SRecLayoutType)
	case *DepositType:
		t.DEPOSIT = node.(*DepositType)
	case *ByteOrderType:
		t.BYTE_ORDER = node.(*ByteOrderType)
	case *DataSizeType:
		t.DATA_SIZE = node.(*DataSizeType)
	case *AlignmentByteType:
		t.ALIGNMENT_BYTE = node.(*AlignmentByteType)
	case *AlignmentWordType:
		t.ALIGNMENT_WORD = node.(*AlignmentWordType)
	case *AlignmentLongType:
		t.ALIGNMENT_LONG = node.(*AlignmentLongType)
	case *AlignmentFloat32IeeeType:
		t.ALIGNMENT_FLOAT32_IEEE = node.(*AlignmentFloat32IeeeType)
	case *AlignmentFloat64IeeeType:
		t.ALIGNMENT_FLOAT64_IEEE = node.(*AlignmentFloat64IeeeType)
	default:
		panic("not implemented yet...")
	}
}

func (t *ModCommonType) MarshalA2L() (result string) {
	return marshalA2L[*ModCommonType](t)
}

func (t *ModCommonType) A2LTag() *string {
	tag := "MOD_COMMON"
	return &tag
}

func (t *ModParType) MapChildNodes(node any) {
	switch node.(type) {
	case *VersionType:
		t.VERSION = node.(*VersionType)
	case *AddrEpkType:
		if t.ADDR_EPK == nil {
			t.ADDR_EPK = make([]*AddrEpkType, 0)
		}

		t.ADDR_EPK = append(t.ADDR_EPK, node.(*AddrEpkType))
	case *EpkType:
		t.EPK = node.(*EpkType)
	case *SupplierType:
		t.SUPPLIER = node.(*SupplierType)
	case *CustomerType:
		t.CUSTOMER = node.(*CustomerType)
	case *CustomerNoType:
		t.CUSTOMER_NO = node.(*CustomerNoType)
	case *UserType:
		t.USER = node.(*UserType)
	case *PhoneNoType:
		t.PHONE_NO = node.(*PhoneNoType)
	case *EcuType:
		t.ECU = node.(*EcuType)
	case *CpuTypeType:
		t.CPU_TYPE = node.(*CpuTypeType)
	case *NoOfInterfacesType:
		t.NO_OF_INTERFACES = node.(*NoOfInterfacesType)
	case *EcuCalibrationOffsetType:
		t.ECU_CALIBRATION_OFFSET = node.(*EcuCalibrationOffsetType)
	case *CalibrationMethodType:
		if t.CALIBRATION_METHOD == nil {
			t.CALIBRATION_METHOD = make([]*CalibrationMethodType, 0)
		}

		t.CALIBRATION_METHOD = append(t.CALIBRATION_METHOD, node.(*CalibrationMethodType))
	case *MemoryLayoutType:
		if t.MEMORY_LAYOUT == nil {
			t.MEMORY_LAYOUT = make([]*MemoryLayoutType, 0)
		}

		t.MEMORY_LAYOUT = append(t.MEMORY_LAYOUT, node.(*MemoryLayoutType))
	case *MemorySegmentType:
		if t.MEMORY_SEGMENT == nil {
			t.MEMORY_SEGMENT = make([]*MemorySegmentType, 0)
		}

		t.MEMORY_SEGMENT = append(t.MEMORY_SEGMENT, node.(*MemorySegmentType))
	case *SystemConstantType:
		if t.SYSTEM_CONSTANT == nil {
			t.SYSTEM_CONSTANT = make([]*SystemConstantType, 0)
		}

		t.SYSTEM_CONSTANT = append(t.SYSTEM_CONSTANT, node.(*SystemConstantType))
	default:
		panic("not implemented yet...")
	}
}

func (t *ModParType) MarshalA2L() (result string) {
	return marshalA2L[*ModParType](t)
}

func (t *ModParType) A2LTag() *string {
	tag := "MOD_PAR"
	return &tag
}

func (t *ModuleType) MapChildNodes(node any) {
	switch node.(type) {
	case *A2MLType:
		t.A2ML = node.(*A2MLType)
	case *ModParType:
		t.MOD_PAR = node.(*ModParType)
	case *ModCommonType:
		t.MOD_COMMON = node.(*ModCommonType)
	case *IfDataType:
		if t.IF_DATA == nil {
			t.IF_DATA = make([]*IfDataType, 0)
		}

		t.IF_DATA = append(t.IF_DATA, node.(*IfDataType))
	case *CharacteristicType:
		if t.CHARACTERISTIC == nil {
			t.CHARACTERISTIC = make([]*CharacteristicType, 0)
		}

		t.CHARACTERISTIC = append(t.CHARACTERISTIC, node.(*CharacteristicType))
	case *AxisPtsType:
		if t.AXIS_PTS == nil {
			t.AXIS_PTS = make([]*AxisPtsType, 0)
		}

		t.AXIS_PTS = append(t.AXIS_PTS, node.(*AxisPtsType))
	case *MeasurementType:
		if t.MEASUREMENT == nil {
			t.MEASUREMENT = make([]*MeasurementType, 0)
		}

		t.MEASUREMENT = append(t.MEASUREMENT, node.(*MeasurementType))
	case *CompuMethodType:
		if t.COMPU_METHOD == nil {
			t.COMPU_METHOD = make([]*CompuMethodType, 0)
		}

		t.COMPU_METHOD = append(t.COMPU_METHOD, node.(*CompuMethodType))
	case *CompuTabType:
		if t.COMPU_TAB == nil {
			t.COMPU_TAB = make([]*CompuTabType, 0)
		}

		t.COMPU_TAB = append(t.COMPU_TAB, node.(*CompuTabType))
	case *CompuVTabType:
		if t.COMPU_VTAB == nil {
			t.COMPU_VTAB = make([]*CompuVTabType, 0)
		}

		t.COMPU_VTAB = append(t.COMPU_VTAB, node.(*CompuVTabType))
	case *CompuVTabRangeType:
		if t.COMPU_VTAB_RANGE == nil {
			t.COMPU_VTAB_RANGE = make([]*CompuVTabRangeType, 0)
		}

		t.COMPU_VTAB_RANGE = append(t.COMPU_VTAB_RANGE, node.(*CompuVTabRangeType))
	case *FunctionType:
		if t.FUNCTION == nil {
			t.FUNCTION = make([]*FunctionType, 0)
		}

		t.FUNCTION = append(t.FUNCTION, node.(*FunctionType))
	case *GroupType:
		if t.GROUP == nil {
			t.GROUP = make([]*GroupType, 0)
		}

		t.GROUP = append(t.GROUP, node.(*GroupType))
	case *RecordLayoutType:
		if t.RECORD_LAYOUT == nil {
			t.RECORD_LAYOUT = make([]*RecordLayoutType, 0)
		}

		t.RECORD_LAYOUT = append(t.RECORD_LAYOUT, node.(*RecordLayoutType))
	case *VariantCodingType:
		t.VARIANT_CODING = node.(*VariantCodingType)
	case *FrameType:
		t.FRAME = node.(*FrameType)
	case *UserRightsType:
		if t.USER_RIGHTS == nil {
			t.USER_RIGHTS = make([]*UserRightsType, 0)
		}

		t.USER_RIGHTS = append(t.USER_RIGHTS, node.(*UserRightsType))
	case *UnitType:
		if t.UNIT == nil {
			t.UNIT = make([]*UnitType, 0)
		}

		t.UNIT = append(t.UNIT, node.(*UnitType))
	default:
		panic("not implemented yet...")
	}
}

func (t *ModuleType) MarshalA2L() (result string) {
	return `/begin MODULE
/end MODULE`
}

func (t *ModuleType) A2LTag() *string {
	tag := "MODULE"
	return &tag
}

func (t *MonotonyType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *MonotonyType) MarshalA2L() (result string) {
	return marshalA2L[*MonotonyType](t)
}

func (t *MonotonyType) A2LTag() (result *string) {
	return result
}

func (t *NoAxisPtsXType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *NoAxisPtsXType) MarshalA2L() (result string) {
	return marshalA2L[*NoAxisPtsXType](t)
}

func (t *NoAxisPtsXType) A2LTag() (result *string) {
	return result
}

func (t *NoAxisPtsYType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *NoAxisPtsYType) MarshalA2L() (result string) {
	return marshalA2L[*NoAxisPtsYType](t)
}

func (t *NoAxisPtsYType) A2LTag() (result *string) {
	return result
}

func (t *NoAxisPtsZType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *NoAxisPtsZType) MarshalA2L() (result string) {
	return marshalA2L[*NoAxisPtsZType](t)
}

func (t *NoAxisPtsZType) A2LTag() (result *string) {
	return result
}

func (t *NoOfInterfacesType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *NoOfInterfacesType) MarshalA2L() (result string) {
	return marshalA2L[*NoOfInterfacesType](t)
}

func (t *NoOfInterfacesType) A2LTag() (result *string) {
	return result
}

func (t *NoRescaleXType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *NoRescaleXType) MarshalA2L() (result string) {
	return marshalA2L[*NoRescaleXType](t)
}

func (t *NoRescaleXType) A2LTag() (result *string) {
	return result
}

func (t *NoRescaleYType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *NoRescaleYType) MarshalA2L() (result string) {
	return marshalA2L[*NoRescaleYType](t)
}

func (t *NoRescaleYType) A2LTag() (result *string) {
	return result
}

func (t *NoRescaleZType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *NoRescaleZType) MarshalA2L() (result string) {
	return marshalA2L[*NoRescaleZType](t)
}

func (t *NoRescaleZType) A2LTag() (result *string) {
	return result
}

func (t *NumberType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *NumberType) MarshalA2L() (result string) {
	return marshalA2L[*NumberType](t)
}

func (t *NumberType) A2LTag() (result *string) {
	return result
}

func (t *OffsetXType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *OffsetXType) MarshalA2L() (result string) {
	return marshalA2L[*OffsetXType](t)
}

func (t *OffsetXType) A2LTag() (result *string) {
	return result
}

func (t *OffsetYType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *OffsetYType) MarshalA2L() (result string) {
	return marshalA2L[*OffsetYType](t)
}

func (t *OffsetYType) A2LTag() (result *string) {
	return result
}

func (t *OffsetZType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *OffsetZType) MarshalA2L() (result string) {
	return marshalA2L[*OffsetZType](t)
}

func (t *OffsetZType) A2LTag() (result *string) {
	return result
}

func (t *OutMeasurementType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *OutMeasurementType) MarshalA2L() (result string) {
	return marshalA2L[*OutMeasurementType](t)
}

func (t *OutMeasurementType) A2LTag() (result *string) {
	return result
}

func (t *PhoneNoType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *PhoneNoType) MarshalA2L() (result string) {
	return marshalA2L[*PhoneNoType](t)
}

func (t *PhoneNoType) A2LTag() (result *string) {
	return result
}

func (t *ProjectNoType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *ProjectNoType) MarshalA2L() (result string) {
	return fmt.Sprintf("PROJECT_NO %s", t.ProjectNumber.A2LString())
}

func (t *ProjectNoType) A2LTag() (result *string) {
	return result
}

func (t *ProjectType) MapChildNodes(node any) {
	switch node.(type) {
	case *HeaderType:
		t.HEADER = node.(*HeaderType)
	case *ModuleType:
		if t.MODULE == nil {
			t.MODULE = make([]*ModuleType, 0)
		}

		t.MODULE = append(t.MODULE, node.(*ModuleType))
	}
}

func (t *ProjectType) MarshalA2L() (result string) {
	return fmt.Sprintf(`/begin PROJECT
	%s
	%s
	%s
	%s
/end PROJECT
`, t.Name.A2LString(), t.LongIdentifier.A2LString(), t.HEADER.MarshalA2L(), t.MODULE[0].MarshalA2L())
}

func (t *ProjectType) A2LTag() *string {
	tag := "PROJECT"
	return &tag
}

func (t *ReadOnlyType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *ReadOnlyType) MarshalA2L() (result string) {
	return marshalA2L[*ReadOnlyType](t)
}

func (t *ReadOnlyType) A2LTag() (result *string) {
	return result
}

func (t *ReadWriteType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *ReadWriteType) MarshalA2L() (result string) {
	return marshalA2L[*ReadWriteType](t)
}

func (t *ReadWriteType) A2LTag() (result *string) {
	return result
}

func (t *RecordLayoutType) MapChildNodes(node any) {
	switch node.(type) {
	case *FncValuesType:
		t.FNC_VALUES = node.(*FncValuesType)
	case *IdentificationType:
		t.IDENTIFICATION = node.(*IdentificationType)
	case *AxisPtsXType:
		t.AXIS_PTS_X = node.(*AxisPtsXType)
	case *AxisPtsYType:
		t.AXIS_PTS_Y = node.(*AxisPtsYType)
	case *AxisPtsZType:
		t.AXIS_PTS_Z = node.(*AxisPtsZType)
	case *AxisRescaleXType:
		t.AXIS_RESCALE_X = node.(*AxisRescaleXType)
	case *AxisRescaleYType:
		t.AXIS_RESCALE_Y = node.(*AxisRescaleYType)
	case *AxisRescaleZType:
		t.AXIS_RESCALE_Z = node.(*AxisRescaleZType)
	case *NoAxisPtsXType:
		t.NO_AXIS_PTS_X = node.(*NoAxisPtsXType)
	case *NoAxisPtsYType:
		t.NO_AXIS_PTS_Y = node.(*NoAxisPtsYType)
	case *NoAxisPtsZType:
		t.NO_AXIS_PTS_Z = node.(*NoAxisPtsZType)
	case *NoRescaleXType:
		t.NO_RESCALE_X = node.(*NoRescaleXType)
	case *NoRescaleYType:
		t.NO_RESCALE_Y = node.(*NoRescaleYType)
	case *NoRescaleZType:
		t.NO_RESCALE_Z = node.(*NoRescaleZType)
	case *FixNoAxisPtsXType:
		t.FIX_NO_AXIS_PTS_X = node.(*FixNoAxisPtsXType)
	case *FixNoAxisPtsYType:
		t.FIX_NO_AXIS_PTS_Y = node.(*FixNoAxisPtsYType)
	case *FixNoAxisPtsZType:
		t.FIX_NO_AXIS_PTS_Z = node.(*FixNoAxisPtsZType)
	case *SrcAddrXType:
		t.SRC_ADDR_X = node.(*SrcAddrXType)
	case *SrcAddrYType:
		t.SRC_ADDR_Y = node.(*SrcAddrYType)
	case *SrcAddrZType:
		t.SRC_ADDR_Z = node.(*SrcAddrZType)
	case *RipAddrXType:
		t.RIP_ADDR_X = node.(*RipAddrXType)
	case *RipAddrYType:
		t.RIP_ADDR_Y = node.(*RipAddrYType)
	case *RipAddrZType:
		t.RIP_ADDR_Z = node.(*RipAddrZType)
	case *RipAddrWType:
		t.RIP_ADDR_W = node.(*RipAddrWType)
	case *ShiftOpXType:
		t.SHIFT_OP_X = node.(*ShiftOpXType)
	case *ShiftOpYType:
		t.SHIFT_OP_Y = node.(*ShiftOpYType)
	case *ShiftOpZType:
		t.SHIFT_OP_Z = node.(*ShiftOpZType)
	case *OffsetXType:
		t.OFFSET_X = node.(*OffsetXType)
	case *OffsetYType:
		t.OFFSET_Y = node.(*OffsetYType)
	case *OffsetZType:
		t.OFFSET_Z = node.(*OffsetZType)
	case *DistOpXType:
		t.DIST_OP_X = node.(*DistOpXType)
	case *DistOpYType:
		t.DIST_OP_Y = node.(*DistOpYType)
	case *DistOpZType:
		t.DIST_OP_Z = node.(*DistOpZType)
	case *AlignmentByteType:
		t.ALIGNMENT_BYTE = node.(*AlignmentByteType)
	case *AlignmentWordType:
		t.ALIGNMENT_WORD = node.(*AlignmentWordType)
	case *AlignmentLongType:
		t.ALIGNMENT_LONG = node.(*AlignmentLongType)
	case *AlignmentFloat32IeeeType:
		t.ALIGNMENT_FLOAT32_IEEE = node.(*AlignmentFloat32IeeeType)
	case *AlignmentFloat64IeeeType:
		t.ALIGNMENT_FLOAT64_IEEE = node.(*AlignmentFloat64IeeeType)
	case *ReservedType:
		if t.RESERVED == nil {
			t.RESERVED = make([]*ReservedType, 0)
		}

		t.RESERVED = append(t.RESERVED, node.(*ReservedType))
	default:
		panic("not implemented yet...")
	}
}

func (t *RecordLayoutType) MarshalA2L() (result string) {
	return marshalA2L[*RecordLayoutType](t)
}

func (t *RecordLayoutType) A2LTag() (result *string) {
	return result
}

func (t *RefCharacteristicType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *RefCharacteristicType) MarshalA2L() (result string) {
	return marshalA2L[*RefCharacteristicType](t)
}

func (t *RefCharacteristicType) A2LTag() (result *string) {
	return result
}

func (t *RefGroupType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *RefGroupType) MarshalA2L() (result string) {
	return marshalA2L[*RefGroupType](t)
}

func (t *RefGroupType) A2LTag() (result *string) {
	return result
}

func (t *RefMeasurementType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *RefMeasurementType) MarshalA2L() (result string) {
	return marshalA2L[*RefMeasurementType](t)
}

func (t *RefMeasurementType) A2LTag() (result *string) {
	return result
}

func (t *RefMemorySegmentType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *RefMemorySegmentType) MarshalA2L() (result string) {
	return marshalA2L[*RefMemorySegmentType](t)
}

func (t *RefMemorySegmentType) A2LTag() (result *string) {
	return result
}

func (t *RefUnitType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *RefUnitType) MarshalA2L() (result string) {
	return marshalA2L[*RefUnitType](t)
}

func (t *RefUnitType) A2LTag() (result *string) {
	return result
}

func (t *ReservedType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *ReservedType) MarshalA2L() (result string) {
	return marshalA2L[*ReservedType](t)
}

func (t *ReservedType) A2LTag() (result *string) {
	return result
}

func (t *RightShiftType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *RightShiftType) MarshalA2L() (result string) {
	return marshalA2L[*RightShiftType](t)
}

func (t *RightShiftType) A2LTag() (result *string) {
	return result
}

func (t *RipAddrWType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *RipAddrWType) MarshalA2L() (result string) {
	return marshalA2L[*RipAddrWType](t)
}

func (t *RipAddrWType) A2LTag() (result *string) {
	return result
}

func (t *RipAddrXType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *RipAddrXType) MarshalA2L() (result string) {
	return marshalA2L[*RipAddrXType](t)
}

func (t *RipAddrXType) A2LTag() (result *string) {
	return result
}

func (t *RipAddrYType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *RipAddrYType) MarshalA2L() (result string) {
	return marshalA2L[*RipAddrYType](t)
}

func (t *RipAddrYType) A2LTag() (result *string) {
	return result
}

func (t *RipAddrZType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *RipAddrZType) MarshalA2L() (result string) {
	return marshalA2L[*RipAddrZType](t)
}

func (t *RipAddrZType) A2LTag() (result *string) {
	return result
}

func (t *RootNodeType) MapChildNodes(node any) {
	switch node.(type) {
	case *ASAP2VersionType:
		t.ASAP2_VERSION = node.(*ASAP2VersionType)
	case *A2MLVersionType:
		t.A2ML_VERSION = node.(*A2MLVersionType)
	case *ProjectType:
		t.PROJECT = node.(*ProjectType)
	default:
		panic("not implemented yet...")
	}
}

func (t *RootNodeType) MarshalA2L() (result string) {
	return fmt.Sprintf("%s\n%s\n%s", t.ASAP2_VERSION.MarshalA2L(), t.A2ML_VERSION.MarshalA2L(), t.PROJECT.MarshalA2L())
}

func (t *RootNodeType) A2LTag() (result *string) {
	return result
}

func (t *RootType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *RootType) MarshalA2L() (result string) {
	return marshalA2L[*RootType](t)
}

func (t *RootType) A2LTag() (result *string) {
	return result
}

func (t *ShiftOpXType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *ShiftOpXType) MarshalA2L() (result string) {
	return marshalA2L[*ShiftOpXType](t)
}

func (t *ShiftOpXType) A2LTag() (result *string) {
	return result
}

func (t *ShiftOpYType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *ShiftOpYType) MarshalA2L() (result string) {
	return marshalA2L[*ShiftOpYType](t)
}

func (t *ShiftOpYType) A2LTag() (result *string) {
	return result
}

func (t *ShiftOpZType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *ShiftOpZType) MarshalA2L() (result string) {
	return marshalA2L[*ShiftOpZType](t)
}

func (t *ShiftOpZType) A2LTag() (result *string) {
	return result
}

func (t *SiExponentsType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *SiExponentsType) MarshalA2L() (result string) {
	return marshalA2L[*SiExponentsType](t)
}

func (t *SiExponentsType) A2LTag() (result *string) {
	return result
}

func (t *SignExtendType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *SignExtendType) MarshalA2L() (result string) {
	return marshalA2L[*SignExtendType](t)
}

func (t *SignExtendType) A2LTag() (result *string) {
	return result
}

func (t *SrcAddrXType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *SrcAddrXType) MarshalA2L() (result string) {
	return marshalA2L[*SrcAddrXType](t)
}

func (t *SrcAddrXType) A2LTag() (result *string) {
	return result
}

func (t *SrcAddrYType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *SrcAddrYType) MarshalA2L() (result string) {
	return marshalA2L[*SrcAddrYType](t)
}

func (t *SrcAddrYType) A2LTag() (result *string) {
	return result
}

func (t *SrcAddrZType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *SrcAddrZType) MarshalA2L() (result string) {
	return marshalA2L[*SrcAddrZType](t)
}

func (t *SrcAddrZType) A2LTag() (result *string) {
	return result
}

func (t *SRecLayoutType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *SRecLayoutType) MarshalA2L() (result string) {
	return marshalA2L[*SRecLayoutType](t)
}

func (t *SRecLayoutType) A2LTag() (result *string) {
	return result
}

func (t *SubFunctionType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *SubFunctionType) MarshalA2L() (result string) {
	return marshalA2L[*SubFunctionType](t)
}

func (t *SubFunctionType) A2LTag() (result *string) {
	return result
}

func (t *SubGroupType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *SubGroupType) MarshalA2L() (result string) {
	return marshalA2L[*SubGroupType](t)
}

func (t *SubGroupType) A2LTag() (result *string) {
	return result
}

func (t *SupplierType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *SupplierType) MarshalA2L() (result string) {
	return marshalA2L[*SupplierType](t)
}

func (t *SupplierType) A2LTag() (result *string) {
	return result
}

func (t *SystemConstantType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *SystemConstantType) MarshalA2L() (result string) {
	return marshalA2L[*SystemConstantType](t)
}

func (t *SystemConstantType) A2LTag() (result *string) {
	return result
}

func (t *UnitConversionType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *UnitConversionType) MarshalA2L() (result string) {
	return marshalA2L[*UnitConversionType](t)
}

func (t *UnitConversionType) A2LTag() (result *string) {
	return result
}

func (t *UnitType) MapChildNodes(node any) {
	switch node.(type) {
	case *SiExponentsType:
		t.SI_EXPONENTS = node.(*SiExponentsType)
	case *RefUnitType:
		t.REF_UNIT = node.(*RefUnitType)
	case *UnitConversionType:
		t.UNIT_CONVERSION = node.(*UnitConversionType)
	default:
		panic("not implemented yet...")
	}
}

func (t *UnitType) MarshalA2L() (result string) {
	return marshalA2L[*UnitType](t)
}

func (t *UnitType) A2LTag() (result *string) {
	return result
}

func (t *UserRightsType) MapChildNodes(node any) {
	switch node.(type) {
	case *RefGroupType:
		if t.REF_GROUP == nil {
			t.REF_GROUP = make([]*RefGroupType, 0)
		}

		t.REF_GROUP = append(t.REF_GROUP, node.(*RefGroupType))
	case *ReadOnlyType:
		t.READ_ONLY = node.(*ReadOnlyType)
	default:
		panic("not implemented yet...")
	}
}

func (t *UserRightsType) MarshalA2L() (result string) {
	return marshalA2L[*UserRightsType](t)
}

func (t *UserRightsType) A2LTag() (result *string) {
	return result
}

func (t *UserType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *UserType) MarshalA2L() (result string) {
	return marshalA2L[*UserType](t)
}

func (t *UserType) A2LTag() (result *string) {
	return result
}

func (t *VarAddressType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *VarAddressType) MarshalA2L() (result string) {
	return marshalA2L[*VarAddressType](t)
}

func (t *VarAddressType) A2LTag() (result *string) {
	return result
}

func (t *VarCharacteristicType) MapChildNodes(node any) {
	switch node.(type) {
	case *VarAddressType:
		t.VAR_ADDRESS = node.(*VarAddressType)
	default:
		panic("not implemented yet...")
	}
}

func (t *VarCharacteristicType) MarshalA2L() (result string) {
	return marshalA2L[*VarCharacteristicType](t)
}

func (t *VarCharacteristicType) A2LTag() (result *string) {
	return result
}

func (t *VarCriterionType) MapChildNodes(node any) {
	switch node.(type) {
	case *VarMeasurementType:
		t.VAR_MEASUREMENT = node.(*VarMeasurementType)
	case *VarSelectionCharacteristicType:
		t.VAR_SELECTION_CHARACTERISTIC = node.(*VarSelectionCharacteristicType)
	default:
		panic("not implemented yet...")
	}
}

func (t *VarCriterionType) MarshalA2L() (result string) {
	return marshalA2L[*VarCriterionType](t)
}

func (t *VarCriterionType) A2LTag() (result *string) {
	return result
}

func (t *VarForbiddenCombType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *VarForbiddenCombType) MarshalA2L() (result string) {
	return marshalA2L[*VarForbiddenCombType](t)
}

func (t *VarForbiddenCombType) A2LTag() (result *string) {
	return result
}

func (t *VariantCodingType) MapChildNodes(node any) {
	switch node.(type) {
	case *VarSeparatorType:
		t.VAR_SEPARATOR = node.(*VarSeparatorType)
	case *VarNamingType:
		t.VAR_NAMING = node.(*VarNamingType)
	case *VarCriterionType:
		if t.VAR_CRITERION == nil {
			t.VAR_CRITERION = make([]*VarCriterionType, 0)
		}

		t.VAR_CRITERION = append(t.VAR_CRITERION, node.(*VarCriterionType))
	case *VarForbiddenCombType:
		if t.VAR_FORBIDDEN_COMB == nil {
			t.VAR_FORBIDDEN_COMB = make([]*VarForbiddenCombType, 0)
		}

		t.VAR_FORBIDDEN_COMB = append(t.VAR_FORBIDDEN_COMB, node.(*VarForbiddenCombType))
	case *VarCharacteristicType:
		if t.VAR_CHARACTERISTIC == nil {
			t.VAR_CHARACTERISTIC = make([]*VarCharacteristicType, 0)
		}

		t.VAR_CHARACTERISTIC = append(t.VAR_CHARACTERISTIC, node.(*VarCharacteristicType))
	default:
		panic("not implemented yet...")
	}
}

func (t *VariantCodingType) MarshalA2L() (result string) {
	return marshalA2L[*VariantCodingType](t)
}

func (t *VariantCodingType) A2LTag() (result *string) {
	return result
}

func (t *VarMeasurementType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *VarMeasurementType) MarshalA2L() (result string) {
	return marshalA2L[*VarMeasurementType](t)
}

func (t *VarMeasurementType) A2LTag() (result *string) {
	return result
}

func (t *VarNamingType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *VarNamingType) MarshalA2L() (result string) {
	return marshalA2L[*VarNamingType](t)
}

func (t *VarNamingType) A2LTag() (result *string) {
	return result
}

func (t *VarSelectionCharacteristicType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *VarSelectionCharacteristicType) MarshalA2L() (result string) {
	return marshalA2L[*VarSelectionCharacteristicType](t)
}

func (t *VarSelectionCharacteristicType) A2LTag() (result *string) {
	return result
}

func (t *VarSeparatorType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *VarSeparatorType) MarshalA2L() (result string) {
	return marshalA2L[*VarSeparatorType](t)
}

func (t *VarSeparatorType) A2LTag() (result *string) {
	return result
}

func (t *VersionType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *VersionType) MarshalA2L() (result string) {
	return fmt.Sprintf("VERSION %s", t.VersionIdentifier.A2LString())
}

func (t *VersionType) A2LTag() (result *string) {
	return result
}

func (t *VirtualCharacteristicType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *VirtualCharacteristicType) MarshalA2L() (result string) {
	return marshalA2L[*VirtualCharacteristicType](t)
}

func (t *VirtualCharacteristicType) A2LTag() (result *string) {
	return result
}

func (t *VirtualType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *VirtualType) MarshalA2L() (result string) {
	return marshalA2L[*VirtualType](t)
}

func (t *VirtualType) A2LTag() (result *string) {
	return result
}
