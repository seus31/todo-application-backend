package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seus31/todo-application-backend/controllers"
	"github.com/seus31/todo-application-backend/repository"
	"github.com/seus31/todo-application-backend/services"
	"gorm.io/gorm"
)

func SetUpCategoryRoutes(router fiber.Router, db *gorm.DB) {
	var categoryRepository = repository.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepository)
	categoryController := controllers.NewCategoryController(categoryService)
	router.Post("/", categoryController.CreateCategory)
	router.Get("/", categoryController.GetCategories)
	router.Get("/:id", categoryController.GetCategory)
	router.Put("/:id", categoryController.UpdateCategory)
	router.Delete("/:id", categoryController.DeleteCategory)
}
