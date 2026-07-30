package a2l

// FUNCTION (chapter 6.3.65), GROUP (chapter 6.3.68), USER_RIGHTS (chapter 6.3.130) and the
// identifier lists they contain.

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrammar_FUNCTION(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		module, ok := parseModule(t, "/begin FUNCTION function \"long identifier\"\n/end FUNCTION")
		if !ok {
			return
		}

		equalNodes(t, []*FunctionType{{
			Name:           identVal("function"),
			LongIdentifier: strVal("long identifier"),
		}}, module.FUNCTION)
	})

	t.Run("optional/all parameters", func(t *testing.T) {
		function, ok := parseFunction(t, `/begin ANNOTATION
ANNOTATION_LABEL "label"
/end ANNOTATION
/begin DEF_CHARACTERISTIC defined
/end DEF_CHARACTERISTIC
/begin REF_CHARACTERISTIC referenced
/end REF_CHARACTERISTIC
/begin IN_MEASUREMENT input
/end IN_MEASUREMENT
/begin OUT_MEASUREMENT output
/end OUT_MEASUREMENT
/begin LOC_MEASUREMENT local
/end LOC_MEASUREMENT
/begin SUB_FUNCTION sub_function
/end SUB_FUNCTION
FUNCTION_VERSION "1.0"`)
		if !ok {
			return
		}

		equalNode(t, &FunctionType{
			Name:           identVal("function"),
			LongIdentifier: strVal(""),
			ANNOTATION: []*AnnotationType{
				{ANNOTATION_LABEL: &AnnotationLabelType{Label: strVal("label")}},
			},
			DEF_CHARACTERISTIC: &DefCharacteristicType{Identifier: []*IdentType{identVal("defined")}},
			REF_CHARACTERISTIC: &RefCharacteristicType{Identifier: []*IdentType{identVal("referenced")}},
			IN_MEASUREMENT:     &InMeasurementType{Identifier: []*IdentType{identVal("input")}},
			OUT_MEASUREMENT:    &OutMeasurementType{Identifier: []*IdentType{identVal("output")}},
			LOC_MEASUREMENT:    &LocMeasurementType{Identifier: []*IdentType{identVal("local")}},
			SUB_FUNCTION:       &SubFunctionType{Identifier: []*IdentType{identVal("sub_function")}},
			FUNCTION_VERSION:   &FunctionVersionType{VersionIdentifier: strVal("1.0")},
		}, function)
	})

	t.Run("reject/missing long identifier", func(t *testing.T) {
		parseFails(t, moduleScope("/begin FUNCTION function\n/end FUNCTION"))
	})

	// ASAP2 1.51 does not declare IF_DATA for FUNCTION.
	t.Run("reject/IF_DATA", func(t *testing.T) {
		deviation(t, "FUNCTION accepts IF_DATA and then fails with an internal error",
			func(t assert.TestingT) {
				parseFailsWithSyntaxError(t, functionScope("/begin IF_DATA XCP\n/end IF_DATA"))
			})
	})
}

func TestGrammar_DEF_CHARACTERISTIC(t *testing.T) {
	testIdentifierListKeyword(t, "DEF_CHARACTERISTIC", functionScope,
		func(identifiers []*IdentType) *DefCharacteristicType {
			return &DefCharacteristicType{Identifier: identifiers}
		},
		func(t assert.TestingT, body string) (*DefCharacteristicType, bool) {
			function, ok := parseFunction(t, body)
			if !ok {
				return nil, false
			}

			return function.DEF_CHARACTERISTIC, assert.NotNil(t, function.DEF_CHARACTERISTIC,
				"DEF_CHARACTERISTIC is missing from the tree")
		})
}

func TestGrammar_REF_CHARACTERISTIC(t *testing.T) {
	testIdentifierListKeyword(t, "REF_CHARACTERISTIC", functionScope,
		func(identifiers []*IdentType) *RefCharacteristicType {
			return &RefCharacteristicType{Identifier: identifiers}
		},
		func(t assert.TestingT, body string) (*RefCharacteristicType, bool) {
			function, ok := parseFunction(t, body)
			if !ok {
				return nil, false
			}

			return function.REF_CHARACTERISTIC, assert.NotNil(t, function.REF_CHARACTERISTIC,
				"REF_CHARACTERISTIC is missing from the tree")
		})

	t.Run("in GROUP", func(t *testing.T) {
		group, ok := parseGroup(t, "/begin REF_CHARACTERISTIC first second\n/end REF_CHARACTERISTIC")
		if !ok {
			return
		}

		equalNode(t, &RefCharacteristicType{Identifier: []*IdentType{identVal("first"), identVal("second")}},
			group.REF_CHARACTERISTIC)
	})
}

func TestGrammar_IN_MEASUREMENT(t *testing.T) {
	testIdentifierListKeyword(t, "IN_MEASUREMENT", functionScope,
		func(identifiers []*IdentType) *InMeasurementType {
			return &InMeasurementType{Identifier: identifiers}
		},
		func(t assert.TestingT, body string) (*InMeasurementType, bool) {
			function, ok := parseFunction(t, body)
			if !ok {
				return nil, false
			}

			return function.IN_MEASUREMENT, assert.NotNil(t, function.IN_MEASUREMENT,
				"IN_MEASUREMENT is missing from the tree")
		})
}

func TestGrammar_OUT_MEASUREMENT(t *testing.T) {
	testIdentifierListKeyword(t, "OUT_MEASUREMENT", functionScope,
		func(identifiers []*IdentType) *OutMeasurementType {
			return &OutMeasurementType{Identifier: identifiers}
		},
		func(t assert.TestingT, body string) (*OutMeasurementType, bool) {
			function, ok := parseFunction(t, body)
			if !ok {
				return nil, false
			}

			return function.OUT_MEASUREMENT, assert.NotNil(t, function.OUT_MEASUREMENT,
				"OUT_MEASUREMENT is missing from the tree")
		})
}

func TestGrammar_LOC_MEASUREMENT(t *testing.T) {
	testIdentifierListKeyword(t, "LOC_MEASUREMENT", functionScope,
		func(identifiers []*IdentType) *LocMeasurementType {
			return &LocMeasurementType{Identifier: identifiers}
		},
		func(t assert.TestingT, body string) (*LocMeasurementType, bool) {
			function, ok := parseFunction(t, body)
			if !ok {
				return nil, false
			}

			return function.LOC_MEASUREMENT, assert.NotNil(t, function.LOC_MEASUREMENT,
				"LOC_MEASUREMENT is missing from the tree")
		})
}

func TestGrammar_SUB_FUNCTION(t *testing.T) {
	testIdentifierListKeyword(t, "SUB_FUNCTION", functionScope,
		func(identifiers []*IdentType) *SubFunctionType {
			return &SubFunctionType{Identifier: identifiers}
		},
		func(t assert.TestingT, body string) (*SubFunctionType, bool) {
			function, ok := parseFunction(t, body)
			if !ok {
				return nil, false
			}

			return function.SUB_FUNCTION, assert.NotNil(t, function.SUB_FUNCTION,
				"SUB_FUNCTION is missing from the tree")
		})
}

func TestGrammar_FUNCTION_VERSION(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		function, ok := parseFunction(t, "FUNCTION_VERSION \"1.0\"")
		if !ok {
			return
		}

		equalNode(t, &FunctionVersionType{VersionIdentifier: strVal("1.0")}, function.FUNCTION_VERSION)
	})

	t.Run("reject/numeric version", func(t *testing.T) {
		parseFails(t, functionScope("FUNCTION_VERSION 1.0"))
	})
}

func TestGrammar_FUNCTION_LIST(t *testing.T) {
	testIdentifierListKeyword(t, "FUNCTION_LIST", characteristicScope,
		func(identifiers []*IdentType) *FunctionListType {
			return &FunctionListType{Name: identifiers}
		},
		func(t assert.TestingT, body string) (*FunctionListType, bool) {
			characteristic, ok := parseCharacteristic(t, body)
			if !ok {
				return nil, false
			}

			return characteristic.FUNCTION_LIST, assert.NotNil(t, characteristic.FUNCTION_LIST,
				"FUNCTION_LIST is missing from the tree")
		})

	t.Run("in GROUP", func(t *testing.T) {
		group, ok := parseGroup(t, "/begin FUNCTION_LIST function\n/end FUNCTION_LIST")
		if !ok {
			return
		}

		equalNode(t, &FunctionListType{Name: []*IdentType{identVal("function")}}, group.FUNCTION_LIST)
	})
}

func TestGrammar_GROUP(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		module, ok := parseModule(t, "/begin GROUP group \"long identifier\"\n/end GROUP")
		if !ok {
			return
		}

		equalNodes(t, []*GroupType{{
			GroupName:           identVal("group"),
			GroupLongIdentifier: strVal("long identifier"),
		}}, module.GROUP)
	})

	t.Run("optional/all parameters", func(t *testing.T) {
		group, ok := parseGroup(t, `/begin ANNOTATION
ANNOTATION_LABEL "label"
/end ANNOTATION
ROOT
/begin REF_CHARACTERISTIC characteristic
/end REF_CHARACTERISTIC
/begin REF_MEASUREMENT measurement
/end REF_MEASUREMENT
/begin FUNCTION_LIST function
/end FUNCTION_LIST
/begin SUB_GROUP sub_group
/end SUB_GROUP`)
		if !ok {
			return
		}

		equalNode(t, &GroupType{
			GroupName:           identVal("group"),
			GroupLongIdentifier: strVal(""),
			ANNOTATION: []*AnnotationType{
				{ANNOTATION_LABEL: &AnnotationLabelType{Label: strVal("label")}},
			},
			ROOT:               &RootType{Present: true},
			REF_CHARACTERISTIC: &RefCharacteristicType{Identifier: []*IdentType{identVal("characteristic")}},
			REF_MEASUREMENT:    &RefMeasurementType{Identifier: []*IdentType{identVal("measurement")}},
			FUNCTION_LIST:      &FunctionListType{Name: []*IdentType{identVal("function")}},
			SUB_GROUP:          &SubGroupType{Identifier: []*IdentType{identVal("sub_group")}},
		}, group)
	})

	t.Run("reject/missing long identifier", func(t *testing.T) {
		parseFails(t, moduleScope("/begin GROUP group\n/end GROUP"))
	})

	// Deviation: the specification declares SUB_GROUP once per group. The grammar accepts a
	// repetition, and only the last one is kept in the tree.
	t.Run("reject/repeated SUB_GROUP", func(t *testing.T) {
		deviation(t, "optional keywords may be repeated", func(t assert.TestingT) {
			parseFails(t, groupScope("/begin SUB_GROUP first\n/end SUB_GROUP\n"+
				"/begin SUB_GROUP second\n/end SUB_GROUP"))
		})
	})
}

func TestGrammar_ROOT(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		group, ok := parseGroup(t, "ROOT")
		if !ok {
			return
		}

		equalNode(t, &RootType{Present: true}, group.ROOT)
	})

	t.Run("reject/with parameter", func(t *testing.T) {
		parseFails(t, groupScope("ROOT root"))
	})

	t.Run("reject/outside of GROUP", func(t *testing.T) {
		parseFails(t, functionScope("ROOT"))
	})
}

func TestGrammar_REF_MEASUREMENT(t *testing.T) {
	testIdentifierListKeyword(t, "REF_MEASUREMENT", groupScope,
		func(identifiers []*IdentType) *RefMeasurementType {
			return &RefMeasurementType{Identifier: identifiers}
		},
		func(t assert.TestingT, body string) (*RefMeasurementType, bool) {
			group, ok := parseGroup(t, body)
			if !ok {
				return nil, false
			}

			return group.REF_MEASUREMENT, assert.NotNil(t, group.REF_MEASUREMENT,
				"REF_MEASUREMENT is missing from the tree")
		})
}

func TestGrammar_SUB_GROUP(t *testing.T) {
	testIdentifierListKeyword(t, "SUB_GROUP", groupScope,
		func(identifiers []*IdentType) *SubGroupType {
			return &SubGroupType{Identifier: identifiers}
		},
		func(t assert.TestingT, body string) (*SubGroupType, bool) {
			group, ok := parseGroup(t, body)
			if !ok {
				return nil, false
			}

			return group.SUB_GROUP, assert.NotNil(t, group.SUB_GROUP, "SUB_GROUP is missing from the tree")
		})
}

func TestGrammar_USER_RIGHTS(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		module, ok := parseModule(t, "/begin USER_RIGHTS user_level\n/end USER_RIGHTS")
		if !ok {
			return
		}

		equalNodes(t, []*UserRightsType{{UserLevelId: identVal("user_level")}}, module.USER_RIGHTS)
	})

	t.Run("optional/all parameters", func(t *testing.T) {
		userRights, ok := parseUserRights(t, "/begin REF_GROUP group\n/end REF_GROUP\nREAD_ONLY")
		if !ok {
			return
		}

		equalNode(t, &UserRightsType{
			UserLevelId: identVal("user_level"),
			REF_GROUP:   []*RefGroupType{{Identifier: []*IdentType{identVal("group")}}},
			READ_ONLY:   &ReadOnlyType{Present: true},
		}, userRights)
	})

	t.Run("list/several REF_GROUP", func(t *testing.T) {
		userRights, ok := parseUserRights(t,
			"/begin REF_GROUP first\n/end REF_GROUP\n/begin REF_GROUP second\n/end REF_GROUP")
		if !ok {
			return
		}

		equalNodes(t, []*RefGroupType{
			{Identifier: []*IdentType{identVal("first")}},
			{Identifier: []*IdentType{identVal("second")}},
		}, userRights.REF_GROUP)
	})

	t.Run("reject/string as user level", func(t *testing.T) {
		parseFails(t, moduleScope("/begin USER_RIGHTS \"user_level\"\n/end USER_RIGHTS"))
	})

	t.Run("reject/missing user level", func(t *testing.T) {
		parseFails(t, moduleScope("/begin USER_RIGHTS\n/end USER_RIGHTS"))
	})
}

func TestGrammar_REF_GROUP(t *testing.T) {
	t.Run("list/no identifier", func(t *testing.T) {
		userRights, ok := parseUserRights(t, "/begin REF_GROUP\n/end REF_GROUP")
		if !ok {
			return
		}

		equalNodes(t, []*RefGroupType{{}}, userRights.REF_GROUP)
	})

	t.Run("list/several identifiers", func(t *testing.T) {
		userRights, ok := parseUserRights(t, "/begin REF_GROUP first second third\n/end REF_GROUP")
		if !ok {
			return
		}

		equalNodes(t, []*RefGroupType{
			{Identifier: []*IdentType{identVal("first"), identVal("second"), identVal("third")}},
		}, userRights.REF_GROUP)
	})

	t.Run("reject/string instead of identifier", func(t *testing.T) {
		parseFails(t, userRightsScope("/begin REF_GROUP \"group\"\n/end REF_GROUP"))
	})
}
