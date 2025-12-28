package services

import (
	"strings"
	"sync"
	"testing"

	"github.com/SANEKNAYMCHIK/task-manager/internal/models"
)

func TestTaskService_Create(t *testing.T) {
	tests := []struct {
		name       string
		req        models.TaskRequest
		wantID     int
		wantError  bool
		wantIsDone bool
	}{
		{
			name: "Create task with all fields",
			req: models.TaskRequest{
				Title:       "Test task 1",
				Description: stringPtr("Test description"),
				IsDone:      boolPtr(false),
			},
			wantIsDone: false,
			wantID:     1,
			wantError:  false,
		},
		{
			name: "Create task with only Description",
			req: models.TaskRequest{
				Title:       "Test task 2",
				Description: stringPtr("Test description"),
			},
			wantIsDone: false,
			wantID:     2,
			wantError:  false,
		},
		{
			name: "Create task with only IsDone",
			req: models.TaskRequest{
				Title:  "Test task 3",
				IsDone: boolPtr(true),
			},
			wantIsDone: true,
			wantID:     3,
			wantError:  false,
		},
		{
			name: "Create task with only Title",
			req: models.TaskRequest{
				Title: "Test task 4",
			},
			wantIsDone: false,
			wantID:     4,
			wantError:  false,
		},
		{
			name:       "Create task without any fields",
			req:        models.TaskRequest{},
			wantIsDone: false,
			wantID:     5,
			wantError:  true,
		},
	}
	data := &sync.Map{}
	service := NewTaskService(data)
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task, err := service.Create(tt.req)
			if tt.wantError && err == nil {
				t.Errorf("Test %d: expected error, got nil", i+1)
			}
			if !tt.wantError && err != nil {
				t.Errorf("Test %d: unexpected error", i+1)
			}
			if !tt.wantError && task != nil {
				if task.ID != i+1 {
					t.Errorf("Test %d: expected ID: %d, but got %d", i+1, tt.wantID, task.ID)
				}
				if task.Title != tt.req.Title {
					t.Errorf("Test %d: expected Title: %s, but got %s", i+1, tt.req.Title, task.Title)
				}
				if task.IsDone != tt.wantIsDone {
					t.Errorf("Test %d: expected IsDone: %t, but got %t", i+1, *tt.req.IsDone, task.IsDone)
				}
				val, ok := data.Load(task.ID)
				if !ok {
					t.Errorf("Test %d: task not found in data store", i+1)
				}
				storedTask, ok := val.(*models.Task)
				if !ok {
					t.Errorf("Test %d: stored value is not *models.Task", i)
				} else if storedTask.Title != task.Title {
					t.Errorf("Test %d: stored task has wrong title", i)
				}
			}
		})
	}
}

func TestTaskService_Update(t *testing.T) {
	data := &sync.Map{}
	service := NewTaskService(data)
	req := models.TaskRequest{
		Title:       "Original title",
		Description: stringPtr("Original description"),
		IsDone:      boolPtr(false),
	}
	createdTask1, _ := service.Create(req)
	tests := []struct {
		name            string
		id              int
		req             models.TaskRequest
		wantError       bool
		wantIsDone      bool
		wantDescription string
	}{
		{
			name: "Update all fields",
			id:   createdTask1.ID,
			req: models.TaskRequest{
				Title:       "New title",
				Description: stringPtr("New description"),
				IsDone:      boolPtr(true),
			},
			wantError:       false,
			wantIsDone:      true,
			wantDescription: "New description",
		},
		{
			name: "Update only title",
			id:   createdTask1.ID,
			req: models.TaskRequest{
				Title: "Some title",
			},
			wantError:       false,
			wantIsDone:      true,
			wantDescription: "New description",
		},
		{
			name: "Update only description",
			id:   createdTask1.ID,
			req: models.TaskRequest{
				Title:       "Some title",
				Description: stringPtr("Extra new description"),
			},
			wantError:       false,
			wantIsDone:      true,
			wantDescription: "Extra new description",
		},
		{
			name: "Update only IsDone",
			id:   createdTask1.ID,
			req: models.TaskRequest{
				Title:  "Some title",
				IsDone: boolPtr(false),
			},
			wantError:       false,
			wantIsDone:      false,
			wantDescription: "Extra new description",
		},
		{
			name: "Update non existing task",
			id:   0,
			req: models.TaskRequest{
				Title:       "New title",
				Description: stringPtr("New description"),
				IsDone:      boolPtr(true),
			},
			wantError:       true,
			wantIsDone:      false,
			wantDescription: "Error",
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task, err := service.Update(tt.id, tt.req)
			if tt.wantError && err == nil {
				t.Errorf("Test %d: expected error, got nil", i+1)
			}
			if !tt.wantError && err != nil {
				t.Errorf("Test %d: unexpected error", i+1)
			}
			if !tt.wantError && task != nil {
				if task.Title != tt.req.Title {
					t.Errorf("Test %d: expected Title: %s, but got %s", i+1, tt.req.Title, task.Title)
				}
				if task.Description != tt.wantDescription {
					t.Errorf("Test %d: expected Description: %s, but got %s", i+1, tt.wantDescription, task.Description)
				}
				if task.IsDone != tt.wantIsDone {
					t.Errorf("Test %d: expected IsDone: %t, but got %t", i+1, tt.wantIsDone, task.IsDone)
				}
				val, ok := data.Load(task.ID)
				if !ok {
					t.Errorf("Test %d: task not found in data store", i+1)
				}
				storedTask, ok := val.(*models.Task)
				if !ok {
					t.Errorf("Test %d: stored value is not *models.Task", i)
				} else if storedTask.Title != task.Title {
					t.Errorf("Test %d: stored task has wrong title", i)
				}
			}
		})
	}
}

func TestTaskService_Delete(t *testing.T) {
	data := &sync.Map{}
	service := NewTaskService(data)
	req := models.TaskRequest{
		Title: "Task to delete",
	}
	createdTask, _ := service.Create(req)

	t.Run("Delete existing task", func(t *testing.T) {
		err := service.Delete(createdTask.ID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		_, exists := data.Load(createdTask.ID)
		if exists {
			t.Error("Task was not deleted from data store")
		}

		_, err = service.Get(createdTask.ID)
		if err == nil {
			t.Error("Expected error when getting deleted task, got nil")
		}
	})

	t.Run("Delete non-existent task", func(t *testing.T) {
		err := service.Delete(0)
		if err == nil {
			t.Error("Expected error for non existing task, got nil")
		}
	})
}

func TestTaskService_Get(t *testing.T) {
	data := &sync.Map{}
	service := NewTaskService(data)
	req := models.TaskRequest{
		Title:       "Test task 1",
		Description: stringPtr("Test description"),
		IsDone:      boolPtr(false),
	}
	createdTask1, err := service.Create(req)
	if err != nil {
		t.Error("Creating error")
	}
	t.Run("Get existing task", func(t *testing.T) {
		task, err := service.Get(createdTask1.ID)
		if err != nil {
			t.Error("Unexpected error")
		}
		if task.ID != createdTask1.ID {
			t.Errorf("Expected ID: %d, but got %d", createdTask1.ID, task.ID)
		}
		if task.Title != createdTask1.Title {
			t.Errorf("Expected Title: %s, but got %s", createdTask1.Title, task.Title)
		}
	})
	t.Run("Get non existing task", func(t *testing.T) {
		_, err := service.Get(0)
		if err == nil {
			t.Error("Something went wrong")
		}
		if !strings.Contains(err.Error(), "not found") {
			t.Errorf("Expected error: %s, but got %s", "not found", err.Error())
		}
	})
}

func TestTaskService_List(t *testing.T) {
	data := &sync.Map{}
	service := NewTaskService(data)

	initialTasks := service.List()
	if len(initialTasks) != 0 {
		t.Errorf("Expected empty list, got %d tasks", len(initialTasks))
	}

	tasksToCreate := []models.TaskRequest{
		{Title: "Task 1"},
		{Title: "Task 2"},
		{Title: "Task 3"},
	}
	createdIDs := make([]int, 0, len(tasksToCreate))
	for _, req := range tasksToCreate {
		task, _ := service.Create(req)
		createdIDs = append(createdIDs, task.ID)
	}

	tasks := service.List()
	if len(tasks) != len(tasksToCreate) {
		t.Errorf("Expected %d tasks, got %d", len(tasksToCreate), len(tasks))
	}

	foundTasks := make(map[int]bool)
	for _, task := range tasks {
		foundTasks[task.ID] = true
		if task.ID < 1 || task.ID > len(tasksToCreate) {
			t.Errorf("Unexpected task ID: %d", task.ID)
		}
	}
	for _, id := range createdIDs {
		if !foundTasks[id] {
			t.Errorf("Task with ID %d not found in list", id)
		}
	}

	titles := make(map[string]bool)
	for _, task := range tasks {
		if titles[task.Title] {
			t.Errorf("Duplicate title found: %s", task.Title)
		}
		titles[task.Title] = true
	}
}

func TestTaskService_Concurrent(t *testing.T) {
	data := &sync.Map{}
	service := NewTaskService(data)
	numGoroutines := 10
	tasksPerGoroutine := 100

	var wg sync.WaitGroup
	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range tasksPerGoroutine {
				req := models.TaskRequest{
					Title: "Task",
				}
				service.Create(req)
			}
		}()
	}
	wg.Wait()

	tasks := service.List()
	expectedLen := numGoroutines * tasksPerGoroutine
	if len(tasks) != expectedLen {
		t.Errorf("Expected %d tasks, got %d", expectedLen, len(tasks))
	}
}

func stringPtr(val string) *string {
	return &val
}

func boolPtr(val bool) *bool {
	return &val
}
