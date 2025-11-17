package database

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func setupTestDatabase(t *testing.T) {
	if err := godotenv.Load("../../configs/.env"); err != nil {
		log.Fatal("Error loading .env file:", err)
		os.Exit(1)
	}

	status := make(chan bool)
	go InitDatabase(status)
	if success := <-status; !success {
		t.Fatal("failed to initialize database")
	}
}

func teardownTestDatabase(t *testing.T) {
	if err := CloseDatabase(); err != nil {
		t.Fatalf("failed to close database: %v", err)
	}
}

func TestCreateTask(t *testing.T) {
	setupTestDatabase(t)
	defer teardownTestDatabase(t)

	taskName := "TestTask"
	idTask, err := CreateTask(taskName)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if idTask == 0 {
		t.Fatal("expected valid task ID, got 0")
	}

	DeleteTask(idTask)
}

func TestGetTaskByName(t *testing.T) {
	setupTestDatabase(t)
	defer teardownTestDatabase(t)

	taskName := "TestTask"
	_, err := CreateTask(taskName)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	task, err := GetTaskByName(taskName)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if task.Name != taskName {
		t.Fatalf("expected task name %s, got %s", taskName, task.Name)
	}
	DeleteTask(task.Id)
}

func TestUpdateTaskStatus(t *testing.T) {
	setupTestDatabase(t)
	defer teardownTestDatabase(t)

	taskName := "TestTask"
	idTask, err := CreateTask(taskName)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = UpdateTaskStatus(idTask, true)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	task, err := GetTaskByName(taskName)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !task.Done {
		t.Fatal("expected task to be marked as done")
	}
	DeleteTask(idTask)
}

func TestDeleteTask(t *testing.T) {
	setupTestDatabase(t)
	defer teardownTestDatabase(t)

	taskName := "TestTask"
	idTask, err := CreateTask(taskName)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = DeleteTask(idTask)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = GetTaskByName(taskName)
	if err == nil {
		t.Fatal("expected error for non-existent task, got nil")
	}
}

func TestGetAllTasks(t *testing.T) {
	setupTestDatabase(t)
	defer teardownTestDatabase(t)

	taskNames := []string{"Task1", "Task2", "Task3"}
	for _, name := range taskNames {
		_, err := CreateTask(name)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	}

	tasks, err := GetAllTasks()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(tasks) < len(taskNames) {
		t.Fatalf("expected at least %d tasks, got %d", len(taskNames), len(tasks))
	}

	for _, task := range tasks {
		DeleteTask(task.Id)
	}
}
