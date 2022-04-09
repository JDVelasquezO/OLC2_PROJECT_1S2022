package Expression

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter/SymbolTable"
)

type ItemStruct struct {
	Row      int
	Col      int
	Id       string
	DataType SymbolTable.DataType
}

func (i ItemStruct) Compile(symbolTable *SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {
	//TODO implement me
	panic("implement me")
}

func NewItemStruct(id string, dataType SymbolTable.DataType, row int, col int) ItemStruct {
	return ItemStruct{
		Id:       id,
		DataType: dataType,
		Row:      row,
		Col:      col,
	}
}

func (i ItemStruct) GetValue(symbolTable SymbolTable.SymbolTable) SymbolTable.ReturnType {
	return SymbolTable.ReturnType{}
}
