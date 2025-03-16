package controllers

import (
	"github.com/gofiber/fiber/v2"
	"pennyWorth/database"
	"pennyWorth/models"
)

func CreateExpense(c *fiber.Ctx) error {

	expense := new(models.Expense)

	if err := c.BodyParser(expense); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid Request"})
	}

	database.DB.Create(expense)
	return c.Status(201).JSON(expense)
}

func UpdateExpense(c *fiber.Ctx) error {
	userID := c.Locals("user").(uint)
	expenseID := c.Params("id")

	var expense models.Expense

	if err := database.DB.Where(
		"id = ? AND user_id = ?",
		userID, expenseID).
		First(&expense).Error; err != nil {
		c.Status(401).JSON(fiber.Map{
			"error": "Expense not found"})
	}

	var updateUserData struct {
		Amount   float64 `json:"amount"`
		Category string  `json:"category"`
		Note     string  `json:"note"`
	}

	if err := c.BodyParser(&updateUserData); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request"})
	}

	if updateUserData.Amount != 0 {
		expense.Amount = updateUserData.Amount
	}

	if updateUserData.Category != "" {
		expense.Category = updateUserData.Category
	}

	if updateUserData.Note != "" {
		expense.Note = updateUserData.Note
	}

	database.DB.Save(&expense)
	return c.JSON(fiber.Map{
		"message": "Expense Updated",
		"expense": expense})
}

func GetAllExpenses(c *fiber.Ctx) error {
	var expenses []models.Expense

	if err := database.DB.Find(&expenses).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to retrieve expenses"})
	}

	return c.JSON(expenses)

}

func GetExpenseByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var expense models.Expense
	result := database.DB.First(&expense, id)

	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Failed to retrieve expense"})
	}

	return c.JSON(expense)

}

func DeleteExpense(c *fiber.Ctx) error {
	id := c.Params("id")

	result := database.DB.Delete(&models.Expense{}, id)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "Expense not found"})
	}
	return c.JSON(fiber.Map{"success": "Expense deleted"})
}
