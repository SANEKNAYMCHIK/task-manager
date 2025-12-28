package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/SANEKNAYMCHIK/task-manager/internal/handlers"
	"github.com/SANEKNAYMCHIK/task-manager/internal/services"
)

func main() {
	serverPort := ":8080"
	data := &sync.Map{}
	taskService := services.NewTaskService(data)
	handler := handlers.NewRouter(taskService)
	httpServer := &http.Server{
		Addr:    serverPort,
		Handler: handler,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v\n", err)
		}
	}()
	<-stop
	log.Println("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}
	wg.Wait()
	log.Println("Server stopped")
}
