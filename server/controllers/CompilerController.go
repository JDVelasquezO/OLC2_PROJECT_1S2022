package controllers

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/parser"
	"OLC2_Project1/server/utilities"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gofiber/fiber/v2"
)

type CodeToCompile struct {
	Code string `form:"code"`
}

func Compiler(c *fiber.Ctx) error {
	data := new(CodeToCompile)
	e := c.BodyParser(data)
	if e != nil {
		return e
	}
	fmt.Println(data.Code)

	codeCompiled := Generator.NewGenerator(data.Code)
	codeCompiled.CleanAll()

	/* Lexical and syntactical analysis */
	codeToSend := antlr.NewInputStream(data.Code)
	lexicalErrors := &utilities.CustomErrorListener{}
	lexer := parser.NewProjectLexer(codeToSend)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexicalErrors)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parseErrors := &utilities.CustomErrorListener{}
	p := parser.NewProjectParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(parseErrors)
	tree := p.Start()
	var listener = utilities.NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	ast := listener.Tree
	interpreter.GlobalTable = SymbolTable.NewSymbolTable("global", nil)

	if ast.ListInstr != nil {
		for i := 0; i < ast.ListInstr.Len(); i++ {
			r := ast.ListInstr.GetValue(i)
			r.(Abstract.Instruction).Compile(interpreter.GlobalTable, codeCompiled)
		}
	}

	codeString := codeCompiled.GetCode()

	return c.Render("index", fiber.Map{
		"Parser":          data.Code,
		"Code3Directions": codeString,
	})
}
