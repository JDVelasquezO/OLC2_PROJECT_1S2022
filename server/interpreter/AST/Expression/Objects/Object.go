package Objects

import (
	"OLC2_Project1/server/interpreter/AST/Expression"
	"OLC2_Project1/server/interpreter/AST/Natives/DecStructs"
	"OLC2_Project1/server/interpreter/SymbolTable"
	arrayList "github.com/colegno/arraylist"
)

type Object struct {
	Id         string
	Attributes *arrayList.List
}

func (o Object) GetValue(symbolTable SymbolTable.SymbolTable) SymbolTable.ReturnType {
	//var obj struct{}
	newObj := symbolTable.GetStruct(o.Id)
	if !symbolTable.ExistsStruct(newObj.(DecStructs.DecStruct).Id) {
		// Error - No existe el struct
	}

	if newObj.(DecStructs.DecStruct).Items.Len() != o.Attributes.Len() {
		// Error - No tiene el mismo numero de atributos
	}

	var object []interface{}
	newTable := SymbolTable.NewSymbolTable("Object", &symbolTable)
	for i := 0; i < o.Attributes.Len(); i++ {
		attrDefined := newObj.(DecStructs.DecStruct).Items.GetValue(i).(Expression.ItemStruct)
		attrInComing := o.Attributes.GetValue(i).(Attribute)

		if attrDefined.Id != attrInComing.Id {
			// Error - No existe ese atributo en ese orden
		}

		attr := attrInComing.GetValue(newTable)
		if attrDefined.DataType != attr.Type {
			// Error - No son los mismos tipos de datos
		}

		object = append(object, attr)
	}

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
