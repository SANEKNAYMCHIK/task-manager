package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	customerrors "github.com/SANEKNAYMCHIK/task-manager/internal/custom_errors"
	"github.com/SANEKNAYMCHIK/task-manager/internal/models"
	"github.com/SANEKNAYMCHIK/task-manager/internal/services"
)

type TaskHandler struct {
	taskService services.TaskServiceInterface
}

func NewTaskHandler(taskService services.TaskServiceInterface) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func (t *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req models.TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Title is required, title and description must be string, is_done must be bool", http.StatusBadRequest)
		return
	}

	createdTask, err := t.taskService.Create(req)
	if err != nil && errors.Is(err, customerrors.ErrTitleIsRequired) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteResponse(w, createdTask, http.StatusCreated)
}

func (t *TaskHandler) ListTask(w http.ResponseWriter, r *http.Request) {
	tasks := t.taskService.List()
	WriteResponse(w, tasks, http.StatusOK)
}

func (t *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r.URL.Path)
	if err != nil && errors.Is(err, customerrors.ErrInvalidID) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	task, err := t.taskService.Get(id)
	if err != nil {
		switch err {
		case customerrors.ErrTaskNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		case customerrors.ErrInvalidData:
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}
	WriteResponse(w, task, http.StatusOK)
}

func (t *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r.URL.Path)
	if err != nil && errors.Is(err, customerrors.ErrInvalidID) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var req models.TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Title is required, title and description must be string, is_done must be bool", http.StatusBadRequest)
		return
	}

	task, err := t.taskService.Update(id, req)
	if err != nil {
		switch err {
		case customerrors.ErrTitleIsRequired:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case customerrors.ErrTaskNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		return
	}
	WriteResponse(w, task, http.StatusOK)
}

func (t *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r.URL.Path)
	if err != nil && errors.Is(err, customerrors.ErrInvalidID) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = t.taskService.Delete(id)
	if err != nil && errors.Is(err, customerrors.ErrTaskNotFound) {
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
		return 0, customerrors.ErrInvalidID
	}
	return id, nil
}

func WriteResponse(w http.ResponseWriter, val any, httpCodeSuccess int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCodeSuccess)
	if err := json.NewEncoder(w).Encode(val); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
