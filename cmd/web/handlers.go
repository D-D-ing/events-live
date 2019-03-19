package main

import (
	"io"
	"net/http"
)

// HealthCheckHandler returns a health check
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	go w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}
