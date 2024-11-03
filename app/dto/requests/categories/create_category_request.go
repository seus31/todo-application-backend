package categories

type CreateCategoryRequest struct {
	CategoryName string `json:"category_name" validate:"required,max=255"`
}
