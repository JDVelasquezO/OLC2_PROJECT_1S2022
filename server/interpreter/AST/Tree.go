package AST

import (
	arrayList "github.com/colegno/arraylist"
)

type Tree struct {
	ListInstr     *arrayList.List
	ListFunctions *arrayList.List
	ListStructs   *arrayList.List
	ListArrays    *arrayList.List
}

func NewTree(list *arrayList.List) Tree {
	tree := Tree{
		ListInstr:     list,
		ListFunctions: nil,
		ListStructs:   nil,
		ListArrays:    nil,
	}
	return tree
}
