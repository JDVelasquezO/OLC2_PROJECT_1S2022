package DecVectors

import (
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment/Vector"
)

type Push struct {
	Id    string
	Value Abstract.Expression
}

func NewPush(id string, value Abstract.Expression) Push {
	return Push{
		Id:    id,
		Value: value,
	}
}

func (p Push) Execute(symbolTable SymbolTable.SymbolTable) interface{} {
	val := symbolTable.GetSymbol(p.Id)
	newVal := val.(Vector.Vector).Push(p.Value, symbolTable)
	symbolTable.ChangeVal(p.Id, newVal)
	return val
}
