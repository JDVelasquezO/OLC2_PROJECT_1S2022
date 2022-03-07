package AST

import (
	arrayList "github.com/colegno/arraylist"
)

type Tree struct {
	ListInstr     *arrayList.List
	ListFunctions *arrayList.List
}

func NewTree(list *arrayList.List) Tree {
	tree := Tree{
		ListInstr:     list,
		ListFunctions: nil,
	}
	return tree
}
