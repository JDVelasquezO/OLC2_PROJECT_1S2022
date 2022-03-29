package Natives

import (
	"OLC2_Project1/server/Generator"
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

func (p Print) Compile(symbolTable SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {
	//TODO implement me
	panic("implement me")
}

func NewPrint(val Abstract.Expression, isBreakLine bool, row int, col int) Print {
	e := Print{val, arrayList.New(), isBreakLine, row, col}
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

	if p.ListIds.Len() > 0 {

		finalMsg := ""
		var lastWord string
		for i := 0; i < p.ListIds.Len(); i++ {
			strAsExpression := p.ListIds.GetValue(i).(Abstract.Expression)
			strFromList := strAsExpression.GetValue(symbolTable)

			// Validacion si no existe el ID
			if strFromList.Value == nil {
				row := p.Row
				col := p.Col
				errors.CounterError += 1
				var msg string
				if typeof(strAsExpression) == "Access.ArrayAccess" {
					msg = "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") El arreglo no cumple las condiciones permitidas \n"
				} else {
					msg = "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") No esta asignada la variable ID: " + strAsExpression.(Expression.Identifier).Id + "\n"
				}
				err := errors.NewError(errors.CounterError, p.Row, p.Col, msg, symbolTable.Name)
				errors.TypeError = append(errors.TypeError, err)
				interpreter.Console += fmt.Sprintf("%v", err.Msg)
				return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
			}

			if strFromList.Type == SymbolTable.ERROR {
				return strFromList
			}

			strToConcat := fmt.Sprintf("%v", strFromList.Value)
			strToCompare := fmt.Sprintf("%v", p.Expressions.GetValue(symbolTable).Value)
			var numberRepKeys int
			if typeof(strAsExpression) == "Access.ArrayAccess" ||
				typeof(strFromList.Value) == "[]interface {}" {
				numberRepKeys = strings.Count(strToCompare, "{:?}")
			} else {
				numberRepKeys = strings.Count(strToCompare, "{}")
			}

			if p.ListIds.Len() != numberRepKeys {
				row := p.Row
				col := p.Col
				errors.CounterError += 1
				msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") No está el formato especificada \n"
				err := errors.NewError(errors.CounterError, p.Row, p.Col, msg, symbolTable.Name)
				errors.TypeError = append(errors.TypeError, err)
				interpreter.Console += fmt.Sprintf("%v", err.Msg)
				return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
			}

			var words []string
			if typeof(strAsExpression) == "Access.ArrayAccess" ||
				typeof(strFromList.Value) == "[]interface {}" {
				words = strings.Split(strToCompare, "{:?}")
			} else {
				words = strings.Split(strToCompare, "{}")
			}
			finalMsg += words[i] + strToConcat
			lastWord = words[i+1]
		}
		finalMsg = finalMsg + lastWord
		if p.isBreakLine {
			finalMsg = finalMsg + "\n"
		}
		interpreter.Console += fmt.Sprintf("%v", finalMsg)

	} else {
		varDec := p.Expressions.GetValue(symbolTable)
		if varDec.Type == SymbolTable.STR {
			returnExp := p.Expressions.GetValue(symbolTable)
			finalMsg := fmt.Sprintf("%v", returnExp.Value)

			if strings.Contains(finalMsg, "{}") {
				row := p.Row
				col := p.Col
				errors.CounterError += 1
				msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") No tiene formato especificado: " + fmt.Sprintf("%v", p.Expressions.GetValue(symbolTable).Value) + "\n"
				err := errors.NewError(errors.CounterError, p.Row, p.Col, msg, symbolTable.Name)
				errors.TypeError = append(errors.TypeError, err)
				interpreter.Console += fmt.Sprintf("%v", err.Msg)
				return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
			}

			if p.isBreakLine {
				finalMsg = finalMsg + "\n"
			}

			interpreter.Console += fmt.Sprintf("%v", finalMsg)
		} else {
			row := p.Row
			col := p.Col
			errors.CounterError += 1
			msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") No tiene formato especificado: " + fmt.Sprintf("%v", p.Expressions.GetValue(symbolTable).Value) + "\n"
			err := errors.NewError(errors.CounterError, p.Row, p.Col, msg, symbolTable.Name)
			errors.TypeError = append(errors.TypeError, err)
			interpreter.Console += fmt.Sprintf("%v", err.Msg)
			return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
		}
	}
	return nil
}
