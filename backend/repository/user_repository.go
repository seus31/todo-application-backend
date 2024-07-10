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

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}
