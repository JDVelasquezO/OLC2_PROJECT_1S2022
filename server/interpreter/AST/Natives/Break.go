package Natives

import (
	"OLC2_Project1/server/interpreter/SymbolTable"
)

type Break struct {
	Row int
	Col int
}

func (b Break) Execute(symbolTable SymbolTable.SymbolTable) interface{} {
	return b
}

func NewBreak(row int, col int) Break {
	return Break{
		Row: row,
		Col: col,
	}
}
