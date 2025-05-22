package tasks

type CreateTasksRequest struct {
	Title    string `json:"title" validate:"required,min=4,max=100"`
	Subtitle string `json:"subtitle"`
	UserID   uint   `json:"userid" validate:"required"`
}

type UpdateTasksRequest struct {
	Title     string `json:"title"`
	Subtitle  string `json:"subtitle"`
	Completed *bool  `json:"completed"`
}

type StatusTaskRequest struct {
	Completed *bool `json:"completed"`
}
