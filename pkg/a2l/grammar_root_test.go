package a2l

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// The byte-order mark declares the encoding of the file (ASAM MCD-2 MC 1.6.1, chapter 1.5.2) and
// is not part of its content. UTF-8 is the preferred encoding, and chapter 1.5.1 requires tools to
// support at least it, so a leading UTF-8 mark has to be accepted and ignored.
func TestGrammar_ByteOrderMark(t *testing.T) {
	const byteOrderMark = "\uFEFF"

	t.Run("accept/before ASAP2_VERSION", func(t *testing.T) {
		tree, ok := parse(t, byteOrderMark+"ASAP2_VERSION 1 61\n"+projectScope(""))
		if !ok {
			return
		}

		equalNode(t, &ASAP2VersionType{
			VersionNo: intVal("1"),
			UpgradeNo: intVal("61"),
		}, tree.ASAP2_VERSION)
	})

	t.Run("accept/before PROJECT", func(t *testing.T) {
		tree, ok := parse(t, byteOrderMark+projectScope(""))
		if !ok || !assert.NotNil(t, tree.PROJECT) {
			return
		}

		equalNode(t, identVal("project"), tree.PROJECT.Name)
	})

	// The mark carries no information besides the encoding, so it must not reach the tree.
	t.Run("tree is identical to the same content without the mark", func(t *testing.T) {
		expected, ok := parseOnly(t, "ASAP2_VERSION 1 61\n"+projectScope(""))
		if !ok {
			return
		}

		actual, ok := parseOnly(t, byteOrderMark+"ASAP2_VERSION 1 61\n"+projectScope(""))
		if !ok {
			return
		}

		equalNode(t, expected, actual)
	})

	// Chapter 1.5.2 defines the mark as a byte sequence at the beginning of the file; the same
	// sequence anywhere else is not a mark and stays invalid.
	t.Run("reject/inside the file", func(t *testing.T) {
		parseFails(t, projectScope(byteOrderMark))
	})
}

func TestGrammar_ASAP2_VERSION(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		tree, ok := parse(t, "ASAP2_VERSION 1 51\n"+projectScope(""))
		if !ok {
			return
		}

		equalNode(t, &ASAP2VersionType{
			VersionNo: intVal("1"),
			UpgradeNo: intVal("51"),
		}, tree.ASAP2_VERSION)
	})

	t.Run("is optional", func(t *testing.T) {
		tree, ok := parse(t, projectScope(""))
		if !ok {
			return
		}

		assert.Nil(t, tree.ASAP2_VERSION)
	})

	t.Run("reject/missing upgrade number", func(t *testing.T) {
		parseFails(t, "ASAP2_VERSION 1\n"+projectScope(""))
	})

	t.Run("reject/float version number", func(t *testing.T) {
		parseFails(t, "ASAP2_VERSION 1.5 1\n"+projectScope(""))
	})
}

func TestGrammar_A2ML_VERSION(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		tree, ok := parse(t, "A2ML_VERSION 1 31\n"+projectScope(""))
		if !ok {
			return
		}

		equalNode(t, &A2MLVersionType{
			VersionNo: intVal("1"),
			UpgradeNo: intVal("31"),
		}, tree.A2ML_VERSION)
	})

	t.Run("follows ASAP2_VERSION", func(t *testing.T) {
		tree, ok := parse(t, "ASAP2_VERSION 1 51\nA2ML_VERSION 1 31\n"+projectScope(""))
		if !ok {
			return
		}

		assert.NotNil(t, tree.ASAP2_VERSION)
		assert.NotNil(t, tree.A2ML_VERSION)
	})

	t.Run("reject/declared before ASAP2_VERSION", func(t *testing.T) {
		parseFails(t, "A2ML_VERSION 1 31\nASAP2_VERSION 1 51\n"+projectScope(""))
	})
}

func TestGrammar_PROJECT(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		project, ok := parseProject(t, "")
		if !ok {
			return
		}

		equalNode(t, &ProjectType{
			Name:           identVal("project"),
			LongIdentifier: strVal(""),
		}, project)
	})

	t.Run("list/no MODULE", func(t *testing.T) {
		project, ok := parseProject(t, "")
		if !ok {
			return
		}

		assert.Empty(t, project.MODULE)
	})

	t.Run("list/several MODULE", func(t *testing.T) {
		project, ok := parseProject(t,
			"/begin MODULE first \"\" /end MODULE\n/begin MODULE second \"\" /end MODULE")
		if !ok {
			return
		}

		if assert.Len(t, project.MODULE, 2) {
			assert.Equal(t, "first", project.MODULE[0].Name.Value)
			assert.Equal(t, "second", project.MODULE[1].Name.Value)
		}
	})

	t.Run("reject/missing long identifier", func(t *testing.T) {
		parseFails(t, fileScope("/begin PROJECT project\n/end PROJECT"))
	})

	t.Run("reject/empty file", func(t *testing.T) {
		parseFails(t, "")
	})

	// The description file consists of a single PROJECT block, content behind it is invalid.
	t.Run("reject/content after PROJECT", func(t *testing.T) {
		parseFails(t, projectScope("")+"\nnot an A2L keyword\n")
	})
}

func TestGrammar_HEADER(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		project, ok := parseProject(t, "/begin HEADER \"comment\"\n/end HEADER")
		if !ok {
			return
		}

		equalNode(t, &HeaderType{Comment: strVal("comment")}, project.HEADER)
	})

	t.Run("optional/VERSION and PROJECT_NO", func(t *testing.T) {
		project, ok := parseProject(t,
			"/begin HEADER \"comment\"\nVERSION \"1.0\"\nPROJECT_NO project_number\n/end HEADER")
		if !ok {
			return
		}

		equalNode(t, &HeaderType{
			Comment:    strVal("comment"),
			VERSION:    &VersionType{VersionIdentifier: strVal("1.0")},
			PROJECT_NO: &ProjectNoType{ProjectNumber: identVal("project_number")},
		}, project.HEADER)
	})

	t.Run("reject/missing comment", func(t *testing.T) {
		parseFails(t, projectScope("/begin HEADER\n/end HEADER"))
	})

	// Chapter 6.3.101 declares [-> HEADER]: at most one HEADER, before the MODULE blocks.
	t.Run("reject/several HEADER", func(t *testing.T) {
		parseFails(t, projectScope(
			"/begin HEADER \"first\"\n/end HEADER\n/begin HEADER \"second\"\n/end HEADER"))
	})

	// Neither version prescribes an order for the optional elements. The prototype of chapter
	// 3.5.99 lists HEADER before MODULE, but that is the order of the listing and not a syntax
	// rule: the AXIS_PTS example of chapter 3.5.18 emits its own optional elements out of
	// prototype order. Only the cardinality, "[-> HEADER]", is prescribed.
	t.Run("HEADER after MODULE", func(t *testing.T) {
		project, ok := parseProject(t,
			"/begin MODULE module \"\" /end MODULE\n/begin HEADER \"comment\"\n/end HEADER")
		if !ok {
			return
		}

		equalNode(t, &HeaderType{Comment: strVal("comment")}, project.HEADER)
		assert.Len(t, project.MODULE, 1)
	})

	t.Run("HEADER between two MODULE", func(t *testing.T) {
		project, ok := parseProject(t,
			"/begin MODULE first \"\" /end MODULE\n/begin HEADER \"comment\"\n/end HEADER\n"+
				"/begin MODULE second \"\" /end MODULE")
		if !ok {
			return
		}

		equalNode(t, &HeaderType{Comment: strVal("comment")}, project.HEADER)

		if assert.Len(t, project.MODULE, 2) {
			assert.Equal(t, "first", project.MODULE[0].Name.Value)
			assert.Equal(t, "second", project.MODULE[1].Name.Value)
		}
	})
}

func TestGrammar_PROJECT_NO(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		project, ok := parseProject(t, "/begin HEADER \"\"\nPROJECT_NO project_number\n/end HEADER")
		if !ok || !assert.NotNil(t, project.HEADER) {
			return
		}

		equalNode(t, &ProjectNoType{ProjectNumber: identVal("project_number")}, project.HEADER.PROJECT_NO)
	})

	t.Run("reject/string parameter", func(t *testing.T) {
		parseFails(t, projectScope("/begin HEADER \"\"\nPROJECT_NO \"project_number\"\n/end HEADER"))
	})
}

func TestGrammar_VERSION(t *testing.T) {
	t.Run("mandatory parameters/in HEADER", func(t *testing.T) {
		project, ok := parseProject(t, "/begin HEADER \"\"\nVERSION \"version\"\n/end HEADER")
		if !ok || !assert.NotNil(t, project.HEADER) {
			return
		}

		equalNode(t, &VersionType{VersionIdentifier: strVal("version")}, project.HEADER.VERSION)
	})

	t.Run("mandatory parameters/in MOD_PAR", func(t *testing.T) {
		module, ok := parseModule(t, "/begin MOD_PAR \"\"\nVERSION \"version\"\n/end MOD_PAR")
		if !ok || !assert.NotNil(t, module.MOD_PAR) {
			return
		}

		equalNode(t, &VersionType{VersionIdentifier: strVal("version")}, module.MOD_PAR.VERSION)
	})

	t.Run("reject/identifier parameter", func(t *testing.T) {
		parseFails(t, projectScope("/begin HEADER \"\"\nVERSION version\n/end HEADER"))
	})
}

func TestGrammar_MODULE(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		module, ok := parseModule(t, "")
		if !ok {
			return
		}

		equalNode(t, &ModuleType{
			Name:           identVal("module"),
			LongIdentifier: strVal(""),
		}, module)
	})

	t.Run("list/several CHARACTERISTIC", func(t *testing.T) {
		module, ok := parseModule(t,
			"/begin CHARACTERISTIC first \"\" VALUE 0x0 record_layout 0 compu_method 0 0 /end CHARACTERISTIC\n"+
				"/begin CHARACTERISTIC second \"\" VALUE 0x0 record_layout 0 compu_method 0 0 /end CHARACTERISTIC")
		if !ok {
			return
		}

		assert.Len(t, module.CHARACTERISTIC, 2)
	})

	t.Run("reject/missing long identifier", func(t *testing.T) {
		parseFails(t, projectScope("/begin MODULE module\n/end MODULE"))
	})

	t.Run("reject/unknown keyword", func(t *testing.T) {
		parseFails(t, moduleScope("NOT_A_KEYWORD 1"))
	})
}
