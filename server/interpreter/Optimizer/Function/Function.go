package Function

import (
	"OLC2_Project1/server/interpreter/Optimizer/AbstractOptimizer"
	"OLC2_Project1/server/interpreter/Optimizer/Control"
	arrayList "github.com/colegno/arraylist"
	"reflect"
)

type Function struct {
	Id            string
	ListInstructs *arrayList.List
	Line          int
	Col           int
}

func NewFunction(line int, col int, name string, listInstructs *arrayList.List) Function {

	return Function{
		Id:            name,
		ListInstructs: listInstructs,
		Line:          line,
		Col:           col,
	}
}

func (f Function) Execute() interface{} {
	for i := 0; i < f.ListInstructs.Len(); i++ {
		instruction := f.ListInstructs.GetValue(i)

		if reflect.TypeOf(instruction) != reflect.TypeOf(Control.If{}) ||
			reflect.TypeOf(instruction) != reflect.TypeOf(Control.GoTo{}) {

			instruction.(AbstractOptimizer.Instruction).Execute()
		}
	}
	return nil
}
