package a2l

// Version gating of the parsed content.
//
// The grammar accepts a superset of ASAP2 1.51: a number of keywords and enumeration values were
// introduced by later versions of the standard (ASAM MCD-2MC 1.60 and 1.70). When the parsed file
// declares its version with ASAP2_VERSION, the constructs which require a newer version than the
// declared one are reported as warnings, or as errors when the check is enforced with
// ParseOptions.EnforceVersionCheck. A file which does not declare a version is not gated, since
// there is nothing to check against.

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/sauci/a2l-grpc/pkg/a2l/parser"
)

type asap2Version struct {
	versionNo, upgradeNo int32
}

func (v asap2Version) atLeast(other asap2Version) bool {
	if v.versionNo != other.versionNo {
		return v.versionNo > other.versionNo
	}

	return v.upgradeNo >= other.upgradeNo
}

func (v asap2Version) String() string {
	return fmt.Sprintf("%d.%d", v.versionNo, v.upgradeNo)
}

var asap2Version160 = asap2Version{versionNo: 1, upgradeNo: 60}

type versionCheckListener struct {
	*parser.BaseA2LListener
	version  *asap2Version
	warnings []SyntaxError
}

func newVersionCheckListener() *versionCheckListener {
	return &versionCheckListener{}
}

func (l *versionCheckListener) warn(message string, at antlr.Token) {
	l.warnings = append(l.warnings, SyntaxError{
		Line:   at.GetLine(),
		Column: at.GetColumn(),
		Msg:    message,
	})
}

func (l *versionCheckListener) require(min asap2Version, construct string, at antlr.Token) {
	if l.version == nil || l.version.atLeast(min) {
		return
	}

	l.warn(fmt.Sprintf("%s requires ASAP2 version %s, but the file declares ASAP2_VERSION %d %d",
		construct, min, l.version.versionNo, l.version.upgradeNo), at)
}

func (l *versionCheckListener) requireValue(min asap2Version, keyword string, value antlr.Token, gated ...string) {
	if value == nil {
		return
	}

	for _, g := range gated {
		if value.GetText() == g {
			l.require(min, fmt.Sprintf("%s %s", keyword, value.GetText()), value)
			return
		}
	}
}

func (l *versionCheckListener) EnterAsap2Version(ctx *parser.Asap2VersionContext) {
	version := &asap2Version{}

	if v := a2lIntToIntType(ctx.GetVersionNo()); v != nil {
		version.versionNo = v.Value
	}

	if v := a2lIntToIntType(ctx.GetUpgradeNo()); v != nil {
		version.upgradeNo = v.Value
	}

	l.version = version
}

func (l *versionCheckListener) EnterStepSize(ctx *parser.StepSizeContext) {
	l.require(asap2Version160, "STEP_SIZE", ctx.GetStart())
}

func (l *versionCheckListener) EnterPhysUnit(ctx *parser.PhysUnitContext) {
	l.require(asap2Version160, "PHYS_UNIT", ctx.GetStart())
}

func (l *versionCheckListener) EnterDiscrete(ctx *parser.DiscreteContext) {
	l.require(asap2Version160, "DISCRETE", ctx.GetStart())
}

func (l *versionCheckListener) EnterSymbolLink(ctx *parser.SymbolLinkContext) {
	l.require(asap2Version160, "SYMBOL_LINK", ctx.GetStart())
}

func (l *versionCheckListener) EnterLayout(ctx *parser.LayoutContext) {
	l.require(asap2Version160, "LAYOUT", ctx.GetStart())
}

func (l *versionCheckListener) EnterCoeffsLinear(ctx *parser.CoeffsLinearContext) {
	l.require(asap2Version160, "COEFFS_LINEAR", ctx.GetStart())
}

func (l *versionCheckListener) EnterStatusStringRef(ctx *parser.StatusStringRefContext) {
	l.require(asap2Version160, "STATUS_STRING_REF", ctx.GetStart())
}

func (l *versionCheckListener) EnterDefaultValueNumeric(ctx *parser.DefaultValueNumericContext) {
	l.require(asap2Version160, "DEFAULT_VALUE_NUMERIC", ctx.GetStart())
}

func (l *versionCheckListener) EnterAlignmentInt64(ctx *parser.AlignmentInt64Context) {
	l.require(asap2Version160, "ALIGNMENT_INT64", ctx.GetStart())
}

func (l *versionCheckListener) EnterStaticRecordLayout(ctx *parser.StaticRecordLayoutContext) {
	l.require(asap2Version160, "STATIC_RECORD_LAYOUT", ctx.GetStart())
}

func (l *versionCheckListener) EnterAxisPts4(ctx *parser.AxisPts4Context) {
	l.require(asap2Version160, "AXIS_PTS_4", ctx.GetStart())
}

func (l *versionCheckListener) EnterAxisPts5(ctx *parser.AxisPts5Context) {
	l.require(asap2Version160, "AXIS_PTS_5", ctx.GetStart())
}

func (l *versionCheckListener) EnterAxisRescale4(ctx *parser.AxisRescale4Context) {
	l.require(asap2Version160, "AXIS_RESCALE_4", ctx.GetStart())
}

func (l *versionCheckListener) EnterAxisRescale5(ctx *parser.AxisRescale5Context) {
	l.require(asap2Version160, "AXIS_RESCALE_5", ctx.GetStart())
}

func (l *versionCheckListener) EnterNoAxisPts4(ctx *parser.NoAxisPts4Context) {
	l.require(asap2Version160, "NO_AXIS_PTS_4", ctx.GetStart())
}

func (l *versionCheckListener) EnterNoAxisPts5(ctx *parser.NoAxisPts5Context) {
	l.require(asap2Version160, "NO_AXIS_PTS_5", ctx.GetStart())
}

func (l *versionCheckListener) EnterNoRescale4(ctx *parser.NoRescale4Context) {
	l.require(asap2Version160, "NO_RESCALE_4", ctx.GetStart())
}

func (l *versionCheckListener) EnterNoRescale5(ctx *parser.NoRescale5Context) {
	l.require(asap2Version160, "NO_RESCALE_5", ctx.GetStart())
}

func (l *versionCheckListener) EnterFixNoAxisPts4(ctx *parser.FixNoAxisPts4Context) {
	l.require(asap2Version160, "FIX_NO_AXIS_PTS_4", ctx.GetStart())
}

func (l *versionCheckListener) EnterFixNoAxisPts5(ctx *parser.FixNoAxisPts5Context) {
	l.require(asap2Version160, "FIX_NO_AXIS_PTS_5", ctx.GetStart())
}

func (l *versionCheckListener) EnterSrcAddr4(ctx *parser.SrcAddr4Context) {
	l.require(asap2Version160, "SRC_ADDR_4", ctx.GetStart())
}

func (l *versionCheckListener) EnterSrcAddr5(ctx *parser.SrcAddr5Context) {
	l.require(asap2Version160, "SRC_ADDR_5", ctx.GetStart())
}

func (l *versionCheckListener) EnterRipAddr4(ctx *parser.RipAddr4Context) {
	l.require(asap2Version160, "RIP_ADDR_4", ctx.GetStart())
}

func (l *versionCheckListener) EnterRipAddr5(ctx *parser.RipAddr5Context) {
	l.require(asap2Version160, "RIP_ADDR_5", ctx.GetStart())
}

func (l *versionCheckListener) EnterShiftOp4(ctx *parser.ShiftOp4Context) {
	l.require(asap2Version160, "SHIFT_OP_4", ctx.GetStart())
}

func (l *versionCheckListener) EnterShiftOp5(ctx *parser.ShiftOp5Context) {
	l.require(asap2Version160, "SHIFT_OP_5", ctx.GetStart())
}

func (l *versionCheckListener) EnterOffset4(ctx *parser.Offset4Context) {
	l.require(asap2Version160, "OFFSET_4", ctx.GetStart())
}

func (l *versionCheckListener) EnterOffset5(ctx *parser.Offset5Context) {
	l.require(asap2Version160, "OFFSET_5", ctx.GetStart())
}

func (l *versionCheckListener) EnterDistOp4(ctx *parser.DistOp4Context) {
	l.require(asap2Version160, "DIST_OP_4", ctx.GetStart())
}

func (l *versionCheckListener) EnterDistOp5(ctx *parser.DistOp5Context) {
	l.require(asap2Version160, "DIST_OP_5", ctx.GetStart())
}

func (l *versionCheckListener) EnterCalibrationHandleText(ctx *parser.CalibrationHandleTextContext) {
	l.require(asap2Version160, "CALIBRATION_HANDLE_TEXT", ctx.GetStart())
}

// ASAP2 1.51 does not declare IF_DATA for FUNCTION and GROUP, ASAM MCD-2 MC 1.6.0 added it.
func (l *versionCheckListener) EnterIfData(ctx *parser.IfDataContext) {
	switch ctx.GetParent().(type) {
	case *parser.FunctionContext:
		l.require(asap2Version160, "IF_DATA in FUNCTION", ctx.GetStart())
	case *parser.GroupContext:
		l.require(asap2Version160, "IF_DATA in GROUP", ctx.GetStart())
	}
}

func (l *versionCheckListener) EnterCharacteristic(ctx *parser.CharacteristicContext) {
	l.requireValue(asap2Version160, "CHARACTERISTIC type", ctx.GetType_(), "CUBE_4", "CUBE_5")
}

func (l *versionCheckListener) EnterMonotony(ctx *parser.MonotonyContext) {
	l.requireValue(asap2Version160, "MONOTONY", ctx.GetMonotony_(), "MONOTONOUS", "STRICT_MON", "NOT_MON")

	// ASAP2 1.51 declares MONOTONY for AXIS_DESCR only, ASAM MCD-2 MC 1.6.0 added it to AXIS_PTS
	if _, ok := ctx.GetParent().(*parser.AxisPtsContext); ok {
		l.require(asap2Version160, "MONOTONY in AXIS_PTS", ctx.GetStart())
	}
}

func (l *versionCheckListener) EnterCompuMethod(ctx *parser.CompuMethodContext) {
	l.requireValue(asap2Version160, "conversion type", ctx.GetConversionType(), "IDENTICAL", "LINEAR")
}

func (l *versionCheckListener) EnterDataType(ctx *parser.DataTypeContext) {
	l.requireValue(asap2Version160, "data type", ctx.GetV(), "A_UINT64", "A_INT64")
}

// EnterVarNaming reports the value ALPHA, which is reserved for a future extension of the
// standard (ASAM MCD-2 MC 1.6.1, chapter 3.5.134) and therefore not valid in any of the
// supported versions. Unlike the version gates above, this warning does not depend on the
// declared version.
func (l *versionCheckListener) EnterVarNaming(ctx *parser.VarNamingContext) {
	if tag := ctx.GetTag(); tag != nil && tag.GetText() == "ALPHA" {
		l.warn("VAR_NAMING ALPHA is reserved for a future extension of the standard", tag)
	}
}

