package Natives

import (
	"OLC2_Project1/server/interpreter/SymbolTable"
	arrayList "github.com/colegno/arraylist"
)

type Case struct {
	Expressions  *arrayList.List
	Instructions *arrayList.List
	Row          int
	Col          int
}

func NewCase(expressions *arrayList.List, instructions *arrayList.List, row int, col int) Case {
	return Case{
		Instructions: instructions,
		Expressions:  expressions,
		Row:          row,
		Col:          col,
	}
}

func (c Case) Execute(table SymbolTable.SymbolTable) interface{} {
	return nil
}
