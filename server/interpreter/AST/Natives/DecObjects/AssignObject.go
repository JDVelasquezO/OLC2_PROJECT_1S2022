package DecObjects

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter/AST/Expression"
	"OLC2_Project1/server/interpreter/AST/Expression/Access"
	"OLC2_Project1/server/interpreter/AST/Expression/Objects"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	arrayList "github.com/colegno/arraylist"
	"reflect"
	"strconv"
)

type AssignObject struct {
	AccessObject Abstract.Expression
	Value        Abstract.Expression
}

func (a AssignObject) Compile(symbolTable *SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {
	idInit := a.AccessObject.(Access.ObjectAccess).GetId(*symbolTable, 0)
	idAttr := a.AccessObject.(Access.ObjectAccess).GetId(*symbolTable, 1)
	newObj := symbolTable.GetObject(idInit)

	generator.AddComment("---- Assign New Val Struct ----")
	tempPosObject := generator.AddTemp()
	tempValObject := generator.AddTemp()
	tempPosAttr := generator.AddTemp()

	generator.AddExpression(tempPosObject, "P", strconv.Itoa(newObj.(Objects.Object).Pos), "+")
	generator.GetStack(tempValObject, tempPosObject)

	var symbol interface{}
	var posAttr int
	for i := 0; i < newObj.(Objects.Object).Attributes.Len(); i++ {
		nameAttr := newObj.(Objects.Object).Attributes.GetValue(i).(SymbolTable.ReturnType).Value
		symbol = nameAttr.(map[string]interface{})[idAttr]
		if symbol != nil {
			posAttr = i
			break
		}
	}

	generator.AddExpression(tempPosAttr, tempValObject, strconv.Itoa(posAttr), "+")
	var val Abstract.Value
	if reflect.TypeOf(symbol) == reflect.TypeOf(Expression.Identifier{}) {
		val = symbol.(Expression.Identifier).Compile(symbolTable, generator).(Abstract.Value)
	} else {
		val = symbol.(Expression.Primitive).Compile(symbolTable, generator).(Abstract.Value)
	}

	if val.Type == SymbolTable.BOOLEAN {
		ValueBoolean(val, tempPosAttr, generator)
	} else {
		generator.SetHeap(tempPosAttr, val.Value)
	}

	return nil
}

func (a AssignObject) Execute(symbolTable SymbolTable.SymbolTable) interface{} {
	idInit := a.AccessObject.(Access.ObjectAccess).GetId(symbolTable, 0)
	idAttr := a.AccessObject.(Access.ObjectAccess).GetId(symbolTable, 1)
	newObj := symbolTable.GetObject(idInit)

	if !symbolTable.ExistsSymbol(newObj.(Objects.Object).Id) {
		// Error - No existe objeto
	}

	var symbol interface{}
	var retVal interface{}
	objRet := *arrayList.New()
	for i := 0; i < newObj.(Objects.Object).Attributes.Len(); i++ {
		nameAttr := newObj.(Objects.Object).Attributes.GetValue(i).(SymbolTable.ReturnType).Value
		symbol = nameAttr.(map[string]interface{})[idAttr]
		if symbol != nil {
			retVal = a.Value

			newAttr := make(map[string]interface{})
			newAttr[idAttr] = retVal
			var newVal SymbolTable.ReturnType
			if reflect.TypeOf(retVal) == reflect.TypeOf(Expression.Identifier{}) {
				newVal = SymbolTable.ReturnType{Value: newAttr, Type: retVal.(Expression.Identifier).GetValue(symbolTable).Type}
			} else {
				newVal = SymbolTable.ReturnType{Value: newAttr, Type: retVal.(Expression.Primitive).Type}
			}
			objRet.Add(newVal)
			//newObj.(Objects.Object).Attributes.Add(newVal)
			//break
		} else {
			oldRetType := newObj.(Objects.Object).Attributes.GetValue(i)
			objRet.Add(oldRetType)
		}
	}

	objToSend := Objects.NewObject(idInit, &objRet)
	symbolTable.ChangeValObject(idInit, objToSend)
	return nil
}

func NewAssignObject(accessObject Abstract.Expression, value Abstract.Expression) AssignObject {
	return AssignObject{
		AccessObject: accessObject,
		Value:        value,
	}
}
