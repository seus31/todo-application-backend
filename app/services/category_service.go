package services

import (
	"context"
	"github.com/seus31/todo-application-backend/dto/requests/categories"
	"github.com/seus31/todo-application-backend/interfaces"
	"github.com/seus31/todo-application-backend/models"
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

func (s *CategoryService) GetCategories(ctx context.Context, limit int, offset int) ([]*models.Category, error) {
	return s.categoryRepo.GetCategories(ctx, limit, offset)
}

func (s *CategoryService) GetCategory(ctx context.Context, id uint) (*models.Category, error) {
	return s.categoryRepo.GetCategoryByID(ctx, id)
}

func (s *CategoryService) UpdateCategory(ctx context.Context, category *models.Category, req categories.UpdateCategoryRequest) (*models.Category, error) {
	category.CategoryName = req.CategoryName
	err := s.categoryRepo.Update(ctx, category)
	if err != nil {
		return nil, err
	}

	updateTask, err := s.categoryRepo.GetCategoryByID(ctx, category.ID)

	return updateTask, err
}

func (s *CategoryService) DeleteCategory(ctx context.Context, category *models.Category) error {
	return s.categoryRepo.Delete(ctx, category)
}
