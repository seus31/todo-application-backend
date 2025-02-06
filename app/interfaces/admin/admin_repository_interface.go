package admin_interfaces

import (
	"context"
	"github.com/seus31/todo-application-backend/models/admin"
)

type AdminRepositoryInterface interface {
	GetAdmins(ctx context.Context, limit int, offset int) ([]*admin_models.Admin, error)
	Create(ctx context.Context, user *admin_models.Admin) error
	GetAdminByID(ctx context.Context, id uint) (*admin_models.Admin, error)
	Update(ctx context.Context, user *admin_models.Admin) error
	Delete(ctx context.Context, user *admin_models.Admin) error
	FindAdminByEmail(ctx context.Context, email string) (*admin_models.Admin, error)
}
