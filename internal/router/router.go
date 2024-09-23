package router

import (
	"github.com/go-chi/chi"
	"github.com/mongj/gds-onecv-swe-assignment/internal/api"
	"github.com/mongj/gds-onecv-swe-assignment/internal/handlers/applicants"
	"github.com/mongj/gds-onecv-swe-assignment/internal/handlers/applications"
	"github.com/mongj/gds-onecv-swe-assignment/internal/handlers/schemes"
	"github.com/mongj/gds-onecv-swe-assignment/internal/middleware"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) chi.Router {
	r := chi.NewRouter()
	middleware.Setup(r, db)
	setupRoutes(r)
	return r
}

func setupRoutes(r chi.Router) {
	r.Route("/api", getAPIRoutes())
}

func getAPIRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/applicants", api.HTTPHandler(applicants.HandleList))
		r.Post("/applicants", api.HTTPHandler(applicants.HandleCreate))
		r.Get("/applications", api.HTTPHandler(applications.HandleList))
		r.Post("/applications", api.HTTPHandler(applications.HandleCreate))
		r.Route("/schemes", func(r chi.Router) {
			r.Get("/", api.HTTPHandler(schemes.HandleList))
			r.Get("/eligible", api.HTTPHandler(schemes.HandleFind))
		})
	}
}
