package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"goapi/internal/config"
	"goapi/internal/database"
	"goapi/internal/handlers"
	"goapi/internal/middleware"
	"goapi/internal/services"
	"goapi/pkg/logger"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize logger
	logger := logger.NewLogger()
	logger.Info("Starting Go API Server...")

	// Load configuration
	cfg := config.LoadConfig()
	logger.Info("Configuration loaded successfully")

	// Initialize database
	db, err := database.NewDatabase(cfg)
	if err != nil {
		logger.Error("Failed to initialize database: %v", err)
		os.Exit(1)
	}
	defer db.Close()

	// Initialize repositories
	userRepo := database.NewUserRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepo)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	testHandler := handlers.NewTestHandler()

	// Setup routes
	router := setupRoutes(userHandler,testHandler)

	// Setup middleware
	handler := setupMiddleware(router, cfg)

	// Create HTTP server
	server := &http.Server{
		Addr:         cfg.Server.Host + ":" + cfg.Server.Port,
		Handler:      handler,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
	}

	// Start server in a goroutine
	go func() {
		logger.Info("Server starting on %s:%s", cfg.Server.Host, cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Failed to start server: %v", err)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Server is shutting down...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown: %v", err)
		os.Exit(1)
	}

	logger.Info("Server exited")
}

// setupRoutes configures all API routes
func setupRoutes(userHandler *handlers.UserHandler, testHandler *handlers.TestHandler) *mux.Router {
	router := mux.NewRouter()

	// API routes
	api := router.PathPrefix("/api").Subrouter()
	
	// User routes
	api.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	api.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	api.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	api.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	api.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")



	api.HandleFunc("/test", testHandler.Test).Methods("GET")

	// Health check endpoint
	router.HandleFunc("/health", healthCheck).Methods("GET")

	return router
}

// setupMiddleware configures all middleware
func setupMiddleware(router *mux.Router, cfg *config.Config) http.Handler {
	// Recovery middleware (should be first)
	handler := middleware.RecoveryMiddleware(router)
	
	// Logging middleware
	handler = middleware.LoggingMiddleware(handler)
	
	// CORS middleware
	handler = middleware.DefaultCORS()(handler)

	return handler
}

// healthCheck handles health check requests
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status":"healthy","timestamp":"%s"}`, time.Now().Format(time.RFC3339))
}
