package Expression

import (
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
)

type ValueRepetition struct {
	Expression1 Abstract.Expression
	Expression2 Abstract.Expression
}

func NewValueRepetition(exp1 Abstract.Expression, exp2 Abstract.Expression) ValueRepetition {
	return ValueRepetition{
		Expression1: exp1,
		Expression2: exp2,
	}
}

func (r ValueRepetition) GetValue(table SymbolTable.SymbolTable) SymbolTable.ReturnType {
	exp1 := r.Expression1.GetValue(table)
	exp2 := r.Expression2.GetValue(table)

	var list []int
	for i := 0; i < exp2.Value.(int); i++ {
		list = append(list, exp1.Value.(int))
	}

	return SymbolTable.ReturnType{
		Value: list,
		Type:  exp1.Type,
	}
}
