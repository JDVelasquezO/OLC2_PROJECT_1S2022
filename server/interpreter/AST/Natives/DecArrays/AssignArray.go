package DecArrays

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment/Array"
	"OLC2_Project1/server/interpreter/errors"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"strconv"
)

type AssignArray struct {
	Val       Abstract.Expression
	Id        string
	ListIndex *arrayList.List
	Row       int
	Col       int
}

func (a AssignArray) Compile(symbolTable SymbolTable.SymbolTable, generator Generator.Generator) interface{} {
	//TODO implement me
	panic("implement me")
}

func (a AssignArray) Execute(symbolTable SymbolTable.SymbolTable) interface{} {
	if a.Val == nil {
		row := 0
		col := 0
		errors.CounterError += 1
		msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") No se puede mezclar esos tipos de datos \n"
		err := errors.NewError(errors.CounterError, row, col, msg, symbolTable.Name)
		errors.TypeError = append(errors.TypeError, err)
		interpreter.Console += fmt.Sprintf("%v", err.Msg)
		return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
	}

	foundSymbol := symbolTable.GetSymbolArray(a.Id)
	symbolArray := foundSymbol.(Array.Array)

	if symbolArray.IsConst {
		row := symbolArray.Row
		col := symbolArray.Col
		errors.CounterError += 1
		msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") No se puede declarar valor a un no MUT \n"
		err := errors.NewError(errors.CounterError, row, col, msg, symbolTable.Name)
		errors.TypeError = append(errors.TypeError, err)
		interpreter.Console += fmt.Sprintf("%v", err.Msg)
		return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
	}

	dimensions := a.GetIntDimensions(symbolTable)
	symbolArray.ChangeValue(dimensions, 0, symbolArray.Values, a.Val, symbolTable, a.Row, a.Col)
	return symbolArray
}

func NewAssignArray(id string, listInArray *arrayList.List, val Abstract.Expression, row int, col int) *AssignArray {
	return &AssignArray{
		Val:       val,
		Id:        id,
		ListIndex: listInArray,
		Row:       row,
		Col:       col,
	}
}

func (a AssignArray) GetIntDimensions(table SymbolTable.SymbolTable) *arrayList.List {
	listDimensions := arrayList.New()

	for x := 0; x < a.ListIndex.Len(); x++ {
		dim := a.ListIndex.GetValue(x)

		valRet := dim.(Abstract.Expression).GetValue(table)

		if valRet.Type != SymbolTable.INTEGER {
			return nil
		}

		listDimensions.Add(valRet.Value)
	}

	return listDimensions
}
