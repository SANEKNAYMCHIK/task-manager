package services

import (
	"sync"
	"sync/atomic"

	customerrors "github.com/SANEKNAYMCHIK/task-manager/internal/custom_errors"
	"github.com/SANEKNAYMCHIK/task-manager/internal/models"
)

type TaskService struct {
	data  *sync.Map
	curID int64
}

func NewTaskService(data *sync.Map) *TaskService {
	return &TaskService{
		data:  data,
		curID: 0,
	}
}

func (s *TaskService) Create(reqTask models.TaskRequest) (*models.Task, error) {
	if reqTask.Title == "" {
		return nil, customerrors.ErrTitleIsRequired
	}
	var tempDescr string
	if reqTask.Description != nil {
		tempDescr = *reqTask.Description
	}
	var tempIsDone bool
	if reqTask.IsDone != nil {
		tempIsDone = *reqTask.IsDone
	}

	id := atomic.AddInt64(&s.curID, 1)
	resp := &models.Task{
		ID:          int(id),
		Title:       reqTask.Title,
		Description: tempDescr,
		IsDone:      tempIsDone,
	}
	s.data.Store(int(id), resp)
	return resp, nil
}

func (s *TaskService) Update(id int, reqTask models.TaskRequest) (*models.Task, error) {
	if reqTask.Title == "" {
		return nil, customerrors.ErrTitleIsRequired
	}
	curData, ok := s.data.Load(id)
	if !ok {
		return nil, customerrors.ErrTaskNotFound
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
		return customerrors.ErrTaskNotFound
	}
	s.data.Delete(id)
	return nil
}

func (s *TaskService) Get(id int) (*models.Task, error) {
	val, ok := s.data.Load(id)
	if !ok {
		return nil, customerrors.ErrTaskNotFound
	}
	task, ok := val.(*models.Task)
	if !ok {
		return nil, customerrors.ErrInvalidData
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
