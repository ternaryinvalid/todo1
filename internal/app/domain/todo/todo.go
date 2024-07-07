package todo

type Task struct {
	TaskID        int    `json:"task_id"`
	UserID        int    `json:"user_id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	EstimatedDate string `json:"estimated_date"`
}

type CreateTaskRequest struct {
	UserID        int
	Title         string `json:"title"`
	Description   string `json:"description"`
	EstimatedDate string `json:"estimated_date"`
}

type DeleteTaskRequest struct {
	UserID int
	TaskID int `json:"task_id"`
}
