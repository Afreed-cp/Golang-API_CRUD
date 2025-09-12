package handlers

import (
	"encoding/json"
	"net/http"
)

// TestHandler handles HTTP requests for user operations
type TestHandler struct {

}

// NewUserHandler creates a new user handler
func NewTestHandler() *TestHandler {
	return &TestHandler{
	}
}


// Test handles GET /api/users
func (h *TestHandler) Test(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    "Test",
	})
}
