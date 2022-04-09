package Module

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter/SymbolTable"
	arrayList "github.com/colegno/arraylist"
)

type Module struct {
	Id           string
	Instructions *arrayList.List
	Access       string
}

func (m Module) Compile(symbolTable *SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {
	//TODO implement me
	panic("implement me")
}

func NewModule(id string, instructions *arrayList.List, access string) Module {
	return Module{
		Instructions: instructions,
		Id:           id,
		Access:       access,
	}
}

func (m Module) Execute(symbolTable SymbolTable.SymbolTable) interface{} {
	return nil
}
