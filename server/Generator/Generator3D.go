package Generator

import "github.com/colegno/arraylist"

type Generator struct {
	Temps *arraylist.List
	Code  string
}

func NewGenerator(code string) Generator {

	temps := arraylist.New()
	generator := Generator{
		Temps: temps,
		Code:  "",
	}

	return generator
}

func (g Generator) CleanAll() {
	g.Code = ""
	g.Temps = arraylist.New()
}

func (g Generator) GetCode() string {
	g.Code = initialHeader(g)
	g.Code += "\n/*------MAIN------*/\n void main() { \n\tP = 1; H = 0;\n\n\t return; \n }"
	return g.Code
}

func initialHeader(g Generator) string {
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
