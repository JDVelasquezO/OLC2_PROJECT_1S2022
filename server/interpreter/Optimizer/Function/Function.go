package Function

import (
	"OLC2_Project1/server/interpreter/Optimizer/AbstractOptimizer"
	"OLC2_Project1/server/interpreter/Optimizer/Assignment"
	"OLC2_Project1/server/interpreter/Optimizer/Control"
	"OLC2_Project1/server/interpreter/Optimizer/Primitive"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"reflect"
	"strings"
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
	// RULE 1
	newList := arrayList.New()
	newListInstr := arrayList.New()
	newListStrs := arrayList.New()
	valToSend := ""
	for i := 0; i < f.ListInstructs.Len(); i++ {
		instruction := f.ListInstructs.GetValue(i)

		newExpr := instruction.(AbstractOptimizer.Instruction).Execute()

		if reflect.TypeOf(instruction) == reflect.TypeOf(Control.Label{}) ||
			reflect.TypeOf(instruction) == reflect.TypeOf(Control.If{}) ||
			reflect.TypeOf(instruction) == reflect.TypeOf(Control.GoTo{}) {
			valToSend = "\t" + newExpr.(string) + "\n"
			newListStrs.Add(valToSend)
			continue
		}

		if reflect.TypeOf(instruction) != reflect.TypeOf(Control.Label{}) {
			tempNew := newExpr.(map[string]interface{})
			if newList.Len() > 0 {
				limit := newList.Len()
				for j := 0; j < limit; j++ {
					tempToCmp := newList.GetValue(j).(map[string]interface{})
					if tempNew["val"] == tempToCmp["val"] {
						tempNew["val"] = tempToCmp["temp"]
						tempNew["rep"] = true
						//newList.Add(newExpr)
						break
					}
				}
			}
			newList.Add(newExpr)
			valToSend = "\t" + fmt.Sprintf("%v", tempNew["temp"]) + " = " + fmt.Sprintf("%v", tempNew["val"]) + ";\n"
			newListStrs.Add(valToSend)

			newExpTemp := Primitive.NewTemp(fmt.Sprintf("%v", tempNew["temp"]), 0, 0)
			newExpVal := Primitive.NewTemp(fmt.Sprintf("%v", tempNew["val"]), 0, 0)
			newInstr := Assignment.NewAssign(newExpTemp, newExpVal, 0, 0)
			newListInstr.Add(newInstr)
		}
	}

	// RULE 2
	newList2 := arrayList.New()
	for i := 0; i < newListInstr.Len(); i++ {
		instruction := newListInstr.GetValue(i)

		newExpr := instruction.(AbstractOptimizer.Instruction).Execute()

		if reflect.TypeOf(instruction) != reflect.TypeOf(Control.Label{}) {

			tempNew := newExpr.(map[string]interface{})
			if newList2.Len() > 0 {
				limit := newList2.Len()
				for j := limit - 1; j >= 0; j-- {
					tempToCmp := newList2.GetValue(j).(map[string]interface{})

					valOfNew := tempNew["val"].(string)
					valOfCmp := tempToCmp["temp"].(string)
					valRightOfCmp := tempToCmp["val"].(string)

					valOfNewArray := Splitter(valOfNew, "+ - * /")
					valOfCmpArray := Splitter(valRightOfCmp, "+ - * /")

					if len(valOfNewArray) == 2 {
						for k := 0; k < 2; k++ {
							if valOfNewArray[k] == valOfCmp && len(valOfCmpArray) == 1 {

								valToReplace := newList2.GetValue(j).(map[string]interface{})["val"].(string)
								newVal := strings.ReplaceAll(valOfNew, valOfNewArray[k], valToReplace)
								tempNew["val"] = newVal
								valToSend = "\t" + tempNew["temp"].(string) + " = " + tempNew["val"].(string) + ";\n"

								strToCmp1 := "\t" + tempNew["temp"].(string) + " = " + valOfNew + ";\n"
								if newListStrs.Contains(strToCmp1) {
									//index := newListStrs.IndexOf(strToCmp1)
									//newListStrs.RemoveAtIndex(index)
									newListStrs.ReplaceAll(strToCmp1, valToSend)
								}
								//newList2.RemoveAtIndex(newList2.Len() - 1)

								strToCmp2 := "\t" + valOfCmp + " = " + valRightOfCmp + ";\n"
								if newListStrs.Contains(strToCmp2) {
									index := newListStrs.IndexOf(strToCmp2)
									newListStrs.RemoveAtIndex(index)
								}
								newList2.RemoveAtIndex(j)

								newList2.Add(newExpr)
								goto ContinueLabel
							}
						}
					}
				}
			ContinueLabel:
			}
			newList2.Add(newExpr)
		}
	}

	valToSend = ""
	for i := 0; i < newListStrs.Len(); i++ {
		valToSend += newListStrs.GetValue(i).(string)
	}

	return valToSend
}

func Splitter(s string, splits string) []string {
	m := make(map[rune]int)
	for _, r := range splits {
		m[r] = 1
	}

	splitter := func(r rune) bool {
		return m[r] == 1
	}

	return strings.FieldsFunc(s, splitter)
}
