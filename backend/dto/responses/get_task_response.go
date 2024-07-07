package responses

type TaskResponse struct {
	ID        uint   `json:"id"`
	TaskName  string `json:"task_name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
