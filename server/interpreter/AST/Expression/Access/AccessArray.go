package Access

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment/Array"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment/Vector"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"reflect"
)

type ArrayAccess struct {
	Id  string
	Dim *arrayList.List
}

func (a ArrayAccess) Compile(symbolTable *SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {
	generator.AddComment("---- Access to Index Array ----")
	temp := generator.AddTemp()
	tempIndex := generator.AddTemp()
	tempSend := generator.AddTemp()

	array := symbolTable.GetSymbolArray(a.Id)
	posArray := array.(Array.Array).Pos
	indexArray := a.Dim.GetValue(0).(Abstract.Expression).GetValue(*symbolTable)

	generator.GetStack(temp, fmt.Sprintf("%v", posArray))
	generator.AddExpression(tempIndex, temp, fmt.Sprintf("%v", indexArray.Value), "+")
	generator.AddExpression(tempIndex, tempIndex, "1", "+")
	generator.GetHeap(tempSend, tempIndex)

	val := Abstract.NewValue(tempSend, array.(Array.Array).DataType, true, "")
	return val
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
	exists := table.ExistsArray(a.Id)

	if !exists {
		fmt.Println("No existe el arreglo")
		return SymbolTable.ReturnType{
			Value: nil,
			Type:  SymbolTable.NULL,
		}
	}

	symbol := table.GetSymbolArray(a.Id)

	if reflect.TypeOf(symbol) == reflect.TypeOf(Vector.Vector{}) {
		val := symbol.(Vector.Vector).GetValue(a.Dim, table)

		if val == nil {
			return SymbolTable.ReturnType{
				Value: nil,
				Type:  SymbolTable.ERROR,
			}
		}

		return SymbolTable.ReturnType{
			Value: val.(Abstract.Expression).GetValue(table).Value,
			Type:  val.(Abstract.Expression).GetValue(table).Type,
		}
	}

	if reflect.TypeOf(symbol) != reflect.TypeOf(Array.Array{}) {
		fmt.Println("No es un arreglo")
		return SymbolTable.ReturnType{
			Value: 0,
			Type:  SymbolTable.NULL,
		}
	}

	objectArr := symbol.(Array.Array)

	if objectArr.ListIntDim.Len() < a.Dim.Len() {
		fmt.Println("Dimensiones invalidas")
		return SymbolTable.ReturnType{
			Value: 0,
			Type:  SymbolTable.NULL,
		}
	}

	dimensions := a.GetIntDimensions(table)
	val := objectArr.GetValue(dimensions, 0, objectArr.Values, table)

	if reflect.TypeOf(val) == reflect.TypeOf(SymbolTable.ReturnType{}) {
		return val.(SymbolTable.ReturnType)
	}

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
