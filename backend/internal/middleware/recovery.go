package middleware

import (
	"log"
	"net/http"

	"goapi/internal/models"
)

// RecoveryMiddleware recovers from panics and returns a 500 error
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				models.WriteInternalServerError(w, "Internal server error")
			}
		}()
		
		next.ServeHTTP(w, r)
	})
}
