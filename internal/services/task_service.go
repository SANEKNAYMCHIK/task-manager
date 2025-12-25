package services

import (
	"sync"

	"github.com/SANEKNAYMCHIK/task-manager/internal/models"
)

type TaskService struct {
	data  *sync.Map
	curID int
}

func NewTaskService(data *sync.Map) *TaskService {
	return &TaskService{
		data:  data,
		curID: 1,
	}
}

func (s *TaskService) Create(reqTask models.CreateTaskRequest) (*models.Task, error) {
	resp := &models.Task{
		ID:          s.curID,
		Title:       reqTask.Title,
		Description: reqTask.Description,
		IsDone:      false,
	}
	s.data.Store(s.curID, resp)
	s.curID++
	return resp, nil
}

func (s *TaskService) Update() {
	return
}

func (s *TaskService) Delete() {
	return
}

func (s *TaskService) Get() {
	return
}

func (s *TaskService) List() {
	return
}
