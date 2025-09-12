package models

import (
	"encoding/json"
	"net/http"
)

// APIError represents an API error response
type APIError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// ErrorResponse represents a standardized error response
type ErrorResponse struct {
	Success bool     `json:"success"`
	Error   APIError `json:"error"`
}

// NewErrorResponse creates a new error response
func NewErrorResponse(message string, code int) ErrorResponse {
	return ErrorResponse{
		Success: false,
		Error: APIError{
			Error:   http.StatusText(code),
			Message: message,
			Code:    code,
		},
	}
}

// WriteError writes an error response to the HTTP response writer
func WriteError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	
	errorResponse := NewErrorResponse(message, code)
	json.NewEncoder(w).Encode(errorResponse)
}

// WriteValidationError writes a validation error response
func WriteValidationError(w http.ResponseWriter, message string) {
	WriteError(w, message, http.StatusBadRequest)
}

// WriteNotFoundError writes a not found error response
func WriteNotFoundError(w http.ResponseWriter, resource string) {
	WriteError(w, resource+" not found", http.StatusNotFound)
}

// WriteInternalServerError writes an internal server error response
func WriteInternalServerError(w http.ResponseWriter, message string) {
	WriteError(w, message, http.StatusInternalServerError)
}

// WriteConflictError writes a conflict error response
func WriteConflictError(w http.ResponseWriter, message string) {
	WriteError(w, message, http.StatusConflict)
}
