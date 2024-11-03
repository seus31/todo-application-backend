package services

import (
	"context"
	"github.com/seus31/todo-application/backend/dto/requests/tasks"
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
	return s.taskRepo.Create(ctx, task)
}

func (s *TaskService) GetTask(ctx context.Context, id uint) (*models.Task, error) {
	return s.taskRepo.GetTaskByID(ctx, id)
}

func (s *TaskService) UpdateTask(ctx context.Context, task *models.Task, req tasks.UpdateTaskRequest) (*models.Task, error) {
	task.TaskName = req.TaskName
	err := s.taskRepo.Update(ctx, task)
	if err != nil {
		return nil, err
	}

	updateTask, err := s.taskRepo.GetTaskByID(ctx, task.ID)

	return updateTask, err
}

func (s *TaskService) DeleteTask(ctx context.Context, task *models.Task) error {
	return s.taskRepo.Delete(ctx, task)
}
