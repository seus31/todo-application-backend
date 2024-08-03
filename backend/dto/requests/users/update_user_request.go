package users

type UpdateUserRequest struct {
	Name     *string `json:"name" validate:"omitempty,max=255"`
	Email    *string `json:"email" validate:"omitempty,email,max=255"`
	Password *string `json:"password" validate:"omitempty,min=8,max=64,containsany=!@#$%^&*()_+-=[]{};:"`
}
