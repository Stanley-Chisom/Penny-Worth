package routes

import (
	"github.com/gofiber/fiber/v2"
	"pennyWorth/controllers"
)

func AuthRoutes(app *fiber.App) {
	app.Post("/api/register", controllers.RegisterUser)
	app.Post("/api/login", controllers.LoginUser)
}
