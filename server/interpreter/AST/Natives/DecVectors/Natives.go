package DecVectors

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment/Vector"
	"reflect"
)

type Natives struct {
	Id        string
	Value     Abstract.Expression
	Operation string
	Index     Abstract.Expression
}

func (n Natives) Compile(symbolTable SymbolTable.SymbolTable, generator Generator.Generator) interface{} {
	//TODO implement me
	panic("implement me")
}

func (n Natives) GetValue(symbolTable SymbolTable.SymbolTable) SymbolTable.ReturnType {
	value := n.Execute(symbolTable)
	if reflect.TypeOf(value).String() == "bool" {
		return SymbolTable.ReturnType{
			Value: value,
			Type:  SymbolTable.BOOLEAN,
		}
	} else {
		return SymbolTable.ReturnType{
			Value: value,
			Type:  SymbolTable.INTEGER,
		}
	}
}

func NewOperation(id string, value Abstract.Expression, operation string, index Abstract.Expression) Natives {
	return Natives{
		Value:     value,
		Id:        id,
		Operation: operation,
		Index:     index,
	}
}

func (n Natives) Execute(symbolTable SymbolTable.SymbolTable) interface{} {
	switch n.Operation {
	case "push":
		newPush := NewPush(n.Id, n.Value)
		newPush.Execute(symbolTable)
	case "remove":
		newRemove := NewRemove(n.Id, n.Value)
		newRemove.Execute(symbolTable)
	case "len", "capacity":
		newLen := NewLen(n.Id)
		return newLen.Execute(symbolTable)
	case "insert":
		index := n.Index.GetValue(symbolTable).Value.(int)
		newInsert := NewInsert(n.Id, n.Value, index)
		newInsert.Execute(symbolTable)
	case "contains":
		val := symbolTable.GetSymbolArray(n.Id)
		return val.(Vector.Vector).Contains(n.Value, symbolTable)
	}
	return nil
}
