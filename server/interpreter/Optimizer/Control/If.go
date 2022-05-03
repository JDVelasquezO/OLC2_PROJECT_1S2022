package Control

import (
	"OLC2_Project1/server/interpreter/Optimizer/AbstractOptimizer"
)

type If struct {
	Condition AbstractOptimizer.Expression
	Row       int
	Col       int
	label     string
}

func NewIf(condition AbstractOptimizer.Expression, label string, row int, col int) If {
	return If{
		Condition: condition,
		label:     label,
		Row:       row,
		Col:       col,
	}
}

func (i If) Execute() interface{} {
	return "if (" + i.Condition.GetValue().(string) + ") goto " + i.label + ";"
}
