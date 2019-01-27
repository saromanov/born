// Package server contains RESTAPI endpoints for Born
package server

import (
	"log"
	"net/http"

	"github.com/qor/auth"
	"github.com/qor/auth/providers/password"
	"github.com/qor/session/manager"
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
	mux := http.NewServeMux()
	authData := auth.New(&auth.Config{})
	authData.RegisterProvider(password.New(&password.Config{}))

	mux.Handle("/auth/", authData.NewServeMux())
	http.HandleFunc("/health", Health)
	log.Printf("starting server at %s", addr)
	log.Fatal(http.ListenAndServe(addr, manager.SessionManager.Middleware(mux)))
}
