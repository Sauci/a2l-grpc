package a2l

import (
	"fmt"
	"strings"
)

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

func (t *A2MLVersionType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("A2ML_VERSION %s %s", t.VersionNo.A2LString(), t.UpgradeNo.A2LString()), indentLevel, indentString)
}

func (t *AddrEpkType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AddrEpkType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("ADDR_EPK %s", t.Address.A2LString()), indentLevel, indentString)
}

func (t *AlignmentByteType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AlignmentByteType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("ALIGNMENT_BYTE %s", t.AlignmentBorder.A2LString()), indentLevel, indentString)
}

func (t *AlignmentFloat32IeeeType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AlignmentFloat32IeeeType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("ALIGNMENT_FLOAT32_IEEE %s", t.AlignmentBorder.A2LString()), indentLevel, indentString)
}

func (t *AlignmentFloat64IeeeType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AlignmentFloat64IeeeType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("ALIGNMENT_FLOAT64_IEEE %s", t.AlignmentBorder.A2LString()), indentLevel, indentString)
}

func (t *AlignmentLongType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AlignmentLongType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("ALIGNMENT_LONG %s", t.AlignmentBorder.A2LString()), indentLevel, indentString)
}

func (t *AlignmentWordType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AlignmentWordType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("ALIGNMENT_WORD %s", t.AlignmentBorder.A2LString()), indentLevel, indentString)
}

func (t *AnnotationLabelType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AnnotationLabelType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("ANNOTATION_LABEL %s", t.Label.A2LString()), indentLevel, indentString)
}

func (t *AnnotationOriginType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AnnotationOriginType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("ANNOTATION_ORIGIN %s", t.Origin.A2LString()), indentLevel, indentString)
}

func (t *AnnotationTextType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AnnotationTextType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin ANNOTATION_TEXT"), indentLevel, indentString)}

	for _, annotationText := range t.AnnotationText {
		tmpResult = append(tmpResult, indentContent(annotationText.A2LString(), indentLevel+1, indentString))
	}

	tmpResult = append(tmpResult, indentContent("/end ANNOTATION_TEXT", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
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

func (t *AnnotationType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin ANNOTATION"), indentLevel, indentString)}

	if t.ANNOTATION_LABEL != nil {
		tmpResult = append(tmpResult, t.ANNOTATION_LABEL.MarshalA2L(indentLevel+1, indentString))
	}

	if t.ANNOTATION_ORIGIN != nil {
		tmpResult = append(tmpResult, t.ANNOTATION_ORIGIN.MarshalA2L(indentLevel+1, indentString))
	}

	if t.ANNOTATION_TEXT != nil {
		tmpResult = append(tmpResult, t.ANNOTATION_TEXT.MarshalA2L(indentLevel+1, indentString))
	}

	tmpResult = append(tmpResult, indentContent("/end ANNOTATION", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *ArraySizeType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *ArraySizeType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("ARRAY_SIZE %s", t.Number.A2LString()), indentLevel, indentString)
}

func (t *ASAP2VersionType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *ASAP2VersionType) MarshalA2L(indentLevel int, indentString string) string {
	return indentContent(fmt.Sprintf("ASAP2_VERSION %s %s", t.VersionNo.A2LString(), t.UpgradeNo.A2LString()), indentLevel, indentString)
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

func (t *AxisDescrType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin AXIS_DESCR %s %s %s %s %s %s",
		t.Attribute,
		t.InputQuantity.A2LString(),
		t.Conversion.A2LString(),
		t.MaxAxisPoints.A2LString(),
		t.LowerLimit.A2LString(),
		t.UpperLimit.A2LString()), indentLevel, indentString)}

	if t.READ_ONLY != nil {
		tmpResult = append(tmpResult, t.READ_ONLY.MarshalA2L(indentLevel+1, indentString))
	}

	if t.FORMAT != nil {
		tmpResult = append(tmpResult, t.FORMAT.MarshalA2L(indentLevel+1, indentString))
	}

	if t.ANNOTATION != nil {
		for _, annotation := range t.ANNOTATION {
			tmpResult = append(tmpResult, annotation.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.AXIS_PTS_REF != nil {
		tmpResult = append(tmpResult, t.AXIS_PTS_REF.MarshalA2L(indentLevel+1, indentString))
	}

	if t.MAX_GRAD != nil {
		tmpResult = append(tmpResult, t.MAX_GRAD.MarshalA2L(indentLevel+1, indentString))
	}

	if t.MONOTONY != nil {
		tmpResult = append(tmpResult, t.MONOTONY.MarshalA2L(indentLevel+1, indentString))
	}

	if t.BYTE_ORDER != nil {
		tmpResult = append(tmpResult, t.BYTE_ORDER.MarshalA2L(indentLevel+1, indentString))
	}

	if t.EXTENDED_LIMITS != nil {
		tmpResult = append(tmpResult, t.EXTENDED_LIMITS.MarshalA2L(indentLevel+1, indentString))
	}

	if t.FIX_AXIS_PAR != nil {
		tmpResult = append(tmpResult, t.FIX_AXIS_PAR.MarshalA2L(indentLevel+1, indentString))
	}

	if t.FIX_AXIS_PAR_DIST != nil {
		tmpResult = append(tmpResult, t.FIX_AXIS_PAR_DIST.MarshalA2L(indentLevel+1, indentString))
	}

	if t.FIX_AXIS_PAR_LIST != nil {
		tmpResult = append(tmpResult, t.FIX_AXIS_PAR_LIST.MarshalA2L(indentLevel+1, indentString))
	}

	if t.DEPOSIT != nil {
		tmpResult = append(tmpResult, t.DEPOSIT.MarshalA2L(indentLevel+1, indentString))
	}

	if t.CURVE_AXIS_REF != nil {
		tmpResult = append(tmpResult, t.CURVE_AXIS_REF.MarshalA2L(indentLevel+1, indentString))
	}

	tmpResult = append(tmpResult, indentContent("/end AXIS_DESCR", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *AxisPtsRefType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AxisPtsRefType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("AXIS_PTS_REF %s", t.AxisPoints.A2LString()), indentLevel, indentString)
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

func (t *AxisPtsType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin AXIS_PTS %s %s %s %s %s %s %s %s %s %s",
		t.Name.A2LString(),
		t.LongIdentifier.A2LString(),
		t.Address.A2LString(),
		t.InputQuantity.A2LString(),
		t.DepositR.A2LString(),
		t.MaxDiff.A2LString(),
		t.Conversion.A2LString(),
		t.MaxAxisPoints.A2LString(),
		t.LowerLimit.A2LString(),
		t.UpperLimit.A2LString()), indentLevel, indentString)}

	if t.DISPLAY_IDENTIFIER != nil {
		tmpResult = append(tmpResult, t.DISPLAY_IDENTIFIER.MarshalA2L(indentLevel+1, indentString))
	}

	if t.READ_ONLY != nil {
		tmpResult = append(tmpResult, t.READ_ONLY.MarshalA2L(indentLevel+1, indentString))
	}

	if t.FORMAT != nil {
		tmpResult = append(tmpResult, t.FORMAT.MarshalA2L(indentLevel+1, indentString))
	}

	if t.DEPOSIT != nil {
		tmpResult = append(tmpResult, t.DEPOSIT.MarshalA2L(indentLevel+1, indentString))
	}

	if t.BYTE_ORDER != nil {
		tmpResult = append(tmpResult, t.BYTE_ORDER.MarshalA2L(indentLevel+1, indentString))
	}

	if t.FUNCTION_LIST != nil {
		tmpResult = append(tmpResult, t.FUNCTION_LIST.MarshalA2L(indentLevel+1, indentString))
	}

	if t.REF_MEMORY_SEGMENT != nil {
		tmpResult = append(tmpResult, t.REF_MEMORY_SEGMENT.MarshalA2L(indentLevel+1, indentString))
	}

	if t.GUARD_RAILS != nil {
		tmpResult = append(tmpResult, t.GUARD_RAILS.MarshalA2L(indentLevel+1, indentString))
	}

	if t.EXTENDED_LIMITS != nil {
		tmpResult = append(tmpResult, t.EXTENDED_LIMITS.MarshalA2L(indentLevel+1, indentString))
	}

	if t.ANNOTATION != nil {
		for _, annotation := range t.ANNOTATION {
			tmpResult = append(tmpResult, annotation.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.IF_DATA != nil {
		for _, ifData := range t.IF_DATA {
			tmpResult = append(tmpResult, ifData.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.CALIBRATION_ACCESS != nil {
		tmpResult = append(tmpResult, t.CALIBRATION_ACCESS.MarshalA2L(indentLevel+1, indentString))
	}

	if t.ECU_ADDRESS_EXTENSION != nil {
		tmpResult = append(tmpResult, t.ECU_ADDRESS_EXTENSION.MarshalA2L(indentLevel+1, indentString))
	}

	tmpResult = append(tmpResult, indentContent("/end AXIS_PTS", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *AxisPtsXType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AxisPtsXType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("AXIS_PTS_X %s %s %s %s",
		t.Position.A2LString(),
		t.DataType.A2LString(),
		t.IndexIncr.A2LString(),
		t.Addressing.A2LString()), indentLevel, indentString)
}

func (t *AxisPtsYType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AxisPtsYType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("AXIS_PTS_Y %s %s %s %s",
		t.Position.A2LString(),
		t.DataType.A2LString(),
		t.IndexIncr.A2LString(),
		t.Addressing.A2LString()), indentLevel, indentString)
}

func (t *AxisPtsZType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AxisPtsZType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("AXIS_PTS_Z %s %s %s %s",
		t.Position.A2LString(),
		t.DataType.A2LString(),
		t.IndexIncr.A2LString(),
		t.Addressing.A2LString()), indentLevel, indentString)
}

func (t *AxisRescaleXType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AxisRescaleXType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("AXIS_RESCALE_X %s %s %s %s %s",
		t.Position.A2LString(),
		t.DataType.A2LString(),
		t.MaxNumberOfRescalePairs.A2LString(),
		t.IndexIncr.A2LString(),
		t.Addressing.A2LString()), indentLevel, indentString)
}

func (t *AxisRescaleYType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AxisRescaleYType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("AXIS_RESCALE_Y %s %s %s %s %s",
		t.Position.A2LString(),
		t.DataType.A2LString(),
		t.MaxNumberOfRescalePairs.A2LString(),
		t.IndexIncr.A2LString(),
		t.Addressing.A2LString()), indentLevel, indentString)
}

func (t *AxisRescaleZType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *AxisRescaleZType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("AXIS_RESCALE_Z %s %s %s %s %s",
		t.Position.A2LString(),
		t.DataType.A2LString(),
		t.MaxNumberOfRescalePairs.A2LString(),
		t.IndexIncr.A2LString(),
		t.Addressing.A2LString()), indentLevel, indentString)
}

func (t *BitMaskType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *BitMaskType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("BIT_MASK %s", t.Mask.A2LString()), indentLevel, indentString)
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

func (t *BitOperationType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin BIT_OPERATION"), indentLevel, indentString)}

	if t.LEFT_SHIFT != nil {
		tmpResult = append(tmpResult, t.LEFT_SHIFT.MarshalA2L(indentLevel+1, indentString))
	}

	if t.RIGHT_SHIFT != nil {
		tmpResult = append(tmpResult, t.RIGHT_SHIFT.MarshalA2L(indentLevel+1, indentString))
	}

	if t.SIGN_EXTEND != nil {
		tmpResult = append(tmpResult, t.SIGN_EXTEND.MarshalA2L(indentLevel+1, indentString))
	}

	tmpResult = append(tmpResult, indentContent("/end BIT_OPERATION", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *ByteOrderType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *ByteOrderType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("BYTE_ORDER %s", t.ByteOrder), indentLevel, indentString)
}

func (t *CalibrationAccessType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *CalibrationAccessType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("CALIBRATION_ACCESS %s", t.Type), indentLevel, indentString)
}

func (t *CalibrationHandleType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *CalibrationHandleType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin CALIBRATION_HANDLE"), indentLevel, indentString)}

	if t.Handle != nil {
		for _, handle := range t.Handle {
			tmpResult = append(tmpResult, indentContent(handle.A2LString(), indentLevel+1, indentString))
		}
	}

	tmpResult = append(tmpResult, indentContent("/end CALIBRATION_HANDLE", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
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

func (t *CalibrationMethodType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin CALIBRATION_METHOD %s %s",
		t.Method.A2LString(),
		t.Version.A2LString()), indentLevel, indentString)}

	if t.CALIBRATION_HANDLE != nil {
		for _, calibrationHandle := range t.CALIBRATION_HANDLE {
			tmpResult = append(tmpResult, calibrationHandle.MarshalA2L(indentLevel+1, indentString))
		}
	}

	tmpResult = append(tmpResult, indentContent("/end CALIBRATION_METHOD", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
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

func (t *CharacteristicType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin CHARACTERISTIC %s %s %s %s %s %s %s %s %s",
		t.Name.A2LString(),
		t.LongIdentifier.A2LString(),
		t.Type,
		t.Address.A2LString(),
		t.Deposit.A2LString(),
		t.MaxDiff.A2LString(),
		t.Conversion.A2LString(),
		t.LowerLimit.A2LString(),
		t.UpperLimit.A2LString()), indentLevel, indentString)}

	if t.DISPLAY_IDENTIFIER != nil {
		tmpResult = append(tmpResult, t.DISPLAY_IDENTIFIER.MarshalA2L(indentLevel+1, indentString))
	}

	if t.FORMAT != nil {
		tmpResult = append(tmpResult, t.FORMAT.MarshalA2L(indentLevel+1, indentString))
	}

	if t.BYTE_ORDER != nil {
		tmpResult = append(tmpResult, t.BYTE_ORDER.MarshalA2L(indentLevel+1, indentString))
	}

	if t.BIT_MASK != nil {
		tmpResult = append(tmpResult, t.BIT_MASK.MarshalA2L(indentLevel+1, indentString))
	}

	if t.FUNCTION_LIST != nil {
		tmpResult = append(tmpResult, t.FUNCTION_LIST.MarshalA2L(indentLevel+1, indentString))
	}

	if t.NUMBER != nil {
		tmpResult = append(tmpResult, t.NUMBER.MarshalA2L(indentLevel+1, indentString))
	}

	if t.EXTENDED_LIMITS != nil {
		tmpResult = append(tmpResult, t.EXTENDED_LIMITS.MarshalA2L(indentLevel+1, indentString))
	}

	if t.READ_ONLY != nil {
		tmpResult = append(tmpResult, t.READ_ONLY.MarshalA2L(indentLevel+1, indentString))
	}

	if t.GUARD_RAILS != nil {
		tmpResult = append(tmpResult, t.GUARD_RAILS.MarshalA2L(indentLevel+1, indentString))
	}

	if t.MAP_LIST != nil {
		tmpResult = append(tmpResult, t.MAP_LIST.MarshalA2L(indentLevel+1, indentString))
	}

	if t.MAX_REFRESH != nil {
		tmpResult = append(tmpResult, t.MAX_REFRESH.MarshalA2L(indentLevel+1, indentString))
	}

	if t.DEPENDENT_CHARACTERISTIC != nil {
		tmpResult = append(tmpResult, t.DEPENDENT_CHARACTERISTIC.MarshalA2L(indentLevel+1, indentString))
	}

	if t.VIRTUAL_CHARACTERISTIC != nil {
		tmpResult = append(tmpResult, t.VIRTUAL_CHARACTERISTIC.MarshalA2L(indentLevel+1, indentString))
	}

	if t.REF_MEMORY_SEGMENT != nil {
		tmpResult = append(tmpResult, t.REF_MEMORY_SEGMENT.MarshalA2L(indentLevel+1, indentString))
	}

	if t.ANNOTATION != nil {
		for _, annotation := range t.ANNOTATION {
			tmpResult = append(tmpResult, annotation.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.COMPARISON_QUANTITY != nil {
		tmpResult = append(tmpResult, t.COMPARISON_QUANTITY.MarshalA2L(indentLevel+1, indentString))
	}

	if t.IF_DATA != nil {
		for _, ifData := range t.IF_DATA {
			tmpResult = append(tmpResult, ifData.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.AXIS_DESCR != nil {
		for _, axisDescr := range t.AXIS_DESCR {
			tmpResult = append(tmpResult, axisDescr.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.CALIBRATION_ACCESS != nil {
		tmpResult = append(tmpResult, t.CALIBRATION_ACCESS.MarshalA2L(indentLevel+1, indentString))
	}

	if t.MATRIX_DIM != nil {
		tmpResult = append(tmpResult, t.MATRIX_DIM.MarshalA2L(indentLevel+1, indentString))
	}

	if t.ECU_ADDRESS_EXTENSION != nil {
		tmpResult = append(tmpResult, t.ECU_ADDRESS_EXTENSION.MarshalA2L(indentLevel+1, indentString))
	}

	tmpResult = append(tmpResult, indentContent("/end CHARACTERISTIC", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *CoeffsType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *CoeffsType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("COEFFS %s %s %s %s %s %s",
		t.A.A2LString(),
		t.B.A2LString(),
		t.C.A2LString(),
		t.D.A2LString(),
		t.E.A2LString(),
		t.F.A2LString()), indentLevel, indentString)
}

func (t *ComparisonQuantityType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *ComparisonQuantityType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("COMPARISON_QUANTITY %s", t.Name.A2LString()), indentLevel, indentString)
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

func (t *CompuMethodType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin COMPU_METHOD %s %s %s %s %s",
		t.Name.A2LString(),
		t.LongIdentifier.A2LString(),
		t.ConversionType,
		t.Format.A2LString(),
		t.Unit.A2LString()), indentLevel, indentString)}

	if t.FORMULA != nil {
		tmpResult = append(tmpResult, t.FORMULA.MarshalA2L(indentLevel+1, indentString))
	}

	if t.COEFFS != nil {
		tmpResult = append(tmpResult, t.COEFFS.MarshalA2L(indentLevel+1, indentString))
	}

	if t.COMPU_TAB_REF != nil {
		tmpResult = append(tmpResult, t.COMPU_TAB_REF.MarshalA2L(indentLevel+1, indentString))
	}

	if t.REF_UNIT != nil {
		tmpResult = append(tmpResult, t.REF_UNIT.MarshalA2L(indentLevel+1, indentString))
	}

	tmpResult = append(tmpResult, indentContent("/end COMPU_METHOD", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *CompuTabRefType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *CompuTabRefType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("COMPU_TAB_REF %s", t.ConversionTable.A2LString()), indentLevel, indentString)
}

func (t *CompuTabType) MapChildNodes(node any) {
	switch node.(type) {
	case *DefaultValueType:
		t.DEFAULT_VALUE = node.(*DefaultValueType)
	default:
		panic("not implemented yet...")
	}
}

func (t *CompuTabType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin COMPU_TAB %s %s %s %s",
		t.Name.A2LString(),
		t.LongIdentifier.A2LString(),
		t.ConversionType,
		t.NumberValuePairs.A2LString()), indentLevel, indentString)}

	if t.InValOutVal != nil {
		for _, inValOutVal := range t.InValOutVal {
			tmpResult = append(tmpResult, indentContent(fmt.Sprintf("%s %s",
				inValOutVal.InVal.A2LString(), inValOutVal.OutVal.A2LString()), indentLevel+1, indentString))
		}
	}

	if t.DEFAULT_VALUE != nil {
		tmpResult = append(tmpResult, t.DEFAULT_VALUE.MarshalA2L(indentLevel+1, indentString))
	}

	tmpResult = append(tmpResult, indentContent("/end COMPU_TAB", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *CompuVTabRangeType) MapChildNodes(node any) {
	switch node.(type) {
	case *DefaultValueType:
		t.DEFAULT_VALUE = node.(*DefaultValueType)
	default:
		panic("not implemented yet...")
	}
}

func (t *CompuVTabRangeType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin COMPU_VTAB_RANGE %s %s %s",
		t.Name.A2LString(),
		t.LongIdentifier.A2LString(),
		t.NumberOfValuesTriples.A2LString()), indentLevel, indentString)}

	if t.InValMinInValMaxOutVal != nil {
		for _, inValMinInValMaxOutVal := range t.InValMinInValMaxOutVal {
			tmpResult = append(tmpResult, indentContent(fmt.Sprintf("%s %s %s",
				inValMinInValMaxOutVal.InValMin.A2LString(),
				inValMinInValMaxOutVal.InValMax.A2LString(),
				inValMinInValMaxOutVal.OutVal.A2LString()), indentLevel+1, indentString))
		}
	}

	if t.DEFAULT_VALUE != nil {
		tmpResult = append(tmpResult, t.DEFAULT_VALUE.MarshalA2L(indentLevel+1, indentString))
	}

	tmpResult = append(tmpResult, indentContent("/end COMPU_VTAB_RANGE", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *CompuVTabType) MapChildNodes(node any) {
	switch node.(type) {
	case *DefaultValueType:
		t.DEFAULT_VALUE = node.(*DefaultValueType)
	default:
		panic("not implemented yet...")
	}
}

func (t *CompuVTabType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin COMPU_VTAB %s %s %s %s",
		t.Name.A2LString(),
		t.LongIdentifier.A2LString(),
		t.ConversionType,
		t.NumberValuePairs.A2LString()), indentLevel, indentString)}

	if t.InValOutVal != nil {
		for _, inValOutVal := range t.InValOutVal {
			tmpResult = append(tmpResult, indentContent(fmt.Sprintf("%s %s",
				inValOutVal.InVal.A2LString(),
				inValOutVal.OutVal.A2LString()), indentLevel+1, indentString))
		}
	}

	if t.DEFAULT_VALUE != nil {
		tmpResult = append(tmpResult, t.DEFAULT_VALUE.MarshalA2L(indentLevel+1, indentString))
	}

	tmpResult = append(tmpResult, indentContent("/end COMPU_VTAB", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *CpuTypeType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *CpuTypeType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("CPU_TYPE %s", t.Cpu.A2LString()), indentLevel, indentString)
}

func (t *CurveAxisRefType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *CurveAxisRefType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("CURVE_AXIS_REF %s", t.CurveAxis.A2LString()), indentLevel, indentString)
}

func (t *CustomerNoType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *CustomerNoType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("CUSTOMER_NO %s", t.Number.A2LString()), indentLevel, indentString)
}

func (t *CustomerType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *CustomerType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("CUSTOMER %s", t.Customer.A2LString()), indentLevel, indentString)
}

func (t *DataSizeType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *DataSizeType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("DATA_SIZE %s", t.Size.A2LString()), indentLevel, indentString)
}

func (t *DefaultValueType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *DefaultValueType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("DEFAULT_VALUE %s", t.DisplayString.A2LString()), indentLevel, indentString)
}

func (t *DefCharacteristicType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *DefCharacteristicType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin DEF_CHARACTERISTIC"), indentLevel, indentString)}

	if t.Identifier != nil {
		for _, identifier := range t.Identifier {
			tmpResult = append(tmpResult, indentContent(identifier.A2LString(), indentLevel+1, indentString))
		}
	}

	tmpResult = append(tmpResult, indentContent("/end DEF_CHARACTERISTIC", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *DependentCharacteristicType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *DependentCharacteristicType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin DEPENDENT_CHARACTERISTIC %s",
		t.Formula.A2LString()), indentLevel, indentString)}

	if t.Characteristic != nil {
		for _, characteristic := range t.Characteristic {
			tmpResult = append(tmpResult, indentContent(characteristic.A2LString(), indentLevel+1, indentString))
		}
	}

	tmpResult = append(tmpResult, indentContent("/end DEPENDENT_CHARACTERISTIC", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *DepositType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *DepositType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("DEPOSIT %s", t.Mode), indentLevel, indentString)
}

func (t *DisplayIdentifierType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *DisplayIdentifierType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("DISPLAY_IDENTIFIER %s", t.DisplayName.A2LString()), indentLevel, indentString)
}

func (t *DistOpXType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *DistOpXType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("DIST_OP_X %s %s",
		t.Position.A2LString(),
		t.DataType.A2LString()), indentLevel, indentString)
}

func (t *DistOpYType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *DistOpYType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("DIST_OP_Y %s %s",
		t.Position.A2LString(),
		t.DataType.A2LString()), indentLevel, indentString)
}

func (t *DistOpZType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *DistOpZType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("DIST_OP_Z %s %s",
		t.Position.A2LString(),
		t.DataType.A2LString()), indentLevel, indentString)
}

func (t *EcuAddressExtensionType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *EcuAddressExtensionType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("ECU_ADDRESS_EXTENSION %s", t.Extension.A2LString()), indentLevel, indentString)
}

func (t *EcuAddressType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *EcuAddressType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("ECU_ADDRESS %s", t.Address.A2LString()), indentLevel, indentString)
}

func (t *EcuCalibrationOffsetType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *EcuCalibrationOffsetType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("ECU_CALIBRATION_OFFSET %s", t.Offset.A2LString()), indentLevel, indentString)
}

func (t *EcuType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *EcuType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("ECU %s", t.ControlUnit.A2LString()), indentLevel, indentString)
}

func (t *EpkType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *EpkType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("EPK %s", t.Identifier.A2LString()), indentLevel, indentString)
}

func (t *ErrorMaskType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *ErrorMaskType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("ERROR_MASK %s", t.Mask.A2LString()), indentLevel, indentString)
}

func (t *ExtendedLimitsType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *ExtendedLimitsType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("EXTENDED_LIMITS %s %s",
		t.LowerLimit.A2LString(),
		t.UpperLimit.A2LString()), indentLevel, indentString)
}

func (t *FixAxisParDistType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *FixAxisParDistType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("FIX_AXIS_PAR_DIST %s %s %s",
		t.Offset.A2LString(),
		t.Distance.A2LString(),
		t.Numberapo.A2LString()), indentLevel, indentString)
}

func (t *FixAxisParListType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *FixAxisParListType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent("/begin FIX_AXIS_PAR_LIST", indentLevel, indentString)}

	if t.AxisPtsValue != nil {
		for _, axisPtsValue := range t.AxisPtsValue {
			tmpResult = append(tmpResult, indentContent(axisPtsValue.A2LString(), indentLevel+1, indentString))
		}
	}

	tmpResult = append(tmpResult, indentContent("/end FIX_AXIS_PAR_LIST", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *FixAxisParType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *FixAxisParType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("FIX_AXIS_PAR %s %s %s",
		t.Offset.A2LString(),
		t.Shift.A2LString(),
		t.Numberapo.A2LString()), indentLevel, indentString)
}

func (t *FixNoAxisPtsXType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *FixNoAxisPtsXType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("FIX_NO_AXIS_PTS_X %s", t.NumberOfAxisPoints.A2LString()), indentLevel, indentString)
}

func (t *FixNoAxisPtsYType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *FixNoAxisPtsYType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("FIX_NO_AXIS_PTS_Y %s", t.NumberOfAxisPoints.A2LString()), indentLevel, indentString)
}

func (t *FixNoAxisPtsZType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *FixNoAxisPtsZType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("FIX_NO_AXIS_PTS_Z %s", t.NumberOfAxisPoints.A2LString()), indentLevel, indentString)
}

func (t *FncValuesType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *FncValuesType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("FNC_VALUES %s %s %s %s",
		t.Position.A2LString(),
		t.DataType.A2LString(),
		t.IndexMode,
		t.AddressType.A2LString()), indentLevel, indentString)
}

func (t *FormatType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *FormatType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("FORMAT %s", t.FormatString.A2LString()), indentLevel, indentString)
}

func (t *FormulaInvType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *FormulaInvType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("FORMULA_INV %s", t.GX.A2LString()), indentLevel, indentString)
}

func (t *FormulaType) MapChildNodes(node any) {
	switch node.(type) {
	case *FormulaInvType:
		t.FORMULA_INV = node.(*FormulaInvType)
	default:
		panic("not implemented yet...")
	}
}

func (t *FormulaType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin FORMULA %s", t.FX.A2LString()), indentLevel, indentString)}

	if t.FORMULA_INV != nil {
		tmpResult = append(tmpResult, t.FORMULA_INV.MarshalA2L(indentLevel+1, indentString))
	}

	tmpResult = append(tmpResult, indentContent("/end FORMULA", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *FrameMeasurementType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *FrameMeasurementType) MarshalA2L(indentLevel int, indentString string) (result string) {
	result = indentContent("FRAME_MEASUREMENT", indentLevel, indentString)

	if t.Identifier != nil {
		for _, identifier := range t.Identifier {
			result += fmt.Sprintf(" %s", identifier.A2LString())
		}
	}

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

func (t *FrameType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin FRAME %s %s %s %s",
		t.Name.A2LString(),
		t.LongIdentifier.A2LString(),
		t.ScalingUnit.A2LString(),
		t.Rate.A2LString()), indentLevel, indentString)}

	if t.FRAME_MEASUREMENT != nil {
		tmpResult = append(tmpResult, t.FRAME_MEASUREMENT.MarshalA2L(indentLevel+1, indentString))
	}

	if t.IF_DATA != nil {
		for _, ifData := range t.IF_DATA {
			tmpResult = append(tmpResult, ifData.MarshalA2L(indentLevel+1, indentString))
		}
	}

	tmpResult = append(tmpResult, indentContent("/end FRAME", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *FunctionListType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *FunctionListType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent("/begin FUNCTION_LIST", indentLevel, indentString)}

	if t.Name != nil {
		for _, name := range t.Name {
			tmpResult = append(tmpResult, indentContent(name.A2LString(), indentLevel+1, indentString))
		}
	}

	tmpResult = append(tmpResult, indentContent("/end FUNCTION_LIST", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
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

func (t *FunctionType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin FUNCTION %s %s",
		t.Name.A2LString(),
		t.LongIdentifier.A2LString()), indentLevel, indentString)}

	if t.ANNOTATION != nil {
		for _, annotation := range t.ANNOTATION {
			tmpResult = append(tmpResult, annotation.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.DEF_CHARACTERISTIC != nil {
		tmpResult = append(tmpResult, t.DEF_CHARACTERISTIC.MarshalA2L(indentLevel+1, indentString))
	}

	if t.REF_CHARACTERISTIC != nil {
		tmpResult = append(tmpResult, t.REF_CHARACTERISTIC.MarshalA2L(indentLevel+1, indentString))
	}

	if t.IN_MEASUREMENT != nil {
		tmpResult = append(tmpResult, t.IN_MEASUREMENT.MarshalA2L(indentLevel+1, indentString))
	}

	if t.OUT_MEASUREMENT != nil {
		tmpResult = append(tmpResult, t.OUT_MEASUREMENT.MarshalA2L(indentLevel+1, indentString))
	}

	if t.LOC_MEASUREMENT != nil {
		tmpResult = append(tmpResult, t.LOC_MEASUREMENT.MarshalA2L(indentLevel+1, indentString))
	}

	if t.SUB_FUNCTION != nil {
		tmpResult = append(tmpResult, t.SUB_FUNCTION.MarshalA2L(indentLevel+1, indentString))
	}

	if t.FUNCTION_VERSION != nil {
		tmpResult = append(tmpResult, t.FUNCTION_VERSION.MarshalA2L(indentLevel+1, indentString))
	}

	tmpResult = append(tmpResult, indentContent("/end FUNCTION", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *FunctionVersionType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *FunctionVersionType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("FUNCTION_VERSION %s", t.VersionIdentifier.A2LString()), indentLevel, indentString)
}

func (t *GroupType) MapChildNodes(node any) {
	switch node.(type) {
	case *AnnotationType:
		if t.ANNOTATION == nil {
			t.ANNOTATION = make([]*AnnotationType, 0)
		}

		t.ANNOTATION = append(t.ANNOTATION, node.(*AnnotationType))
	case *RootType:
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

func (t *GroupType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin GROUP %s %s",
		t.GroupName.A2LString(),
		t.GroupLongIdentifier.A2LString()), indentLevel, indentString)}

	if t.ANNOTATION != nil {
		for _, annotation := range t.ANNOTATION {
			tmpResult = append(tmpResult, annotation.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.ROOT != nil {
		tmpResult = append(tmpResult, t.ROOT.MarshalA2L(indentLevel+1, indentString))
	}

	if t.REF_CHARACTERISTIC != nil {
		tmpResult = append(tmpResult, t.REF_CHARACTERISTIC.MarshalA2L(indentLevel+1, indentString))
	}

	if t.REF_MEASUREMENT != nil {
		tmpResult = append(tmpResult, t.REF_MEASUREMENT.MarshalA2L(indentLevel+1, indentString))
	}

	if t.FUNCTION_LIST != nil {
		tmpResult = append(tmpResult, t.FUNCTION_LIST.MarshalA2L(indentLevel+1, indentString))
	}

	if t.SUB_GROUP != nil {
		tmpResult = append(tmpResult, t.SUB_GROUP.MarshalA2L(indentLevel+1, indentString))
	}

	tmpResult = append(tmpResult, indentContent("/end GROUP", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *GuardRailsType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *GuardRailsType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent("GUARD_RAILS", indentLevel, indentString)
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

func (t *HeaderType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin HEADER %s", t.Comment.A2LString()), indentLevel, indentString)}

	if t.VERSION != nil {
		tmpResult = append(tmpResult, t.VERSION.MarshalA2L(indentLevel+1, indentString))
	}

	if t.PROJECT_NO != nil {
		tmpResult = append(tmpResult, t.PROJECT_NO.MarshalA2L(indentLevel+1, indentString))
	}

	tmpResult = append(tmpResult, indentContent("/end HEADER", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *IdentificationType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *IdentificationType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("IDENTIFICATION %s %s", t.Position.A2LString(), t.DataType.A2LString()), indentLevel, indentString)
}

func (t *InMeasurementType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *InMeasurementType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent("/begin IN_MEASUREMENT", indentLevel, indentString)}

	if t.Identifier != nil {
		for _, identifier := range t.Identifier {
			tmpResult = append(tmpResult, indentContent(identifier.A2LString(), indentLevel+1, indentString))
		}
	}

	tmpResult = append(tmpResult, indentContent("/end IN_MEASUREMENT", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *LeftShiftType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *LeftShiftType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("LEFT_SHIFT %s", t.BitCount.A2LString()), indentLevel, indentString)
}

func (t *LocMeasurementType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *LocMeasurementType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent("/begin LOC_MEASUREMENT", indentLevel, indentString)}

	if t.Identifier != nil {
		for _, identifier := range t.Identifier {
			tmpResult = append(tmpResult, indentContent(identifier.A2LString(), indentLevel+1, indentString))
		}
	}

	tmpResult = append(tmpResult, indentContent("/end LOC_MEASUREMENT", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *MapListType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *MapListType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent("/begin MAP_LIST", indentLevel, indentString)}

	if t.Name != nil {
		for _, name := range t.Name {
			tmpResult = append(tmpResult, indentContent(name.A2LString(), indentLevel+1, indentString))
		}
	}

	tmpResult = append(tmpResult, indentContent("/end MAP_LIST", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *MatrixDimType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *MatrixDimType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("MATRIX_DIM %s %s %s", t.XDim.A2LString(), t.YDim.A2LString(), t.ZDim.A2LString()), indentLevel, indentString)
}

func (t *MaxGradType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *MaxGradType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("MAX_GRAD %s", t.MaxGradient.A2LString()), indentLevel, indentString)
}

func (t *MaxRefreshType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *MaxRefreshType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("MAX_REFRESH %s %s", t.ScalingUnit.A2LString(), t.Rate.A2LString()), indentLevel, indentString)
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

func (t *MeasurementType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin MEASUREMENT %s %s %s %s %s %s %s %s",
		t.Name.A2LString(),
		t.LongIdentifier.A2LString(),
		t.DataType.A2LString(),
		t.Conversion.A2LString(),
		t.Resolution.A2LString(),
		t.Accuracy.A2LString(),
		t.LowerLimit.A2LString(),
		t.UpperLimit.A2LString()), indentLevel, indentString)}

	if t.DISPLAY_IDENTIFIER != nil {
		tmpResult = append(tmpResult, t.DISPLAY_IDENTIFIER.MarshalA2L(indentLevel+1, indentString))
	}

	if t.READ_WRITE != nil {
		tmpResult = append(tmpResult, t.READ_WRITE.MarshalA2L(indentLevel+1, indentString))
	}

	if t.FORMAT != nil {
		tmpResult = append(tmpResult, t.FORMAT.MarshalA2L(indentLevel+1, indentString))
	}

	if t.ARRAY_SIZE != nil {
		tmpResult = append(tmpResult, t.ARRAY_SIZE.MarshalA2L(indentLevel+1, indentString))
	}

	if t.BIT_MASK != nil {
		tmpResult = append(tmpResult, t.BIT_MASK.MarshalA2L(indentLevel+1, indentString))
	}

	if t.BIT_OPERATION != nil {
		tmpResult = append(tmpResult, t.BIT_OPERATION.MarshalA2L(indentLevel+1, indentString))
	}

	if t.BYTE_ORDER != nil {
		tmpResult = append(tmpResult, t.BYTE_ORDER.MarshalA2L(indentLevel+1, indentString))
	}

	if t.MAX_REFRESH != nil {
		tmpResult = append(tmpResult, t.MAX_REFRESH.MarshalA2L(indentLevel+1, indentString))
	}

	if t.VIRTUAL != nil {
		tmpResult = append(tmpResult, t.VIRTUAL.MarshalA2L(indentLevel+1, indentString))
	}

	if t.FUNCTION_LIST != nil {
		tmpResult = append(tmpResult, t.FUNCTION_LIST.MarshalA2L(indentLevel+1, indentString))
	}

	if t.ECU_ADDRESS != nil {
		tmpResult = append(tmpResult, t.ECU_ADDRESS.MarshalA2L(indentLevel+1, indentString))
	}

	if t.ERROR_MASK != nil {
		tmpResult = append(tmpResult, t.ERROR_MASK.MarshalA2L(indentLevel+1, indentString))
	}

	if t.REF_MEMORY_SEGMENT != nil {
		tmpResult = append(tmpResult, t.REF_MEMORY_SEGMENT.MarshalA2L(indentLevel+1, indentString))
	}

	if t.ANNOTATION != nil {
		for _, annotation := range t.ANNOTATION {
			tmpResult = append(tmpResult, annotation.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.IF_DATA != nil {
		for _, ifData := range t.IF_DATA {
			tmpResult = append(tmpResult, ifData.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.MATRIX_DIM != nil {
		tmpResult = append(tmpResult, t.MATRIX_DIM.MarshalA2L(indentLevel+1, indentString))
	}

	if t.ECU_ADDRESS_EXTENSION != nil {
		tmpResult = append(tmpResult, t.ECU_ADDRESS_EXTENSION.MarshalA2L(indentLevel+1, indentString))
	}

	tmpResult = append(tmpResult, indentContent("/end MEASUREMENT", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
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

func (t *MemoryLayoutType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin MEMORY_LAYOUT %s %s %s %s %s %s %s %s",
		t.PrgType,
		t.Address.A2LString(),
		t.Size.A2LString(),
		t.Offset[0].A2LString(),
		t.Offset[1].A2LString(),
		t.Offset[2].A2LString(),
		t.Offset[3].A2LString(),
		t.Offset[4].A2LString()), indentLevel, indentString)}

	if t.IF_DATA != nil {
		for _, ifData := range t.IF_DATA {
			tmpResult = append(tmpResult, ifData.MarshalA2L(indentLevel+1, indentString))
		}
	}

	tmpResult = append(tmpResult, indentContent("/end MEMORY_LAYOUT", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
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

func (t *MemorySegmentType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin MEMORY_SEGMENT %s %s %s %s %s %s %s %s %s %s %s %s",
		t.Name.A2LString(),
		t.LongIdentifier.A2LString(),
		t.PrgType,
		t.MemoryType,
		t.Attribute,
		t.Address.A2LString(),
		t.Size.A2LString(),
		t.Offset[0].A2LString(),
		t.Offset[1].A2LString(),
		t.Offset[2].A2LString(),
		t.Offset[3].A2LString(),
		t.Offset[4].A2LString()), indentLevel, indentString)}

	if t.IF_DATA != nil {
		for _, ifData := range t.IF_DATA {
			tmpResult = append(tmpResult, ifData.MarshalA2L(indentLevel+1, indentString))
		}
	}

	tmpResult = append(tmpResult, indentContent("/end MEMORY_SEGMENT", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
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

func (t *ModCommonType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin MOD_COMMON %s", t.Comment.A2LString()), indentLevel, indentString)}

	if t.S_REC_LAYOUT != nil {
		tmpResult = append(tmpResult, t.S_REC_LAYOUT.MarshalA2L(indentLevel+1, indentString))
	}

	if t.DEPOSIT != nil {
		tmpResult = append(tmpResult, t.DEPOSIT.MarshalA2L(indentLevel+1, indentString))
	}

	if t.BYTE_ORDER != nil {
		tmpResult = append(tmpResult, t.BYTE_ORDER.MarshalA2L(indentLevel+1, indentString))
	}

	if t.DATA_SIZE != nil {
		tmpResult = append(tmpResult, t.DATA_SIZE.MarshalA2L(indentLevel+1, indentString))
	}

	if t.ALIGNMENT_BYTE != nil {
		tmpResult = append(tmpResult, t.ALIGNMENT_BYTE.MarshalA2L(indentLevel+1, indentString))
	}

	if t.ALIGNMENT_WORD != nil {
		tmpResult = append(tmpResult, t.ALIGNMENT_WORD.MarshalA2L(indentLevel+1, indentString))
	}

	if t.ALIGNMENT_LONG != nil {
		tmpResult = append(tmpResult, t.ALIGNMENT_LONG.MarshalA2L(indentLevel+1, indentString))
	}

	if t.ALIGNMENT_FLOAT32_IEEE != nil {
		tmpResult = append(tmpResult, t.ALIGNMENT_FLOAT32_IEEE.MarshalA2L(indentLevel+1, indentString))
	}

	if t.ALIGNMENT_FLOAT64_IEEE != nil {
		tmpResult = append(tmpResult, t.ALIGNMENT_FLOAT64_IEEE.MarshalA2L(indentLevel+1, indentString))
	}

	tmpResult = append(tmpResult, indentContent("/end MOD_COMMON", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
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

func (t *ModParType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin MOD_PAR %s", t.Comment.A2LString()), indentLevel, indentString)}

	if t.VERSION != nil {
		tmpResult = append(tmpResult, t.VERSION.MarshalA2L(indentLevel+1, indentString))
	}

	if t.ADDR_EPK != nil {
		for _, addrEpk := range t.ADDR_EPK {
			tmpResult = append(tmpResult, addrEpk.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.EPK != nil {
		tmpResult = append(tmpResult, t.EPK.MarshalA2L(indentLevel+1, indentString))
	}

	if t.SUPPLIER != nil {
		tmpResult = append(tmpResult, t.SUPPLIER.MarshalA2L(indentLevel+1, indentString))
	}

	if t.CUSTOMER != nil {
		tmpResult = append(tmpResult, t.CUSTOMER.MarshalA2L(indentLevel+1, indentString))
	}

	if t.CUSTOMER_NO != nil {
		tmpResult = append(tmpResult, t.CUSTOMER_NO.MarshalA2L(indentLevel+1, indentString))
	}

	if t.USER != nil {
		tmpResult = append(tmpResult, t.USER.MarshalA2L(indentLevel+1, indentString))
	}

	if t.PHONE_NO != nil {
		tmpResult = append(tmpResult, t.PHONE_NO.MarshalA2L(indentLevel+1, indentString))
	}

	if t.ECU != nil {
		tmpResult = append(tmpResult, t.ECU.MarshalA2L(indentLevel+1, indentString))
	}

	if t.CPU_TYPE != nil {
		tmpResult = append(tmpResult, t.CPU_TYPE.MarshalA2L(indentLevel+1, indentString))
	}

	if t.NO_OF_INTERFACES != nil {
		tmpResult = append(tmpResult, t.NO_OF_INTERFACES.MarshalA2L(indentLevel+1, indentString))
	}

	if t.ECU_CALIBRATION_OFFSET != nil {
		tmpResult = append(tmpResult, t.ECU_CALIBRATION_OFFSET.MarshalA2L(indentLevel+1, indentString))
	}

	if t.CALIBRATION_METHOD != nil {
		for _, calibrationMethod := range t.CALIBRATION_METHOD {
			tmpResult = append(tmpResult, calibrationMethod.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.MEMORY_LAYOUT != nil {
		for _, memoryLayout := range t.MEMORY_LAYOUT {
			tmpResult = append(tmpResult, memoryLayout.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.MEMORY_SEGMENT != nil {
		for _, memorySegment := range t.MEMORY_SEGMENT {
			tmpResult = append(tmpResult, memorySegment.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.SYSTEM_CONSTANT != nil {
		for _, systemConstant := range t.SYSTEM_CONSTANT {
			tmpResult = append(tmpResult, systemConstant.MarshalA2L(indentLevel+1, indentString))
		}
	}

	tmpResult = append(tmpResult, indentContent("/end MOD_PAR", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
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

func (t *ModuleType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin MODULE %s %s",
		t.Name.A2LString(),
		t.LongIdentifier.A2LString()), indentLevel, indentString)}

	if t.A2ML != nil {
		tmpResult = append(tmpResult, t.A2ML.MarshalA2L(indentLevel+1, indentString))
	}

	if t.MOD_PAR != nil {
		tmpResult = append(tmpResult, t.MOD_PAR.MarshalA2L(indentLevel+1, indentString))
	}

	if t.MOD_COMMON != nil {
		tmpResult = append(tmpResult, t.MOD_COMMON.MarshalA2L(indentLevel+1, indentString))
	}

	if t.IF_DATA != nil {
		for _, ifData := range t.IF_DATA {
			tmpResult = append(tmpResult, ifData.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.CHARACTERISTIC != nil {
		for _, characteristic := range t.CHARACTERISTIC {
			tmpResult = append(tmpResult, characteristic.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.AXIS_PTS != nil {
		for _, axisPts := range t.AXIS_PTS {
			tmpResult = append(tmpResult, axisPts.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.MEASUREMENT != nil {
		for _, measurement := range t.MEASUREMENT {
			tmpResult = append(tmpResult, measurement.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.COMPU_METHOD != nil {
		for _, compuMethod := range t.COMPU_METHOD {
			tmpResult = append(tmpResult, compuMethod.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.COMPU_TAB != nil {
		for _, compuTab := range t.COMPU_TAB {
			tmpResult = append(tmpResult, compuTab.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.COMPU_VTAB != nil {
		for _, compuVTab := range t.COMPU_VTAB {
			tmpResult = append(tmpResult, compuVTab.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.COMPU_VTAB_RANGE != nil {
		for _, compuVTabRange := range t.COMPU_VTAB_RANGE {
			tmpResult = append(tmpResult, compuVTabRange.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.FUNCTION != nil {
		for _, function := range t.FUNCTION {
			tmpResult = append(tmpResult, function.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.GROUP != nil {
		for _, group := range t.GROUP {
			tmpResult = append(tmpResult, group.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.RECORD_LAYOUT != nil {
		for _, recordLayout := range t.RECORD_LAYOUT {
			tmpResult = append(tmpResult, recordLayout.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.VARIANT_CODING != nil {
		tmpResult = append(tmpResult, t.VARIANT_CODING.MarshalA2L(indentLevel+1, indentString))
	}

	if t.FRAME != nil {
		tmpResult = append(tmpResult, t.FRAME.MarshalA2L(indentLevel+1, indentString))
	}

	if t.USER_RIGHTS != nil {
		for _, userRights := range t.USER_RIGHTS {
			tmpResult = append(tmpResult, userRights.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.UNIT != nil {
		for _, unit := range t.UNIT {
			tmpResult = append(tmpResult, unit.MarshalA2L(indentLevel+1, indentString))
		}
	}

	tmpResult = append(tmpResult, indentContent("/end MODULE", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *MonotonyType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *MonotonyType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("MONOTONY %s", t.Monotony), indentLevel, indentString)
}

func (t *NoAxisPtsXType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *NoAxisPtsXType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("NO_AXIS_PTS_X %s %s", t.Position.A2LString(), t.DataType.A2LString()), indentLevel, indentString)
}

func (t *NoAxisPtsYType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *NoAxisPtsYType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("NO_AXIS_PTS_Y %s %s", t.Position.A2LString(), t.DataType.A2LString()), indentLevel, indentString)
}

func (t *NoAxisPtsZType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *NoAxisPtsZType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("NO_AXIS_PTS_Z %s %s", t.Position.A2LString(), t.DataType.A2LString()), indentLevel, indentString)
}

func (t *NoOfInterfacesType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *NoOfInterfacesType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("NO_OF_INTERFACES %s", t.Num.A2LString()), indentLevel, indentString)
}

func (t *NoRescaleXType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *NoRescaleXType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("NO_RESCALE_X %s %s", t.Position.A2LString(), t.DataType.A2LString()), indentLevel, indentString)
}

func (t *NoRescaleYType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *NoRescaleYType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("NO_RESCALE_Y %s %s", t.Position.A2LString(), t.DataType.A2LString()), indentLevel, indentString)
}

func (t *NoRescaleZType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *NoRescaleZType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("NO_RESCALE_Z %s %s", t.Position.A2LString(), t.DataType.A2LString()), indentLevel, indentString)
}

func (t *NumberType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *NumberType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("NUMBER %s", t.Number.A2LString()), indentLevel, indentString)
}

func (t *OffsetXType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *OffsetXType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("OFFSET_X %s %s", t.Position.A2LString(), t.DataType.A2LString()), indentLevel, indentString)
}

func (t *OffsetYType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *OffsetYType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("OFFSET_Y %s %s", t.Position.A2LString(), t.DataType.A2LString()), indentLevel, indentString)
}

func (t *OffsetZType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *OffsetZType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("OFFSET_Z %s %s", t.Position.A2LString(), t.DataType.A2LString()), indentLevel, indentString)
}

func (t *OutMeasurementType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *OutMeasurementType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent("/begin OUT_MEASUREMENT", indentLevel, indentString)}

	if t.Identifier != nil {
		for _, identifier := range t.Identifier {
			tmpResult = append(tmpResult, indentContent(identifier.A2LString(), indentLevel+1, indentString))
		}
	}

	tmpResult = append(tmpResult, indentContent("/end OUT_MEASUREMENT", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *PhoneNoType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *PhoneNoType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("PHONE_NO %s", t.TelNum.A2LString()), indentLevel, indentString)
}

func (t *ProjectNoType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *ProjectNoType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("PROJECT_NO %s", t.ProjectNumber.A2LString()), indentLevel, indentString)
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

func (t *ProjectType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin PROJECT %s %s",
		t.Name.A2LString(),
		t.LongIdentifier.A2LString()), indentLevel, indentString)}

	if t.HEADER != nil {
		tmpResult = append(tmpResult, t.HEADER.MarshalA2L(indentLevel+1, indentString))
	}

	if t.MODULE != nil {
		for _, module := range t.MODULE {
			tmpResult = append(tmpResult, module.MarshalA2L(indentLevel+1, indentString))
		}
	}

	tmpResult = append(tmpResult, indentContent("/end PROJECT", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *ReadOnlyType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *ReadOnlyType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent("READ_ONLY", indentLevel, indentString)
}

func (t *ReadWriteType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *ReadWriteType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent("READ_WRITE", indentLevel, indentString)
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

func (t *RecordLayoutType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin RECORD_LAYOUT %s", t.Name.A2LString()), indentLevel, indentString)}

	if t.FNC_VALUES != nil {
		tmpResult = append(tmpResult, t.FNC_VALUES.MarshalA2L(indentLevel+1, indentString))
	}

	if t.IDENTIFICATION != nil {
		tmpResult = append(tmpResult, t.IDENTIFICATION.MarshalA2L(indentLevel+1, indentString))
	}

	if t.AXIS_PTS_X != nil {
		tmpResult = append(tmpResult, t.AXIS_PTS_X.MarshalA2L(indentLevel+1, indentString))
	}

	if t.AXIS_PTS_Y != nil {
		tmpResult = append(tmpResult, t.AXIS_PTS_Y.MarshalA2L(indentLevel+1, indentString))
	}

	if t.AXIS_PTS_Z != nil {
		tmpResult = append(tmpResult, t.AXIS_PTS_Z.MarshalA2L(indentLevel+1, indentString))
	}

	if t.AXIS_RESCALE_X != nil {
		tmpResult = append(tmpResult, t.AXIS_RESCALE_X.MarshalA2L(indentLevel+1, indentString))
	}

	if t.AXIS_RESCALE_Y != nil {
		tmpResult = append(tmpResult, t.AXIS_RESCALE_Y.MarshalA2L(indentLevel+1, indentString))
	}

	if t.AXIS_RESCALE_Z != nil {
		tmpResult = append(tmpResult, t.AXIS_RESCALE_Z.MarshalA2L(indentLevel+1, indentString))
	}

	if t.NO_AXIS_PTS_X != nil {
		tmpResult = append(tmpResult, t.NO_AXIS_PTS_X.MarshalA2L(indentLevel+1, indentString))
	}

	if t.NO_AXIS_PTS_Y != nil {
		tmpResult = append(tmpResult, t.NO_AXIS_PTS_Y.MarshalA2L(indentLevel+1, indentString))
	}

	if t.NO_AXIS_PTS_Z != nil {
		tmpResult = append(tmpResult, t.NO_AXIS_PTS_Z.MarshalA2L(indentLevel+1, indentString))
	}

	if t.NO_RESCALE_X != nil {
		tmpResult = append(tmpResult, t.NO_RESCALE_X.MarshalA2L(indentLevel+1, indentString))
	}

	if t.NO_RESCALE_Y != nil {
		tmpResult = append(tmpResult, t.NO_RESCALE_Y.MarshalA2L(indentLevel+1, indentString))
	}

	if t.NO_RESCALE_Z != nil {
		tmpResult = append(tmpResult, t.NO_RESCALE_Z.MarshalA2L(indentLevel+1, indentString))
	}

	if t.FIX_NO_AXIS_PTS_X != nil {
		tmpResult = append(tmpResult, t.FIX_NO_AXIS_PTS_X.MarshalA2L(indentLevel+1, indentString))
	}

	if t.FIX_NO_AXIS_PTS_Y != nil {
		tmpResult = append(tmpResult, t.FIX_NO_AXIS_PTS_Y.MarshalA2L(indentLevel+1, indentString))
	}

	if t.FIX_NO_AXIS_PTS_Z != nil {
		tmpResult = append(tmpResult, t.FIX_NO_AXIS_PTS_Z.MarshalA2L(indentLevel+1, indentString))
	}

	if t.SRC_ADDR_X != nil {
		tmpResult = append(tmpResult, t.SRC_ADDR_X.MarshalA2L(indentLevel+1, indentString))
	}

	if t.SRC_ADDR_Y != nil {
		tmpResult = append(tmpResult, t.SRC_ADDR_Y.MarshalA2L(indentLevel+1, indentString))
	}

	if t.SRC_ADDR_Z != nil {
		tmpResult = append(tmpResult, t.SRC_ADDR_Z.MarshalA2L(indentLevel+1, indentString))
	}

	if t.RIP_ADDR_X != nil {
		tmpResult = append(tmpResult, t.RIP_ADDR_X.MarshalA2L(indentLevel+1, indentString))
	}

	if t.RIP_ADDR_Y != nil {
		tmpResult = append(tmpResult, t.RIP_ADDR_Y.MarshalA2L(indentLevel+1, indentString))
	}

	if t.RIP_ADDR_Z != nil {
		tmpResult = append(tmpResult, t.RIP_ADDR_Z.MarshalA2L(indentLevel+1, indentString))
	}

	if t.RIP_ADDR_W != nil {
		tmpResult = append(tmpResult, t.RIP_ADDR_W.MarshalA2L(indentLevel+1, indentString))
	}

	if t.SHIFT_OP_X != nil {
		tmpResult = append(tmpResult, t.SHIFT_OP_X.MarshalA2L(indentLevel+1, indentString))
	}

	if t.SHIFT_OP_Y != nil {
		tmpResult = append(tmpResult, t.SHIFT_OP_Y.MarshalA2L(indentLevel+1, indentString))
	}

	if t.SHIFT_OP_Z != nil {
		tmpResult = append(tmpResult, t.SHIFT_OP_Z.MarshalA2L(indentLevel+1, indentString))
	}

	if t.OFFSET_X != nil {
		tmpResult = append(tmpResult, t.OFFSET_X.MarshalA2L(indentLevel+1, indentString))
	}

	if t.OFFSET_Y != nil {
		tmpResult = append(tmpResult, t.OFFSET_Y.MarshalA2L(indentLevel+1, indentString))
	}

	if t.OFFSET_Z != nil {
		tmpResult = append(tmpResult, t.OFFSET_Z.MarshalA2L(indentLevel+1, indentString))
	}

	if t.DIST_OP_X != nil {
		tmpResult = append(tmpResult, t.DIST_OP_X.MarshalA2L(indentLevel+1, indentString))
	}

	if t.DIST_OP_Y != nil {
		tmpResult = append(tmpResult, t.DIST_OP_Y.MarshalA2L(indentLevel+1, indentString))
	}

	if t.DIST_OP_Z != nil {
		tmpResult = append(tmpResult, t.DIST_OP_Z.MarshalA2L(indentLevel+1, indentString))
	}

	if t.ALIGNMENT_BYTE != nil {
		tmpResult = append(tmpResult, t.ALIGNMENT_BYTE.MarshalA2L(indentLevel+1, indentString))
	}

	if t.ALIGNMENT_WORD != nil {
		tmpResult = append(tmpResult, t.ALIGNMENT_WORD.MarshalA2L(indentLevel+1, indentString))
	}

	if t.ALIGNMENT_LONG != nil {
		tmpResult = append(tmpResult, t.ALIGNMENT_LONG.MarshalA2L(indentLevel+1, indentString))
	}

	if t.ALIGNMENT_FLOAT32_IEEE != nil {
		tmpResult = append(tmpResult, t.ALIGNMENT_FLOAT32_IEEE.MarshalA2L(indentLevel+1, indentString))
	}

	if t.ALIGNMENT_FLOAT64_IEEE != nil {
		tmpResult = append(tmpResult, t.ALIGNMENT_FLOAT64_IEEE.MarshalA2L(indentLevel+1, indentString))
	}

	if t.RESERVED != nil {
		for _, reserved := range t.RESERVED {
			tmpResult = append(tmpResult, reserved.MarshalA2L(indentLevel+1, indentString))
		}
	}

	tmpResult = append(tmpResult, indentContent("/end RECORD_LAYOUT", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *RefCharacteristicType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *RefCharacteristicType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent("/begin REF_CHARACTERISTIC", indentLevel, indentString)}

	if t.Identifier != nil {
		for _, identifier := range t.Identifier {
			tmpResult = append(tmpResult, indentContent(identifier.A2LString(), indentLevel+1, indentString))
		}
	}

	tmpResult = append(tmpResult, indentContent("/end REF_CHARACTERISTIC", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *RefGroupType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *RefGroupType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent("/begin REF_GROUP", indentLevel, indentString)}

	if t.Identifier != nil {
		for _, identifier := range t.Identifier {
			tmpResult = append(tmpResult, indentContent(identifier.A2LString(), indentLevel+1, indentString))
		}
	}

	tmpResult = append(tmpResult, indentContent("/end REF_GROUP", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *RefMeasurementType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *RefMeasurementType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent("/begin REF_MEASUREMENT", indentLevel, indentString)}

	if t.Identifier != nil {
		for _, identifier := range t.Identifier {
			tmpResult = append(tmpResult, indentContent(identifier.A2LString(), indentLevel+1, indentString))
		}
	}

	tmpResult = append(tmpResult, indentContent("/end REF_MEASUREMENT", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *RefMemorySegmentType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *RefMemorySegmentType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("REF_MEMORY_SEGMENT %s", t.Name.A2LString()), indentLevel, indentString)
}

func (t *RefUnitType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *RefUnitType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("REF_UNIT %s", t.Unit.A2LString()), indentLevel, indentString)
}

func (t *ReservedType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *ReservedType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("RESERVED %s %s", t.Position.A2LString(), t.DataSize), indentLevel, indentString)
}

func (t *RightShiftType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *RightShiftType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("RIGHT_SHIFT %s", t.BitCount.A2LString()), indentLevel, indentString)
}

func (t *RipAddrWType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *RipAddrWType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("RIP_ADDR_W %s %s", t.Position.A2LString(), t.DataType.A2LString()), indentLevel, indentString)
}

func (t *RipAddrXType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *RipAddrXType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("RIP_ADDR_X %s %s", t.Position.A2LString(), t.DataType.A2LString()), indentLevel, indentString)
}

func (t *RipAddrYType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *RipAddrYType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("RIP_ADDR_Y %s %s", t.Position.A2LString(), t.DataType.A2LString()), indentLevel, indentString)
}

func (t *RipAddrZType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *RipAddrZType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("RIP_ADDR_Z %s %s", t.Position.A2LString(), t.DataType.A2LString()), indentLevel, indentString)
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

func (t *RootNodeType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := make([]string, 0)

	if t.ASAP2_VERSION != nil {
		tmpResult = append(tmpResult, t.ASAP2_VERSION.MarshalA2L(indentLevel, indentString))
	}

	if t.A2ML_VERSION != nil {
		tmpResult = append(tmpResult, t.A2ML_VERSION.MarshalA2L(indentLevel, indentString))
	}

	if t.PROJECT != nil {
		tmpResult = append(tmpResult, t.PROJECT.MarshalA2L(indentLevel, indentString))
	}

	return strings.Join(tmpResult, "\n")
}

func (t *RootType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *RootType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent("ROOT", indentLevel, indentString)
}

func (t *ShiftOpXType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *ShiftOpXType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("SHIFT_OP_X %s %s", t.Position.A2LString(), t.DataType.A2LString()), indentLevel, indentString)
}

func (t *ShiftOpYType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *ShiftOpYType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("SHIFT_OP_Y %s %s", t.Position.A2LString(), t.DataType.A2LString()), indentLevel, indentString)
}

func (t *ShiftOpZType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *ShiftOpZType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("SHIFT_OP_Z %s %s", t.Position.A2LString(), t.DataType.A2LString()), indentLevel, indentString)
}

func (t *SiExponentsType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *SiExponentsType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("SI_EXPONENTS %s %s %s %s %s %s %s",
		t.Length.A2LString(),
		t.Mass.A2LString(),
		t.Time.A2LString(),
		t.ElectricCurrent.A2LString(),
		t.Temperature.A2LString(),
		t.AmountOfSubstance.A2LString(),
		t.LuminousIntensity.A2LString()), indentLevel, indentString)
}

func (t *SignExtendType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *SignExtendType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent("SIGN_EXTEND", indentLevel, indentString)
}

func (t *SrcAddrXType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *SrcAddrXType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("SRC_ADDR_X %s %s", t.Position.A2LString(), t.DataType.A2LString()), indentLevel, indentString)
}

func (t *SrcAddrYType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *SrcAddrYType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("SRC_ADDR_Y %s %s", t.Position.A2LString(), t.DataType.A2LString()), indentLevel, indentString)
}

func (t *SrcAddrZType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *SrcAddrZType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("SRC_ADDR_Z %s %s", t.Position.A2LString(), t.DataType.A2LString()), indentLevel, indentString)
}

func (t *SRecLayoutType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *SRecLayoutType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("S_REC_LAYOUT %s", t.Name.A2LString()), indentLevel, indentString)
}

func (t *SubFunctionType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *SubFunctionType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent("/begin SUB_FUNCTION", indentLevel, indentString)}

	if t.Identifier != nil {
		for _, identifier := range t.Identifier {
			tmpResult = append(tmpResult, indentContent(identifier.A2LString(), indentLevel+1, indentString))
		}
	}

	tmpResult = append(tmpResult, indentContent("/end SUB_FUNCTION", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *SubGroupType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *SubGroupType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent("/begin SUB_GROUP", indentLevel, indentString)}

	if t.Identifier != nil {
		for _, identifier := range t.Identifier {
			tmpResult = append(tmpResult, indentContent(identifier.A2LString(), indentLevel+1, indentString))
		}
	}

	tmpResult = append(tmpResult, indentContent("/end SUB_GROUP", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *SupplierType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *SupplierType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("SUPPLIER %s", t.Manufacturer.A2LString()), indentLevel, indentString)
}

func (t *SystemConstantType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *SystemConstantType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("SYSTEM_CONSTANT %s %s", t.Name.A2LString(), t.Value.A2LString()), indentLevel, indentString)
}

func (t *UnitConversionType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *UnitConversionType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("UNIT_CONVERSION %s %s", t.Gradient.A2LString(), t.Offset.A2LString()), indentLevel, indentString)
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

func (t *UnitType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin UNIT %s %s %s %s",
		t.Name.A2LString(),
		t.LongIdentifier.A2LString(),
		t.Display.A2LString(),
		t.Type), indentLevel, indentString)}

	if t.SI_EXPONENTS != nil {
		tmpResult = append(tmpResult, t.SI_EXPONENTS.MarshalA2L(indentLevel+1, indentString))
	}

	if t.REF_UNIT != nil {
		tmpResult = append(tmpResult, t.REF_UNIT.MarshalA2L(indentLevel+1, indentString))
	}

	if t.UNIT_CONVERSION != nil {
		tmpResult = append(tmpResult, t.UNIT_CONVERSION.MarshalA2L(indentLevel+1, indentString))
	}

	tmpResult = append(tmpResult, indentContent("/end UNIT", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
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

func (t *UserRightsType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin USER_RIGHTS %s", t.UserLevelId.A2LString()), indentLevel, indentString)}

	if t.REF_GROUP != nil {
		for _, refGroup := range t.REF_GROUP {
			tmpResult = append(tmpResult, refGroup.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.READ_ONLY != nil {
		tmpResult = append(tmpResult, t.READ_ONLY.MarshalA2L(indentLevel+1, indentString))
	}

	tmpResult = append(tmpResult, indentContent("/end USER_RIGHTS", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *UserType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *UserType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("USER %s", t.UserName.A2LString()), indentLevel, indentString)
}

func (t *VarAddressType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *VarAddressType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent("/begin VAR_ADDRESS", indentLevel, indentString)}

	if t.Address != nil {
		for _, address := range t.Address {
			tmpResult = append(tmpResult, indentContent(address.A2LString(), indentLevel+1, indentString))
		}
	}

	tmpResult = append(tmpResult, indentContent("/end VAR_ADDRESS", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *VarCharacteristicType) MapChildNodes(node any) {
	switch node.(type) {
	case *VarAddressType:
		t.VAR_ADDRESS = node.(*VarAddressType)
	default:
		panic("not implemented yet...")
	}
}

func (t *VarCharacteristicType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin VAR_CHARACTERISTIC %s", t.Name.A2LString()), indentLevel, indentString)}

	if t.CriterionName != nil {
		for _, criterionName := range t.CriterionName {
			tmpResult[0] += fmt.Sprintf(" %s", criterionName.A2LString())
		}
	}

	if t.VAR_ADDRESS != nil {
		tmpResult = append(tmpResult, t.VAR_ADDRESS.MarshalA2L(indentLevel+1, indentString))
	}

	tmpResult = append(tmpResult, indentContent("/end VAR_CHARACTERISTIC", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
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

func (t *VarCriterionType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin VAR_CRITERION %s %s",
		t.Name.A2LString(),
		t.LongIdentifier.A2LString()), indentLevel, indentString)}

	if t.Value != nil {
		for _, value := range t.Value {
			tmpResult[0] += fmt.Sprintf(" %s", value.A2LString())
		}
	}

	if t.VAR_MEASUREMENT != nil {
		tmpResult = append(tmpResult, t.VAR_MEASUREMENT.MarshalA2L(indentLevel+1, indentString))
	}

	if t.VAR_SELECTION_CHARACTERISTIC != nil {
		tmpResult = append(tmpResult, t.VAR_SELECTION_CHARACTERISTIC.MarshalA2L(indentLevel+1, indentString))
	}

	tmpResult = append(tmpResult, indentContent("/end VAR_CRITERION", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *VarForbiddenCombType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *VarForbiddenCombType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent("/begin VAR_FORBIDDEN_COMB", indentLevel, indentString)}

	if t.CriterionNameCriterionValue != nil {
		for _, criterionNameCriterionValue := range t.CriterionNameCriterionValue {
			tmpResult[0] += fmt.Sprintf(" %s %s",
				criterionNameCriterionValue.CriterionName.A2LString(),
				criterionNameCriterionValue.CriterionValue.A2LString())
		}
	}

	tmpResult = append(tmpResult, indentContent("/end VAR_FORBIDDEN_COMB", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
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

func (t *VariantCodingType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent("/begin VARIANT_CODING", indentLevel, indentString)}

	if t.VAR_SEPARATOR != nil {
		tmpResult = append(tmpResult, t.VAR_SEPARATOR.MarshalA2L(indentLevel+1, indentString))
	}

	if t.VAR_NAMING != nil {
		tmpResult = append(tmpResult, t.VAR_NAMING.MarshalA2L(indentLevel+1, indentString))
	}

	if t.VAR_CRITERION != nil {
		for _, varCriterion := range t.VAR_CRITERION {
			tmpResult = append(tmpResult, varCriterion.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.VAR_FORBIDDEN_COMB != nil {
		for _, varForbiddenComb := range t.VAR_FORBIDDEN_COMB {
			tmpResult = append(tmpResult, varForbiddenComb.MarshalA2L(indentLevel+1, indentString))
		}
	}

	if t.VAR_CHARACTERISTIC != nil {
		for _, varCharacteristic := range t.VAR_CHARACTERISTIC {
			tmpResult = append(tmpResult, varCharacteristic.MarshalA2L(indentLevel+1, indentString))
		}
	}

	tmpResult = append(tmpResult, indentContent("/end VARIANT_CODING", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *VarMeasurementType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *VarMeasurementType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("VAR_MEASUREMENT %s", t.Name.A2LString()), indentLevel, indentString)
}

func (t *VarNamingType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *VarNamingType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("VAR_NAMING %s", t.Tag), indentLevel, indentString)
}

func (t *VarSelectionCharacteristicType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *VarSelectionCharacteristicType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("VAR_SELECTION_CHARACTERISTIC %s", t.Name.A2LString()), indentLevel, indentString)
}

func (t *VarSeparatorType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *VarSeparatorType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("VAR_SEPARATOR %s", t.Separator.A2LString()), indentLevel, indentString)
}

func (t *VersionType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *VersionType) MarshalA2L(indentLevel int, indentString string) (result string) {
	return indentContent(fmt.Sprintf("VERSION %s", t.VersionIdentifier.A2LString()), indentLevel, indentString)
}

func (t *VirtualCharacteristicType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *VirtualCharacteristicType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent(fmt.Sprintf("/begin VIRTUAL_CHARACTERISTIC %s", t.Formula.A2LString()), indentLevel, indentString)}

	if t.Characteristic != nil {
		for _, characteristic := range t.Characteristic {
			tmpResult = append(tmpResult, indentContent(characteristic.A2LString(), indentLevel+1, indentString))
		}
	}

	tmpResult = append(tmpResult, indentContent("/end VIRTUAL_CHARACTERISTIC", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}

func (t *VirtualType) MapChildNodes(_ any) {
	panic("leaf node")
}

func (t *VirtualType) MarshalA2L(indentLevel int, indentString string) (result string) {
	tmpResult := []string{indentContent("/begin VIRTUAL", indentLevel, indentString)}

	if t.MeasuringChannel != nil {
		for _, measuringChannel := range t.MeasuringChannel {
			tmpResult[0] += fmt.Sprintf(" %s", measuringChannel.A2LString())
		}
	}

	tmpResult = append(tmpResult, indentContent("/end VIRTUAL", indentLevel, indentString))

	return strings.Join(tmpResult, "\n")
}
