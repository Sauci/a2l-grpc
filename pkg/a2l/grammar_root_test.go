package a2l

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

	t.Run("reject/HEADER after MODULE", func(t *testing.T) {
		parseFails(t, projectScope(
			"/begin MODULE module \"\" /end MODULE\n/begin HEADER \"comment\"\n/end HEADER"))
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
