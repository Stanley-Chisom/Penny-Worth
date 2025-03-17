package controllers

import (
	"github.com/gofiber/fiber/v2"
	"pennyWorth/database"
	"pennyWorth/models"
)


func GetDashBoardMetrics(c *fiber.Ctx) error {
	var totalExpenses float64
	database.DB.Model(&models.Expense{}).Select(
		"SUM(amount)").Scan(&totalExpenses)

	var highestSpending struct {
		Category string
		Total    float64
	}

	database.DB.Raw(`
			SELECT category, SUM(amount) as total
			FROM expenses 
			GROUP BY category
			ORDER BY total DESC
			LIMIT 1
			
		`).Scan(&highestSpending)

	return c.JSON(fiber.Map{
		"total_spent":          totalExpenses,
		"highest_spending_cat": highestSpending.Category,
		"highest_spending_amt": highestSpending.Total,
	})
}

func GetMonthlySummary(c *fiber.Ctx) error {
	var result []struct {
		Category string
		Total    float64
	}
	database.DB.Model(&models.Expense{}).
		Select("category, SUM(amount) as total").
		Group("category").
		Find(&result)

	return c.JSON(result)
}
