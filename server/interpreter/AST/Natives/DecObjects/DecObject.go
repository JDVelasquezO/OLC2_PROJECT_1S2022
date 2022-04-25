package DecObjects

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter/AST/Expression"
	"OLC2_Project1/server/interpreter/AST/Expression/Objects"
	"OLC2_Project1/server/interpreter/AST/Natives/DecStructs"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	arrayList "github.com/colegno/arraylist"
	"strconv"
)

type DecObjects struct {
	Id         string
	IdObject   string
	Attributes *arrayList.List
}

func (d DecObjects) Compile(symbolTable *SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {
	value := symbolTable.GetObject(d.Id)

	generator.AddComment("---- Declare Struct Instance ----")
	temp := generator.AddTemp()
	generator.AddExpression(temp, "P", strconv.Itoa(value.(Objects.Object).Pos), "+")
	generator.SetStack(temp, "H", true)
	generator.AddExpression("H", "H", strconv.Itoa(d.Attributes.Len()), "+")

	generator.AddComment("---- Assign Value to Attributes ----")
	for i := 0; i < d.Attributes.Len(); i++ {
		tempPosObj := generator.AddTemp()
		tempRef := generator.AddTemp()
		tempAttribute := generator.AddTemp()

		generator.AddExpression(tempPosObj, "P", strconv.Itoa(value.(Objects.Object).Pos), "+")
		generator.GetStack(tempRef, tempPosObj)
		generator.AddExpression(tempAttribute, tempRef, strconv.Itoa(i), "+")

		val := d.Attributes.GetValue(i).(Objects.Attribute).Value.Compile(symbolTable, generator)
		valToSend := val.(Abstract.Value).Value

		generator.SetHeap(tempAttribute, valToSend)
	}
	return nil
}

func NewDecObjects(id string, idObject string, attributes *arrayList.List) DecObjects {
	return DecObjects{
		Id:         id,
		IdObject:   idObject,
		Attributes: attributes,
	}
}

func (d DecObjects) Execute(symbolTable SymbolTable.SymbolTable) interface{} {
	newObj := symbolTable.GetStruct(d.IdObject)
	if !symbolTable.ExistsStruct(newObj.(DecStructs.DecStruct).Id) {
		// Error - No existe el struct
	}

	if newObj.(DecStructs.DecStruct).Items.Len() != d.Attributes.Len() {
		// Error - No tiene el mismo numero de atributos
	}
	//
	object := *arrayList.New()
	newTable := SymbolTable.NewSymbolTable("Object", &symbolTable)
	for i := 0; i < d.Attributes.Len(); i++ {
		attrDefined := newObj.(DecStructs.DecStruct).Items.GetValue(i).(Expression.ItemStruct)
		attrInComing := d.Attributes.GetValue(i).(Objects.Attribute)

		if attrDefined.Id != attrInComing.Id {
			// Error - No existe ese atributo en ese orden
		}

		attr := attrInComing.GetValue(newTable)
		if attrDefined.DataType != attr.Type {
			// Error - No son los mismos tipos de datos
		}

		object.Add(attr)
	}

	objCreated := Objects.NewObject(d.Id, &object)
	objCreated.SetPos(symbolTable.SizeTable)
	symbolTable.AddObject(d.Id, objCreated)

	return SymbolTable.ReturnType{
		Value: objCreated,
		Type:  SymbolTable.OBJECT,
	}
}
