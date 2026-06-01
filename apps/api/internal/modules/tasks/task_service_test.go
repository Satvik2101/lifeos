package tasks

import (
	"errors"
	"testing"

	api "lifeos-api/internal/server/api"
)

func TestCreateTask(t *testing.T) {
	repo := NewFakeTaskRepository()
	service := NewTaskService(repo)

	task := service.Create(api.CreateTaskRequest{
		Title: "Test task",
	})

	if task.Id == "" {
		t.Fatal("expected task ID to be set")
	}

	if task.Title != "Test task" {
		t.Fatalf("expected title %q, got %q", "Test task", task.Title)
	}
}

func TestGetTask(t *testing.T) {
	repo := NewFakeTaskRepository()
	service := NewTaskService(repo)

	created := service.Create(api.CreateTaskRequest{
		Title: "Test task",
	})

	task, err := service.Get(created.Id)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if task.Id != created.Id {
		t.Fatalf("expected id %q, got %q", created.Id, task.Id)
	}
}

func TestGetTaskNotFound(t *testing.T) {
	repo := NewFakeTaskRepository()
	service := NewTaskService(repo)

	_, err := service.Get("missing")

	if !errors.Is(err, ErrTaskNotFound) {
		t.Fatalf("expected ErrTaskNotFound, got %v", err)
	}
}

func TestListTasks(t *testing.T) {
	repo := NewFakeTaskRepository()
	service := NewTaskService(repo)

	title1 := "Task 1"
	title2 := "Task 2"
	service.Create(api.CreateTaskRequest{
		Title: title1,
	})

	service.Create(api.CreateTaskRequest{
		Title: title2,
	})

	tasks := service.List()

	if len(tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(tasks))
	}

	if (tasks[0].Title != title1 && tasks[1].Title != title2) && (tasks[0].Title != title2 && tasks[1].Title != title1) {
		t.Fatalf("expected tasks with titles %q and %q, got %q and %q", title1, title2, tasks[0].Title, tasks[1].Title)
	}

}

func TestUpdateTask(t *testing.T) {
	repo := NewFakeTaskRepository()
	service := NewTaskService(repo)

	created := service.Create(api.CreateTaskRequest{
		Title: "Old title",
	})

	updated, err := service.Update(created.Id, api.UpdateTaskRequest{
		Title: new("New title"),
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if updated.Title != "New title" {
		t.Fatalf("expected title %q, got %q", "New title", updated.Title)
	}
}

func TestUpdateTaskNotFound(t *testing.T) {
	repo := NewFakeTaskRepository()
	service := NewTaskService(repo)

	_, err := service.Update("missing", api.UpdateTaskRequest{
		Title: new("New title"),
	})

	if !errors.Is(err, ErrTaskNotFound) {
		t.Fatalf("expected ErrTaskNotFound, got %v", err)
	}
}

func TestDeleteTask(t *testing.T) {
	repo := NewFakeTaskRepository()
	service := NewTaskService(repo)

	created := service.Create(api.CreateTaskRequest{
		Title: "Delete me",
	})

	err := service.Delete(created.Id)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err = service.Get(created.Id)

	if !errors.Is(err, ErrTaskNotFound) {
		t.Fatalf("expected ErrTaskNotFound, got %v", err)
	}
}

func TestDeleteTaskNotFound(t *testing.T) {
	repo := NewFakeTaskRepository()
	service := NewTaskService(repo)

	err := service.Delete("missing")

	if !errors.Is(err, ErrTaskNotFound) {
		t.Fatalf("expected ErrTaskNotFound, got %v", err)
	}
}
