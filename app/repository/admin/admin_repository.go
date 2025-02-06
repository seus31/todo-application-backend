package admin_repository

import (
	"context"
	"github.com/seus31/todo-application-backend/interfaces/admin"
	"github.com/seus31/todo-application-backend/models/admin"
	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) admin_interfaces.AdminRepositoryInterface {
	return &AdminRepository{db: db}
}

func (r *AdminRepository) GetAdmins(ctx context.Context, limit int, offset int) ([]*admin_models.Admin, error) {
	var admin []*admin_models.Admin
	if err := r.db.Limit(limit).Offset(offset).Find(&admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *AdminRepository) Create(ctx context.Context, admin *admin_models.Admin) error {
	return r.db.WithContext(ctx).Create(admin).Error
}

func (r *AdminRepository) GetAdminByID(ctx context.Context, id uint) (*admin_models.Admin, error) {
	var admin admin_models.Admin
	if err := r.db.WithContext(ctx).First(&admin, id).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *AdminRepository) Update(ctx context.Context, admin *admin_models.Admin) error {
	return r.db.WithContext(ctx).Model(admin).Save(admin).Error
}

func (r *AdminRepository) Delete(ctx context.Context, admin *admin_models.Admin) error {
	return r.db.WithContext(ctx).Delete(&admin, admin.ID).Error
}

func (r *AdminRepository) FindAdminByEmail(ctx context.Context, email string) (*admin_models.Admin, error) {
	var admin admin_models.Admin
	if err := r.db.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}
