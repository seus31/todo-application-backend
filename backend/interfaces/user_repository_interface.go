package interfaces

import (
	"context"
	"github.com/seus31/todo-application/backend/models"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context, user *models.User) error
}
