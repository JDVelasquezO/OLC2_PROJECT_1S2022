package DecVectors

import (
	"OLC2_Project1/server/interpreter"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment/Vector"
	"OLC2_Project1/server/interpreter/errors"
	"fmt"
	"reflect"
	"strconv"
)

type Remove struct {
	Id    string
	Index Abstract.Expression
}

func NewRemove(id string, index Abstract.Expression) Remove {
	return Remove{
		Index: index,
		Id:    id,
	}
}

func (r Remove) GetValue(symbolTable SymbolTable.SymbolTable) SymbolTable.ReturnType {
	//TODO implement me
	panic("implement me")
}

func (r Remove) Execute(symbolTable SymbolTable.SymbolTable) interface{} {
	val := symbolTable.GetSymbolArray(r.Id)
	if val.(Vector.Vector).IsConst {
		row := val.(Vector.Vector).Row
		col := val.(Vector.Vector).Col
		id := val.(Vector.Vector).Id
		errors.CounterError += 1
		msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") Error: Variable no mutable: " + id + " \n"
		err := errors.NewError(errors.CounterError, row, col, msg, symbolTable.Name)
		errors.TypeError = append(errors.TypeError, err)
		interpreter.Console += fmt.Sprintf("%v", err.Msg)
		return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
	}
	index := r.Index.GetValue(symbolTable).Value
	newVal := val.(Vector.Vector).Remove(index.(int), symbolTable)

	if reflect.TypeOf(newVal) == reflect.TypeOf(SymbolTable.ReturnType{}) {
		if newVal.(SymbolTable.ReturnType).Type == SymbolTable.ERROR {
			return newVal
		}
	}

	symbolTable.ChangeValVector(r.Id, newVal)
	return val
}
