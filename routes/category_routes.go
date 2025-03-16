package routes

import (
	"github.com/gofiber/fiber/v2"
	"pennyWorth/controllers"
	"pennyWorth/middlewares"
)

func CategoryRoutes(app *fiber.App) {
	categoryGroup := app.Get("/api/category", middlewares.AuthMiddleware)

	categoryGroup.Post("/", controllers.CreateCategory)
	categoryGroup.Get("/", controllers.GetCategories)
}
