package tasks

import (
	"github.com/go-playground/validator/v10"
)

type GetTasksRequest struct {
	Limit int `query:"limit" validate:"required,min=1,max=100"`
	Page  int `query:"page" validate:"required,min=1"`
}

var getTasksRequestValidate *validator.Validate

func init() {
	getTasksRequestValidate = validator.New()
}

func GetTasksRequestValidator() *validator.Validate {
	return getTasksRequestValidate
}

func (r *GetTasksRequest) Validate() error {
	return getTasksRequestValidate.Struct(r)
}
