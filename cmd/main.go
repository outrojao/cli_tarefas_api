package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"cli_tasks_api/internal/database"
	"cli_tasks_api/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("configs/.env"); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	initializeStatus := make(chan bool, 2)
	go database.InitDatabase(initializeStatus)
	defer func() {
		if err := database.CloseDatabase(); err != nil {
			log.Println("error closing database:", err)
		}
	}()
	go server.InitApi(initializeStatus)

	dbOK := <-initializeStatus
	apiOK := <-initializeStatus
	if !dbOK || !apiOK {
		log.Fatal("Failed to initialize API components.")
	}

	fmt.Println("API is running")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	<-shutdown
	log.Println("Shutdown signal received, exiting")
}
