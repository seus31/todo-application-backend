package repository

import (
	"context"
	"github.com/seus31/todo-application/backend/interfaces"
	"github.com/seus31/todo-application/backend/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepositoryInterface {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUsers(ctx context.Context, limit int, offset int) ([]*models.User, error) {
	var users []*models.User
	if err := r.db.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) GetUserById(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
