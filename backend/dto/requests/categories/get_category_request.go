package categories

type GetCategoryRequest struct {
	ID uint `params:"id" validate:"required,min=1"`
}
