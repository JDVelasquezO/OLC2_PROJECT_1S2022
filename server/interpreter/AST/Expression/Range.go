package Expression

import (
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
)

type Range struct {
	Row   int
	Col   int
	Expr1 Abstract.Expression
	Expr2 Abstract.Expression
	Type  SymbolTable.DataType
}

func NewRange(row int, col int, expr1 Abstract.Expression, expr2 Abstract.Expression) Range {
	return Range{
		Row:   row,
		Col:   col,
		Expr1: expr1,
		Expr2: expr2,
	}
}

func (r Range) GetValue(symbolTable SymbolTable.SymbolTable) SymbolTable.ReturnType {
	return SymbolTable.ReturnType{
		Type:  r.Type,
		Value: r.Expr1,
	}
}
