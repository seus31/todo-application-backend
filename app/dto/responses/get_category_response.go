package responses

type CategoryResponse struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
