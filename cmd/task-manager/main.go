package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/SANEKNAYMCHIK/task-manager/internal/handlers"
	"github.com/SANEKNAYMCHIK/task-manager/internal/services"
)

func main() {
	serverPort := "8080"
	data := &sync.Map{}
	taskService := services.NewTaskService(data)
	mux := handlers.NewRouter(taskService)
	if err := http.ListenAndServe(":"+serverPort, mux); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}

}
