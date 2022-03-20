package DecVectors

import (
	"OLC2_Project1/server/interpreter/AST/Expression"
	"OLC2_Project1/server/interpreter/AST/ExpressionSpecial"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment/Array"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment/Vector"
	"fmt"
	"reflect"
)

type DecVector struct {
	Length       int
	Id           string
	ValueInitial Abstract.Expression
	Type         SymbolTable.DataType
	IsMut        bool
	Positions    Abstract.Expression
}

func (d DecVector) GetValue(symbolTable SymbolTable.SymbolTable) SymbolTable.ReturnType {
	//TODO implement me
	panic("implement me")
}

func NewDecVector(length int, id string, init Abstract.Expression,
	dataType SymbolTable.DataType, isMut bool) DecVector {
	return DecVector{
		Length:       length,
		Id:           id,
		ValueInitial: init,
		Type:         dataType,
		IsMut:        isMut,
	}
}

func (d DecVector) Execute(table SymbolTable.SymbolTable) interface{} {

	var objectVector Vector.Vector
	listValues := make([]interface{}, 0)
	objectVector.Value = listValues

	if table.ExistsSymbol(d.Id) {
		fmt.Printf("Error, variable %s ya declarada", d.Id)
	} else {
		symbol := objectVector
		symbol.Id = d.Id
		table.AddNewSymbol(d.Id, symbol)
	}

	if d.ValueInitial != nil {
		valueDec := d.ValueInitial.GetValue(table).Value
		dataType := valueDec.(SymbolTable.ReturnType).Type
		newArr := valueDec.(SymbolTable.ReturnType).Value.(Array.Array).Values

		if reflect.TypeOf(d.ValueInitial) == reflect.TypeOf(ExpressionSpecial.ValueArray{}) {
			for i := 0; i < newArr[1].(int); i++ {
				newExpr := Expression.NewPrimitive(newArr[0], dataType, 0, 0)
				newPush := NewPush(d.Id, newExpr)
				newPush.Execute(table)
			}
		} else {
			for i := 0; i < len(newArr); i++ {
				newExpr := Expression.NewPrimitive(valueDec.(SymbolTable.ReturnType).Value.(Array.Array).Values[i], dataType, 0, 0)
				newPush := NewPush(d.Id, newExpr)
				newPush.Execute(table)
			}
		}
	}

	return nil
}
