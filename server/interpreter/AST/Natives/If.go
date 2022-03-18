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
	Row               int
	Col               int
}

func NewIf(condition Abstract.Expression, listInstructs *arrayList.List,
	listIfElse *arrayList.List, listInstructsElse *arrayList.List, row int, col int) If {

	return If{
		Condition:         condition,
		ListInstructs:     listInstructs,
		ListIfElse:        listIfElse,
		ListInstructsElse: listInstructsElse,
		Row:               row,
		Col:               col,
	}
}

func (i If) GetValue(symbolTable SymbolTable.SymbolTable) SymbolTable.ReturnType {
	val := i.Execute(symbolTable)
	if val != nil {
		return val.(SymbolTable.ReturnType)
	}

	return SymbolTable.ReturnType{
		Type:  SymbolTable.ERROR,
		Value: nil,
	}
}

func (i If) Execute(table SymbolTable.SymbolTable) interface{} {

	returnPrincipal := i.Condition.GetValue(table)

	if returnPrincipal.Type != SymbolTable.BOOLEAN {
		interpreter.Console += "Error no es booleano"
		return nil
	}

	//var retVal interface{}
	if returnPrincipal.Value.(bool) {
		newTable := SymbolTable.NewSymbolTable("if", &table)
		var valueRet interface{}
		for j := 0; j < i.ListInstructs.Len(); j++ {
			instruct := i.ListInstructs.GetValue(j).(Abstract.Instruction)
			valueRet = instruct.Execute(newTable)
			if valueRet != nil && typeof(valueRet) != "SymbolTable.ReturnType" {

				if typeof(valueRet) == "Natives.Break" || typeof(valueRet) == "Natives.Continue" {
					return valueRet
				}

				newTable.AddNewSymbol(valueRet.(SymbolTable.Symbol).Id, valueRet.(SymbolTable.Symbol))
			}
		}
		return valueRet
	} else {
		if i.ListIfElse != nil {
			var valueRet interface{}
			for _, instr := range i.ListIfElse.ToArray() {
				newIf := instr.(If)
				returnConditionNewIf := newIf.Condition.GetValue(table)
				if returnConditionNewIf.Type != SymbolTable.BOOLEAN {
					return nil
				}

				if returnConditionNewIf.Value.(bool) {
					tableNewIf := SymbolTable.NewSymbolTable("Else-If", &table)
					for j := 0; j < newIf.ListInstructs.Len(); j++ {
						instr := newIf.ListInstructs.GetValue(j).(Abstract.Instruction)
						valueRet = instr.Execute(tableNewIf)
					}
					return valueRet
				}
			}
		}

		if i.ListInstructsElse != nil {
			newTable := SymbolTable.NewSymbolTable("else", &table)
			var valueRet interface{}
			for j := 0; j < i.ListInstructsElse.Len(); j++ {
				instruct := i.ListInstructsElse.GetValue(j).(Abstract.Instruction)
				valueRet = instruct.Execute(newTable)

				if valueRet == nil {
					return nil
				}

				if valueRet.(SymbolTable.ReturnType).Type == SymbolTable.ERROR {
					return nil
				}
			}
			return valueRet
		} else {
			return nil
		}
	}
}
