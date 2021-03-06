// Package server contains RESTAPI endpoints for Born
package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/qor/auth"
	"github.com/qor/auth/providers/password"
	"github.com/qor/session/manager"
	"github.com/saromanov/born/provider"
	"github.com/saromanov/born/store"
)

const (
	defaultAddress = "127.0.0.1:2310"
	contentType    = "application/json"
)

type handler struct {
	handler func(provider.Provider, store.Store, http.ResponseWriter, *http.Request)
	P       provider.Provider
	S       store.Store
}

func (handle *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", contentType)
	handle.handler(handle.P, handle.S, w, r)
}

// Health provides healthcheck of the server
func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// Setup Provides setup of the server
func Setup(p provider.Provider, s store.Store, addr string) {
	if addr == "" {
		addr = defaultAddress
	}
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/v1/builds", &handler{P: p, S: s, handler: createBuild}).Methods("POST")
	authData := auth.New(&auth.Config{})
	authData.RegisterProvider(password.New(&password.Config{}))

	router.Handle("v1//auth/", authData.NewServeMux())
	router.HandleFunc("v1/health", Health)
	log.Printf("starting server at %s", addr)
	log.Fatal(http.ListenAndServe(addr, manager.SessionManager.Middleware(router)))
}
