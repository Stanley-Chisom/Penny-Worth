package routes

import (
	"github.com/gofiber/fiber/v2"
	"pennyWorth/controllers"
	"pennyWorth/middlewares"
)

func ExpenseRoutes(app *fiber.App) {
	expenseGroup := app.Group("/api/expenses", middlewares.AuthMiddleware)

	expenseGroup.Post("/", controllers.CreateExpense)
	expenseGroup.Get("/", controllers.GetAllExpenses)
	expenseGroup.Get("/:id", controllers.GetExpenseByID)
	expenseGroup.Patch("/:id", controllers.UpdateExpense)
	expenseGroup.Delete("/:id", controllers.DeleteExpense)
}
