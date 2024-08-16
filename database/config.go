package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

// DB is a global variable to hold the database connection
var DB *pgx.Conn

// Connect initializes the database connection
func Connect() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Connect to PostgreSQL
	connectionString := os.Getenv("DATABASE_URL")
	var err error
	DB, err = pgx.Connect(context.Background(), connectionString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
}
