package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seus31/todo-application-backend/controllers"
	"github.com/seus31/todo-application-backend/repository"
	"github.com/seus31/todo-application-backend/services"
	"gorm.io/gorm"
)

func SetUpTaskRoutes(router fiber.Router, db *gorm.DB) {
	var taskRepository = repository.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepository)
	taskController := controllers.NewTaskController(taskService)
	router.Get("/", taskController.GetTasks)
	router.Post("/", taskController.CreateTask)
	router.Get("/:id", taskController.GetTask)
	router.Put("/:id", taskController.UpdateTask)
	router.Delete("/:id", taskController.DeleteTask)
}
