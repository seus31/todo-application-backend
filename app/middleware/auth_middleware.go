package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/seus31/todo-application-backend/config"
	"time"
)

func AuthMiddleware(c *fiber.Ctx) error {
	secretKey := config.Config("SECRET_KEY")
	userToken := c.Get("X-User-Token")

	if userToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized: User Token not provided",
		})
	}

	token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized: Invalid Token",
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized: Invalid Token Claims",
		})
	}

	exp := int64(claims["exp"].(float64))
	if time.Now().Unix() > exp {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token expired",
		})
	}

	idValue, ok := claims["id"].(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized: Invalid User ID",
		})
	}
	userID := uint(idValue)

	c.Locals("userToken", userToken)
	c.Locals("userName", claims["name"])
	c.Locals("userID", userID)

	return c.Next()
}
