package SymbolTable

import (
	"strings"
)

type SymbolTable struct {
	Name          string
	Before        *SymbolTable
	Table         map[string]interface{} // string: Abstract
	FunctionTable map[string]interface{}
	ArrayTable    map[string]interface{}
	StructTable   map[string]interface{}
	ObjectTable   map[string]interface{}
	ModuleTable   map[string]interface{}
	BreakLabel    string
	ContinueLabel string
	ReturnLabel   string
	SizeTable     int
}

func NewSymbolTable(name string, before *SymbolTable) SymbolTable {
	table := make(map[string]interface{})
	funcTable := make(map[string]interface{})
	arrTable := make(map[string]interface{})
	structTable := make(map[string]interface{})
	objTable := make(map[string]interface{})
	modTable := make(map[string]interface{})

	in := SymbolTable{
		Name:          name,
		Before:        before,
		Table:         table,
		FunctionTable: funcTable,
		ArrayTable:    arrTable,
		StructTable:   structTable,
		ObjectTable:   objTable,
		ModuleTable:   modTable,
		BreakLabel:    "",
		ContinueLabel: "",
		ReturnLabel:   "",
		SizeTable:     0,
	}

	return in
}

func (table *SymbolTable) AddNewSymbol(id string, symbol interface{}) {
	newId := strings.ToLower(id)
	table.Table[newId] = symbol
	table.SizeTable += 1
	symbol.(Symbol).SetPosition(table.SizeTable)
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

func (table *SymbolTable) ExistsSymbolInEnvironment(id string) bool {
	newId := strings.ToLower(id)

	for key, _ := range table.Table {
		if key == newId {
			return true
		}
	}

	return false
}

func (table *SymbolTable) GetSymbol(id string) interface{} {
	newId := strings.ToLower(id)

	for actualTable := table; actualTable != nil; actualTable = actualTable.Before {

		for key, symbol := range actualTable.Table {
			if key == newId {
				return symbol
			}
		}
	}

	var symbolNil Symbol
	return symbolNil
}

func (table *SymbolTable) GetSymbolArray(id string) interface{} {
	newId := strings.ToLower(id)

	for actualTable := table; actualTable != nil; actualTable = actualTable.Before {

		for key, symbol := range actualTable.ArrayTable {
			if key == newId {
				return symbol
			}
		}
	}

	var symbolNil Symbol
	return symbolNil
}

func (table *SymbolTable) ChangeVal(id string, newSymbol interface{}) {
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

func (table *SymbolTable) AddArray(id string, symbol interface{}) {
	idFinal := strings.ToLower(id)
	table.ArrayTable[idFinal] = symbol
}

func (table *SymbolTable) ExistsArray(id string) bool {
	idFinal := strings.ToLower(id)

	for env := table; env != nil; env = env.Before {
		for key, _ := range env.ArrayTable {
			if key == idFinal {
				return true
			}
		}
	}

	return false
}

func (table *SymbolTable) ChangeValArray(id string, newSymbol interface{}) {
	newId := strings.ToLower(id)

	for actualTable := table; actualTable != nil; actualTable = actualTable.Before {

		for key, _ := range actualTable.ArrayTable {
			if key == newId {
				actualTable.ArrayTable[newId] = newSymbol.(Symbol).Value
				return
			}
		}
	}
}

func (table *SymbolTable) ChangeValVector(id string, newSymbol interface{}) {
	newId := strings.ToLower(id)

	for actualTable := table; actualTable != nil; actualTable = actualTable.Before {

		for key, _ := range actualTable.ArrayTable {
			if key == newId {
				actualTable.ArrayTable[newId] = newSymbol
				return
			}
		}
	}
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

func (table *SymbolTable) ExistsStruct(id string) bool {
	idFinal := strings.ToLower(id)

	for actualTable := table; actualTable != nil; actualTable = actualTable.Before {
		for key, _ := range actualTable.StructTable {
			if key == idFinal {
				return true
			}
		}
	}

	return false
}

func (table *SymbolTable) GetStruct(id string) interface{} {
	idFinal := strings.ToLower(id)

	for actualTable := table; actualTable != nil; actualTable = actualTable.Before {
		for key, structElement := range actualTable.StructTable {
			if key == idFinal {
				return structElement
			}
		}
	}

	return nil
}

func (table *SymbolTable) AddStruct(id string, symbol interface{}) {
	idFinal := strings.ToLower(id)
	table.StructTable[idFinal] = symbol
}

func (table *SymbolTable) AddObject(id string, symbol interface{}) {
	idFinal := strings.ToLower(id)
	table.ObjectTable[idFinal] = symbol
}

func (table *SymbolTable) GetObject(id string) interface{} {
	idFinal := strings.ToLower(id)

	for actualTable := table; actualTable != nil; actualTable = actualTable.Before {
		for key, structElement := range actualTable.ObjectTable {
			if key == idFinal {
				return structElement
			}
		}
	}

	return nil
}

func (table *SymbolTable) ChangeValObject(id string, newSymbol interface{}) {
	newId := strings.ToLower(id)

	for actualTable := table; actualTable != nil; actualTable = actualTable.Before {

		for key, _ := range actualTable.ObjectTable {
			if key == newId {
				actualTable.ObjectTable[newId] = newSymbol
				return
			}
		}
	}
}

func (table *SymbolTable) AddModule(id string, symbol interface{}) {
	idFinal := strings.ToLower(id)
	table.ModuleTable[idFinal] = symbol
}
