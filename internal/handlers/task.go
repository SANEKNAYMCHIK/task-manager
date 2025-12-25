package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SANEKNAYMCHIK/task-manager/internal/models"
	"github.com/SANEKNAYMCHIK/task-manager/internal/services"
)

type TaskHandler struct {
	taskService *services.TaskService
}

func NewTaskHandler(taskService *services.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func (t *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Title and description are required", http.StatusBadRequest)
		return
	}
	if req.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	createdTask, err := t.taskService.Create(req)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdTask); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
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
