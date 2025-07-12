package http

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func StartHTTPServer() {

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodOptions, http.MethodDelete},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}))
	
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/users", func(r chi.Router) {
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {})
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {})
		r.Get("/{userID}", func(w http.ResponseWriter, r *http.Request) {})
		r.Put("/{userID}", func(w http.ResponseWriter, r *http.Request) {})
		r.Delete("/{userID}", func(w http.ResponseWriter, r *http.Request) {})
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("error while init http server on port %s", ":8080")
	}
}
