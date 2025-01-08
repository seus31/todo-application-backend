package users

import "github.com/go-playground/validator/v10"

type GetUserRequest struct {
	ID uint `params:"id" validate:"required,min=1"`
}

var getUserRequestValidate *validator.Validate

func init() {
	getUserRequestValidate = validator.New()
}

func GetUserRequestValidator() *validator.Validate {
	return getUserRequestValidate
}

func (r *GetUserRequest) Validate() error {
	return getUserRequestValidate.Struct(r)
}
