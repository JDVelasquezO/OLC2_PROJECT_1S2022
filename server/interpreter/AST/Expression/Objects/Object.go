package Objects

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter/SymbolTable"
	arrayList "github.com/colegno/arraylist"
)

type Object struct {
	Id         string
	Attributes *arrayList.List
}

func (o Object) Compile(symbolTable *SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {
	//TODO implement me
	panic("implement me")
}

func (o Object) GetValue(symbolTable SymbolTable.SymbolTable) SymbolTable.ReturnType {
	object := make(map[string]interface{})
	object[o.Id] = o.Attributes
	return SymbolTable.ReturnType{
		Value: object,
		Type:  SymbolTable.OBJECT,
	}
}

func NewObject(id string, attributes *arrayList.List) Object {
	return Object{
		Id:         id,
		Attributes: attributes,
	}
}
