package routes

import (
	"OLC2_Project1/server/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Parser": "fn main() {\n    let x = 10;\n    let y = 25;\n\n    let z = x + y;\n\n    println!(\"Suma de x + y = {} \",z);\n}",
			"Res":    "",
			"Err":    "",
		})
	})
	app.Post("/analyze", controllers.Analyze)
}
