package Generator

import (
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/errors"
	"fmt"
	"github.com/colegno/arraylist"
	"reflect"
	"strconv"
)

type Generator struct {
	Temps          *arraylist.List
	Code           string
	CountLabel     int
	CountTemp      int
	TempsRecovered map[string]interface{}
	InNatives      bool
	InFunc         bool
	PrintStr       bool
	ConcatStr      bool
	CompareStr     bool
	Natives        string
	Functions      string
	Errors         *arraylist.List
}

func NewGenerator(code string) Generator {

	temps := arraylist.New()
	errors := arraylist.New()
	generator := Generator{
		Temps:          temps,
		Code:           "",
		CountLabel:     0,
		CountTemp:      0,
		TempsRecovered: make(map[string]interface{}),
		Errors:         errors,
		InNatives:      false,
		InFunc:         false,
		PrintStr:       false,
		ConcatStr:      false,
		CompareStr:     false,
		Natives:        "",
		Functions:      "",
	}

	return generator
}

func (g *Generator) CleanAll() {
	g.Code = ""
	g.Temps = arraylist.New()
}

func (g *Generator) GetCode() string {
	code := initialHeader(g)
	code += g.Natives
	code += g.Functions
	code += "\n/*------MAIN------*/\n void main() { \n\tP = 0; H = 0;\n\n" + g.Code + "\n\t return; \n }"
	return code
}

func initialHeader(g *Generator) string {
	header := "/*------HEADER------*/\n"
	header += "#include <stdio.h>\n"
	header += "float heap[30101999];\n"
	header += "float stack[30101999];\n"
	header += "float P;\n"
	header += "float H;\n"

	if g.Temps.Len() > 0 {
		header += "float "
		for i := 0; i < g.Temps.Len(); i++ {
			header += g.Temps.GetValue(i).(string)
			if i != (g.Temps.Len() - 1) {
				header += ", "
			}
		}
		header += ";\n"
	}

	return header
}

func (g *Generator) NewLabel() string {
	label := "L" + strconv.Itoa(g.CountLabel)
	g.CountLabel += 1
	return label
}

func (g *Generator) SetLabel(label string) {
	g.CodeInFunction(label+":\n", "\t")
}

func (g *Generator) FreeAllTemps() {
	g.TempsRecovered = make(map[string]interface{})
}

func (g *Generator) SaveTemps(table SymbolTable.Symbol) int {
	size := 0
	if len(g.TempsRecovered) > 0 {
		temp := g.AddTemp()
		g.FreeTemp(temp)

		g.AddComment("---- Start Keep Temps ----")
		g.AddExpression(temp, "P", strconv.Itoa(table.Size), "+")
		for val := range g.TempsRecovered {
			size += 1
			g.SetStack(temp, val, false)
			if size != len(g.TempsRecovered) {
				g.AddExpression(temp, temp, "1", "+")
			}
		}
		g.AddComment("---- End Keep Temps ----")
	}

	ptr := table.Size
	table.Size = ptr + size
	return ptr
}

func (g *Generator) RecoverTemps(env SymbolTable.Symbol, pos int) {
	if len(g.TempsRecovered) > 0 {
		temp := g.AddTemp()
		g.FreeTemp(temp)

		size := 0

		g.AddComment("---- Start Recover Temps ----")
		g.AddExpression(temp, "P", strconv.Itoa(pos), "+")
		for value := range g.TempsRecovered {
			size += 1
			g.GetStack(value, temp)
			if size != len(g.TempsRecovered) {
				g.AddExpression(temp, temp, "1", "+")
			}

		}
		env.Size = pos
		g.AddComment("---- Fin Recover Temps ----")
	}
}

func (g *Generator) FreeTemp(temp string) {
	if _, found := g.TempsRecovered[temp]; found {
		delete(g.TempsRecovered, temp)
	}
}

func (g *Generator) AddBeginFunc(id string, dataType SymbolTable.DataType) {
	if !g.InNatives {
		g.InFunc = true
	}
	g.CodeInFunction("void "+id+"() {\n", "")
}

func (g *Generator) AddEndFunc() {
	g.CodeInFunction("return; \n}\n", "\t")
	if !g.InNatives {
		g.InFunc = false
	}
}

func (g *Generator) CodeInFunction(code string, tab string) {
	if g.InNatives {
		if g.Natives == "" {
			g.Natives = g.Natives + "\n/*------NATIVES------*/\n"
		}
		g.Natives = g.Natives + tab + code
	} else if g.InFunc {
		if g.Functions == "" {
			g.Functions = g.Functions + "\n/*-----FUNCTIONS-----*/\n"
		}
		g.Functions = g.Functions + tab + code
	} else {
		g.Code = g.Code + "\t" + code
	}
}

func (g *Generator) GetFreeTemp(temp string) {
	if temp == g.TempsRecovered[temp] {
		delete(g.TempsRecovered, temp)
	}
}

func (g *Generator) SetStack(pos interface{}, value interface{}, freeValue bool) {
	if freeValue {
		if typeof(value) == "int" {
			g.GetFreeTemp(strconv.Itoa(value.(int)))
			//newPos := strconv.Itoa(pos.(int))
			g.CodeInFunction("stack[(int)"+pos.(string)+"] = "+strconv.Itoa(value.(int))+";\n", "\t")
		} else if typeof(value) == "float64" {
			g.GetFreeTemp(fmt.Sprintf("%v", value))
			g.CodeInFunction("stack[(int)"+pos.(string)+"] = "+fmt.Sprintf("%v", value)+";\n", "\t")
		} else {
			g.GetFreeTemp(value.(string))
			if len(value.(string)) > 1 || value.(string) == "H" {
				g.CodeInFunction("stack[(int)"+pos.(string)+"] = "+value.(string)+";\n", "\t")
			} else {
				g.CodeInFunction("stack[(int)"+pos.(string)+"] = '"+value.(string)+"';\n", "\t")
			}
		}
	}
}

func (g *Generator) GetStack(place string, pos string) {
	g.GetFreeTemp(pos)
	g.CodeInFunction(place+"=stack[(int)"+pos+"];\n", "\t")
}

func (g *Generator) GetHeap(place string, pos string) {
	g.GetFreeTemp(pos)
	g.CodeInFunction(place+"=heap[(int)"+pos+"];\n", "\t")
}

func (g *Generator) SetHeap(pos interface{}, value interface{}) {
	g.GetFreeTemp(pos.(string))
	if typeof(value) == "int" {
		g.GetFreeTemp(strconv.Itoa(value.(int)))
		g.CodeInFunction("heap[(int)"+pos.(string)+"]="+strconv.Itoa(value.(int))+";\n", "\t")
	} else if typeof(value) == "float64" {
		g.GetFreeTemp(fmt.Sprintf("%v", value))
		g.CodeInFunction("heap[(int)"+pos.(string)+"]="+fmt.Sprintf("%v", value)+";\n", "\t")
	} else {
		g.GetFreeTemp(value.(string))
		if len(value.(string)) == 1 {
			g.CodeInFunction("heap[(int)"+pos.(string)+"]='"+value.(string)+"';\n", "\t")
		} else {
			g.CodeInFunction("heap[(int)"+pos.(string)+"]="+value.(string)+";\n", "\t")
		}
	}
}

func (g *Generator) NextHeap() {
	g.CodeInFunction("H = H + 1;\n", "\t")
}

func (g *Generator) AddGoTo(label string) {
	g.CodeInFunction("goto "+label+";\n", "\t")
}

func (g *Generator) AddTemp() string {
	temp := "t" + strconv.Itoa(g.CountTemp)
	g.CountTemp += 1
	g.Temps.Add(temp)
	g.TempsRecovered[temp] = temp
	return temp
}

func (g *Generator) AddExpression(res string, left string, right string, ope string) {
	g.GetFreeTemp(left)
	g.GetFreeTemp(right)
	var newCode string
	if right == "" && ope == "" {
		newCode = res + " = " + left + ";\n"
	} else {
		newCode = res + " = " + left + " " + ope + " " + right + ";\n"
	}
	g.CodeInFunction(newCode, "\t")
}

func (g *Generator) AddIf(left string, right string, ope string, label string) {
	g.GetFreeTemp(left)
	g.GetFreeTemp(right)
	g.CodeInFunction("if ("+left+" "+ope+" "+right+") goto "+label+";\n", "\t")
}

func (g *Generator) AddPrint(format string, dataType string, value string) {
	g.GetFreeTemp(value)
	g.CodeInFunction("printf(\"%"+format+"\", ("+dataType+")"+value+");\n", "\t")
}

func (g *Generator) AddError(msg string, line int, col int) {
	g.Errors.Add(errors.NewError(0, line, col, msg, "local"))
}

func (g *Generator) PrintString() {
	if g.PrintStr {
		return
	}

	g.PrintStr = true
	g.InNatives = true
	g.AddBeginFunc("print", SymbolTable.VOID)
	retLbl1 := g.NewLabel()
	cmpLbl1 := g.NewLabel()
	tempP := g.AddTemp()
	tempH := g.AddTemp()
	g.AddExpression(tempP, "P", "1", "+")
	g.GetStack(tempH, tempP)

	tempCmp := g.AddTemp()
	g.SetLabel(cmpLbl1)
	g.GetHeap(tempCmp, tempH)
	g.AddIf(tempCmp, "-1", "==", retLbl1)

	g.AddPrint("c", "char", tempCmp)
	g.AddExpression(tempH, tempH, "1", "+")
	g.AddGoTo(cmpLbl1)
	g.SetLabel(retLbl1)
	g.AddEndFunc()
	g.InNatives = false
	g.GetFreeTemp(tempP)
	g.GetFreeTemp(tempH)
	g.GetFreeTemp(tempCmp)
}

func (g *Generator) ConcatString() {
	if g.ConcatStr {
		return
	}

	g.ConcatStr = true
	g.InNatives = true
	g.AddBeginFunc("concat", SymbolTable.STRING)

	t2 := g.AddTemp()
	t3 := g.AddTemp()
	t4 := g.AddTemp()
	t5 := g.AddTemp()

	g.AddExpression(t2, "H", "", "")
	g.AddExpression(t3, "P", "1", "+")
	g.GetStack(t5, t3)
	g.AddExpression(t4, "P", "2", "+")

	l1 := g.NewLabel()
	l2 := g.NewLabel()

	g.SetLabel(l1)
	t6 := g.AddTemp()
	g.GetHeap(t6, t5)

	g.AddIf(t6, "-1", "==", l2)
	g.SetHeap("H", t6)
	g.NextHeap()
	g.AddExpression(t5, t5, "1", "+")
	g.AddGoTo(l1)

	g.SetLabel(l2)
	g.GetStack(t5, t4)

	l0 := g.NewLabel()
	l3 := g.NewLabel()
	g.SetLabel(l3)
	g.GetHeap(t6, t5)

	g.AddIf(t6, "-1", "==", l0)
	g.SetHeap("H", t6)
	g.NextHeap()
	g.AddExpression(t5, t5, "1", "+")
	g.AddGoTo(l3)

	g.SetLabel(l0)
	g.SetHeap("H", "-1")
	g.NextHeap()
	g.SetStack("P", t2, true)

	g.AddEndFunc()
	g.InNatives = false
	g.GetFreeTemp(t2)
	g.GetFreeTemp(t3)
	g.GetFreeTemp(t4)
	g.GetFreeTemp(t5)
	g.GetFreeTemp(t6)
}

func (g *Generator) CompareString() {
	if g.CompareStr {
		return
	}

	g.CompareStr = true
	g.InNatives = true
	g.AddBeginFunc("cmp", SymbolTable.STRING)

	t2 := g.AddTemp()
	g.AddExpression(t2, "P", "1", "+")
	t3 := g.AddTemp()
	g.GetStack(t3, t2)
	g.AddExpression(t2, t2, "1", "+")
	t4 := g.AddTemp()
	g.GetStack(t4, t2)

	l0 := g.NewLabel()
	l1 := g.NewLabel()
	l2 := g.NewLabel()
	l3 := g.NewLabel()

	g.SetLabel(l1)
	t5 := g.AddTemp()
	g.GetHeap(t5, t3)
	t6 := g.AddTemp()
	g.GetHeap(t6, t4)

	g.AddIf(t5, t6, "!=", l3)
	g.AddIf(t5, "-1", "==", l2)
	g.AddExpression(t3, t3, "1", "+")
	g.AddExpression(t4, t4, "1", "+")
	g.AddGoTo(l1)

	g.SetLabel(l2)
	g.SetStack("P", "1", true)
	g.AddGoTo(l0)

	g.SetLabel(l3)
	g.SetStack("P", "0", true)
	g.SetLabel(l0)

	g.AddEndFunc()
	g.InNatives = false
	g.GetFreeTemp(t2)
	g.GetFreeTemp(t3)
	g.GetFreeTemp(t4)
	g.GetFreeTemp(t5)
	g.GetFreeTemp(t6)
}

func (g *Generator) NewEnv(size int) {
	g.CodeInFunction("P = P + "+strconv.Itoa(size)+";\n", "\t")
}

func (g *Generator) SetEnv(size int) {
	g.CodeInFunction("P = P - "+strconv.Itoa(size)+";\n", "\t")
}

func (g *Generator) SetTemp(temp string, size int) {
	g.CodeInFunction(temp+" = P + "+strconv.Itoa(size)+";\n", "\t")
}

func (g *Generator) CallFunc(id string) {
	g.CodeInFunction(id+"();\n", "\t")
}

func (g *Generator) AddOperationMod(res string, left string, right string) {
	g.CodeInFunction(res+"=(int)"+left+" % "+right+";\n", "\t")
}

func (g *Generator) AddOperationPow(res string, base string, exponent string) {
	g.CodeInFunction(res+"=pow("+base+", "+exponent+");\n", "\t")
}

func (g *Generator) AddComment(comment string) {
	g.CodeInFunction("/* "+comment+" */\n", "\t")
}

func (g *Generator) ToInt(res string, exp string) {
	g.CodeInFunction(res+"=(int)"+exp+";\n", "\t")
}

func (g *Generator) ToFloat(res string, exp string) {
	g.CodeInFunction(res+"=(double)"+exp+";\n", "\t")
}

func (g *Generator) PrintTrue() {
	g.AddPrint("c", "char", "116")
	g.AddPrint("c", "char", "114")
	g.AddPrint("c", "char", "117")
	g.AddPrint("c", "char", "101")
}

func (g *Generator) PrintFalse() {
	g.AddPrint("c", "char", "102")
	g.AddPrint("c", "char", "97")
	g.AddPrint("c", "char", "108")
	g.AddPrint("c", "char", "115")
	g.AddPrint("c", "char", "101")
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}
