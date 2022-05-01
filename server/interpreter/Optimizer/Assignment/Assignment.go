package Assignment

import (
	"OLC2_Project1/server/interpreter/Optimizer/AbstractOptimizer"
)

type Assign struct {
	Row          int
	Col          int
	Value        AbstractOptimizer.Expression
	Temp         AbstractOptimizer.Expression
	InstrDeleted bool
}

func NewAssign(temp AbstractOptimizer.Expression, value AbstractOptimizer.Expression, row int, col int) Assign {
	return Assign{
		Value: value,
		Temp:  temp,
		Row:   row,
		Col:   col,
	}
}

func (a Assign) Execute() interface{} {

	opAssign := a.Value.GetValue()

	return opAssign
}
