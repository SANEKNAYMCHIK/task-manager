package services

import "github.com/SANEKNAYMCHIK/task-manager/internal/models"

type TaskServiceInterface interface {
	Create(reqTask models.TaskRequest) (*models.Task, error)
	Update(id int, reqTask models.TaskRequest) (*models.Task, error)
	Delete(id int) error
	Get(id int) (*models.Task, error)
	List() []*models.Task
}
