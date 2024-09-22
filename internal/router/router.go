package router

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mongj/gds-onecv-swe-assignment/internal/handlers/applicants"
	"github.com/mongj/gds-onecv-swe-assignment/internal/handlers/schemes"
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
	r.Route("/api", getAPIRoutes())
}

func getAPIRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/applicants", applicants.HandleList)
		r.Post("/applicants", applicants.HandleCreate)
		r.Get("/schemes", schemes.HandleList)
		r.Get("/schemes/eligible", schemes.HandleFind)
		r.Get("/applications", applicants.HandleList)
		r.Post("/applications", applicants.HandleCreate)
	}
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
