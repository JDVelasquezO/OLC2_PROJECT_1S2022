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

type Push struct {
	Id    string
	Value Abstract.Expression
}

func (p Push) GetValue(symbolTable SymbolTable.SymbolTable) SymbolTable.ReturnType {
	//TODO implement me
	panic("implement me")
}

func NewPush(id string, value Abstract.Expression) Push {
	return Push{
		Id:    id,
		Value: value,
	}
}

func (p Push) Execute(symbolTable SymbolTable.SymbolTable) interface{} {
	val := symbolTable.GetSymbolArray(p.Id)
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
	newVal := val.(Vector.Vector).Push(p.Value, symbolTable)
	symbolTable.ChangeValVector(p.Id, newVal)
	return val
}

func (p Push) ExecuteFirstTime(symbolTable SymbolTable.SymbolTable) interface{} {
	val := symbolTable.GetSymbol(p.Id)
	newVal := val.(Vector.Vector).Push(p.Value, symbolTable)
	symbolTable.ChangeVal(p.Id, newVal)
	return val
}
