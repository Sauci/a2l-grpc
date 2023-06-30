package a2l

import "fmt"

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

func (n *FloatType) A2LString() string {
	return fmt.Sprintf("%v", n.Value)
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
