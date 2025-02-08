package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/seus31/todo-application-backend/services"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	if err := c.authService.Register(ctx); err != nil {
		if errors.Is(err, services.ErrFailedToRegisterUser) {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		} else {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	token, userId, err := c.authService.Login(ctx)
	if err != nil {
		if errors.Is(err, services.ErrFailedToParseRequest) ||
			errors.Is(err, services.ErrInvalidCredentials) {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"token": token, "userId": userId})
}
