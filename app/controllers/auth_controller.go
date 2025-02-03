package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/seus31/todo-application-backend/services"
	"github.com/seus31/todo-application-backend/utils"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	if err := c.authService.Register(ctx); err != nil {
		if errors.Is(err, services.ErrDuplicateName) ||
			errors.Is(err, services.ErrDuplicateEmail) ||
			errors.Is(err, services.ErrPasswordMismatch) ||
			errors.Is(err, services.ErrFailedToParseRequest) ||
			errors.Is(err, services.ErrFailedToHashPassword) {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		if errors.Is(err, services.ErrFailedToRegisterUser) {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var input struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	token, err := c.authService.Login(utils.GetContextFromFiber(ctx), input.Name, input.Password)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	return ctx.JSON(fiber.Map{"token": token})
}
