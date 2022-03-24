package Access

import (
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
)

type ModuleAccess struct {
	Id       string
	NameFunc Abstract.Expression
}

func NewAccessModule(id string, nameFunc Abstract.Expression) ModuleAccess {
	return ModuleAccess{
		Id:       id,
		NameFunc: nameFunc,
	}
}

func (a ModuleAccess) GetValue(table SymbolTable.SymbolTable) SymbolTable.ReturnType {
	return SymbolTable.ReturnType{}
}
