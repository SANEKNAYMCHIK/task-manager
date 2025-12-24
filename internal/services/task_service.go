package services

import "sync"

type TaskService struct {
	data *sync.Map
}

func NewTaskService(data *sync.Map) *TaskService {
	return &TaskService{data: data}
}

func (s *TaskService) Create() {
	return
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
