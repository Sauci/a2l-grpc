package a2l

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrammar_ANNOTATION(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		characteristic, ok := parseCharacteristic(t, "/begin ANNOTATION\n/end ANNOTATION")
		if !ok {
			return
		}

		equalNodes(t, []*AnnotationType{{}}, characteristic.ANNOTATION)
	})

	t.Run("optional/all parameters", func(t *testing.T) {
		characteristic, ok := parseCharacteristic(t, `/begin ANNOTATION
ANNOTATION_LABEL "label"
ANNOTATION_ORIGIN "origin"
/begin ANNOTATION_TEXT "text"
/end ANNOTATION_TEXT
/end ANNOTATION`)
		if !ok {
			return
		}

		equalNodes(t, []*AnnotationType{{
			ANNOTATION_LABEL:  &AnnotationLabelType{Label: strVal("label")},
			ANNOTATION_ORIGIN: &AnnotationOriginType{Origin: strVal("origin")},
			ANNOTATION_TEXT:   &AnnotationTextType{AnnotationText: []*StringType{strVal("text")}},
		}}, characteristic.ANNOTATION)
	})

	t.Run("list/several ANNOTATION", func(t *testing.T) {
		characteristic, ok := parseCharacteristic(t, `/begin ANNOTATION
ANNOTATION_LABEL "first"
/end ANNOTATION
/begin ANNOTATION
ANNOTATION_LABEL "second"
/end ANNOTATION`)
		if !ok {
			return
		}

		equalNodes(t, []*AnnotationType{
			{ANNOTATION_LABEL: &AnnotationLabelType{Label: strVal("first")}},
			{ANNOTATION_LABEL: &AnnotationLabelType{Label: strVal("second")}},
		}, characteristic.ANNOTATION)
	})

	t.Run("in AXIS_DESCR", func(t *testing.T) {
		axisDescr, ok := parseAxisDescr(t, "/begin ANNOTATION\n/end ANNOTATION")
		if !ok {
			return
		}

		assert.Len(t, axisDescr.ANNOTATION, 1)
	})

	t.Run("reject/missing /end", func(t *testing.T) {
		parseFails(t, characteristicScope("/begin ANNOTATION"))
	})
}

func TestGrammar_ANNOTATION_LABEL(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		characteristic, ok := parseCharacteristic(t,
			"/begin ANNOTATION\nANNOTATION_LABEL \"label\"\n/end ANNOTATION")
		if !ok || !assert.Len(t, characteristic.ANNOTATION, 1) {
			return
		}

		equalNode(t, &AnnotationLabelType{Label: strVal("label")}, characteristic.ANNOTATION[0].ANNOTATION_LABEL)
	})

	t.Run("reject/identifier parameter", func(t *testing.T) {
		parseFails(t, characteristicScope("/begin ANNOTATION\nANNOTATION_LABEL label\n/end ANNOTATION"))
	})

	t.Run("reject/missing label", func(t *testing.T) {
		parseFails(t, characteristicScope("/begin ANNOTATION\nANNOTATION_LABEL\n/end ANNOTATION"))
	})
}

func TestGrammar_ANNOTATION_ORIGIN(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		characteristic, ok := parseCharacteristic(t,
			"/begin ANNOTATION\nANNOTATION_ORIGIN \"origin\"\n/end ANNOTATION")
		if !ok || !assert.Len(t, characteristic.ANNOTATION, 1) {
			return
		}

		equalNode(t, &AnnotationOriginType{Origin: strVal("origin")}, characteristic.ANNOTATION[0].ANNOTATION_ORIGIN)
	})

	t.Run("reject/missing origin", func(t *testing.T) {
		parseFails(t, characteristicScope("/begin ANNOTATION\nANNOTATION_ORIGIN\n/end ANNOTATION"))
	})
}

func TestGrammar_ANNOTATION_TEXT(t *testing.T) {
	type testCaseType struct {
		name  string
		body  string
		texts []*StringType
	}

	for _, testCase := range []testCaseType{
		{
			name:  "list/no text",
			body:  "/begin ANNOTATION_TEXT\n/end ANNOTATION_TEXT",
			texts: nil,
		},
		{
			name:  "list/single text",
			body:  "/begin ANNOTATION_TEXT \"first line\"\n/end ANNOTATION_TEXT",
			texts: []*StringType{strVal("first line")},
		},
		{
			name: "list/several texts",
			body: "/begin ANNOTATION_TEXT \"first line\" \"second line\" \"third line\"\n/end ANNOTATION_TEXT",
			texts: []*StringType{
				strVal("first line"),
				strVal("second line"),
				strVal("third line"),
			},
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			characteristic, ok := parseCharacteristic(t, "/begin ANNOTATION\n"+testCase.body+"\n/end ANNOTATION")
			if !ok || !assert.Len(t, characteristic.ANNOTATION, 1) {
				return
			}

			annotationText := characteristic.ANNOTATION[0].ANNOTATION_TEXT
			if !assert.NotNil(t, annotationText, "ANNOTATION_TEXT is missing from the tree") {
				return
			}

			equalNodes(t, testCase.texts, annotationText.AnnotationText)
		})
	}

	t.Run("escaped double quote", func(t *testing.T) {
		characteristic, ok := parseCharacteristic(t,
			"/begin ANNOTATION\n/begin ANNOTATION_TEXT \"escaped \\\" quote\"\n/end ANNOTATION_TEXT\n/end ANNOTATION")
		if !ok || !assert.Len(t, characteristic.ANNOTATION, 1) {
			return
		}

		equalNodes(t, []*StringType{{Value: "escaped \\\" quote"}},
			characteristic.ANNOTATION[0].ANNOTATION_TEXT.AnnotationText)
	})

	t.Run("reject/identifier parameter", func(t *testing.T) {
		parseFails(t, characteristicScope(
			"/begin ANNOTATION\n/begin ANNOTATION_TEXT text\n/end ANNOTATION_TEXT\n/end ANNOTATION"))
	})
}
