package tasks

import (
	"github.com/go-playground/validator/v10"
	"github.com/seus31/todo-application-backend/models"
)

type CreateTaskRequest struct {
	TaskName string            `json:"task_name" validate:"required,max=255"`
	ParentID *uint             `json:"parent_id" validate:"omitempty,numeric"`
	DueDate  *string           `json:"due_date" validate:"omitempty,datetime=2006-01-02"`
	DueTime  *string           `json:"due_time" validate:"omitempty,datetime=15:04"`
	Status   models.TaskStatus `json:"status" validate:"required,valid_task_status"`
	Priority models.Priority   `json:"priority" validate:"required,valid_priority"`
}

var createTaskRequestValidate *validator.Validate

func init() {
	createTaskRequestValidate = models.GetTaskValidator()
}

func CreateTaskRequestValidator() *validator.Validate {
	return createTaskRequestValidate
}

func (r *CreateTaskRequest) Validate() error {
	return createTaskRequestValidate.Struct(r)
}
