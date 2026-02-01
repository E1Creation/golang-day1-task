package config

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func InitDB(connectionString string) (*sql.DB, error) {
	log.Printf("Attempting to connect to database...")

	// Open database
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Printf("Failed to open database: %v", err)
		return nil, err
	}

	log.Printf("Database opened, testing connection...")

	// Test connection
	err = db.Ping()
	if err != nil {
		log.Printf("Failed to ping database: %v", err)
		return nil, err
	}

	// Set connection pool settings for Supabase
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(1 * time.Minute)

	log.Println("Database connected successfully")
	return db, nil
}
