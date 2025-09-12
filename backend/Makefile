# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=goapi
BINARY_UNIX=$(BINARY_NAME)_unix

# Build the application
build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/api

# Build for Linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v ./cmd/api

# Clean build artifacts
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# Run the application
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/api
	./$(BINARY_NAME)

# Run tests
test:
	$(GOTEST) -v ./...

# Run tests with coverage
test-coverage:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out

# Download dependencies
deps:
	$(GOMOD) download
	$(GOMOD) tidy

# Format code
fmt:
	$(GOCMD) fmt ./...

# Lint code
lint:
	golangci-lint run

# Docker commands
docker-build:
	docker build -t $(BINARY_NAME) .

docker-run:
	docker run -p 8080:8080 $(BINARY_NAME)

docker-compose-up:
	docker-compose up --build

docker-compose-down:
	docker-compose down

docker-compose-logs:
	docker-compose logs -f

# Development commands
dev:
	air

# Install development tools
install-tools:
	$(GOGET) -u github.com/cosmtrek/air
	$(GOGET) -u github.com/golangci/golangci-lint/cmd/golangci-lint

# Database commands
db-migrate:
	# Add migration commands here if using a migration tool

# Help
help:
	@echo "Available commands:"
	@echo "  build          - Build the application"
	@echo "  build-linux    - Build for Linux"
	@echo "  clean          - Clean build artifacts"
	@echo "  run            - Run the application"
	@echo "  test           - Run tests"
	@echo "  test-coverage  - Run tests with coverage"
	@echo "  deps           - Download dependencies"
	@echo "  fmt            - Format code"
	@echo "  lint           - Lint code"
	@echo "  docker-build   - Build Docker image"
	@echo "  docker-run     - Run Docker container"
	@echo "  docker-compose-up   - Start with docker-compose"
	@echo "  docker-compose-down - Stop docker-compose"
	@echo "  docker-compose-logs - Show docker-compose logs"
	@echo "  dev            - Run in development mode with air"
	@echo "  install-tools  - Install development tools"
	@echo "  help           - Show this help message"

.PHONY: build build-linux clean run test test-coverage deps fmt lint docker-build docker-run docker-compose-up docker-compose-down docker-compose-logs dev install-tools help
