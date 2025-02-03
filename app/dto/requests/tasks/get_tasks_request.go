package tasks

import (
	"github.com/go-playground/validator/v10"
)

type GetTasksRequest struct {
	Limit int `query:"limit" validate:"omitempty,min=1,max=100" default:"20"`
	Page  int `query:"page" validate:"omitempty,min=1" default:"1"`
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
