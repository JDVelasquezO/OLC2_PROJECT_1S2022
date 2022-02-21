package Expression

import (
	"OLC2_Project1/server/interpreter/SymbolTable"
)

type Primitive struct {
	Value interface{}
	Type  SymbolTable.DataType
	Row   int
	Col   int
}

func (p Primitive) GetValue(symbolTable SymbolTable.SymbolTable) SymbolTable.ReturnType {
	return SymbolTable.ReturnType{
		Type:  p.Type,
		Value: p.Value,
	}
}

func NewPrimitive(val interface{}, types SymbolTable.DataType, row int, col int) Primitive {
	e := Primitive{val, types, row, col}
	return e
}
