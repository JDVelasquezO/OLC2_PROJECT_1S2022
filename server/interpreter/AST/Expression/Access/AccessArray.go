package Access

import (
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"reflect"
)

type ArrayAccess struct {
	Id  string
	Dim *arrayList.List
}

func NewAccessArray(id string, dims *arrayList.List) ArrayAccess {
	return ArrayAccess{
		Id:  id,
		Dim: dims,
	}
}

func (a ArrayAccess) Execute(table SymbolTable.SymbolTable) interface{} {
	res := a.GetValue(table)
	return res
}

func (a ArrayAccess) GetValue(table SymbolTable.SymbolTable) SymbolTable.ReturnType {
	exists := table.ExistsSymbol(a.Id)

	if !exists {
		fmt.Println("No existe el arreglo")
		return SymbolTable.ReturnType{
			Value: 0,
			Type:  SymbolTable.NULL,
		}
	}

	symbol := table.GetSymbol(a.Id)
	if reflect.TypeOf(symbol) != reflect.TypeOf(Environment.Array{}) {
		fmt.Println("No es un arreglo")
		return SymbolTable.ReturnType{
			Value: 0,
			Type:  SymbolTable.NULL,
		}
	}

	objectArr := symbol.(Environment.Array)

	if objectArr.ListIntDim.Len() != a.Dim.Len() {
		fmt.Println("Dimensiones invalidas")
		return SymbolTable.ReturnType{
			Value: 0,
			Type:  SymbolTable.NULL,
		}
	}

	dimensions := a.GetIntDimensions(table)
	val := objectArr.GetValue(dimensions, 0, objectArr.Values)
	return SymbolTable.ReturnType{
		Value: val,
		Type:  objectArr.DataType,
	}
}

func (a ArrayAccess) GetIntDimensions(table SymbolTable.SymbolTable) *arrayList.List {
	listDimensions := arrayList.New()

	for x := 0; x < a.Dim.Len(); x++ {
		dim := a.Dim.GetValue(x)

		valRet := dim.(Abstract.Expression).GetValue(table)

		if valRet.Type != SymbolTable.INTEGER {
			return nil
		}

		listDimensions.Add(valRet.Value)
	}

	return listDimensions
}
