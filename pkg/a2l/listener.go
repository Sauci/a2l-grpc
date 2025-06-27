package a2l

import "github.com/ngamber/a2l-grpc/pkg/a2l/parser"

type Listener struct {
	*parser.BaseA2LListener
	root  RootNodeType
	stack []IIndentedNode
}

func (n *Listener) Push(node IIndentedNode) {
	n.stack = append(n.stack, node)
}

func (n *Listener) Last() (node IIndentedNode) {
	return n.stack[len(n.stack)-1]
}

func (n *Listener) Pop() (node IIndentedNode) {
	node = n.stack[len(n.stack)-1]
	n.stack = n.stack[:len(n.stack)-1]

	return node
}

func (n *Listener) Tree() *RootNodeType {
	return &n.root
}

func NewListener() *Listener {
	listener := &Listener{stack: make([]IIndentedNode, 0)}

	listener.Push(&listener.root)

	return listener
}
