package controllers

import (
	"OLC2_Project1/server/interpreter"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment"
	"OLC2_Project1/server/interpreter/errors"
	"OLC2_Project1/server/parser"
	"OLC2_Project1/server/utilities"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gofiber/fiber/v2"
	"reflect"
	"strconv"
	"time"
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
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	ast := listener.Tree
	interpreter.GlobalTable = SymbolTable.NewSymbolTable("global", nil)

	if len(lexicalErrors.Errors) == 0 && len(parseErrors.Errors) == 0 {
		start := time.Now()
		if ast.ListInstr != nil {
			// First time for functions
			for i := 0; i < ast.ListInstr.Len(); i++ {
				r := ast.ListInstr.GetValue(i)
				if typeof(r) == "Environment.Function" {
					interpreter.GlobalTable.AddFunction(r.(Environment.Function).Id, r.(Environment.Function))
				}
			}

			// Second time for Main
			if interpreter.GlobalTable.ExistsFunction("main") {
				listInstructs := interpreter.GlobalTable.GetFunction("main").(Environment.Function).ListInstructs
				for i := 0; i < listInstructs.Len(); i++ {
					r := listInstructs.GetValue(i)

					if typeof(r) == "Natives.Break" || typeof(r) == "Natives.Continue" || typeof(r) == "Natives.return" {
						errors.CounterError += 1
						msg := "(ERROR) No se puede declarar Break sin un ciclo \n"
						err := errors.NewError(errors.CounterError, 0, 0, msg, interpreter.GlobalTable.Name)
						errors.TypeError = append(errors.TypeError, err)
						interpreter.Console += fmt.Sprintf("%v", msg)
						continue
					}

					r.(Abstract.Instruction).Execute(interpreter.GlobalTable)
				}
			} else {
				errors.CounterError += 1
				msg := "(ERROR) No existe método MAIN \n"
				err := errors.NewError(errors.CounterError, 0, 0, msg, interpreter.GlobalTable.Name)
				errors.TypeError = append(errors.TypeError, err)
				interpreter.Console += fmt.Sprintf("%v", msg)
			}
		}
		t := time.Now()
		elapsed := t.Sub(start)
		interpreter.Console += "\n\n Ejecutado en " + elapsed.Round(elapsed).String()
	}

	return c.Render("index", fiber.Map{
		"Parser": data.Code,
		"Res":    interpreter.Console,
		"Err":    errors.TypeError,
	})
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}
