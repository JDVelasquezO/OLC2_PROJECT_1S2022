package Function

import (
	"OLC2_Project1/server/interpreter/Optimizer/AbstractOptimizer"
	"OLC2_Project1/server/interpreter/Optimizer/Assignment"
	"OLC2_Project1/server/interpreter/Optimizer/Control"
	"OLC2_Project1/server/interpreter/Optimizer/Primitive"
	"OLC2_Project1/server/interpreter/rules"
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

		if reflect.TypeOf(newExpr).String() == "string" {
			valToKeep := "\t" + newExpr.(string) + "\n"
			newListStrs.Add(valToKeep)
			continue
		} else {
			objExpr := newExpr.(map[string]interface{})
			if strings.Contains(fmt.Sprintf("%v", objExpr["val"]), "H") ||
				strings.Contains(fmt.Sprintf("%v", objExpr["val"]), "P") {

				valToKeep := "\t" + objExpr["temp"].(string) + " = " + objExpr["val"].(string) + ";\n"
				newListStrs.Add(valToKeep)
				continue
			}
		}

		if reflect.TypeOf(instruction) != reflect.TypeOf(Control.Label{}) {
			tempNew := newExpr.(map[string]interface{})
			if newList.Len() > 0 {
				limit := newList.Len()
				for j := 0; j < limit; j++ {
					tempToCmp := newList.GetValue(j).(map[string]interface{})
					if tempNew["val"] == tempToCmp["val"] {
						valToSend = "\t/*---- Se aplico regla 1 ----*/\n"
						newListStrs.Add(valToSend)

						original := tempNew["temp"].(string) + " = " + tempNew["val"].(string)
						tempNew["val"] = tempToCmp["temp"]
						tempNew["rep"] = true
						optimized := tempNew["temp"].(string) + " = " + tempNew["val"].(string)
						row := tempNew["row"].(int)
						rule := rules.NewRule("Regla 1", original, optimized, row)
						rules.CounterRule += 1
						rules.TypeRule = append(rules.TypeRule, rule)

						break
					}
				}
			}
			newList.Add(newExpr)
			valToSend = "\t" + fmt.Sprintf("%v", tempNew["temp"]) + " = " + fmt.Sprintf("%v", tempNew["val"]) + ";\n"
			newListStrs.Add(valToSend)

			newExpTemp := Primitive.NewTemp(fmt.Sprintf("%v", tempNew["temp"]), tempNew["row"].(int), 0)
			newExpVal := Primitive.NewTemp(fmt.Sprintf("%v", tempNew["val"]), tempNew["row"].(int), 0)
			newInstr := Assignment.NewAssign(newExpTemp, newExpVal, tempNew["row"].(int), 0)
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

								original := tempNew["temp"].(string) + " = " + tempNew["val"].(string)
								tempNew["val"] = newVal
								optimized := tempNew["temp"].(string) + " = " + tempNew["val"].(string)
								row := tempNew["row"].(int)
								rule := rules.NewRule("Regla 2 o 4", original, optimized, row)
								rules.CounterRule += 1
								rules.TypeRule = append(rules.TypeRule, rule)

								commentToSend := "\t/*---- Se aplico regla (2 o 4) y 3 ----*/\n"
								valToSend = commentToSend + "\t" + tempNew["temp"].(string) + " = " + tempNew["val"].(string) + ";\n"

								strToCmp1 := "\t" + tempNew["temp"].(string) + " = " + valOfNew + ";\n"
								if newListStrs.Contains(strToCmp1) {
									newListStrs.ReplaceAll(strToCmp1, valToSend)
								}

								strToCmp2 := "\t" + valOfCmp + " = " + valRightOfCmp + ";\n"
								if newListStrs.Contains(strToCmp2) {
									index := newListStrs.IndexOf(strToCmp2)
									newListStrs.RemoveAtIndex(index)

									original := valOfCmp + " = " + valRightOfCmp
									rule := rules.NewRule("Regla 3", original, "desaparece", tempToCmp["row"].(int))
									rules.CounterRule += 1
									rules.TypeRule = append(rules.TypeRule, rule)

									for l := 0; l < newListStrs.Len(); l++ {
										strToValue := newListStrs.GetValue(l)
										if strings.Contains(strToValue.(string), valOfCmp) {
											strChanged := strings.ReplaceAll(strToValue.(string), valOfCmp, valRightOfCmp)
											newListStrs.ReplaceAll(strToValue, strChanged)
											continue
										}
									}
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
	if newListStrs.Len() > 0 {
		for i := 0; i < newListStrs.Len(); i++ {
			valToSend += newListStrs.GetValue(i).(string)
		}
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
