package Objects

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter/AST/Expression"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
)

type Attribute struct {
	Id    string
	Value Abstract.Expression
}

func (a Attribute) Compile(symbolTable SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {
	//TODO implement me
	panic("implement me")
}

func (a Attribute) GetValue(symbolTable SymbolTable.SymbolTable) SymbolTable.ReturnType {
	attr := make(map[string]interface{})
	attr[a.Id] = a.Value
	return SymbolTable.ReturnType{
		Value: attr,
		Type:  a.Value.(Expression.Primitive).Type,
	}
}

func NewAttribute(id string, attributes Abstract.Expression) Attribute {
	return Attribute{
		Id:    id,
		Value: attributes,
	}
}
