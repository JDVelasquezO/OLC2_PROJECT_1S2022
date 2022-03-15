package Natives

import (
	"OLC2_Project1/server/interpreter/SymbolTable"
	arrayList "github.com/colegno/arraylist"
)

type Case struct {
	Expressions  *arrayList.List
	Instructions *arrayList.List
}

func NewCase(expressions *arrayList.List, instructions *arrayList.List) Case {
	return Case{
		Instructions: instructions,
		Expressions:  expressions,
	}
}

func (c Case) Execute(table SymbolTable.SymbolTable) interface{} {
	return nil
}
