package utils

import (
	"encoding/json"
	"net/http"
)

// JSONResponse writes a JSON response with the given status and data.
// If status is 204 No Content, it writes the header only.
func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if status == http.StatusNoContent {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// ErrorResponse writes a standardized JSON error response with the given status and message.
func ErrorResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(struct {
		Error string `json:"error"`
	}{Error: message})
}
