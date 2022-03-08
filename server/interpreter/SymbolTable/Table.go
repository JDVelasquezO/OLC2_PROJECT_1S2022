package SymbolTable

import "strings"

type SymbolTable struct {
	Name          string
	Before        *SymbolTable
	Table         map[string]interface{} // string: Abstract
	FunctionTable map[string]interface{}
}

func NewSymbolTable(name string, before *SymbolTable) SymbolTable {
	table := make(map[string]interface{})
	funcTable := make(map[string]interface{})

	in := SymbolTable{
		Name:          name,
		Before:        before,
		Table:         table,
		FunctionTable: funcTable,
	}

	return in
}

func (table *SymbolTable) AddNewSymbol(id string, symbol Symbol) {
	newId := strings.ToLower(id)
	table.Table[newId] = symbol
}

func (table *SymbolTable) ExistsSymbol(id string) bool {
	newId := strings.ToLower(id)

	for actualTable := table; actualTable != nil; actualTable = actualTable.Before {

		for key, _ := range actualTable.Table {
			if key == newId {
				return true
			}
		}
	}

	return false
}

func (table *SymbolTable) GetSymbol(id string) Symbol {
	newId := strings.ToLower(id)

	for actualTable := table; actualTable != nil; actualTable = actualTable.Before {

		for key, symbol := range actualTable.Table {
			if key == newId {
				return symbol.(Symbol)
			}
		}
	}

	var symbolNil Symbol
	return symbolNil
}

func (table *SymbolTable) ChangeVal(id string, newSymbol Symbol) {
	newId := strings.ToLower(id)

	for actualTable := table; actualTable != nil; actualTable = actualTable.Before {

		for key, _ := range actualTable.Table {
			if key == newId {
				actualTable.Table[newId] = newSymbol
				return
			}
		}
	}
}

func (table *SymbolTable) AddFunction(id string, symbol interface{}) {
	idFinal := strings.ToLower(id)
	table.FunctionTable[idFinal] = symbol
}

func (table *SymbolTable) ExistsFunction(id string) bool {
	idFinal := strings.ToLower(id)

	for env := table; env != nil; env = env.Before {
		for key, _ := range env.FunctionTable {
			if key == idFinal {
				return true
			}
		}
	}

	return false
}

func (table *SymbolTable) GetFunction(id string) interface{} {
	idFinal := strings.ToLower(id)

	for actualTable := table; actualTable != nil; actualTable = actualTable.Before {
		for key, symbol := range actualTable.FunctionTable {
			if key == idFinal {
				return symbol
			}
		}
	}

	return nil
}
