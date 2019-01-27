package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	structs "github.com/saromanov/born/structs/v1"
)

// createBuild provides creating of the new build
func createBuild(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-BORN-TOKEN")
	decoder := json.NewDecoder(r.Body)
	var payload *structs.BuildRequest
	err := decoder.Decode(payload)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to unmarshal input: %v", err), http.StatusBadRequest)
		return
	}
	if payload.Repo == "" {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "repo is not defined", http.StatusBadRequest)
		return
	}
	payload.UserID = token
	w.WriteHeader(http.StatusOK)
}
