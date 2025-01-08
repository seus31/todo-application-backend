package categories

import "github.com/go-playground/validator/v10"

type UpdateCategoryRequest struct {
	CategoryName string `json:"category_name" validate:"required,max=255"`
}

var updateCategoryRequestValidate *validator.Validate

func init() {
	updateCategoryRequestValidate = validator.New()
}

func UpdateCategoryRequestValidator() *validator.Validate {
	return updateCategoryRequestValidate
}

func (r *UpdateCategoryRequest) Validate() error {
	return updateCategoryRequestValidate.Struct(r)
}
