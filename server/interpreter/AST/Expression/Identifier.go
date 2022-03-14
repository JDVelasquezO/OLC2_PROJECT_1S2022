package Expression

import (
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment/Array"
	"reflect"
)

type Identifier struct {
	Id  string
	Row int
	Col int
}

func (id Identifier) Execute(symbolTable SymbolTable.SymbolTable) interface{} {
	res := id.GetValue(symbolTable)

	return res
}

func NewIdentifier(id string, row int, col int) Identifier {
	return Identifier{
		Id:  id,
		Row: row,
		Col: col,
	}
}

func (id Identifier) GetValue(table SymbolTable.SymbolTable) SymbolTable.ReturnType {
	var founded = table.ExistsSymbol(id.Id)

	if !founded {
		return SymbolTable.ReturnType{Value: nil, Type: SymbolTable.NULL}
	}

	symbol := table.GetSymbol(id.Id)

	if reflect.TypeOf(symbol) == reflect.TypeOf(SymbolTable.Symbol{}) {
		data := symbol.(SymbolTable.Symbol)
		return SymbolTable.ReturnType{
			Value: data.Value,
			Type:  data.DataType,
		}
	}

	if typeof(symbol) == "Array.Array" {
		dataType := symbol.(Array.Array).Symbol.DataType
		values := symbol.(Array.Array).Values
		return SymbolTable.ReturnType{
			Value: values,
			Type:  dataType,
		}
	}

	return SymbolTable.ReturnType{Value: -1, Type: SymbolTable.NULL}
}
