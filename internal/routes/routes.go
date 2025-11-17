package routes

import (
	"cli_tasks_api/internal/handlers"
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/create", handlers.CreateTask)
	http.HandleFunc("/do/", handlers.DoTask)
	http.HandleFunc("/remove/", handlers.RemoveTask)
	http.HandleFunc("/list", handlers.ListTasks)
	http.HandleFunc("/health", handlers.HealthCheck)
}
