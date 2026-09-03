package a2l

// Version gating: the grammar accepts a superset of ASAP2 1.51. When the parsed file declares its
// version with ASAP2_VERSION, every construct which requires a newer version of the standard than
// the declared one is reported as a warning; with ParseOptions.EnforceVersionCheck the file is
// rejected instead. A file which does not declare a version is not gated; when the check is
// enforced, the missing ASAP2_VERSION is reported instead.

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
			ParseOptions{})
		assert.NoError(t, err)
		assert.Empty(t, warnings)
	})

	// ASAP2_VERSION is mandatory since ASAM MCD-2 MC 1.6.1 (chapters 1.4.4 and 3.5.16). A file
	// which does not declare it cannot be gated at all, so a caller which asked for the check to
	// be enforced is told about it instead of silently getting no verification at all.
	t.Run("missing ASAP2_VERSION is reported when the check is enforced", func(t *testing.T) {
		_, _, err := GetTreeFromStringWithOptions(characteristicScope("STEP_SIZE 0.5"),
			ParseOptions{EnforceVersionCheck: true})
		if assert.Error(t, err, "the enforced check should report that it cannot verify anything") {
			assert.Contains(t, err.Error(), "ASAP2_VERSION is missing")
		}
	})

	t.Run("a declared version is not reported as missing", func(t *testing.T) {
		_, warnings, err := GetTreeFromStringWithOptions("ASAP2_VERSION 1 61\n"+
			characteristicScope("STEP_SIZE 0.5"), ParseOptions{EnforceVersionCheck: true})
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

// Chapter 3.5.1 fixes the notation of the keyword prototypes: "Optional keywords are shown with
// help of square brackets, which include an arrow followed by the keyword. If the keyword can be
// used multiple times this is shown with help of asterisk after the closing bracket, e.g.
// [-> keyword]*." A keyword written without the asterisk may appear at most once. The grammar
// cannot express that, every optional element sits in a "( a | b | c )*" loop, so a repeat parses
// and the tree builder keeps only the last occurrence: without this check the earlier ones are
// dropped without a word.
func TestGrammar_DuplicateSingleOccurrenceKeyword(t *testing.T) {
	t.Run("a repeated single occurrence keyword is reported", func(t *testing.T) {
		_, warnings, err := GetTreeFromStringWithOptions(
			characteristicScope("BYTE_ORDER MSB_LAST\nBYTE_ORDER MSB_FIRST"), ParseOptions{})
		if !assert.NoError(t, err, "without enforcement the file should still parse") {
			return
		}

		if assert.NotEmpty(t, warnings, "the repeated keyword should have been reported") {
			assert.Contains(t, warnings[0].String(), "BYTE_ORDER")
			assert.Contains(t, warnings[0].String(), "at most once")
		}
	})

	// The elements the specification marks with an asterisk must stay silent.
	t.Run("a repeatable keyword is not reported", func(t *testing.T) {
		_, warnings, err := GetTreeFromStringWithOptions(characteristicScope(
			"/begin ANNOTATION\n/end ANNOTATION\n/begin ANNOTATION\n/end ANNOTATION"), ParseOptions{})
		assert.NoError(t, err)
		assert.Empty(t, warnings)
	})

	t.Run("a single occurrence keyword used once is not reported", func(t *testing.T) {
		_, warnings, err := GetTreeFromStringWithOptions(
			characteristicScope("BYTE_ORDER MSB_LAST"), ParseOptions{})
		assert.NoError(t, err)
		assert.Empty(t, warnings)
	})

	t.Run("rejected when the check is enforced", func(t *testing.T) {
		_, _, err := GetTreeFromStringWithOptions(
			"ASAP2_VERSION 1 61\n"+characteristicScope("BYTE_ORDER MSB_LAST\nBYTE_ORDER MSB_FIRST"),
			ParseOptions{EnforceVersionCheck: true})
		if assert.Error(t, err) {
			assert.Contains(t, err.Error(), "BYTE_ORDER")
		}
	})
}

// ASAP2 1.51 (chapter 6.3.26) declares CALIBRATION_HANDLE as "{-> CALIBRATION_HANDLE}*", ASAM
// MCD-2 MC 1.6.1 (chapter 3.5.28) reduced it to "[-> CALIBRATION_HANDLE]". The grammar accepts
// several of them, so that a 1.51 file is not rejected, which means the check of the cardinality
// depends on the declared version and cannot be read from the labels of the grammar.
func TestGrammar_VersionCheck_ReducedCardinality(t *testing.T) {
	body := modParScope("/begin CALIBRATION_METHOD \"InCircuit\" 2\n" +
		"/begin CALIBRATION_HANDLE 0x10\n/end CALIBRATION_HANDLE\n" +
		"/begin CALIBRATION_HANDLE 0x20\n/end CALIBRATION_HANDLE\n" +
		"/end CALIBRATION_METHOD")

	t.Run("warned when the file declares 1.61", func(t *testing.T) {
		_, warnings, err := GetTreeFromStringWithOptions("ASAP2_VERSION 1 61\n"+body, ParseOptions{})
		if !assert.NoError(t, err, "without enforcement the file should still parse") {
			return
		}

		if assert.Len(t, warnings, 1, "only the occurrences beyond the first one are reported") {
			assert.Contains(t, warnings[0].String(), "CALIBRATION_HANDLE")
			assert.Contains(t, warnings[0].String(), "at most once since ASAP2 version 1.61")
		}
	})

	t.Run("still accepted when the file declares 1.51", func(t *testing.T) {
		_, warnings, err := GetTreeFromStringWithOptions("ASAP2_VERSION 1 51\n"+body,
			ParseOptions{EnforceVersionCheck: true})
		assert.NoError(t, err)
		assert.Empty(t, warnings)
	})

	t.Run("a single occurrence is never reported", func(t *testing.T) {
		_, warnings, err := GetTreeFromStringWithOptions("ASAP2_VERSION 1 61\n"+
			modParScope("/begin CALIBRATION_METHOD \"InCircuit\" 2\n"+
				"/begin CALIBRATION_HANDLE 0x10\n/end CALIBRATION_HANDLE\n"+
				"/end CALIBRATION_METHOD"),
			ParseOptions{EnforceVersionCheck: true})
		assert.NoError(t, err)
		assert.Empty(t, warnings)
	})
}

// ASAM MCD-2 MC 1.6.1 chapter 1.4.4: "The current version ASAM MCD-2 MC V 1.6.1 defined in this
// document does not support the following not usable keywords anymore: - S_REC_LAYOUT -
// NO_RESCALE_Y / _Z / _4 / _5 (reduced to NO_RESCALE_X) - AXIS_RESCALE_Y / _Z / _4 / _5 (reduced
// to AXIS_RESCALE_X)". They stay accepted, because the grammar covers ASAP2 1.51 as well, and are
// reported when the file declares the version which withdrew them.
func TestGrammar_VersionCheck_WithdrawnKeywords(t *testing.T) {
	type removalType struct {
		construct string
		body      string
	}

	for _, removal := range []removalType{
		{construct: "S_REC_LAYOUT", body: modCommonScope("S_REC_LAYOUT record_layout")},
		{construct: "AXIS_RESCALE_Y", body: recordLayoutScope("AXIS_RESCALE_Y 1 UWORD 4 INDEX_INCR DIRECT")},
		{construct: "AXIS_RESCALE_Z", body: recordLayoutScope("AXIS_RESCALE_Z 1 UWORD 4 INDEX_INCR DIRECT")},
		{construct: "AXIS_RESCALE_4", body: recordLayoutScope("AXIS_RESCALE_4 1 UWORD 4 INDEX_INCR DIRECT")},
		{construct: "AXIS_RESCALE_5", body: recordLayoutScope("AXIS_RESCALE_5 1 UWORD 4 INDEX_INCR DIRECT")},
		{construct: "NO_RESCALE_Y", body: recordLayoutScope("NO_RESCALE_Y 1 UWORD")},
		{construct: "NO_RESCALE_Z", body: recordLayoutScope("NO_RESCALE_Z 1 UWORD")},
		{construct: "NO_RESCALE_4", body: recordLayoutScope("NO_RESCALE_4 1 UWORD")},
		{construct: "NO_RESCALE_5", body: recordLayoutScope("NO_RESCALE_5 1 UWORD")},
	} {
		t.Run(removal.construct+"/warned when the file declares 1.61", func(t *testing.T) {
			_, warnings, err := GetTreeFromStringWithOptions("ASAP2_VERSION 1 61\n"+removal.body,
				ParseOptions{})
			if !assert.NoError(t, err, "without enforcement the file should still parse") {
				return
			}

			if assert.NotEmpty(t, warnings, "the withdrawn keyword should have been warned about") {
				assert.Contains(t, warnings[0].String(), removal.construct)
				assert.Contains(t, warnings[0].String(), "was removed in ASAP2 version 1.61")
			}
		})

		t.Run(removal.construct+"/rejected when the check is enforced", func(t *testing.T) {
			_, _, err := GetTreeFromStringWithOptions("ASAP2_VERSION 1 61\n"+removal.body,
				ParseOptions{EnforceVersionCheck: true})
			if assert.Error(t, err) {
				assert.Contains(t, err.Error(), removal.construct)
			}
		})

		t.Run(removal.construct+"/still accepted when the file declares 1.60", func(t *testing.T) {
			_, warnings, err := GetTreeFromStringWithOptions("ASAP2_VERSION 1 60\n"+removal.body,
				ParseOptions{EnforceVersionCheck: true})
			assert.NoError(t, err)
			assert.Empty(t, warnings)
		})
	}

	// The X variants of both families survived, they must never be reported.
	t.Run("the surviving X variants are not reported", func(t *testing.T) {
		_, warnings, err := GetTreeFromStringWithOptions("ASAP2_VERSION 1 61\n"+
			recordLayoutScope("AXIS_RESCALE_X 1 UWORD 4 INDEX_INCR DIRECT\nNO_RESCALE_X 2 UWORD"),
			ParseOptions{EnforceVersionCheck: true})
		assert.NoError(t, err)
		assert.Empty(t, warnings)
	})
}
