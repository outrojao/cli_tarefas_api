package handlers

import (
	"cli_tasks_api/internal/database"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}

	var payload struct {
		TaskName string `json:"task_name"`
	}
	if err := json.Unmarshal(body, &payload); err != nil || payload.TaskName == "" {
		http.Error(w, "invalid task data", http.StatusBadRequest)
		return
	}

	idTask, err := database.CreateTask(payload.TaskName)
	if err != nil {
		http.Error(w, "failed to create task", http.StatusInternalServerError)
		return
	}
	_ = idTask

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func DoTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	taskName := strings.TrimPrefix(r.URL.Path, "/do/")
	if taskName == "" || taskName == r.URL.Path {
		http.Error(w, "Missing task_name parameter", http.StatusBadRequest)
		return
	}
	if unescaped, err := url.PathUnescape(taskName); err == nil {
		taskName = unescaped
	}

	t, err := database.GetTaskByName(taskName)
	if err != nil {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	if t.Done {
		http.Error(w, "task already done", http.StatusBadRequest)
		return
	}

	if err := database.UpdateTaskStatus(t.Id, true); err != nil {
		http.Error(w, "failed to update task status", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func RemoveTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	taskName := strings.TrimPrefix(r.URL.Path, "/remove/")
	if taskName == "" || taskName == r.URL.Path {
		http.Error(w, "Missing task_name parameter", http.StatusBadRequest)
		return
	}
	if unescaped, err := url.PathUnescape(taskName); err == nil {
		taskName = unescaped
	}

	t, err := database.GetTaskByName(taskName)
	if err != nil {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	if err := database.DeleteTask(t.Id); err != nil {
		http.Error(w, "failed to delete task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func ListTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tasks, err := database.GetAllTasks()
	if err != nil {
		http.Error(w, "failed to fetch tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
}
