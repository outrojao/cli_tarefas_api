package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDatabase(status chan<- bool) {
	db, err := ConectDatabase()
	if err != nil {
		status <- false
		log.Println("Error connecting to database:", err)
		return
	}
	DB = db
	if err := DB.Ping(); err != nil {
		status <- false
		log.Println("Error pinging database:", err)
		return
	}
	status <- true
}

func ConectDatabase() (*sql.DB, error) {
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	sslmode := os.Getenv("DB_SSLMODE")

	connection := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", user, dbname, password, host, sslmode)
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CloseDatabase() error {
	if DB == nil {
		return nil
	}
	if err := DB.Close(); err != nil {
		log.Println("Error closing database:", err)
		return err
	}
	DB = nil
	return nil
}
