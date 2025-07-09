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
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch,
			http.MethodOptions, http.MethodDelete},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/admin", func(r chi.Router) {
		r.Route("/supplier", func(r chi.Router) {
			r.Post("/", func(w http.ResponseWriter, r *http.Request) {})
			r.Get("/{supplierID}", func(w http.ResponseWriter, r *http.Request) {})
		})
	})

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World !"))
	})

	log.Default().Printf("http server started at %s", ":8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("error while init http server on port %s", ":8080")
	}
}
