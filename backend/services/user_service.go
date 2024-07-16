package services

import (
	"context"
	"github.com/seus31/todo-application/backend/interfaces"
	"github.com/seus31/todo-application/backend/models"
)

type UserService struct {
	userRepo interfaces.UserRepositoryInterface
}

func NewUserService(repo interfaces.UserRepositoryInterface) *UserService {
	return &UserService{
		userRepo: repo,
	}
}

func (s *UserService) GetUsers(ctx context.Context, offset int, limit int) ([]*models.User, error) {
	return s.userRepo.GetUsers(ctx, offset, limit)
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {

	return s.userRepo.Create(ctx, user)
}
