package repository

import (
	"context"
	"github.com/seus31/todo-application-backend/interfaces"
	"github.com/seus31/todo-application-backend/models"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) interfaces.TaskRepositoryInterface {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetTasks(ctx context.Context, userId uint, limit int, offset int) ([]*models.Task, error) {
	var tasks []*models.Task
	if err := r.db.WithContext(ctx).Where("user_id = ?", userId).Limit(limit).Offset(offset).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) Create(ctx context.Context, task *models.Task) error {
	return r.db.WithContext(ctx).Create(task).Error
}

func (r *TaskRepository) GetTaskByID(ctx context.Context, id uint) (*models.Task, error) {
	var task models.Task
	if err := r.db.WithContext(ctx).First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) Update(ctx context.Context, task *models.Task) error {
	return r.db.WithContext(ctx).Model(task).Save(task).Error
}

func (r *TaskRepository) Delete(ctx context.Context, task *models.Task) error {
	return r.db.WithContext(ctx).Delete(&task, task.ID).Error
}
