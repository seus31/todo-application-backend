package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seus31/todo-application-backend/controllers"
	"github.com/seus31/todo-application-backend/repository"
	"github.com/seus31/todo-application-backend/services"
	"gorm.io/gorm"
)

func SetUpAuthRoutes(router fiber.Router, db *gorm.DB) {
	var userRepository = repository.NewUserRepository(db)
	authService := services.NewAuthService(userRepository)
	authController := controllers.NewAuthController(authService)
	router.Post("/login", authController.Login)
	router.Post("/register", authController.Register)
}
