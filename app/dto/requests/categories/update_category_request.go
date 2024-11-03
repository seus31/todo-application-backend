package categories

type UpdateCategoryRequest struct {
	CategoryName string `json:"category_name" validate:"required,max=255"`
}
