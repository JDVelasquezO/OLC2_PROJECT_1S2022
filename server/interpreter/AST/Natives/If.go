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

	if returnPrincipal.Value == "true" {
		returnPrincipal.Value = true
		returnPrincipal.Type = SymbolTable.BOOLEAN
	}

	if returnPrincipal.Value == "false" {
		returnPrincipal.Value = false
		returnPrincipal.Type = SymbolTable.BOOLEAN
	}

	//var retVal interface{}
	if returnPrincipal.Value.(bool) {
		newTable := SymbolTable.NewSymbolTable("if", &table)
		var valueRet interface{}
		for j := 0; j < i.ListInstructs.Len(); j++ {
			instruct := i.ListInstructs.GetValue(j).(Abstract.Instruction)
			valueRet = instruct.Execute(newTable)
			if valueRet != nil && typeof(valueRet) != "SymbolTable.ReturnType" {
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

func ValidateElseInstructs(instr Abstract.Instruction, newTable SymbolTable.SymbolTable) SymbolTable.ReturnType {
	typeVal := typeof(instr)
	var row int
	var col int
	switch typeVal {
	case "Expression.Primitive":
		row = instr.(Expression.Primitive).Row
		col = instr.(Expression.Primitive).Col
	case "Expression.Operation":
		row = instr.(Expression.Operation).Row
		col = instr.(Expression.Operation).Col
	case "Expression.Identifier":
		row = instr.(Expression.Identifier).Row
		col = instr.(Expression.Identifier).Col
	}
	errors.CounterError += 1
	msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ")  Tipos de datos en IF incorrectos \n"
	err := errors.NewError(errors.CounterError, 0, 0, msg, newTable.Name)
	errors.TypeError = append(errors.TypeError, err)
	interpreter.Console += fmt.Sprintf("%v", msg)
	return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err}
}
