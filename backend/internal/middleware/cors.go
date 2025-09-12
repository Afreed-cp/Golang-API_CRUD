package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

// CORSConfig holds CORS configuration
type CORSConfig struct {
	AllowedOrigins []string
	AllowedMethods []string
	AllowedHeaders []string
}

// NewCORS creates a new CORS middleware
func NewCORS(config CORSConfig) func(http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   config.AllowedOrigins,
		AllowedMethods:   config.AllowedMethods,
		AllowedHeaders:   config.AllowedHeaders,
		AllowCredentials: false,
		Debug:            false,
	})

	return c.Handler
}

// DefaultCORS returns a default CORS configuration
func DefaultCORS() func(http.Handler) http.Handler {
	return NewCORS(CORSConfig{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})
}
