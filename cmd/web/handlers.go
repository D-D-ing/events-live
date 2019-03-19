package main

import (
	"io"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	go w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}
