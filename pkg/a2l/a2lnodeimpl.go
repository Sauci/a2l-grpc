package a2l

import (
	"github.com/sauci/a2l-grpc/pkg/a2l/parser"
)

func (n *Listener) EnterA2mlVersion(ctx *parser.A2mlVersionContext) {
	n.Push(&A2MLVersionType{
		VersionNo: a2lIntToIntType(ctx.GetVersionNo()),
		UpgradeNo: a2lIntToIntType(ctx.GetUpgradeNo()),
	})
}

func (n *Listener) ExitA2mlVersion(_ *parser.A2mlVersionContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterAddrEpk(ctx *parser.AddrEpkContext) {
	n.Push(&AddrEpkType{Address: a2lLongToLongType(ctx.GetAddress())})
}

func (n *Listener) ExitAddrEpk(_ *parser.AddrEpkContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterAlignmentByte(ctx *parser.AlignmentByteContext) {
	n.Push(&AlignmentByteType{AlignmentBorder: a2lIntToIntType(ctx.GetAlignmentBorder())})
}

func (n *Listener) ExitAlignmentByte(_ *parser.AlignmentByteContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterAlignmentFloat32Ieee(ctx *parser.AlignmentFloat32IeeeContext) {
	n.Push(&AlignmentFloat32IeeeType{AlignmentBorder: a2lIntToIntType(ctx.GetAlignmentBorder())})
}

func (n *Listener) ExitAlignmentFloat32Ieee(_ *parser.AlignmentFloat32IeeeContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterAlignmentFloat64Ieee(ctx *parser.AlignmentFloat64IeeeContext) {
	n.Push(&AlignmentFloat64IeeeType{AlignmentBorder: a2lIntToIntType(ctx.GetAlignmentBorder())})
}

func (n *Listener) ExitAlignmentFloat64Ieee(_ *parser.AlignmentFloat64IeeeContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterAlignmentLong(ctx *parser.AlignmentLongContext) {
	n.Push(&AlignmentLongType{AlignmentBorder: a2lIntToIntType(ctx.GetAlignmentBorder())})
}

func (n *Listener) ExitAlignmentLong(_ *parser.AlignmentLongContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterAlignmentWord(ctx *parser.AlignmentWordContext) {
	n.Push(&AlignmentWordType{AlignmentBorder: a2lIntToIntType(ctx.GetAlignmentBorder())})
}

func (n *Listener) ExitAlignmentWord(_ *parser.AlignmentWordContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterAnnotationLabel(ctx *parser.AnnotationLabelContext) {
	n.Push(&AnnotationLabelType{Label: a2lStringToStringType(ctx.GetLabel())})
}

func (n *Listener) ExitAnnotationLabel(_ *parser.AnnotationLabelContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterAnnotationOrigin(ctx *parser.AnnotationOriginContext) {
	n.Push(&AnnotationOriginType{Origin: a2lStringToStringType(ctx.GetOrigin())})
}

func (n *Listener) ExitAnnotationOrigin(_ *parser.AnnotationOriginContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterAnnotationText(ctx *parser.AnnotationTextContext) {
	e := &AnnotationTextType{AnnotationText: make([]*StringType, 0)}

	for _, annotation := range ctx.GetAnnotation_text() {
		e.AnnotationText = append(e.AnnotationText, a2lStringToStringType(annotation))
	}

	n.Push(e)
}

func (n *Listener) ExitAnnotationText(_ *parser.AnnotationTextContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterAnnotation(_ *parser.AnnotationContext) {
	n.Push(&AnnotationType{})
}

func (n *Listener) ExitAnnotation(_ *parser.AnnotationContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterArraySize(ctx *parser.ArraySizeContext) {
	n.Push(&ArraySizeType{Number: a2lIntToIntType(ctx.GetNumber_())})
}

func (n *Listener) ExitArraySize(_ *parser.ArraySizeContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterAsap2Version(ctx *parser.Asap2VersionContext) {
	n.Push(&ASAP2VersionType{
		VersionNo: a2lIntToIntType(ctx.GetVersionNo()),
		UpgradeNo: a2lIntToIntType(ctx.GetUpgradeNo()),
	})
}

func (n *Listener) ExitAsap2Version(_ *parser.Asap2VersionContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterAxisDescr(ctx *parser.AxisDescrContext) {
	n.Push(&AxisDescrType{
		Attribute:     ctx.GetAttribute().GetText(),
		InputQuantity: identifierToIdentType(ctx.GetInputQuantity()),
		Conversion:    identifierToIdentType(ctx.GetConversion()),
		MaxAxisPoints: a2lIntToIntType(ctx.GetMaxAxisPoints()),
		LowerLimit:    floatToFloatType(ctx.GetLowerLimit()),
		UpperLimit:    floatToFloatType(ctx.GetUpperLimit()),
	})
}

func (n *Listener) ExitAxisDescr(_ *parser.AxisDescrContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterAxisPtsRef(ctx *parser.AxisPtsRefContext) {
	n.Push(&AxisPtsRefType{AxisPoints: identifierToIdentType(ctx.GetAxisPoints())})
}

func (n *Listener) ExitAxisPtsRef(_ *parser.AxisPtsRefContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterAxisPts(ctx *parser.AxisPtsContext) {
	n.Push(&AxisPtsType{
		Name:           identifierToIdentType(ctx.GetName()),
		LongIdentifier: a2lStringToStringType(ctx.GetLongIdentifier()),
		Address:        a2lLongToLongType(ctx.GetAddress()),
		InputQuantity:  identifierToIdentType(ctx.GetInputQuantity()),
		DepositR:       identifierToIdentType(ctx.GetVDeposit()),
		MaxDiff:        floatToFloatType(ctx.GetMaxDiff()),
		Conversion:     identifierToIdentType(ctx.GetConversion()),
		MaxAxisPoints:  a2lIntToIntType(ctx.GetMaxAxisPoints()),
		LowerLimit:     floatToFloatType(ctx.GetLowerLimit()),
		UpperLimit:     floatToFloatType(ctx.GetUpperLimit()),
	})
}

func (n *Listener) ExitAxisPts(_ *parser.AxisPtsContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterAxisPtsX(ctx *parser.AxisPtsXContext) {
	n.Push(&AxisPtsXType{
		Position:   a2lIntToIntType(ctx.GetPosition()),
		DataType:   &DataTypeType{Value: ctx.GetDatatype().GetText()},
		IndexIncr:  &IndexOrderType{Value: ctx.GetIndexIncr().GetText()},
		Addressing: &AddrTypeType{Value: ctx.GetAddressing().GetText()},
	})
}

func (n *Listener) ExitAxisPtsX(_ *parser.AxisPtsXContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterAxisPtsY(ctx *parser.AxisPtsYContext) {
	n.Push(&AxisPtsYType{
		Position:   a2lIntToIntType(ctx.GetPosition()),
		DataType:   &DataTypeType{Value: ctx.GetDatatype().GetText()},
		IndexIncr:  &IndexOrderType{Value: ctx.GetIndexIncr().GetText()},
		Addressing: &AddrTypeType{Value: ctx.GetAddressing().GetText()},
	})
}

func (n *Listener) ExitAxisPtsY(_ *parser.AxisPtsYContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterAxisPtsZ(ctx *parser.AxisPtsZContext) {
	n.Push(&AxisPtsZType{
		Position:   a2lIntToIntType(ctx.GetPosition()),
		DataType:   &DataTypeType{Value: ctx.GetDatatype().GetText()},
		IndexIncr:  &IndexOrderType{Value: ctx.GetIndexIncr().GetText()},
		Addressing: &AddrTypeType{Value: ctx.GetAddressing().GetText()},
	})
}

func (n *Listener) ExitAxisPtsZ(_ *parser.AxisPtsZContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterAxisRescaleX(ctx *parser.AxisRescaleXContext) {
	n.Push(&AxisRescaleXType{
		Position:                a2lIntToIntType(ctx.GetPosition()),
		DataType:                &DataTypeType{Value: ctx.GetDatatype().GetText()},
		MaxNumberOfRescalePairs: a2lIntToIntType(ctx.GetMaxNumberOfRescalePairs()),
		IndexIncr:               &IndexOrderType{Value: ctx.GetIndexIncr().GetText()},
		Addressing:              &AddrTypeType{Value: ctx.GetAddressing().GetText()},
	})
}

func (n *Listener) ExitAxisRescaleX(_ *parser.AxisRescaleXContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterAxisRescaleY(ctx *parser.AxisRescaleYContext) {
	n.Push(&AxisRescaleYType{
		Position:                a2lIntToIntType(ctx.GetPosition()),
		DataType:                &DataTypeType{Value: ctx.GetDatatype().GetText()},
		MaxNumberOfRescalePairs: a2lIntToIntType(ctx.GetMaxNumberOfRescalePairs()),
		IndexIncr:               &IndexOrderType{Value: ctx.GetIndexIncr().GetText()},
		Addressing:              &AddrTypeType{Value: ctx.GetAddressing().GetText()},
	})
}

func (n *Listener) ExitAxisRescaleY(_ *parser.AxisRescaleYContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterAxisRescaleZ(ctx *parser.AxisRescaleZContext) {
	n.Push(&AxisRescaleZType{
		Position:                a2lIntToIntType(ctx.GetPosition()),
		DataType:                &DataTypeType{Value: ctx.GetDatatype().GetText()},
		MaxNumberOfRescalePairs: a2lIntToIntType(ctx.GetMaxNumberOfRescalePairs()),
		IndexIncr:               &IndexOrderType{Value: ctx.GetIndexIncr().GetText()},
		Addressing:              &AddrTypeType{Value: ctx.GetAddressing().GetText()},
	})
}

func (n *Listener) ExitAxisRescaleZ(_ *parser.AxisRescaleZContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterBitMask(ctx *parser.BitMaskContext) {
	n.Push(&BitMaskType{Mask: a2lLongToLongType(ctx.GetMask())})
}

func (n *Listener) ExitBitMask(_ *parser.BitMaskContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterBitOperation(_ *parser.BitOperationContext) {
	n.Push(&BitOperationType{})
}

func (n *Listener) ExitBitOperation(_ *parser.BitOperationContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterByteOrder(ctx *parser.ByteOrderContext) {
	n.Push(&ByteOrderType{ByteOrder: ctx.GetByteOrder_().GetText()})
}

func (n *Listener) ExitByteOrder(_ *parser.ByteOrderContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterCalibrationAccess(ctx *parser.CalibrationAccessContext) {
	n.Push(&CalibrationAccessType{Type: ctx.GetType_().GetText()})
}

func (n *Listener) ExitCalibrationAccess(_ *parser.CalibrationAccessContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterCalibrationHandle(ctx *parser.CalibrationHandleContext) {
	e := &CalibrationHandleType{Handle: make([]*LongType, 0)}

	for _, handle := range ctx.GetHandle() {
		e.Handle = append(e.Handle, a2lLongToLongType(handle))
	}

	n.Push(e)
}

func (n *Listener) ExitCalibrationHandle(_ *parser.CalibrationHandleContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterCalibrationMethod(ctx *parser.CalibrationMethodContext) {
	n.Push(&CalibrationMethodType{
		Method:  a2lStringToStringType(ctx.GetMethod()),
		Version: a2lLongToLongType(ctx.GetVersion_()),
	})
}

func (n *Listener) ExitCalibrationMethod(_ *parser.CalibrationMethodContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterCharacteristic(ctx *parser.CharacteristicContext) {
	n.Push(&CharacteristicType{
		Name:           identifierToIdentType(ctx.GetName()),
		LongIdentifier: a2lStringToStringType(ctx.GetLongIdentifier()),
		Type:           ctx.GetType_().GetText(),
		Address:        a2lLongToLongType(ctx.GetAddress()),
		Deposit:        identifierToIdentType(ctx.GetDeposit_()),
		MaxDiff:        floatToFloatType(ctx.GetMaxDiff()),
		Conversion:     identifierToIdentType(ctx.GetConversion()),
		LowerLimit:     floatToFloatType(ctx.GetLowerLimit()),
		UpperLimit:     floatToFloatType(ctx.GetUpperLimit()),
	})
}

func (n *Listener) ExitCharacteristic(_ *parser.CharacteristicContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterCoeffs(ctx *parser.CoeffsContext) {
	n.Push(&CoeffsType{
		A: floatToFloatType(ctx.GetA()),
		B: floatToFloatType(ctx.GetB()),
		C: floatToFloatType(ctx.GetC()),
		D: floatToFloatType(ctx.GetD()),
		E: floatToFloatType(ctx.GetE()),
		F: floatToFloatType(ctx.GetF()),
	})
}

func (n *Listener) ExitCoeffs(_ *parser.CoeffsContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterComparisonQuantity(ctx *parser.ComparisonQuantityContext) {
	n.Push(&ComparisonQuantityType{Name: identifierToIdentType(ctx.GetName())})
}

func (n *Listener) ExitComparisonQuantity(_ *parser.ComparisonQuantityContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterCompuMethod(ctx *parser.CompuMethodContext) {
	n.Push(&CompuMethodType{
		Name:           identifierToIdentType(ctx.GetName()),
		LongIdentifier: a2lStringToStringType(ctx.GetLongIdentifier()),
		ConversionType: ctx.GetConversionType().GetText(),
		Format:         a2lStringToStringType(ctx.GetVFormat()),
		Unit:           a2lStringToStringType(ctx.GetVUnit()),
	})
}

func (n *Listener) ExitCompuMethod(_ *parser.CompuMethodContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterCompuTabRef(ctx *parser.CompuTabRefContext) {
	n.Push(&CompuTabRefType{ConversionTable: identifierToIdentType(ctx.GetConversionTable())})
}

func (n *Listener) ExitCompuTabRef(_ *parser.CompuTabRefContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterCompuTab(ctx *parser.CompuTabContext) {
	e := &CompuTabType{Name: identifierToIdentType(ctx.GetName()),
		LongIdentifier:   a2lStringToStringType(ctx.GetLongIdentifier()),
		ConversionType:   ctx.GetConversionType().GetText(),
		NumberValuePairs: a2lIntToIntType(ctx.GetNumberValuePairs()),
		InValOutVal:      make([]*CompuTabType_InValOutValType, 0)}

	for i := 0; i < len(ctx.GetInVal()); i++ {
		e.InValOutVal = append(e.InValOutVal, &CompuTabType_InValOutValType{
			InVal:  floatToFloatType(ctx.GetInVal()[i]),
			OutVal: floatToFloatType(ctx.GetOutVal()[i])})
	}

	n.Push(e)
}

func (n *Listener) ExitCompuTab(_ *parser.CompuTabContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterCompuVTabRange(ctx *parser.CompuVTabRangeContext) {
	e := &CompuVTabRangeType{Name: identifierToIdentType(ctx.GetName()),
		LongIdentifier:         a2lStringToStringType(ctx.GetLongIdentifier()),
		NumberOfValuesTriples:  a2lIntToIntType(ctx.GetNumberValueTriples()),
		InValMinInValMaxOutVal: make([]*CompuVTabRangeType_InValMinInValMaxOutValType, 0)}

	for i := 0; i < len(ctx.GetInValMin()); i++ {
		e.InValMinInValMaxOutVal = append(e.InValMinInValMaxOutVal, &CompuVTabRangeType_InValMinInValMaxOutValType{
			InValMin: floatToFloatType(ctx.GetInValMin()[i]),
			InValMax: floatToFloatType(ctx.GetInValMax()[i]),
			OutVal:   a2lStringToStringType(ctx.GetOutVal()[i]),
		})
	}

	n.Push(e)
}

func (n *Listener) ExitCompuVTabRange(_ *parser.CompuVTabRangeContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterCompuVTab(ctx *parser.CompuVTabContext) {
	e := &CompuVTabType{Name: identifierToIdentType(ctx.GetName()),
		LongIdentifier:   a2lStringToStringType(ctx.GetLongIdentifier()),
		ConversionType:   ctx.GetConversionType().GetText(),
		NumberValuePairs: a2lIntToIntType(ctx.GetNumberValuePairs()),
		InValOutVal:      make([]*CompuVTabType_InValOutValType, 0)}

	for i := 0; i < len(ctx.GetInVal()); i++ {
		e.InValOutVal = append(e.InValOutVal, &CompuVTabType_InValOutValType{
			InVal:  floatToFloatType(ctx.GetInVal()[i]),
			OutVal: a2lStringToStringType(ctx.GetOutVal()[i])})
	}

	n.Push(e)
}

func (n *Listener) ExitCompuVTab(_ *parser.CompuVTabContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterCpuType(ctx *parser.CpuTypeContext) {
	n.Push(&CpuTypeType{Cpu: a2lStringToStringType(ctx.GetCPU())})
}

func (n *Listener) ExitCpuType(_ *parser.CpuTypeContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterCurveAxisRef(ctx *parser.CurveAxisRefContext) {
	n.Push(&CurveAxisRefType{CurveAxis: identifierToIdentType(ctx.GetCurveAxis())})
}

func (n *Listener) ExitCurveAxisRef(_ *parser.CurveAxisRefContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterCustomerNo(ctx *parser.CustomerNoContext) {
	n.Push(&CustomerNoType{Number: a2lStringToStringType(ctx.GetNumber_())})
}

func (n *Listener) ExitCustomerNo(_ *parser.CustomerNoContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterCustomer(ctx *parser.CustomerContext) {
	n.Push(&CustomerType{Customer: a2lStringToStringType(ctx.GetCustomer_())})
}

func (n *Listener) ExitCustomer(_ *parser.CustomerContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterDataSize(ctx *parser.DataSizeContext) {
	n.Push(&DataSizeType{Value: ctx.GetText()})
}

func (n *Listener) ExitDataSize(_ *parser.DataSizeContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterDefaultValue(ctx *parser.DefaultValueContext) {
	n.Push(&DefaultValueType{DisplayString: a2lStringToStringType(ctx.GetDisplay_string())})
}

func (n *Listener) ExitDefaultValue(_ *parser.DefaultValueContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterDefCharacteristic(ctx *parser.DefCharacteristicContext) {
	e := &DefCharacteristicType{Identifier: make([]*IdentType, 0)}

	for _, identifier := range ctx.GetIdentifier() {
		e.Identifier = append(e.Identifier, identifierToIdentType(identifier))
	}

	n.Push(e)
}

func (n *Listener) ExitDefCharacteristic(_ *parser.DefCharacteristicContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterDependentCharacteristic(ctx *parser.DependentCharacteristicContext) {
	e := &DependentCharacteristicType{Formula: a2lStringToStringType(ctx.GetFormula_()), Characteristic: make([]*IdentType, 0)}

	for _, characteristic := range ctx.GetCharacteristic_() {
		e.Characteristic = append(e.Characteristic, identifierToIdentType(characteristic))
	}

	n.Push(e)
}

func (n *Listener) ExitDependentCharacteristic(_ *parser.DependentCharacteristicContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterDeposit(ctx *parser.DepositContext) {
	n.Push(&DepositType{Mode: ctx.GetMode_().GetText()})
}

func (n *Listener) ExitDeposit(_ *parser.DepositContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterDisplayIdentifier(ctx *parser.DisplayIdentifierContext) {
	n.Push(&DisplayIdentifierType{DisplayName: identifierToIdentType(ctx.GetDisplay_name())})
}

func (n *Listener) ExitDisplayIdentifier(_ *parser.DisplayIdentifierContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterDistOpX(ctx *parser.DistOpXContext) {
	n.Push(&DistOpXType{
		Position: a2lIntToIntType(ctx.GetPosition()),
		DataType: &DataTypeType{Value: ctx.GetDatatype().GetText()},
	})
}

func (n *Listener) ExitDistOpX(_ *parser.DistOpXContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterDistOpY(ctx *parser.DistOpYContext) {
	n.Push(&DistOpYType{
		Position: a2lIntToIntType(ctx.GetPosition()),
		DataType: &DataTypeType{Value: ctx.GetDatatype().GetText()},
	})
}

func (n *Listener) ExitDistOpY(_ *parser.DistOpYContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterDistOpZ(ctx *parser.DistOpZContext) {
	n.Push(&DistOpZType{
		Position: a2lIntToIntType(ctx.GetPosition()),
		DataType: &DataTypeType{Value: ctx.GetDatatype().GetText()},
	})
}

func (n *Listener) ExitDistOpZ(_ *parser.DistOpZContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterEcuAddressExtension(ctx *parser.EcuAddressExtensionContext) {
	n.Push(&EcuAddressExtensionType{Extension: a2lIntToIntType(ctx.IntegerValue())})
}

func (n *Listener) ExitEcuAddressExtension(_ *parser.EcuAddressExtensionContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterEcuAddress(ctx *parser.EcuAddressContext) {
	n.Push(&EcuAddressType{Address: a2lLongToLongType(ctx.GetAddress())})
}

func (n *Listener) ExitEcuAddress(_ *parser.EcuAddressContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterEcuCalibrationOffset(ctx *parser.EcuCalibrationOffsetContext) {
	n.Push(&EcuCalibrationOffsetType{Offset: a2lLongToLongType(ctx.GetOffset())})
}

func (n *Listener) ExitEcuCalibrationOffset(_ *parser.EcuCalibrationOffsetContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterEcu(ctx *parser.EcuContext) {
	n.Push(&EcuType{ControlUnit: a2lStringToStringType(ctx.GetControlUnit())})
}

func (n *Listener) ExitEcu(_ *parser.EcuContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterEpk(ctx *parser.EpkContext) {
	n.Push(&EpkType{Identifier: a2lStringToStringType(ctx.GetIdentifier())})
}

func (n *Listener) ExitEpk(_ *parser.EpkContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterErrorMask(ctx *parser.ErrorMaskContext) {
	n.Push(&ErrorMaskType{Mask: a2lLongToLongType(ctx.GetMask())})
}

func (n *Listener) ExitErrorMask(_ *parser.ErrorMaskContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterExtendedLimits(ctx *parser.ExtendedLimitsContext) {
	n.Push(&ExtendedLimitsType{
		LowerLimit: floatToFloatType(ctx.GetLowerLimit()),
		UpperLimit: floatToFloatType(ctx.GetUpperLimit()),
	})
}

func (n *Listener) ExitExtendedLimits(_ *parser.ExtendedLimitsContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterFixAxisParDist(ctx *parser.FixAxisParDistContext) {
	n.Push(&FixAxisParDistType{
		Offset:    a2lIntToIntType(ctx.GetOffset()),
		Distance:  a2lIntToIntType(ctx.GetDistance()),
		Numberapo: a2lIntToIntType(ctx.GetNumberapo()),
	})
}

func (n *Listener) ExitFixAxisParDist(_ *parser.FixAxisParDistContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterFixAxisParList(ctx *parser.FixAxisParListContext) {
	e := &FixAxisParListType{AxisPtsValue: make([]*FloatType, 0)}

	for _, axisPtsValue := range ctx.GetAxisPts_Value() {
		e.AxisPtsValue = append(e.AxisPtsValue, floatToFloatType(axisPtsValue))
	}

	n.Push(e)
}

func (n *Listener) ExitFixAxisParList(_ *parser.FixAxisParListContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterFixAxisPar(ctx *parser.FixAxisParContext) {
	n.Push(&FixAxisParType{
		Offset:    a2lIntToIntType(ctx.GetOffset()),
		Shift:     a2lIntToIntType(ctx.GetShift()),
		Numberapo: a2lIntToIntType(ctx.GetNumberapo()),
	})
}

func (n *Listener) ExitFixAxisPar(_ *parser.FixAxisParContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterFixNoAxisPtsX(ctx *parser.FixNoAxisPtsXContext) {
	n.Push(&FixNoAxisPtsXType{NumberOfAxisPoints: a2lIntToIntType(ctx.GetNumberOfAxisPoints())})
}

func (n *Listener) ExitFixNoAxisPtsX(_ *parser.FixNoAxisPtsXContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterFixNoAxisPtsY(ctx *parser.FixNoAxisPtsYContext) {
	n.Push(&FixNoAxisPtsYType{NumberOfAxisPoints: a2lIntToIntType(ctx.GetNumberOfAxisPoints())})
}

func (n *Listener) ExitFixNoAxisPtsY(_ *parser.FixNoAxisPtsYContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterFixNoAxisPtsZ(ctx *parser.FixNoAxisPtsZContext) {
	n.Push(&FixNoAxisPtsZType{NumberOfAxisPoints: a2lIntToIntType(ctx.GetNumberOfAxisPoints())})
}

func (n *Listener) ExitFixNoAxisPtsZ(_ *parser.FixNoAxisPtsZContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterFncValues(ctx *parser.FncValuesContext) {
	n.Push(&FncValuesType{
		Position:    a2lIntToIntType(ctx.GetPosition()),
		DataType:    &DataTypeType{Value: ctx.GetDatatype().GetText()},
		IndexMode:   ctx.GetIndexMode().GetText(),
		AddressType: &AddrTypeType{Value: ctx.GetAddresstype().GetText()},
	})
}

func (n *Listener) ExitFncValues(_ *parser.FncValuesContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterFormat(ctx *parser.FormatContext) {
	n.Push(&FormatType{FormatString: a2lStringToStringType(ctx.GetFormatString())})
}

func (n *Listener) ExitFormat(_ *parser.FormatContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterFormulaInv(ctx *parser.FormulaInvContext) {
	n.Push(&FormulaInvType{GX: a2lStringToStringType(ctx.GetG_x())})
}

func (n *Listener) ExitFormulaInv(_ *parser.FormulaInvContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterFormula(ctx *parser.FormulaContext) {
	n.Push(&FormulaType{FX: a2lStringToStringType(ctx.GetF_x())})
}

func (n *Listener) ExitFormula(_ *parser.FormulaContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterFrameMeasurement(ctx *parser.FrameMeasurementContext) {
	e := &FrameMeasurementType{Identifier: make([]*IdentType, 0)}

	for _, identifier := range ctx.GetIdentifier() {
		e.Identifier = append(e.Identifier, identifierToIdentType(identifier))
	}

	n.Push(e)
}

func (n *Listener) ExitFrameMeasurement(_ *parser.FrameMeasurementContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterFrame(ctx *parser.FrameContext) {
	n.Push(&FrameType{
		Name:           identifierToIdentType(ctx.GetName()),
		LongIdentifier: a2lStringToStringType(ctx.GetLongIdentifier()),
		ScalingUnit:    a2lIntToIntType(ctx.GetScalingUnit()),
		Rate:           a2lLongToLongType(ctx.GetRate()),
	})
}

func (n *Listener) ExitFrame(_ *parser.FrameContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterFunctionList(ctx *parser.FunctionListContext) {
	e := &FunctionListType{Name: make([]*IdentType, 0)}

	for _, name := range ctx.GetName() {
		e.Name = append(e.Name, identifierToIdentType(name))
	}

	n.Push(e)
}

func (n *Listener) ExitFunctionList(_ *parser.FunctionListContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterFunction(ctx *parser.FunctionContext) {
	n.Push(&FunctionType{
		Name:           identifierToIdentType(ctx.GetName()),
		LongIdentifier: a2lStringToStringType(ctx.GetLongIdentifier()),
	})
}

func (n *Listener) ExitFunction(_ *parser.FunctionContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterFunctionVersion(ctx *parser.FunctionVersionContext) {
	n.Push(&FunctionVersionType{VersionIdentifier: a2lStringToStringType(ctx.GetVersionIdentifier())})
}

func (n *Listener) ExitFunctionVersion(_ *parser.FunctionVersionContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterGroup(ctx *parser.GroupContext) {
	n.Push(&GroupType{
		GroupName:           identifierToIdentType(ctx.GetGroupName()),
		GroupLongIdentifier: a2lStringToStringType(ctx.GetGroupLongIdentifier()),
	})
}

func (n *Listener) ExitGroup(_ *parser.GroupContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterGuardRails(_ *parser.GuardRailsContext) {
	n.Push(&GuardRailsType{})
}

func (n *Listener) ExitGuardRails(_ *parser.GuardRailsContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterHeader(ctx *parser.HeaderContext) {
	n.Push(&HeaderType{Comment: a2lStringToStringType(ctx.GetComment())})
}

func (n *Listener) ExitHeader(_ *parser.HeaderContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterIdentification(ctx *parser.IdentificationContext) {
	n.Push(&IdentificationType{
		Position: a2lIntToIntType(ctx.GetPosition()),
		DataType: &DataTypeType{Value: ctx.GetDatatype().GetText()},
	})
}

func (n *Listener) ExitIdentification(_ *parser.IdentificationContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterInMeasurement(ctx *parser.InMeasurementContext) {
	e := &InMeasurementType{Identifier: make([]*IdentType, 0)}

	for _, identifier := range ctx.GetIdentifier() {
		e.Identifier = append(e.Identifier, identifierToIdentType(identifier))
	}

	n.Push(e)
}

func (n *Listener) ExitInMeasurement(_ *parser.InMeasurementContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterLeftShift(ctx *parser.LeftShiftContext) {
	n.Push(&LeftShiftType{BitCount: a2lLongToLongType(ctx.GetBitcount())})
}

func (n *Listener) ExitLeftShift(_ *parser.LeftShiftContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterLocMeasurement(ctx *parser.LocMeasurementContext) {
	e := &LocMeasurementType{}

	for _, identifier := range ctx.GetIdentifier() {
		e.Identifier = append(e.Identifier, identifierToIdentType(identifier))
	}

	n.Push(e)
}

func (n *Listener) ExitLocMeasurement(_ *parser.LocMeasurementContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterMapList(ctx *parser.MapListContext) {
	e := &MapListType{Name: make([]*IdentType, 0)}

	for _, name := range ctx.GetName() {
		e.Name = append(e.Name, identifierToIdentType(name))
	}

	n.Push(e)
}

func (n *Listener) ExitMapList(_ *parser.MapListContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterMatrixDim(ctx *parser.MatrixDimContext) {
	n.Push(&MatrixDimType{
		XDim: a2lIntToIntType(ctx.GetXDim()),
		YDim: a2lIntToIntType(ctx.GetYDim()),
		ZDim: a2lIntToIntType(ctx.GetZDim()),
	})
}

func (n *Listener) ExitMatrixDim(_ *parser.MatrixDimContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterMaxGrad(ctx *parser.MaxGradContext) {
	n.Push(&MaxGradType{MaxGradient: floatToFloatType(ctx.GetMaxGradient())})
}

func (n *Listener) ExitMaxGrad(_ *parser.MaxGradContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterMaxRefresh(ctx *parser.MaxRefreshContext) {
	n.Push(&MaxRefreshType{
		ScalingUnit: a2lIntToIntType(ctx.GetScalingUnit()),
		Rate:        a2lLongToLongType(ctx.GetRate()),
	})
}

func (n *Listener) ExitMaxRefresh(_ *parser.MaxRefreshContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterMeasurement(ctx *parser.MeasurementContext) {
	n.Push(&MeasurementType{
		Name:           identifierToIdentType(ctx.GetName()),
		LongIdentifier: a2lStringToStringType(ctx.GetLongIdentifier()),
		DataType:       &DataTypeType{Value: ctx.DataType().GetText()},
		Conversion:     identifierToIdentType(ctx.GetConversion()),
		Resolution:     a2lIntToIntType(ctx.GetResolution()),
		Accuracy:       floatToFloatType(ctx.GetAccuracy()),
		LowerLimit:     floatToFloatType(ctx.GetLowerLimit()),
		UpperLimit:     floatToFloatType(ctx.GetUpperLimit()),
	})
}

func (n *Listener) ExitMeasurement(_ *parser.MeasurementContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterMemoryLayout(ctx *parser.MemoryLayoutContext) {
	n.Push(&MemoryLayoutType{
		PrgType: ctx.GetPrgType().GetText(),
		Address: a2lLongToLongType(ctx.GetAddress()),
		Size:    a2lLongToLongType(ctx.GetSize()),
		Offset: []*LongType{
			a2lLongToLongType(ctx.GetOffset_0()),
			a2lLongToLongType(ctx.GetOffset_1()),
			a2lLongToLongType(ctx.GetOffset_2()),
			a2lLongToLongType(ctx.GetOffset_3()),
			a2lLongToLongType(ctx.GetOffset_4()),
		},
	})
}

func (n *Listener) ExitMemoryLayout(_ *parser.MemoryLayoutContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterMemorySegment(ctx *parser.MemorySegmentContext) {
	n.Push(&MemorySegmentType{
		Name:           identifierToIdentType(ctx.GetName()),
		LongIdentifier: a2lStringToStringType(ctx.GetLongIdentifier()),
		PrgType:        ctx.GetPrgType().GetText(),
		MemoryType:     ctx.GetMemoryType().GetText(),
		Attribute:      ctx.GetAttribute().GetText(),
		Address:        a2lLongToLongType(ctx.GetAddress()),
		Size:           a2lLongToLongType(ctx.GetSize()),
		Offset: []*LongType{
			a2lLongToLongType(ctx.GetOffset_0()),
			a2lLongToLongType(ctx.GetOffset_1()),
			a2lLongToLongType(ctx.GetOffset_2()),
			a2lLongToLongType(ctx.GetOffset_3()),
			a2lLongToLongType(ctx.GetOffset_4())},
	})
}

func (n *Listener) ExitMemorySegment(_ *parser.MemorySegmentContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterModCommon(ctx *parser.ModCommonContext) {
	n.Push(&ModCommonType{Comment: a2lStringToStringType(ctx.GetComment())})
}

func (n *Listener) ExitModCommon(_ *parser.ModCommonContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterModPar(ctx *parser.ModParContext) {
	n.Push(&ModParType{Comment: a2lStringToStringType(ctx.GetComment())})
}

func (n *Listener) ExitModPar(_ *parser.ModParContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterModule(ctx *parser.ModuleContext) {
	n.Push(&ModuleType{
		Name:           identifierToIdentType(ctx.GetName()),
		LongIdentifier: a2lStringToStringType(ctx.GetLongIdentifier()),
	})
}

func (n *Listener) ExitModule(_ *parser.ModuleContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterMonotony(ctx *parser.MonotonyContext) {
	n.Push(&MonotonyType{Monotony: ctx.GetMonotony_().GetText()})
}

func (n *Listener) ExitMonotony(_ *parser.MonotonyContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterNoAxisPtsX(ctx *parser.NoAxisPtsXContext) {
	n.Push(&NoAxisPtsXType{
		Position: a2lIntToIntType(ctx.GetPosition()),
		DataType: &DataTypeType{Value: ctx.GetDatatype().GetText()},
	})
}

func (n *Listener) ExitNoAxisPtsX(_ *parser.NoAxisPtsXContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterNoAxisPtsY(ctx *parser.NoAxisPtsYContext) {
	n.Push(&NoAxisPtsYType{
		Position: a2lIntToIntType(ctx.GetPosition()),
		DataType: &DataTypeType{Value: ctx.GetDatatype().GetText()},
	})
}

func (n *Listener) ExitNoAxisPtsY(_ *parser.NoAxisPtsYContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterNoAxisPtsZ(ctx *parser.NoAxisPtsZContext) {
	n.Push(&NoAxisPtsZType{
		Position: a2lIntToIntType(ctx.GetPosition()),
		DataType: &DataTypeType{Value: ctx.GetDatatype().GetText()},
	})
}

func (n *Listener) ExitNoAxisPtsZ(_ *parser.NoAxisPtsZContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterNoOfInterfaces(ctx *parser.NoOfInterfacesContext) {
	n.Push(&NoOfInterfacesType{Num: a2lIntToIntType(ctx.GetNum())})
}

func (n *Listener) ExitNoOfInterfaces(_ *parser.NoOfInterfacesContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterNoRescaleX(ctx *parser.NoRescaleXContext) {
	n.Push(&NoRescaleXType{
		Position: a2lIntToIntType(ctx.GetPosition()),
		DataType: &DataTypeType{Value: ctx.GetDatatype().GetText()},
	})
}

func (n *Listener) ExitNoRescaleX(_ *parser.NoRescaleXContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterNoRescaleY(ctx *parser.NoRescaleYContext) {
	n.Push(&NoRescaleYType{
		Position: a2lIntToIntType(ctx.GetPosition()),
		DataType: &DataTypeType{Value: ctx.GetDatatype().GetText()},
	})
}

func (n *Listener) ExitNoRescaleY(_ *parser.NoRescaleYContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterNoRescaleZ(ctx *parser.NoRescaleZContext) {
	n.Push(&NoRescaleZType{
		Position: a2lIntToIntType(ctx.GetPosition()),
		DataType: &DataTypeType{Value: ctx.GetDatatype().GetText()},
	})
}

func (n *Listener) ExitNoRescaleZ(_ *parser.NoRescaleZContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterNumber(ctx *parser.NumberContext) {
	n.Push(&NumberType{Number: a2lIntToIntType(ctx.GetNumber_())})
}

func (n *Listener) ExitNumber(_ *parser.NumberContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterOffsetX(ctx *parser.OffsetXContext) {
	n.Push(&OffsetXType{
		Position: a2lIntToIntType(ctx.GetPosition()),
		DataType: &DataTypeType{Value: ctx.GetDatatype().GetText()},
	})
}

func (n *Listener) ExitOffsetX(_ *parser.OffsetXContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterOffsetY(ctx *parser.OffsetYContext) {
	n.Push(&OffsetYType{
		Position: a2lIntToIntType(ctx.GetPosition()),
		DataType: &DataTypeType{Value: ctx.GetDatatype().GetText()},
	})
}

func (n *Listener) ExitOffsetY(_ *parser.OffsetYContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterOffsetZ(ctx *parser.OffsetZContext) {
	n.Push(&OffsetZType{
		Position: a2lIntToIntType(ctx.GetPosition()),
		DataType: &DataTypeType{Value: ctx.GetDatatype().GetText()},
	})
}

func (n *Listener) ExitOffsetZ(_ *parser.OffsetZContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterOutMeasurement(ctx *parser.OutMeasurementContext) {
	e := &OutMeasurementType{Identifier: make([]*IdentType, 0)}

	for _, identifier := range ctx.GetIdentifier() {
		e.Identifier = append(e.Identifier, identifierToIdentType(identifier))
	}

	n.Push(e)
}

func (n *Listener) ExitOutMeasurement(_ *parser.OutMeasurementContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterPhoneNo(ctx *parser.PhoneNoContext) {
	n.Push(&PhoneNoType{TelNum: a2lStringToStringType(ctx.GetTelnum())})
}

func (n *Listener) ExitPhoneNo(_ *parser.PhoneNoContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterProjectNo(ctx *parser.ProjectNoContext) {
	n.Push(&ProjectNoType{ProjectNumber: identifierToIdentType(ctx.GetProjectNumber())})
}

func (n *Listener) ExitProjectNo(_ *parser.ProjectNoContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterProject(ctx *parser.ProjectContext) {
	n.Push(&ProjectType{
		Name:           identifierToIdentType(ctx.GetName()),
		LongIdentifier: a2lStringToStringType(ctx.GetLongIdentifier()),
	})
}

func (n *Listener) ExitProject(_ *parser.ProjectContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterReadOnly(_ *parser.ReadOnlyContext) {
	n.Push(&ReadOnlyType{})
}

func (n *Listener) ExitReadOnly(_ *parser.ReadOnlyContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterReadWrite(_ *parser.ReadWriteContext) {
	n.Push(&ReadWriteType{})
}

func (n *Listener) ExitReadWrite(_ *parser.ReadWriteContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterRecordLayout(ctx *parser.RecordLayoutContext) {
	n.Push(&RecordLayoutType{Name: identifierToIdentType(ctx.GetName())})
}

func (n *Listener) ExitRecordLayout(_ *parser.RecordLayoutContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterRefCharacteristic(ctx *parser.RefCharacteristicContext) {
	e := &RefCharacteristicType{Identifier: make([]*IdentType, 0)}

	for _, identifier := range ctx.GetIdentifier() {
		e.Identifier = append(e.Identifier, identifierToIdentType(identifier))
	}

	n.Push(e)
}

func (n *Listener) ExitRefCharacteristic(_ *parser.RefCharacteristicContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterRefGroup(ctx *parser.RefGroupContext) {
	e := &RefGroupType{Identifier: make([]*IdentType, 0)}

	for _, identifier := range ctx.GetIdentifier() {
		e.Identifier = append(e.Identifier, identifierToIdentType(identifier))
	}

	n.Push(e)
}

func (n *Listener) ExitRefGroup(_ *parser.RefGroupContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterRefMeasurement(ctx *parser.RefMeasurementContext) {
	e := &RefMeasurementType{Identifier: make([]*IdentType, 0)}

	for _, identifier := range ctx.GetIdentifier() {
		e.Identifier = append(e.Identifier, identifierToIdentType(identifier))
	}

	n.Push(e)
}

func (n *Listener) ExitRefMeasurement(_ *parser.RefMeasurementContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterRefMemorySegment(ctx *parser.RefMemorySegmentContext) {
	n.Push(&RefMemorySegmentType{Name: identifierToIdentType(ctx.GetName())})
}

func (n *Listener) ExitRefMemorySegment(_ *parser.RefMemorySegmentContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterRefUnit(ctx *parser.RefUnitContext) {
	n.Push(&RefUnitType{Unit: identifierToIdentType(ctx.GetUnit_())})
}

func (n *Listener) ExitRefUnit(_ *parser.RefUnitContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterReserved(ctx *parser.ReservedContext) {
	n.Push(&ReservedType{
		Position: a2lIntToIntType(ctx.GetPosition()),
		DataSize: &DataSizeType{Value: ctx.GetDataSize_().GetText()},
	})
}

func (n *Listener) ExitReserved(_ *parser.ReservedContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterRightShift(ctx *parser.RightShiftContext) {
	n.Push(&RightShiftType{BitCount: a2lLongToLongType(ctx.GetBitcount())})
}

func (n *Listener) ExitRightShift(_ *parser.RightShiftContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterRipAddrW(ctx *parser.RipAddrWContext) {
	n.Push(&RipAddrWType{
		Position: a2lIntToIntType(ctx.GetPosition()),
		DataType: &DataTypeType{Value: ctx.GetDatatype().GetText()},
	})
}

func (n *Listener) ExitRipAddrW(_ *parser.RipAddrWContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterRipAddrX(ctx *parser.RipAddrXContext) {
	n.Push(&RipAddrXType{
		Position: a2lIntToIntType(ctx.GetPosition()),
		DataType: &DataTypeType{Value: ctx.GetDatatype().GetText()},
	})
}

func (n *Listener) ExitRipAddrX(_ *parser.RipAddrXContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterRipAddrY(ctx *parser.RipAddrYContext) {
	n.Push(&RipAddrYType{
		Position: a2lIntToIntType(ctx.GetPosition()),
		DataType: &DataTypeType{Value: ctx.GetDatatype().GetText()},
	})
}

func (n *Listener) ExitRipAddrY(_ *parser.RipAddrYContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterRipAddrZ(ctx *parser.RipAddrZContext) {
	n.Push(&RipAddrZType{
		Position: a2lIntToIntType(ctx.GetPosition()),
		DataType: &DataTypeType{Value: ctx.GetDatatype().GetText()},
	})
}

func (n *Listener) ExitRipAddrZ(_ *parser.RipAddrZContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterRoot(_ *parser.RootContext) {
	n.Push(&RootType{})
}

func (n *Listener) ExitRoot(_ *parser.RootContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterShiftOpX(ctx *parser.ShiftOpXContext) {
	n.Push(&ShiftOpXType{
		Position: a2lIntToIntType(ctx.GetPosition()),
		DataType: &DataTypeType{Value: ctx.GetDatatype().GetText()},
	})
}

func (n *Listener) ExitShiftOpX(_ *parser.ShiftOpXContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterShiftOpY(ctx *parser.ShiftOpYContext) {
	n.Push(&ShiftOpYType{
		Position: a2lIntToIntType(ctx.GetPosition()),
		DataType: &DataTypeType{Value: ctx.GetDatatype().GetText()},
	})
}

func (n *Listener) ExitShiftOpY(_ *parser.ShiftOpYContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterShiftOpZ(ctx *parser.ShiftOpZContext) {
	n.Push(&ShiftOpZType{
		Position: a2lIntToIntType(ctx.GetPosition()),
		DataType: &DataTypeType{Value: ctx.GetDatatype().GetText()},
	})
}

func (n *Listener) ExitShiftOpZ(_ *parser.ShiftOpZContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterSiExponents(ctx *parser.SiExponentsContext) {
	n.Push(&SiExponentsType{
		Length:            a2lIntToIntType(ctx.GetLength()),
		Mass:              a2lIntToIntType(ctx.GetMass()),
		Time:              a2lIntToIntType(ctx.GetTime()),
		ElectricCurrent:   a2lIntToIntType(ctx.GetElectricCurrent()),
		Temperature:       a2lIntToIntType(ctx.GetTemperature()),
		AmountOfSubstance: a2lIntToIntType(ctx.GetAmountOfSubstance()),
		LuminousIntensity: a2lIntToIntType(ctx.GetLuminousIntensity()),
	})
}

func (n *Listener) ExitSiExponents(_ *parser.SiExponentsContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterSignExtend(_ *parser.SignExtendContext) {
	n.Push(&SignExtendType{})
}

func (n *Listener) ExitSignExtend(_ *parser.SignExtendContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterSrcAddrX(ctx *parser.SrcAddrXContext) {
	n.Push(&SrcAddrXType{
		Position: a2lIntToIntType(ctx.GetPosition()),
		DataType: &DataTypeType{Value: ctx.GetDatatype().GetText()},
	})
}

func (n *Listener) ExitSrcAddrX(_ *parser.SrcAddrXContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterSrcAddrY(ctx *parser.SrcAddrYContext) {
	n.Push(&SrcAddrYType{
		Position: a2lIntToIntType(ctx.GetPosition()),
		DataType: &DataTypeType{Value: ctx.GetDatatype().GetText()},
	})
}

func (n *Listener) ExitSrcAddrY(_ *parser.SrcAddrYContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterSrcAddrZ(ctx *parser.SrcAddrZContext) {
	n.Push(&SrcAddrZType{
		Position: a2lIntToIntType(ctx.GetPosition()),
		DataType: &DataTypeType{Value: ctx.GetDatatype().GetText()},
	})
}

func (n *Listener) ExitSrcAddrZ(_ *parser.SrcAddrZContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterSRecLayout(ctx *parser.SRecLayoutContext) {
	n.Push(&SRecLayoutType{Name: identifierToIdentType(ctx.IdentifierValue())})
}

func (n *Listener) ExitSRecLayout(_ *parser.SRecLayoutContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterSubFunction(ctx *parser.SubFunctionContext) {
	e := &SubFunctionType{Identifier: make([]*IdentType, 0)}

	for _, identifier := range ctx.GetIdentifier() {
		e.Identifier = append(e.Identifier, identifierToIdentType(identifier))
	}

	n.Push(e)
}

func (n *Listener) ExitSubFunction(_ *parser.SubFunctionContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterSubGroup(ctx *parser.SubGroupContext) {
	e := &SubGroupType{Identifier: make([]*IdentType, 0)}

	for _, identifier := range ctx.GetIdentifier() {
		e.Identifier = append(e.Identifier, identifierToIdentType(identifier))
	}

	n.Push(e)
}

func (n *Listener) ExitSubGroup(_ *parser.SubGroupContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterSupplier(ctx *parser.SupplierContext) {
	n.Push(&SupplierType{Manufacturer: a2lStringToStringType(ctx.GetManufacturer())})
}

func (n *Listener) ExitSupplier(_ *parser.SupplierContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterSystemConstant(ctx *parser.SystemConstantContext) {
	n.Push(&SystemConstantType{
		Name:  a2lStringToStringType(ctx.GetName()),
		Value: a2lStringToStringType(ctx.GetValue_()),
	})
}

func (n *Listener) ExitSystemConstant(_ *parser.SystemConstantContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterUnitConversion(ctx *parser.UnitConversionContext) {
	n.Push(&UnitConversionType{
		Gradient: floatToFloatType(ctx.GetGradient()),
		Offset:   floatToFloatType(ctx.GetOffset()),
	})
}

func (n *Listener) ExitUnitConversion(_ *parser.UnitConversionContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterUnit(ctx *parser.UnitContext) {
	n.Push(&UnitType{
		Name:           identifierToIdentType(ctx.GetName()),
		LongIdentifier: a2lStringToStringType(ctx.GetLongIdentifier()),
		Display:        a2lStringToStringType(ctx.GetDisplay()),
		Type:           ctx.GetType_().GetText(),
	})
}

func (n *Listener) ExitUnit(_ *parser.UnitContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterUserRights(ctx *parser.UserRightsContext) {
	n.Push(&UserRightsType{UserLevelId: identifierToIdentType(ctx.GetUserLevelId())})
}

func (n *Listener) ExitUserRights(_ *parser.UserRightsContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterUser(ctx *parser.UserContext) {
	n.Push(&UserType{UserName: a2lStringToStringType(ctx.GetUserName())})
}

func (n *Listener) ExitUser(_ *parser.UserContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterVarAddress(ctx *parser.VarAddressContext) {
	e := &VarAddressType{Address: make([]*LongType, 0)}

	for _, address := range ctx.GetAddress() {
		e.Address = append(e.Address, a2lLongToLongType(address))
	}

	n.Push(e)
}

func (n *Listener) ExitVarAddress(_ *parser.VarAddressContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterVarCharacteristic(ctx *parser.VarCharacteristicContext) {
	e := &VarCharacteristicType{
		Name:          identifierToIdentType(ctx.GetName()),
		CriterionName: make([]*IdentType, 0),
	}

	for _, criterionName := range ctx.GetCriterionName() {
		e.CriterionName = append(e.CriterionName, identifierToIdentType(criterionName))
	}

	n.Push(e)
}

func (n *Listener) ExitVarCharacteristic(_ *parser.VarCharacteristicContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterVarCriterion(ctx *parser.VarCriterionContext) {
	e := &VarCriterionType{
		Name:           identifierToIdentType(ctx.GetName()),
		LongIdentifier: a2lStringToStringType(ctx.GetLongIdentifier()),
		Value:          make([]*IdentType, 0),
	}

	for _, value := range ctx.GetValue_() {
		e.Value = append(e.Value, identifierToIdentType(value))
	}

	n.Push(e)
}

func (n *Listener) ExitVarCriterion(_ *parser.VarCriterionContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterVarForbiddenComb(ctx *parser.VarForbiddenCombContext) {
	e := &VarForbiddenCombType{CriterionNameCriterionValue: make([]*VarForbiddenCombType_CriterionType, 0)}

	for i := 0; i < len(ctx.GetCriterionName()); i++ {
		e.CriterionNameCriterionValue = append(e.CriterionNameCriterionValue, &VarForbiddenCombType_CriterionType{
			CriterionName:  identifierToIdentType(ctx.GetCriterionName()[i]),
			CriterionValue: identifierToIdentType(ctx.GetCriterionValue()[i]),
		})
	}

	n.Push(e)
}

func (n *Listener) ExitVarForbiddenComb(_ *parser.VarForbiddenCombContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterVariantCoding(_ *parser.VariantCodingContext) {
	n.Push(&VariantCodingType{})
}

func (n *Listener) ExitVariantCoding(_ *parser.VariantCodingContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterVarMeasurement(ctx *parser.VarMeasurementContext) {
	n.Push(&VarMeasurementType{Name: identifierToIdentType(ctx.GetName())})
}

func (n *Listener) ExitVarMeasurement(_ *parser.VarMeasurementContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterVarNaming(ctx *parser.VarNamingContext) {
	n.Push(&VarNamingType{Tag: ctx.GetTag().GetText()})
}

func (n *Listener) ExitVarNaming(_ *parser.VarNamingContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterVarSelectionCharacteristic(ctx *parser.VarSelectionCharacteristicContext) {
	n.Push(&VarSelectionCharacteristicType{Name: identifierToIdentType(ctx.GetName())})
}

func (n *Listener) ExitVarSelectionCharacteristic(_ *parser.VarSelectionCharacteristicContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterVarSeparator(ctx *parser.VarSeparatorContext) {
	n.Push(&VarSeparatorType{Separator: a2lStringToStringType(ctx.GetSeparator())})
}

func (n *Listener) ExitVarSeparator(_ *parser.VarSeparatorContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterVersion(ctx *parser.VersionContext) {
	n.Push(&VersionType{VersionIdentifier: a2lStringToStringType(ctx.GetVersionIdentifier())})
}

func (n *Listener) ExitVersion(_ *parser.VersionContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterVirtualCharacteristic(ctx *parser.VirtualCharacteristicContext) {
	e := &VirtualCharacteristicType{
		Formula:        a2lStringToStringType(ctx.GetFormula_()),
		Characteristic: make([]*IdentType, 0),
	}

	for _, characteristic := range ctx.GetCharacteristic_() {
		e.Characteristic = append(e.Characteristic, identifierToIdentType(characteristic))
	}

	n.Push(e)
}

func (n *Listener) ExitVirtualCharacteristic(_ *parser.VirtualCharacteristicContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}

func (n *Listener) EnterVirtual(ctx *parser.VirtualContext) {
	e := &VirtualType{MeasuringChannel: make([]*IdentType, 0)}

	for _, measuringChannel := range ctx.GetMeasuringChannel() {
		e.MeasuringChannel = append(e.MeasuringChannel, identifierToIdentType(measuringChannel))
	}

	n.Push(e)
}

func (n *Listener) ExitVirtual(_ *parser.VirtualContext) {
	node := n.Pop()

	n.Last().MapChildNodes(node)
}
