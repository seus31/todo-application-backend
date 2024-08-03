package users

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,max=255"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=8,max=64,containsany=!@#$%^&*()_+-=[]{};:"`
}
