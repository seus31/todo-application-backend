package users

import "github.com/go-playground/validator/v10"

type CreateUserRequest struct {
	Name            string `json:"name" validate:"required,max=255"`
	Email           string `json:"email" validate:"required,email,max=255"`
	Password        string `json:"password" validate:"required,min=8,max=64,containsany=!@#$%^&*()_+-=[]{};:"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

var createUserRequestValidate *validator.Validate

func init() {
	createUserRequestValidate = validator.New()
}

func CreateUserRequestValidator() *validator.Validate {
	return createUserRequestValidate
}

func (r *CreateUserRequest) Validate() error {
	return createUserRequestValidate.Struct(r)
}
