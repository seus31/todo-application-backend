package interfaces

import (
	"context"
	"github.com/seus31/todo-application-backend/models"
)

type CategoryRepositoryInterface interface {
	Create(ctx context.Context, category *models.Category) error
	GetCategories(ctx context.Context, limit int, offset int) ([]*models.Category, error)
	GetCategoryByID(ctx context.Context, id uint) (*models.Category, error)
	Update(ctx context.Context, category *models.Category) error
	Delete(ctx context.Context, category *models.Category) error
}
