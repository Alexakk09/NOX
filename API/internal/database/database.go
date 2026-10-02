package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func Connect() {
	// Load .env
	if err := godotenv.Load(); err != nil {
	log.Println(".env file not found, using system environment variables")
}

	// Build connection string
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	// Create connection pool
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	// Test connection
	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal("Database ping failed:", err)
	}

	DB = pool

	log.Println("✅ Connected to PostgreSQL!")
}