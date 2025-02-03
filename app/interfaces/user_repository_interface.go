package interfaces

import (
	"context"
	"github.com/seus31/todo-application-backend/models"
)

type UserRepositoryInterface interface {
	GetUsers(ctx context.Context, limit int, offset int) ([]*models.User, error)
	Create(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id uint) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, user *models.User) error
	FindUserByName(ctx context.Context, name string) (*models.User, error)
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
}
