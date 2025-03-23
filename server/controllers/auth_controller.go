package controllers

import (
	"pennyWorth/database"
	"pennyWorth/models"
	"pennyWorth/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password), 14)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to hash password"})
	}
	user.Password = string(hashedPassword)

	database.DB.Create(&user)
	return c.Status(201).JSON(fiber.Map{"message": "User registered"})
}

func LoginUser(c *fiber.Ctx) error {
	var input models.User

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	var user models.User
	database.DB.Where("username = ?", input.Username).First(&user)

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(input.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid credentials"})
	}

	token, err := jwt.New(jwt.SigningMethodHS256).SignedString(utils.JwtSecret)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{"token": token})
}
