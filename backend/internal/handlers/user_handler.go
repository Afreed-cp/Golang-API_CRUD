package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"goapi/internal/models"
	"goapi/internal/services"

	"github.com/gorilla/mux"
)

// UserHandler handles HTTP requests for user operations
type UserHandler struct {
	userService *services.UserService
}

// NewUserHandler creates a new user handler
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetUsers handles GET /api/users
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		models.WriteInternalServerError(w, "Failed to retrieve users")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    users,
	})
}

// GetUser handles GET /api/users/{id}
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		models.WriteValidationError(w, "Invalid user ID")
		return
	}

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		models.WriteNotFoundError(w, "User")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    user,
	})
}

// CreateUser handles POST /api/users
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		models.WriteValidationError(w, "Invalid JSON payload")
		return
	}

	// Basic validation
	if req.Name == "" || req.Email == "" {
		models.WriteValidationError(w, "Name and email are required")
		return
	}

	user, err := h.userService.CreateUser(req)
	if err != nil {
		if err.Error() == "email already exists" {
			models.WriteConflictError(w, "Email already exists")
			return
		}
		models.WriteInternalServerError(w, "Failed to create user")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    user,
	})
}

// UpdateUser handles PUT /api/users/{id}
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		models.WriteValidationError(w, "Invalid user ID")
		return
	}

	var req models.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		models.WriteValidationError(w, "Invalid JSON payload")
		return
	}

	// Basic validation
	if req.Name == "" || req.Email == "" {
		models.WriteValidationError(w, "Name and email are required")
		return
	}

	user, err := h.userService.UpdateUser(id, req)
	if err != nil {
		if err.Error() == "user not found" {
			models.WriteNotFoundError(w, "User")
			return
		}
		if err.Error() == "email already exists" {
			models.WriteConflictError(w, "Email already exists")
			return
		}
		models.WriteInternalServerError(w, "Failed to update user")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    user,
	})
}

// DeleteUser handles DELETE /api/users/{id}
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		models.WriteValidationError(w, "Invalid user ID")
		return
	}

	err = h.userService.DeleteUser(id)
	if err != nil {
		if err.Error() == "user not found" {
			models.WriteNotFoundError(w, "User")
			return
		}
		models.WriteInternalServerError(w, "Failed to delete user")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
