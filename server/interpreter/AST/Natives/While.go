package Natives

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	arrayList "github.com/colegno/arraylist"
)

type While struct {
	Condition     Abstract.Expression
	ListInstructs *arrayList.List
	Row           int
	Col           int
}

func (w While) Compile(symbolTable *SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {
	generator.AddComment("---- Start While ----")
	newTable := SymbolTable.NewSymbolTable("while", symbolTable)
	newTable.SizeTable = symbolTable.SizeTable

	continueLabel := generator.NewLabel()
	generator.SetLabel(continueLabel)
	condition := w.Condition.Compile(symbolTable, generator)
	newTable.BreakLabel = condition.(Abstract.Value).FalseLabel
	newTable.ContinueLabel = continueLabel

	generator.SetLabel(condition.(Abstract.Value).TrueLabel)

	for i := 0; i < w.ListInstructs.Len(); i++ {
		w.ListInstructs.GetValue(i).(Abstract.Instruction).Compile(&newTable, generator)
	}

	generator.AddGoTo(continueLabel)
	generator.SetLabel(condition.(Abstract.Value).FalseLabel)

	return nil
}

func NewWhile(condition Abstract.Expression, listInstructs *arrayList.List, row int, col int) While {

	return While{
		Condition:     condition,
		ListInstructs: listInstructs,
		Row:           row,
		Col:           col,
	}
}

func (w While) Execute(table SymbolTable.SymbolTable) interface{} {

	returnPrincipal := w.Condition.GetValue(table)

	if returnPrincipal.Type != SymbolTable.BOOLEAN {
		interpreter.Console += "Error no es booleano"
		return nil
	}

	//var retVal interface{}
	var newTable SymbolTable.SymbolTable
StartLoop:
	for returnPrincipal.Value == "true" || returnPrincipal.Value.(bool) {

		if newTable.Table == nil {
			newTable = SymbolTable.NewSymbolTable("while", &table)
		} else {
			oldTable := newTable
			newTable = SymbolTable.NewSymbolTable("while", &oldTable)
		}

		for j := 0; j < w.ListInstructs.Len(); j++ {
			instruct := w.ListInstructs.GetValue(j).(Abstract.Instruction)
			newSymbol := instruct.Execute(newTable)

			if newSymbol != nil {

				if typeof(newSymbol) == "string" {
					return newSymbol
				}

				if typeof(newSymbol) == "Natives.Break" {
					return nil
				}

				if typeof(newSymbol) == "Natives.Continue" {
					break
				}

				if typeof(newSymbol) == "SymbolTable.ReturnType" {
					return newSymbol
				}
				valueToSend := newSymbol.(*SymbolTable.Symbol)
				newTable.AddNewSymbol(newSymbol.(*SymbolTable.Symbol).Id, valueToSend)
			}
			//fmt.Println(otherTable)
		}
		returnPrincipal = w.Condition.GetValue(newTable)
		if returnPrincipal.Value.(bool) {
			goto StartLoop
		} else {
			return nil
		}
	}
	return nil
}
