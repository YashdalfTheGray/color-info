package handlers

import (
	"fmt"
	"net/http"
)

// PingHandler responds with string "pong" when GET /ping
func PingHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "pong\n")
}
