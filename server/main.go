package main

import (
	"pennyWorth/database"
	"pennyWorth/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.Connect()
	database.Migrate()

	routes.AuthRoutes(app)
	routes.ExpenseRoutes(app)
	routes.CategoryRoutes(app)
	routes.DashboardRoutes(app)

	app.Listen(":5050")
}
