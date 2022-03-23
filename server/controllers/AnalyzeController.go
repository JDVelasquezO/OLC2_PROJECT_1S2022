package controllers

import (
	"OLC2_Project1/server/interpreter"
	"OLC2_Project1/server/interpreter/AST/Expression"
	"OLC2_Project1/server/interpreter/AST/Natives"
	"OLC2_Project1/server/interpreter/AST/Natives/DecArrays"
	"OLC2_Project1/server/interpreter/AST/Natives/DecStructs"
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
			// First time for functions and
			for i := 0; i < ast.ListInstr.Len(); i++ {
				r := ast.ListInstr.GetValue(i)
				if typeof(r) == "Environment.Function" {
					interpreter.GlobalTable.AddFunction(r.(Environment.Function).Id, r.(Environment.Function))
				}

				if typeof(r) == "DecStructs.DecStruct" {
					interpreter.GlobalTable.AddStruct(r.(DecStructs.DecStruct).Id, r.(DecStructs.DecStruct))
				}

				if typeof(r) == "DecArrays.DecArray" {
					interpreter.GlobalTable.AddArray(r.(DecArrays.DecArray).Id, r.(DecArrays.DecArray))
				}
			}

			// Second time for Main
			if interpreter.GlobalTable.ExistsFunction("main") {
				listInstructs := interpreter.GlobalTable.GetFunction("main").(Environment.Function).ListInstructs
				for i := 0; i < listInstructs.Len(); i++ {
					r := listInstructs.GetValue(i)

					if typeof(r) == "Natives.Break" || typeof(r) == "Natives.Continue" || typeof(r) == "Natives.Return" {
						errors.CounterError += 1

						var row int
						var col int
						var transfer string
						if typeof(r) == "Natives.Break" {
							row = r.(Natives.Break).Row
							col = r.(Natives.Break).Col
							transfer = "Break"
						}

						if typeof(r) == "Natives.Continue" {
							row = r.(Natives.Continue).Row
							col = r.(Natives.Continue).Col
							transfer = "Continue"
						}

						if typeof(r) == "Natives.Return" {
							row = r.(Natives.Return).Row
							col = r.(Natives.Return).Col
							transfer = "Return"
						}

						msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") Error: No se puede declarar " + transfer + " sin un ciclo \n"
						err := errors.NewError(errors.CounterError, row, col, msg, interpreter.GlobalTable.Name)
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

	symbols := interpreter.GlobalTable
	var listParams []interface{}
	var listVarsFunction []interface{}
	for _, function := range symbols.FunctionTable {
		for i := 0; i < function.(Environment.Function).ListParams.Len(); i++ {
			m := make(map[string]string)
			m["Id"] = function.(Environment.Function).ListParams.GetValue(i).(*Natives.Declaration).ListIds.GetValue(0).(Expression.Identifier).Id
			m["Col"] = fmt.Sprintf("%v", function.(Environment.Function).ListParams.GetValue(i).(*Natives.Declaration).ListIds.GetValue(0).(Expression.Identifier).Col)
			m["Row"] = fmt.Sprintf("%v", function.(Environment.Function).ListParams.GetValue(i).(*Natives.Declaration).ListIds.GetValue(0).(Expression.Identifier).Row)
			m["DataType"] = fmt.Sprintf("%v", function.(Environment.Function).ListParams.GetValue(i).(*Natives.Declaration).DataType)
			listParams = append(listParams, m)
		}

		if function.(Environment.Function).Id != "main" {
			for i := 0; i < function.(Environment.Function).ListInstructs.Len(); i++ {
				if typeof(function.(Environment.Function).ListInstructs.GetValue(i)) == "*Natives.Declaration" {
					m2 := make(map[string]string)
					m2["Id"] = function.(Environment.Function).ListInstructs.GetValue(i).(*Natives.Declaration).ListIds.GetValue(0).(Expression.Identifier).Id
					m2["Col"] = fmt.Sprintf("%v", function.(Environment.Function).ListInstructs.GetValue(i).(*Natives.Declaration).ListIds.GetValue(0).(Expression.Identifier).Col)
					m2["Row"] = fmt.Sprintf("%v", function.(Environment.Function).ListInstructs.GetValue(i).(*Natives.Declaration).ListIds.GetValue(0).(Expression.Identifier).Row)
					m2["DataType"] = fmt.Sprintf("%v", function.(Environment.Function).ListInstructs.GetValue(i).(*Natives.Declaration).DataType)
					listVarsFunction = append(listVarsFunction, m2)
				}
			}
		}
	}
	errorsArray := errors.TypeError
	//fmt.Println(symbols)
	return c.Render("index", fiber.Map{
		"Parser":       data.Code,
		"Res":          interpreter.Console,
		"Err":          errorsArray,
		"Symbols":      symbols.Table,
		"Arrays":       symbols.ArrayTable,
		"Functions":    symbols.FunctionTable,
		"Params":       listParams,
		"VarsFunction": listVarsFunction,
	})
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}
