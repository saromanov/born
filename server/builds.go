package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/saromanov/born/internal/build"
	"github.com/saromanov/born/provider"
	"github.com/saromanov/born/store"
	structs "github.com/saromanov/born/structs/v1"
)

// createBuild provides creating of the new build
func createBuild(p provider.Provider, s store.Store, w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-BORN-TOKEN")
	decoder := json.NewDecoder(r.Body)
	payload := &structs.BuildRequest{}
	err := decoder.Decode(payload)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to unmarshal input: %v", err), http.StatusBadRequest)
		return
	}
	if payload.Repo == "" {
		http.Error(w, "repo is not defined", http.StatusBadRequest)
		return
	}
	b := &build.Build{
		Repo: payload.Repo,
		P:    p,
	}
	err = b.Create()
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to create build: %v", err), http.StatusBadRequest)
		return
	}
	payload.UserID = token
	w.WriteHeader(http.StatusOK)
}
