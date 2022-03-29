package Expression

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
)

type Primitive struct {
	Value interface{}
	Type  SymbolTable.DataType
	Row   int
	Col   int
}

func (p Primitive) Compile(symbolTable SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {
	if p.Type == SymbolTable.INTEGER || p.Type == SymbolTable.FLOAT {
		return Abstract.NewValue(p.Value, p.Type, false, "")
	}

	return nil
}

func (p Primitive) Execute(symbolTable SymbolTable.SymbolTable) interface{} {
	res := p.GetValue(symbolTable)

	return res
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
