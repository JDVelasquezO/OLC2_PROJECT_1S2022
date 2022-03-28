package DecStructs

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment/Struct"
	arrayList "github.com/colegno/arraylist"
)

type DecStruct struct {
	Id    string
	Items *arrayList.List
}

func (d DecStruct) Compile(symbolTable SymbolTable.SymbolTable, generator Generator.Generator) interface{} {
	//TODO implement me
	panic("implement me")
}

func (d DecStruct) Execute(symbolTable SymbolTable.SymbolTable) interface{} {
	template := Struct.NewStruct(d.Id, d.Items)
	existsStruct := symbolTable.ExistsStruct(template.Id)

	if !existsStruct {
		symbolTable.AddStruct(template.Id, template)
		return nil
	}

	// Error clase existente
	return nil
}

func NewDecStruct(id string, items *arrayList.List) DecStruct {
	return DecStruct{
		Id:    id,
		Items: items,
	}
}
