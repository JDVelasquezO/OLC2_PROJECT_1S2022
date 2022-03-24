package Access

import (
	"OLC2_Project1/server/interpreter/AST/Expression"
	"OLC2_Project1/server/interpreter/AST/Expression/Objects"
	"OLC2_Project1/server/interpreter/SymbolTable"
	arrayList "github.com/colegno/arraylist"
	"reflect"
)

type ObjectAccess struct {
	listAccess *arrayList.List
}

func (o ObjectAccess) GetValue(symbolTable SymbolTable.SymbolTable) SymbolTable.ReturnType {
	exprInit := o.listAccess.GetValue(0)

	if reflect.TypeOf(exprInit) == reflect.TypeOf(Expression.Identifier{}) {
		id := exprInit.(Expression.Identifier).Id

		if !symbolTable.ExistsSymbol(id) {
			// Error - No existe objeto
		}

		object := symbolTable.GetObject(id)
		value := object.(Objects.Object)
		newList := o.listAccess.Clone()
		newList.RemoveAtIndex(0)

		return o.GetValueRecursion(symbolTable, newList, value)

		//fmt.Println(object)
	}

	return SymbolTable.ReturnType{}
}

func (o ObjectAccess) GetValueRecursion(symbolTable SymbolTable.SymbolTable, listAccess *arrayList.List, object Objects.Object) SymbolTable.ReturnType {
	exprInit := listAccess.GetValue(0)

	if reflect.TypeOf(exprInit) == reflect.TypeOf(Expression.Identifier{}) {
		id := exprInit.(Expression.Identifier).Id

		if !symbolTable.ExistsSymbol(id) {
			// Error - No existe objeto
		}

		var symbol interface{}
		var retVal interface{}
		for i := 0; i < object.Attributes.Len(); i++ {
			nameAttr := object.Attributes.GetValue(i).(SymbolTable.ReturnType).Value
			symbol = nameAttr.(map[string]interface{})[id]
			if symbol != nil {
				retVal = symbol
				break
			}
		}
		return SymbolTable.ReturnType{
			Value: retVal.(Expression.Primitive).Value,
			Type:  retVal.(Expression.Primitive).Type,
		}
	}

	return SymbolTable.ReturnType{}
}

func NewObjectAccess(listAccess *arrayList.List) ObjectAccess {
	return ObjectAccess{
		listAccess: listAccess,
	}
}
