package DecVectors

import (
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment/Vector"
	"fmt"
)

type DecVector struct {
	Length       int
	Id           string
	ValueInitial Abstract.Expression
	Type         SymbolTable.DataType
	IsMut        bool
	Positions    Abstract.Expression
}

func NewDecVector(length int, id string, init Abstract.Expression,
	dataType SymbolTable.DataType, isMut bool) DecVector {
	return DecVector{
		Length:       length,
		Id:           id,
		ValueInitial: init,
		Type:         dataType,
		IsMut:        isMut,
	}
}

func (d DecVector) Execute(table SymbolTable.SymbolTable) interface{} {
	var objectVector Vector.Vector

	if table.ExistsSymbol(d.Id) {
		fmt.Printf("Error, variable %s ya declarada", d.Id)
	} else {
		symbol := objectVector
		symbol.Id = d.Id
		table.AddNewSymbol(d.Id, symbol)
	}

	return nil
}
