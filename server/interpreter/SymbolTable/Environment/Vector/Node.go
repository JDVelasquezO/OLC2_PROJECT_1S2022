package Vector

import "OLC2_Project1/server/interpreter/Abstract"

type Node struct {
	Value Abstract.Expression
	Next  *Node
}

func NewNode(value Abstract.Expression) Node {
	return Node{
		Value: value,
	}
}
