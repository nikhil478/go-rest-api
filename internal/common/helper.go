package common

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SendResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
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
