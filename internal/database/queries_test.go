package database_test

import (
	"cli_tasks_api/internal/database"
	"cli_tasks_api/internal/utils"
	"testing"
)

func TestCreateTask(t *testing.T) {
	utils.SetupTestDatabase(t)
	defer utils.TeardownTestDatabase(t)

	taskName := "TestTask"
	idTask, err := database.CreateTask(taskName)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if idTask == 0 {
		t.Fatal("expected valid task ID, got 0")
	}

	err = database.DeleteTask(idTask)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestGetTaskByName(t *testing.T) {
	utils.SetupTestDatabase(t)
	defer utils.TeardownTestDatabase(t)

	taskName := "TestTask"
	_, err := database.CreateTask(taskName)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	task, err := database.GetTaskByName(taskName)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if task.Name != taskName {
		t.Fatalf("expected task name %s, got %s", taskName, task.Name)
	}
	err = database.DeleteTask(task.Id)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestUpdateTaskStatus(t *testing.T) {
	utils.SetupTestDatabase(t)
	defer utils.TeardownTestDatabase(t)

	taskName := "TestTask"
	idTask, err := database.CreateTask(taskName)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = database.UpdateTaskStatus(idTask, true)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	task, err := database.GetTaskByName(taskName)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !task.Done {
		t.Fatal("expected task to be marked as done")
	}
	err = database.DeleteTask(idTask)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestDeleteTask(t *testing.T) {
	utils.SetupTestDatabase(t)
	defer utils.TeardownTestDatabase(t)

	taskName := "TestTask"
	idTask, err := database.CreateTask(taskName)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = database.DeleteTask(idTask)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = database.GetTaskByName(taskName)
	if err == nil {
		t.Fatal("expected error for non-existent task, got nil")
	}
}

func TestGetAllTasks(t *testing.T) {
	utils.SetupTestDatabase(t)
	defer utils.TeardownTestDatabase(t)

	taskNames := []string{"Task1", "Task2", "Task3"}
	for _, name := range taskNames {
		_, err := database.CreateTask(name)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	}

	tasks, err := database.GetAllTasks()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(tasks) < len(taskNames) {
		t.Fatalf("expected at least %d tasks, got %d", len(taskNames), len(tasks))
	}

	for _, task := range tasks {
		err = database.DeleteTask(task.Id)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	}
}
