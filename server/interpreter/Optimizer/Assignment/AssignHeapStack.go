package Assignment

import (
	"OLC2_Project1/server/interpreter/Optimizer/AbstractOptimizer"
	"fmt"
)

type AssignHeapStack struct {
	Row          int
	Col          int
	Value        AbstractOptimizer.Expression
	Temp         AbstractOptimizer.Expression
	InstrDeleted bool
	Structure    string
}

func NewAssignHeapStack(temp AbstractOptimizer.Expression, value AbstractOptimizer.Expression,
	row int, col int, structure string) AssignHeapStack {
	return AssignHeapStack{
		Value:     value,
		Temp:      temp,
		Row:       row,
		Col:       col,
		Structure: structure,
	}
}

func (a AssignHeapStack) Execute() interface{} {

	opAssign := a.Value.GetValue()
	temp := a.Temp.GetValue()

	return a.Structure + "[(int)" + fmt.Sprintf("%v", opAssign) + "] = " + fmt.Sprintf("%v", temp) + ";"
}
