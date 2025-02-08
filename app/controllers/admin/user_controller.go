package admin_controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seus31/todo-application-backend/dto/requests/users"
	"github.com/seus31/todo-application-backend/dto/responses"
	"github.com/seus31/todo-application-backend/models"
	"github.com/seus31/todo-application-backend/services/admin"
	"github.com/seus31/todo-application-backend/utils"
	"strconv"
	"time"
)

type UserController struct {
	UserService *admin_services.UserService
}

func NewUserController(userService *admin_services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) CreateUser(ctx *fiber.Ctx) error {
	var req users.CreateUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request parsing failed"})
	}

	validate := users.CreateUserRequestValidator()
	if err := utils.ValidateStruct(validate, &req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashPassword,
	}

	if err := uc.UserService.CreateUser(utils.GetContextFromFiber(ctx), user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(user)
}

func (uc *UserController) GetUsers(ctx *fiber.Ctx) error {
	var req users.GetUsersRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid parameters"})
	}

	validate := users.GetUsersRequestValidator()
	if err := utils.ValidateStruct(validate, &req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	offset := (req.Page - 1) * req.Limit
	usersData, err := uc.UserService.GetUsers(utils.GetContextFromFiber(ctx), req.Limit, offset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get users"})
	}

	return ctx.Status(fiber.StatusOK).JSON(usersData)
}

func (uc *UserController) GetUser(ctx *fiber.Ctx) error {
	var req users.GetUserRequest
	if err := ctx.ParamsParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid parameters"})
	}

	validate := users.GetUserRequestValidator()
	if err := utils.ValidateStruct(validate, &req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := uc.UserService.GetUser(utils.GetContextFromFiber(ctx), req.ID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	response := responses.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (uc *UserController) UpdateUser(ctx *fiber.Ctx) error {
	var req users.UpdateUserRequest
	userId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	validate := users.UpdateUserRequestValidator()
	if err := utils.ValidateStruct(validate, req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := uc.UserService.GetUser(utils.GetContextFromFiber(ctx), uint(userId))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.Password != nil {
		hashPassword, err := utils.HashPassword(*req.Password)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		user.Password = hashPassword
	}

	updatedUser, err := uc.UserService.UpdateUser(utils.GetContextFromFiber(ctx), user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot update user"})
	}

	response := responses.UserResponse{
		ID:        updatedUser.ID,
		Name:      updatedUser.Name,
		Email:     updatedUser.Email,
		CreatedAt: updatedUser.CreatedAt.Format(time.RFC3339),
		UpdatedAt: updatedUser.UpdatedAt.Format(time.RFC3339),
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (uc *UserController) DeleteUser(ctx *fiber.Ctx) error {
	userId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	user, err := uc.UserService.GetUser(utils.GetContextFromFiber(ctx), uint(userId))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	if err := uc.UserService.DeleteUser(utils.GetContextFromFiber(ctx), user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot delete user"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}
