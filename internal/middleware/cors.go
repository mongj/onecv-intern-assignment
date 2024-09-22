package middleware

import (
	"net/http"

	"github.com/go-chi/cors"
)

func corsMiddleware() func(http.Handler) http.Handler {
	options := cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins for now since we don't have a frontend
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}

	return cors.Handler(options)
}
