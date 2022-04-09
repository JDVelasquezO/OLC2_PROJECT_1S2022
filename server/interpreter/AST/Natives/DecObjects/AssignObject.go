package DecObjects

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter/AST/Expression"
	"OLC2_Project1/server/interpreter/AST/Expression/Access"
	"OLC2_Project1/server/interpreter/AST/Expression/Objects"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	arrayList "github.com/colegno/arraylist"
)

type AssignObject struct {
	AccessObject Abstract.Expression
	Value        Abstract.Expression
}

func (a AssignObject) Compile(symbolTable *SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {
	//TODO implement me
	panic("implement me")
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
			newVal := SymbolTable.ReturnType{Value: newAttr, Type: retVal.(Expression.Primitive).Type}
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
