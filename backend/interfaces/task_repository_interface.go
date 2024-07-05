package interfaces

import (
	"context"
	"github.com/seus31/todo-application/backend/models"
)

type TaskRepositoryInterface interface {
	GetTasks(ctx context.Context, limit int, offset int) ([]*models.Task, error)
	Create(ctx context.Context, task *models.Task) error
}
