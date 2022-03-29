package controllers

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter"
	"OLC2_Project1/server/interpreter/AST/Natives"
	"OLC2_Project1/server/interpreter/AST/Natives/DecArrays"
	"OLC2_Project1/server/interpreter/AST/Natives/DecStructs"
	"OLC2_Project1/server/interpreter/AST/Natives/Module"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment"
	"OLC2_Project1/server/interpreter/errors"
	"OLC2_Project1/server/parser"
	"OLC2_Project1/server/utilities"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gofiber/fiber/v2"
	"strconv"
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
			if typeof(r) == "Environment.Function" {
				interpreter.GlobalTable.AddFunction(r.(Environment.Function).Id, r.(Environment.Function))
			}

			if typeof(r) == "DecStructs.DecStruct" {
				interpreter.GlobalTable.AddStruct(r.(DecStructs.DecStruct).Id, r.(DecStructs.DecStruct))
			}

			if typeof(r) == "DecArrays.DecArray" {
				interpreter.GlobalTable.AddArray(r.(DecArrays.DecArray).Id, r.(DecArrays.DecArray))
			}

			if typeof(r) == "Module.Module" {
				interpreter.GlobalTable.AddModule(r.(Module.Module).Id, r.(Module.Module))
			}
		}

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
		}
	}

	if ast.ListInstr != nil {
		for i := 0; i < ast.ListInstr.Len(); i++ {
			r := ast.ListInstr.GetValue(i)
			if interpreter.GlobalTable.ExistsFunction("main") {
				newTable := SymbolTable.NewSymbolTable("main", &interpreter.GlobalTable)
				newTable.SizeTable = 1
				codeCompiled.FreeAllTemps()
				listInstructs := interpreter.GlobalTable.GetFunction("main").(Environment.Function).ListInstructs
				for i := 0; i < listInstructs.Len(); i++ {
					instr := r.(Environment.Function).ListInstructs.GetValue(i)
					instr.(Abstract.Instruction).Compile(newTable, &codeCompiled)
				}
				codeCompiled.FreeAllTemps()
			} else {
				r.(Abstract.Instruction).Compile(interpreter.GlobalTable, &codeCompiled)
			}
		}
	}

	codeString := codeCompiled.GetCode()

	return c.Render("index", fiber.Map{
		"Parser":          data.Code,
		"Code3Directions": codeString,
	})
}
