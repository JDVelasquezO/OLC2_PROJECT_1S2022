package routes

import (
	"OLC2_Project1/server/controllers"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	//m := make(map[string]interface{})
	//m["key"] = "-"
	//var it []interface{}
	//it = append(it, m)
	newTable := SymbolTable.NewSymbolTable("test", nil)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Parser":  "fn main() {\n    let x = 10;\n    let y = 25;\n\n    let z = x + y;\n\n    println!(\"Suma de x + y = {} \",z);\n}",
			"Res":     "",
			"Err":     "",
			"Symbols": newTable.Table,
		})
	})
	app.Post("/analyze", controllers.Analyze)
}
