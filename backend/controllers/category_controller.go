package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seus31/todo-application/backend/dto/requests/categories"
	"github.com/seus31/todo-application/backend/dto/responses"
	"github.com/seus31/todo-application/backend/models"
	"github.com/seus31/todo-application/backend/services"
	"github.com/seus31/todo-application/backend/utils"
	"log"
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

func (cc *CategoryController) GetCategories(ctx *fiber.Ctx) error {
	var req categories.GetCategoriesRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid parameters"})
	}

	if err := utils.ValidateStruct(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	offset := (req.Page - 1) * req.Limit
	categoriesData, err := cc.CategoryService.GetCategories(utils.GetContextFromFiber(ctx), req.Limit, offset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch categories"})
	}

	return ctx.Status(fiber.StatusOK).JSON(categoriesData)
}

func (cc *CategoryController) GetCategory(ctx *fiber.Ctx) error {
	var req categories.GetCategoryRequest
	if err := ctx.ParamsParser(&req); err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid category ID"})
	}

	if err := utils.ValidateStruct(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	category, err := cc.CategoryService.GetCategory(utils.GetContextFromFiber(ctx), req.ID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Category not found"})
	}

	response := responses.CategoryResponse{
		ID:           category.ID,
		CategoryName: category.CategoryName,
		CreatedAt:    category.CreatedAt.Format(time.RFC3339),
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
