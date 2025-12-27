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

func (s *TaskService) Create(reqTask models.TaskRequest) (*models.Task, error) {
	var tempDescr string
	if reqTask.Description != nil {
		tempDescr = *reqTask.Description
	}
	var tempIsDone bool
	if reqTask.IsDone != nil {
		tempIsDone = *reqTask.IsDone
	}
	resp := &models.Task{
		ID:          s.curID,
		Title:       reqTask.Title,
		Description: tempDescr,
		IsDone:      tempIsDone,
	}
	s.data.Store(s.curID, resp)
	s.curID++
	return resp, nil
}

func (s *TaskService) Update(id int, reqTask models.TaskRequest) (*models.Task, error) {
	curData, ok := s.data.Load(id)
	if !ok {
		return nil, fmt.Errorf("Not found")
	}
	tempDescr := curData.(*models.Task).Description
	if reqTask.Description != nil {
		tempDescr = *reqTask.Description
	}
	tempIsDone := curData.(*models.Task).IsDone
	if reqTask.IsDone != nil {
		tempIsDone = *reqTask.IsDone
	}
	resp := &models.Task{
		ID:          id,
		Title:       reqTask.Title,
		Description: tempDescr,
		IsDone:      tempIsDone,
	}
	s.data.Store(id, resp)
	return resp, nil
}

func (s *TaskService) Delete(id int) error {
	_, ok := s.data.Load(id)
	if !ok {
		return fmt.Errorf("Not Found")
	}
	s.data.Delete(id)
	return nil
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
