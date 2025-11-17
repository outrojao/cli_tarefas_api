package handlers

import (
	"cli_tasks_api/internal/database"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
)

func setupTestDatabase(t *testing.T) {
	if err := godotenv.Load("../../configs/.env"); err != nil {
		log.Fatal("Error loading .env file:", err)
		os.Exit(1)
	}

	status := make(chan bool)
	go database.InitDatabase(status)
	if success := <-status; !success {
		t.Fatal("failed to initialize database")
	}
}

func teardownTestDatabase(t *testing.T) {
	if err := database.CloseDatabase(); err != nil {
		t.Fatalf("failed to close database: %v", err)
	}
}

func TestCreateTaskApi(t *testing.T) {
	setupTestDatabase(t)
	defer teardownTestDatabase(t)
	payload := `{"task_name":"TestTask"}`
	req := httptest.NewRequest(http.MethodPost, "/create", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	CreateTask(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d, body: %s", http.StatusCreated, rr.Code, rr.Body.String())
	}
}

func TestDoTaskApi(t *testing.T) {
	setupTestDatabase(t)
	defer teardownTestDatabase(t)
	taskName := "TestTask"

	taskNameEscaped := url.PathEscape(taskName)
	url := "/do/" + taskNameEscaped
	req := httptest.NewRequest(http.MethodPut, url, nil)
	rr := httptest.NewRecorder()

	DoTask(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d, body: %s", http.StatusOK, rr.Code, rr.Body.String())
	}
}

func TestRemoveTaskApi(t *testing.T) {
	setupTestDatabase(t)
	defer teardownTestDatabase(t)
	taskName := "TestTask"

	taskNameEscaped := url.PathEscape(taskName)
	url := "/remove/" + taskNameEscaped
	req := httptest.NewRequest(http.MethodDelete, url, nil)
	rr := httptest.NewRecorder()

	RemoveTask(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d, body: %s", http.StatusOK, rr.Code, rr.Body.String())
	}
}

func TestListTasksApi(t *testing.T) {
	setupTestDatabase(t)
	defer teardownTestDatabase(t)
	req := httptest.NewRequest(http.MethodGet, "/list", nil)
	rr := httptest.NewRecorder()

	ListTasks(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d, body: %s", http.StatusOK, rr.Code, rr.Body.String())
	}
}

func TestHealthCheckApi(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()

	HealthCheck(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d, body: %s", http.StatusOK, rr.Code, rr.Body.String())
	}
}
