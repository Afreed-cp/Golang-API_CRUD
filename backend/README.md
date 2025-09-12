# Go API Project

A production-ready REST API built with Go, featuring clean architecture, proper error handling, and comprehensive middleware.

## 🏗️ Architecture

This project follows the **Clean Architecture** pattern with the following structure:

```
backend/
├── cmd/
│   └── api/                 # Application entry point
│       └── main.go
├── internal/                # Private application code
│   ├── config/             # Configuration management
│   ├── database/           # Database layer
│   │   ├── database.go     # Database connection
│   │   └── user_repository.go
│   ├── handlers/           # HTTP handlers
│   │   └── user_handler.go
│   ├── middleware/         # HTTP middleware
│   │   ├── cors.go
│   │   ├── logging.go
│   │   └── recovery.go
│   ├── models/             # Data models
│   │   ├── user.go
│   │   └── errors.go
│   └── services/           # Business logic
│       └── user_service.go
├── pkg/                    # Public packages
│   ├── logger/             # Logging utilities
│   └── utils/              # Utility functions
├── configs/                # Configuration files
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── README.md
```

## 🚀 Features

- **Clean Architecture**: Separation of concerns with handlers, services, and repositories
- **Configuration Management**: Environment-based configuration
- **Database Integration**: PostgreSQL with connection pooling
- **Middleware**: CORS, logging, and recovery middleware
- **Error Handling**: Standardized error responses
- **Docker Support**: Containerized application with Docker Compose
- **Health Checks**: Built-in health check endpoint
- **Graceful Shutdown**: Proper server shutdown handling
- **Validation**: Input validation and sanitization

## 📋 Prerequisites

- Go 1.21 or higher
- PostgreSQL 12 or higher
- Docker and Docker Compose (optional)

## 🛠️ Installation

### Option 1: Local Development

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd GoAPIProject/backend
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   cp env.example .env
   # Edit .env with your configuration
   ```

4. **Set up PostgreSQL database**
   ```bash
   # Create database
   createdb postgres
   # Or use your preferred method
   ```

5. **Run the application**
   ```bash
   make run
   # Or
   go run cmd/api/main.go
   ```

### Option 2: Docker Compose

1. **Start the application with Docker Compose**
   ```bash
   docker-compose up --build
   ```

This will start both the API and PostgreSQL database.

## 🔧 Configuration

The application can be configured using environment variables:

| Variable | Default | Description |
|----------|---------|-------------|
| `SERVER_HOST` | `0.0.0.0` | Server host |
| `SERVER_PORT` | `8080` | Server port |
| `SERVER_READ_TIMEOUT` | `30` | Read timeout in seconds |
| `SERVER_WRITE_TIMEOUT` | `30` | Write timeout in seconds |
| `DB_HOST` | `localhost` | Database host |
| `DB_PORT` | `5432` | Database port |
| `DB_USER` | `postgres` | Database user |
| `DB_PASSWORD` | `pass` | Database password |
| `DB_NAME` | `postgres` | Database name |
| `DB_SSLMODE` | `disable` | SSL mode |
| `LOG_LEVEL` | `info` | Log level |
| `LOG_FORMAT` | `json` | Log format |

## 📚 API Endpoints

### Users

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/api/users` | Get all users |
| `GET` | `/api/users/{id}` | Get user by ID |
| `POST` | `/api/users` | Create new user |
| `PUT` | `/api/users/{id}` | Update user |
| `DELETE` | `/api/users/{id}` | Delete user |

### Health Check

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/health` | Health check |

## 📝 API Examples

### Create User
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com"
  }'
```

### Get All Users
```bash
curl http://localhost:8080/api/users
```

### Update User
```bash
curl -X PUT http://localhost:8080/api/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Doe",
    "email": "jane@example.com"
  }'
```

### Delete User
```bash
curl -X DELETE http://localhost:8080/api/users/1
```

## 🧪 Testing

```bash
# Run tests
make test

# Run tests with coverage
make test-coverage
```

## 🐳 Docker

### Build Docker Image
```bash
make docker-build
```

### Run with Docker Compose
```bash
make docker-compose-up
```

### View Logs
```bash
make docker-compose-logs
```

## 🔧 Development

### Hot Reload
```bash
# Install air for hot reload
make install-tools

# Run with hot reload
make dev
```

### Code Formatting
```bash
make fmt
```

### Linting
```bash
make lint
```

## 📊 Monitoring

The application includes:
- Health check endpoint at `/health`
- Request logging middleware
- Graceful shutdown handling
- Database connection pooling

## 🚀 Production Deployment

1. **Set production environment variables**
2. **Use a production database**
3. **Enable SSL/TLS**
4. **Set up monitoring and logging**
5. **Use a reverse proxy (nginx)**
6. **Set up CI/CD pipeline**

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License.

## 🆘 Support

For support, please open an issue in the repository.