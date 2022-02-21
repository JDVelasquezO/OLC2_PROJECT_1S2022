package SymbolTable

import "strings"

type SymbolTable struct {
	Name   string
	Before *SymbolTable
	Table  map[string]interface{} // string: Abstract
}

func NewSymbolTable(name string, before *SymbolTable) SymbolTable {
	in := SymbolTable{Name: name, Before: before}
	in.Table = make(map[string]interface{})
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
				table.Table[newId] = newSymbol
				return
			}
		}
	}
}
