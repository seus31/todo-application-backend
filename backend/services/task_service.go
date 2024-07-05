package services

import (
	"context"
	"github.com/seus31/todo-application/backend/interfaces"
	"github.com/seus31/todo-application/backend/models"
)

type TaskService struct {
	taskRepo interfaces.TaskRepositoryInterface
}

func NewTaskService(repo interfaces.TaskRepositoryInterface) *TaskService {
	return &TaskService{
		taskRepo: repo,
	}
}

func (s *TaskService) GetTasks(ctx context.Context, limit int, offset int) ([]*models.Task, error) {
	return s.taskRepo.GetTasks(ctx, limit, offset)
}

func (s *TaskService) CreateTask(ctx context.Context, task *models.Task) error {
	if err := task.Validate(); err != nil {
		return err
	}

	return s.taskRepo.Create(ctx, task)
}
