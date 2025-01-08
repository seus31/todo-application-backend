package users

import "github.com/go-playground/validator/v10"

type UpdateUserRequest struct {
	Name     *string `json:"name" validate:"omitempty,max=255"`
	Email    *string `json:"email" validate:"omitempty,email,max=255"`
	Password *string `json:"password" validate:"omitempty,min=8,max=64,containsany=!@#$%^&*()_+-=[]{};:"`
}

var updateUserRequestValidate *validator.Validate

func init() {
	updateUserRequestValidate = validator.New()
}

func UpdateUserRequestValidator() *validator.Validate {
	return updateUserRequestValidate
}

func (r *UpdateUserRequest) Validate() error {
	return updateUserRequestValidate.Struct(r)
}
