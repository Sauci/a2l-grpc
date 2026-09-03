package a2l

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrammar_Terminal_IDENT(t *testing.T) {
	for _, identifier := range []string{
		"identifier",
		"_leading_underscore",
		"with_digits_0123",
		"partial.identifier",
		"array[0]",
		"array[SYMBOLIC_INDEX]",
		"instance.element[2].member",
		// Chapter 3.2 requires a letter or an underscore for the first character of the
		// identifier; a partial string may begin with a digit, as the chapter names the partial
		// string explicitly where it means one ("brackets must occur in pairs at the end of a
		// partial string"). Such identifiers occur in the wild, e.g. "SFB_R_FFO_DE.Properties.1".
		"digit.1.partial",
		"digit.0partial",
		"SFB_R_FFO_DE.Properties.1.Qly",
		// ALPHA is named by chapter 3.5.134 as reserved for a future extension, but it is not an
		// enum value of any published version and is absent from the index of keywords and enum
		// values which chapter 3.2 declares to be the list an identifier must not match. Chapter
		// 3.5.29 uses it as an identifier in its own example:
		// "/begin DEPENDENT_CHARACTERISTIC "sin(X1)" ALPHA /end DEPENDENT_CHARACTERISTIC".
		"ALPHA",
	} {
		t.Run("accept/"+identifier, func(t *testing.T) {
			module, ok := parseModule(t,
				"/begin CHARACTERISTIC "+identifier+" \"\" VALUE 0x0 record_layout 0 compu_method 0 0\n"+
					"/end CHARACTERISTIC")
			if !ok || !assert.Len(t, module.CHARACTERISTIC, 1) {
				return
			}

			equalNode(t, identVal(identifier), module.CHARACTERISTIC[0].Name)
		})
	}

	t.Run("case sensitivity", func(t *testing.T) {
		module, ok := parseModule(t,
			"/begin CHARACTERISTIC x \"\" VALUE 0x0 record_layout 0 compu_method 0 0 /end CHARACTERISTIC\n"+
				"/begin CHARACTERISTIC X \"\" VALUE 0x0 record_layout 0 compu_method 0 0 /end CHARACTERISTIC")
		if !ok || !assert.Len(t, module.CHARACTERISTIC, 2) {
			return
		}

		assert.Equal(t, "x", module.CHARACTERISTIC[0].Name.Value)
		assert.Equal(t, "X", module.CHARACTERISTIC[1].Name.Value)
	})

	// Chapter 3.2 describes an identifier as a "hierarchical concatenation of partial strings
	// separated by points", so a partial string which is not there is malformed.
	for _, identifier := range []string{
		"trailing.",
		"empty..partial",
		".leading",
	} {
		t.Run("reject/malformed partial identifier "+identifier, func(t *testing.T) {
			parseFails(t, moduleScope(
				"/begin CHARACTERISTIC "+identifier+" \"\" VALUE 0x0 record_layout 0 compu_method 0 0\n"+
					"/end CHARACTERISTIC"))
		})
	}

	// Chapter 3.2: the brackets at the end of a partial identifier "must contain a number or an
	// alphabetic string (description of the index of an array element)", and the symbolic form is
	// "a symbolic string which is defined as an enumerator of an ENUM definition of the C
	// program". A signed literal is not such a number, and a C enumerator contains no point.
	for _, identifier := range []string{
		"array[-1]",
		"array[+1]",
		"array[symbolic.index]",
	} {
		t.Run("reject/malformed array index "+identifier, func(t *testing.T) {
			parseFails(t, moduleScope(
				"/begin CHARACTERISTIC "+identifier+" \"\" VALUE 0x0 record_layout 0 compu_method 0 0\n"+
					"/end CHARACTERISTIC"))
		})
	}

	t.Run("reject/leading digit", func(t *testing.T) {
		parseFails(t, moduleScope(
			"/begin CHARACTERISTIC 0identifier \"\" VALUE 0x0 record_layout 0 compu_method 0 0\n/end CHARACTERISTIC"))
	})

	// An identifier must not be one of the ASAP2 keywords.
	t.Run("reject/keyword used as identifier", func(t *testing.T) {
		parseFails(t, moduleScope(
			"/begin CHARACTERISTIC VALUE \"\" VALUE 0x0 record_layout 0 compu_method 0 0\n/end CHARACTERISTIC"))
	})
}

func TestGrammar_Terminal_STRING(t *testing.T) {
	type testCaseType struct {
		name    string
		literal string
		value   string
	}

	// Note: the tree stores the source form of a string, the escape sequences allowed by ASAP2
	// 1.51 (chapter 6.2) and ASAM MCD-2 MC 1.6.1 (chapter 3.2) are not decoded. The expected
	// values below therefore contain the backslashes literally.
	for _, testCase := range []testCaseType{
		{name: "accept/empty", literal: `""`, value: ""},
		{name: "accept/spaces", literal: `"with spaces"`, value: "with spaces"},
		{name: "accept/keyword", literal: `"/begin PROJECT"`, value: "/begin PROJECT"},
		{name: "accept/escaped quote", literal: `"escaped \" quote"`, value: `escaped \" quote`},
		{name: "accept/escaped carriage return", literal: `"line\r\n"`, value: `line\r\n`},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			module, ok := parseModule(t, "/begin MOD_COMMON "+testCase.literal+"\n/end MOD_COMMON")
			if !ok || !assert.NotNil(t, module.MOD_COMMON) {
				return
			}

			equalNode(t, strVal(testCase.value), module.MOD_COMMON.Comment)
		})
	}

	// Both versions allow a double inverted comma inside a string to be escaped by doubling it,
	// as an alternative to the backslash notation (compatibility with ASAP2 V1.2 and prior).
	t.Run("accept/doubled quote", func(t *testing.T) {
		module, ok := parseModule(t, "/begin MOD_COMMON \"doubled \"\"quote\"\"\"\n/end MOD_COMMON")
		if !ok || !assert.NotNil(t, module.MOD_COMMON) {
			return
		}

		equalNode(t, strVal("doubled \"\"quote\"\""), module.MOD_COMMON.Comment)
	})

	// ASAM MCD-2 MC 1.6.1 (chapter 3.2) allows the escape sequences \", \', \\, \n, \r and \t
	// (ASAP2 1.51 allowed \r\n and \" only); any other character behind a backslash is invalid.
	t.Run("accept/tab and backslash escapes", func(t *testing.T) {
		module, ok := parseModule(t, "/begin MOD_COMMON \"tab\\t and backslash \\\\\"\n/end MOD_COMMON")
		if !ok || !assert.NotNil(t, module.MOD_COMMON) {
			return
		}

		equalNode(t, strVal("tab\\t and backslash \\\\"), module.MOD_COMMON.Comment)
	})

	t.Run("reject/escape sequence not allowed by the specification", func(t *testing.T) {
		parseFails(t, moduleScope("/begin MOD_COMMON \"invalid \\q escape\"\n/end MOD_COMMON"))
	})

	// A malformed escape must not desynchronize the lexer. When the string does not lex as one
	// token, the lexer resumes inside its content, the closing quote opens a new string which
	// swallows the rest of the file, and the reported errors end up far away from the real fault.
	// The whole literal is therefore consumed and the offending sequence reported on its own.
	t.Run("reject/invalid escape is reported at the string, without a cascade", func(t *testing.T) {
		_, err := GetTreeFromString(moduleScope(
			"/begin MOD_COMMON \"a path C:\\data\\x\"\n/end MOD_COMMON\n" +
				"/begin MOD_PAR \"unrelated\"\n/end MOD_PAR"))
		if !assert.Error(t, err, "the invalid escape sequence should be rejected") {
			return
		}

		assert.Contains(t, err.Error(), "escape sequence",
			"the error should name the offending escape sequence")
		assert.NotContains(t, err.Error(), "token recognition error",
			"the string should lex as a single token instead of desynchronizing the lexer")
		assert.NotContains(t, err.Error(), "error while building A2L tree",
			"the parser should report a syntax error instead of an internal error")
	})

	t.Run("reject/unterminated string", func(t *testing.T) {
		parseFails(t, moduleScope("/begin MOD_COMMON \"unterminated\n/end MOD_COMMON"))
	})
}

func TestGrammar_Terminal_INT(t *testing.T) {
	type testCaseType struct {
		literal string
		value   int32
	}

	for _, testCase := range []testCaseType{
		{literal: "0", value: 0},
		{literal: "1", value: 1},
		{literal: "255", value: 255},
		{literal: "-1", value: -1},
		{literal: "+1", value: 1},
		{literal: "0001", value: 1},
	} {
		t.Run("accept/"+testCase.literal, func(t *testing.T) {
			modCommon, ok := parseModCommon(t, "DATA_SIZE "+testCase.literal)
			if !ok || !assert.NotNil(t, modCommon.DATA_SIZE) {
				return
			}

			assert.Equal(t, testCase.value, modCommon.DATA_SIZE.Size.Value)
			assert.Equal(t, uint32(10), modCommon.DATA_SIZE.Size.Base)
		})
	}

	t.Run("reject/space between sign and digits", func(t *testing.T) {
		parseFails(t, modCommonScope("DATA_SIZE - 1"))
	})

	t.Run("reject/thousands separator", func(t *testing.T) {
		parseFails(t, modCommonScope("DATA_SIZE 1'000"))
	})
}

func TestGrammar_Terminal_HEX(t *testing.T) {
	type testCaseType struct {
		literal string
		value   int64
	}

	for _, testCase := range []testCaseType{
		{literal: "0x0", value: 0},
		{literal: "0xFF", value: 255},
		{literal: "0xff", value: 255},
		{literal: "0X10", value: 16},
		{literal: "0x00001234", value: 0x1234},
	} {
		t.Run("accept/"+testCase.literal, func(t *testing.T) {
			module, ok := parseModule(t,
				"/begin CHARACTERISTIC characteristic \"\" VALUE "+testCase.literal+
					" record_layout 0 compu_method 0 0\n/end CHARACTERISTIC")
			if !ok || !assert.Len(t, module.CHARACTERISTIC, 1) {
				return
			}

			address := module.CHARACTERISTIC[0].Address
			if assert.NotNil(t, address) {
				assert.Equal(t, testCase.value, address.Value)
				assert.Equal(t, uint32(16), address.Base)
			}
		})
	}

	t.Run("reject/missing 0x prefix", func(t *testing.T) {
		parseFails(t, moduleScope(
			"/begin CHARACTERISTIC characteristic \"\" VALUE FF record_layout 0 compu_method 0 0\n"+
				"/end CHARACTERISTIC"))
	})

	t.Run("reject/negative hexadecimal value", func(t *testing.T) {
		parseFails(t, moduleScope(
			"/begin CHARACTERISTIC characteristic \"\" VALUE -0x10 record_layout 0 compu_method 0 0\n"+
				"/end CHARACTERISTIC"))
	})
}

func TestGrammar_Terminal_FLOAT(t *testing.T) {
	type testCaseType struct {
		literal string
		value   float64
	}

	for _, testCase := range []testCaseType{
		{literal: "0", value: 0},
		{literal: "0.5", value: 0.5},
		{literal: "-0.5", value: -0.5},
		{literal: "+0.5", value: 0.5},
		{literal: ".5", value: 0.5},
		{literal: "1.", value: 1},
		{literal: "1e3", value: 1000},
		{literal: "1.5e-3", value: 0.0015},
		{literal: "1.5e+3", value: 1500},
	} {
		t.Run("accept/"+testCase.literal, func(t *testing.T) {
			tree, ok := parse(t, axisDescrScope("MAX_GRAD "+testCase.literal))
			if !ok {
				return
			}

			axisDescr, ok := axisDescrOf(t, tree)
			if !ok || !assert.NotNil(t, axisDescr.MAX_GRAD) {
				return
			}

			assert.Equal(t, testCase.value, axisDescr.MAX_GRAD.MaxGradient.Value)
		})
	}

	// The source form of a float value is part of the tree and is reproduced verbatim by the
	// serializer ("12E-2" is the exponential notation example of ASAM MCD-2 MC 1.6.1, chapter
	// 3.2).
	for _, literal := range []string{"1.50", "+0.5", ".5", "1.", "1.5e3", "1e3", "12E-2", "1.5e-3"} {
		t.Run("layout of the value is preserved/"+literal, func(t *testing.T) {
			parse(t, axisDescrScope("MAX_GRAD "+literal))
		})
	}

	// Both versions attach the note about the hexadecimal notation to the integer data types only
	// (ASAP2 1.51 chapter 6.2, ASAM MCD-2 MC 1.6.1 chapter 3.2); neither states the restriction
	// explicitly, so this is the reading the grammar implements, see floatValue.
	t.Run("reject/hexadecimal value", func(t *testing.T) {
		_, err := GetTreeFromString(axisDescrScope("MAX_GRAD 0x10"))
		if assert.Error(t, err, "a hexadecimal value should not be accepted as a float") {
			assert.NotContains(t, err.Error(), "error while building A2L tree",
				"the parser should report a syntax error instead of an internal error")
		}
	})

	t.Run("reject/comma as decimal separator", func(t *testing.T) {
		parseFails(t, axisDescrScope("MAX_GRAD 1,5"))
	})
}

func TestGrammar_Terminal_Comment(t *testing.T) {
	t.Run("accept/block comment", func(t *testing.T) {
		parse(t, projectScope("/* a comment */"))
	})

	t.Run("accept/block comment between parameters", func(t *testing.T) {
		module, ok := parseModule(t,
			"/begin MOD_COMMON /* comment */ \"comment\" /* comment */\n/end MOD_COMMON")
		if !ok || !assert.NotNil(t, module.MOD_COMMON) {
			return
		}

		equalNode(t, strVal("comment"), module.MOD_COMMON.Comment)
	})

	t.Run("accept/line comment", func(t *testing.T) {
		parse(t, projectScope("// a comment"))
	})

	// A line comment on the last line of a file needs no terminating line break.
	t.Run("accept/line comment at end of file", func(t *testing.T) {
		parse(t, projectScope("")+"// a comment without line break")
	})

	t.Run("reject/unterminated block comment", func(t *testing.T) {
		parseFails(t, projectScope("/* unterminated"))
	})
}

func TestGrammar_Terminal_BlockDelimiter(t *testing.T) {
	t.Run("reject/uppercase begin", func(t *testing.T) {
		parseFails(t, "/BEGIN PROJECT project \"\"\n/end PROJECT")
	})

	t.Run("reject/mismatched block name", func(t *testing.T) {
		parseFails(t, "/begin PROJECT project \"\"\n/end MODULE")
	})

	t.Run("reject/missing end", func(t *testing.T) {
		parseFails(t, "/begin PROJECT project \"\"")
	})

	// ASAP2 1.51 (chapter 6.3) declares the short form "<keyword> { <description_body> }" as an
	// alternative to the /begin and /end delimiters, but no longer recommends it since version
	// 1.31. ASAM MCD-2 MC 1.6.1 (chapter 1.4.2) removed it: "Since ASAM MCD-2 MC V1.6.0 always
	// brackets of the form '/begin' '/end' are requested. Curly brackets '{' '}' are no longer
	// supported."
	t.Run("reject/short delimiters", func(t *testing.T) {
		parseFails(t, "PROJECT { project \"\" }")
	})
}
