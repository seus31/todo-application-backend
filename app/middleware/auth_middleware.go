package middleware

import "github.com/gofiber/fiber/v2"

func AuthMiddleware(c *fiber.Ctx) error {
	userToken := c.Get("X-User-Token")
	if userToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized: User Token not provided",
		})
	}
	c.Locals("userToken", userToken)
	return c.Next()
}
