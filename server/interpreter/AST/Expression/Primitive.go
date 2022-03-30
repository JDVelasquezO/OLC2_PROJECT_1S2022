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
	if p.Type == SymbolTable.INTEGER || p.Type == SymbolTable.FLOAT || p.Type == SymbolTable.CHAR {
		return Abstract.NewValue(p.Value, p.Type, false, "")
	} else if p.Type == SymbolTable.BOOLEAN {
		res := Abstract.NewValue(p.Value, p.Type, false, "")
		newValue := CheckLabels(generator, res)

		if newValue.Value.(bool) {
			generator.AddGoTo(newValue.TrueLabel)
			generator.AddGoTo(newValue.FalseLabel)
		} else {
			generator.AddGoTo(newValue.FalseLabel)
			generator.AddGoTo(newValue.TrueLabel)
		}

		res.TrueLabel = newValue.TrueLabel
		res.FalseLabel = newValue.FalseLabel

		return res
	}

	return nil
}

func CheckLabels(generator *Generator.Generator, value Abstract.Value) Abstract.Value {
	if value.TrueLabel == "" {
		value.TrueLabel = generator.NewLabel()
	}

	if value.FalseLabel == "" {
		value.FalseLabel = generator.NewLabel()
	}

	return value
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
