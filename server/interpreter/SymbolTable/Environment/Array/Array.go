package Array

import (
	"OLC2_Project1/server/interpreter"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/errors"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"reflect"
	"strconv"
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

func (a Array) GetValue(list *arrayList.List, indexLevel int, values []interface{}, table SymbolTable.SymbolTable) interface{} {
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
				return a.GetValue(listClone, indexLevel+1, subArray.([]interface{}), table)
			}
		}
	} else {
		if indexPiv > lengthLevel-1 {
			row := a.Row
			col := a.Col
			errors.CounterError += 1
			msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") Desbordamiento de memoria. Índice invalido \n"
			err := errors.NewError(errors.CounterError, row, col, msg, table.Name)
			errors.TypeError = append(errors.TypeError, err)
			interpreter.Console += fmt.Sprintf("%v", err.Msg)
			return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
		}
		return values[indexPiv]
	}
}

func (a Array) ChangeValue(list *arrayList.List, indexLevel int, values []interface{},
	newVal interface{}, symbolTable SymbolTable.SymbolTable, row int, col int) {
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
				a.ChangeValue(listClone, indexLevel+1, subArray.([]interface{}), newVal, symbolTable, row, col)
			}
		}
	} else {
		if indexPiv > lengthLevel-1 {
			fmt.Println("Error 3")
			return
		}

		newValExpr := newVal.(Abstract.Expression)
		//newType := newValExpr.GetValue(symbolTable).Type
		if reflect.TypeOf(values[indexPiv]) != reflect.TypeOf(newValExpr.GetValue(symbolTable).Value) {
			errors.CounterError += 1
			msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") Tipos de datos incompatibles \n"
			err := errors.NewError(errors.CounterError, row, col, msg, symbolTable.Name)
			errors.TypeError = append(errors.TypeError, err)
			interpreter.Console += fmt.Sprintf("%v", err.Msg)
			return
		} else {
			values[indexPiv] = newValExpr.GetValue(symbolTable).Value
		}
	}
}
