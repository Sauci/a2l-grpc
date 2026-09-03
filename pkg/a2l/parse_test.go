package a2l

// Behaviour of the parsing entry points which is not tied to a single keyword: how a failure is
// reported, and how the include mechanism of ASAM MCD-2 MC 1.6.1 chapter 4 is handled.

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// The messages of the parser quote the offending part of the file, so they must never be used as
// a format string: a per cent sign in the parsed content would otherwise be taken for a verb and
// the message would reach the caller mangled, e.g. as "%!'(NOVERB)".
func TestParse_ErrorMessageQuotingTheContentIsNotFormatted(t *testing.T) {
	_, err := GetTreeFromString(moduleScope("%"))
	if !assert.Error(t, err, "a stray per cent sign should be rejected") {
		return
	}

	assert.NotContains(t, err.Error(), "%!", "the message should not be interpreted as a format string")
	assert.NotContains(t, err.Error(), "(MISSING)")
	assert.Contains(t, err.Error(), "%", "the message should still quote the offending character")
}

// ASAM MCD-2 MC 1.6.1 chapter 4 (ASAP2 1.51 chapter 5): "/include <filename>". Resolving the
// reference needs the file system the file was read from, which the parser does not have, so each
// occurrence is reported with the name of the file it refers to instead of being dropped.
func TestParse_Include(t *testing.T) {
	// Appendix B.2, MST_ABS.A2L
	t.Run("reported with the name of the file", func(t *testing.T) {
		_, err := GetTreeFromString(projectScope("/include engine_ecu.a2l"))
		if !assert.Error(t, err, "an unresolved include should be reported") {
			return
		}

		assert.Contains(t, err.Error(), "/include engine_ecu.a2l")
		assert.Contains(t, err.Error(), "is not resolved by this parser")
	})

	// Chapter 4: "The filename may be put between quotation marks. If the filename contains
	// spaces or path information the quotation marks are required." and "The relative path uses
	// backslashes without escape sequences", so the escape sequences of a string do not apply.
	t.Run("accepts every notation of the filename", func(t *testing.T) {
		for _, filename := range []string{
			`ENGINE_ECU.A2L`,
			`"C:\ENG_ECU.A2L"`,
			`"..\includes\ABS_ECU.A2L"`,
			`"SPEC_ECU.A2L"`,
			`"\\MyServer\VariableDescriptions\ESP_ECU.A2L"`,
			`"with space.a2l"`,
		} {
			_, err := GetTreeFromString(projectScope("/include " + filename))
			if !assert.Error(t, err, "for %s", filename) {
				continue
			}

			assert.Contains(t, err.Error(), filename)
			assert.Equal(t, 1, strings.Count(err.Error(), "\n")+1,
				"one include should produce exactly one message, for %s", filename)
		}
	})

	// The statement is a text replacement and "there is no restriction where to use include", so
	// it must not throw the parser off wherever it appears.
	t.Run("does not desynchronize the parser", func(t *testing.T) {
		_, err := GetTreeFromString(fileScope(
			"/include header.a2l\n" +
				"/begin PROJECT project \"\"\n" +
				"/include modules.a2l\n" +
				"/begin MODULE module \"\"\n/end MODULE\n" +
				"/end PROJECT"))
		if !assert.Error(t, err) {
			return
		}

		assert.Contains(t, err.Error(), "/include header.a2l")
		assert.Contains(t, err.Error(), "/include modules.a2l")
		assert.NotContains(t, err.Error(), "mismatched input",
			"the include should not be reported as a syntax error of its own")
	})

	// A comment behind the statement belongs to the file, not to the filename.
	t.Run("stops at a trailing comment", func(t *testing.T) {
		_, err := GetTreeFromString(projectScope("/include engine_ecu.a2l /* the engine ECU */"))
		if !assert.Error(t, err) {
			return
		}

		assert.Contains(t, err.Error(), "/include engine_ecu.a2l")
		assert.NotContains(t, err.Error(), "the engine ECU")
	})
}

// Chapter 3.2 declares int and uint as 2 byte integers, so the 32 bit node which carries them is
// already permissive. A wider literal used to be truncated silently, which both changed the value
// and produced a tree which could no longer be serialized back to A2L.
func TestParse_IntegerParameterOutOfRange(t *testing.T) {
	t.Run("reported instead of truncated", func(t *testing.T) {
		_, err := GetTreeFromString(modCommonScope("ALIGNMENT_BYTE 0xFFFFFFFF"))
		if !assert.Error(t, err, "a literal which does not fit should be reported") {
			return
		}

		assert.Contains(t, err.Error(), "0xFFFFFFFF")
		assert.Contains(t, err.Error(), "value out of range")
	})

	t.Run("reported with its position", func(t *testing.T) {
		_, err := GetTreeFromString("ASAP2_VERSION 1 0x100000000\n" +
			projectScope("/begin MODULE module \"\"\n/end MODULE"))
		if !assert.Error(t, err) {
			return
		}

		assert.Contains(t, err.Error(), "1:16", "the message should name the line and the column")
	})

	// An address is a ulong, which the 64 bit node holds without truncating.
	t.Run("an unsigned 32 bit address is accepted", func(t *testing.T) {
		module, ok := parseModule(t, "/begin CHARACTERISTIC characteristic \"\" VALUE 0xFFFFFFFF "+
			"record_layout 0 compu_method 0 0\n/end CHARACTERISTIC")
		if !ok || !assert.Len(t, module.CHARACTERISTIC, 1) {
			return
		}

		assert.Equal(t, int64(0xFFFFFFFF), module.CHARACTERISTIC[0].Address.Value)
	})
}

// An identifier whose partial strings are separated by points is decomposed as chapter 3.2
// describes it. The chapter lists two limitations: "the first character must be a letter or an
// underscore, brackets must occur in pairs at the end of a partial string". The first one is about
// the first character of the identifier, not of every partial string; where the chapter means a
// partial string it says so, as the second limitation does.
func TestParse_PartialIdentifier(t *testing.T) {
	// https://github.com/Sauci/pya2l/issues/17: a real characteristic name whose third partial
	// string is a number.
	t.Run("a partial identifier may begin with a digit", func(t *testing.T) {
		module, ok := parseModule(t, "/begin CHARACTERISTIC SFB_R_FFO_DE.Properties.1.Qly \"\" "+
			"VALUE 36453 record_layout 0 compu_method 0 255\n"+
			"SYMBOL_LINK \"SFB_R_FFO_DE.Properties.1.Qly\" 0\n/end CHARACTERISTIC")
		if !ok || !assert.Len(t, module.CHARACTERISTIC, 1) {
			return
		}

		assert.Equal(t, "SFB_R_FFO_DE.Properties.1.Qly", module.CHARACTERISTIC[0].Name.Value)
	})

	t.Run("a partial identifier must not be empty", func(t *testing.T) {
		for _, identifier := range []string{"trailing.", "empty..partial"} {
			_, err := GetTreeFromString(moduleScope("/begin CHARACTERISTIC " + identifier +
				" \"\" VALUE 0x0 record_layout 0 compu_method 0 0\n/end CHARACTERISTIC"))
			if !assert.Error(t, err, "identifier %s", identifier) {
				continue
			}

			assert.Contains(t, err.Error(), "a partial identifier must not be empty")
		}
	})

	// Chapter 3.2: the brackets "must contain a number or an alphabetic string", the latter being
	// "a symbolic string which is defined as an enumerator of an ENUM definition of the C program".
	t.Run("a symbolic array index is a single enumerator", func(t *testing.T) {
		parse(t, moduleScope("/begin CHARACTERISTIC array[INDEX] \"\" VALUE 0x0 record_layout 0 "+
			"compu_method 0 0\n/end CHARACTERISTIC"))

		_, err := GetTreeFromString(moduleScope("/begin CHARACTERISTIC array[symbolic.index] \"\" " +
			"VALUE 0x0 record_layout 0 compu_method 0 0\n/end CHARACTERISTIC"))
		if assert.Error(t, err) {
			assert.Contains(t, err.Error(), "a symbolic index is a single enumerator")
		}
	})
}
