package Natives

import (
	"OLC2_Project1/server/interpreter"
	"OLC2_Project1/server/interpreter/AST/Expression"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/errors"
	"fmt"
	"strconv"
)

type Print struct {
	Expressions Abstract.Expression
	isBreakLine bool
	Row         int
	Col         int
}

func NewPrint(val Abstract.Expression, isBreakLine bool, row int, col int) Print {
	e := Print{val, isBreakLine, row, col}
	return e
}

func (p Print) Execute(symbolTable SymbolTable.SymbolTable) interface{} {

	if p.Expressions == nil {
		return nil
	}

	varDec := p.Expressions.GetValue(symbolTable)

	if varDec.Type == SymbolTable.ERROR {
		interpreter.Console += fmt.Sprintf("%v", varDec.Value.(errors.Error).Msg)
		return nil
	}

	if varDec.Value == nil {
		row := p.Row
		col := p.Col
		errors.CounterError += 1
		msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") No esta asignada la variable ID: " + p.Expressions.(Expression.Identifier).Id
		err := errors.NewError(errors.CounterError, p.Row, p.Col, msg, symbolTable.Name)
		errors.TypeError = append(errors.TypeError, err)
		interpreter.Console += fmt.Sprintf("%v", err.Msg)
		return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
	}

	returnExp := p.Expressions.GetValue(symbolTable)

	finalMsg := fmt.Sprintf("%v", returnExp.Value)
	if p.isBreakLine {
		finalMsg = finalMsg + "\n"
	}

	interpreter.Console += fmt.Sprintf("%v", finalMsg)

	return nil
}
