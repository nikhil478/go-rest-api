package http

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/nikhil478/go-rest-api/internal/http/handlers"
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
		r.Post("/", handlers.CreateUser)
		r.Get("/", handlers.GetAllUser)
		r.Get("/{userID}", handlers.GetUserByID)
		r.Put("/{userID}", handlers.UpdateUser)
		r.Delete("/{userID}", handlers.DeleteUser)
	})

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Default().Printf("Server started on port %s", ":8080")
		if err := http.ListenAndServe(":8080", r); err != nil {
			log.Fatalf("error while init http server on port %s", ":8080")
		}
	}()

	time.Sleep(5 * time.Second)
	<-quit
}
