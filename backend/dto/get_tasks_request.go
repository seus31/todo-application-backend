package dto

type GetTasksRequest struct {
	Limit int `query:"limit" validate:"required,min=1,max=100"`
	Page  int `query:"page" validate:"required,min=1"`
}
