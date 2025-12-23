package handlers

import (
	"net/http"

	"github.com/SANEKNAYMCHIK/task-manager/internal/services"
)

type TaskHandler struct {
	taskService *services.TaskService
}

func NewTaskHandler(taskService *services.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func (t *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	return
}

func (t *TaskHandler) ListTask(w http.ResponseWriter, r *http.Request) {
	return
}

func (t *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	return
}

func (t *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	return
}

func (t *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	return
}
