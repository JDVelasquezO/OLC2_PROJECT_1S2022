package AST

import (
	"OLC2_Project1/server/interpreter/SymbolTable/Environment"
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

func (t *Tree) AddFunction(nameFunc Environment.Function) {
	t.ListFunctions.Add(nameFunc)
}
