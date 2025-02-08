package admin_routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seus31/todo-application-backend/controllers/admin"
	"github.com/seus31/todo-application-backend/repository"
	"github.com/seus31/todo-application-backend/services/admin"
	"gorm.io/gorm"
)

func SetUpUserRoutes(router fiber.Router, db *gorm.DB) {
	var userRepository = repository.NewUserRepository(db)
	userService := admin_services.NewUserService(userRepository)
	userController := admin_controllers.NewUserController(userService)
	router.Get("/", userController.GetUsers)
	router.Post("/", userController.CreateUser)
	router.Get("/:id", userController.GetUser)
	router.Put("/:id", userController.UpdateUser)
	router.Delete("/:id", userController.DeleteUser)
}
