package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seus31/todo-application/backend/controllers"
	"github.com/seus31/todo-application/backend/repository"
	"github.com/seus31/todo-application/backend/services"
	"gorm.io/gorm"
)

func SetUpUserRoutes(router fiber.Router, db *gorm.DB) {
	var userRepository = repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
	router.Get("/", userController.GetUsers)
	router.Post("/", userController.CreateUser)
	router.Get("/:id", userController.GetUser)
	router.Put("/:id", userController.UpdateUser)
}
