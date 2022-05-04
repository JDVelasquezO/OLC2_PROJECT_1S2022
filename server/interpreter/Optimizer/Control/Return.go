package Control

import (
	"OLC2_Project1/server/interpreter/Optimizer/AbstractOptimizer"
	"fmt"
)

type Return struct {
	Row   int
	Col   int
	Value AbstractOptimizer.Expression
}

func NewReturn(value AbstractOptimizer.Expression, row int, col int) Return {
	return Return{
		Value: value,
		Row:   row,
		Col:   row,
	}
}

func (r Return) Execute() interface{} {
	if r.Value != nil {
		return "return " + fmt.Sprintf("%v", r.Value.GetValue()) + ";"
	}
	return "return 0;"
}
