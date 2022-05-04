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
		OpRight:  opRight,
		Operator: operator,
		Row:      row,
		Col:      col,
	}
}

func (o Operation) GetValue() interface{} {
	opLeft := fmt.Sprintf("%v", o.OpLeft.GetValue())
	if o.OpRight == nil {
		return o.Operator + opLeft
	}
	opRight := fmt.Sprintf("%v", o.OpRight.GetValue())
	return opLeft + " " + o.Operator + " " + opRight
}
