package a2l

import (
	"fmt"
	"strconv"
)

type A2LStringer interface {
	A2LString() string
}

func (n *IntType) A2LString() string {
	var format string

	if n.Base == 10 {
		format = fmt.Sprintf("%%0%vd", n.Size)
	} else {
		format = fmt.Sprintf("0x%%0%vX", n.Size)
	}

	return fmt.Sprintf(format, n.Value)
}

// A2LString returns the value as written in the parsed file. A value which was not created by the
// parser carries no source form and is formatted from its numeric value instead.
func (n *FloatType) A2LString() string {
	var format string

	if n.Source != "" {
		format = n.Source
	} else {
		format = strconv.FormatFloat(n.Value, 'f', -1, 64)
	}

	return format
}

func (n *LongType) A2LString() string {
	var format string

	if n.Base == 10 {
		format = fmt.Sprintf("%%0%vd", n.Size)
	} else {
		format = fmt.Sprintf("0x%%0%vX", n.Size)
	}

	return fmt.Sprintf(format, n.Value)
}

func (n *StringType) A2LString() string {
	return fmt.Sprintf("\"%v\"", n.Value)
}

func (n *IdentType) A2LString() string {
	return fmt.Sprintf("%v", n.Value)
}

func (n *TagType) A2LString() string {
	return fmt.Sprintf("\"%v\"", n.Value)
}

func (n *DataTypeType) A2LString() string {
	return fmt.Sprintf("%v", n.Value)
}

func (n *IndexOrderType) A2LString() string {
	return fmt.Sprintf("%v", n.Value)
}

func (n *AddrTypeType) A2LString() string {
	return fmt.Sprintf("%v", n.Value)
}
