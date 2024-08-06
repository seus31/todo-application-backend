package services

import (
	"context"
	"github.com/seus31/todo-application/backend/interfaces"
	"github.com/seus31/todo-application/backend/models"
)

type CategoryService struct {
	categoryRepo interfaces.CategoryRepositoryInterface
}

func NewCategoryService(repo interfaces.CategoryRepositoryInterface) *CategoryService {
	return &CategoryService{
		categoryRepo: repo,
	}
}

func (s *CategoryService) CreateCategory(ctx context.Context, category *models.Category) error {
	return s.categoryRepo.Create(ctx, category)
}
