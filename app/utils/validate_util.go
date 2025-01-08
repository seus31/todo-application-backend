package utils

import (
	"github.com/go-playground/validator/v10"
)

func ValidateStruct(validate *validator.Validate, s interface{}) error {
	return validate.Struct(s)
}
