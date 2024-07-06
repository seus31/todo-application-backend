package requests

type GetTaskRequest struct {
	ID uint `params:"id" validate:"required,min=1"`
}
