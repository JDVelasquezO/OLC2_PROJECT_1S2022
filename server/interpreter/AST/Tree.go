package AST

import arrayList "github.com/colegno/arraylist"

type Tree struct {
	ListInstr *arrayList.List
}

func NewTree(list *arrayList.List) Tree {
	tree := Tree{ListInstr: list}
	return tree
}
