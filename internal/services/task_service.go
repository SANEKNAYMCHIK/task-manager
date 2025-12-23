package services

type TaskService struct {
	_ int
}

func NewTaskService() *TaskService {
	return &TaskService{}
}
