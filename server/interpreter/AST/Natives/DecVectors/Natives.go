package DecVectors

import (
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
)

type Natives struct {
	Id        string
	Value     Abstract.Expression
	Operation string
}

func (n Natives) GetValue(symbolTable SymbolTable.SymbolTable) SymbolTable.ReturnType {
	//TODO implement me
	panic("implement me")
}

func NewOperation(id string, value Abstract.Expression, operation string) Natives {
	return Natives{
		Value:     value,
		Id:        id,
		Operation: operation,
	}
}

func (n Natives) Execute(symbolTable SymbolTable.SymbolTable) interface{} {
	switch n.Operation {
	case "push":
		newPush := NewPush(n.Id, n.Value)
		newPush.ExecuteFirstTime(symbolTable)
	case "remove":
		newRemove := NewRemove(n.Id, n.Value)
		newRemove.Execute(symbolTable)
	}
	return nil
}
