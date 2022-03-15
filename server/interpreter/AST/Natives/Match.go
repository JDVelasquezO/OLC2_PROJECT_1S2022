package Natives

import (
	"OLC2_Project1/server/interpreter/AST/Expression"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	arrayList "github.com/colegno/arraylist"
)

type Match struct {
	ExprToValue Abstract.Expression
	ListMatches *arrayList.List
}

func NewMatch(exprValue Abstract.Expression, listMatches *arrayList.List) Match {
	return Match{
		ListMatches: listMatches,
		ExprToValue: exprValue,
	}
}

func (m Match) GetValue(symbolTable SymbolTable.SymbolTable) SymbolTable.ReturnType {
	return SymbolTable.ReturnType{}
}

func (m Match) Execute(table SymbolTable.SymbolTable) interface{} {
	returnPrincipal := m.ExprToValue.GetValue(table)

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
						valInstructs.GetValue(k).(Abstract.Instruction).Execute(newTable)
						return nil
					}
				}
			} else if typeof(valEvaluate) == "Expression.Primitive" {
				valCase = valEvaluate.(Expression.Primitive).Value
			}
			switch returnPrincipal.Value {
			case valCase:
				for k := 0; k < valInstructs.Len(); k++ {
					newTable := SymbolTable.NewSymbolTable("Match", &table)
					valInstructs.GetValue(k).(Abstract.Instruction).Execute(newTable)
					return nil
				}
			}
		}
	}

	return returnPrincipal
}
