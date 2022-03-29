package Natives

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
)

type Break struct {
	Row        int
	Col        int
	Expression Abstract.Expression
}

func (b Break) Compile(symbolTable SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {
	return nil
}

func (b Break) Execute(symbolTable SymbolTable.SymbolTable) interface{} {
	if b.Expression == nil {
		return b
	}

	value := b.Expression.GetValue(symbolTable)

	return SymbolTable.ReturnType{
		Type:  value.Type,
		Value: value.Value,
	}
}

func NewBreak(row int, col int, expression Abstract.Expression) Break {
	return Break{
		Row:        row,
		Col:        col,
		Expression: expression,
	}
}
