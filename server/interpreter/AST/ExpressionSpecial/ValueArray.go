package ExpressionSpecial

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment/Array"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type ValueArray struct {
	Expressions *arrayList.List
	Exp1        Abstract.Expression
	Exp2        Abstract.Expression
}

func (v ValueArray) Compile(symbolTable *SymbolTable.SymbolTable, generator *Generator.Generator) interface{} {
	return nil
}

func NewValueArray(expressions *arrayList.List, exp1 Abstract.Expression, exp2 Abstract.Expression) ValueArray {
	return ValueArray{
		Expressions: expressions,
		Exp1:        exp1,
		Exp2:        exp2,
	}
}

func (v ValueArray) GetValue(table SymbolTable.SymbolTable) SymbolTable.ReturnType {
	data, dtype := v.GetData(table)
	return SymbolTable.ReturnType{
		Value: data,
		Type:  dtype,
	}
}

func (v ValueArray) GetData(table SymbolTable.SymbolTable) (interface{}, SymbolTable.DataType) {
	dtype := SymbolTable.NULL
	data := arrayList.New()

	if v.Expressions != nil {
		for i := 0; i < v.Expressions.Len(); i++ {
			expr := v.Expressions.GetValue(i).(Abstract.Expression)
			valExpr := expr.GetValue(table)

			if i == 0 {
				dtype = valExpr.Type
				data.Add(valExpr)
			} else {
				if dtype != valExpr.Type {
					fmt.Println("Error1")
					return nil, SymbolTable.NULL
				}
				data.Add(valExpr)
			}
		}
	}

	if v.Exp2 != nil {
		expr := v.Exp1.GetValue(table)
		for i := 0; i < v.Exp2.GetValue(table).Value.(int); i++ {
			valExpr := expr

			if i == 0 {
				dtype = valExpr.Type
				data.Add(valExpr)
			} else {
				if dtype != valExpr.Type {
					fmt.Println("Error1")
					return nil, SymbolTable.NULL
				}
				data.Add(valExpr)
			}
		}
	}

	listDimensions := arrayList.New()
	listDimensions.Add(data.Len())

	dataType := SymbolTable.NULL
	s := make([]interface{}, data.Len())

	for i := 0; i < data.Len(); i++ {
		d := data.GetValue(i)
		valueData := d.(SymbolTable.ReturnType)

		if valueData.Type == SymbolTable.NULL {
			fmt.Println("Error2")
			return nil, SymbolTable.NULL
		}

		if valueData.Type != SymbolTable.ARRAY {
			if i == 0 {
				dataType = valueData.Type
			} else {
				if dataType != valueData.Type {
					fmt.Println("Error4")
					return nil, SymbolTable.NULL
				}
			}
			s[i] = valueData.Value
			continue
		}

		valObject := valueData.Value.(SymbolTable.ReturnType)
		objectArray := valObject.Value.(Array.Array)

		if i == 0 {
			dataType = valObject.Type
			listDimensions.AddAll(objectArray.ListIntDim.ToArray())
		} else {
			if dataType != valObject.Type {
				fmt.Println("Error5")
				return nil, SymbolTable.NULL
			}
		}

		s[i] = objectArray.Values
	}

	object := Array.NewArray("", SymbolTable.INTEGER, s, listDimensions)
	objectVal := SymbolTable.ReturnType{
		Value: object,
		Type:  dataType,
	}

	return objectVal, SymbolTable.ARRAY
}

func (v ValueArray) resizeArray(data SymbolTable.DataType) bool {
	switch data {
	case SymbolTable.INTEGER:
	case SymbolTable.BOOLEAN:
	case SymbolTable.STRING:
	case SymbolTable.FLOAT:
		return false
	}
	return true
}
