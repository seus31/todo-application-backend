package admin_routes

import (
	"github.com/gofiber/fiber/v2"
	admin_controllers "github.com/seus31/todo-application-backend/controllers/admin"
	admin_repository "github.com/seus31/todo-application-backend/repository/admin"
	admin_services "github.com/seus31/todo-application-backend/services/admin"
	"gorm.io/gorm"
)

func SetUpAuthRoutes(router fiber.Router, db *gorm.DB) {
	var adminRepository = admin_repository.NewAdminRepository(db)
	authService := admin_services.NewAuthService(adminRepository)
	authController := admin_controllers.NewAuthController(authService)
	router.Post("/login", authController.Login)
}
