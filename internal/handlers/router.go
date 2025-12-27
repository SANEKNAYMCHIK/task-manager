package handlers

import (
	"net/http"

	"github.com/SANEKNAYMCHIK/task-manager/internal/services"
)

func NewRouter(taskService *services.TaskService) *http.ServeMux {
	mux := http.NewServeMux()
	taskHandler := NewTaskHandler(taskService)

	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			// POST Создать новую задачу
			taskHandler.CreateTask(w, r)
		case http.MethodGet:
			// GET Получить список всех задач
			taskHandler.ListTask(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// GET Получить задачу по идентификатору
			taskHandler.GetTask(w, r, "/todos/")
		case http.MethodPut:
			// PUT Обновить задачу по идентификатору
			taskHandler.UpdateTask(w, r)
		case http.MethodDelete:
			// DELETE Удалить задачу по идентификатору
			taskHandler.DeleteTask(w, r)
		}
	})

	return mux
}
