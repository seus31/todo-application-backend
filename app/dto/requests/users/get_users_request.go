package users

import "github.com/go-playground/validator/v10"

type GetUsersRequest struct {
	Limit int `query:"limit" validate:"required,min=1,max=100"`
	Page  int `query:"page" validate:"required,min=1"`
}

var getUsersRequestValidate *validator.Validate

func init() {
	getUsersRequestValidate = validator.New()
}

func GetUsersRequestValidator() *validator.Validate {
	return getUsersRequestValidate
}

func (r *GetUsersRequest) Validate() error {
	return getUsersRequestValidate.Struct(r)
}
