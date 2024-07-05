package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seus31/todo-application/backend/controllers"
	"github.com/seus31/todo-application/backend/interfaces"
	"github.com/seus31/todo-application/backend/repository"
	"github.com/seus31/todo-application/backend/services"
	"gorm.io/gorm"
)

func SetUpTaskRoutes(router fiber.Router, db *gorm.DB) {
	var taskRepository interfaces.TaskRepositoryInterface = repository.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepository)
	taskController := controllers.NewUTaskController(taskService)
	router.Get("/tasks", taskController.GetTasks)
	router.Post("/task", taskController.CreateTask)
}
