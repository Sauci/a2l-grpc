package a2l

// Keywords describing a single property of an adjustable object, a measurement object or an axis.

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrammar_BIT_MASK(t *testing.T) {
	t.Run("mandatory parameters/in CHARACTERISTIC", func(t *testing.T) {
		characteristic, ok := parseCharacteristic(t, "BIT_MASK 0x40")
		if !ok {
			return
		}

		equalNode(t, &BitMaskType{Mask: longVal("0x40")}, characteristic.BIT_MASK)
	})

	t.Run("mandatory parameters/in MEASUREMENT", func(t *testing.T) {
		measurement, ok := parseMeasurement(t, "BIT_MASK 0x40")
		if !ok {
			return
		}

		equalNode(t, &BitMaskType{Mask: longVal("0x40")}, measurement.BIT_MASK)
	})

	t.Run("reject/missing mask", func(t *testing.T) {
		parseFails(t, characteristicScope("BIT_MASK"))
	})

	t.Run("reject/float mask", func(t *testing.T) {
		parseFails(t, characteristicScope("BIT_MASK 1.0"))
	})
}

func TestGrammar_CALIBRATION_ACCESS(t *testing.T) {
	for _, accessType := range []string{
		"CALIBRATION", "NO_CALIBRATION", "NOT_IN_MCD_SYSTEM", "OFFLINE_CALIBRATION",
	} {
		t.Run("enum/"+accessType, func(t *testing.T) {
			characteristic, ok := parseCharacteristic(t, "CALIBRATION_ACCESS "+accessType)
			if !ok {
				return
			}

			equalNode(t, &CalibrationAccessType{Type: accessType}, characteristic.CALIBRATION_ACCESS)
		})
	}

	t.Run("mandatory parameters/in AXIS_PTS", func(t *testing.T) {
		axisPts, ok := parseAxisPts(t, "CALIBRATION_ACCESS NO_CALIBRATION")
		if !ok {
			return
		}

		equalNode(t, &CalibrationAccessType{Type: "NO_CALIBRATION"}, axisPts.CALIBRATION_ACCESS)
	})

	t.Run("reject/unknown access type", func(t *testing.T) {
		parseFails(t, characteristicScope("CALIBRATION_ACCESS ONLINE_CALIBRATION"))
	})

	t.Run("reject/missing access type", func(t *testing.T) {
		parseFails(t, characteristicScope("CALIBRATION_ACCESS"))
	})
}

func TestGrammar_COMPARISON_QUANTITY(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		characteristic, ok := parseCharacteristic(t, "COMPARISON_QUANTITY measurement")
		if !ok {
			return
		}

		equalNode(t, &ComparisonQuantityType{Name: identVal("measurement")}, characteristic.COMPARISON_QUANTITY)
	})

	t.Run("reject/string parameter", func(t *testing.T) {
		parseFails(t, characteristicScope("COMPARISON_QUANTITY \"measurement\""))
	})
}

func TestGrammar_DEPENDENT_CHARACTERISTIC(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		characteristic, ok := parseCharacteristic(t,
			"/begin DEPENDENT_CHARACTERISTIC \"X1+X2\"\n/end DEPENDENT_CHARACTERISTIC")
		if !ok {
			return
		}

		equalNode(t, &DependentCharacteristicType{Formula: strVal("X1+X2")},
			characteristic.DEPENDENT_CHARACTERISTIC)
	})

	t.Run("list/several characteristics", func(t *testing.T) {
		characteristic, ok := parseCharacteristic(t,
			"/begin DEPENDENT_CHARACTERISTIC \"X1+X2\" first second\n/end DEPENDENT_CHARACTERISTIC")
		if !ok {
			return
		}

		equalNode(t, &DependentCharacteristicType{
			Formula:        strVal("X1+X2"),
			Characteristic: []*IdentType{identVal("first"), identVal("second")},
		}, characteristic.DEPENDENT_CHARACTERISTIC)
	})

	t.Run("reject/missing formula", func(t *testing.T) {
		parseFails(t, characteristicScope(
			"/begin DEPENDENT_CHARACTERISTIC\n/end DEPENDENT_CHARACTERISTIC"))
	})

	t.Run("reject/identifier as formula", func(t *testing.T) {
		parseFails(t, characteristicScope(
			"/begin DEPENDENT_CHARACTERISTIC formula first\n/end DEPENDENT_CHARACTERISTIC"))
	})
}

func TestGrammar_VIRTUAL_CHARACTERISTIC(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		characteristic, ok := parseCharacteristic(t,
			"/begin VIRTUAL_CHARACTERISTIC \"X1*2\"\n/end VIRTUAL_CHARACTERISTIC")
		if !ok {
			return
		}

		equalNode(t, &VirtualCharacteristicType{Formula: strVal("X1*2")},
			characteristic.VIRTUAL_CHARACTERISTIC)
	})

	t.Run("list/several characteristics", func(t *testing.T) {
		characteristic, ok := parseCharacteristic(t,
			"/begin VIRTUAL_CHARACTERISTIC \"X1*2\" first second\n/end VIRTUAL_CHARACTERISTIC")
		if !ok {
			return
		}

		equalNode(t, &VirtualCharacteristicType{
			Formula:        strVal("X1*2"),
			Characteristic: []*IdentType{identVal("first"), identVal("second")},
		}, characteristic.VIRTUAL_CHARACTERISTIC)
	})

	t.Run("reject/missing formula", func(t *testing.T) {
		parseFails(t, characteristicScope("/begin VIRTUAL_CHARACTERISTIC\n/end VIRTUAL_CHARACTERISTIC"))
	})
}

func TestGrammar_DISPLAY_IDENTIFIER(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		characteristic, ok := parseCharacteristic(t, "DISPLAY_IDENTIFIER display_name")
		if !ok {
			return
		}

		equalNode(t, &DisplayIdentifierType{DisplayName: identVal("display_name")},
			characteristic.DISPLAY_IDENTIFIER)
	})

	t.Run("reject/string parameter", func(t *testing.T) {
		parseFails(t, characteristicScope("DISPLAY_IDENTIFIER \"display name\""))
	})
}

func TestGrammar_ECU_ADDRESS_EXTENSION(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		characteristic, ok := parseCharacteristic(t, "ECU_ADDRESS_EXTENSION 2")
		if !ok {
			return
		}

		equalNode(t, &EcuAddressExtensionType{Extension: intVal("2")}, characteristic.ECU_ADDRESS_EXTENSION)
	})

	t.Run("reject/float extension", func(t *testing.T) {
		parseFails(t, characteristicScope("ECU_ADDRESS_EXTENSION 2.0"))
	})
}

func TestGrammar_EXTENDED_LIMITS(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		characteristic, ok := parseCharacteristic(t, "EXTENDED_LIMITS -100.5 100.5")
		if !ok {
			return
		}

		equalNode(t, &ExtendedLimitsType{
			LowerLimit: floatVal("-100.5"),
			UpperLimit: floatVal("100.5"),
		}, characteristic.EXTENDED_LIMITS)
	})

	t.Run("reject/missing upper limit", func(t *testing.T) {
		parseFails(t, characteristicScope("EXTENDED_LIMITS -100"))
	})
}

func TestGrammar_FORMAT(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		characteristic, ok := parseCharacteristic(t, "FORMAT \"%8.4\"")
		if !ok {
			return
		}

		equalNode(t, &FormatType{FormatString: strVal("%8.4")}, characteristic.FORMAT)
	})

	t.Run("reject/unquoted format", func(t *testing.T) {
		parseFails(t, characteristicScope("FORMAT %8.4"))
	})
}

func TestGrammar_GUARD_RAILS(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		characteristic, ok := parseCharacteristic(t, "GUARD_RAILS")
		if !ok {
			return
		}

		equalNode(t, &GuardRailsType{Present: true}, characteristic.GUARD_RAILS)
	})

	t.Run("mandatory parameters/in AXIS_PTS", func(t *testing.T) {
		axisPts, ok := parseAxisPts(t, "GUARD_RAILS")
		if !ok {
			return
		}

		equalNode(t, &GuardRailsType{Present: true}, axisPts.GUARD_RAILS)
	})

	t.Run("reject/with parameter", func(t *testing.T) {
		parseFails(t, characteristicScope("GUARD_RAILS 1"))
	})
}

func TestGrammar_MAP_LIST(t *testing.T) {
	testIdentifierListKeyword(t, "MAP_LIST", characteristicScope,
		func(identifiers []*IdentType) *MapListType {
			return &MapListType{Name: identifiers}
		},
		func(t assert.TestingT, body string) (*MapListType, bool) {
			characteristic, ok := parseCharacteristic(t, body)
			if !ok {
				return nil, false
			}

			return characteristic.MAP_LIST, assert.NotNil(t, characteristic.MAP_LIST,
				"MAP_LIST is missing from the tree")
		})
}

func TestGrammar_MAX_REFRESH(t *testing.T) {
	t.Run("mandatory parameters/in CHARACTERISTIC", func(t *testing.T) {
		characteristic, ok := parseCharacteristic(t, "MAX_REFRESH 3 15")
		if !ok {
			return
		}

		equalNode(t, &MaxRefreshType{ScalingUnit: intVal("3"), Rate: longVal("15")},
			characteristic.MAX_REFRESH)
	})

	t.Run("mandatory parameters/in MEASUREMENT", func(t *testing.T) {
		measurement, ok := parseMeasurement(t, "MAX_REFRESH 3 15")
		if !ok {
			return
		}

		equalNode(t, &MaxRefreshType{ScalingUnit: intVal("3"), Rate: longVal("15")}, measurement.MAX_REFRESH)
	})

	t.Run("reject/missing rate", func(t *testing.T) {
		parseFails(t, characteristicScope("MAX_REFRESH 3"))
	})

	t.Run("reject/float scaling unit", func(t *testing.T) {
		parseFails(t, characteristicScope("MAX_REFRESH 3.0 15"))
	})
}

func TestGrammar_NUMBER(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		characteristic, ok := parseCharacteristic(t, "NUMBER 16")
		if !ok {
			return
		}

		equalNode(t, &NumberType{Number: intVal("16")}, characteristic.NUMBER)
	})

	t.Run("reject/float number", func(t *testing.T) {
		parseFails(t, characteristicScope("NUMBER 16.0"))
	})

	// NUMBER is declared for CHARACTERISTIC only, a measurement object uses ARRAY_SIZE.
	t.Run("reject/in MEASUREMENT", func(t *testing.T) {
		parseFails(t, measurementScope("NUMBER 16"))
	})
}

func TestGrammar_READ_ONLY(t *testing.T) {
	t.Run("mandatory parameters/in CHARACTERISTIC", func(t *testing.T) {
		characteristic, ok := parseCharacteristic(t, "READ_ONLY")
		if !ok {
			return
		}

		equalNode(t, &ReadOnlyType{Present: true}, characteristic.READ_ONLY)
	})

	t.Run("mandatory parameters/in AXIS_DESCR", func(t *testing.T) {
		axisDescr, ok := parseAxisDescr(t, "READ_ONLY")
		if !ok {
			return
		}

		equalNode(t, &ReadOnlyType{Present: true}, axisDescr.READ_ONLY)
	})

	t.Run("mandatory parameters/in USER_RIGHTS", func(t *testing.T) {
		userRights, ok := parseUserRights(t, "READ_ONLY")
		if !ok {
			return
		}

		equalNode(t, &ReadOnlyType{Present: true}, userRights.READ_ONLY)
	})

	t.Run("reject/with parameter", func(t *testing.T) {
		parseFails(t, characteristicScope("READ_ONLY 1"))
	})
}

func TestGrammar_REF_MEMORY_SEGMENT(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		characteristic, ok := parseCharacteristic(t, "REF_MEMORY_SEGMENT memory_segment")
		if !ok {
			return
		}

		equalNode(t, &RefMemorySegmentType{Name: identVal("memory_segment")},
			characteristic.REF_MEMORY_SEGMENT)
	})

	t.Run("reject/string parameter", func(t *testing.T) {
		parseFails(t, characteristicScope("REF_MEMORY_SEGMENT \"memory_segment\""))
	})
}

// DISCRETE, PHYS_UNIT, STEP_SIZE and SYMBOL_LINK are not part of ASAP2 1.51, they belong to a
// later version of the standard. They are parsed into the tree, but the serializer does not write
// them back.
func TestGrammar_DISCRETE(t *testing.T) {
	t.Run("mandatory parameters/in CHARACTERISTIC", func(t *testing.T) {
		deviation(t, "CHARACTERISTIC serializer ignores DISCRETE, PHYS_UNIT and STEP_SIZE",
			func(t assert.TestingT) {
				characteristic, ok := parseCharacteristic(t, "DISCRETE")
				if !ok {
					return
				}

				equalNode(t, &DiscreteType{Present: true}, characteristic.DISCRETE)
			})
	})

	t.Run("reject/with parameter", func(t *testing.T) {
		parseFails(t, characteristicScope("DISCRETE 1"))
	})
}

func TestGrammar_PHYS_UNIT(t *testing.T) {
	t.Run("mandatory parameters/in CHARACTERISTIC", func(t *testing.T) {
		deviation(t, "CHARACTERISTIC serializer ignores DISCRETE, PHYS_UNIT and STEP_SIZE",
			func(t assert.TestingT) {
				characteristic, ok := parseCharacteristic(t, "PHYS_UNIT \"km/h\"")
				if !ok {
					return
				}

				equalNode(t, &PhysUnitType{Unit: strVal("km/h")}, characteristic.PHYS_UNIT)
			})
	})

	t.Run("reject/identifier parameter", func(t *testing.T) {
		parseFails(t, characteristicScope("PHYS_UNIT km_per_h"))
	})
}

func TestGrammar_STEP_SIZE(t *testing.T) {
	t.Run("mandatory parameters/in CHARACTERISTIC", func(t *testing.T) {
		deviation(t, "CHARACTERISTIC serializer ignores DISCRETE, PHYS_UNIT and STEP_SIZE",
			func(t assert.TestingT) {
				characteristic, ok := parseCharacteristic(t, "STEP_SIZE 0.25")
				if !ok {
					return
				}

				equalNode(t, &StepSizeType{StepSize: floatVal("0.25")}, characteristic.STEP_SIZE)
			})
	})

	t.Run("reject/missing step size", func(t *testing.T) {
		parseFails(t, characteristicScope("STEP_SIZE"))
	})
}

func TestGrammar_SYMBOL_LINK(t *testing.T) {
	t.Run("mandatory parameters/in CHARACTERISTIC", func(t *testing.T) {
		deviation(t, "CHARACTERISTIC serializer ignores SYMBOL_LINK", func(t assert.TestingT) {
			characteristic, ok := parseCharacteristic(t, "SYMBOL_LINK \"symbol\" 4")
			if !ok {
				return
			}

			equalNode(t, &SymbolLinkType{SymbolName: strVal("symbol"), Offset: longVal("4")},
				characteristic.SYMBOL_LINK)
		})
	})

	t.Run("reject/missing offset", func(t *testing.T) {
		parseFails(t, characteristicScope("SYMBOL_LINK \"symbol\""))
	})

	t.Run("reject/identifier as symbol name", func(t *testing.T) {
		parseFails(t, characteristicScope("SYMBOL_LINK symbol 4"))
	})
}
