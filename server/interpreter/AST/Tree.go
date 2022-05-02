package AST

import (
	arrayList "github.com/colegno/arraylist"
)

type Tree struct {
	ListInstr     *arrayList.List
	ListFunctions *arrayList.List
	ListStructs   *arrayList.List
	ListArrays    *arrayList.List
	ListTemps     *arrayList.List
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

func NewTreeOptimizer(list *arrayList.List, listTemps *arrayList.List) Tree {
	tree := Tree{
		ListInstr: list,
		ListTemps: listTemps,
	}
	return tree
}
