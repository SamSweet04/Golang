package utils


import (
	"encoding/json"
	"net/http"
)

// respondWithError writes a JSON error response to the HTTP response writer.
func respondWithError(w http.ResponseWriter, code int, message string) {
	resp := map[string]string{"error": message}
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

// respondWithSuccess writes a JSON success response to the HTTP response writer.
func respondWithSuccess(w http.ResponseWriter, data interface{}) {
	resp := map[string]interface{}{"data": data}
	json.NewEncoder(w).Encode(resp)
}