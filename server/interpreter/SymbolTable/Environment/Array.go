package Environment

import (
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"reflect"
)

type Array struct {
	SymbolTable.Symbol
	Values     []interface{}
	ListIntDim *arrayList.List
}

func NewArray(id string, dataType SymbolTable.DataType, values []interface{}, listDims *arrayList.List) Array {
	symbol := SymbolTable.NewSymbolArray(0, 0, id, dataType)
	return Array{
		Symbol:     symbol,
		Values:     values,
		ListIntDim: listDims,
	}
}

func (a Array) GetValue(list *arrayList.List, indexLevel int, values []interface{}) interface{} {
	listClone := list.Clone()

	indexPiv := listClone.GetValue(0).(int)
	lengthLevel := a.ListIntDim.GetValue(indexLevel).(int)
	listClone.RemoveAtIndex(0)

	fmt.Printf("%v -> %v", indexPiv, indexLevel)
	if listClone.Len() > 0 {
		if indexPiv > lengthLevel-1 {
			fmt.Println("Error 1")
			return nil
		} else {
			subArray := values[indexPiv]
			if reflect.TypeOf(subArray) != reflect.TypeOf([]interface{}{}) {
				fmt.Println("Error 2")
				return nil
			} else {
				return a.GetValue(listClone, indexLevel+1, subArray.([]interface{}))
			}
		}
	} else {
		if indexPiv > lengthLevel-1 {
			fmt.Println("Error 3")
			return nil
		}
		return values[indexPiv]
	}
}

func (a Array) ChangeValue(list *arrayList.List, indexLevel int, values []interface{},
	newVal interface{}, symbolTable SymbolTable.SymbolTable) {
	listClone := list.Clone()

	indexPiv := listClone.GetValue(0).(int)
	lengthLevel := a.ListIntDim.GetValue(indexLevel).(int)
	listClone.RemoveAtIndex(0)

	fmt.Printf("%v -> %v", indexPiv, indexLevel)
	if listClone.Len() > 0 {
		if indexPiv > lengthLevel-1 {
			fmt.Println("Error 1")
			return
		} else {
			subArray := values[indexPiv]
			if reflect.TypeOf(subArray) != reflect.TypeOf([]interface{}{}) {
				fmt.Println("Error 2")
				return
			} else {
				a.ChangeValue(listClone, indexLevel+1, subArray.([]interface{}), newVal, symbolTable)
			}
		}
	} else {
		if indexPiv > lengthLevel-1 {
			fmt.Println("Error 3")
			return
		}
		values[indexPiv] = newVal.(Abstract.Expression).GetValue(symbolTable).Value
	}
}
