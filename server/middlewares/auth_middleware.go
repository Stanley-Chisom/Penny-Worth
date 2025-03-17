package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"pennyWorth/utils"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Get the Authorization header
	authHeader := c.Get("Authorization")

	// Check if the header is missing or incorrectly formatted
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(401).JSON(fiber.Map{"error": "Missing or invalid token"})
	}

	// Extract only the token part
	token := strings.TrimPrefix(authHeader, "Bearer ")

	// Verify JWT Token
	userID, err := utils.VerifyJWT(token)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Store user ID in request context
	c.Locals("user", userID)
	return c.Next()
}
