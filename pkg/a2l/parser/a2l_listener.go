// Code generated from grammar/A2L.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // A2L

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A2LListener is a complete listener for a parse tree produced by A2LParser.
type A2LListener interface {
	antlr.ParseTreeListener

	// EnterA2lFile is called when entering the a2lFile production.
	EnterA2lFile(c *A2lFileContext)

	// EnterAlignmentByte is called when entering the alignmentByte production.
	EnterAlignmentByte(c *AlignmentByteContext)

	// EnterAlignmentFloat16Ieee is called when entering the alignmentFloat16Ieee production.
	EnterAlignmentFloat16Ieee(c *AlignmentFloat16IeeeContext)

	// EnterAlignmentFloat32Ieee is called when entering the alignmentFloat32Ieee production.
	EnterAlignmentFloat32Ieee(c *AlignmentFloat32IeeeContext)

	// EnterAlignmentFloat64Ieee is called when entering the alignmentFloat64Ieee production.
	EnterAlignmentFloat64Ieee(c *AlignmentFloat64IeeeContext)

	// EnterAlignmentInt64 is called when entering the alignmentInt64 production.
	EnterAlignmentInt64(c *AlignmentInt64Context)

	// EnterAlignmentLong is called when entering the alignmentLong production.
	EnterAlignmentLong(c *AlignmentLongContext)

	// EnterAlignmentWord is called when entering the alignmentWord production.
	EnterAlignmentWord(c *AlignmentWordContext)

	// EnterAnnotation is called when entering the annotation production.
	EnterAnnotation(c *AnnotationContext)

	// EnterAnnotationLabel is called when entering the annotationLabel production.
	EnterAnnotationLabel(c *AnnotationLabelContext)

	// EnterAnnotationOrigin is called when entering the annotationOrigin production.
	EnterAnnotationOrigin(c *AnnotationOriginContext)

	// EnterAnnotationText is called when entering the annotationText production.
	EnterAnnotationText(c *AnnotationTextContext)

	// EnterBitMask is called when entering the bitMask production.
	EnterBitMask(c *BitMaskContext)

	// EnterByteOrder is called when entering the byteOrder production.
	EnterByteOrder(c *ByteOrderContext)

	// EnterCalibrationAccess is called when entering the calibrationAccess production.
	EnterCalibrationAccess(c *CalibrationAccessContext)

	// EnterDefaultValue is called when entering the defaultValue production.
	EnterDefaultValue(c *DefaultValueContext)

	// EnterDeposit is called when entering the deposit production.
	EnterDeposit(c *DepositContext)

	// EnterDiscrete is called when entering the discrete production.
	EnterDiscrete(c *DiscreteContext)

	// EnterDisplayIdentifier is called when entering the displayIdentifier production.
	EnterDisplayIdentifier(c *DisplayIdentifierContext)

	// EnterEcuAddressExtension is called when entering the ecuAddressExtension production.
	EnterEcuAddressExtension(c *EcuAddressExtensionContext)

	// EnterExtendedLimits is called when entering the extendedLimits production.
	EnterExtendedLimits(c *ExtendedLimitsContext)

	// EnterFormat is called when entering the format production.
	EnterFormat(c *FormatContext)

	// EnterFunctionList is called when entering the functionList production.
	EnterFunctionList(c *FunctionListContext)

	// EnterGuardRails is called when entering the guardRails production.
	EnterGuardRails(c *GuardRailsContext)

	// EnterMatrixDim is called when entering the matrixDim production.
	EnterMatrixDim(c *MatrixDimContext)

	// EnterMaxRefresh is called when entering the maxRefresh production.
	EnterMaxRefresh(c *MaxRefreshContext)

	// EnterMonotony is called when entering the monotony production.
	EnterMonotony(c *MonotonyContext)

	// EnterPhysUnit is called when entering the physUnit production.
	EnterPhysUnit(c *PhysUnitContext)

	// EnterReadOnly is called when entering the readOnly production.
	EnterReadOnly(c *ReadOnlyContext)

	// EnterRefCharacteristic is called when entering the refCharacteristic production.
	EnterRefCharacteristic(c *RefCharacteristicContext)

	// EnterRefMemorySegment is called when entering the refMemorySegment production.
	EnterRefMemorySegment(c *RefMemorySegmentContext)

	// EnterRefUnit is called when entering the refUnit production.
	EnterRefUnit(c *RefUnitContext)

	// EnterStepSize is called when entering the stepSize production.
	EnterStepSize(c *StepSizeContext)

	// EnterSymbolLink is called when entering the symbolLink production.
	EnterSymbolLink(c *SymbolLinkContext)

	// EnterVersion is called when entering the version production.
	EnterVersion(c *VersionContext)

	// EnterAsap2Version is called when entering the asap2Version production.
	EnterAsap2Version(c *Asap2VersionContext)

	// EnterA2mlVersion is called when entering the a2mlVersion production.
	EnterA2mlVersion(c *A2mlVersionContext)

	// EnterProject is called when entering the project production.
	EnterProject(c *ProjectContext)

	// EnterHeader is called when entering the header production.
	EnterHeader(c *HeaderContext)

	// EnterProjectNo is called when entering the projectNo production.
	EnterProjectNo(c *ProjectNoContext)

	// EnterModule is called when entering the module production.
	EnterModule(c *ModuleContext)

	// EnterAxisPts is called when entering the axisPts production.
	EnterAxisPts(c *AxisPtsContext)

	// EnterCharacteristic is called when entering the characteristic production.
	EnterCharacteristic(c *CharacteristicContext)

	// EnterAxisDescr is called when entering the axisDescr production.
	EnterAxisDescr(c *AxisDescrContext)

	// EnterAxisPtsRef is called when entering the axisPtsRef production.
	EnterAxisPtsRef(c *AxisPtsRefContext)

	// EnterCurveAxisRef is called when entering the curveAxisRef production.
	EnterCurveAxisRef(c *CurveAxisRefContext)

	// EnterFixAxisPar is called when entering the fixAxisPar production.
	EnterFixAxisPar(c *FixAxisParContext)

	// EnterFixAxisParDist is called when entering the fixAxisParDist production.
	EnterFixAxisParDist(c *FixAxisParDistContext)

	// EnterFixAxisParList is called when entering the fixAxisParList production.
	EnterFixAxisParList(c *FixAxisParListContext)

	// EnterMaxGrad is called when entering the maxGrad production.
	EnterMaxGrad(c *MaxGradContext)

	// EnterComparisonQuantity is called when entering the comparisonQuantity production.
	EnterComparisonQuantity(c *ComparisonQuantityContext)

	// EnterDependentCharacteristic is called when entering the dependentCharacteristic production.
	EnterDependentCharacteristic(c *DependentCharacteristicContext)

	// EnterMapList is called when entering the mapList production.
	EnterMapList(c *MapListContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterVirtualCharacteristic is called when entering the virtualCharacteristic production.
	EnterVirtualCharacteristic(c *VirtualCharacteristicContext)

	// EnterCompuMethod is called when entering the compuMethod production.
	EnterCompuMethod(c *CompuMethodContext)

	// EnterCoeffs is called when entering the coeffs production.
	EnterCoeffs(c *CoeffsContext)

	// EnterCoeffsLinear is called when entering the coeffsLinear production.
	EnterCoeffsLinear(c *CoeffsLinearContext)

	// EnterCompuTabRef is called when entering the compuTabRef production.
	EnterCompuTabRef(c *CompuTabRefContext)

	// EnterFormula is called when entering the formula production.
	EnterFormula(c *FormulaContext)

	// EnterFormulaInv is called when entering the formulaInv production.
	EnterFormulaInv(c *FormulaInvContext)

	// EnterStatusStringRef is called when entering the statusStringRef production.
	EnterStatusStringRef(c *StatusStringRefContext)

	// EnterCompuTab is called when entering the compuTab production.
	EnterCompuTab(c *CompuTabContext)

	// EnterDefaultValueNumeric is called when entering the defaultValueNumeric production.
	EnterDefaultValueNumeric(c *DefaultValueNumericContext)

	// EnterCompuVTab is called when entering the compuVTab production.
	EnterCompuVTab(c *CompuVTabContext)

	// EnterCompuVTabRange is called when entering the compuVTabRange production.
	EnterCompuVTabRange(c *CompuVTabRangeContext)

	// EnterFrame is called when entering the frame production.
	EnterFrame(c *FrameContext)

	// EnterFrameMeasurement is called when entering the frameMeasurement production.
	EnterFrameMeasurement(c *FrameMeasurementContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterDefCharacteristic is called when entering the defCharacteristic production.
	EnterDefCharacteristic(c *DefCharacteristicContext)

	// EnterFunctionVersion is called when entering the functionVersion production.
	EnterFunctionVersion(c *FunctionVersionContext)

	// EnterInMeasurement is called when entering the inMeasurement production.
	EnterInMeasurement(c *InMeasurementContext)

	// EnterLocMeasurement is called when entering the locMeasurement production.
	EnterLocMeasurement(c *LocMeasurementContext)

	// EnterOutMeasurement is called when entering the outMeasurement production.
	EnterOutMeasurement(c *OutMeasurementContext)

	// EnterSubFunction is called when entering the subFunction production.
	EnterSubFunction(c *SubFunctionContext)

	// EnterGroup is called when entering the group production.
	EnterGroup(c *GroupContext)

	// EnterRefMeasurement is called when entering the refMeasurement production.
	EnterRefMeasurement(c *RefMeasurementContext)

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterSubGroup is called when entering the subGroup production.
	EnterSubGroup(c *SubGroupContext)

	// EnterInstance is called when entering the instance production.
	EnterInstance(c *InstanceContext)

	// EnterMeasurement is called when entering the measurement production.
	EnterMeasurement(c *MeasurementContext)

	// EnterArraySize is called when entering the arraySize production.
	EnterArraySize(c *ArraySizeContext)

	// EnterBitOperation is called when entering the bitOperation production.
	EnterBitOperation(c *BitOperationContext)

	// EnterLeftShift is called when entering the leftShift production.
	EnterLeftShift(c *LeftShiftContext)

	// EnterRightShift is called when entering the rightShift production.
	EnterRightShift(c *RightShiftContext)

	// EnterSignExtend is called when entering the signExtend production.
	EnterSignExtend(c *SignExtendContext)

	// EnterEcuAddress is called when entering the ecuAddress production.
	EnterEcuAddress(c *EcuAddressContext)

	// EnterErrorMask is called when entering the errorMask production.
	EnterErrorMask(c *ErrorMaskContext)

	// EnterLayout is called when entering the layout production.
	EnterLayout(c *LayoutContext)

	// EnterReadWrite is called when entering the readWrite production.
	EnterReadWrite(c *ReadWriteContext)

	// EnterVirtual is called when entering the virtual production.
	EnterVirtual(c *VirtualContext)

	// EnterModCommon is called when entering the modCommon production.
	EnterModCommon(c *ModCommonContext)

	// EnterDataSize is called when entering the dataSize production.
	EnterDataSize(c *DataSizeContext)

	// EnterSRecLayout is called when entering the sRecLayout production.
	EnterSRecLayout(c *SRecLayoutContext)

	// EnterModPar is called when entering the modPar production.
	EnterModPar(c *ModParContext)

	// EnterAddrEpk is called when entering the addrEpk production.
	EnterAddrEpk(c *AddrEpkContext)

	// EnterCalibrationMethod is called when entering the calibrationMethod production.
	EnterCalibrationMethod(c *CalibrationMethodContext)

	// EnterCalibrationHandle is called when entering the calibrationHandle production.
	EnterCalibrationHandle(c *CalibrationHandleContext)

	// EnterCpuType is called when entering the cpuType production.
	EnterCpuType(c *CpuTypeContext)

	// EnterCustomer is called when entering the customer production.
	EnterCustomer(c *CustomerContext)

	// EnterCustomerNo is called when entering the customerNo production.
	EnterCustomerNo(c *CustomerNoContext)

	// EnterEcu is called when entering the ecu production.
	EnterEcu(c *EcuContext)

	// EnterEcuCalibrationOffset is called when entering the ecuCalibrationOffset production.
	EnterEcuCalibrationOffset(c *EcuCalibrationOffsetContext)

	// EnterEpk is called when entering the epk production.
	EnterEpk(c *EpkContext)

	// EnterMemoryLayout is called when entering the memoryLayout production.
	EnterMemoryLayout(c *MemoryLayoutContext)

	// EnterMemorySegment is called when entering the memorySegment production.
	EnterMemorySegment(c *MemorySegmentContext)

	// EnterNoOfInterfaces is called when entering the noOfInterfaces production.
	EnterNoOfInterfaces(c *NoOfInterfacesContext)

	// EnterPhoneNo is called when entering the phoneNo production.
	EnterPhoneNo(c *PhoneNoContext)

	// EnterSupplier is called when entering the supplier production.
	EnterSupplier(c *SupplierContext)

	// EnterSystemConstant is called when entering the systemConstant production.
	EnterSystemConstant(c *SystemConstantContext)

	// EnterUser is called when entering the user production.
	EnterUser(c *UserContext)

	// EnterRecordLayout is called when entering the recordLayout production.
	EnterRecordLayout(c *RecordLayoutContext)

	// EnterAxisPtsX is called when entering the axisPtsX production.
	EnterAxisPtsX(c *AxisPtsXContext)

	// EnterAxisPtsY is called when entering the axisPtsY production.
	EnterAxisPtsY(c *AxisPtsYContext)

	// EnterAxisPtsZ is called when entering the axisPtsZ production.
	EnterAxisPtsZ(c *AxisPtsZContext)

	// EnterAxisPts4 is called when entering the axisPts4 production.
	EnterAxisPts4(c *AxisPts4Context)

	// EnterAxisPts5 is called when entering the axisPts5 production.
	EnterAxisPts5(c *AxisPts5Context)

	// EnterAxisRescaleX is called when entering the axisRescaleX production.
	EnterAxisRescaleX(c *AxisRescaleXContext)

	// EnterAxisRescaleY is called when entering the axisRescaleY production.
	EnterAxisRescaleY(c *AxisRescaleYContext)

	// EnterAxisRescaleZ is called when entering the axisRescaleZ production.
	EnterAxisRescaleZ(c *AxisRescaleZContext)

	// EnterAxisRescale4 is called when entering the axisRescale4 production.
	EnterAxisRescale4(c *AxisRescale4Context)

	// EnterAxisRescale5 is called when entering the axisRescale5 production.
	EnterAxisRescale5(c *AxisRescale5Context)

	// EnterDistOpX is called when entering the distOpX production.
	EnterDistOpX(c *DistOpXContext)

	// EnterDistOpY is called when entering the distOpY production.
	EnterDistOpY(c *DistOpYContext)

	// EnterDistOpZ is called when entering the distOpZ production.
	EnterDistOpZ(c *DistOpZContext)

	// EnterDistOp4 is called when entering the distOp4 production.
	EnterDistOp4(c *DistOp4Context)

	// EnterDistOp5 is called when entering the distOp5 production.
	EnterDistOp5(c *DistOp5Context)

	// EnterFixNoAxisPtsX is called when entering the fixNoAxisPtsX production.
	EnterFixNoAxisPtsX(c *FixNoAxisPtsXContext)

	// EnterFixNoAxisPtsY is called when entering the fixNoAxisPtsY production.
	EnterFixNoAxisPtsY(c *FixNoAxisPtsYContext)

	// EnterFixNoAxisPtsZ is called when entering the fixNoAxisPtsZ production.
	EnterFixNoAxisPtsZ(c *FixNoAxisPtsZContext)

	// EnterFixNoAxisPts4 is called when entering the fixNoAxisPts4 production.
	EnterFixNoAxisPts4(c *FixNoAxisPts4Context)

	// EnterFixNoAxisPts5 is called when entering the fixNoAxisPts5 production.
	EnterFixNoAxisPts5(c *FixNoAxisPts5Context)

	// EnterFncValues is called when entering the fncValues production.
	EnterFncValues(c *FncValuesContext)

	// EnterIdentification is called when entering the identification production.
	EnterIdentification(c *IdentificationContext)

	// EnterNoAxisPtsX is called when entering the noAxisPtsX production.
	EnterNoAxisPtsX(c *NoAxisPtsXContext)

	// EnterNoAxisPtsY is called when entering the noAxisPtsY production.
	EnterNoAxisPtsY(c *NoAxisPtsYContext)

	// EnterNoAxisPtsZ is called when entering the noAxisPtsZ production.
	EnterNoAxisPtsZ(c *NoAxisPtsZContext)

	// EnterNoAxisPts4 is called when entering the noAxisPts4 production.
	EnterNoAxisPts4(c *NoAxisPts4Context)

	// EnterNoAxisPts5 is called when entering the noAxisPts5 production.
	EnterNoAxisPts5(c *NoAxisPts5Context)

	// EnterStaticRecordLayout is called when entering the staticRecordLayout production.
	EnterStaticRecordLayout(c *StaticRecordLayoutContext)

	// EnterNoRescaleX is called when entering the noRescaleX production.
	EnterNoRescaleX(c *NoRescaleXContext)

	// EnterNoRescaleY is called when entering the noRescaleY production.
	EnterNoRescaleY(c *NoRescaleYContext)

	// EnterNoRescaleZ is called when entering the noRescaleZ production.
	EnterNoRescaleZ(c *NoRescaleZContext)

	// EnterNoRescale4 is called when entering the noRescale4 production.
	EnterNoRescale4(c *NoRescale4Context)

	// EnterNoRescale5 is called when entering the noRescale5 production.
	EnterNoRescale5(c *NoRescale5Context)

	// EnterOffsetX is called when entering the offsetX production.
	EnterOffsetX(c *OffsetXContext)

	// EnterOffsetY is called when entering the offsetY production.
	EnterOffsetY(c *OffsetYContext)

	// EnterOffsetZ is called when entering the offsetZ production.
	EnterOffsetZ(c *OffsetZContext)

	// EnterOffset4 is called when entering the offset4 production.
	EnterOffset4(c *Offset4Context)

	// EnterOffset5 is called when entering the offset5 production.
	EnterOffset5(c *Offset5Context)

	// EnterReserved is called when entering the reserved production.
	EnterReserved(c *ReservedContext)

	// EnterRipAddrW is called when entering the ripAddrW production.
	EnterRipAddrW(c *RipAddrWContext)

	// EnterRipAddrX is called when entering the ripAddrX production.
	EnterRipAddrX(c *RipAddrXContext)

	// EnterRipAddrY is called when entering the ripAddrY production.
	EnterRipAddrY(c *RipAddrYContext)

	// EnterRipAddrZ is called when entering the ripAddrZ production.
	EnterRipAddrZ(c *RipAddrZContext)

	// EnterRipAddr4 is called when entering the ripAddr4 production.
	EnterRipAddr4(c *RipAddr4Context)

	// EnterRipAddr5 is called when entering the ripAddr5 production.
	EnterRipAddr5(c *RipAddr5Context)

	// EnterShiftOpX is called when entering the shiftOpX production.
	EnterShiftOpX(c *ShiftOpXContext)

	// EnterShiftOpY is called when entering the shiftOpY production.
	EnterShiftOpY(c *ShiftOpYContext)

	// EnterShiftOpZ is called when entering the shiftOpZ production.
	EnterShiftOpZ(c *ShiftOpZContext)

	// EnterShiftOp4 is called when entering the shiftOp4 production.
	EnterShiftOp4(c *ShiftOp4Context)

	// EnterShiftOp5 is called when entering the shiftOp5 production.
	EnterShiftOp5(c *ShiftOp5Context)

	// EnterSrcAddrX is called when entering the srcAddrX production.
	EnterSrcAddrX(c *SrcAddrXContext)

	// EnterSrcAddrY is called when entering the srcAddrY production.
	EnterSrcAddrY(c *SrcAddrYContext)

	// EnterSrcAddrZ is called when entering the srcAddrZ production.
	EnterSrcAddrZ(c *SrcAddrZContext)

	// EnterSrcAddr4 is called when entering the srcAddr4 production.
	EnterSrcAddr4(c *SrcAddr4Context)

	// EnterSrcAddr5 is called when entering the srcAddr5 production.
	EnterSrcAddr5(c *SrcAddr5Context)

	// EnterTypedefCharacteristic is called when entering the typedefCharacteristic production.
	EnterTypedefCharacteristic(c *TypedefCharacteristicContext)

	// EnterTypedefMeasurement is called when entering the typedefMeasurement production.
	EnterTypedefMeasurement(c *TypedefMeasurementContext)

	// EnterTypedefStructure is called when entering the typedefStructure production.
	EnterTypedefStructure(c *TypedefStructureContext)

	// EnterStructureComponent is called when entering the structureComponent production.
	EnterStructureComponent(c *StructureComponentContext)

	// EnterUnit is called when entering the unit production.
	EnterUnit(c *UnitContext)

	// EnterSiExponents is called when entering the siExponents production.
	EnterSiExponents(c *SiExponentsContext)

	// EnterUnitConversion is called when entering the unitConversion production.
	EnterUnitConversion(c *UnitConversionContext)

	// EnterUserRights is called when entering the userRights production.
	EnterUserRights(c *UserRightsContext)

	// EnterRefGroup is called when entering the refGroup production.
	EnterRefGroup(c *RefGroupContext)

	// EnterVariantCoding is called when entering the variantCoding production.
	EnterVariantCoding(c *VariantCodingContext)

	// EnterVarCharacteristic is called when entering the varCharacteristic production.
	EnterVarCharacteristic(c *VarCharacteristicContext)

	// EnterVarAddress is called when entering the varAddress production.
	EnterVarAddress(c *VarAddressContext)

	// EnterVarCriterion is called when entering the varCriterion production.
	EnterVarCriterion(c *VarCriterionContext)

	// EnterVarMeasurement is called when entering the varMeasurement production.
	EnterVarMeasurement(c *VarMeasurementContext)

	// EnterVarSelectionCharacteristic is called when entering the varSelectionCharacteristic production.
	EnterVarSelectionCharacteristic(c *VarSelectionCharacteristicContext)

	// EnterVarForbiddenComb is called when entering the varForbiddenComb production.
	EnterVarForbiddenComb(c *VarForbiddenCombContext)

	// EnterVarNaming is called when entering the varNaming production.
	EnterVarNaming(c *VarNamingContext)

	// EnterVarSeparator is called when entering the varSeparator production.
	EnterVarSeparator(c *VarSeparatorContext)

	// EnterIntegerValue is called when entering the integerValue production.
	EnterIntegerValue(c *IntegerValueContext)

	// EnterNumericValue is called when entering the numericValue production.
	EnterNumericValue(c *NumericValueContext)

	// EnterStringValue is called when entering the stringValue production.
	EnterStringValue(c *StringValueContext)

	// EnterIdentifierValue is called when entering the identifierValue production.
	EnterIdentifierValue(c *IdentifierValueContext)

	// EnterPartialIdentifier is called when entering the partialIdentifier production.
	EnterPartialIdentifier(c *PartialIdentifierContext)

	// EnterArraySpecifier is called when entering the arraySpecifier production.
	EnterArraySpecifier(c *ArraySpecifierContext)

	// EnterDataType is called when entering the dataType production.
	EnterDataType(c *DataTypeContext)

	// EnterDatasize is called when entering the datasize production.
	EnterDatasize(c *DatasizeContext)

	// EnterAddrtype is called when entering the addrtype production.
	EnterAddrtype(c *AddrtypeContext)

	// EnterByteOrderValue is called when entering the byteOrderValue production.
	EnterByteOrderValue(c *ByteOrderValueContext)

	// EnterIndexorder is called when entering the indexorder production.
	EnterIndexorder(c *IndexorderContext)

	// EnterLinkType is called when entering the linkType production.
	EnterLinkType(c *LinkTypeContext)

	// EnterA2ml is called when entering the a2ml production.
	EnterA2ml(c *A2mlContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterTypeDefinition is called when entering the typeDefinition production.
	EnterTypeDefinition(c *TypeDefinitionContext)

	// EnterA2mlTypeName is called when entering the a2mlTypeName production.
	EnterA2mlTypeName(c *A2mlTypeNameContext)

	// EnterPredefinedTypeName is called when entering the predefinedTypeName production.
	EnterPredefinedTypeName(c *PredefinedTypeNameContext)

	// EnterBlockDefinition is called when entering the blockDefinition production.
	EnterBlockDefinition(c *BlockDefinitionContext)

	// EnterEnumTypeName is called when entering the enumTypeName production.
	EnterEnumTypeName(c *EnumTypeNameContext)

	// EnterEnumeratorList is called when entering the enumeratorList production.
	EnterEnumeratorList(c *EnumeratorListContext)

	// EnterEnumerator is called when entering the enumerator production.
	EnterEnumerator(c *EnumeratorContext)

	// EnterStructTypeName is called when entering the structTypeName production.
	EnterStructTypeName(c *StructTypeNameContext)

	// EnterStructMember is called when entering the structMember production.
	EnterStructMember(c *StructMemberContext)

	// EnterMember is called when entering the member production.
	EnterMember(c *MemberContext)

	// EnterTaggedStructTypeName is called when entering the taggedStructTypeName production.
	EnterTaggedStructTypeName(c *TaggedStructTypeNameContext)

	// EnterTaggedStructMember is called when entering the taggedStructMember production.
	EnterTaggedStructMember(c *TaggedStructMemberContext)

	// EnterTaggedStructDefinition is called when entering the taggedStructDefinition production.
	EnterTaggedStructDefinition(c *TaggedStructDefinitionContext)

	// EnterTaggedUnionTypeName is called when entering the taggedUnionTypeName production.
	EnterTaggedUnionTypeName(c *TaggedUnionTypeNameContext)

	// EnterTaggedUnionMember is called when entering the taggedUnionMember production.
	EnterTaggedUnionMember(c *TaggedUnionMemberContext)

	// EnterTagValue is called when entering the tagValue production.
	EnterTagValue(c *TagValueContext)

	// EnterIfData is called when entering the ifData production.
	EnterIfData(c *IfDataContext)

	// EnterGenericParameter is called when entering the genericParameter production.
	EnterGenericParameter(c *GenericParameterContext)

	// EnterGenericNode is called when entering the genericNode production.
	EnterGenericNode(c *GenericNodeContext)

	// ExitA2lFile is called when exiting the a2lFile production.
	ExitA2lFile(c *A2lFileContext)

	// ExitAlignmentByte is called when exiting the alignmentByte production.
	ExitAlignmentByte(c *AlignmentByteContext)

	// ExitAlignmentFloat16Ieee is called when exiting the alignmentFloat16Ieee production.
	ExitAlignmentFloat16Ieee(c *AlignmentFloat16IeeeContext)

	// ExitAlignmentFloat32Ieee is called when exiting the alignmentFloat32Ieee production.
	ExitAlignmentFloat32Ieee(c *AlignmentFloat32IeeeContext)

	// ExitAlignmentFloat64Ieee is called when exiting the alignmentFloat64Ieee production.
	ExitAlignmentFloat64Ieee(c *AlignmentFloat64IeeeContext)

	// ExitAlignmentInt64 is called when exiting the alignmentInt64 production.
	ExitAlignmentInt64(c *AlignmentInt64Context)

	// ExitAlignmentLong is called when exiting the alignmentLong production.
	ExitAlignmentLong(c *AlignmentLongContext)

	// ExitAlignmentWord is called when exiting the alignmentWord production.
	ExitAlignmentWord(c *AlignmentWordContext)

	// ExitAnnotation is called when exiting the annotation production.
	ExitAnnotation(c *AnnotationContext)

	// ExitAnnotationLabel is called when exiting the annotationLabel production.
	ExitAnnotationLabel(c *AnnotationLabelContext)

	// ExitAnnotationOrigin is called when exiting the annotationOrigin production.
	ExitAnnotationOrigin(c *AnnotationOriginContext)

	// ExitAnnotationText is called when exiting the annotationText production.
	ExitAnnotationText(c *AnnotationTextContext)

	// ExitBitMask is called when exiting the bitMask production.
	ExitBitMask(c *BitMaskContext)

	// ExitByteOrder is called when exiting the byteOrder production.
	ExitByteOrder(c *ByteOrderContext)

	// ExitCalibrationAccess is called when exiting the calibrationAccess production.
	ExitCalibrationAccess(c *CalibrationAccessContext)

	// ExitDefaultValue is called when exiting the defaultValue production.
	ExitDefaultValue(c *DefaultValueContext)

	// ExitDeposit is called when exiting the deposit production.
	ExitDeposit(c *DepositContext)

	// ExitDiscrete is called when exiting the discrete production.
	ExitDiscrete(c *DiscreteContext)

	// ExitDisplayIdentifier is called when exiting the displayIdentifier production.
	ExitDisplayIdentifier(c *DisplayIdentifierContext)

	// ExitEcuAddressExtension is called when exiting the ecuAddressExtension production.
	ExitEcuAddressExtension(c *EcuAddressExtensionContext)

	// ExitExtendedLimits is called when exiting the extendedLimits production.
	ExitExtendedLimits(c *ExtendedLimitsContext)

	// ExitFormat is called when exiting the format production.
	ExitFormat(c *FormatContext)

	// ExitFunctionList is called when exiting the functionList production.
	ExitFunctionList(c *FunctionListContext)

	// ExitGuardRails is called when exiting the guardRails production.
	ExitGuardRails(c *GuardRailsContext)

	// ExitMatrixDim is called when exiting the matrixDim production.
	ExitMatrixDim(c *MatrixDimContext)

	// ExitMaxRefresh is called when exiting the maxRefresh production.
	ExitMaxRefresh(c *MaxRefreshContext)

	// ExitMonotony is called when exiting the monotony production.
	ExitMonotony(c *MonotonyContext)

	// ExitPhysUnit is called when exiting the physUnit production.
	ExitPhysUnit(c *PhysUnitContext)

	// ExitReadOnly is called when exiting the readOnly production.
	ExitReadOnly(c *ReadOnlyContext)

	// ExitRefCharacteristic is called when exiting the refCharacteristic production.
	ExitRefCharacteristic(c *RefCharacteristicContext)

	// ExitRefMemorySegment is called when exiting the refMemorySegment production.
	ExitRefMemorySegment(c *RefMemorySegmentContext)

	// ExitRefUnit is called when exiting the refUnit production.
	ExitRefUnit(c *RefUnitContext)

	// ExitStepSize is called when exiting the stepSize production.
	ExitStepSize(c *StepSizeContext)

	// ExitSymbolLink is called when exiting the symbolLink production.
	ExitSymbolLink(c *SymbolLinkContext)

	// ExitVersion is called when exiting the version production.
	ExitVersion(c *VersionContext)

	// ExitAsap2Version is called when exiting the asap2Version production.
	ExitAsap2Version(c *Asap2VersionContext)

	// ExitA2mlVersion is called when exiting the a2mlVersion production.
	ExitA2mlVersion(c *A2mlVersionContext)

	// ExitProject is called when exiting the project production.
	ExitProject(c *ProjectContext)

	// ExitHeader is called when exiting the header production.
	ExitHeader(c *HeaderContext)

	// ExitProjectNo is called when exiting the projectNo production.
	ExitProjectNo(c *ProjectNoContext)

	// ExitModule is called when exiting the module production.
	ExitModule(c *ModuleContext)

	// ExitAxisPts is called when exiting the axisPts production.
	ExitAxisPts(c *AxisPtsContext)

	// ExitCharacteristic is called when exiting the characteristic production.
	ExitCharacteristic(c *CharacteristicContext)

	// ExitAxisDescr is called when exiting the axisDescr production.
	ExitAxisDescr(c *AxisDescrContext)

	// ExitAxisPtsRef is called when exiting the axisPtsRef production.
	ExitAxisPtsRef(c *AxisPtsRefContext)

	// ExitCurveAxisRef is called when exiting the curveAxisRef production.
	ExitCurveAxisRef(c *CurveAxisRefContext)

	// ExitFixAxisPar is called when exiting the fixAxisPar production.
	ExitFixAxisPar(c *FixAxisParContext)

	// ExitFixAxisParDist is called when exiting the fixAxisParDist production.
	ExitFixAxisParDist(c *FixAxisParDistContext)

	// ExitFixAxisParList is called when exiting the fixAxisParList production.
	ExitFixAxisParList(c *FixAxisParListContext)

	// ExitMaxGrad is called when exiting the maxGrad production.
	ExitMaxGrad(c *MaxGradContext)

	// ExitComparisonQuantity is called when exiting the comparisonQuantity production.
	ExitComparisonQuantity(c *ComparisonQuantityContext)

	// ExitDependentCharacteristic is called when exiting the dependentCharacteristic production.
	ExitDependentCharacteristic(c *DependentCharacteristicContext)

	// ExitMapList is called when exiting the mapList production.
	ExitMapList(c *MapListContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitVirtualCharacteristic is called when exiting the virtualCharacteristic production.
	ExitVirtualCharacteristic(c *VirtualCharacteristicContext)

	// ExitCompuMethod is called when exiting the compuMethod production.
	ExitCompuMethod(c *CompuMethodContext)

	// ExitCoeffs is called when exiting the coeffs production.
	ExitCoeffs(c *CoeffsContext)

	// ExitCoeffsLinear is called when exiting the coeffsLinear production.
	ExitCoeffsLinear(c *CoeffsLinearContext)

	// ExitCompuTabRef is called when exiting the compuTabRef production.
	ExitCompuTabRef(c *CompuTabRefContext)

	// ExitFormula is called when exiting the formula production.
	ExitFormula(c *FormulaContext)

	// ExitFormulaInv is called when exiting the formulaInv production.
	ExitFormulaInv(c *FormulaInvContext)

	// ExitStatusStringRef is called when exiting the statusStringRef production.
	ExitStatusStringRef(c *StatusStringRefContext)

	// ExitCompuTab is called when exiting the compuTab production.
	ExitCompuTab(c *CompuTabContext)

	// ExitDefaultValueNumeric is called when exiting the defaultValueNumeric production.
	ExitDefaultValueNumeric(c *DefaultValueNumericContext)

	// ExitCompuVTab is called when exiting the compuVTab production.
	ExitCompuVTab(c *CompuVTabContext)

	// ExitCompuVTabRange is called when exiting the compuVTabRange production.
	ExitCompuVTabRange(c *CompuVTabRangeContext)

	// ExitFrame is called when exiting the frame production.
	ExitFrame(c *FrameContext)

	// ExitFrameMeasurement is called when exiting the frameMeasurement production.
	ExitFrameMeasurement(c *FrameMeasurementContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitDefCharacteristic is called when exiting the defCharacteristic production.
	ExitDefCharacteristic(c *DefCharacteristicContext)

	// ExitFunctionVersion is called when exiting the functionVersion production.
	ExitFunctionVersion(c *FunctionVersionContext)

	// ExitInMeasurement is called when exiting the inMeasurement production.
	ExitInMeasurement(c *InMeasurementContext)

	// ExitLocMeasurement is called when exiting the locMeasurement production.
	ExitLocMeasurement(c *LocMeasurementContext)

	// ExitOutMeasurement is called when exiting the outMeasurement production.
	ExitOutMeasurement(c *OutMeasurementContext)

	// ExitSubFunction is called when exiting the subFunction production.
	ExitSubFunction(c *SubFunctionContext)

	// ExitGroup is called when exiting the group production.
	ExitGroup(c *GroupContext)

	// ExitRefMeasurement is called when exiting the refMeasurement production.
	ExitRefMeasurement(c *RefMeasurementContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitSubGroup is called when exiting the subGroup production.
	ExitSubGroup(c *SubGroupContext)

	// ExitInstance is called when exiting the instance production.
	ExitInstance(c *InstanceContext)

	// ExitMeasurement is called when exiting the measurement production.
	ExitMeasurement(c *MeasurementContext)

	// ExitArraySize is called when exiting the arraySize production.
	ExitArraySize(c *ArraySizeContext)

	// ExitBitOperation is called when exiting the bitOperation production.
	ExitBitOperation(c *BitOperationContext)

	// ExitLeftShift is called when exiting the leftShift production.
	ExitLeftShift(c *LeftShiftContext)

	// ExitRightShift is called when exiting the rightShift production.
	ExitRightShift(c *RightShiftContext)

	// ExitSignExtend is called when exiting the signExtend production.
	ExitSignExtend(c *SignExtendContext)

	// ExitEcuAddress is called when exiting the ecuAddress production.
	ExitEcuAddress(c *EcuAddressContext)

	// ExitErrorMask is called when exiting the errorMask production.
	ExitErrorMask(c *ErrorMaskContext)

	// ExitLayout is called when exiting the layout production.
	ExitLayout(c *LayoutContext)

	// ExitReadWrite is called when exiting the readWrite production.
	ExitReadWrite(c *ReadWriteContext)

	// ExitVirtual is called when exiting the virtual production.
	ExitVirtual(c *VirtualContext)

	// ExitModCommon is called when exiting the modCommon production.
	ExitModCommon(c *ModCommonContext)

	// ExitDataSize is called when exiting the dataSize production.
	ExitDataSize(c *DataSizeContext)

	// ExitSRecLayout is called when exiting the sRecLayout production.
	ExitSRecLayout(c *SRecLayoutContext)

	// ExitModPar is called when exiting the modPar production.
	ExitModPar(c *ModParContext)

	// ExitAddrEpk is called when exiting the addrEpk production.
	ExitAddrEpk(c *AddrEpkContext)

	// ExitCalibrationMethod is called when exiting the calibrationMethod production.
	ExitCalibrationMethod(c *CalibrationMethodContext)

	// ExitCalibrationHandle is called when exiting the calibrationHandle production.
	ExitCalibrationHandle(c *CalibrationHandleContext)

	// ExitCpuType is called when exiting the cpuType production.
	ExitCpuType(c *CpuTypeContext)

	// ExitCustomer is called when exiting the customer production.
	ExitCustomer(c *CustomerContext)

	// ExitCustomerNo is called when exiting the customerNo production.
	ExitCustomerNo(c *CustomerNoContext)

	// ExitEcu is called when exiting the ecu production.
	ExitEcu(c *EcuContext)

	// ExitEcuCalibrationOffset is called when exiting the ecuCalibrationOffset production.
	ExitEcuCalibrationOffset(c *EcuCalibrationOffsetContext)

	// ExitEpk is called when exiting the epk production.
	ExitEpk(c *EpkContext)

	// ExitMemoryLayout is called when exiting the memoryLayout production.
	ExitMemoryLayout(c *MemoryLayoutContext)

	// ExitMemorySegment is called when exiting the memorySegment production.
	ExitMemorySegment(c *MemorySegmentContext)

	// ExitNoOfInterfaces is called when exiting the noOfInterfaces production.
	ExitNoOfInterfaces(c *NoOfInterfacesContext)

	// ExitPhoneNo is called when exiting the phoneNo production.
	ExitPhoneNo(c *PhoneNoContext)

	// ExitSupplier is called when exiting the supplier production.
	ExitSupplier(c *SupplierContext)

	// ExitSystemConstant is called when exiting the systemConstant production.
	ExitSystemConstant(c *SystemConstantContext)

	// ExitUser is called when exiting the user production.
	ExitUser(c *UserContext)

	// ExitRecordLayout is called when exiting the recordLayout production.
	ExitRecordLayout(c *RecordLayoutContext)

	// ExitAxisPtsX is called when exiting the axisPtsX production.
	ExitAxisPtsX(c *AxisPtsXContext)

	// ExitAxisPtsY is called when exiting the axisPtsY production.
	ExitAxisPtsY(c *AxisPtsYContext)

	// ExitAxisPtsZ is called when exiting the axisPtsZ production.
	ExitAxisPtsZ(c *AxisPtsZContext)

	// ExitAxisPts4 is called when exiting the axisPts4 production.
	ExitAxisPts4(c *AxisPts4Context)

	// ExitAxisPts5 is called when exiting the axisPts5 production.
	ExitAxisPts5(c *AxisPts5Context)

	// ExitAxisRescaleX is called when exiting the axisRescaleX production.
	ExitAxisRescaleX(c *AxisRescaleXContext)

	// ExitAxisRescaleY is called when exiting the axisRescaleY production.
	ExitAxisRescaleY(c *AxisRescaleYContext)

	// ExitAxisRescaleZ is called when exiting the axisRescaleZ production.
	ExitAxisRescaleZ(c *AxisRescaleZContext)

	// ExitAxisRescale4 is called when exiting the axisRescale4 production.
	ExitAxisRescale4(c *AxisRescale4Context)

	// ExitAxisRescale5 is called when exiting the axisRescale5 production.
	ExitAxisRescale5(c *AxisRescale5Context)

	// ExitDistOpX is called when exiting the distOpX production.
	ExitDistOpX(c *DistOpXContext)

	// ExitDistOpY is called when exiting the distOpY production.
	ExitDistOpY(c *DistOpYContext)

	// ExitDistOpZ is called when exiting the distOpZ production.
	ExitDistOpZ(c *DistOpZContext)

	// ExitDistOp4 is called when exiting the distOp4 production.
	ExitDistOp4(c *DistOp4Context)

	// ExitDistOp5 is called when exiting the distOp5 production.
	ExitDistOp5(c *DistOp5Context)

	// ExitFixNoAxisPtsX is called when exiting the fixNoAxisPtsX production.
	ExitFixNoAxisPtsX(c *FixNoAxisPtsXContext)

	// ExitFixNoAxisPtsY is called when exiting the fixNoAxisPtsY production.
	ExitFixNoAxisPtsY(c *FixNoAxisPtsYContext)

	// ExitFixNoAxisPtsZ is called when exiting the fixNoAxisPtsZ production.
	ExitFixNoAxisPtsZ(c *FixNoAxisPtsZContext)

	// ExitFixNoAxisPts4 is called when exiting the fixNoAxisPts4 production.
	ExitFixNoAxisPts4(c *FixNoAxisPts4Context)

	// ExitFixNoAxisPts5 is called when exiting the fixNoAxisPts5 production.
	ExitFixNoAxisPts5(c *FixNoAxisPts5Context)

	// ExitFncValues is called when exiting the fncValues production.
	ExitFncValues(c *FncValuesContext)

	// ExitIdentification is called when exiting the identification production.
	ExitIdentification(c *IdentificationContext)

	// ExitNoAxisPtsX is called when exiting the noAxisPtsX production.
	ExitNoAxisPtsX(c *NoAxisPtsXContext)

	// ExitNoAxisPtsY is called when exiting the noAxisPtsY production.
	ExitNoAxisPtsY(c *NoAxisPtsYContext)

	// ExitNoAxisPtsZ is called when exiting the noAxisPtsZ production.
	ExitNoAxisPtsZ(c *NoAxisPtsZContext)

	// ExitNoAxisPts4 is called when exiting the noAxisPts4 production.
	ExitNoAxisPts4(c *NoAxisPts4Context)

	// ExitNoAxisPts5 is called when exiting the noAxisPts5 production.
	ExitNoAxisPts5(c *NoAxisPts5Context)

	// ExitStaticRecordLayout is called when exiting the staticRecordLayout production.
	ExitStaticRecordLayout(c *StaticRecordLayoutContext)

	// ExitNoRescaleX is called when exiting the noRescaleX production.
	ExitNoRescaleX(c *NoRescaleXContext)

	// ExitNoRescaleY is called when exiting the noRescaleY production.
	ExitNoRescaleY(c *NoRescaleYContext)

	// ExitNoRescaleZ is called when exiting the noRescaleZ production.
	ExitNoRescaleZ(c *NoRescaleZContext)

	// ExitNoRescale4 is called when exiting the noRescale4 production.
	ExitNoRescale4(c *NoRescale4Context)

	// ExitNoRescale5 is called when exiting the noRescale5 production.
	ExitNoRescale5(c *NoRescale5Context)

	// ExitOffsetX is called when exiting the offsetX production.
	ExitOffsetX(c *OffsetXContext)

	// ExitOffsetY is called when exiting the offsetY production.
	ExitOffsetY(c *OffsetYContext)

	// ExitOffsetZ is called when exiting the offsetZ production.
	ExitOffsetZ(c *OffsetZContext)

	// ExitOffset4 is called when exiting the offset4 production.
	ExitOffset4(c *Offset4Context)

	// ExitOffset5 is called when exiting the offset5 production.
	ExitOffset5(c *Offset5Context)

	// ExitReserved is called when exiting the reserved production.
	ExitReserved(c *ReservedContext)

	// ExitRipAddrW is called when exiting the ripAddrW production.
	ExitRipAddrW(c *RipAddrWContext)

	// ExitRipAddrX is called when exiting the ripAddrX production.
	ExitRipAddrX(c *RipAddrXContext)

	// ExitRipAddrY is called when exiting the ripAddrY production.
	ExitRipAddrY(c *RipAddrYContext)

	// ExitRipAddrZ is called when exiting the ripAddrZ production.
	ExitRipAddrZ(c *RipAddrZContext)

	// ExitRipAddr4 is called when exiting the ripAddr4 production.
	ExitRipAddr4(c *RipAddr4Context)

	// ExitRipAddr5 is called when exiting the ripAddr5 production.
	ExitRipAddr5(c *RipAddr5Context)

	// ExitShiftOpX is called when exiting the shiftOpX production.
	ExitShiftOpX(c *ShiftOpXContext)

	// ExitShiftOpY is called when exiting the shiftOpY production.
	ExitShiftOpY(c *ShiftOpYContext)

	// ExitShiftOpZ is called when exiting the shiftOpZ production.
	ExitShiftOpZ(c *ShiftOpZContext)

	// ExitShiftOp4 is called when exiting the shiftOp4 production.
	ExitShiftOp4(c *ShiftOp4Context)

	// ExitShiftOp5 is called when exiting the shiftOp5 production.
	ExitShiftOp5(c *ShiftOp5Context)

	// ExitSrcAddrX is called when exiting the srcAddrX production.
	ExitSrcAddrX(c *SrcAddrXContext)

	// ExitSrcAddrY is called when exiting the srcAddrY production.
	ExitSrcAddrY(c *SrcAddrYContext)

	// ExitSrcAddrZ is called when exiting the srcAddrZ production.
	ExitSrcAddrZ(c *SrcAddrZContext)

	// ExitSrcAddr4 is called when exiting the srcAddr4 production.
	ExitSrcAddr4(c *SrcAddr4Context)

	// ExitSrcAddr5 is called when exiting the srcAddr5 production.
	ExitSrcAddr5(c *SrcAddr5Context)

	// ExitTypedefCharacteristic is called when exiting the typedefCharacteristic production.
	ExitTypedefCharacteristic(c *TypedefCharacteristicContext)

	// ExitTypedefMeasurement is called when exiting the typedefMeasurement production.
	ExitTypedefMeasurement(c *TypedefMeasurementContext)

	// ExitTypedefStructure is called when exiting the typedefStructure production.
	ExitTypedefStructure(c *TypedefStructureContext)

	// ExitStructureComponent is called when exiting the structureComponent production.
	ExitStructureComponent(c *StructureComponentContext)

	// ExitUnit is called when exiting the unit production.
	ExitUnit(c *UnitContext)

	// ExitSiExponents is called when exiting the siExponents production.
	ExitSiExponents(c *SiExponentsContext)

	// ExitUnitConversion is called when exiting the unitConversion production.
	ExitUnitConversion(c *UnitConversionContext)

	// ExitUserRights is called when exiting the userRights production.
	ExitUserRights(c *UserRightsContext)

	// ExitRefGroup is called when exiting the refGroup production.
	ExitRefGroup(c *RefGroupContext)

	// ExitVariantCoding is called when exiting the variantCoding production.
	ExitVariantCoding(c *VariantCodingContext)

	// ExitVarCharacteristic is called when exiting the varCharacteristic production.
	ExitVarCharacteristic(c *VarCharacteristicContext)

	// ExitVarAddress is called when exiting the varAddress production.
	ExitVarAddress(c *VarAddressContext)

	// ExitVarCriterion is called when exiting the varCriterion production.
	ExitVarCriterion(c *VarCriterionContext)

	// ExitVarMeasurement is called when exiting the varMeasurement production.
	ExitVarMeasurement(c *VarMeasurementContext)

	// ExitVarSelectionCharacteristic is called when exiting the varSelectionCharacteristic production.
	ExitVarSelectionCharacteristic(c *VarSelectionCharacteristicContext)

	// ExitVarForbiddenComb is called when exiting the varForbiddenComb production.
	ExitVarForbiddenComb(c *VarForbiddenCombContext)

	// ExitVarNaming is called when exiting the varNaming production.
	ExitVarNaming(c *VarNamingContext)

	// ExitVarSeparator is called when exiting the varSeparator production.
	ExitVarSeparator(c *VarSeparatorContext)

	// ExitIntegerValue is called when exiting the integerValue production.
	ExitIntegerValue(c *IntegerValueContext)

	// ExitNumericValue is called when exiting the numericValue production.
	ExitNumericValue(c *NumericValueContext)

	// ExitStringValue is called when exiting the stringValue production.
	ExitStringValue(c *StringValueContext)

	// ExitIdentifierValue is called when exiting the identifierValue production.
	ExitIdentifierValue(c *IdentifierValueContext)

	// ExitPartialIdentifier is called when exiting the partialIdentifier production.
	ExitPartialIdentifier(c *PartialIdentifierContext)

	// ExitArraySpecifier is called when exiting the arraySpecifier production.
	ExitArraySpecifier(c *ArraySpecifierContext)

	// ExitDataType is called when exiting the dataType production.
	ExitDataType(c *DataTypeContext)

	// ExitDatasize is called when exiting the datasize production.
	ExitDatasize(c *DatasizeContext)

	// ExitAddrtype is called when exiting the addrtype production.
	ExitAddrtype(c *AddrtypeContext)

	// ExitByteOrderValue is called when exiting the byteOrderValue production.
	ExitByteOrderValue(c *ByteOrderValueContext)

	// ExitIndexorder is called when exiting the indexorder production.
	ExitIndexorder(c *IndexorderContext)

	// ExitLinkType is called when exiting the linkType production.
	ExitLinkType(c *LinkTypeContext)

	// ExitA2ml is called when exiting the a2ml production.
	ExitA2ml(c *A2mlContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitTypeDefinition is called when exiting the typeDefinition production.
	ExitTypeDefinition(c *TypeDefinitionContext)

	// ExitA2mlTypeName is called when exiting the a2mlTypeName production.
	ExitA2mlTypeName(c *A2mlTypeNameContext)

	// ExitPredefinedTypeName is called when exiting the predefinedTypeName production.
	ExitPredefinedTypeName(c *PredefinedTypeNameContext)

	// ExitBlockDefinition is called when exiting the blockDefinition production.
	ExitBlockDefinition(c *BlockDefinitionContext)

	// ExitEnumTypeName is called when exiting the enumTypeName production.
	ExitEnumTypeName(c *EnumTypeNameContext)

	// ExitEnumeratorList is called when exiting the enumeratorList production.
	ExitEnumeratorList(c *EnumeratorListContext)

	// ExitEnumerator is called when exiting the enumerator production.
	ExitEnumerator(c *EnumeratorContext)

	// ExitStructTypeName is called when exiting the structTypeName production.
	ExitStructTypeName(c *StructTypeNameContext)

	// ExitStructMember is called when exiting the structMember production.
	ExitStructMember(c *StructMemberContext)

	// ExitMember is called when exiting the member production.
	ExitMember(c *MemberContext)

	// ExitTaggedStructTypeName is called when exiting the taggedStructTypeName production.
	ExitTaggedStructTypeName(c *TaggedStructTypeNameContext)

	// ExitTaggedStructMember is called when exiting the taggedStructMember production.
	ExitTaggedStructMember(c *TaggedStructMemberContext)

	// ExitTaggedStructDefinition is called when exiting the taggedStructDefinition production.
	ExitTaggedStructDefinition(c *TaggedStructDefinitionContext)

	// ExitTaggedUnionTypeName is called when exiting the taggedUnionTypeName production.
	ExitTaggedUnionTypeName(c *TaggedUnionTypeNameContext)

	// ExitTaggedUnionMember is called when exiting the taggedUnionMember production.
	ExitTaggedUnionMember(c *TaggedUnionMemberContext)

	// ExitTagValue is called when exiting the tagValue production.
	ExitTagValue(c *TagValueContext)

	// ExitIfData is called when exiting the ifData production.
	ExitIfData(c *IfDataContext)

	// ExitGenericParameter is called when exiting the genericParameter production.
	ExitGenericParameter(c *GenericParameterContext)

	// ExitGenericNode is called when exiting the genericNode production.
	ExitGenericNode(c *GenericNodeContext)
}
