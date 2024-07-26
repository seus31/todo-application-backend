package requests

type GetUserRequest struct {
	ID uint `params:"id" validate:"required,min=1"`
}
