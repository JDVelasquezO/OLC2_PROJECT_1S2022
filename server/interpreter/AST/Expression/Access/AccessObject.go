package Access

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter/AST/Expression"
	"OLC2_Project1/server/interpreter/AST/Expression/Objects"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	arrayList "github.com/colegno/arraylist"
	"reflect"
	"strconv"
)

type ObjectAccess struct {
	listAccess *arrayList.List
}

func (o ObjectAccess) Compile(symbolTable *SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {
	exprInit := o.listAccess.GetValue(0)
	id := exprInit.(Expression.Identifier).Id

	object := symbolTable.GetObject(id)
	value := object.(Objects.Object)

	tempPosObj := generator.AddTemp()
	tempObj := generator.AddTemp()
	tempPosAttribute := generator.AddTemp()
	tempAttribute := generator.AddTemp()

	newList := o.listAccess.Clone()
	newList.RemoveAtIndex(0)
	attribute := o.GetAttribute(symbolTable, newList, value)
	//fmt.Println(attribute)
	typeAttr := attribute["type"].(SymbolTable.DataType)
	posAttr := attribute["pos"].(int)

	generator.AddExpression(tempPosObj, "P", strconv.Itoa(value.Pos), "+")
	generator.GetStack(tempObj, tempPosObj)
	generator.AddExpression(tempPosAttribute, tempObj, strconv.Itoa(posAttr), "+")
	generator.GetHeap(tempAttribute, tempPosAttribute)

	val := Abstract.NewValue(tempAttribute, typeAttr, true, "")
	if val.Type == SymbolTable.BOOLEAN {
		labelTrue := generator.NewLabel()
		labelFalse := generator.NewLabel()

		val.TrueLabel = labelTrue
		val.FalseLabel = labelFalse
	}
	return val
}

func (o ObjectAccess) GetAttribute(symbolTable *SymbolTable.SymbolTable, listAccess *arrayList.List,
	object Objects.Object) map[string]interface{} {
	exprInit := listAccess.GetValue(0)

	if reflect.TypeOf(exprInit) == reflect.TypeOf(Expression.Identifier{}) {
		id := exprInit.(Expression.Identifier).Id

		if !symbolTable.ExistsSymbol(id) {
			// Error - No existe objeto
		}

		var symbol interface{}
		mapToSend := make(map[string]interface{})
		for i := 0; i < object.Attributes.Len(); i++ {
			nameAttr := object.Attributes.GetValue(i).(SymbolTable.ReturnType).Value
			symbol = nameAttr.(map[string]interface{})[id]
			if symbol != nil {
				//retVal = symbol
				mapToSend["type"] = symbol.(Expression.Primitive).Type
				mapToSend["pos"] = i
				return mapToSend
			}
		}
		return nil
	}

	return nil
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

		return o.GetValueRecursion(&symbolTable, newList, value)

		//fmt.Println(object)
	}

	return SymbolTable.ReturnType{}
}

func (o ObjectAccess) GetValueRecursion(symbolTable *SymbolTable.SymbolTable, listAccess *arrayList.List, object Objects.Object) SymbolTable.ReturnType {
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

func (o ObjectAccess) GetId(symbolTable SymbolTable.SymbolTable, pos int) string {
	return o.listAccess.GetValue(pos).(Expression.Identifier).Id
}
