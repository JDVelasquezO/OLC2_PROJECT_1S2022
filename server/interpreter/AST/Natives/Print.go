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

func (p Print) Compile(symbolTable *SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {

	//var retVal Abstract.Value
	//if p.ListIds.Len() > 0 {
	//	for i := 0; i < p.ListIds.Len(); i++ {
	//		strAsExpression := p.ListIds.GetValue(i).(Abstract.Expression)
	//
	//		left := strAsExpression.Compile(symbolTable, generator)
	//		right := p.Expressions.Compile(symbolTable, generator)
	//		generator.ConcatString()
	//		paramTemp := generator.AddTemp()
	//		generator.AddExpression(paramTemp, "P", strconv.Itoa(symbolTable.SizeTable), "+")
	//
	//		// Val left
	//		generator.AddExpression(paramTemp, paramTemp, "1", "+")
	//		generator.SetStack(paramTemp, left.(Abstract.Value).Value, true)
	//
	//		// Val right
	//		generator.AddExpression(paramTemp, paramTemp, "1", "+")
	//		generator.SetStack(paramTemp, right.(Abstract.Value).Value, true)
	//
	//		generator.NewEnv(symbolTable.SizeTable)
	//		generator.CallFunc("concat")
	//		temp := generator.AddTemp()
	//
	//		generator.GetStack(temp, "P")
	//		generator.SetEnv(symbolTable.SizeTable)
	//
	//		retVal = Abstract.NewValue(temp, SymbolTable.STRING, true, "")
	//		retVal.Size = left.(Abstract.Value).Size + right.(Abstract.Value).Size
	//		//return retVal
	//	}
	//}

	generator.AddComment("---- Start Print ----")
	var res0 interface{}
	if p.ListIds.Len() == 0 {
		res0 = p.Expressions.(Abstract.Expression).Compile(symbolTable, generator)
	} else {
		res0 = p.ListIds.GetValue(0).(Abstract.Expression).Compile(symbolTable, generator)
	}
	//fmt.Println(res0)

	generator.AddComment("---- Print ----")
	if res0.(Abstract.Value).Type == SymbolTable.INTEGER {
		newVal := fmt.Sprintf("%v", res0.(Abstract.Value).Value)
		generator.AddPrint("d", "int", newVal)
	} else if res0.(Abstract.Value).Type == SymbolTable.FLOAT {
		newVal := fmt.Sprintf("%v", res0.(Abstract.Value).Value)
		generator.AddPrint("f", "double", newVal)
	} else if res0.(Abstract.Value).Type == SymbolTable.BOOLEAN {
		TypeBoolean(res0.(Abstract.Value), generator)
	} else if res0.(Abstract.Value).Type == SymbolTable.STRING || res0.(Abstract.Value).Type == SymbolTable.STR {
		TypeString(res0.(Abstract.Value).Value.(string), *symbolTable, generator)
	}

	if p.isBreakLine {
		generator.AddPrint("c", "char", "10")
	}

	return nil
}

func TypeString(value string, table SymbolTable.SymbolTable, generator *Generator.Generator) {
	generator.PrintString()
	paramTemp1 := generator.AddTemp()
	generator.AddExpression(paramTemp1, "P", strconv.Itoa(table.SizeTable), "+")
	generator.AddExpression(paramTemp1, paramTemp1, "1", "+")
	generator.SetStack(paramTemp1, value, true)
	generator.NewEnv(table.SizeTable)
	generator.CallFunc("print")

	temp := generator.AddTemp()
	generator.GetStack(temp, "P")
	generator.SetEnv(table.SizeTable)
}

func TypeBoolean(value Abstract.Value, generator *Generator.Generator) {
	exitLabel := generator.NewLabel()

	generator.SetLabel(value.TrueLabel)
	generator.PrintTrue()
	generator.AddGoTo(exitLabel)

	generator.SetLabel(value.FalseLabel)
	generator.PrintFalse()
	generator.AddGoTo(exitLabel)

	generator.SetLabel(exitLabel)
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
		interpreter.FinalMsg = finalMsg

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
			interpreter.FinalMsg = finalMsg
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
	return interpreter.FinalMsg
}
