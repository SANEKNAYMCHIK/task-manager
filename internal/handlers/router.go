package handlers

import (
	"net/http"

	"github.com/SANEKNAYMCHIK/task-manager/internal/services"
)

func NewRouter(taskService *services.TaskService) *http.ServeMux {
	mux := http.NewServeMux()
	taskHandler := NewTaskHandler(taskService)

	// POST Создать новую задачу
	mux.HandleFunc("/todos", taskHandler.CreateTask)

	// GET Получить список всех задач
	mux.HandleFunc("/todos", taskHandler.ListTask)

	// GET Получить задачу по идентификатору
	mux.HandleFunc("/todos/{id}", taskHandler.GetTask)

	// PUT Обновить задачу по идентификатору
	mux.HandleFunc("/todos/{id}", taskHandler.UpdateTask)

	// DELETE Удалить задачу по идентификатору
	mux.HandleFunc("/todos/{id}", taskHandler.DeleteTask)

	return mux
}
