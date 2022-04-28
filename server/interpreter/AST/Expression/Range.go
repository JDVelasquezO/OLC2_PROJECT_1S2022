package Expression

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	arrayList "github.com/colegno/arraylist"
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

func (r Range) Compile(symbolTable *SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {
	value1 := r.Expr1.Compile(symbolTable, generator)
	value2 := r.Expr2.Compile(symbolTable, generator)

	auxList := *arrayList.New()
	auxList.Add(value1)
	auxList.Add(value2)

	value := Abstract.NewValue(auxList, r.Type, false, "")
	return value
}

func (r Range) GetValue(symbolTable SymbolTable.SymbolTable) SymbolTable.ReturnType {
	return SymbolTable.ReturnType{
		Type:  r.Type,
		Value: r.Expr1,
	}
}
