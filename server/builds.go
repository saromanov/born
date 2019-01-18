package server

import (
	"fmt"
	"net/http"
	"encoding/json"

	structs "github.com/saromanov/drone/structs/v1"
)

// createBuild provides creating of the new build
func createBuild(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-BORN-TOKEN")
	decoder := json.NewDecoder(r.Body)
	var payload *structs.BuildRequest
	err = decoder.Decode(payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonapi.WriteBasicError(w, "400", err.Error(), "Cannot create channel")
		return
	}

	w.WriteHeader(http.StatusOK)
}
