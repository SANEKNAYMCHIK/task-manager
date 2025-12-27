package models

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
}

type TaskRequest struct {
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	IsDone      *bool   `json:"is_done,omitempty"`
}
