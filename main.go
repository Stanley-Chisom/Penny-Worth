package main

import (
	"pennyWorth/database"
	"pennyWorth/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	app := fiber.New()

	database.Connect()
	database.Migrate()

	app.Post("/api/expenses", createExpense)
	app.Get("/api/expenses", getAllExpenses)
	app.Get("/api/expenses/:id", getExpenseByID)
	app.Delete("/api/expenses/:id", deleteExpense)

	app.Get("api/dashboard", getDashBoardMetrics)
	app.Get("/api/summary", getMonthlySummary)

	app.Post("api/categories", createCategory)
	app.Get("/api/categories", getCategories)

	app.Get("api/profile", getUserProfile)
	app.Patch("/api/profile", updateUserProfile)

	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: []byte("secret"),
	// }))

	app.Listen(":5050")
}

func createExpense(c *fiber.Ctx) error {

	expense := new(models.Expense)

	if err := c.BodyParser(expense); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid Request"})
	}

	database.DB.Create(expense)
	return c.Status(201).JSON(expense)
}

func getAllExpenses(c *fiber.Ctx) error {
	var expenses []models.Expense

	if err := database.DB.Find(&expenses).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to retrieve expenses"})
	}

	return c.JSON(expenses)

}

func getExpenseByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var expense models.Expense
	result := database.DB.First(&expense, id)

	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Failed to retrieve expense"})
	}

	return c.JSON(expense)

}

func deleteExpense(c *fiber.Ctx) error {
	id := c.Params("id")

	result := database.DB.Delete(&models.Expense{}, id)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "Expense not found"})
	}
	return c.JSON(fiber.Map{"success": "Expense deleted"})
}

func getMonthlySummary(c *fiber.Ctx) error {
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

func registerUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Invalid request"})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword(
		[]byte(user.Password), 14)
	user.Password = string(hashedPassword)

	database.DB.Create(&user)
	return c.Status(201).JSON(fiber.Map{"message": "User registered"})
}

func loginUser(c *fiber.Ctx) error {
	var input models.User

	if err := c.BodyParser(input); err != nil {
		c.Status(404).JSON(fiber.Map{"error": "Invalid request"})
	}

	var user models.User
	database.DB.Where("username = ?", input.Username).First(&user)

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Invalid request"})
	}

	token, _ := jwt.New(jwt.SigningMethodHS256).SignedString([]byte("secret"))
	return c.JSON(fiber.Map{"token": token})
}

func createCategory(c *fiber.Ctx) error {
	category := new(models.Category)

	if err := c.BodyParser(category); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request"})
	}
	database.DB.Create(&category)
	return c.Status(201).JSON(category)
}

func getCategories(c *fiber.Ctx) error {
	var categories []models.Category
	database.DB.Find(&categories)
	return c.JSON(categories)
}

func getDashBoardMetrics(c *fiber.Ctx) error {
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

func getUserProfile(c *fiber.Ctx) error {
	userID := c.Locals("user").(float64)
	var user models.User

	if err := database.DB.First(&user, uint(userID)).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found"})
	}
	return c.JSON(user)
}

func updateUserProfile(c *fiber.Ctx) error {
	userID := c.Locals("user").(float64)
	var user models.User

	if err := database.DB.First(&user, uint(userID)).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found"})
	}

	var updateUserData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&updateUserData); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "User not found"})
	}

	if updateUserData.Email != "" {
		user.Email = updateUserData.Email
	}

	if updateUserData.Password != "" {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(updateUserData.Password), 14)
		user.Password = string(hashedPassword)
	}

	database.DB.Save(&user)
	return c.Status(201).JSON(fiber.Map{
		"success": "User profile updated"})
}
