package routes

import (
	"github.com/gofiber/fiber/v2"
	"pennyWorth/controllers"
	"pennyWorth/middlewares"
)

func DashboardRoutes(app *fiber.App) {
	dashboardGroup := app.Group("/api/dashboard", middlewares.AuthMiddleware)

	dashboardGroup.Get("/", controllers.GetDashBoardMetrics)
	dashboardGroup.Get("/summary", controllers.GetMonthlySummary)
}
