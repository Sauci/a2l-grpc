// Code generated from grammar/A2L.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // A2L

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by A2LParser.
type A2LVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by A2LParser#a2lFile.
	VisitA2lFile(ctx *A2lFileContext) interface{}

	// Visit a parse tree produced by A2LParser#alignmentByte.
	VisitAlignmentByte(ctx *AlignmentByteContext) interface{}

	// Visit a parse tree produced by A2LParser#alignmentFloat16Ieee.
	VisitAlignmentFloat16Ieee(ctx *AlignmentFloat16IeeeContext) interface{}

	// Visit a parse tree produced by A2LParser#alignmentFloat32Ieee.
	VisitAlignmentFloat32Ieee(ctx *AlignmentFloat32IeeeContext) interface{}

	// Visit a parse tree produced by A2LParser#alignmentFloat64Ieee.
	VisitAlignmentFloat64Ieee(ctx *AlignmentFloat64IeeeContext) interface{}

	// Visit a parse tree produced by A2LParser#alignmentInt64.
	VisitAlignmentInt64(ctx *AlignmentInt64Context) interface{}

	// Visit a parse tree produced by A2LParser#alignmentLong.
	VisitAlignmentLong(ctx *AlignmentLongContext) interface{}

	// Visit a parse tree produced by A2LParser#alignmentWord.
	VisitAlignmentWord(ctx *AlignmentWordContext) interface{}

	// Visit a parse tree produced by A2LParser#annotation.
	VisitAnnotation(ctx *AnnotationContext) interface{}

	// Visit a parse tree produced by A2LParser#annotationLabel.
	VisitAnnotationLabel(ctx *AnnotationLabelContext) interface{}

	// Visit a parse tree produced by A2LParser#annotationOrigin.
	VisitAnnotationOrigin(ctx *AnnotationOriginContext) interface{}

	// Visit a parse tree produced by A2LParser#annotationText.
	VisitAnnotationText(ctx *AnnotationTextContext) interface{}

	// Visit a parse tree produced by A2LParser#bitMask.
	VisitBitMask(ctx *BitMaskContext) interface{}

	// Visit a parse tree produced by A2LParser#byteOrder.
	VisitByteOrder(ctx *ByteOrderContext) interface{}

	// Visit a parse tree produced by A2LParser#calibrationAccess.
	VisitCalibrationAccess(ctx *CalibrationAccessContext) interface{}

	// Visit a parse tree produced by A2LParser#defaultValue.
	VisitDefaultValue(ctx *DefaultValueContext) interface{}

	// Visit a parse tree produced by A2LParser#deposit.
	VisitDeposit(ctx *DepositContext) interface{}

	// Visit a parse tree produced by A2LParser#discrete.
	VisitDiscrete(ctx *DiscreteContext) interface{}

	// Visit a parse tree produced by A2LParser#displayIdentifier.
	VisitDisplayIdentifier(ctx *DisplayIdentifierContext) interface{}

	// Visit a parse tree produced by A2LParser#ecuAddressExtension.
	VisitEcuAddressExtension(ctx *EcuAddressExtensionContext) interface{}

	// Visit a parse tree produced by A2LParser#extendedLimits.
	VisitExtendedLimits(ctx *ExtendedLimitsContext) interface{}

	// Visit a parse tree produced by A2LParser#format.
	VisitFormat(ctx *FormatContext) interface{}

	// Visit a parse tree produced by A2LParser#functionList.
	VisitFunctionList(ctx *FunctionListContext) interface{}

	// Visit a parse tree produced by A2LParser#guardRails.
	VisitGuardRails(ctx *GuardRailsContext) interface{}

	// Visit a parse tree produced by A2LParser#matrixDim.
	VisitMatrixDim(ctx *MatrixDimContext) interface{}

	// Visit a parse tree produced by A2LParser#maxRefresh.
	VisitMaxRefresh(ctx *MaxRefreshContext) interface{}

	// Visit a parse tree produced by A2LParser#monotony.
	VisitMonotony(ctx *MonotonyContext) interface{}

	// Visit a parse tree produced by A2LParser#physUnit.
	VisitPhysUnit(ctx *PhysUnitContext) interface{}

	// Visit a parse tree produced by A2LParser#readOnly.
	VisitReadOnly(ctx *ReadOnlyContext) interface{}

	// Visit a parse tree produced by A2LParser#refCharacteristic.
	VisitRefCharacteristic(ctx *RefCharacteristicContext) interface{}

	// Visit a parse tree produced by A2LParser#refMemorySegment.
	VisitRefMemorySegment(ctx *RefMemorySegmentContext) interface{}

	// Visit a parse tree produced by A2LParser#refUnit.
	VisitRefUnit(ctx *RefUnitContext) interface{}

	// Visit a parse tree produced by A2LParser#stepSize.
	VisitStepSize(ctx *StepSizeContext) interface{}

	// Visit a parse tree produced by A2LParser#symbolLink.
	VisitSymbolLink(ctx *SymbolLinkContext) interface{}

	// Visit a parse tree produced by A2LParser#version.
	VisitVersion(ctx *VersionContext) interface{}

	// Visit a parse tree produced by A2LParser#asap2Version.
	VisitAsap2Version(ctx *Asap2VersionContext) interface{}

	// Visit a parse tree produced by A2LParser#a2mlVersion.
	VisitA2mlVersion(ctx *A2mlVersionContext) interface{}

	// Visit a parse tree produced by A2LParser#project.
	VisitProject(ctx *ProjectContext) interface{}

	// Visit a parse tree produced by A2LParser#header.
	VisitHeader(ctx *HeaderContext) interface{}

	// Visit a parse tree produced by A2LParser#projectNo.
	VisitProjectNo(ctx *ProjectNoContext) interface{}

	// Visit a parse tree produced by A2LParser#module.
	VisitModule(ctx *ModuleContext) interface{}

	// Visit a parse tree produced by A2LParser#axisPts.
	VisitAxisPts(ctx *AxisPtsContext) interface{}

	// Visit a parse tree produced by A2LParser#characteristic.
	VisitCharacteristic(ctx *CharacteristicContext) interface{}

	// Visit a parse tree produced by A2LParser#axisDescr.
	VisitAxisDescr(ctx *AxisDescrContext) interface{}

	// Visit a parse tree produced by A2LParser#axisPtsRef.
	VisitAxisPtsRef(ctx *AxisPtsRefContext) interface{}

	// Visit a parse tree produced by A2LParser#curveAxisRef.
	VisitCurveAxisRef(ctx *CurveAxisRefContext) interface{}

	// Visit a parse tree produced by A2LParser#fixAxisPar.
	VisitFixAxisPar(ctx *FixAxisParContext) interface{}

	// Visit a parse tree produced by A2LParser#fixAxisParDist.
	VisitFixAxisParDist(ctx *FixAxisParDistContext) interface{}

	// Visit a parse tree produced by A2LParser#fixAxisParList.
	VisitFixAxisParList(ctx *FixAxisParListContext) interface{}

	// Visit a parse tree produced by A2LParser#maxGrad.
	VisitMaxGrad(ctx *MaxGradContext) interface{}

	// Visit a parse tree produced by A2LParser#comparisonQuantity.
	VisitComparisonQuantity(ctx *ComparisonQuantityContext) interface{}

	// Visit a parse tree produced by A2LParser#dependentCharacteristic.
	VisitDependentCharacteristic(ctx *DependentCharacteristicContext) interface{}

	// Visit a parse tree produced by A2LParser#mapList.
	VisitMapList(ctx *MapListContext) interface{}

	// Visit a parse tree produced by A2LParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by A2LParser#virtualCharacteristic.
	VisitVirtualCharacteristic(ctx *VirtualCharacteristicContext) interface{}

	// Visit a parse tree produced by A2LParser#compuMethod.
	VisitCompuMethod(ctx *CompuMethodContext) interface{}

	// Visit a parse tree produced by A2LParser#coeffs.
	VisitCoeffs(ctx *CoeffsContext) interface{}

	// Visit a parse tree produced by A2LParser#coeffsLinear.
	VisitCoeffsLinear(ctx *CoeffsLinearContext) interface{}

	// Visit a parse tree produced by A2LParser#compuTabRef.
	VisitCompuTabRef(ctx *CompuTabRefContext) interface{}

	// Visit a parse tree produced by A2LParser#formula.
	VisitFormula(ctx *FormulaContext) interface{}

	// Visit a parse tree produced by A2LParser#formulaInv.
	VisitFormulaInv(ctx *FormulaInvContext) interface{}

	// Visit a parse tree produced by A2LParser#statusStringRef.
	VisitStatusStringRef(ctx *StatusStringRefContext) interface{}

	// Visit a parse tree produced by A2LParser#compuTab.
	VisitCompuTab(ctx *CompuTabContext) interface{}

	// Visit a parse tree produced by A2LParser#defaultValueNumeric.
	VisitDefaultValueNumeric(ctx *DefaultValueNumericContext) interface{}

	// Visit a parse tree produced by A2LParser#compuVTab.
	VisitCompuVTab(ctx *CompuVTabContext) interface{}

	// Visit a parse tree produced by A2LParser#compuVTabRange.
	VisitCompuVTabRange(ctx *CompuVTabRangeContext) interface{}

	// Visit a parse tree produced by A2LParser#frame.
	VisitFrame(ctx *FrameContext) interface{}

	// Visit a parse tree produced by A2LParser#frameMeasurement.
	VisitFrameMeasurement(ctx *FrameMeasurementContext) interface{}

	// Visit a parse tree produced by A2LParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by A2LParser#defCharacteristic.
	VisitDefCharacteristic(ctx *DefCharacteristicContext) interface{}

	// Visit a parse tree produced by A2LParser#functionVersion.
	VisitFunctionVersion(ctx *FunctionVersionContext) interface{}

	// Visit a parse tree produced by A2LParser#inMeasurement.
	VisitInMeasurement(ctx *InMeasurementContext) interface{}

	// Visit a parse tree produced by A2LParser#locMeasurement.
	VisitLocMeasurement(ctx *LocMeasurementContext) interface{}

	// Visit a parse tree produced by A2LParser#outMeasurement.
	VisitOutMeasurement(ctx *OutMeasurementContext) interface{}

	// Visit a parse tree produced by A2LParser#subFunction.
	VisitSubFunction(ctx *SubFunctionContext) interface{}

	// Visit a parse tree produced by A2LParser#group.
	VisitGroup(ctx *GroupContext) interface{}

	// Visit a parse tree produced by A2LParser#refMeasurement.
	VisitRefMeasurement(ctx *RefMeasurementContext) interface{}

	// Visit a parse tree produced by A2LParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by A2LParser#subGroup.
	VisitSubGroup(ctx *SubGroupContext) interface{}

	// Visit a parse tree produced by A2LParser#instance.
	VisitInstance(ctx *InstanceContext) interface{}

	// Visit a parse tree produced by A2LParser#measurement.
	VisitMeasurement(ctx *MeasurementContext) interface{}

	// Visit a parse tree produced by A2LParser#arraySize.
	VisitArraySize(ctx *ArraySizeContext) interface{}

	// Visit a parse tree produced by A2LParser#bitOperation.
	VisitBitOperation(ctx *BitOperationContext) interface{}

	// Visit a parse tree produced by A2LParser#leftShift.
	VisitLeftShift(ctx *LeftShiftContext) interface{}

	// Visit a parse tree produced by A2LParser#rightShift.
	VisitRightShift(ctx *RightShiftContext) interface{}

	// Visit a parse tree produced by A2LParser#signExtend.
	VisitSignExtend(ctx *SignExtendContext) interface{}

	// Visit a parse tree produced by A2LParser#ecuAddress.
	VisitEcuAddress(ctx *EcuAddressContext) interface{}

	// Visit a parse tree produced by A2LParser#errorMask.
	VisitErrorMask(ctx *ErrorMaskContext) interface{}

	// Visit a parse tree produced by A2LParser#layout.
	VisitLayout(ctx *LayoutContext) interface{}

	// Visit a parse tree produced by A2LParser#readWrite.
	VisitReadWrite(ctx *ReadWriteContext) interface{}

	// Visit a parse tree produced by A2LParser#virtual.
	VisitVirtual(ctx *VirtualContext) interface{}

	// Visit a parse tree produced by A2LParser#modCommon.
	VisitModCommon(ctx *ModCommonContext) interface{}

	// Visit a parse tree produced by A2LParser#dataSize.
	VisitDataSize(ctx *DataSizeContext) interface{}

	// Visit a parse tree produced by A2LParser#sRecLayout.
	VisitSRecLayout(ctx *SRecLayoutContext) interface{}

	// Visit a parse tree produced by A2LParser#modPar.
	VisitModPar(ctx *ModParContext) interface{}

	// Visit a parse tree produced by A2LParser#addrEpk.
	VisitAddrEpk(ctx *AddrEpkContext) interface{}

	// Visit a parse tree produced by A2LParser#calibrationMethod.
	VisitCalibrationMethod(ctx *CalibrationMethodContext) interface{}

	// Visit a parse tree produced by A2LParser#calibrationHandle.
	VisitCalibrationHandle(ctx *CalibrationHandleContext) interface{}

	// Visit a parse tree produced by A2LParser#cpuType.
	VisitCpuType(ctx *CpuTypeContext) interface{}

	// Visit a parse tree produced by A2LParser#customer.
	VisitCustomer(ctx *CustomerContext) interface{}

	// Visit a parse tree produced by A2LParser#customerNo.
	VisitCustomerNo(ctx *CustomerNoContext) interface{}

	// Visit a parse tree produced by A2LParser#ecu.
	VisitEcu(ctx *EcuContext) interface{}

	// Visit a parse tree produced by A2LParser#ecuCalibrationOffset.
	VisitEcuCalibrationOffset(ctx *EcuCalibrationOffsetContext) interface{}

	// Visit a parse tree produced by A2LParser#epk.
	VisitEpk(ctx *EpkContext) interface{}

	// Visit a parse tree produced by A2LParser#memoryLayout.
	VisitMemoryLayout(ctx *MemoryLayoutContext) interface{}

	// Visit a parse tree produced by A2LParser#memorySegment.
	VisitMemorySegment(ctx *MemorySegmentContext) interface{}

	// Visit a parse tree produced by A2LParser#noOfInterfaces.
	VisitNoOfInterfaces(ctx *NoOfInterfacesContext) interface{}

	// Visit a parse tree produced by A2LParser#phoneNo.
	VisitPhoneNo(ctx *PhoneNoContext) interface{}

	// Visit a parse tree produced by A2LParser#supplier.
	VisitSupplier(ctx *SupplierContext) interface{}

	// Visit a parse tree produced by A2LParser#systemConstant.
	VisitSystemConstant(ctx *SystemConstantContext) interface{}

	// Visit a parse tree produced by A2LParser#user.
	VisitUser(ctx *UserContext) interface{}

	// Visit a parse tree produced by A2LParser#recordLayout.
	VisitRecordLayout(ctx *RecordLayoutContext) interface{}

	// Visit a parse tree produced by A2LParser#axisPtsX.
	VisitAxisPtsX(ctx *AxisPtsXContext) interface{}

	// Visit a parse tree produced by A2LParser#axisPtsY.
	VisitAxisPtsY(ctx *AxisPtsYContext) interface{}

	// Visit a parse tree produced by A2LParser#axisPtsZ.
	VisitAxisPtsZ(ctx *AxisPtsZContext) interface{}

	// Visit a parse tree produced by A2LParser#axisPts4.
	VisitAxisPts4(ctx *AxisPts4Context) interface{}

	// Visit a parse tree produced by A2LParser#axisPts5.
	VisitAxisPts5(ctx *AxisPts5Context) interface{}

	// Visit a parse tree produced by A2LParser#axisRescaleX.
	VisitAxisRescaleX(ctx *AxisRescaleXContext) interface{}

	// Visit a parse tree produced by A2LParser#axisRescaleY.
	VisitAxisRescaleY(ctx *AxisRescaleYContext) interface{}

	// Visit a parse tree produced by A2LParser#axisRescaleZ.
	VisitAxisRescaleZ(ctx *AxisRescaleZContext) interface{}

	// Visit a parse tree produced by A2LParser#axisRescale4.
	VisitAxisRescale4(ctx *AxisRescale4Context) interface{}

	// Visit a parse tree produced by A2LParser#axisRescale5.
	VisitAxisRescale5(ctx *AxisRescale5Context) interface{}

	// Visit a parse tree produced by A2LParser#distOpX.
	VisitDistOpX(ctx *DistOpXContext) interface{}

	// Visit a parse tree produced by A2LParser#distOpY.
	VisitDistOpY(ctx *DistOpYContext) interface{}

	// Visit a parse tree produced by A2LParser#distOpZ.
	VisitDistOpZ(ctx *DistOpZContext) interface{}

	// Visit a parse tree produced by A2LParser#distOp4.
	VisitDistOp4(ctx *DistOp4Context) interface{}

	// Visit a parse tree produced by A2LParser#distOp5.
	VisitDistOp5(ctx *DistOp5Context) interface{}

	// Visit a parse tree produced by A2LParser#fixNoAxisPtsX.
	VisitFixNoAxisPtsX(ctx *FixNoAxisPtsXContext) interface{}

	// Visit a parse tree produced by A2LParser#fixNoAxisPtsY.
	VisitFixNoAxisPtsY(ctx *FixNoAxisPtsYContext) interface{}

	// Visit a parse tree produced by A2LParser#fixNoAxisPtsZ.
	VisitFixNoAxisPtsZ(ctx *FixNoAxisPtsZContext) interface{}

	// Visit a parse tree produced by A2LParser#fixNoAxisPts4.
	VisitFixNoAxisPts4(ctx *FixNoAxisPts4Context) interface{}

	// Visit a parse tree produced by A2LParser#fixNoAxisPts5.
	VisitFixNoAxisPts5(ctx *FixNoAxisPts5Context) interface{}

	// Visit a parse tree produced by A2LParser#fncValues.
	VisitFncValues(ctx *FncValuesContext) interface{}

	// Visit a parse tree produced by A2LParser#identification.
	VisitIdentification(ctx *IdentificationContext) interface{}

	// Visit a parse tree produced by A2LParser#noAxisPtsX.
	VisitNoAxisPtsX(ctx *NoAxisPtsXContext) interface{}

	// Visit a parse tree produced by A2LParser#noAxisPtsY.
	VisitNoAxisPtsY(ctx *NoAxisPtsYContext) interface{}

	// Visit a parse tree produced by A2LParser#noAxisPtsZ.
	VisitNoAxisPtsZ(ctx *NoAxisPtsZContext) interface{}

	// Visit a parse tree produced by A2LParser#noAxisPts4.
	VisitNoAxisPts4(ctx *NoAxisPts4Context) interface{}

	// Visit a parse tree produced by A2LParser#noAxisPts5.
	VisitNoAxisPts5(ctx *NoAxisPts5Context) interface{}

	// Visit a parse tree produced by A2LParser#staticRecordLayout.
	VisitStaticRecordLayout(ctx *StaticRecordLayoutContext) interface{}

	// Visit a parse tree produced by A2LParser#noRescaleX.
	VisitNoRescaleX(ctx *NoRescaleXContext) interface{}

	// Visit a parse tree produced by A2LParser#noRescaleY.
	VisitNoRescaleY(ctx *NoRescaleYContext) interface{}

	// Visit a parse tree produced by A2LParser#noRescaleZ.
	VisitNoRescaleZ(ctx *NoRescaleZContext) interface{}

	// Visit a parse tree produced by A2LParser#noRescale4.
	VisitNoRescale4(ctx *NoRescale4Context) interface{}

	// Visit a parse tree produced by A2LParser#noRescale5.
	VisitNoRescale5(ctx *NoRescale5Context) interface{}

	// Visit a parse tree produced by A2LParser#offsetX.
	VisitOffsetX(ctx *OffsetXContext) interface{}

	// Visit a parse tree produced by A2LParser#offsetY.
	VisitOffsetY(ctx *OffsetYContext) interface{}

	// Visit a parse tree produced by A2LParser#offsetZ.
	VisitOffsetZ(ctx *OffsetZContext) interface{}

	// Visit a parse tree produced by A2LParser#offset4.
	VisitOffset4(ctx *Offset4Context) interface{}

	// Visit a parse tree produced by A2LParser#offset5.
	VisitOffset5(ctx *Offset5Context) interface{}

	// Visit a parse tree produced by A2LParser#reserved.
	VisitReserved(ctx *ReservedContext) interface{}

	// Visit a parse tree produced by A2LParser#ripAddrW.
	VisitRipAddrW(ctx *RipAddrWContext) interface{}

	// Visit a parse tree produced by A2LParser#ripAddrX.
	VisitRipAddrX(ctx *RipAddrXContext) interface{}

	// Visit a parse tree produced by A2LParser#ripAddrY.
	VisitRipAddrY(ctx *RipAddrYContext) interface{}

	// Visit a parse tree produced by A2LParser#ripAddrZ.
	VisitRipAddrZ(ctx *RipAddrZContext) interface{}

	// Visit a parse tree produced by A2LParser#ripAddr4.
	VisitRipAddr4(ctx *RipAddr4Context) interface{}

	// Visit a parse tree produced by A2LParser#ripAddr5.
	VisitRipAddr5(ctx *RipAddr5Context) interface{}

	// Visit a parse tree produced by A2LParser#shiftOpX.
	VisitShiftOpX(ctx *ShiftOpXContext) interface{}

	// Visit a parse tree produced by A2LParser#shiftOpY.
	VisitShiftOpY(ctx *ShiftOpYContext) interface{}

	// Visit a parse tree produced by A2LParser#shiftOpZ.
	VisitShiftOpZ(ctx *ShiftOpZContext) interface{}

	// Visit a parse tree produced by A2LParser#shiftOp4.
	VisitShiftOp4(ctx *ShiftOp4Context) interface{}

	// Visit a parse tree produced by A2LParser#shiftOp5.
	VisitShiftOp5(ctx *ShiftOp5Context) interface{}

	// Visit a parse tree produced by A2LParser#srcAddrX.
	VisitSrcAddrX(ctx *SrcAddrXContext) interface{}

	// Visit a parse tree produced by A2LParser#srcAddrY.
	VisitSrcAddrY(ctx *SrcAddrYContext) interface{}

	// Visit a parse tree produced by A2LParser#srcAddrZ.
	VisitSrcAddrZ(ctx *SrcAddrZContext) interface{}

	// Visit a parse tree produced by A2LParser#srcAddr4.
	VisitSrcAddr4(ctx *SrcAddr4Context) interface{}

	// Visit a parse tree produced by A2LParser#srcAddr5.
	VisitSrcAddr5(ctx *SrcAddr5Context) interface{}

	// Visit a parse tree produced by A2LParser#typedefCharacteristic.
	VisitTypedefCharacteristic(ctx *TypedefCharacteristicContext) interface{}

	// Visit a parse tree produced by A2LParser#typedefMeasurement.
	VisitTypedefMeasurement(ctx *TypedefMeasurementContext) interface{}

	// Visit a parse tree produced by A2LParser#typedefStructure.
	VisitTypedefStructure(ctx *TypedefStructureContext) interface{}

	// Visit a parse tree produced by A2LParser#structureComponent.
	VisitStructureComponent(ctx *StructureComponentContext) interface{}

	// Visit a parse tree produced by A2LParser#unit.
	VisitUnit(ctx *UnitContext) interface{}

	// Visit a parse tree produced by A2LParser#siExponents.
	VisitSiExponents(ctx *SiExponentsContext) interface{}

	// Visit a parse tree produced by A2LParser#unitConversion.
	VisitUnitConversion(ctx *UnitConversionContext) interface{}

	// Visit a parse tree produced by A2LParser#userRights.
	VisitUserRights(ctx *UserRightsContext) interface{}

	// Visit a parse tree produced by A2LParser#refGroup.
	VisitRefGroup(ctx *RefGroupContext) interface{}

	// Visit a parse tree produced by A2LParser#variantCoding.
	VisitVariantCoding(ctx *VariantCodingContext) interface{}

	// Visit a parse tree produced by A2LParser#varCharacteristic.
	VisitVarCharacteristic(ctx *VarCharacteristicContext) interface{}

	// Visit a parse tree produced by A2LParser#varAddress.
	VisitVarAddress(ctx *VarAddressContext) interface{}

	// Visit a parse tree produced by A2LParser#varCriterion.
	VisitVarCriterion(ctx *VarCriterionContext) interface{}

	// Visit a parse tree produced by A2LParser#varMeasurement.
	VisitVarMeasurement(ctx *VarMeasurementContext) interface{}

	// Visit a parse tree produced by A2LParser#varSelectionCharacteristic.
	VisitVarSelectionCharacteristic(ctx *VarSelectionCharacteristicContext) interface{}

	// Visit a parse tree produced by A2LParser#varForbiddenComb.
	VisitVarForbiddenComb(ctx *VarForbiddenCombContext) interface{}

	// Visit a parse tree produced by A2LParser#varNaming.
	VisitVarNaming(ctx *VarNamingContext) interface{}

	// Visit a parse tree produced by A2LParser#varSeparator.
	VisitVarSeparator(ctx *VarSeparatorContext) interface{}

	// Visit a parse tree produced by A2LParser#integerValue.
	VisitIntegerValue(ctx *IntegerValueContext) interface{}

	// Visit a parse tree produced by A2LParser#numericValue.
	VisitNumericValue(ctx *NumericValueContext) interface{}

	// Visit a parse tree produced by A2LParser#stringValue.
	VisitStringValue(ctx *StringValueContext) interface{}

	// Visit a parse tree produced by A2LParser#identifierValue.
	VisitIdentifierValue(ctx *IdentifierValueContext) interface{}

	// Visit a parse tree produced by A2LParser#partialIdentifier.
	VisitPartialIdentifier(ctx *PartialIdentifierContext) interface{}

	// Visit a parse tree produced by A2LParser#arraySpecifier.
	VisitArraySpecifier(ctx *ArraySpecifierContext) interface{}

	// Visit a parse tree produced by A2LParser#dataType.
	VisitDataType(ctx *DataTypeContext) interface{}

	// Visit a parse tree produced by A2LParser#datasize.
	VisitDatasize(ctx *DatasizeContext) interface{}

	// Visit a parse tree produced by A2LParser#addrtype.
	VisitAddrtype(ctx *AddrtypeContext) interface{}

	// Visit a parse tree produced by A2LParser#byteOrderValue.
	VisitByteOrderValue(ctx *ByteOrderValueContext) interface{}

	// Visit a parse tree produced by A2LParser#indexorder.
	VisitIndexorder(ctx *IndexorderContext) interface{}

	// Visit a parse tree produced by A2LParser#linkType.
	VisitLinkType(ctx *LinkTypeContext) interface{}

	// Visit a parse tree produced by A2LParser#a2ml.
	VisitA2ml(ctx *A2mlContext) interface{}

	// Visit a parse tree produced by A2LParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by A2LParser#typeDefinition.
	VisitTypeDefinition(ctx *TypeDefinitionContext) interface{}

	// Visit a parse tree produced by A2LParser#a2mlTypeName.
	VisitA2mlTypeName(ctx *A2mlTypeNameContext) interface{}

	// Visit a parse tree produced by A2LParser#predefinedTypeName.
	VisitPredefinedTypeName(ctx *PredefinedTypeNameContext) interface{}

	// Visit a parse tree produced by A2LParser#blockDefinition.
	VisitBlockDefinition(ctx *BlockDefinitionContext) interface{}

	// Visit a parse tree produced by A2LParser#enumTypeName.
	VisitEnumTypeName(ctx *EnumTypeNameContext) interface{}

	// Visit a parse tree produced by A2LParser#enumeratorList.
	VisitEnumeratorList(ctx *EnumeratorListContext) interface{}

	// Visit a parse tree produced by A2LParser#enumerator.
	VisitEnumerator(ctx *EnumeratorContext) interface{}

	// Visit a parse tree produced by A2LParser#structTypeName.
	VisitStructTypeName(ctx *StructTypeNameContext) interface{}

	// Visit a parse tree produced by A2LParser#structMember.
	VisitStructMember(ctx *StructMemberContext) interface{}

	// Visit a parse tree produced by A2LParser#member.
	VisitMember(ctx *MemberContext) interface{}

	// Visit a parse tree produced by A2LParser#taggedStructTypeName.
	VisitTaggedStructTypeName(ctx *TaggedStructTypeNameContext) interface{}

	// Visit a parse tree produced by A2LParser#taggedStructMember.
	VisitTaggedStructMember(ctx *TaggedStructMemberContext) interface{}

	// Visit a parse tree produced by A2LParser#taggedStructDefinition.
	VisitTaggedStructDefinition(ctx *TaggedStructDefinitionContext) interface{}

	// Visit a parse tree produced by A2LParser#taggedUnionTypeName.
	VisitTaggedUnionTypeName(ctx *TaggedUnionTypeNameContext) interface{}

	// Visit a parse tree produced by A2LParser#taggedUnionMember.
	VisitTaggedUnionMember(ctx *TaggedUnionMemberContext) interface{}

	// Visit a parse tree produced by A2LParser#tagValue.
	VisitTagValue(ctx *TagValueContext) interface{}

	// Visit a parse tree produced by A2LParser#ifData.
	VisitIfData(ctx *IfDataContext) interface{}

	// Visit a parse tree produced by A2LParser#genericParameter.
	VisitGenericParameter(ctx *GenericParameterContext) interface{}

	// Visit a parse tree produced by A2LParser#genericNode.
	VisitGenericNode(ctx *GenericNodeContext) interface{}
}
