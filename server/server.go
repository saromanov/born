// Package server contains RESTAPI endpoints for Born
package server

import "net/http"

// Health provides healthcheck of the server
func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// Setup Provides setup of the server
func Setup(addr string) {

}
