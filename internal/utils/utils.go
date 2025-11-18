package utils

import (
	"cli_tasks_api/internal/database"
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func SetupTestDatabase(t *testing.T) {
	if err := godotenv.Load("../../configs/database.env"); err != nil {
		log.Println("Error loading .env file:", err)
		return
	}

	status := make(chan bool)
	go database.InitDatabase(status)
	if success := <-status; !success {
		t.Fatal("failed to initialize database")
	}
}

func TeardownTestDatabase(t *testing.T) {
	if err := database.CloseDatabase(); err != nil {
		t.Fatalf("failed to close database: %v", err)
	}
}
