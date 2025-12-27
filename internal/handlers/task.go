package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
	var req models.TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Title is required, title and description must be string, is_done must be bool", http.StatusBadRequest)
		return
	}
	if req.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	createdTask, err := t.taskService.Create(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
	tasks := t.taskService.List()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (t *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	task, err := t.taskService.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (t *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var req models.TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Title is required, title and description must be string, is_done must be bool", http.StatusBadRequest)
		return
	}
	if req.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	task, err := t.taskService.Update(id, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (t *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = t.taskService.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getID(path string) (int, error) {
	path = strings.Trim(path, "/")
	parts := strings.Split(path, "/")
	id, err := strconv.Atoi(parts[1])
	if (err != nil) || (id < 0) {
		return 0, fmt.Errorf("Invalid ID")
	}
	return id, nil
}
