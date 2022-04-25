package DecStructs

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment/Struct"
	arrayList "github.com/colegno/arraylist"
	"strings"
)

type DecStruct struct {
	Id    string
	Items *arrayList.List
	Size  int
}

func (d DecStruct) Compile(symbolTable *SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {
	symbol := interpreter.GlobalTable.StructTable[strings.ToLower(d.Id)].(DecStruct)
	symbol.SetSize(d.Items.Len())
	interpreter.GlobalTable.StructTable[strings.ToLower(d.Id)] = symbol
	return nil
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

func (d *DecStruct) SetSize(size int) {
	d.Size = size
}
