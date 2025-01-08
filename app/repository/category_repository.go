package repository

import (
	"context"
	"github.com/seus31/todo-application-backend/interfaces"
	"github.com/seus31/todo-application-backend/models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) interfaces.CategoryRepositoryInterface {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(ctx context.Context, category *models.Category) error {
	return r.db.WithContext(ctx).Create(category).Error
}

func (r *CategoryRepository) GetCategories(ctx context.Context, limit int, offset int) ([]*models.Category, error) {
	var categories []*models.Category
	if err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) GetCategoryByID(ctx context.Context, id uint) (*models.Category, error) {
	var category models.Category
	if err := r.db.WithContext(ctx).First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) Update(ctx context.Context, category *models.Category) error {
	return r.db.WithContext(ctx).Model(category).Updates(category).Error
}

func (r *CategoryRepository) Delete(ctx context.Context, category *models.Category) error {
	return r.db.WithContext(ctx).Delete(&category, category.ID).Error
}
