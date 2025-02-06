package admin_controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/seus31/todo-application-backend/services"
	admin_services "github.com/seus31/todo-application-backend/services/admin"
)

type AuthController struct {
	authService *admin_services.AuthService
}

func NewAuthController(authService *admin_services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	token, err := c.authService.Login(ctx)
	if err != nil {
		if errors.Is(err, services.ErrFailedToParseRequest) ||
			errors.Is(err, services.ErrInvalidCredentials) {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"token": token})
}
