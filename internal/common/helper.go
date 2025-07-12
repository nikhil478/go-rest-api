package common

import (
	"fmt"
	"net/http"
)


func SendPlainResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	_, err := fmt.Fprint(w, data)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}


func SendErrorResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusConflict)

	_, err := fmt.Fprint(w, data)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}