package Natives

import (
	"OLC2_Project1/server/interpreter"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/errors"
	"fmt"
	"strconv"
)

type Print struct {
	Expressions     Abstract.Expression
	ExpressionAfter Abstract.Expression
	isBreakLine     bool
	Row             int
	Col             int
}

func NewPrint(val Abstract.Expression, isBreakLine bool, row int, col int) Print {
	e := Print{val, nil, isBreakLine, row, col}
	return e
}

func NewPrintWithAfter(val Abstract.Expression, after Abstract.Expression, isBreakLine bool, row int, col int) Print {
	e := Print{val, after, isBreakLine, row, col}
	return e
}

func (p Print) Execute(symbolTable SymbolTable.SymbolTable) interface{} {

	if p.Expressions == nil {
		return nil
	}

	if p.ExpressionAfter != nil {

	} else {
		varDec := p.Expressions.GetValue(symbolTable)
		if varDec.Type == SymbolTable.STRING {
			returnExp := p.Expressions.GetValue(symbolTable)

			finalMsg := fmt.Sprintf("%v", returnExp.Value)
			if p.isBreakLine {
				finalMsg = finalMsg + "\n"
			}

			interpreter.Console += fmt.Sprintf("%v", finalMsg)
		} else {
			row := p.Row
			col := p.Col
			errors.CounterError += 1
			msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") No tiene formato especificado: " + fmt.Sprintf("%v", p.Expressions.GetValue(symbolTable).Value)
			err := errors.NewError(errors.CounterError, p.Row, p.Col, msg, symbolTable.Name)
			errors.TypeError = append(errors.TypeError, err)
			interpreter.Console += fmt.Sprintf("%v", err.Msg)
			return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
		}
	}

	//varDec := p.Expressions.GetValue(symbolTable)
	//// Validacion si no existe el ID
	//if varDec.Value == nil {
	//	row := p.Row
	//	col := p.Col
	//	errors.CounterError += 1
	//	msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") No esta asignada la variable ID: " + p.Expressions.(Expression.Identifier).Id
	//	err := errors.NewError(errors.CounterError, p.Row, p.Col, msg, symbolTable.Name)
	//	errors.TypeError = append(errors.TypeError, err)
	//	interpreter.Console += fmt.Sprintf("%v", err.Msg)
	//	return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
	//}

	return nil
}
