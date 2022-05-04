package Assignment

import (
	"OLC2_Project1/server/interpreter/Optimizer/AbstractOptimizer"
	"fmt"
)

type GetHeapStack struct {
	Row       int
	Col       int
	Value     AbstractOptimizer.Expression
	Temp      AbstractOptimizer.Expression
	Structure string
}

func NewGetHeapStack(temp AbstractOptimizer.Expression, value AbstractOptimizer.Expression,
	row int, col int, structure string) GetHeapStack {
	return GetHeapStack{
		Value:     value,
		Temp:      temp,
		Row:       row,
		Col:       col,
		Structure: structure,
	}
}

func (g GetHeapStack) Execute() interface{} {

	opAssign := g.Value.GetValue()
	temp := g.Temp.GetValue()

	return fmt.Sprintf("%v", temp) + " = " + g.Structure + "[(int)" + fmt.Sprintf("%v", opAssign) + "];"
}
