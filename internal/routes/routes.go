package routes

import (
	"cli_tasks_api/internal/handlers"
	"cli_tasks_api/internal/middleware"
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/create", middleware.AuthMiddleware(handlers.CreateTask))
	http.HandleFunc("/do/", middleware.AuthMiddleware(handlers.DoTask))
	http.HandleFunc("/remove/", middleware.AuthMiddleware(handlers.RemoveTask))
	http.HandleFunc("/list", middleware.AuthMiddleware(handlers.ListTasks))
	http.HandleFunc("/health", handlers.HealthCheck)
}
