package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/gorm"
)

type contextKey string

func Setup(db *gorm.DB) chi.Router {
	r := chi.NewRouter()
	setUpMiddleware(r, db)
	setupRoutes(r)
	return r
}

func setupRoutes(r chi.Router) {
	// Add routes here
	r.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
}

func setUpMiddleware(r chi.Router, db *gorm.DB) {
	// Injects a request ID in the context of each request
	r.Use(middleware.RequestID)
	// Sets a http.Request's RemoteAddr to that of either the X-Forwarded-For or X-Real-IP header
	r.Use(middleware.RealIP)
	// Recovers from panics and return a 500 Internal Service Error
	r.Use(middleware.Recoverer)
	// Returns a 504 Gateway Timeout after 10 seconds if the request hasn't completed
	r.Use(middleware.Timeout(10 * time.Second))
	// CORS
	r.Use(corsMiddleware())
	// Add database middleware which injects the db instance into the request context
	r.Use(databaseMiddleware(db))
}
