package admin_auth

import "github.com/go-playground/validator/v10"

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

var loginRequestValidate *validator.Validate

func init() {
	loginRequestValidate = validator.New()
}

func LoginRequestValidator() *validator.Validate {
	return loginRequestValidate
}

func (r *LoginRequest) Validate() error {
	return loginRequestValidate.Struct(r)
}
