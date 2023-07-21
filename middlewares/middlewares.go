package middlewares

import (
	"gostarter-backend/helpers/token"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func JwtAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := token.TokenValid(c)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
		}
		return c.Next()
	}
}
