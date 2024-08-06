package interfaces

import (
	"context"
	"github.com/seus31/todo-application/backend/models"
)

type CategoryRepositoryInterface interface {
	Create(ctx context.Context, category *models.Category) error
}
