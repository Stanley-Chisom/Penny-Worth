package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"pennyWorth/utils"
)

func AuthMiddleware(c *fiber.Ctx) error {
	userID, err := utils.VerifyJWT(c)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}
	c.Locals("user", userID)
	return c.Next()
}
