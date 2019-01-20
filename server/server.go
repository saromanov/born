// Package server contains RESTAPI endpoints for Born
package server

import (
	"log"
	"net/http"
)

const defaultAddress = "127.0.0.1:2310"

// Health provides healthcheck of the server
func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// Setup Provides setup of the server
func Setup(addr string) {
	if addr == "" {
		addr = defaultAddress
	}
	http.HandleFunc("/health", Health)
	log.Fatal(http.ListenAndServe(addr, nil))
}
