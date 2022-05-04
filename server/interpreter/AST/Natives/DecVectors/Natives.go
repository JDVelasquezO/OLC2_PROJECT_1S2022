package DecVectors

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment/Vector"
	"fmt"
	"reflect"
)

type Natives struct {
	Id        string
	Value     Abstract.Expression
	Operation string
	Index     Abstract.Expression
}

func (n Natives) Compile(symbolTable *SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {
	id := symbolTable.GetSymbol(n.Id)
	if n.Operation == "abs" {
		if id.(*SymbolTable.Symbol).DataType == SymbolTable.INTEGER {
			if id.(*SymbolTable.Symbol).Value.(int) < 0 {
				temp := generator.AddTemp()
				generator.AddExpression(temp, fmt.Sprintf("%v", id.(*SymbolTable.Symbol).Value), "-1", "*")
				val := Abstract.NewValue(temp, SymbolTable.INTEGER, true, "")
				return val
			}
		} else {
			if id.(*SymbolTable.Symbol).Value.(float64) < 0 {
				temp := generator.AddTemp()
				generator.AddExpression(temp, fmt.Sprintf("%v", id.(*SymbolTable.Symbol).Value), "-1", "*")
				val := Abstract.NewValue(temp, SymbolTable.FLOAT, true, "")
				return val
			}
		}

	}
	val := Abstract.NewValue(id.(*SymbolTable.Symbol).Value, SymbolTable.INTEGER, false, "")
	return val
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
	case "abs":
		symbol := symbolTable.GetSymbol(n.Id)
		if symbol.(*SymbolTable.Symbol).DataType == SymbolTable.INTEGER {
			if symbol.(*SymbolTable.Symbol).Value.(int) < 0 {
				return symbol.(*SymbolTable.Symbol).Value.(int) * -1
			}
			return symbol.(*SymbolTable.Symbol).Value.(int)
		} else {
			if symbol.(*SymbolTable.Symbol).Value.(float64) < 0 {
				return symbol.(*SymbolTable.Symbol).Value.(float64) * -1.0
			}
			return symbol.(*SymbolTable.Symbol).Value.(float64)
		}
	}
	return nil
}
