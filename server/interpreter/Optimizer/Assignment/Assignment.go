package Assignment

import (
	"OLC2_Project1/server/interpreter/Optimizer/AbstractOptimizer"
	"fmt"
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
	temp := a.Temp.GetValue()

	if temp.(string) == "P" || temp.(string) == "H" {
		return temp.(string) + " = " + fmt.Sprintf("%v", opAssign) + ";"
	}

	typeRet := make(map[string]interface{})
	typeRet["temp"] = temp
	typeRet["val"] = opAssign
	typeRet["rep"] = false

	return typeRet
}
