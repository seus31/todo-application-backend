package categories

import "github.com/go-playground/validator/v10"

type GetCategoriesRequest struct {
	Limit int `query:"limit" validate:"required,min=1,max=100"`
	Page  int `query:"page" validate:"required,min=1"`
}

var getCategoriesRequestValidate *validator.Validate

func init() {
	getCategoriesRequestValidate = validator.New()
}

func GetCategoriesRequestValidator() *validator.Validate {
	return getCategoriesRequestValidate
}

func (r *GetCategoriesRequest) Validate() error {
	return getCategoriesRequestValidate.Struct(r)
}
