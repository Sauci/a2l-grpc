package a2l

import "fmt"

type A2LStringer interface {
	A2LString() string
}

func (n *IntType) A2LString() string {
	return fmt.Sprintf("%v", n.Value)
}

func (n *FloatType) A2LString() string {
	return fmt.Sprintf("%v", n.Value)
}

func (n *LongType) A2LString() string {
	return fmt.Sprintf("%v", n.Value)
}

func (n *StringType) A2LString() string {
	return fmt.Sprintf("\"%v\"", n.Value)
}

func (n *IdentType) A2LString() string {
	return fmt.Sprintf("%v", n.Value)
}
