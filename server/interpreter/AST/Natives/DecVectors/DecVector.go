package DecVectors

import (
	"OLC2_Project1/server/interpreter/AST/Expression"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment/Array"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment/Vector"
	"fmt"
	"reflect"
)

type DecVector struct {
	Length          int
	LengthAsExpr    Abstract.Expression
	Id              string
	ValueInitial    Abstract.Expression
	Type            SymbolTable.DataType
	IsMut           bool
	Positions       Abstract.Expression
	IsCapacityFixed bool
}

func (d DecVector) GetValue(symbolTable SymbolTable.SymbolTable) SymbolTable.ReturnType {
	//TODO implement me
	panic("implement me")
}

func NewDecVector(length int, lengthAsExpr Abstract.Expression, id string, init Abstract.Expression,
	dataType SymbolTable.DataType, isMut bool, isCapacityFixed bool) DecVector {
	return DecVector{
		Length:          length,
		LengthAsExpr:    lengthAsExpr,
		Id:              id,
		ValueInitial:    init,
		Type:            dataType,
		IsMut:           isMut,
		IsCapacityFixed: isCapacityFixed,
	}
}

func (d DecVector) Execute(table SymbolTable.SymbolTable) interface{} {
	var objectVector Vector.Vector
	listValues := make([]interface{}, 0)
	objectVector.Value = listValues

	if !d.IsMut {
		objectVector.IsConst = true
	}

	if d.LengthAsExpr != nil {
		objectVector.Length = d.LengthAsExpr.GetValue(table).Value.(int)
	}

	if d.IsCapacityFixed {
		objectVector.IsCapacityFixed = true
	}

	if table.ExistsSymbol(d.Id) && d.IsMut {
		fmt.Printf("Error, variable %s ya declarada", d.Id)
	} else {
		symbol := objectVector
		symbol.Id = d.Id
		table.AddNewSymbol(d.Id, symbol)
	}

	if d.ValueInitial != nil {
		valueDec := d.ValueInitial.GetValue(table).Value
		dataType := d.ValueInitial.GetValue(table).Type
		if reflect.TypeOf(valueDec) == reflect.TypeOf(SymbolTable.ReturnType{}) {
			newArr := valueDec.(SymbolTable.ReturnType).Value.(Array.Array).Values

			for i := 0; i < len(newArr); i++ {
				newExpr := Expression.NewPrimitive(valueDec.(SymbolTable.ReturnType).Value.(Array.Array).Values[i], dataType, 0, 0)
				newPush := NewPush(d.Id, newExpr)
				newPush.ExecuteFirstTime(table)
			}
		} else {
			for i := 0; i < objectVector.Length; i++ {
				newExpr := Expression.NewPrimitive(valueDec, dataType, 0, 0)
				newPush := NewPush(d.Id, newExpr)
				newPush.ExecuteFirstTime(table)
			}
		}
	}

	return nil
}
