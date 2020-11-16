// Code generated from grammar/A2L.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // A2L

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseA2LListener is a complete listener for a parse tree produced by A2LParser.
type BaseA2LListener struct{}

var _ A2LListener = &BaseA2LListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseA2LListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseA2LListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseA2LListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseA2LListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterA2lFile is called when production a2lFile is entered.
func (s *BaseA2LListener) EnterA2lFile(ctx *A2lFileContext) {}

// ExitA2lFile is called when production a2lFile is exited.
func (s *BaseA2LListener) ExitA2lFile(ctx *A2lFileContext) {}

// EnterAlignmentByte is called when production alignmentByte is entered.
func (s *BaseA2LListener) EnterAlignmentByte(ctx *AlignmentByteContext) {}

// ExitAlignmentByte is called when production alignmentByte is exited.
func (s *BaseA2LListener) ExitAlignmentByte(ctx *AlignmentByteContext) {}

// EnterAlignmentFloat16Ieee is called when production alignmentFloat16Ieee is entered.
func (s *BaseA2LListener) EnterAlignmentFloat16Ieee(ctx *AlignmentFloat16IeeeContext) {}

// ExitAlignmentFloat16Ieee is called when production alignmentFloat16Ieee is exited.
func (s *BaseA2LListener) ExitAlignmentFloat16Ieee(ctx *AlignmentFloat16IeeeContext) {}

// EnterAlignmentFloat32Ieee is called when production alignmentFloat32Ieee is entered.
func (s *BaseA2LListener) EnterAlignmentFloat32Ieee(ctx *AlignmentFloat32IeeeContext) {}

// ExitAlignmentFloat32Ieee is called when production alignmentFloat32Ieee is exited.
func (s *BaseA2LListener) ExitAlignmentFloat32Ieee(ctx *AlignmentFloat32IeeeContext) {}

// EnterAlignmentFloat64Ieee is called when production alignmentFloat64Ieee is entered.
func (s *BaseA2LListener) EnterAlignmentFloat64Ieee(ctx *AlignmentFloat64IeeeContext) {}

// ExitAlignmentFloat64Ieee is called when production alignmentFloat64Ieee is exited.
func (s *BaseA2LListener) ExitAlignmentFloat64Ieee(ctx *AlignmentFloat64IeeeContext) {}

// EnterAlignmentInt64 is called when production alignmentInt64 is entered.
func (s *BaseA2LListener) EnterAlignmentInt64(ctx *AlignmentInt64Context) {}

// ExitAlignmentInt64 is called when production alignmentInt64 is exited.
func (s *BaseA2LListener) ExitAlignmentInt64(ctx *AlignmentInt64Context) {}

// EnterAlignmentLong is called when production alignmentLong is entered.
func (s *BaseA2LListener) EnterAlignmentLong(ctx *AlignmentLongContext) {}

// ExitAlignmentLong is called when production alignmentLong is exited.
func (s *BaseA2LListener) ExitAlignmentLong(ctx *AlignmentLongContext) {}

// EnterAlignmentWord is called when production alignmentWord is entered.
func (s *BaseA2LListener) EnterAlignmentWord(ctx *AlignmentWordContext) {}

// ExitAlignmentWord is called when production alignmentWord is exited.
func (s *BaseA2LListener) ExitAlignmentWord(ctx *AlignmentWordContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *BaseA2LListener) EnterAnnotation(ctx *AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *BaseA2LListener) ExitAnnotation(ctx *AnnotationContext) {}

// EnterAnnotationLabel is called when production annotationLabel is entered.
func (s *BaseA2LListener) EnterAnnotationLabel(ctx *AnnotationLabelContext) {}

// ExitAnnotationLabel is called when production annotationLabel is exited.
func (s *BaseA2LListener) ExitAnnotationLabel(ctx *AnnotationLabelContext) {}

// EnterAnnotationOrigin is called when production annotationOrigin is entered.
func (s *BaseA2LListener) EnterAnnotationOrigin(ctx *AnnotationOriginContext) {}

// ExitAnnotationOrigin is called when production annotationOrigin is exited.
func (s *BaseA2LListener) ExitAnnotationOrigin(ctx *AnnotationOriginContext) {}

// EnterAnnotationText is called when production annotationText is entered.
func (s *BaseA2LListener) EnterAnnotationText(ctx *AnnotationTextContext) {}

// ExitAnnotationText is called when production annotationText is exited.
func (s *BaseA2LListener) ExitAnnotationText(ctx *AnnotationTextContext) {}

// EnterBitMask is called when production bitMask is entered.
func (s *BaseA2LListener) EnterBitMask(ctx *BitMaskContext) {}

// ExitBitMask is called when production bitMask is exited.
func (s *BaseA2LListener) ExitBitMask(ctx *BitMaskContext) {}

// EnterByteOrder is called when production byteOrder is entered.
func (s *BaseA2LListener) EnterByteOrder(ctx *ByteOrderContext) {}

// ExitByteOrder is called when production byteOrder is exited.
func (s *BaseA2LListener) ExitByteOrder(ctx *ByteOrderContext) {}

// EnterCalibrationAccess is called when production calibrationAccess is entered.
func (s *BaseA2LListener) EnterCalibrationAccess(ctx *CalibrationAccessContext) {}

// ExitCalibrationAccess is called when production calibrationAccess is exited.
func (s *BaseA2LListener) ExitCalibrationAccess(ctx *CalibrationAccessContext) {}

// EnterDefaultValue is called when production defaultValue is entered.
func (s *BaseA2LListener) EnterDefaultValue(ctx *DefaultValueContext) {}

// ExitDefaultValue is called when production defaultValue is exited.
func (s *BaseA2LListener) ExitDefaultValue(ctx *DefaultValueContext) {}

// EnterDeposit is called when production deposit is entered.
func (s *BaseA2LListener) EnterDeposit(ctx *DepositContext) {}

// ExitDeposit is called when production deposit is exited.
func (s *BaseA2LListener) ExitDeposit(ctx *DepositContext) {}

// EnterDiscrete is called when production discrete is entered.
func (s *BaseA2LListener) EnterDiscrete(ctx *DiscreteContext) {}

// ExitDiscrete is called when production discrete is exited.
func (s *BaseA2LListener) ExitDiscrete(ctx *DiscreteContext) {}

// EnterDisplayIdentifier is called when production displayIdentifier is entered.
func (s *BaseA2LListener) EnterDisplayIdentifier(ctx *DisplayIdentifierContext) {}

// ExitDisplayIdentifier is called when production displayIdentifier is exited.
func (s *BaseA2LListener) ExitDisplayIdentifier(ctx *DisplayIdentifierContext) {}

// EnterEcuAddressExtension is called when production ecuAddressExtension is entered.
func (s *BaseA2LListener) EnterEcuAddressExtension(ctx *EcuAddressExtensionContext) {}

// ExitEcuAddressExtension is called when production ecuAddressExtension is exited.
func (s *BaseA2LListener) ExitEcuAddressExtension(ctx *EcuAddressExtensionContext) {}

// EnterExtendedLimits is called when production extendedLimits is entered.
func (s *BaseA2LListener) EnterExtendedLimits(ctx *ExtendedLimitsContext) {}

// ExitExtendedLimits is called when production extendedLimits is exited.
func (s *BaseA2LListener) ExitExtendedLimits(ctx *ExtendedLimitsContext) {}

// EnterFormat is called when production format is entered.
func (s *BaseA2LListener) EnterFormat(ctx *FormatContext) {}

// ExitFormat is called when production format is exited.
func (s *BaseA2LListener) ExitFormat(ctx *FormatContext) {}

// EnterFunctionList is called when production functionList is entered.
func (s *BaseA2LListener) EnterFunctionList(ctx *FunctionListContext) {}

// ExitFunctionList is called when production functionList is exited.
func (s *BaseA2LListener) ExitFunctionList(ctx *FunctionListContext) {}

// EnterGuardRails is called when production guardRails is entered.
func (s *BaseA2LListener) EnterGuardRails(ctx *GuardRailsContext) {}

// ExitGuardRails is called when production guardRails is exited.
func (s *BaseA2LListener) ExitGuardRails(ctx *GuardRailsContext) {}

// EnterMatrixDim is called when production matrixDim is entered.
func (s *BaseA2LListener) EnterMatrixDim(ctx *MatrixDimContext) {}

// ExitMatrixDim is called when production matrixDim is exited.
func (s *BaseA2LListener) ExitMatrixDim(ctx *MatrixDimContext) {}

// EnterMaxRefresh is called when production maxRefresh is entered.
func (s *BaseA2LListener) EnterMaxRefresh(ctx *MaxRefreshContext) {}

// ExitMaxRefresh is called when production maxRefresh is exited.
func (s *BaseA2LListener) ExitMaxRefresh(ctx *MaxRefreshContext) {}

// EnterMonotony is called when production monotony is entered.
func (s *BaseA2LListener) EnterMonotony(ctx *MonotonyContext) {}

// ExitMonotony is called when production monotony is exited.
func (s *BaseA2LListener) ExitMonotony(ctx *MonotonyContext) {}

// EnterPhysUnit is called when production physUnit is entered.
func (s *BaseA2LListener) EnterPhysUnit(ctx *PhysUnitContext) {}

// ExitPhysUnit is called when production physUnit is exited.
func (s *BaseA2LListener) ExitPhysUnit(ctx *PhysUnitContext) {}

// EnterReadOnly is called when production readOnly is entered.
func (s *BaseA2LListener) EnterReadOnly(ctx *ReadOnlyContext) {}

// ExitReadOnly is called when production readOnly is exited.
func (s *BaseA2LListener) ExitReadOnly(ctx *ReadOnlyContext) {}

// EnterRefCharacteristic is called when production refCharacteristic is entered.
func (s *BaseA2LListener) EnterRefCharacteristic(ctx *RefCharacteristicContext) {}

// ExitRefCharacteristic is called when production refCharacteristic is exited.
func (s *BaseA2LListener) ExitRefCharacteristic(ctx *RefCharacteristicContext) {}

// EnterRefMemorySegment is called when production refMemorySegment is entered.
func (s *BaseA2LListener) EnterRefMemorySegment(ctx *RefMemorySegmentContext) {}

// ExitRefMemorySegment is called when production refMemorySegment is exited.
func (s *BaseA2LListener) ExitRefMemorySegment(ctx *RefMemorySegmentContext) {}

// EnterRefUnit is called when production refUnit is entered.
func (s *BaseA2LListener) EnterRefUnit(ctx *RefUnitContext) {}

// ExitRefUnit is called when production refUnit is exited.
func (s *BaseA2LListener) ExitRefUnit(ctx *RefUnitContext) {}

// EnterStepSize is called when production stepSize is entered.
func (s *BaseA2LListener) EnterStepSize(ctx *StepSizeContext) {}

// ExitStepSize is called when production stepSize is exited.
func (s *BaseA2LListener) ExitStepSize(ctx *StepSizeContext) {}

// EnterSymbolLink is called when production symbolLink is entered.
func (s *BaseA2LListener) EnterSymbolLink(ctx *SymbolLinkContext) {}

// ExitSymbolLink is called when production symbolLink is exited.
func (s *BaseA2LListener) ExitSymbolLink(ctx *SymbolLinkContext) {}

// EnterVersion is called when production version is entered.
func (s *BaseA2LListener) EnterVersion(ctx *VersionContext) {}

// ExitVersion is called when production version is exited.
func (s *BaseA2LListener) ExitVersion(ctx *VersionContext) {}

// EnterAsap2Version is called when production asap2Version is entered.
func (s *BaseA2LListener) EnterAsap2Version(ctx *Asap2VersionContext) {}

// ExitAsap2Version is called when production asap2Version is exited.
func (s *BaseA2LListener) ExitAsap2Version(ctx *Asap2VersionContext) {}

// EnterA2mlVersion is called when production a2mlVersion is entered.
func (s *BaseA2LListener) EnterA2mlVersion(ctx *A2mlVersionContext) {}

// ExitA2mlVersion is called when production a2mlVersion is exited.
func (s *BaseA2LListener) ExitA2mlVersion(ctx *A2mlVersionContext) {}

// EnterProject is called when production project is entered.
func (s *BaseA2LListener) EnterProject(ctx *ProjectContext) {}

// ExitProject is called when production project is exited.
func (s *BaseA2LListener) ExitProject(ctx *ProjectContext) {}

// EnterHeader is called when production header is entered.
func (s *BaseA2LListener) EnterHeader(ctx *HeaderContext) {}

// ExitHeader is called when production header is exited.
func (s *BaseA2LListener) ExitHeader(ctx *HeaderContext) {}

// EnterProjectNo is called when production projectNo is entered.
func (s *BaseA2LListener) EnterProjectNo(ctx *ProjectNoContext) {}

// ExitProjectNo is called when production projectNo is exited.
func (s *BaseA2LListener) ExitProjectNo(ctx *ProjectNoContext) {}

// EnterModule is called when production module is entered.
func (s *BaseA2LListener) EnterModule(ctx *ModuleContext) {}

// ExitModule is called when production module is exited.
func (s *BaseA2LListener) ExitModule(ctx *ModuleContext) {}

// EnterAxisPts is called when production axisPts is entered.
func (s *BaseA2LListener) EnterAxisPts(ctx *AxisPtsContext) {}

// ExitAxisPts is called when production axisPts is exited.
func (s *BaseA2LListener) ExitAxisPts(ctx *AxisPtsContext) {}

// EnterCharacteristic is called when production characteristic is entered.
func (s *BaseA2LListener) EnterCharacteristic(ctx *CharacteristicContext) {}

// ExitCharacteristic is called when production characteristic is exited.
func (s *BaseA2LListener) ExitCharacteristic(ctx *CharacteristicContext) {}

// EnterAxisDescr is called when production axisDescr is entered.
func (s *BaseA2LListener) EnterAxisDescr(ctx *AxisDescrContext) {}

// ExitAxisDescr is called when production axisDescr is exited.
func (s *BaseA2LListener) ExitAxisDescr(ctx *AxisDescrContext) {}

// EnterAxisPtsRef is called when production axisPtsRef is entered.
func (s *BaseA2LListener) EnterAxisPtsRef(ctx *AxisPtsRefContext) {}

// ExitAxisPtsRef is called when production axisPtsRef is exited.
func (s *BaseA2LListener) ExitAxisPtsRef(ctx *AxisPtsRefContext) {}

// EnterCurveAxisRef is called when production curveAxisRef is entered.
func (s *BaseA2LListener) EnterCurveAxisRef(ctx *CurveAxisRefContext) {}

// ExitCurveAxisRef is called when production curveAxisRef is exited.
func (s *BaseA2LListener) ExitCurveAxisRef(ctx *CurveAxisRefContext) {}

// EnterFixAxisPar is called when production fixAxisPar is entered.
func (s *BaseA2LListener) EnterFixAxisPar(ctx *FixAxisParContext) {}

// ExitFixAxisPar is called when production fixAxisPar is exited.
func (s *BaseA2LListener) ExitFixAxisPar(ctx *FixAxisParContext) {}

// EnterFixAxisParDist is called when production fixAxisParDist is entered.
func (s *BaseA2LListener) EnterFixAxisParDist(ctx *FixAxisParDistContext) {}

// ExitFixAxisParDist is called when production fixAxisParDist is exited.
func (s *BaseA2LListener) ExitFixAxisParDist(ctx *FixAxisParDistContext) {}

// EnterFixAxisParList is called when production fixAxisParList is entered.
func (s *BaseA2LListener) EnterFixAxisParList(ctx *FixAxisParListContext) {}

// ExitFixAxisParList is called when production fixAxisParList is exited.
func (s *BaseA2LListener) ExitFixAxisParList(ctx *FixAxisParListContext) {}

// EnterMaxGrad is called when production maxGrad is entered.
func (s *BaseA2LListener) EnterMaxGrad(ctx *MaxGradContext) {}

// ExitMaxGrad is called when production maxGrad is exited.
func (s *BaseA2LListener) ExitMaxGrad(ctx *MaxGradContext) {}

// EnterComparisonQuantity is called when production comparisonQuantity is entered.
func (s *BaseA2LListener) EnterComparisonQuantity(ctx *ComparisonQuantityContext) {}

// ExitComparisonQuantity is called when production comparisonQuantity is exited.
func (s *BaseA2LListener) ExitComparisonQuantity(ctx *ComparisonQuantityContext) {}

// EnterDependentCharacteristic is called when production dependentCharacteristic is entered.
func (s *BaseA2LListener) EnterDependentCharacteristic(ctx *DependentCharacteristicContext) {}

// ExitDependentCharacteristic is called when production dependentCharacteristic is exited.
func (s *BaseA2LListener) ExitDependentCharacteristic(ctx *DependentCharacteristicContext) {}

// EnterMapList is called when production mapList is entered.
func (s *BaseA2LListener) EnterMapList(ctx *MapListContext) {}

// ExitMapList is called when production mapList is exited.
func (s *BaseA2LListener) ExitMapList(ctx *MapListContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseA2LListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseA2LListener) ExitNumber(ctx *NumberContext) {}

// EnterVirtualCharacteristic is called when production virtualCharacteristic is entered.
func (s *BaseA2LListener) EnterVirtualCharacteristic(ctx *VirtualCharacteristicContext) {}

// ExitVirtualCharacteristic is called when production virtualCharacteristic is exited.
func (s *BaseA2LListener) ExitVirtualCharacteristic(ctx *VirtualCharacteristicContext) {}

// EnterCompuMethod is called when production compuMethod is entered.
func (s *BaseA2LListener) EnterCompuMethod(ctx *CompuMethodContext) {}

// ExitCompuMethod is called when production compuMethod is exited.
func (s *BaseA2LListener) ExitCompuMethod(ctx *CompuMethodContext) {}

// EnterCoeffs is called when production coeffs is entered.
func (s *BaseA2LListener) EnterCoeffs(ctx *CoeffsContext) {}

// ExitCoeffs is called when production coeffs is exited.
func (s *BaseA2LListener) ExitCoeffs(ctx *CoeffsContext) {}

// EnterCoeffsLinear is called when production coeffsLinear is entered.
func (s *BaseA2LListener) EnterCoeffsLinear(ctx *CoeffsLinearContext) {}

// ExitCoeffsLinear is called when production coeffsLinear is exited.
func (s *BaseA2LListener) ExitCoeffsLinear(ctx *CoeffsLinearContext) {}

// EnterCompuTabRef is called when production compuTabRef is entered.
func (s *BaseA2LListener) EnterCompuTabRef(ctx *CompuTabRefContext) {}

// ExitCompuTabRef is called when production compuTabRef is exited.
func (s *BaseA2LListener) ExitCompuTabRef(ctx *CompuTabRefContext) {}

// EnterFormula is called when production formula is entered.
func (s *BaseA2LListener) EnterFormula(ctx *FormulaContext) {}

// ExitFormula is called when production formula is exited.
func (s *BaseA2LListener) ExitFormula(ctx *FormulaContext) {}

// EnterFormulaInv is called when production formulaInv is entered.
func (s *BaseA2LListener) EnterFormulaInv(ctx *FormulaInvContext) {}

// ExitFormulaInv is called when production formulaInv is exited.
func (s *BaseA2LListener) ExitFormulaInv(ctx *FormulaInvContext) {}

// EnterStatusStringRef is called when production statusStringRef is entered.
func (s *BaseA2LListener) EnterStatusStringRef(ctx *StatusStringRefContext) {}

// ExitStatusStringRef is called when production statusStringRef is exited.
func (s *BaseA2LListener) ExitStatusStringRef(ctx *StatusStringRefContext) {}

// EnterCompuTab is called when production compuTab is entered.
func (s *BaseA2LListener) EnterCompuTab(ctx *CompuTabContext) {}

// ExitCompuTab is called when production compuTab is exited.
func (s *BaseA2LListener) ExitCompuTab(ctx *CompuTabContext) {}

// EnterDefaultValueNumeric is called when production defaultValueNumeric is entered.
func (s *BaseA2LListener) EnterDefaultValueNumeric(ctx *DefaultValueNumericContext) {}

// ExitDefaultValueNumeric is called when production defaultValueNumeric is exited.
func (s *BaseA2LListener) ExitDefaultValueNumeric(ctx *DefaultValueNumericContext) {}

// EnterCompuVTab is called when production compuVTab is entered.
func (s *BaseA2LListener) EnterCompuVTab(ctx *CompuVTabContext) {}

// ExitCompuVTab is called when production compuVTab is exited.
func (s *BaseA2LListener) ExitCompuVTab(ctx *CompuVTabContext) {}

// EnterCompuVTabRange is called when production compuVTabRange is entered.
func (s *BaseA2LListener) EnterCompuVTabRange(ctx *CompuVTabRangeContext) {}

// ExitCompuVTabRange is called when production compuVTabRange is exited.
func (s *BaseA2LListener) ExitCompuVTabRange(ctx *CompuVTabRangeContext) {}

// EnterFrame is called when production frame is entered.
func (s *BaseA2LListener) EnterFrame(ctx *FrameContext) {}

// ExitFrame is called when production frame is exited.
func (s *BaseA2LListener) ExitFrame(ctx *FrameContext) {}

// EnterFrameMeasurement is called when production frameMeasurement is entered.
func (s *BaseA2LListener) EnterFrameMeasurement(ctx *FrameMeasurementContext) {}

// ExitFrameMeasurement is called when production frameMeasurement is exited.
func (s *BaseA2LListener) ExitFrameMeasurement(ctx *FrameMeasurementContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseA2LListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseA2LListener) ExitFunction(ctx *FunctionContext) {}

// EnterDefCharacteristic is called when production defCharacteristic is entered.
func (s *BaseA2LListener) EnterDefCharacteristic(ctx *DefCharacteristicContext) {}

// ExitDefCharacteristic is called when production defCharacteristic is exited.
func (s *BaseA2LListener) ExitDefCharacteristic(ctx *DefCharacteristicContext) {}

// EnterFunctionVersion is called when production functionVersion is entered.
func (s *BaseA2LListener) EnterFunctionVersion(ctx *FunctionVersionContext) {}

// ExitFunctionVersion is called when production functionVersion is exited.
func (s *BaseA2LListener) ExitFunctionVersion(ctx *FunctionVersionContext) {}

// EnterInMeasurement is called when production inMeasurement is entered.
func (s *BaseA2LListener) EnterInMeasurement(ctx *InMeasurementContext) {}

// ExitInMeasurement is called when production inMeasurement is exited.
func (s *BaseA2LListener) ExitInMeasurement(ctx *InMeasurementContext) {}

// EnterLocMeasurement is called when production locMeasurement is entered.
func (s *BaseA2LListener) EnterLocMeasurement(ctx *LocMeasurementContext) {}

// ExitLocMeasurement is called when production locMeasurement is exited.
func (s *BaseA2LListener) ExitLocMeasurement(ctx *LocMeasurementContext) {}

// EnterOutMeasurement is called when production outMeasurement is entered.
func (s *BaseA2LListener) EnterOutMeasurement(ctx *OutMeasurementContext) {}

// ExitOutMeasurement is called when production outMeasurement is exited.
func (s *BaseA2LListener) ExitOutMeasurement(ctx *OutMeasurementContext) {}

// EnterSubFunction is called when production subFunction is entered.
func (s *BaseA2LListener) EnterSubFunction(ctx *SubFunctionContext) {}

// ExitSubFunction is called when production subFunction is exited.
func (s *BaseA2LListener) ExitSubFunction(ctx *SubFunctionContext) {}

// EnterGroup is called when production group is entered.
func (s *BaseA2LListener) EnterGroup(ctx *GroupContext) {}

// ExitGroup is called when production group is exited.
func (s *BaseA2LListener) ExitGroup(ctx *GroupContext) {}

// EnterRefMeasurement is called when production refMeasurement is entered.
func (s *BaseA2LListener) EnterRefMeasurement(ctx *RefMeasurementContext) {}

// ExitRefMeasurement is called when production refMeasurement is exited.
func (s *BaseA2LListener) ExitRefMeasurement(ctx *RefMeasurementContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseA2LListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseA2LListener) ExitRoot(ctx *RootContext) {}

// EnterSubGroup is called when production subGroup is entered.
func (s *BaseA2LListener) EnterSubGroup(ctx *SubGroupContext) {}

// ExitSubGroup is called when production subGroup is exited.
func (s *BaseA2LListener) ExitSubGroup(ctx *SubGroupContext) {}

// EnterInstance is called when production instance is entered.
func (s *BaseA2LListener) EnterInstance(ctx *InstanceContext) {}

// ExitInstance is called when production instance is exited.
func (s *BaseA2LListener) ExitInstance(ctx *InstanceContext) {}

// EnterMeasurement is called when production measurement is entered.
func (s *BaseA2LListener) EnterMeasurement(ctx *MeasurementContext) {}

// ExitMeasurement is called when production measurement is exited.
func (s *BaseA2LListener) ExitMeasurement(ctx *MeasurementContext) {}

// EnterArraySize is called when production arraySize is entered.
func (s *BaseA2LListener) EnterArraySize(ctx *ArraySizeContext) {}

// ExitArraySize is called when production arraySize is exited.
func (s *BaseA2LListener) ExitArraySize(ctx *ArraySizeContext) {}

// EnterBitOperation is called when production bitOperation is entered.
func (s *BaseA2LListener) EnterBitOperation(ctx *BitOperationContext) {}

// ExitBitOperation is called when production bitOperation is exited.
func (s *BaseA2LListener) ExitBitOperation(ctx *BitOperationContext) {}

// EnterLeftShift is called when production leftShift is entered.
func (s *BaseA2LListener) EnterLeftShift(ctx *LeftShiftContext) {}

// ExitLeftShift is called when production leftShift is exited.
func (s *BaseA2LListener) ExitLeftShift(ctx *LeftShiftContext) {}

// EnterRightShift is called when production rightShift is entered.
func (s *BaseA2LListener) EnterRightShift(ctx *RightShiftContext) {}

// ExitRightShift is called when production rightShift is exited.
func (s *BaseA2LListener) ExitRightShift(ctx *RightShiftContext) {}

// EnterSignExtend is called when production signExtend is entered.
func (s *BaseA2LListener) EnterSignExtend(ctx *SignExtendContext) {}

// ExitSignExtend is called when production signExtend is exited.
func (s *BaseA2LListener) ExitSignExtend(ctx *SignExtendContext) {}

// EnterEcuAddress is called when production ecuAddress is entered.
func (s *BaseA2LListener) EnterEcuAddress(ctx *EcuAddressContext) {}

// ExitEcuAddress is called when production ecuAddress is exited.
func (s *BaseA2LListener) ExitEcuAddress(ctx *EcuAddressContext) {}

// EnterErrorMask is called when production errorMask is entered.
func (s *BaseA2LListener) EnterErrorMask(ctx *ErrorMaskContext) {}

// ExitErrorMask is called when production errorMask is exited.
func (s *BaseA2LListener) ExitErrorMask(ctx *ErrorMaskContext) {}

// EnterLayout is called when production layout is entered.
func (s *BaseA2LListener) EnterLayout(ctx *LayoutContext) {}

// ExitLayout is called when production layout is exited.
func (s *BaseA2LListener) ExitLayout(ctx *LayoutContext) {}

// EnterReadWrite is called when production readWrite is entered.
func (s *BaseA2LListener) EnterReadWrite(ctx *ReadWriteContext) {}

// ExitReadWrite is called when production readWrite is exited.
func (s *BaseA2LListener) ExitReadWrite(ctx *ReadWriteContext) {}

// EnterVirtual is called when production virtual is entered.
func (s *BaseA2LListener) EnterVirtual(ctx *VirtualContext) {}

// ExitVirtual is called when production virtual is exited.
func (s *BaseA2LListener) ExitVirtual(ctx *VirtualContext) {}

// EnterModCommon is called when production modCommon is entered.
func (s *BaseA2LListener) EnterModCommon(ctx *ModCommonContext) {}

// ExitModCommon is called when production modCommon is exited.
func (s *BaseA2LListener) ExitModCommon(ctx *ModCommonContext) {}

// EnterDataSize is called when production dataSize is entered.
func (s *BaseA2LListener) EnterDataSize(ctx *DataSizeContext) {}

// ExitDataSize is called when production dataSize is exited.
func (s *BaseA2LListener) ExitDataSize(ctx *DataSizeContext) {}

// EnterSRecLayout is called when production sRecLayout is entered.
func (s *BaseA2LListener) EnterSRecLayout(ctx *SRecLayoutContext) {}

// ExitSRecLayout is called when production sRecLayout is exited.
func (s *BaseA2LListener) ExitSRecLayout(ctx *SRecLayoutContext) {}

// EnterModPar is called when production modPar is entered.
func (s *BaseA2LListener) EnterModPar(ctx *ModParContext) {}

// ExitModPar is called when production modPar is exited.
func (s *BaseA2LListener) ExitModPar(ctx *ModParContext) {}

// EnterAddrEpk is called when production addrEpk is entered.
func (s *BaseA2LListener) EnterAddrEpk(ctx *AddrEpkContext) {}

// ExitAddrEpk is called when production addrEpk is exited.
func (s *BaseA2LListener) ExitAddrEpk(ctx *AddrEpkContext) {}

// EnterCalibrationMethod is called when production calibrationMethod is entered.
func (s *BaseA2LListener) EnterCalibrationMethod(ctx *CalibrationMethodContext) {}

// ExitCalibrationMethod is called when production calibrationMethod is exited.
func (s *BaseA2LListener) ExitCalibrationMethod(ctx *CalibrationMethodContext) {}

// EnterCalibrationHandle is called when production calibrationHandle is entered.
func (s *BaseA2LListener) EnterCalibrationHandle(ctx *CalibrationHandleContext) {}

// ExitCalibrationHandle is called when production calibrationHandle is exited.
func (s *BaseA2LListener) ExitCalibrationHandle(ctx *CalibrationHandleContext) {}

// EnterCpuType is called when production cpuType is entered.
func (s *BaseA2LListener) EnterCpuType(ctx *CpuTypeContext) {}

// ExitCpuType is called when production cpuType is exited.
func (s *BaseA2LListener) ExitCpuType(ctx *CpuTypeContext) {}

// EnterCustomer is called when production customer is entered.
func (s *BaseA2LListener) EnterCustomer(ctx *CustomerContext) {}

// ExitCustomer is called when production customer is exited.
func (s *BaseA2LListener) ExitCustomer(ctx *CustomerContext) {}

// EnterCustomerNo is called when production customerNo is entered.
func (s *BaseA2LListener) EnterCustomerNo(ctx *CustomerNoContext) {}

// ExitCustomerNo is called when production customerNo is exited.
func (s *BaseA2LListener) ExitCustomerNo(ctx *CustomerNoContext) {}

// EnterEcu is called when production ecu is entered.
func (s *BaseA2LListener) EnterEcu(ctx *EcuContext) {}

// ExitEcu is called when production ecu is exited.
func (s *BaseA2LListener) ExitEcu(ctx *EcuContext) {}

// EnterEcuCalibrationOffset is called when production ecuCalibrationOffset is entered.
func (s *BaseA2LListener) EnterEcuCalibrationOffset(ctx *EcuCalibrationOffsetContext) {}

// ExitEcuCalibrationOffset is called when production ecuCalibrationOffset is exited.
func (s *BaseA2LListener) ExitEcuCalibrationOffset(ctx *EcuCalibrationOffsetContext) {}

// EnterEpk is called when production epk is entered.
func (s *BaseA2LListener) EnterEpk(ctx *EpkContext) {}

// ExitEpk is called when production epk is exited.
func (s *BaseA2LListener) ExitEpk(ctx *EpkContext) {}

// EnterMemoryLayout is called when production memoryLayout is entered.
func (s *BaseA2LListener) EnterMemoryLayout(ctx *MemoryLayoutContext) {}

// ExitMemoryLayout is called when production memoryLayout is exited.
func (s *BaseA2LListener) ExitMemoryLayout(ctx *MemoryLayoutContext) {}

// EnterMemorySegment is called when production memorySegment is entered.
func (s *BaseA2LListener) EnterMemorySegment(ctx *MemorySegmentContext) {}

// ExitMemorySegment is called when production memorySegment is exited.
func (s *BaseA2LListener) ExitMemorySegment(ctx *MemorySegmentContext) {}

// EnterNoOfInterfaces is called when production noOfInterfaces is entered.
func (s *BaseA2LListener) EnterNoOfInterfaces(ctx *NoOfInterfacesContext) {}

// ExitNoOfInterfaces is called when production noOfInterfaces is exited.
func (s *BaseA2LListener) ExitNoOfInterfaces(ctx *NoOfInterfacesContext) {}

// EnterPhoneNo is called when production phoneNo is entered.
func (s *BaseA2LListener) EnterPhoneNo(ctx *PhoneNoContext) {}

// ExitPhoneNo is called when production phoneNo is exited.
func (s *BaseA2LListener) ExitPhoneNo(ctx *PhoneNoContext) {}

// EnterSupplier is called when production supplier is entered.
func (s *BaseA2LListener) EnterSupplier(ctx *SupplierContext) {}

// ExitSupplier is called when production supplier is exited.
func (s *BaseA2LListener) ExitSupplier(ctx *SupplierContext) {}

// EnterSystemConstant is called when production systemConstant is entered.
func (s *BaseA2LListener) EnterSystemConstant(ctx *SystemConstantContext) {}

// ExitSystemConstant is called when production systemConstant is exited.
func (s *BaseA2LListener) ExitSystemConstant(ctx *SystemConstantContext) {}

// EnterUser is called when production user is entered.
func (s *BaseA2LListener) EnterUser(ctx *UserContext) {}

// ExitUser is called when production user is exited.
func (s *BaseA2LListener) ExitUser(ctx *UserContext) {}

// EnterRecordLayout is called when production recordLayout is entered.
func (s *BaseA2LListener) EnterRecordLayout(ctx *RecordLayoutContext) {}

// ExitRecordLayout is called when production recordLayout is exited.
func (s *BaseA2LListener) ExitRecordLayout(ctx *RecordLayoutContext) {}

// EnterAxisPtsX is called when production axisPtsX is entered.
func (s *BaseA2LListener) EnterAxisPtsX(ctx *AxisPtsXContext) {}

// ExitAxisPtsX is called when production axisPtsX is exited.
func (s *BaseA2LListener) ExitAxisPtsX(ctx *AxisPtsXContext) {}

// EnterAxisPtsY is called when production axisPtsY is entered.
func (s *BaseA2LListener) EnterAxisPtsY(ctx *AxisPtsYContext) {}

// ExitAxisPtsY is called when production axisPtsY is exited.
func (s *BaseA2LListener) ExitAxisPtsY(ctx *AxisPtsYContext) {}

// EnterAxisPtsZ is called when production axisPtsZ is entered.
func (s *BaseA2LListener) EnterAxisPtsZ(ctx *AxisPtsZContext) {}

// ExitAxisPtsZ is called when production axisPtsZ is exited.
func (s *BaseA2LListener) ExitAxisPtsZ(ctx *AxisPtsZContext) {}

// EnterAxisPts4 is called when production axisPts4 is entered.
func (s *BaseA2LListener) EnterAxisPts4(ctx *AxisPts4Context) {}

// ExitAxisPts4 is called when production axisPts4 is exited.
func (s *BaseA2LListener) ExitAxisPts4(ctx *AxisPts4Context) {}

// EnterAxisPts5 is called when production axisPts5 is entered.
func (s *BaseA2LListener) EnterAxisPts5(ctx *AxisPts5Context) {}

// ExitAxisPts5 is called when production axisPts5 is exited.
func (s *BaseA2LListener) ExitAxisPts5(ctx *AxisPts5Context) {}

// EnterAxisRescaleX is called when production axisRescaleX is entered.
func (s *BaseA2LListener) EnterAxisRescaleX(ctx *AxisRescaleXContext) {}

// ExitAxisRescaleX is called when production axisRescaleX is exited.
func (s *BaseA2LListener) ExitAxisRescaleX(ctx *AxisRescaleXContext) {}

// EnterAxisRescaleY is called when production axisRescaleY is entered.
func (s *BaseA2LListener) EnterAxisRescaleY(ctx *AxisRescaleYContext) {}

// ExitAxisRescaleY is called when production axisRescaleY is exited.
func (s *BaseA2LListener) ExitAxisRescaleY(ctx *AxisRescaleYContext) {}

// EnterAxisRescaleZ is called when production axisRescaleZ is entered.
func (s *BaseA2LListener) EnterAxisRescaleZ(ctx *AxisRescaleZContext) {}

// ExitAxisRescaleZ is called when production axisRescaleZ is exited.
func (s *BaseA2LListener) ExitAxisRescaleZ(ctx *AxisRescaleZContext) {}

// EnterAxisRescale4 is called when production axisRescale4 is entered.
func (s *BaseA2LListener) EnterAxisRescale4(ctx *AxisRescale4Context) {}

// ExitAxisRescale4 is called when production axisRescale4 is exited.
func (s *BaseA2LListener) ExitAxisRescale4(ctx *AxisRescale4Context) {}

// EnterAxisRescale5 is called when production axisRescale5 is entered.
func (s *BaseA2LListener) EnterAxisRescale5(ctx *AxisRescale5Context) {}

// ExitAxisRescale5 is called when production axisRescale5 is exited.
func (s *BaseA2LListener) ExitAxisRescale5(ctx *AxisRescale5Context) {}

// EnterDistOpX is called when production distOpX is entered.
func (s *BaseA2LListener) EnterDistOpX(ctx *DistOpXContext) {}

// ExitDistOpX is called when production distOpX is exited.
func (s *BaseA2LListener) ExitDistOpX(ctx *DistOpXContext) {}

// EnterDistOpY is called when production distOpY is entered.
func (s *BaseA2LListener) EnterDistOpY(ctx *DistOpYContext) {}

// ExitDistOpY is called when production distOpY is exited.
func (s *BaseA2LListener) ExitDistOpY(ctx *DistOpYContext) {}

// EnterDistOpZ is called when production distOpZ is entered.
func (s *BaseA2LListener) EnterDistOpZ(ctx *DistOpZContext) {}

// ExitDistOpZ is called when production distOpZ is exited.
func (s *BaseA2LListener) ExitDistOpZ(ctx *DistOpZContext) {}

// EnterDistOp4 is called when production distOp4 is entered.
func (s *BaseA2LListener) EnterDistOp4(ctx *DistOp4Context) {}

// ExitDistOp4 is called when production distOp4 is exited.
func (s *BaseA2LListener) ExitDistOp4(ctx *DistOp4Context) {}

// EnterDistOp5 is called when production distOp5 is entered.
func (s *BaseA2LListener) EnterDistOp5(ctx *DistOp5Context) {}

// ExitDistOp5 is called when production distOp5 is exited.
func (s *BaseA2LListener) ExitDistOp5(ctx *DistOp5Context) {}

// EnterFixNoAxisPtsX is called when production fixNoAxisPtsX is entered.
func (s *BaseA2LListener) EnterFixNoAxisPtsX(ctx *FixNoAxisPtsXContext) {}

// ExitFixNoAxisPtsX is called when production fixNoAxisPtsX is exited.
func (s *BaseA2LListener) ExitFixNoAxisPtsX(ctx *FixNoAxisPtsXContext) {}

// EnterFixNoAxisPtsY is called when production fixNoAxisPtsY is entered.
func (s *BaseA2LListener) EnterFixNoAxisPtsY(ctx *FixNoAxisPtsYContext) {}

// ExitFixNoAxisPtsY is called when production fixNoAxisPtsY is exited.
func (s *BaseA2LListener) ExitFixNoAxisPtsY(ctx *FixNoAxisPtsYContext) {}

// EnterFixNoAxisPtsZ is called when production fixNoAxisPtsZ is entered.
func (s *BaseA2LListener) EnterFixNoAxisPtsZ(ctx *FixNoAxisPtsZContext) {}

// ExitFixNoAxisPtsZ is called when production fixNoAxisPtsZ is exited.
func (s *BaseA2LListener) ExitFixNoAxisPtsZ(ctx *FixNoAxisPtsZContext) {}

// EnterFixNoAxisPts4 is called when production fixNoAxisPts4 is entered.
func (s *BaseA2LListener) EnterFixNoAxisPts4(ctx *FixNoAxisPts4Context) {}

// ExitFixNoAxisPts4 is called when production fixNoAxisPts4 is exited.
func (s *BaseA2LListener) ExitFixNoAxisPts4(ctx *FixNoAxisPts4Context) {}

// EnterFixNoAxisPts5 is called when production fixNoAxisPts5 is entered.
func (s *BaseA2LListener) EnterFixNoAxisPts5(ctx *FixNoAxisPts5Context) {}

// ExitFixNoAxisPts5 is called when production fixNoAxisPts5 is exited.
func (s *BaseA2LListener) ExitFixNoAxisPts5(ctx *FixNoAxisPts5Context) {}

// EnterFncValues is called when production fncValues is entered.
func (s *BaseA2LListener) EnterFncValues(ctx *FncValuesContext) {}

// ExitFncValues is called when production fncValues is exited.
func (s *BaseA2LListener) ExitFncValues(ctx *FncValuesContext) {}

// EnterIdentification is called when production identification is entered.
func (s *BaseA2LListener) EnterIdentification(ctx *IdentificationContext) {}

// ExitIdentification is called when production identification is exited.
func (s *BaseA2LListener) ExitIdentification(ctx *IdentificationContext) {}

// EnterNoAxisPtsX is called when production noAxisPtsX is entered.
func (s *BaseA2LListener) EnterNoAxisPtsX(ctx *NoAxisPtsXContext) {}

// ExitNoAxisPtsX is called when production noAxisPtsX is exited.
func (s *BaseA2LListener) ExitNoAxisPtsX(ctx *NoAxisPtsXContext) {}

// EnterNoAxisPtsY is called when production noAxisPtsY is entered.
func (s *BaseA2LListener) EnterNoAxisPtsY(ctx *NoAxisPtsYContext) {}

// ExitNoAxisPtsY is called when production noAxisPtsY is exited.
func (s *BaseA2LListener) ExitNoAxisPtsY(ctx *NoAxisPtsYContext) {}

// EnterNoAxisPtsZ is called when production noAxisPtsZ is entered.
func (s *BaseA2LListener) EnterNoAxisPtsZ(ctx *NoAxisPtsZContext) {}

// ExitNoAxisPtsZ is called when production noAxisPtsZ is exited.
func (s *BaseA2LListener) ExitNoAxisPtsZ(ctx *NoAxisPtsZContext) {}

// EnterNoAxisPts4 is called when production noAxisPts4 is entered.
func (s *BaseA2LListener) EnterNoAxisPts4(ctx *NoAxisPts4Context) {}

// ExitNoAxisPts4 is called when production noAxisPts4 is exited.
func (s *BaseA2LListener) ExitNoAxisPts4(ctx *NoAxisPts4Context) {}

// EnterNoAxisPts5 is called when production noAxisPts5 is entered.
func (s *BaseA2LListener) EnterNoAxisPts5(ctx *NoAxisPts5Context) {}

// ExitNoAxisPts5 is called when production noAxisPts5 is exited.
func (s *BaseA2LListener) ExitNoAxisPts5(ctx *NoAxisPts5Context) {}

// EnterStaticRecordLayout is called when production staticRecordLayout is entered.
func (s *BaseA2LListener) EnterStaticRecordLayout(ctx *StaticRecordLayoutContext) {}

// ExitStaticRecordLayout is called when production staticRecordLayout is exited.
func (s *BaseA2LListener) ExitStaticRecordLayout(ctx *StaticRecordLayoutContext) {}

// EnterNoRescaleX is called when production noRescaleX is entered.
func (s *BaseA2LListener) EnterNoRescaleX(ctx *NoRescaleXContext) {}

// ExitNoRescaleX is called when production noRescaleX is exited.
func (s *BaseA2LListener) ExitNoRescaleX(ctx *NoRescaleXContext) {}

// EnterNoRescaleY is called when production noRescaleY is entered.
func (s *BaseA2LListener) EnterNoRescaleY(ctx *NoRescaleYContext) {}

// ExitNoRescaleY is called when production noRescaleY is exited.
func (s *BaseA2LListener) ExitNoRescaleY(ctx *NoRescaleYContext) {}

// EnterNoRescaleZ is called when production noRescaleZ is entered.
func (s *BaseA2LListener) EnterNoRescaleZ(ctx *NoRescaleZContext) {}

// ExitNoRescaleZ is called when production noRescaleZ is exited.
func (s *BaseA2LListener) ExitNoRescaleZ(ctx *NoRescaleZContext) {}

// EnterNoRescale4 is called when production noRescale4 is entered.
func (s *BaseA2LListener) EnterNoRescale4(ctx *NoRescale4Context) {}

// ExitNoRescale4 is called when production noRescale4 is exited.
func (s *BaseA2LListener) ExitNoRescale4(ctx *NoRescale4Context) {}

// EnterNoRescale5 is called when production noRescale5 is entered.
func (s *BaseA2LListener) EnterNoRescale5(ctx *NoRescale5Context) {}

// ExitNoRescale5 is called when production noRescale5 is exited.
func (s *BaseA2LListener) ExitNoRescale5(ctx *NoRescale5Context) {}

// EnterOffsetX is called when production offsetX is entered.
func (s *BaseA2LListener) EnterOffsetX(ctx *OffsetXContext) {}

// ExitOffsetX is called when production offsetX is exited.
func (s *BaseA2LListener) ExitOffsetX(ctx *OffsetXContext) {}

// EnterOffsetY is called when production offsetY is entered.
func (s *BaseA2LListener) EnterOffsetY(ctx *OffsetYContext) {}

// ExitOffsetY is called when production offsetY is exited.
func (s *BaseA2LListener) ExitOffsetY(ctx *OffsetYContext) {}

// EnterOffsetZ is called when production offsetZ is entered.
func (s *BaseA2LListener) EnterOffsetZ(ctx *OffsetZContext) {}

// ExitOffsetZ is called when production offsetZ is exited.
func (s *BaseA2LListener) ExitOffsetZ(ctx *OffsetZContext) {}

// EnterOffset4 is called when production offset4 is entered.
func (s *BaseA2LListener) EnterOffset4(ctx *Offset4Context) {}

// ExitOffset4 is called when production offset4 is exited.
func (s *BaseA2LListener) ExitOffset4(ctx *Offset4Context) {}

// EnterOffset5 is called when production offset5 is entered.
func (s *BaseA2LListener) EnterOffset5(ctx *Offset5Context) {}

// ExitOffset5 is called when production offset5 is exited.
func (s *BaseA2LListener) ExitOffset5(ctx *Offset5Context) {}

// EnterReserved is called when production reserved is entered.
func (s *BaseA2LListener) EnterReserved(ctx *ReservedContext) {}

// ExitReserved is called when production reserved is exited.
func (s *BaseA2LListener) ExitReserved(ctx *ReservedContext) {}

// EnterRipAddrW is called when production ripAddrW is entered.
func (s *BaseA2LListener) EnterRipAddrW(ctx *RipAddrWContext) {}

// ExitRipAddrW is called when production ripAddrW is exited.
func (s *BaseA2LListener) ExitRipAddrW(ctx *RipAddrWContext) {}

// EnterRipAddrX is called when production ripAddrX is entered.
func (s *BaseA2LListener) EnterRipAddrX(ctx *RipAddrXContext) {}

// ExitRipAddrX is called when production ripAddrX is exited.
func (s *BaseA2LListener) ExitRipAddrX(ctx *RipAddrXContext) {}

// EnterRipAddrY is called when production ripAddrY is entered.
func (s *BaseA2LListener) EnterRipAddrY(ctx *RipAddrYContext) {}

// ExitRipAddrY is called when production ripAddrY is exited.
func (s *BaseA2LListener) ExitRipAddrY(ctx *RipAddrYContext) {}

// EnterRipAddrZ is called when production ripAddrZ is entered.
func (s *BaseA2LListener) EnterRipAddrZ(ctx *RipAddrZContext) {}

// ExitRipAddrZ is called when production ripAddrZ is exited.
func (s *BaseA2LListener) ExitRipAddrZ(ctx *RipAddrZContext) {}

// EnterRipAddr4 is called when production ripAddr4 is entered.
func (s *BaseA2LListener) EnterRipAddr4(ctx *RipAddr4Context) {}

// ExitRipAddr4 is called when production ripAddr4 is exited.
func (s *BaseA2LListener) ExitRipAddr4(ctx *RipAddr4Context) {}

// EnterRipAddr5 is called when production ripAddr5 is entered.
func (s *BaseA2LListener) EnterRipAddr5(ctx *RipAddr5Context) {}

// ExitRipAddr5 is called when production ripAddr5 is exited.
func (s *BaseA2LListener) ExitRipAddr5(ctx *RipAddr5Context) {}

// EnterShiftOpX is called when production shiftOpX is entered.
func (s *BaseA2LListener) EnterShiftOpX(ctx *ShiftOpXContext) {}

// ExitShiftOpX is called when production shiftOpX is exited.
func (s *BaseA2LListener) ExitShiftOpX(ctx *ShiftOpXContext) {}

// EnterShiftOpY is called when production shiftOpY is entered.
func (s *BaseA2LListener) EnterShiftOpY(ctx *ShiftOpYContext) {}

// ExitShiftOpY is called when production shiftOpY is exited.
func (s *BaseA2LListener) ExitShiftOpY(ctx *ShiftOpYContext) {}

// EnterShiftOpZ is called when production shiftOpZ is entered.
func (s *BaseA2LListener) EnterShiftOpZ(ctx *ShiftOpZContext) {}

// ExitShiftOpZ is called when production shiftOpZ is exited.
func (s *BaseA2LListener) ExitShiftOpZ(ctx *ShiftOpZContext) {}

// EnterShiftOp4 is called when production shiftOp4 is entered.
func (s *BaseA2LListener) EnterShiftOp4(ctx *ShiftOp4Context) {}

// ExitShiftOp4 is called when production shiftOp4 is exited.
func (s *BaseA2LListener) ExitShiftOp4(ctx *ShiftOp4Context) {}

// EnterShiftOp5 is called when production shiftOp5 is entered.
func (s *BaseA2LListener) EnterShiftOp5(ctx *ShiftOp5Context) {}

// ExitShiftOp5 is called when production shiftOp5 is exited.
func (s *BaseA2LListener) ExitShiftOp5(ctx *ShiftOp5Context) {}

// EnterSrcAddrX is called when production srcAddrX is entered.
func (s *BaseA2LListener) EnterSrcAddrX(ctx *SrcAddrXContext) {}

// ExitSrcAddrX is called when production srcAddrX is exited.
func (s *BaseA2LListener) ExitSrcAddrX(ctx *SrcAddrXContext) {}

// EnterSrcAddrY is called when production srcAddrY is entered.
func (s *BaseA2LListener) EnterSrcAddrY(ctx *SrcAddrYContext) {}

// ExitSrcAddrY is called when production srcAddrY is exited.
func (s *BaseA2LListener) ExitSrcAddrY(ctx *SrcAddrYContext) {}

// EnterSrcAddrZ is called when production srcAddrZ is entered.
func (s *BaseA2LListener) EnterSrcAddrZ(ctx *SrcAddrZContext) {}

// ExitSrcAddrZ is called when production srcAddrZ is exited.
func (s *BaseA2LListener) ExitSrcAddrZ(ctx *SrcAddrZContext) {}

// EnterSrcAddr4 is called when production srcAddr4 is entered.
func (s *BaseA2LListener) EnterSrcAddr4(ctx *SrcAddr4Context) {}

// ExitSrcAddr4 is called when production srcAddr4 is exited.
func (s *BaseA2LListener) ExitSrcAddr4(ctx *SrcAddr4Context) {}

// EnterSrcAddr5 is called when production srcAddr5 is entered.
func (s *BaseA2LListener) EnterSrcAddr5(ctx *SrcAddr5Context) {}

// ExitSrcAddr5 is called when production srcAddr5 is exited.
func (s *BaseA2LListener) ExitSrcAddr5(ctx *SrcAddr5Context) {}

// EnterTypedefCharacteristic is called when production typedefCharacteristic is entered.
func (s *BaseA2LListener) EnterTypedefCharacteristic(ctx *TypedefCharacteristicContext) {}

// ExitTypedefCharacteristic is called when production typedefCharacteristic is exited.
func (s *BaseA2LListener) ExitTypedefCharacteristic(ctx *TypedefCharacteristicContext) {}

// EnterTypedefMeasurement is called when production typedefMeasurement is entered.
func (s *BaseA2LListener) EnterTypedefMeasurement(ctx *TypedefMeasurementContext) {}

// ExitTypedefMeasurement is called when production typedefMeasurement is exited.
func (s *BaseA2LListener) ExitTypedefMeasurement(ctx *TypedefMeasurementContext) {}

// EnterTypedefStructure is called when production typedefStructure is entered.
func (s *BaseA2LListener) EnterTypedefStructure(ctx *TypedefStructureContext) {}

// ExitTypedefStructure is called when production typedefStructure is exited.
func (s *BaseA2LListener) ExitTypedefStructure(ctx *TypedefStructureContext) {}

// EnterStructureComponent is called when production structureComponent is entered.
func (s *BaseA2LListener) EnterStructureComponent(ctx *StructureComponentContext) {}

// ExitStructureComponent is called when production structureComponent is exited.
func (s *BaseA2LListener) ExitStructureComponent(ctx *StructureComponentContext) {}

// EnterUnit is called when production unit is entered.
func (s *BaseA2LListener) EnterUnit(ctx *UnitContext) {}

// ExitUnit is called when production unit is exited.
func (s *BaseA2LListener) ExitUnit(ctx *UnitContext) {}

// EnterSiExponents is called when production siExponents is entered.
func (s *BaseA2LListener) EnterSiExponents(ctx *SiExponentsContext) {}

// ExitSiExponents is called when production siExponents is exited.
func (s *BaseA2LListener) ExitSiExponents(ctx *SiExponentsContext) {}

// EnterUnitConversion is called when production unitConversion is entered.
func (s *BaseA2LListener) EnterUnitConversion(ctx *UnitConversionContext) {}

// ExitUnitConversion is called when production unitConversion is exited.
func (s *BaseA2LListener) ExitUnitConversion(ctx *UnitConversionContext) {}

// EnterUserRights is called when production userRights is entered.
func (s *BaseA2LListener) EnterUserRights(ctx *UserRightsContext) {}

// ExitUserRights is called when production userRights is exited.
func (s *BaseA2LListener) ExitUserRights(ctx *UserRightsContext) {}

// EnterRefGroup is called when production refGroup is entered.
func (s *BaseA2LListener) EnterRefGroup(ctx *RefGroupContext) {}

// ExitRefGroup is called when production refGroup is exited.
func (s *BaseA2LListener) ExitRefGroup(ctx *RefGroupContext) {}

// EnterVariantCoding is called when production variantCoding is entered.
func (s *BaseA2LListener) EnterVariantCoding(ctx *VariantCodingContext) {}

// ExitVariantCoding is called when production variantCoding is exited.
func (s *BaseA2LListener) ExitVariantCoding(ctx *VariantCodingContext) {}

// EnterVarCharacteristic is called when production varCharacteristic is entered.
func (s *BaseA2LListener) EnterVarCharacteristic(ctx *VarCharacteristicContext) {}

// ExitVarCharacteristic is called when production varCharacteristic is exited.
func (s *BaseA2LListener) ExitVarCharacteristic(ctx *VarCharacteristicContext) {}

// EnterVarAddress is called when production varAddress is entered.
func (s *BaseA2LListener) EnterVarAddress(ctx *VarAddressContext) {}

// ExitVarAddress is called when production varAddress is exited.
func (s *BaseA2LListener) ExitVarAddress(ctx *VarAddressContext) {}

// EnterVarCriterion is called when production varCriterion is entered.
func (s *BaseA2LListener) EnterVarCriterion(ctx *VarCriterionContext) {}

// ExitVarCriterion is called when production varCriterion is exited.
func (s *BaseA2LListener) ExitVarCriterion(ctx *VarCriterionContext) {}

// EnterVarMeasurement is called when production varMeasurement is entered.
func (s *BaseA2LListener) EnterVarMeasurement(ctx *VarMeasurementContext) {}

// ExitVarMeasurement is called when production varMeasurement is exited.
func (s *BaseA2LListener) ExitVarMeasurement(ctx *VarMeasurementContext) {}

// EnterVarSelectionCharacteristic is called when production varSelectionCharacteristic is entered.
func (s *BaseA2LListener) EnterVarSelectionCharacteristic(ctx *VarSelectionCharacteristicContext) {}

// ExitVarSelectionCharacteristic is called when production varSelectionCharacteristic is exited.
func (s *BaseA2LListener) ExitVarSelectionCharacteristic(ctx *VarSelectionCharacteristicContext) {}

// EnterVarForbiddenComb is called when production varForbiddenComb is entered.
func (s *BaseA2LListener) EnterVarForbiddenComb(ctx *VarForbiddenCombContext) {}

// ExitVarForbiddenComb is called when production varForbiddenComb is exited.
func (s *BaseA2LListener) ExitVarForbiddenComb(ctx *VarForbiddenCombContext) {}

// EnterVarNaming is called when production varNaming is entered.
func (s *BaseA2LListener) EnterVarNaming(ctx *VarNamingContext) {}

// ExitVarNaming is called when production varNaming is exited.
func (s *BaseA2LListener) ExitVarNaming(ctx *VarNamingContext) {}

// EnterVarSeparator is called when production varSeparator is entered.
func (s *BaseA2LListener) EnterVarSeparator(ctx *VarSeparatorContext) {}

// ExitVarSeparator is called when production varSeparator is exited.
func (s *BaseA2LListener) ExitVarSeparator(ctx *VarSeparatorContext) {}

// EnterIntegerValue is called when production integerValue is entered.
func (s *BaseA2LListener) EnterIntegerValue(ctx *IntegerValueContext) {}

// ExitIntegerValue is called when production integerValue is exited.
func (s *BaseA2LListener) ExitIntegerValue(ctx *IntegerValueContext) {}

// EnterNumericValue is called when production numericValue is entered.
func (s *BaseA2LListener) EnterNumericValue(ctx *NumericValueContext) {}

// ExitNumericValue is called when production numericValue is exited.
func (s *BaseA2LListener) ExitNumericValue(ctx *NumericValueContext) {}

// EnterStringValue is called when production stringValue is entered.
func (s *BaseA2LListener) EnterStringValue(ctx *StringValueContext) {}

// ExitStringValue is called when production stringValue is exited.
func (s *BaseA2LListener) ExitStringValue(ctx *StringValueContext) {}

// EnterIdentifierValue is called when production identifierValue is entered.
func (s *BaseA2LListener) EnterIdentifierValue(ctx *IdentifierValueContext) {}

// ExitIdentifierValue is called when production identifierValue is exited.
func (s *BaseA2LListener) ExitIdentifierValue(ctx *IdentifierValueContext) {}

// EnterPartialIdentifier is called when production partialIdentifier is entered.
func (s *BaseA2LListener) EnterPartialIdentifier(ctx *PartialIdentifierContext) {}

// ExitPartialIdentifier is called when production partialIdentifier is exited.
func (s *BaseA2LListener) ExitPartialIdentifier(ctx *PartialIdentifierContext) {}

// EnterArraySpecifier is called when production arraySpecifier is entered.
func (s *BaseA2LListener) EnterArraySpecifier(ctx *ArraySpecifierContext) {}

// ExitArraySpecifier is called when production arraySpecifier is exited.
func (s *BaseA2LListener) ExitArraySpecifier(ctx *ArraySpecifierContext) {}

// EnterDataType is called when production dataType is entered.
func (s *BaseA2LListener) EnterDataType(ctx *DataTypeContext) {}

// ExitDataType is called when production dataType is exited.
func (s *BaseA2LListener) ExitDataType(ctx *DataTypeContext) {}

// EnterDatasize is called when production datasize is entered.
func (s *BaseA2LListener) EnterDatasize(ctx *DatasizeContext) {}

// ExitDatasize is called when production datasize is exited.
func (s *BaseA2LListener) ExitDatasize(ctx *DatasizeContext) {}

// EnterAddrtype is called when production addrtype is entered.
func (s *BaseA2LListener) EnterAddrtype(ctx *AddrtypeContext) {}

// ExitAddrtype is called when production addrtype is exited.
func (s *BaseA2LListener) ExitAddrtype(ctx *AddrtypeContext) {}

// EnterByteOrderValue is called when production byteOrderValue is entered.
func (s *BaseA2LListener) EnterByteOrderValue(ctx *ByteOrderValueContext) {}

// ExitByteOrderValue is called when production byteOrderValue is exited.
func (s *BaseA2LListener) ExitByteOrderValue(ctx *ByteOrderValueContext) {}

// EnterIndexorder is called when production indexorder is entered.
func (s *BaseA2LListener) EnterIndexorder(ctx *IndexorderContext) {}

// ExitIndexorder is called when production indexorder is exited.
func (s *BaseA2LListener) ExitIndexorder(ctx *IndexorderContext) {}

// EnterLinkType is called when production linkType is entered.
func (s *BaseA2LListener) EnterLinkType(ctx *LinkTypeContext) {}

// ExitLinkType is called when production linkType is exited.
func (s *BaseA2LListener) ExitLinkType(ctx *LinkTypeContext) {}

// EnterA2ml is called when production a2ml is entered.
func (s *BaseA2LListener) EnterA2ml(ctx *A2mlContext) {}

// ExitA2ml is called when production a2ml is exited.
func (s *BaseA2LListener) ExitA2ml(ctx *A2mlContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseA2LListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseA2LListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterTypeDefinition is called when production typeDefinition is entered.
func (s *BaseA2LListener) EnterTypeDefinition(ctx *TypeDefinitionContext) {}

// ExitTypeDefinition is called when production typeDefinition is exited.
func (s *BaseA2LListener) ExitTypeDefinition(ctx *TypeDefinitionContext) {}

// EnterA2mlTypeName is called when production a2mlTypeName is entered.
func (s *BaseA2LListener) EnterA2mlTypeName(ctx *A2mlTypeNameContext) {}

// ExitA2mlTypeName is called when production a2mlTypeName is exited.
func (s *BaseA2LListener) ExitA2mlTypeName(ctx *A2mlTypeNameContext) {}

// EnterPredefinedTypeName is called when production predefinedTypeName is entered.
func (s *BaseA2LListener) EnterPredefinedTypeName(ctx *PredefinedTypeNameContext) {}

// ExitPredefinedTypeName is called when production predefinedTypeName is exited.
func (s *BaseA2LListener) ExitPredefinedTypeName(ctx *PredefinedTypeNameContext) {}

// EnterBlockDefinition is called when production blockDefinition is entered.
func (s *BaseA2LListener) EnterBlockDefinition(ctx *BlockDefinitionContext) {}

// ExitBlockDefinition is called when production blockDefinition is exited.
func (s *BaseA2LListener) ExitBlockDefinition(ctx *BlockDefinitionContext) {}

// EnterEnumTypeName is called when production enumTypeName is entered.
func (s *BaseA2LListener) EnterEnumTypeName(ctx *EnumTypeNameContext) {}

// ExitEnumTypeName is called when production enumTypeName is exited.
func (s *BaseA2LListener) ExitEnumTypeName(ctx *EnumTypeNameContext) {}

// EnterEnumeratorList is called when production enumeratorList is entered.
func (s *BaseA2LListener) EnterEnumeratorList(ctx *EnumeratorListContext) {}

// ExitEnumeratorList is called when production enumeratorList is exited.
func (s *BaseA2LListener) ExitEnumeratorList(ctx *EnumeratorListContext) {}

// EnterEnumerator is called when production enumerator is entered.
func (s *BaseA2LListener) EnterEnumerator(ctx *EnumeratorContext) {}

// ExitEnumerator is called when production enumerator is exited.
func (s *BaseA2LListener) ExitEnumerator(ctx *EnumeratorContext) {}

// EnterStructTypeName is called when production structTypeName is entered.
func (s *BaseA2LListener) EnterStructTypeName(ctx *StructTypeNameContext) {}

// ExitStructTypeName is called when production structTypeName is exited.
func (s *BaseA2LListener) ExitStructTypeName(ctx *StructTypeNameContext) {}

// EnterStructMember is called when production structMember is entered.
func (s *BaseA2LListener) EnterStructMember(ctx *StructMemberContext) {}

// ExitStructMember is called when production structMember is exited.
func (s *BaseA2LListener) ExitStructMember(ctx *StructMemberContext) {}

// EnterMember is called when production member is entered.
func (s *BaseA2LListener) EnterMember(ctx *MemberContext) {}

// ExitMember is called when production member is exited.
func (s *BaseA2LListener) ExitMember(ctx *MemberContext) {}

// EnterTaggedStructTypeName is called when production taggedStructTypeName is entered.
func (s *BaseA2LListener) EnterTaggedStructTypeName(ctx *TaggedStructTypeNameContext) {}

// ExitTaggedStructTypeName is called when production taggedStructTypeName is exited.
func (s *BaseA2LListener) ExitTaggedStructTypeName(ctx *TaggedStructTypeNameContext) {}

// EnterTaggedStructMember is called when production taggedStructMember is entered.
func (s *BaseA2LListener) EnterTaggedStructMember(ctx *TaggedStructMemberContext) {}

// ExitTaggedStructMember is called when production taggedStructMember is exited.
func (s *BaseA2LListener) ExitTaggedStructMember(ctx *TaggedStructMemberContext) {}

// EnterTaggedStructDefinition is called when production taggedStructDefinition is entered.
func (s *BaseA2LListener) EnterTaggedStructDefinition(ctx *TaggedStructDefinitionContext) {}

// ExitTaggedStructDefinition is called when production taggedStructDefinition is exited.
func (s *BaseA2LListener) ExitTaggedStructDefinition(ctx *TaggedStructDefinitionContext) {}

// EnterTaggedUnionTypeName is called when production taggedUnionTypeName is entered.
func (s *BaseA2LListener) EnterTaggedUnionTypeName(ctx *TaggedUnionTypeNameContext) {}

// ExitTaggedUnionTypeName is called when production taggedUnionTypeName is exited.
func (s *BaseA2LListener) ExitTaggedUnionTypeName(ctx *TaggedUnionTypeNameContext) {}

// EnterTaggedUnionMember is called when production taggedUnionMember is entered.
func (s *BaseA2LListener) EnterTaggedUnionMember(ctx *TaggedUnionMemberContext) {}

// ExitTaggedUnionMember is called when production taggedUnionMember is exited.
func (s *BaseA2LListener) ExitTaggedUnionMember(ctx *TaggedUnionMemberContext) {}

// EnterTagValue is called when production tagValue is entered.
func (s *BaseA2LListener) EnterTagValue(ctx *TagValueContext) {}

// ExitTagValue is called when production tagValue is exited.
func (s *BaseA2LListener) ExitTagValue(ctx *TagValueContext) {}

// EnterIfData is called when production ifData is entered.
func (s *BaseA2LListener) EnterIfData(ctx *IfDataContext) {}

// ExitIfData is called when production ifData is exited.
func (s *BaseA2LListener) ExitIfData(ctx *IfDataContext) {}

// EnterGenericParameter is called when production genericParameter is entered.
func (s *BaseA2LListener) EnterGenericParameter(ctx *GenericParameterContext) {}

// ExitGenericParameter is called when production genericParameter is exited.
func (s *BaseA2LListener) ExitGenericParameter(ctx *GenericParameterContext) {}

// EnterGenericNode is called when production genericNode is entered.
func (s *BaseA2LListener) EnterGenericNode(ctx *GenericNodeContext) {}

// ExitGenericNode is called when production genericNode is exited.
func (s *BaseA2LListener) ExitGenericNode(ctx *GenericNodeContext) {}
