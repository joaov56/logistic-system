package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"logistic-system/internal/delivery/application"
	"logistic-system/internal/delivery/infrastructure"
	"logistic-system/internal/delivery/interfaces"
	"logistic-system/pkg/logger"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// Initialize logger
	logger := logger.New()

	// Get database configuration from environment variables
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "logistics")

	// Connect to database
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("Failed to connect to database: %v", err)
		os.Exit(1)
	}
	defer db.Close()

	// Test database connection
	if err := db.Ping(); err != nil {
		logger.Error("Failed to ping database: %v", err)
		os.Exit(1)
	}

	// Create database schema
	if err := createSchema(db); err != nil {
		logger.Error("Failed to create schema: %v", err)
		os.Exit(1)
	}

	// Initialize repository
	repo := infrastructure.NewPostgresRepository(db)

	// Initialize service
	service := application.NewService(repo)

	// Initialize HTTP handler
	handler := interfaces.NewHandler(service)

	// Create router
	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	// Start server
	port := getEnv("PORT", "8080")
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Info("Server starting on port %s", port)
	if err := server.ListenAndServe(); err != nil {
		logger.Error("Server failed to start: %v", err)
		os.Exit(1)
	}
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// createSchema creates the database schema
func createSchema(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS deliveries (
			id VARCHAR(36) PRIMARY KEY,
			order_id VARCHAR(36) NOT NULL,
			customer_id VARCHAR(36) NOT NULL,
			address TEXT NOT NULL,
			status VARCHAR(20) NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL,
			delivered_at TIMESTAMP
		);
	`
	_, err := db.Exec(query)
	return err
} 