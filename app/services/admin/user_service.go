package admin_services

import (
	"context"
	"github.com/seus31/todo-application-backend/interfaces"
	"github.com/seus31/todo-application-backend/models"
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

func (s *UserService) GetUser(ctx context.Context, id uint) (*models.User, error) {
	return s.userRepo.GetUserByID(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	err := s.userRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	updateUser, err := s.userRepo.GetUserByID(ctx, user.ID)

	return updateUser, err
}

func (s *UserService) DeleteUser(ctx context.Context, user *models.User) error {
	return s.userRepo.Delete(ctx, user)
}
