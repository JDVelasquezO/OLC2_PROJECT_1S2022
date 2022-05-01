package routes

import (
	"OLC2_Project1/server/controllers"
	"OLC2_Project1/server/interpreter/Optimizer"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/errors"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	newTable := SymbolTable.NewSymbolTable("test", nil)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Parser":    "fn main() {\n    let x = 10;\n    let y = 25;\n\n    let z = x + y;\n\n    println!(\"Suma de x + y = {} \",z);\n}",
			"Res":       "",
			"Err":       errors.TypeError,
			"Symbols":   newTable.Table,
			"Errors":    newTable.ArrayTable,
			"Functions": newTable.FunctionTable,
			"Parser3D":  "/*------HEADER------*/\n#include <stdio.h>\n#include <math.h>\nfloat heap[30101999];\nfloat stack[30101999];\nfloat P;\nfloat H;\nfloat t0;\n\n/*------MAIN------*/\n void main() { \n\tP = 0; H = 0;\n\n\t/* ---- Declaration ---- */\n\t/* ---- Arithmetic Operation ---- */\n\tt0 = 3 + 2;\n\t/* ---- Assign Value ---- */\n\tstack[(int)0] = t0;\n\n\t return; \n }",
		})
	})
	app.Post("/analyze", controllers.Analyze)
	app.Post("/compiler", controllers.Compiler)
	app.Post("/optimizer", Optimizer.Optimize)
}
