package Natives

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	arrayList "github.com/colegno/arraylist"
)

type Loop struct {
	Instructions *arrayList.List
}

func (l Loop) Compile(symbolTable SymbolTable.SymbolTable, generator Generator.Generator) interface{} {
	//TODO implement me
	panic("implement me")
}

func NewLoop(instructions *arrayList.List) Loop {
	return Loop{
		Instructions: instructions,
	}
}

func (l Loop) GetValue(symbolTable SymbolTable.SymbolTable) SymbolTable.ReturnType {
	val := l.Execute(symbolTable)
	if val != nil {
		return val.(SymbolTable.ReturnType)
	}

	return SymbolTable.ReturnType{
		Type:  SymbolTable.ERROR,
		Value: nil,
	}
}

func (l Loop) Execute(symbolTable SymbolTable.SymbolTable) interface{} {

	var newTable SymbolTable.SymbolTable
StartLoop:
	if newTable.Table == nil {
		newTable = SymbolTable.NewSymbolTable("loop", &symbolTable)
	} else {
		oldTable := newTable
		newTable = SymbolTable.NewSymbolTable("loop", &oldTable)
	}

	for j := 0; j < l.Instructions.Len(); j++ {
		instruct := l.Instructions.GetValue(j).(Abstract.Instruction)
		newSymbol := instruct.Execute(newTable)

		if newSymbol != nil {
			if typeof(newSymbol) == "Natives.Break" {
				goto EndLoop
			}

			if typeof(newSymbol) == "Natives.Continue" {
				break
			}

			if typeof(newSymbol) == "SymbolTable.ReturnType" {
				return newSymbol
			}
			newTable.AddNewSymbol(newSymbol.(SymbolTable.Symbol).Id, newSymbol.(SymbolTable.Symbol))
		}
	}
	goto StartLoop

EndLoop:
	return nil
}
