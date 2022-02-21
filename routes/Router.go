package routes

import (
	"OLC2_Project1/server/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/analyze", controllers.Analyze)
}
