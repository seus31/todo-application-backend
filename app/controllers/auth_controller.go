package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seus31/todo-application-backend/dto/requests/users"
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
	var req users.CreateUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request parsing failed"})
	}

	validate := users.CreateUserRequestValidator()
	if err := utils.ValidateStruct(validate, &req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if c.authService.CheckUserByName(utils.GetContextFromFiber(ctx), req.Name) == false {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Username already exists"})
	}

	if c.authService.CheckUserByEmail(utils.GetContextFromFiber(ctx), req.Email) == false {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email already exists"})
	}

	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := c.authService.Register(utils.GetContextFromFiber(ctx), req.Name, req.Email, hashPassword); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
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
