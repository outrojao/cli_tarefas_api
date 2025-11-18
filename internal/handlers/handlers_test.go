package handlers

import (
	"cli_tasks_api/internal/utils"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestCreateTaskApi(t *testing.T) {
	utils.SetupTestDatabase(t)
	defer utils.TeardownTestDatabase(t)
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
	utils.SetupTestDatabase(t)
	defer utils.TeardownTestDatabase(t)
	taskName := "TestTask"

	taskNameEscaped := url.PathEscape(taskName)
	url := "/do/" + taskNameEscaped
	req := httptest.NewRequest(http.MethodPut, url, nil)
	rr := httptest.NewRecorder()

	DoTask(rr, req)

	if rr.Code != http.StatusOK && rr.Code != http.StatusNotModified {
		t.Fatalf("expected status %d or %d, got %d, body: %s", http.StatusOK, http.StatusNotModified, rr.Code, rr.Body.String())
	}
}

func TestRemoveTaskApi(t *testing.T) {
	utils.SetupTestDatabase(t)
	defer utils.TeardownTestDatabase(t)
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
	utils.SetupTestDatabase(t)
	defer utils.TeardownTestDatabase(t)
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
