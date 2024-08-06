package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seus31/todo-application/backend/dto/requests/categories"
	"github.com/seus31/todo-application/backend/dto/responses"
	"github.com/seus31/todo-application/backend/models"
	"github.com/seus31/todo-application/backend/services"
	"github.com/seus31/todo-application/backend/utils"
	"time"
)

type CategoryController struct {
	CategoryService *services.CategoryService
}

func NewCategoryController(categoryService *services.CategoryService) *CategoryController {
	return &CategoryController{
		CategoryService: categoryService,
	}
}

func (cc *CategoryController) CreateCategory(ctx *fiber.Ctx) error {
	var req categories.CreateCategoryRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request parsing failed"})
	}

	if err := utils.ValidateStruct(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	category := &models.Category{
		CategoryName: req.CategoryName,
	}

	if err := cc.CategoryService.CreateCategory(utils.GetContextFromFiber(ctx), category); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create category"})
	}

	response := responses.CategoryResponse{
		ID:           category.ID,
		CategoryName: category.CategoryName,
		CreatedAt:    category.CreatedAt.Format(time.RFC3339),
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}
