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
	value := n.Execute(symbolTable)
	return SymbolTable.ReturnType{
		Value: value,
		Type:  SymbolTable.INTEGER,
	}
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
		return newPush.Execute(symbolTable)
	case "remove":
		newRemove := NewRemove(n.Id, n.Value)
		return newRemove.Execute(symbolTable)
	case "len", "capacity":
		newLen := NewLen(n.Id)
		return newLen.Execute(symbolTable)
	}
	return nil
}
