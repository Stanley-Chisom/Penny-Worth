package controllers

import (
	"github.com/gofiber/fiber/v2"
	"pennyWorth/database"
	"pennyWorth/models"
)

func CreateCategory(c *fiber.Ctx) error {
	category := new(models.Category)

	if err := c.BodyParser(category); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request"})
	}
	database.DB.Create(&category)
	return c.Status(201).JSON(category)
}

func GetCategories(c *fiber.Ctx) error {
	var categories []models.Category
	database.DB.Find(&categories)
	return c.JSON(categories)
}
