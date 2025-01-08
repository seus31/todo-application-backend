package tasks

import "github.com/go-playground/validator/v10"

type UpdateTaskRequest struct {
	TaskName string `json:"task_name" validate:"required,max=255"`
}

var updateTaskRequestValidate *validator.Validate

func init() {
	updateTaskRequestValidate = validator.New()
}

func UpdateTaskRequestValidator() *validator.Validate {
	return updateTaskRequestValidate
}

func (r *UpdateTaskRequest) Validate() error {
	return updateTaskRequestValidate.Struct(r)
}
