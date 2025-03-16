package routes

import (
	"pennyWorth/controllers"
	"pennyWorth/middlewares"

	"github.com/gofiber/fiber/v2"
)

func UpdateUserRoutes(app *fiber.App) {
	updateProfileGroup := app.Group("/api/profile", middlewares.AuthMiddleware)

	updateProfileGroup.Get("/", controllers.GetUserProfile)
	updateProfileGroup.Patch("/", controllers.UpdateUserProfile)
}
