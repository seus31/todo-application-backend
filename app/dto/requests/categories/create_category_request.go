package categories

import "github.com/go-playground/validator/v10"

type CreateCategoryRequest struct {
	CategoryName string `json:"category_name" validate:"required,max=255"`
}

var createCategoryRequestValidate *validator.Validate

func init() {
	createCategoryRequestValidate = validator.New()
}

func CreateCategoryRequestValidator() *validator.Validate {
	return createCategoryRequestValidate
}

func (r *CreateCategoryRequest) Validate() error {
	return createCategoryRequestValidate.Struct(r)
}
