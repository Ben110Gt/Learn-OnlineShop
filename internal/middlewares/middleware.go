package middlewares

import (
	"OnlineShop/Api/internal/util"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware ตรวจสอบ JWT token
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header is missing",
			})
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header format must be Bearer {token}",
			})
		}

		tokenString := parts[1]
		claims, err := util.ValidateToken(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		// Save claims to locals (เหมือน c.Set ใน Gin)
		c.Locals("user_id", claims.UserID)
		c.Locals("phone", claims.Phone)
		c.Locals("role", claims.Role)

		return c.Next()
	}
}

// UserOnly middleware
func UserOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		roleI := c.Locals("role")
		role, ok := roleI.(string)
		if !ok || role != "user" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "User access only",
			})
		}
		return c.Next()
	}
}

// AdminOnly middleware
func AdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		roleI := c.Locals("role")
		role, ok := roleI.(string)
		if !ok || role != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Admin access only",
			})
		}
		return c.Next()
	}
}
