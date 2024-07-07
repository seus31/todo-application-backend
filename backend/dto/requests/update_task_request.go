package requests

type UpdateTaskRequest struct {
	TaskName string `json:"task_name" validate:"required,max=255"`
}
