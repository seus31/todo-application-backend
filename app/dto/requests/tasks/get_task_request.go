package tasks

import (
	"github.com/go-playground/validator/v10"
)

type GetTaskRequest struct {
	ID int `params:"id" validate:"required,min=1"`
}

var getTaskRequestValidate *validator.Validate

func init() {
	getTaskRequestValidate = validator.New()
}

func GetTaskRequestValidator() *validator.Validate {
	return getTaskRequestValidate
}

func (r *GetTaskRequest) Validate() error {
	return getTaskRequestValidate.Struct(r)
}
