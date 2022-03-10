package Expression

import (
	"OLC2_Project1/server/interpreter/SymbolTable"
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

	if reflect.TypeOf(symbol) == reflect.TypeOf(SymbolTable.ReturnType{}) {
		data := symbol.(SymbolTable.Symbol)
		return SymbolTable.ReturnType{
			Value: data.Value,
			Type:  data.DataType,
		}
	}

	return SymbolTable.ReturnType{Value: -1, Type: SymbolTable.NULL}
}
