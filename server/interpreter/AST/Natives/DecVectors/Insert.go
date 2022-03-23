package DecVectors

import (
	"OLC2_Project1/server/interpreter"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment/Vector"
	"OLC2_Project1/server/interpreter/errors"
	"fmt"
	"strconv"
)

type Insert struct {
	Id    string
	Value Abstract.Expression
	Index int
}

func NewInsert(id string, value Abstract.Expression, index int) Insert {
	return Insert{
		Id:    id,
		Value: value,
		Index: index,
	}
}

func (i Insert) GetValue(symbolTable SymbolTable.SymbolTable) SymbolTable.ReturnType {
	//TODO implement me
	panic("implement me")
}

func (i Insert) Execute(symbolTable SymbolTable.SymbolTable) interface{} {
	val := symbolTable.GetSymbolArray(i.Id)
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

	newVal := val.(Vector.Vector).Insert(i.Value, i.Index, symbolTable)
	symbolTable.ChangeValVector(i.Id, newVal)
	return SymbolTable.ReturnType{
		Value: val,
		Type:  SymbolTable.ARRAY,
	}
	//return val
}
