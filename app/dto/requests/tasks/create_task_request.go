package tasks

type CreateTaskRequest struct {
	TaskName string `json:"task_name" validate:"required,max=255"`
}
