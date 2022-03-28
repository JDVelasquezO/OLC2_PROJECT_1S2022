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
)

type Match struct {
	ExprToValue Abstract.Expression
	ListMatches *arrayList.List
}

func (m Match) Compile(symbolTable SymbolTable.SymbolTable, generator Generator.Generator) interface{} {
	//TODO implement me
	panic("implement me")
}

func NewMatch(exprValue Abstract.Expression, listMatches *arrayList.List) Match {
	return Match{
		ListMatches: listMatches,
		ExprToValue: exprValue,
	}
}

func (m Match) GetValue(symbolTable SymbolTable.SymbolTable) SymbolTable.ReturnType {
	val := m.Execute(symbolTable)
	if val != nil {
		return val.(SymbolTable.ReturnType)
	}

	return SymbolTable.ReturnType{
		Type:  SymbolTable.ERROR,
		Value: nil,
	}
}

func (m Match) Execute(table SymbolTable.SymbolTable) interface{} {
	returnPrincipal := m.ExprToValue.GetValue(table)

	// Primera pasada para errores semánticos:
	for i := 0; i < m.ListMatches.Len(); i++ {
		val := m.ListMatches.GetValue(i)
		for j := 0; j < val.(Case).Expressions.Len(); j++ {
			valEvaluate := val.(Case).Expressions.GetValue(j)
			if valEvaluate.(Abstract.Expression).GetValue(table).Type != returnPrincipal.Type {
				if typeof(valEvaluate) == "Expression.Identifier" {
					if valEvaluate.(Expression.Identifier).Id != "_" {
						row := val.(Case).Row
						col := val.(Case).Col
						errors.CounterError += 1
						msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") Los tipos de datos no coinciden \n"
						err := errors.NewError(errors.CounterError, row, col, msg, table.Name)
						errors.TypeError = append(errors.TypeError, err)
						interpreter.Console += fmt.Sprintf("%v", err.Msg)
						return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
					} else {
						continue
					}
				}
			}
		}
	}

	for i := 0; i < m.ListMatches.Len(); i++ {
		val := m.ListMatches.GetValue(i)

		for j := 0; j < val.(Case).Expressions.Len(); j++ {

			valEvaluate := val.(Case).Expressions.GetValue(j)
			valInstructs := val.(Case).Instructions

			var valCase interface{}
			if typeof(valEvaluate) == "Expression.Identifier" {
				valCaseId := valEvaluate.(Expression.Identifier).Id
				valCase = valEvaluate.(Expression.Identifier).GetValue(table).Value
				if valCase == nil && valCaseId == "_" {
					for k := 0; k < valInstructs.Len(); k++ {
						newTable := SymbolTable.NewSymbolTable("Default", &table)
						instr := valInstructs.GetValue(k).(Abstract.Instruction).Execute(newTable)
						if instr != nil {
							return instr
						}
						return nil
					}
				}
			} else {
				valCase = valEvaluate.(Abstract.Expression).GetValue(table).Value
			}
			switch returnPrincipal.Value {
			case valCase:
				for k := 0; k < valInstructs.Len(); k++ {
					newTable := SymbolTable.NewSymbolTable("Match", &table)
					instr := valInstructs.GetValue(k).(Abstract.Instruction).Execute(newTable)
					if instr != nil {
						return instr
					}
					return nil
				}
			}
		}
	}

	return nil
}
