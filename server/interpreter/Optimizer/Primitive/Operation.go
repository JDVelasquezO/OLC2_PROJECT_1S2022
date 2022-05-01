package Primitive

import (
	"OLC2_Project1/server/interpreter/Optimizer/AbstractOptimizer"
	"fmt"
)

type Operation struct {
	OpLeft   AbstractOptimizer.Expression
	Operator string
	OpRight  AbstractOptimizer.Expression
	Row      int
	Col      int
}

func NewOperation(opLeft AbstractOptimizer.Expression, operator string, opRight AbstractOptimizer.Expression,
	row int, col int) Operation {

	return Operation{
		OpLeft:   opLeft,
		OpRight:  opLeft,
		Operator: operator,
		Row:      row,
		Col:      col,
	}
}

func (o Operation) GetValue() interface{} {
	return fmt.Sprintf("%v", o.OpLeft.GetValue()) + o.Operator + fmt.Sprintf("%v", o.OpRight.GetValue())
}
