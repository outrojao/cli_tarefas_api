package server

import (
	"cli_tasks_api/internal/routes"
	"log"
	"net"
	"net/http"
)

func InitApi(status chan<- bool) {
	routes.InitRoutes()

	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Printf("Failed to start API server: %v\n", err)
		status <- false
		return
	}

	status <- true

	if err := http.Serve(ln, nil); err != nil {
		log.Printf("Failed to start HTTP server: %v\n", err)
	}
}
