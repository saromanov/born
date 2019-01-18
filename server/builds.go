package server

import (
	"fmt"
	"net/http"
)

// createBuild provides creating of the new build
func createBuild(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-BORN-TOKEN")
	fmt.Println(token)
}
