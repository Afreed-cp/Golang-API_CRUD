#!/bin/bash

# Go API Backend Startup Script

echo "Starting Go API Backend..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed or not in PATH"
    exit 1
fi

# Check if PostgreSQL is running
if ! pg_isready -q; then
    echo "Error: PostgreSQL is not running"
    echo "Please start PostgreSQL before running the backend"
    echo "Or use: docker-compose up -d postgres"
    exit 1
fi

# Set environment variables
export SERVER_HOST=0.0.0.0
export SERVER_PORT=8080
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=pass
export DB_NAME=postgres
export DB_SSLMODE=disable
export LOG_LEVEL=info

# Install dependencies
echo "Installing Go dependencies..."
go mod tidy

# Run the Go application
echo "Starting server on port 8080..."
echo "API will be available at: http://localhost:8080"
echo "Health check: http://localhost:8080/health"
echo "Press Ctrl+C to stop the server"
echo ""

go run cmd/api/main.go
