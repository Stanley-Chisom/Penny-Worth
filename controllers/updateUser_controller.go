package controllers

import (
	"pennyWorth/database"
	"pennyWorth/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)


func GetUserProfile(c *fiber.Ctx) error {
	userID := c.Locals("user").(float64)
	var user models.User

	if err := database.DB.First(&user, uint(userID)).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found"})
	}
	return c.JSON(user)
}

func UpdateUserProfile(c *fiber.Ctx) error {
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
		hashedPassword, _ := bcrypt.GenerateFromPassword(
			[]byte(updateUserData.Password), 14)
		user.Password = string(hashedPassword)
	}

	database.DB.Save(&user)
	return c.Status(201).JSON(fiber.Map{
		"success": "User profile updated"})
}
