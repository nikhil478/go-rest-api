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
	"github.com/spf13/viper"
)

func StartHTTPServer() {

	viper.AddConfigPath(".")
	viper.SetConfigName("secrets")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

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
		r.Get("/id", handlers.GetUserByID)
		r.Put("/id", handlers.UpdateUser)
		r.Delete("/id", handlers.DeleteUser)
	})

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	port := viper.GetString("app.port")

	go func() {
		log.Default().Printf("Server started on port %s", port)
		if err := http.ListenAndServe(port, r); err != nil {
			log.Fatalf("error while init http server on port %s", port)
		}
	}()

	time.Sleep(5 * time.Second)
	<-quit
}
