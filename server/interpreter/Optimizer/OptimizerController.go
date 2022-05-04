package Optimizer

import (
	"OLC2_Project1/server/interpreter"
	"OLC2_Project1/server/interpreter/Optimizer/AbstractOptimizer"
	"OLC2_Project1/server/interpreter/Optimizer/Analyzer/parser"
	"OLC2_Project1/server/interpreter/Optimizer/Function"
	"OLC2_Project1/server/utilities"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type CodeToSend struct {
	Code string `form:"code"`
}

func Optimize(c *fiber.Ctx) error {

	data := new(CodeToSend)
	e := c.BodyParser(data)
	if e != nil {
		return e
	}
	fmt.Println(data.Code)

	codeToSend := antlr.NewInputStream(data.Code)

	lexicalErrors := &utilities.CustomErrorListener{}
	lexer := parser.NewOptimizerLexer(codeToSend)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexicalErrors)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parseErrors := &utilities.CustomErrorListener{}
	p := parser.NewOptimizerParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(parseErrors)

	p.BuildParseTrees = true
	tree := p.Start()

	if len(parseErrors.Errors) > 0 {
		fmt.Printf("Parser %d errors found\n", len(parseErrors.Errors))
		for _, e := range parseErrors.Errors {
			row := strconv.Itoa(e.Row)
			col := strconv.Itoa(e.Column)
			interpreter.Console += "(" + row + ", " + col + ") " + e.Msg + "\n"
			fmt.Println("\t", e)
		}
	}

	var listener = utilities.NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	ast := tree.GetTree()

	resToSend := "/*------HEADER------*/\n#include <stdio.h>\nfloat heap[30101999];\nfloat stack[30101999];\nfloat P;\nfloat H;\n"

	resToSend += "float "
	for i := 0; i < ast.ListTemps.Len(); i++ {
		if i < ast.ListTemps.Len()-1 {
			resToSend += ast.ListTemps.GetValue(i).(string) + ", "
		} else {
			resToSend += ast.ListTemps.GetValue(i).(string)
		}
	}
	resToSend += ";\n\n"

	//resToSend += "/*------MAIN------*/\nvoid main() {\n"
	if ast.ListInstr != nil {
		for i := 0; i < ast.ListInstr.Len(); i++ {
			r := ast.ListInstr.GetValue(i)
			if r != nil {
				resToSend += "void " + r.(Function.Function).Id + "() {\n"
				resToSend += r.(AbstractOptimizer.Instruction).Execute().(string)
				resToSend += "}\n\n"
			}
		}
	}
	resToSend = resToSend[:len(resToSend)-1]

	return c.Render("index", fiber.Map{
		"Parser3D":                 data.Code,
		"Code3DirectionsOptimized": resToSend,
		"Parser":                   "fn main() {\n    let x = 10;\n    let y = 25;\n\n    let z = x + y;\n\n    print!(\"Suma de x + y = \");\n  \tprintln!(\"{}\", z);\n}",
	})
}
