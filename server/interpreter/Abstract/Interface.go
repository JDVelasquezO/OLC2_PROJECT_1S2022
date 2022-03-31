package Abstract

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter/SymbolTable"
)

type Expression interface {
	GetValue(symbolTable SymbolTable.SymbolTable) SymbolTable.ReturnType
	Compile(symbolTable SymbolTable.SymbolTable, generator *Generator.Generator) interface{}
}

type Instruction interface {
	Execute(symbolTable SymbolTable.SymbolTable) interface{}
	Compile(symbolTable SymbolTable.SymbolTable, generator *Generator.Generator) interface{}
}
