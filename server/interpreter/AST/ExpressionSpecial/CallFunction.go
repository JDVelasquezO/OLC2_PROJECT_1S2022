package ExpressionSpecial

import (
	"OLC2_Project1/server/interpreter"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment"
	"OLC2_Project1/server/interpreter/errors"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"reflect"
	"strconv"
)

type CallFunction struct {
	IdFunction      string
	ListExpressions *arrayList.List
	Row             int
	Col             int
}

func NewCallFunction(idFunction string, listExpressions *arrayList.List, row int, col int) CallFunction {
	return CallFunction{
		IdFunction:      idFunction,
		ListExpressions: listExpressions,
		Row:             row,
		Col:             col,
	}
}

func (c CallFunction) GetValue(table SymbolTable.SymbolTable) SymbolTable.ReturnType {
	existsFunction := table.ExistsFunction(c.IdFunction)

	if !existsFunction {
		errors.CounterError += 1
		row := c.Row
		col := c.Col
		msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") Error: La funcion \"" + c.IdFunction + "\" No existe \n"
		err := errors.NewError(errors.CounterError, row, col, msg, table.Name)
		errors.TypeError = append(errors.TypeError, err)
		interpreter.Console += fmt.Sprintf("%v", err.Msg)
		return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
	}

	envFunction := SymbolTable.NewSymbolTable("function", &table)
	function := table.GetFunction(c.IdFunction).(Environment.Function)

	var newDataType string
	switch function.DataType {
	case SymbolTable.INTEGER:
		newDataType = "i64"
	case SymbolTable.FLOAT:
		newDataType = "f64"
	case SymbolTable.STRING:
		newDataType = "String"
	case SymbolTable.STR:
		newDataType = "&str"
	case SymbolTable.CHAR:
		newDataType = "char"
	case SymbolTable.BOOLEAN:
		newDataType = "bool"
	case SymbolTable.VOID:
		newDataType = "void"
	}

	cloneFunc := Environment.NewFunction(c.Row, c.Col, function.Id, function.ListParams.Clone(), function.ListInstructs.Clone(), newDataType)
	finished := cloneFunc.ExecuteParams(envFunction, c.ListExpressions)

	if !finished {
		msg := "(" + strconv.Itoa(c.Row) + ", " + strconv.Itoa(c.Col) + ") El parámetro no existe ni concuerda \n"
		return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: msg}
	}

	val := cloneFunc.Execute(envFunction)

	if (reflect.TypeOf(val) != reflect.TypeOf(SymbolTable.ReturnType{})) {
		return SymbolTable.ReturnType{}
	}

	return val.(SymbolTable.ReturnType)
}

func (c CallFunction) Execute(table SymbolTable.SymbolTable) interface{} {
	valRet := c.GetValue(table)
	return valRet
}
