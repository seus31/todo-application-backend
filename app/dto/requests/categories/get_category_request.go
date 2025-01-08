package categories

import "github.com/go-playground/validator/v10"

type GetCategoryRequest struct {
	ID uint `params:"id" validate:"required,min=1"`
}

var getCategoryRequestValidate *validator.Validate

func init() {
	getCategoryRequestValidate = validator.New()
}

func GetCategoryRequestValidator() *validator.Validate {
	return getCategoryRequestValidate
}

func (r *GetCategoryRequest) Validate() error {
	return getCategoryRequestValidate.Struct(r)
}
