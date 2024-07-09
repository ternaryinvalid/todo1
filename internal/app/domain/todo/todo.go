package todo

type Task struct {
	TaskID        int    `json:"task_id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	EstimatedDate string `json:"estimated_date"`
	Done          bool   `json:"done"`
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

type DoneTaskRequest struct {
	UserID int
	TaskID int `json:"task_id"`
}
