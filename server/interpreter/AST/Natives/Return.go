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

func (r Return) Compile(symbolTable SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {
	//TODO implement me
	panic("implement me")
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
