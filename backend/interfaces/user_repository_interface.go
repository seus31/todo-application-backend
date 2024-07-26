package interfaces

import (
	"context"
	"github.com/seus31/todo-application/backend/models"
)

type UserRepositoryInterface interface {
	GetUsers(ctx context.Context, limit int, offset int) ([]*models.User, error)
	Create(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id uint) (*models.User, error)
}
