package controllers

import (
	"OLC2_Project1/server/interpreter"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/errors"
	"OLC2_Project1/server/parser"
	"OLC2_Project1/server/utilities"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type CodeToSend struct {
	Code string `form:"code"`
}

func Analyze(c *fiber.Ctx) error {
	//var data map[string]string
	//e := c.BodyParser(&data)

	data := new(CodeToSend)
	e := c.BodyParser(data)
	if e != nil {
		return e
	}
	fmt.Println(data.Code)

	errors.TypeError = make([]errors.Error, 0)
	errors.CounterError = 0

	codeToSend := antlr.NewInputStream(data.Code)
	interpreter.Console = ""

	// lexical analysis:
	lexicalErrors := &utilities.CustomErrorListener{}
	lexer := parser.NewProjectLexer(codeToSend)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexicalErrors)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// syntactical analysis
	parseErrors := &utilities.CustomErrorListener{}
	p := parser.NewProjectParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(parseErrors)

	p.BuildParseTrees = true
	tree := p.Start()

	if len(lexicalErrors.Errors) > 0 {
		fmt.Printf("Lexer %d errors found\n", len(lexicalErrors.Errors))
		for _, e := range lexicalErrors.Errors {
			errors.CounterError += 1
			err := errors.Error{
				Id:    errors.CounterError,
				Type:  "Lexical",
				Line:  e.Row,
				Col:   e.Column,
				Msg:   e.Msg,
				Ambit: "global",
			}
			errors.TypeError = append(errors.TypeError, err)
			fmt.Println("\t", errors.TypeError)
		}
	}

	if len(parseErrors.Errors) > 0 {
		fmt.Printf("Parser %d errors found\n", len(parseErrors.Errors))
		for _, e := range parseErrors.Errors {
			row := strconv.Itoa(e.Row)
			col := strconv.Itoa(e.Column)
			interpreter.Console += "(" + row + ", " + col + ") " + e.Msg + "\n"
			errors.CounterError += 1
			err := errors.Error{
				Id:    errors.CounterError,
				Type:  "Syntactic",
				Line:  e.Row,
				Col:   e.Column,
				Msg:   e.Msg,
				Ambit: "global",
			}
			errors.TypeError = append(errors.TypeError, err)
			fmt.Println("\t", e)
		}
	}

	var listener = utilities.NewTreeShapeListener()
	//if len(lexicalErrors.Errors) == 0 && len(parseErrors.Errors) == 0 {
	//	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	//}
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	ast := listener.Tree
	globalTable := SymbolTable.NewSymbolTable("global", nil)

	if ast.ListInstr != nil {
		for i := 0; i < ast.ListInstr.Len(); i++ {
			r := ast.ListInstr.GetValue(i)
			if r != nil {
				r.(Abstract.Instruction).Execute(globalTable)
			}
		}
	}

	return c.JSON(fiber.Map{
		"parser": data.Code,
		"res":    interpreter.Console,
		"err":    errors.TypeError,
	})
}
