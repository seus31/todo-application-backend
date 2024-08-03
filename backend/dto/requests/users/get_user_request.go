package users

type GetUserRequest struct {
	ID uint `params:"id" validate:"required,min=1"`
}
