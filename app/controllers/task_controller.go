package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seus31/todo-application-backend/dto/requests/tasks"
	"github.com/seus31/todo-application-backend/dto/responses"
	"github.com/seus31/todo-application-backend/models"
	"github.com/seus31/todo-application-backend/services"
	"github.com/seus31/todo-application-backend/utils"
	"log"
	"strconv"
	"time"
)

type TaskController struct {
	TaskService *services.TaskService
}

func NewTaskController(taskService *services.TaskService) *TaskController {
	return &TaskController{
		TaskService: taskService,
	}
}

func (tc *TaskController) GetTasks(ctx *fiber.Ctx) error {
	var req tasks.GetTasksRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid parameters"})
	}

	validate := tasks.GetTasksRequestValidator()
	if err := utils.ValidateStruct(validate, &req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	offset := (req.Page - 1) * req.Limit
	tasksData, err := tc.TaskService.GetTasks(utils.GetContextFromFiber(ctx), req.Limit, offset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
	}

	return ctx.JSON(tasksData)
}

func (tc *TaskController) CreateTask(ctx *fiber.Ctx) error {
	var req tasks.CreateTaskRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	validate := tasks.CreateTaskRequestValidator()
	if err := utils.ValidateStruct(validate, &req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	task := &models.Task{
		TaskName: req.TaskName,
		UserID:   req.UserID,
		ParentID: req.ParentID,
		DueDate:  req.DueDate,
		DueTime:  req.DueTime,
		Status:   req.Status,
		Priority: req.Priority,
	}

	if err := tc.TaskService.CreateTask(utils.GetContextFromFiber(ctx), task); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create task"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(task)
}

func (tc *TaskController) GetTask(ctx *fiber.Ctx) error {
	var req tasks.GetTaskRequest
	if err := ctx.ParamsParser(&req); err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid task ID"})
	}

	validate := tasks.GetTaskRequestValidator()
	if err := utils.ValidateStruct(validate, req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	task, err := tc.TaskService.GetTask(utils.GetContextFromFiber(ctx), uint(req.ID))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}

	response := responses.TaskResponse{
		ID:        task.ID,
		TaskName:  task.TaskName,
		CreatedAt: task.CreatedAt.Format(time.RFC3339),
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (tc *TaskController) UpdateTask(ctx *fiber.Ctx) error {
	var req tasks.UpdateTaskRequest
	taskId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid task ID"})
	}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	validate := tasks.UpdateTaskRequestValidator()
	if err := utils.ValidateStruct(validate, req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	task, err := tc.TaskService.GetTask(utils.GetContextFromFiber(ctx), uint(taskId))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}

	updatedTask, err := tc.TaskService.UpdateTask(utils.GetContextFromFiber(ctx), task, req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot update task"})
	}

	response := responses.TaskResponse{
		ID:        updatedTask.ID,
		TaskName:  updatedTask.TaskName,
		CreatedAt: updatedTask.CreatedAt.Format(time.RFC3339),
		UpdatedAt: updatedTask.UpdatedAt.Format(time.RFC3339),
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (tc *TaskController) DeleteTask(ctx *fiber.Ctx) error {
	taskId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid task ID"})
	}

	task, err := tc.TaskService.GetTask(utils.GetContextFromFiber(ctx), uint(taskId))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}

	if err := tc.TaskService.DeleteTask(utils.GetContextFromFiber(ctx), task); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot delete task"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}
