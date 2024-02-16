package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func setContentType(w http.ResponseWriter, code int, contentType string) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(code)
}

func writeJSON(w http.ResponseWriter, code int, i interface{}) error {
	setContentType(w, code, "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(i); err != nil {
		return fmt.Errorf("http.writeJSON error: %w", err)
	}
	return nil
}
