package Natives

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter/SymbolTable"
)

type Continue struct {
	Row int
	Col int
}

func (c Continue) Execute(symbolTable SymbolTable.SymbolTable) interface{} {
	return c
}

func (c Continue) Compile(symbolTable SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {
	return nil
}

func NewContinue(row int, col int) Continue {
	return Continue{
		Row: row,
		Col: col,
	}
}
