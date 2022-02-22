package routes

import (
	"OLC2_Project1/server/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Parser": "",
			"Res":    "",
			"Err":    "",
		})
	})
	app.Post("/analyze", controllers.Analyze)
}
