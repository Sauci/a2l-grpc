// Code generated from grammar/A2L.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // A2L

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BaseA2LVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseA2LVisitor) VisitA2lFile(ctx *A2lFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitAlignmentByte(ctx *AlignmentByteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitAlignmentFloat16Ieee(ctx *AlignmentFloat16IeeeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitAlignmentFloat32Ieee(ctx *AlignmentFloat32IeeeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitAlignmentFloat64Ieee(ctx *AlignmentFloat64IeeeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitAlignmentInt64(ctx *AlignmentInt64Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitAlignmentLong(ctx *AlignmentLongContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitAlignmentWord(ctx *AlignmentWordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitAnnotation(ctx *AnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitAnnotationLabel(ctx *AnnotationLabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitAnnotationOrigin(ctx *AnnotationOriginContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitAnnotationText(ctx *AnnotationTextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitBitMask(ctx *BitMaskContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitByteOrder(ctx *ByteOrderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitCalibrationAccess(ctx *CalibrationAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitDefaultValue(ctx *DefaultValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitDeposit(ctx *DepositContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitDiscrete(ctx *DiscreteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitDisplayIdentifier(ctx *DisplayIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitEcuAddressExtension(ctx *EcuAddressExtensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitExtendedLimits(ctx *ExtendedLimitsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitFormat(ctx *FormatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitFunctionList(ctx *FunctionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitGuardRails(ctx *GuardRailsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitMatrixDim(ctx *MatrixDimContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitMaxRefresh(ctx *MaxRefreshContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitMonotony(ctx *MonotonyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitPhysUnit(ctx *PhysUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitReadOnly(ctx *ReadOnlyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitRefCharacteristic(ctx *RefCharacteristicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitRefMemorySegment(ctx *RefMemorySegmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitRefUnit(ctx *RefUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitStepSize(ctx *StepSizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitSymbolLink(ctx *SymbolLinkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitVersion(ctx *VersionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitAsap2Version(ctx *Asap2VersionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitA2mlVersion(ctx *A2mlVersionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitProject(ctx *ProjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitHeader(ctx *HeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitProjectNo(ctx *ProjectNoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitModule(ctx *ModuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitAxisPts(ctx *AxisPtsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitCharacteristic(ctx *CharacteristicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitAxisDescr(ctx *AxisDescrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitAxisPtsRef(ctx *AxisPtsRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitCurveAxisRef(ctx *CurveAxisRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitFixAxisPar(ctx *FixAxisParContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitFixAxisParDist(ctx *FixAxisParDistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitFixAxisParList(ctx *FixAxisParListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitMaxGrad(ctx *MaxGradContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitComparisonQuantity(ctx *ComparisonQuantityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitDependentCharacteristic(ctx *DependentCharacteristicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitMapList(ctx *MapListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitVirtualCharacteristic(ctx *VirtualCharacteristicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitCompuMethod(ctx *CompuMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitCoeffs(ctx *CoeffsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitCoeffsLinear(ctx *CoeffsLinearContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitCompuTabRef(ctx *CompuTabRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitFormula(ctx *FormulaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitFormulaInv(ctx *FormulaInvContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitStatusStringRef(ctx *StatusStringRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitCompuTab(ctx *CompuTabContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitDefaultValueNumeric(ctx *DefaultValueNumericContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitCompuVTab(ctx *CompuVTabContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitCompuVTabRange(ctx *CompuVTabRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitFrame(ctx *FrameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitFrameMeasurement(ctx *FrameMeasurementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitDefCharacteristic(ctx *DefCharacteristicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitFunctionVersion(ctx *FunctionVersionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitInMeasurement(ctx *InMeasurementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitLocMeasurement(ctx *LocMeasurementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitOutMeasurement(ctx *OutMeasurementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitSubFunction(ctx *SubFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitGroup(ctx *GroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitRefMeasurement(ctx *RefMeasurementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitRoot(ctx *RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitSubGroup(ctx *SubGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitInstance(ctx *InstanceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitMeasurement(ctx *MeasurementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitArraySize(ctx *ArraySizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitBitOperation(ctx *BitOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitLeftShift(ctx *LeftShiftContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitRightShift(ctx *RightShiftContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitSignExtend(ctx *SignExtendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitEcuAddress(ctx *EcuAddressContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitErrorMask(ctx *ErrorMaskContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitLayout(ctx *LayoutContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitReadWrite(ctx *ReadWriteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitVirtual(ctx *VirtualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitModCommon(ctx *ModCommonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitDataSize(ctx *DataSizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitSRecLayout(ctx *SRecLayoutContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitModPar(ctx *ModParContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitAddrEpk(ctx *AddrEpkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitCalibrationMethod(ctx *CalibrationMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitCalibrationHandle(ctx *CalibrationHandleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitCpuType(ctx *CpuTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitCustomer(ctx *CustomerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitCustomerNo(ctx *CustomerNoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitEcu(ctx *EcuContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitEcuCalibrationOffset(ctx *EcuCalibrationOffsetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitEpk(ctx *EpkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitMemoryLayout(ctx *MemoryLayoutContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitMemorySegment(ctx *MemorySegmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitNoOfInterfaces(ctx *NoOfInterfacesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitPhoneNo(ctx *PhoneNoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitSupplier(ctx *SupplierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitSystemConstant(ctx *SystemConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitUser(ctx *UserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitRecordLayout(ctx *RecordLayoutContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitAxisPtsX(ctx *AxisPtsXContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitAxisPtsY(ctx *AxisPtsYContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitAxisPtsZ(ctx *AxisPtsZContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitAxisPts4(ctx *AxisPts4Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitAxisPts5(ctx *AxisPts5Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitAxisRescaleX(ctx *AxisRescaleXContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitAxisRescaleY(ctx *AxisRescaleYContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitAxisRescaleZ(ctx *AxisRescaleZContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitAxisRescale4(ctx *AxisRescale4Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitAxisRescale5(ctx *AxisRescale5Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitDistOpX(ctx *DistOpXContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitDistOpY(ctx *DistOpYContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitDistOpZ(ctx *DistOpZContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitDistOp4(ctx *DistOp4Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitDistOp5(ctx *DistOp5Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitFixNoAxisPtsX(ctx *FixNoAxisPtsXContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitFixNoAxisPtsY(ctx *FixNoAxisPtsYContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitFixNoAxisPtsZ(ctx *FixNoAxisPtsZContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitFixNoAxisPts4(ctx *FixNoAxisPts4Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitFixNoAxisPts5(ctx *FixNoAxisPts5Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitFncValues(ctx *FncValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitIdentification(ctx *IdentificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitNoAxisPtsX(ctx *NoAxisPtsXContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitNoAxisPtsY(ctx *NoAxisPtsYContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitNoAxisPtsZ(ctx *NoAxisPtsZContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitNoAxisPts4(ctx *NoAxisPts4Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitNoAxisPts5(ctx *NoAxisPts5Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitStaticRecordLayout(ctx *StaticRecordLayoutContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitNoRescaleX(ctx *NoRescaleXContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitNoRescaleY(ctx *NoRescaleYContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitNoRescaleZ(ctx *NoRescaleZContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitNoRescale4(ctx *NoRescale4Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitNoRescale5(ctx *NoRescale5Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitOffsetX(ctx *OffsetXContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitOffsetY(ctx *OffsetYContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitOffsetZ(ctx *OffsetZContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitOffset4(ctx *Offset4Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitOffset5(ctx *Offset5Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitReserved(ctx *ReservedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitRipAddrW(ctx *RipAddrWContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitRipAddrX(ctx *RipAddrXContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitRipAddrY(ctx *RipAddrYContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitRipAddrZ(ctx *RipAddrZContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitRipAddr4(ctx *RipAddr4Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitRipAddr5(ctx *RipAddr5Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitShiftOpX(ctx *ShiftOpXContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitShiftOpY(ctx *ShiftOpYContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitShiftOpZ(ctx *ShiftOpZContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitShiftOp4(ctx *ShiftOp4Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitShiftOp5(ctx *ShiftOp5Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitSrcAddrX(ctx *SrcAddrXContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitSrcAddrY(ctx *SrcAddrYContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitSrcAddrZ(ctx *SrcAddrZContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitSrcAddr4(ctx *SrcAddr4Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitSrcAddr5(ctx *SrcAddr5Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitTypedefCharacteristic(ctx *TypedefCharacteristicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitTypedefMeasurement(ctx *TypedefMeasurementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitTypedefStructure(ctx *TypedefStructureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitStructureComponent(ctx *StructureComponentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitUnit(ctx *UnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitSiExponents(ctx *SiExponentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitUnitConversion(ctx *UnitConversionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitUserRights(ctx *UserRightsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitRefGroup(ctx *RefGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitVariantCoding(ctx *VariantCodingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitVarCharacteristic(ctx *VarCharacteristicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitVarAddress(ctx *VarAddressContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitVarCriterion(ctx *VarCriterionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitVarMeasurement(ctx *VarMeasurementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitVarSelectionCharacteristic(ctx *VarSelectionCharacteristicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitVarForbiddenComb(ctx *VarForbiddenCombContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitVarNaming(ctx *VarNamingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitVarSeparator(ctx *VarSeparatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitIntegerValue(ctx *IntegerValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitNumericValue(ctx *NumericValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitStringValue(ctx *StringValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitIdentifierValue(ctx *IdentifierValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitPartialIdentifier(ctx *PartialIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitArraySpecifier(ctx *ArraySpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitDataType(ctx *DataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitDatasize(ctx *DatasizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitAddrtype(ctx *AddrtypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitByteOrderValue(ctx *ByteOrderValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitIndexorder(ctx *IndexorderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitLinkType(ctx *LinkTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitA2ml(ctx *A2mlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitTypeDefinition(ctx *TypeDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitA2mlTypeName(ctx *A2mlTypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitPredefinedTypeName(ctx *PredefinedTypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitBlockDefinition(ctx *BlockDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitEnumTypeName(ctx *EnumTypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitEnumeratorList(ctx *EnumeratorListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitEnumerator(ctx *EnumeratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitStructTypeName(ctx *StructTypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitStructMember(ctx *StructMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitMember(ctx *MemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitTaggedStructTypeName(ctx *TaggedStructTypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitTaggedStructMember(ctx *TaggedStructMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitTaggedStructDefinition(ctx *TaggedStructDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitTaggedUnionTypeName(ctx *TaggedUnionTypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitTaggedUnionMember(ctx *TaggedUnionMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitTagValue(ctx *TagValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitIfData(ctx *IfDataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitGenericParameter(ctx *GenericParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseA2LVisitor) VisitGenericNode(ctx *GenericNodeContext) interface{} {
	return v.VisitChildren(ctx)
}
