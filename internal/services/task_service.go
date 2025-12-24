package services

import "sync"

type TaskService struct {
	data *sync.Map
}

func NewTaskService(data *sync.Map) *TaskService {
	return &TaskService{data: data}
}
