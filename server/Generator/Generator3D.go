package Generator

import (
	"OLC2_Project1/server/interpreter/SymbolTable"
	"github.com/colegno/arraylist"
	"reflect"
	"strconv"
)

type Generator struct {
	Temps          *arraylist.List
	Code           string
	CountLabel     int
	TempsRecovered map[string]interface{}
	InNatives      bool
	InFunc         bool
	Natives        string
	Functions      string
}

func NewGenerator(code string) Generator {

	temps := arraylist.New()
	generator := Generator{
		Temps:          temps,
		Code:           "",
		CountLabel:     0,
		TempsRecovered: make(map[string]interface{}),
		InNatives:      false,
		InFunc:         false,
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
	code += "\n/*------MAIN------*/\n void main() { \n\tP = 1; H = 0;\n\n" + g.Code + "\n\t return; \n }"
	return code
}

func initialHeader(g *Generator) string {
	header := "/*------HEADER------*/\n"
	header += "#include <stdio.h>\n"
	header += "#include <math.h>\n"
	header += "float heap[30101999];\n"
	header += "float stack[30101999];\n"
	header += "float P;\n"
	header += "float H;\n"

	if g.Temps.Len() > 0 {
		header += "float "
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

func (g *Generator) SetStack(pos int, value interface{}, freeValue bool) {
	g.GetFreeTemp(strconv.Itoa(pos))
	if freeValue {
		if typeof(value) == "int" {
			g.GetFreeTemp(strconv.Itoa(value.(int)))
			g.CodeInFunction("stack[(int)"+strconv.Itoa(pos)+"] = "+strconv.Itoa(value.(int))+";\n", "\t")
		} else {
			g.GetFreeTemp(value.(string))
			g.CodeInFunction("stack[(int)"+strconv.Itoa(pos)+"] = "+value.(string)+";\n", "\t")
		}
	}
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}
