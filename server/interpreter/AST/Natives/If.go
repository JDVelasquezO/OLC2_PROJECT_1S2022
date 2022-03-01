package Natives

import (
	"OLC2_Project1/server/interpreter"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	arrayList "github.com/colegno/arraylist"
)

type If struct {
	Condition         Abstract.Expression
	ListInstructs     *arrayList.List
	ListIfElse        *arrayList.List
	ListInstructsElse *arrayList.List
}

func NewIf(condition Abstract.Expression, listInstructs *arrayList.List,
	listIfElse *arrayList.List, listInstructsElse *arrayList.List) If {

	return If{
		Condition:         condition,
		ListInstructs:     listInstructs,
		ListIfElse:        listIfElse,
		ListInstructsElse: listInstructsElse,
	}
}

func (i If) Execute(table SymbolTable.SymbolTable) interface{} {

	returnPrincipal := i.Condition.GetValue(table)

	if returnPrincipal.Type != SymbolTable.BOOLEAN {
		interpreter.Console += "Error no es booleano"
		return nil
	}

	if returnPrincipal.Value == "true" || returnPrincipal.Value.(bool) {
		newTable := SymbolTable.NewSymbolTable("if", &table)
		for j := 0; j < i.ListInstructs.Len(); j++ {
			instruct := i.ListInstructs.GetValue(j).(Abstract.Instruction)
			instruct.Execute(newTable)
		}
	} else {
		if i.ListIfElse != nil {
			newTable := SymbolTable.NewSymbolTable("else", &table)
			for j := 0; j < i.ListInstructsElse.Len(); j++ {
				instruct := i.ListInstructsElse.GetValue(j).(Abstract.Instruction)
				instruct.Execute(newTable)
			}
		} else {
			return nil
		}
	}

	return nil
}
