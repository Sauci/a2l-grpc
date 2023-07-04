package a2l

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntType_A2LString(t *testing.T) {
	type testCaseType struct {
		input  *IntType
		output string
	}

	for _, tc := range []testCaseType{
		{
			input:  &IntType{Value: 1, Base: 10, Size: 1},
			output: "1",
		},
		{
			input:  &IntType{Value: 1, Base: 16, Size: 1},
			output: "0x1",
		},
	} {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.output, tc.input.A2LString())
		})
	}
}

func TestLongType_A2LString(t *testing.T) {
	type testCaseType struct {
		input  *LongType
		output string
	}

	for _, tc := range []testCaseType{
		{
			input:  &LongType{Value: 1, Base: 10, Size: 1},
			output: "1",
		},
		{
			input:  &LongType{Value: 1, Base: 16, Size: 1},
			output: "0x1",
		},
	} {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.output, tc.input.A2LString())
		})
	}
}
