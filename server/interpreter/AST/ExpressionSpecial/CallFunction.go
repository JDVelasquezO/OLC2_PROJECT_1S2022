package ExpressionSpecial

import (
	"OLC2_Project1/server/interpreter"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment"
	"OLC2_Project1/server/interpreter/errors"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"reflect"
)

type CallFunction struct {
	IdFunction      string
	ListExpressions *arrayList.List
}

func NewCallFunction(idFunction string, listExpressions *arrayList.List) CallFunction {
	return CallFunction{
		IdFunction:      idFunction,
		ListExpressions: listExpressions,
	}
}

func (c CallFunction) GetValue(table SymbolTable.SymbolTable) SymbolTable.ReturnType {
	existsFunction := table.ExistsFunction(c.IdFunction)

	if !existsFunction {
		errors.CounterError += 1
		msg := "La funcion \"" + c.IdFunction + "\" No existe \n"
		err := errors.NewError(errors.CounterError, 0, 0, msg, table.Name)
		errors.TypeError = append(errors.TypeError, err)
		interpreter.Console += fmt.Sprintf("%v", err.Msg)
		return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
	}

	envFunction := SymbolTable.NewSymbolTable("function", &table)
	function := table.GetFunction(c.IdFunction).(Environment.Function)
	cloneFunc := Environment.NewFunction(0, 0, function.Id, function.ListParams.Clone(), function.ListInstructs.Clone(), function.DataType)
	finished := cloneFunc.ExecuteParams(envFunction, c.ListExpressions)

	if !finished {
		return SymbolTable.ReturnType{
			Value: -1,
			Type:  SymbolTable.NULL,
		}
	}

	val := cloneFunc.Execute(envFunction)

	if (reflect.TypeOf(val) != reflect.TypeOf(SymbolTable.ReturnType{})) {
		return SymbolTable.ReturnType{}
	}

	return val.(SymbolTable.ReturnType)
}

func (c CallFunction) Execute(table SymbolTable.SymbolTable) interface{} {
	c.GetValue(table)
	return nil
}
