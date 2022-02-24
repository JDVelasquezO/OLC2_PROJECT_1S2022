package Natives

import (
	"OLC2_Project1/server/interpreter"
	"OLC2_Project1/server/interpreter/AST/Expression"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/errors"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"strconv"
	"strings"
)

type Print struct {
	Expressions Abstract.Expression
	ListIds     *arrayList.List
	isBreakLine bool
	Row         int
	Col         int
}

func NewPrint(val Abstract.Expression, isBreakLine bool, row int, col int) Print {
	e := Print{val, nil, isBreakLine, row, col}
	return e
}

func NewPrintWithAfter(val Abstract.Expression, listIds *arrayList.List, isBreakLine bool, row int, col int) Print {
	e := Print{val, listIds, isBreakLine, row, col}
	return e
}

func (p Print) Execute(symbolTable SymbolTable.SymbolTable) interface{} {

	if p.Expressions == nil {
		return nil
	}

	if p.ListIds == nil {
		row := p.Row
		col := p.Col
		errors.CounterError += 1
		msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") No tiene el formato especificado la variable ID: "
		err := errors.NewError(errors.CounterError, p.Row, p.Col, msg, symbolTable.Name)
		errors.TypeError = append(errors.TypeError, err)
		interpreter.Console += fmt.Sprintf("%v", err.Msg)
		return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
	}

	if p.ListIds.Len() > 0 {

		finalMsg := ""
		for i := 0; i < p.ListIds.Len(); i++ {
			strFromList := p.ListIds.GetValue(i).(Abstract.Expression).GetValue(symbolTable)

			// Validacion si no existe el ID
			if strFromList.Value == nil {
				row := p.Row
				col := p.Col
				errors.CounterError += 1
				msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") No esta asignada la variable ID: " + p.ListIds.GetValue(i).(Expression.Identifier).Id
				err := errors.NewError(errors.CounterError, p.Row, p.Col, msg, symbolTable.Name)
				errors.TypeError = append(errors.TypeError, err)
				interpreter.Console += fmt.Sprintf("%v", err.Msg)
				return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
			}

			if strFromList.Type == SymbolTable.ERROR {
				interpreter.Console += strFromList.Value.(errors.Error).Msg
				return strFromList
			}

			strToConcat := fmt.Sprintf("%v", strFromList.Value)
			strToCompare := fmt.Sprintf("%v", p.Expressions.GetValue(symbolTable).Value)
			numberRepKeys := strings.Count(strToCompare, "{}")

			if p.ListIds.Len() != numberRepKeys {
				row := p.Row
				col := p.Col
				errors.CounterError += 1
				msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") No está el formato especificada"
				err := errors.NewError(errors.CounterError, p.Row, p.Col, msg, symbolTable.Name)
				errors.TypeError = append(errors.TypeError, err)
				interpreter.Console += fmt.Sprintf("%v", err.Msg)
				return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
			}

			words := strings.Split(strToCompare, "{}")
			finalMsg += words[i] + strToConcat
		}
		if p.isBreakLine {
			finalMsg = finalMsg + "\n"
		}
		interpreter.Console += fmt.Sprintf("%v", finalMsg)

	} else {
		varDec := p.Expressions.GetValue(symbolTable)
		if varDec.Type == SymbolTable.STR {
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
	return nil
}
