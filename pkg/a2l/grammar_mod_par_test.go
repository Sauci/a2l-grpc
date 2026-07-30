package a2l

// MOD_PAR (chapter 6.3.91) and the keywords it contains.

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrammar_MOD_PAR(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		modPar, ok := parseModPar(t, "")
		if !ok {
			return
		}

		equalNode(t, &ModParType{Comment: strVal("")}, modPar)
	})

	t.Run("optional/single occurrence keywords", func(t *testing.T) {
		modPar, ok := parseModPar(t, `VERSION "version"
EPK "epk"
SUPPLIER "supplier"
CUSTOMER "customer"
CUSTOMER_NO "customer number"
USER "user"
PHONE_NO "phone number"
ECU "control unit"
CPU_TYPE "cpu"
NO_OF_INTERFACES 2
ECU_CALIBRATION_OFFSET 0x1000`)
		if !ok {
			return
		}

		equalNode(t, &ModParType{
			Comment:                strVal(""),
			VERSION:                &VersionType{VersionIdentifier: strVal("version")},
			EPK:                    &EpkType{Identifier: strVal("epk")},
			SUPPLIER:               &SupplierType{Manufacturer: strVal("supplier")},
			CUSTOMER:               &CustomerType{Customer: strVal("customer")},
			CUSTOMER_NO:            &CustomerNoType{Number: strVal("customer number")},
			USER:                   &UserType{UserName: strVal("user")},
			PHONE_NO:               &PhoneNoType{TelNum: strVal("phone number")},
			ECU:                    &EcuType{ControlUnit: strVal("control unit")},
			CPU_TYPE:               &CpuTypeType{Cpu: strVal("cpu")},
			NO_OF_INTERFACES:       &NoOfInterfacesType{Num: intVal("2")},
			ECU_CALIBRATION_OFFSET: &EcuCalibrationOffsetType{Offset: longVal("0x1000")},
		}, modPar)
	})

	t.Run("list/several ADDR_EPK", func(t *testing.T) {
		modPar, ok := parseModPar(t, "ADDR_EPK 0x1000\nADDR_EPK 0x2000")
		if !ok {
			return
		}

		equalNodes(t, []*AddrEpkType{
			{Address: longVal("0x1000")},
			{Address: longVal("0x2000")},
		}, modPar.ADDR_EPK)
	})

	t.Run("list/several SYSTEM_CONSTANT", func(t *testing.T) {
		modPar, ok := parseModPar(t, "SYSTEM_CONSTANT \"first\" \"1\"\nSYSTEM_CONSTANT \"second\" \"2\"")
		if !ok {
			return
		}

		equalNodes(t, []*SystemConstantType{
			{Name: strVal("first"), Value: strVal("1")},
			{Name: strVal("second"), Value: strVal("2")},
		}, modPar.SYSTEM_CONSTANT)
	})

	t.Run("reject/missing comment", func(t *testing.T) {
		parseFails(t, moduleScope("/begin MOD_PAR\n/end MOD_PAR"))
	})

	// Deviation: EPK is declared once by the specification.
	t.Run("reject/repeated EPK", func(t *testing.T) {
		deviation(t, "optional keywords may be repeated", func(t assert.TestingT) {
			parseFails(t, modParScope("EPK \"first\"\nEPK \"second\""))
		})
	})
}

func TestGrammar_ADDR_EPK(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		modPar, ok := parseModPar(t, "ADDR_EPK 0x12345678")
		if !ok {
			return
		}

		equalNodes(t, []*AddrEpkType{{Address: longVal("0x12345678")}}, modPar.ADDR_EPK)
	})

	t.Run("reject/missing address", func(t *testing.T) {
		parseFails(t, modParScope("ADDR_EPK"))
	})

	t.Run("reject/string address", func(t *testing.T) {
		parseFails(t, modParScope("ADDR_EPK \"0x1000\""))
	})
}

func TestGrammar_EPK(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		modPar, ok := parseModPar(t, "EPK \"EPROM identifier\"")
		if !ok {
			return
		}

		equalNode(t, &EpkType{Identifier: strVal("EPROM identifier")}, modPar.EPK)
	})

	t.Run("reject/identifier parameter", func(t *testing.T) {
		parseFails(t, modParScope("EPK identifier"))
	})
}

func TestGrammar_SUPPLIER(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		modPar, ok := parseModPar(t, "SUPPLIER \"manufacturer\"")
		if !ok {
			return
		}

		equalNode(t, &SupplierType{Manufacturer: strVal("manufacturer")}, modPar.SUPPLIER)
	})

	t.Run("reject/missing manufacturer", func(t *testing.T) {
		parseFails(t, modParScope("SUPPLIER"))
	})
}

func TestGrammar_CUSTOMER(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		modPar, ok := parseModPar(t, "CUSTOMER \"customer\"")
		if !ok {
			return
		}

		equalNode(t, &CustomerType{Customer: strVal("customer")}, modPar.CUSTOMER)
	})

	t.Run("reject/missing customer", func(t *testing.T) {
		parseFails(t, modParScope("CUSTOMER"))
	})
}

func TestGrammar_CUSTOMER_NO(t *testing.T) {
	// The customer number is a string, not a number.
	t.Run("mandatory parameters", func(t *testing.T) {
		modPar, ok := parseModPar(t, "CUSTOMER_NO \"1234\"")
		if !ok {
			return
		}

		equalNode(t, &CustomerNoType{Number: strVal("1234")}, modPar.CUSTOMER_NO)
	})

	t.Run("reject/numeric parameter", func(t *testing.T) {
		parseFails(t, modParScope("CUSTOMER_NO 1234"))
	})
}

func TestGrammar_USER(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		modPar, ok := parseModPar(t, "USER \"user name\"")
		if !ok {
			return
		}

		equalNode(t, &UserType{UserName: strVal("user name")}, modPar.USER)
	})

	t.Run("reject/missing user name", func(t *testing.T) {
		parseFails(t, modParScope("USER"))
	})
}

func TestGrammar_PHONE_NO(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		modPar, ok := parseModPar(t, "PHONE_NO \"+41 00 000 00 00\"")
		if !ok {
			return
		}

		equalNode(t, &PhoneNoType{TelNum: strVal("+41 00 000 00 00")}, modPar.PHONE_NO)
	})

	t.Run("reject/numeric parameter", func(t *testing.T) {
		parseFails(t, modParScope("PHONE_NO 123456"))
	})
}

func TestGrammar_ECU(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		modPar, ok := parseModPar(t, "ECU \"control unit\"")
		if !ok {
			return
		}

		equalNode(t, &EcuType{ControlUnit: strVal("control unit")}, modPar.ECU)
	})

	t.Run("reject/missing control unit", func(t *testing.T) {
		parseFails(t, modParScope("ECU"))
	})
}

func TestGrammar_CPU_TYPE(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		modPar, ok := parseModPar(t, "CPU_TYPE \"INTEL 4711\"")
		if !ok {
			return
		}

		equalNode(t, &CpuTypeType{Cpu: strVal("INTEL 4711")}, modPar.CPU_TYPE)
	})

	t.Run("reject/identifier parameter", func(t *testing.T) {
		parseFails(t, modParScope("CPU_TYPE INTEL"))
	})
}

func TestGrammar_NO_OF_INTERFACES(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		modPar, ok := parseModPar(t, "NO_OF_INTERFACES 2")
		if !ok {
			return
		}

		equalNode(t, &NoOfInterfacesType{Num: intVal("2")}, modPar.NO_OF_INTERFACES)
	})

	t.Run("reject/float number of interfaces", func(t *testing.T) {
		parseFails(t, modParScope("NO_OF_INTERFACES 2.0"))
	})
}

func TestGrammar_ECU_CALIBRATION_OFFSET(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		modPar, ok := parseModPar(t, "ECU_CALIBRATION_OFFSET 0x1000")
		if !ok {
			return
		}

		equalNode(t, &EcuCalibrationOffsetType{Offset: longVal("0x1000")}, modPar.ECU_CALIBRATION_OFFSET)
	})

	t.Run("accept/negative offset", func(t *testing.T) {
		modPar, ok := parseModPar(t, "ECU_CALIBRATION_OFFSET -16")
		if !ok {
			return
		}

		equalNode(t, &EcuCalibrationOffsetType{Offset: longVal("-16")}, modPar.ECU_CALIBRATION_OFFSET)
	})

	t.Run("reject/missing offset", func(t *testing.T) {
		parseFails(t, modParScope("ECU_CALIBRATION_OFFSET"))
	})
}

func TestGrammar_SYSTEM_CONSTANT(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		modPar, ok := parseModPar(t, "SYSTEM_CONSTANT \"constant\" \"value\"")
		if !ok {
			return
		}

		equalNodes(t, []*SystemConstantType{{Name: strVal("constant"), Value: strVal("value")}},
			modPar.SYSTEM_CONSTANT)
	})

	// The value of a system constant is a string, even when it holds a number.
	t.Run("reject/numeric value", func(t *testing.T) {
		parseFails(t, modParScope("SYSTEM_CONSTANT \"constant\" 1"))
	})

	t.Run("reject/missing value", func(t *testing.T) {
		parseFails(t, modParScope("SYSTEM_CONSTANT \"constant\""))
	})
}

func TestGrammar_CALIBRATION_METHOD(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		modPar, ok := parseModPar(t, "/begin CALIBRATION_METHOD \"InCircuit\" 2\n/end CALIBRATION_METHOD")
		if !ok {
			return
		}

		equalNodes(t, []*CalibrationMethodType{{
			Method:  strVal("InCircuit"),
			Version: longVal("2"),
		}}, modPar.CALIBRATION_METHOD)
	})

	t.Run("optional/CALIBRATION_HANDLE", func(t *testing.T) {
		modPar, ok := parseModPar(t, `/begin CALIBRATION_METHOD "InCircuit" 2
/begin CALIBRATION_HANDLE 0x10000 0x20000
/end CALIBRATION_HANDLE
/end CALIBRATION_METHOD`)
		if !ok {
			return
		}

		equalNodes(t, []*CalibrationMethodType{{
			Method:  strVal("InCircuit"),
			Version: longVal("2"),
			CALIBRATION_HANDLE: []*CalibrationHandleType{
				{Handle: []*LongType{longVal("0x10000"), longVal("0x20000")}},
			},
		}}, modPar.CALIBRATION_METHOD)
	})

	t.Run("list/several CALIBRATION_HANDLE", func(t *testing.T) {
		modPar, ok := parseModPar(t, `/begin CALIBRATION_METHOD "InCircuit" 2
/begin CALIBRATION_HANDLE 0x10000
/end CALIBRATION_HANDLE
/begin CALIBRATION_HANDLE 0x20000
/end CALIBRATION_HANDLE
/end CALIBRATION_METHOD`)
		if !ok || !assert.Len(t, modPar.CALIBRATION_METHOD, 1) {
			return
		}

		equalNodes(t, []*CalibrationHandleType{
			{Handle: []*LongType{longVal("0x10000")}},
			{Handle: []*LongType{longVal("0x20000")}},
		}, modPar.CALIBRATION_METHOD[0].CALIBRATION_HANDLE)
	})

	t.Run("reject/missing version", func(t *testing.T) {
		parseFails(t, modParScope("/begin CALIBRATION_METHOD \"InCircuit\"\n/end CALIBRATION_METHOD"))
	})

	t.Run("reject/identifier as method", func(t *testing.T) {
		parseFails(t, modParScope("/begin CALIBRATION_METHOD InCircuit 2\n/end CALIBRATION_METHOD"))
	})
}

func TestGrammar_CALIBRATION_HANDLE(t *testing.T) {
	type testCaseType struct {
		name    string
		content string
		handles []*LongType
	}

	for _, testCase := range []testCaseType{
		{name: "list/no handle", content: "", handles: nil},
		{name: "list/single handle", content: "0x10000", handles: []*LongType{longVal("0x10000")}},
		{
			name:    "list/several handles",
			content: "0x10000 0x20000 0x30000",
			handles: []*LongType{longVal("0x10000"), longVal("0x20000"), longVal("0x30000")},
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			modPar, ok := parseModPar(t, "/begin CALIBRATION_METHOD \"InCircuit\" 2\n"+
				"/begin CALIBRATION_HANDLE "+testCase.content+"\n/end CALIBRATION_HANDLE\n/end CALIBRATION_METHOD")
			if !ok || !assert.Len(t, modPar.CALIBRATION_METHOD, 1) {
				return
			}

			equalNodes(t, []*CalibrationHandleType{{Handle: testCase.handles}},
				modPar.CALIBRATION_METHOD[0].CALIBRATION_HANDLE)
		})
	}

	t.Run("reject/string handle", func(t *testing.T) {
		parseFails(t, modParScope("/begin CALIBRATION_METHOD \"InCircuit\" 2\n"+
			"/begin CALIBRATION_HANDLE \"0x10000\"\n/end CALIBRATION_HANDLE\n/end CALIBRATION_METHOD"))
	})
}

func TestGrammar_MEMORY_LAYOUT(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		modPar, ok := parseModPar(t,
			"/begin MEMORY_LAYOUT PRG_CODE 0x1000 0x200 -1 -1 -1 -1 -1\n/end MEMORY_LAYOUT")
		if !ok {
			return
		}

		equalNodes(t, []*MemoryLayoutType{{
			PrgType: "PRG_CODE",
			Address: longVal("0x1000"),
			Size:    longVal("0x200"),
			Offset: []*LongType{
				longVal("-1"), longVal("-1"), longVal("-1"), longVal("-1"), longVal("-1"),
			},
		}}, modPar.MEMORY_LAYOUT)
	})

	for _, prgType := range []string{"PRG_CODE", "PRG_DATA", "PRG_RESERVED"} {
		t.Run("enum/"+prgType, func(t *testing.T) {
			modPar, ok := parseModPar(t,
				"/begin MEMORY_LAYOUT "+prgType+" 0x0 0x0 0 0 0 0 0\n/end MEMORY_LAYOUT")
			if !ok || !assert.Len(t, modPar.MEMORY_LAYOUT, 1) {
				return
			}

			assert.Equal(t, prgType, modPar.MEMORY_LAYOUT[0].PrgType)
		})
	}

	t.Run("optional/IF_DATA", func(t *testing.T) {
		modPar, ok := parseModPar(t, "/begin MEMORY_LAYOUT PRG_CODE 0x0 0x0 0 0 0 0 0\n"+
			"/begin IF_DATA XCP\n/end IF_DATA\n/end MEMORY_LAYOUT")
		if !ok || !assert.Len(t, modPar.MEMORY_LAYOUT, 1) {
			return
		}

		equalNodes(t, []*IfDataType{{Name: identVal("XCP")}}, modPar.MEMORY_LAYOUT[0].IF_DATA)
	})

	// The specification declares exactly five offsets.
	t.Run("reject/four offsets", func(t *testing.T) {
		parseFails(t, modParScope("/begin MEMORY_LAYOUT PRG_CODE 0x0 0x0 0 0 0 0\n/end MEMORY_LAYOUT"))
	})

	t.Run("reject/six offsets", func(t *testing.T) {
		parseFails(t, modParScope("/begin MEMORY_LAYOUT PRG_CODE 0x0 0x0 0 0 0 0 0 0\n/end MEMORY_LAYOUT"))
	})

	t.Run("reject/unknown program type", func(t *testing.T) {
		parseFails(t, modParScope("/begin MEMORY_LAYOUT PRG_UNKNOWN 0x0 0x0 0 0 0 0 0\n/end MEMORY_LAYOUT"))
	})
}

func TestGrammar_MEMORY_SEGMENT(t *testing.T) {
	t.Run("mandatory parameters", func(t *testing.T) {
		modPar, ok := parseModPar(t, "/begin MEMORY_SEGMENT memory_segment \"long identifier\" "+
			"DATA FLASH INTERN 0x1000 0x200 -1 -1 -1 -1 -1\n/end MEMORY_SEGMENT")
		if !ok {
			return
		}

		equalNodes(t, []*MemorySegmentType{{
			Name:           identVal("memory_segment"),
			LongIdentifier: strVal("long identifier"),
			PrgType:        "DATA",
			MemoryType:     "FLASH",
			Attribute:      "INTERN",
			Address:        longVal("0x1000"),
			Size:           longVal("0x200"),
			Offset: []*LongType{
				longVal("-1"), longVal("-1"), longVal("-1"), longVal("-1"), longVal("-1"),
			},
		}}, modPar.MEMORY_SEGMENT)
	})

	for _, prgType := range []string{
		"CODE", "DATA", "OFFLINE_DATA", "VARIABLES", "SERAM", "RESERVED",
		"CALIBRATION_VARIABLES", "EXCLUDE_FROM_FLASH",
	} {
		t.Run("enum/program type "+prgType, func(t *testing.T) {
			modPar, ok := parseModPar(t, "/begin MEMORY_SEGMENT memory_segment \"\" "+prgType+
				" RAM INTERN 0x0 0x0 0 0 0 0 0\n/end MEMORY_SEGMENT")
			if !ok || !assert.Len(t, modPar.MEMORY_SEGMENT, 1) {
				return
			}

			assert.Equal(t, prgType, modPar.MEMORY_SEGMENT[0].PrgType)
		})
	}

	for _, memoryType := range []string{"RAM", "EEPROM", "EPROM", "ROM", "REGISTER", "FLASH"} {
		t.Run("enum/memory type "+memoryType, func(t *testing.T) {
			modPar, ok := parseModPar(t, "/begin MEMORY_SEGMENT memory_segment \"\" DATA "+memoryType+
				" INTERN 0x0 0x0 0 0 0 0 0\n/end MEMORY_SEGMENT")
			if !ok || !assert.Len(t, modPar.MEMORY_SEGMENT, 1) {
				return
			}

			assert.Equal(t, memoryType, modPar.MEMORY_SEGMENT[0].MemoryType)
		})
	}

	for _, attribute := range []string{"INTERN", "EXTERN"} {
		t.Run("enum/attribute "+attribute, func(t *testing.T) {
			modPar, ok := parseModPar(t, "/begin MEMORY_SEGMENT memory_segment \"\" DATA RAM "+attribute+
				" 0x0 0x0 0 0 0 0 0\n/end MEMORY_SEGMENT")
			if !ok || !assert.Len(t, modPar.MEMORY_SEGMENT, 1) {
				return
			}

			assert.Equal(t, attribute, modPar.MEMORY_SEGMENT[0].Attribute)
		})
	}

	t.Run("optional/IF_DATA", func(t *testing.T) {
		modPar, ok := parseModPar(t, "/begin MEMORY_SEGMENT memory_segment \"\" DATA RAM INTERN 0x0 0x0 0 0 0 0 0\n"+
			"/begin IF_DATA XCP\n/end IF_DATA\n/end MEMORY_SEGMENT")
		if !ok || !assert.Len(t, modPar.MEMORY_SEGMENT, 1) {
			return
		}

		equalNodes(t, []*IfDataType{{Name: identVal("XCP")}}, modPar.MEMORY_SEGMENT[0].IF_DATA)
	})

	t.Run("reject/missing attribute", func(t *testing.T) {
		parseFails(t, modParScope(
			"/begin MEMORY_SEGMENT memory_segment \"\" DATA RAM 0x0 0x0 0 0 0 0 0\n/end MEMORY_SEGMENT"))
	})

	t.Run("reject/unknown memory type", func(t *testing.T) {
		parseFails(t, modParScope(
			"/begin MEMORY_SEGMENT memory_segment \"\" DATA MRAM INTERN 0x0 0x0 0 0 0 0 0\n/end MEMORY_SEGMENT"))
	})
}
