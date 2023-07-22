package middlewares

import (
	"gostarter-backend/helpers/token"
	"gostarter-backend/models"
	"gostarter-backend/response"
	"gostarter-backend/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func JwtAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := token.TokenValid(c)
		if err != nil {
			return response.APIResponse(c, http.StatusUnauthorized, "Token invalid", err.Error())
		}
		authUser, err := token.ExtractTokenUser(c)
		if err != nil {
			return response.APIResponse(c, http.StatusUnauthorized, "Token user invalid", err.Error())
		}
		userService := services.UserService{}
		user, err := userService.GetUserByUUID(authUser.UUID)
		if err != nil {
			return response.APIResponse(c, http.StatusUnauthorized, "Get user failed", nil)
		}
		c.Locals("authUser", user)
		return c.Next()
	}
}

func JwtAuthRolesMiddleware(allowedRoles ...models.RoleType) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authUser, ok := c.Locals("authUser").(*models.User)
		if !ok || authUser == nil {
			return response.APIResponse(c, http.StatusForbidden, "Auth user failed", nil)
		}
		// Pemeriksaan roles
		if !checkRole(authUser.Role, allowedRoles) {
			return response.APIResponse(c, http.StatusForbidden, "Access Forbidden", nil)
		}
		// Roles sesuai, lanjutkan request ke handler selanjutnya
		return c.Next()
	}
}

func checkRole(userRole models.RoleType, allowedRoles []models.RoleType) bool {
	for _, role := range allowedRoles {
		if userRole == role {
			return true
		}
	}
	return false
}
