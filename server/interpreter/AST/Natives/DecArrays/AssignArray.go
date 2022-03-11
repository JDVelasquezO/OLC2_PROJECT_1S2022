package DecArrays

import (
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	arrayList "github.com/colegno/arraylist"
)

type AssignArray struct {
	Val       Abstract.Expression
	Id        string
	ListIndex *arrayList.List
}

func (a AssignArray) Execute(symbolTable SymbolTable.SymbolTable) interface{} {
	//TODO implement me
	panic("implement me")
}

func NewAssignArray(id string, listInArray *arrayList.List, val Abstract.Expression) *AssignArray {
	return &AssignArray{
		Val:       val,
		Id:        id,
		ListIndex: listInArray,
	}
}
