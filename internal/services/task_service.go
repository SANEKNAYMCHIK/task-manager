package services

import (
	"fmt"
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

func (s *TaskService) Get(id int) (*models.Task, error) {
	val, ok := s.data.Load(id)
	if !ok {
		return nil, fmt.Errorf("Not found")
	}
	task, ok := val.(*models.Task)
	if !ok {
		return nil, fmt.Errorf("Invalid task data")
	}
	return task, nil
}

func (s *TaskService) List() []*models.Task {
	res := []*models.Task{}
	s.data.Range(func(key, value any) bool {
		if task, ok := value.(*models.Task); ok {
			res = append(res, task)
		}
		return true
	})
	return res
}
