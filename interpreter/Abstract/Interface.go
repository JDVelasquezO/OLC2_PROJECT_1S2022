package Abstract

import "OLC2_Project1/server/interpreter/SymbolTable"

type Expression interface {
	GetValue(symbolTable SymbolTable.SymbolTable) SymbolTable.ReturnType
}

type Instruction interface {
	Execute(symbolTable SymbolTable.SymbolTable) interface{}
}
