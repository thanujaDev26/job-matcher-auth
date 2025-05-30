// middleware/jwt.go
package middleware

import (
	"gofiber-auth/utils"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		email, err := utils.ParseJWT(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		c.Locals("userEmail", email)

		return c.Next()
	}
}
