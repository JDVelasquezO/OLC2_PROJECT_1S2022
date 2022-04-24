package Natives

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
)

type Return struct {
	Row        int
	Col        int
	Expression Abstract.Expression
	Type       SymbolTable.DataType
	Result     interface{}
}

func (r Return) Compile(symbolTable *SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {
	value := r.Expression.Compile(symbolTable, generator)

	if value.(Abstract.Value).Type == SymbolTable.BOOLEAN {
		tempLabel := generator.NewLabel()

		generator.SetLabel(value.(Abstract.Value).TrueLabel)
		generator.SetStack("P", "1", true)
		generator.AddGoTo(tempLabel)

		generator.SetLabel(value.(Abstract.Value).FalseLabel)
		generator.SetStack("P", "0", true)
		generator.SetLabel(tempLabel)
	} else {
		generator.SetStack("P", value.(Abstract.Value).Value, true)
	}

	generator.AddGoTo(symbolTable.ReturnLabel)
	return nil
}

func NewReturn(row int, col int, expression Abstract.Expression) Return {
	return Return{
		Row:        row,
		Col:        col,
		Expression: expression,
	}
}

func (r Return) Execute(symbolTable SymbolTable.SymbolTable) interface{} {
	if r.Expression == nil {
		r.Type = SymbolTable.VOID
		return r
	}

	value := r.Expression.GetValue(symbolTable)
	dataType := value.Type
	result := value.Value

	r.Type = dataType
	r.Result = result

	return SymbolTable.ReturnType{
		Type:  r.Type,
		Value: r.Result,
	}
}
