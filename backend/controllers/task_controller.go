package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seus31/todo-application/backend/dto"
	"github.com/seus31/todo-application/backend/models"
	"github.com/seus31/todo-application/backend/services"
	"github.com/seus31/todo-application/backend/utils"
)

type TaskController struct {
	TaskService *services.TaskService
}

func NewUTaskController(taskService *services.TaskService) *TaskController {
	return &TaskController{
		TaskService: taskService,
	}
}

func (tc *TaskController) GetTasks(ctx *fiber.Ctx) error {
	req := new(dto.GetTasksRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid query parameters"})
	}

	if err := utils.ValidateStruct(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	offset := (req.Page - 1) * req.Limit
	tasks, err := tc.TaskService.GetTasks(utils.GetContextFromFiber(ctx), req.Limit, offset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
	}

	return ctx.JSON(tasks)
}

func (tc *TaskController) CreateTask(ctx *fiber.Ctx) error {
	var req dto.CreateTaskRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request parsing failed"})
	}

	if err := utils.ValidateStruct(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	task := &models.Task{
		TaskName: req.TaskName,
	}

	if err := tc.TaskService.CreateTask(utils.GetContextFromFiber(ctx), task); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create task"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(task)
}
