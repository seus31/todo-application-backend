package interfaces

import (
	"context"
	"github.com/seus31/todo-application/backend/models"
)

type TaskRepositoryInterface interface {
	Create(ctx context.Context, task *models.Task) error
}
