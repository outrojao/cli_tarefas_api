package utils

import (
	"cli_tasks_api/internal/database"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func SetupTestDatabase(t *testing.T) {
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

func TeardownTestDatabase(t *testing.T) {
	if err := database.CloseDatabase(); err != nil {
		t.Fatalf("failed to close database: %v", err)
	}
}
