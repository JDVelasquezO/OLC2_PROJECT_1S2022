package controllers

import (
	"OLC2_Project1/server/Generator"
	"OLC2_Project1/server/interpreter"
	"OLC2_Project1/server/interpreter/AST/Natives/DecArrays"
	"OLC2_Project1/server/interpreter/AST/Natives/DecStructs"
	"OLC2_Project1/server/interpreter/AST/Natives/Module"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment"
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

	if len(parseErrors.Errors) > 0 {
		fmt.Printf("Parser %d errors found\n", len(parseErrors.Errors))
		for _, e := range parseErrors.Errors {
			row := strconv.Itoa(e.Row)
			col := strconv.Itoa(e.Column)
			interpreter.Console += "(" + row + ", " + col + ") " + e.Msg + "\n"
			fmt.Println("\t", e)
		}
	}

	ast := listener.Tree
	interpreter.GlobalTable = SymbolTable.NewSymbolTable("global", nil)

	if len(lexicalErrors.Errors) == 0 && len(parseErrors.Errors) == 0 {
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
				newTable := SymbolTable.NewSymbolTable("main", &interpreter.GlobalTable)
				codeCompiled.FreeAllTemps()
				listInstructs := interpreter.GlobalTable.GetFunction("main").(Environment.Function).ListInstructs
				for i := 0; i < listInstructs.Len(); i++ {
					r := listInstructs.GetValue(i)
					r.(Abstract.Instruction).Execute(newTable)

					codeCompiled.FreeAllTemps()
					interpreter.FlagIf = false
					r.(Abstract.Instruction).Compile(&newTable, &codeCompiled)
				}
				codeCompiled.FreeAllTemps()
			}
		}

	}
	codeString := codeCompiled.GetCode()

	return c.Render("index", fiber.Map{
		"Parser":          data.Code,
		"Code3Directions": codeString,
	})
}
