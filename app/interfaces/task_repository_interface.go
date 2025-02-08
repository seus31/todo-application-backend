package interfaces

import (
	"context"
	"github.com/seus31/todo-application-backend/models"
)

type TaskRepositoryInterface interface {
	GetTasks(ctx context.Context, userId uint, limit int, offset int) ([]*models.Task, error)
	Create(ctx context.Context, task *models.Task) error
	GetTaskByID(ctx context.Context, id uint) (*models.Task, error)
	Update(ctx context.Context, task *models.Task) error
	Delete(ctx context.Context, task *models.Task) error
}
