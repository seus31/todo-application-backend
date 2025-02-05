package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/seus31/todo-application-backend/services"
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
	tasksData, err := tc.TaskService.GetTasks(ctx)
	if err != nil {
		if errors.Is(err, services.ErrUnexpectedError) {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		} else {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
	}

	return ctx.JSON(tasksData)
}

func (tc *TaskController) CreateTask(ctx *fiber.Ctx) error {
	task, err := tc.TaskService.CreateTask(ctx)
	if err != nil {
		if !errors.Is(err, services.ErrUnexpectedError) {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		} else {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
	}

	return ctx.Status(fiber.StatusCreated).JSON(task)
}

func (tc *TaskController) GetTask(ctx *fiber.Ctx) error {
	task, err := tc.TaskService.GetTask(ctx)
	if err != nil {
		if !errors.Is(err, services.ErrUnexpectedError) {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		} else if !errors.Is(err, services.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		} else {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(&task)
}

func (tc *TaskController) UpdateTask(ctx *fiber.Ctx) error {
	task, err := tc.TaskService.UpdateTask(ctx)
	if err != nil {
		if !errors.Is(err, services.ErrUnexpectedError) {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		} else if !errors.Is(err, services.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		} else {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(&task)
}

func (tc *TaskController) DeleteTask(ctx *fiber.Ctx) error {
	err := tc.TaskService.DeleteTask(ctx)
	if err != nil {
		if !errors.Is(err, services.ErrUnexpectedError) {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		} else if !errors.Is(err, services.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		} else {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK"})
}
