package DecVectors

import (
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment/Vector"
)

type Len struct {
	Id string
}

func NewLen(id string) Len {
	return Len{
		Id: id,
	}
}

func (l Len) GetValue(symbolTable SymbolTable.SymbolTable) SymbolTable.ReturnType {
	len_ := l.Execute(symbolTable)
	return SymbolTable.ReturnType{
		Value: len_,
		Type:  SymbolTable.INTEGER,
	}
}

func (l Len) Execute(symbolTable SymbolTable.SymbolTable) interface{} {
	val := symbolTable.GetSymbolArray(l.Id)
	newVal := val.(Vector.Vector).GetLength()
	return newVal
}
