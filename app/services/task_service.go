package services

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/seus31/todo-application-backend/dto/requests/tasks"
	"github.com/seus31/todo-application-backend/dto/responses"
	"github.com/seus31/todo-application-backend/interfaces"
	"github.com/seus31/todo-application-backend/models"
	"github.com/seus31/todo-application-backend/utils"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type TaskService struct {
	taskRepo interfaces.TaskRepositoryInterface
}

func NewTaskService(repo interfaces.TaskRepositoryInterface) *TaskService {
	return &TaskService{
		taskRepo: repo,
	}
}

func (s *TaskService) GetTasks(ctx *fiber.Ctx) ([]*models.Task, error) {
	var req tasks.GetTasksRequest
	contextData := utils.GetContextFromFiber(ctx)

	if err := ctx.QueryParser(&req); err != nil {
		return nil, ErrFailedToParseRequest
	}

	validate := tasks.GetTasksRequestValidator()
	if err := utils.ValidateStruct(validate, &req); err != nil {
		return nil, err
	}

	offset := (req.Page - 1) * req.Limit
	tasksData, err := s.taskRepo.GetTasks(contextData, ctx.Locals("userID").(uint), req.Limit, offset)

	if err != nil {
		return nil, ErrUnexpectedError
	}

	return tasksData, nil
}

func (s *TaskService) CreateTask(ctx *fiber.Ctx) (*models.Task, error) {
	var req tasks.CreateTaskRequest
	contextData := utils.GetContextFromFiber(ctx)

	if err := ctx.BodyParser(&req); err != nil {
		return nil, ErrFailedToParseRequest
	}

	validate := tasks.CreateTaskRequestValidator()
	if err := utils.ValidateStruct(validate, &req); err != nil {
		return nil, err
	}

	task := models.NewTask(
		req.TaskName,
		ctx.Locals("userID").(uint),
		models.WithParentID(req.ParentID),
		models.WithDueDate(req.DueDate),
		models.WithDueTime(req.DueTime),
		models.WithStatus(req.Status),
		models.WithPriority(req.Priority),
	)

	if err := s.taskRepo.Create(contextData, task); err != nil {
		return nil, ErrUnexpectedError
	}

	return task, nil
}

func (s *TaskService) GetTask(ctx *fiber.Ctx) (*responses.TaskResponse, error) {
	var req tasks.GetTaskRequest
	contextData := utils.GetContextFromFiber(ctx)

	if err := ctx.ParamsParser(&req); err != nil {
		return nil, ErrInvalidTaskID
	}

	validate := tasks.GetTaskRequestValidator()
	if err := utils.ValidateStruct(validate, req); err != nil {
		return nil, err
	}

	task, err := s.taskRepo.GetTaskByID(contextData, uint(req.ID))
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		} else {
			return nil, ErrUnexpectedError
		}
	}

	response := &responses.TaskResponse{
		ID:        task.ID,
		TaskName:  task.TaskName,
		CreatedAt: task.CreatedAt.Format(time.RFC3339),
	}

	return response, nil
}

func (s *TaskService) UpdateTask(ctx *fiber.Ctx) (*responses.TaskResponse, error) {
	var req tasks.UpdateTaskRequest
	contextData := utils.GetContextFromFiber(ctx)

	taskId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return nil, ErrInvalidTaskID
	}

	if err := ctx.BodyParser(&req); err != nil {
		return nil, ErrInvalidParameters
	}

	validate := tasks.UpdateTaskRequestValidator()
	if err := utils.ValidateStruct(validate, req); err != nil {
		return nil, err
	}

	task, err := s.taskRepo.GetTaskByID(contextData, uint(taskId))
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		} else {
			return nil, ErrUnexpectedError
		}
	}

	task.TaskName = req.TaskName
	if req.ParentID == nil {
		fmt.Println(req.ParentID)
		task.ParentID = nil
	} else {
		task.ParentID = req.ParentID
	}
	task.DueDate = req.DueDate
	task.DueTime = req.DueTime
	task.Status = req.Status
	task.Priority = req.Priority

	err = s.taskRepo.Update(contextData, task)
	if err != nil {
		return nil, ErrUnexpectedError
	}

	response := &responses.TaskResponse{
		ID:        task.ID,
		TaskName:  task.TaskName,
		CreatedAt: task.CreatedAt.Format(time.RFC3339),
		UpdatedAt: task.UpdatedAt.Format(time.RFC3339),
	}

	return response, nil
}

func (s *TaskService) DeleteTask(ctx *fiber.Ctx) error {
	contextData := utils.GetContextFromFiber(ctx)

	taskId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ErrInvalidTaskID
	}

	task, err := s.taskRepo.GetTaskByID(contextData, uint(taskId))
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrRecordNotFound
		} else {
			return ErrUnexpectedError
		}
	}

	err = s.taskRepo.Delete(contextData, task)
	if err != nil {
		return ErrUnexpectedError
	}

	return nil
}
