package http

import (
	"log"
	"net/http"
)

func StartHTTPServer() {
	handler := http.ServeMux{}

	handler.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World !"))
	})

	log.Default().Printf("http server started at %s", ":8080")

	if err := http.ListenAndServe(":8080", &handler); err != nil {
		log.Fatalf("error while init http server on port %s", ":8080")
	}
}
