package DecArrays

import (
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment"
	"fmt"
	"reflect"
)

type DecArray struct {
	Length       int
	Id           string
	ValueInitial Abstract.Expression
	Type         SymbolTable.DataType
}

func NewDecArray(length int, id string, init Abstract.Expression, dataType SymbolTable.DataType) DecArray {
	return DecArray{
		Length:       length,
		Id:           id,
		ValueInitial: init,
		Type:         dataType,
	}
}

func (d DecArray) Execute(table SymbolTable.SymbolTable) interface{} {
	valueDec := d.ValueInitial.GetValue(table)

	if valueDec.Type != SymbolTable.VOID {
		fmt.Printf("%v", valueDec)
	}

	if valueDec.Type != SymbolTable.ARRAY {
		fmt.Println("Err1")
		return nil
	}

	if reflect.TypeOf(valueDec.Value) != reflect.TypeOf(SymbolTable.ReturnType{}) {
		fmt.Println("Err2")
		return nil
	}

	objectArray := valueDec.Value.(SymbolTable.ReturnType)

	if objectArray.Type != d.Type {
		fmt.Println("Err3")
		return nil
	}

	if objectArray.Value.(Environment.Array).ListIntDim.Len() > d.Length {
		fmt.Printf("Error, variable %s no es posible declarar", d.Id)
	}

	if table.ExistsSymbol(d.Id) {
		fmt.Printf("Error, variable %s ya declarada", d.Id)
	} else {
		symbol := objectArray.Value.(Environment.Array)
		symbol.Id = d.Id
		table.AddNewSymbol(d.Id, symbol)
	}

	fmt.Printf("%v", table.Table)
	return nil
}
