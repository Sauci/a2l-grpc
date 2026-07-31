package a2l

// Version gating: the grammar accepts a superset of ASAP2 1.51. When the parsed file declares its
// version with ASAP2_VERSION, every construct which requires a newer version of the standard than
// the declared one is reported as a warning; with ParseOptions.EnforceVersionCheck the file is
// rejected instead. A file which does not declare a version is not gated.

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrammar_VersionCheck(t *testing.T) {
	type gateType struct {
		construct string
		min       string
		accept    string
		body      string
	}

	gates := []gateType{
		{construct: "STEP_SIZE", min: "1.60", accept: "1 60",
			body: characteristicScope("STEP_SIZE 0.5")},
		{construct: "PHYS_UNIT", min: "1.60", accept: "1 60",
			body: characteristicScope("PHYS_UNIT \"km/h\"")},
		{construct: "DISCRETE", min: "1.60", accept: "1 60",
			body: characteristicScope("DISCRETE")},
		{construct: "SYMBOL_LINK", min: "1.60", accept: "1 60",
			body: characteristicScope("SYMBOL_LINK \"symbol\" 4")},
		{construct: "LAYOUT", min: "1.60", accept: "1 60",
			body: measurementScope("LAYOUT ROW_DIR")},
		{construct: "COEFFS_LINEAR", min: "1.60", accept: "1 60",
			body: compuMethodScope("COEFFS_LINEAR 1 0")},
		{construct: "STATUS_STRING_REF", min: "1.60", accept: "1 60",
			body: compuMethodScope("STATUS_STRING_REF compu_vtab")},
		{construct: "DEFAULT_VALUE_NUMERIC", min: "1.60", accept: "1 60",
			body: moduleScope("/begin COMPU_TAB compu_tab \"\" TAB_INTP 1 0 1 " +
				"DEFAULT_VALUE_NUMERIC 1.5\n/end COMPU_TAB")},
		{construct: "ALIGNMENT_INT64", min: "1.60", accept: "1 60",
			body: modCommonScope("ALIGNMENT_INT64 8")},
		{construct: "STATIC_RECORD_LAYOUT", min: "1.60", accept: "1 60",
			body: recordLayoutScope("STATIC_RECORD_LAYOUT")},
		{construct: "CALIBRATION_HANDLE_TEXT", min: "1.60", accept: "1 60",
			body: modParScope("/begin CALIBRATION_METHOD \"InCircuit\" 2\n" +
				"/begin CALIBRATION_HANDLE 0x10\nCALIBRATION_HANDLE_TEXT \"Torque\"\n" +
				"/end CALIBRATION_HANDLE\n/end CALIBRATION_METHOD")},
		{construct: "MONOTONY in AXIS_PTS", min: "1.60", accept: "1 60",
			body: axisPtsScope("MONOTONY MON_INCREASE")},
		{construct: "IF_DATA in FUNCTION", min: "1.60", accept: "1 60",
			body: functionScope("/begin IF_DATA XCP\n/end IF_DATA")},
		{construct: "IF_DATA in GROUP", min: "1.60", accept: "1 60",
			body: groupScope("/begin IF_DATA XCP\n/end IF_DATA")},
		{construct: "CHARACTERISTIC type CUBE_4", min: "1.60", accept: "1 60",
			body: moduleScope("/begin CHARACTERISTIC characteristic \"\" CUBE_4 0x0 record_layout 0 " +
				"compu_method 0 0\n/end CHARACTERISTIC")},
		{construct: "CHARACTERISTIC type CUBE_5", min: "1.60", accept: "1 60",
			body: moduleScope("/begin CHARACTERISTIC characteristic \"\" CUBE_5 0x0 record_layout 0 " +
				"compu_method 0 0\n/end CHARACTERISTIC")},
		{construct: "MONOTONY MONOTONOUS", min: "1.60", accept: "1 60",
			body: axisDescrScope("MONOTONY MONOTONOUS")},
		{construct: "MONOTONY STRICT_MON", min: "1.60", accept: "1 60",
			body: axisDescrScope("MONOTONY STRICT_MON")},
		{construct: "MONOTONY NOT_MON", min: "1.60", accept: "1 60",
			body: axisDescrScope("MONOTONY NOT_MON")},
		{construct: "conversion type IDENTICAL", min: "1.60", accept: "1 60",
			body: moduleScope("/begin COMPU_METHOD compu_method \"\" IDENTICAL \"%4.2\" \"unit\"\n" +
				"/end COMPU_METHOD")},
		{construct: "conversion type LINEAR", min: "1.60", accept: "1 60",
			body: moduleScope("/begin COMPU_METHOD compu_method \"\" LINEAR \"%4.2\" \"unit\"\n" +
				"/end COMPU_METHOD")},
		{construct: "data type A_UINT64", min: "1.60", accept: "1 60",
			body: moduleScope("/begin MEASUREMENT measurement \"\" A_UINT64 compu_method 0 0 0 0\n" +
				"/end MEASUREMENT")},
		{construct: "data type A_INT64", min: "1.60", accept: "1 60",
			body: moduleScope("/begin MEASUREMENT measurement \"\" A_INT64 compu_method 0 0 0 0\n" +
				"/end MEASUREMENT")},
	}

	for _, keyword := range []string{"NO_AXIS_PTS", "NO_RESCALE", "SRC_ADDR", "RIP_ADDR", "SHIFT_OP", "OFFSET", "DIST_OP"} {
		for _, dimension := range []string{"_4", "_5"} {
			gates = append(gates, gateType{construct: keyword + dimension, min: "1.60", accept: "1 60",
				body: recordLayoutScope(keyword + dimension + " 1 UBYTE")})
		}
	}

	for _, dimension := range []string{"_4", "_5"} {
		gates = append(gates,
			gateType{construct: "AXIS_PTS" + dimension, min: "1.60", accept: "1 60",
				body: recordLayoutScope("AXIS_PTS" + dimension + " 1 UBYTE INDEX_INCR DIRECT")},
			gateType{construct: "AXIS_RESCALE" + dimension, min: "1.60", accept: "1 60",
				body: recordLayoutScope("AXIS_RESCALE" + dimension + " 1 UBYTE 4 INDEX_INCR DIRECT")},
			gateType{construct: "FIX_NO_AXIS_PTS" + dimension, min: "1.60", accept: "1 60",
				body: recordLayoutScope("FIX_NO_AXIS_PTS" + dimension + " 17")})
	}

	for _, gate := range gates {
		t.Run(gate.construct+"/warned when the file declares 1.51", func(t *testing.T) {
			_, warnings, err := GetTreeFromStringWithOptions("ASAP2_VERSION 1 51\n"+gate.body, ParseOptions{})
			if !assert.NoError(t, err, "without enforcement the file should still parse") {
				return
			}

			if assert.NotEmpty(t, warnings, "the construct should have been warned about") {
				assert.Contains(t, warnings[0].String(), gate.construct)
				assert.Contains(t, warnings[0].String(), "requires ASAP2 version "+gate.min)
			}
		})

		t.Run(gate.construct+"/rejected when the check is enforced", func(t *testing.T) {
			_, _, err := GetTreeFromStringWithOptions("ASAP2_VERSION 1 51\n"+gate.body,
				ParseOptions{EnforceVersionCheck: true})
			if assert.Error(t, err, "with enforcement the parser should have rejected the construct") {
				assert.Contains(t, err.Error(), gate.construct)
				assert.Contains(t, err.Error(), "requires ASAP2 version "+gate.min)
			}
		})

		t.Run(gate.construct+"/no warning when the file declares "+gate.min, func(t *testing.T) {
			_, warnings, err := GetTreeFromStringWithOptions("ASAP2_VERSION "+gate.accept+"\n"+gate.body,
				ParseOptions{EnforceVersionCheck: true})
			assert.NoError(t, err)
			assert.Empty(t, warnings)
		})
	}

	// backward compatibility: a file declaring an older version than its content conforms to is
	// still parsed by default, GetTreeFromString does not fail on it
	t.Run("a lying 1.51 file still parses by default", func(t *testing.T) {
		tree, err := GetTreeFromString("ASAP2_VERSION 1 51\n" + characteristicScope("STEP_SIZE 0.5"))
		if !assert.NoError(t, err) {
			return
		}

		characteristic, ok := characteristicOf(t, tree)
		if ok {
			equalNode(t, &StepSizeType{StepSize: floatVal("0.5")}, characteristic.STEP_SIZE)
		}
	})

	t.Run("no declared version is not gated", func(t *testing.T) {
		_, warnings, err := GetTreeFromStringWithOptions(characteristicScope("STEP_SIZE 0.5"),
			ParseOptions{EnforceVersionCheck: true})
		assert.NoError(t, err)
		assert.Empty(t, warnings)
	})

	t.Run("a 1.60 file does not gate 1.60 keywords used together", func(t *testing.T) {
		_, warnings, err := GetTreeFromStringWithOptions("ASAP2_VERSION 1 60\n"+
			characteristicScope("STEP_SIZE 0.5\nPHYS_UNIT \"km/h\"\nDISCRETE"),
			ParseOptions{EnforceVersionCheck: true})
		assert.NoError(t, err)
		assert.Empty(t, warnings)
	})

	t.Run("the warning names the position and the declared version", func(t *testing.T) {
		_, warnings, err := GetTreeFromStringWithOptions("ASAP2_VERSION 1 51\n"+
			characteristicScope("STEP_SIZE 0.5"), ParseOptions{})
		assert.NoError(t, err)

		if assert.NotEmpty(t, warnings) {
			assert.Contains(t, warnings[0].String(),
				"STEP_SIZE requires ASAP2 version 1.60, but the file declares ASAP2_VERSION 1 51")
		}
	})

	// VAR_NAMING ALPHA is reserved for a future extension of the standard (ASAM MCD-2 MC 1.6.1,
	// chapter 3.5.134): it is valid in none of the supported versions, so the warning does not
	// depend on the declared version.
	t.Run("VAR_NAMING ALPHA is warned about independently of the declared version", func(t *testing.T) {
		for _, prefix := range []string{"", "ASAP2_VERSION 1 51\n", "ASAP2_VERSION 1 61\n"} {
			_, warnings, err := GetTreeFromStringWithOptions(prefix+variantCodingScope("VAR_NAMING ALPHA"),
				ParseOptions{})
			assert.NoError(t, err)

			if assert.NotEmpty(t, warnings, "prefix %q", prefix) {
				assert.Contains(t, warnings[0].String(),
					"VAR_NAMING ALPHA is reserved for a future extension of the standard")
			}
		}

		_, _, err := GetTreeFromStringWithOptions(variantCodingScope("VAR_NAMING ALPHA"),
			ParseOptions{EnforceVersionCheck: true})
		assert.Error(t, err)
	})

	t.Run("one warning per offending construct", func(t *testing.T) {
		_, warnings, err := GetTreeFromStringWithOptions("ASAP2_VERSION 1 51\n"+
			characteristicScope("STEP_SIZE 0.5\nDISCRETE\nPHYS_UNIT \"km/h\""), ParseOptions{})
		assert.NoError(t, err)
		assert.Len(t, warnings, 3)
	})
}
