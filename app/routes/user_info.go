package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seus31/todo-application-backend/controllers"
	"github.com/seus31/todo-application-backend/repository"
	"github.com/seus31/todo-application-backend/services"
	"gorm.io/gorm"
)

func SetUpUserInfoRoutes(router fiber.Router, db *gorm.DB) {
	var userRepository = repository.NewUserRepository(db)
	userInfoService := services.NewUserInfoService(userRepository)
	userInfoController := controllers.NewUserInfoController(userInfoService)
	router.Get("/info", userInfoController.Info)
}
