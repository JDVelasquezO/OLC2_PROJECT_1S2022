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
			"Parser":    "fn main() {\n    let x = 10;\n    let y = 25;\n\n    let z = x + y;\n\n    print!(\"Suma de x + y = \");\n  \tprintln!(\"{}\", z);\n}",
			"Res":       "",
			"Err":       errors.TypeError,
			"Symbols":   newTable.Table,
			"Errors":    newTable.ArrayTable,
			"Functions": newTable.FunctionTable,
			"Parser3D":  "/*------HEADER------*/\n#include <stdio.h>\nfloat heap[30101999];\nfloat stack[30101999];\nfloat P;\nfloat H;\nfloat t0, t1, t2, t3;\n\n/*------MAIN------*/\nvoid main() {\n  t1 = 6 / 2;\n  t2 = 3 * t1;\n  t3 = 4 - 2;\n  t4 = t2 * t3;\n  t5 = 2 + 5;\n  t6 = t2 * t3;\n  t7 = t6 + t5;\n}\n",
		})
	})
	app.Post("/analyze", controllers.Analyze)
	app.Post("/compiler", controllers.Compiler)
	app.Post("/optimizer", Optimizer.Optimize)
}
