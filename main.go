package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"logistic-system/internal/delivery/application"
	"logistic-system/internal/delivery/domain"
	"logistic-system/internal/delivery/infrastructure"
	"logistic-system/internal/delivery/interfaces"
	"logistic-system/pkg/logger"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	logger := logger.New()

	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "logistics")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Failed to connect to database: %v", err)
		os.Exit(1)
	}

	if err := db.AutoMigrate(&domain.Delivery{}); err != nil {
		logger.Error("Failed to migrate database: %v", err)
		os.Exit(1)
	}

	repo := infrastructure.NewPostgresRepository(db)

	service := application.NewService(repo)

	handler := interfaces.NewHandler(service)

	router := mux.NewRouter()
	handler.RegisterRoutes(router)

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

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
} 