package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/yashdalfthegray/color-info/models"
)

// NewStatusHandler configures a status handler with a server id
// and returns it. The status handler will respond with a JSON
// { status: "ok", serverId: "some string" }
func NewStatusHandler(serverID string) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(models.Status{
			Status:   "ok",
			ServerID: serverID,
		}); err != nil {
			panic(err)
		}
	}
}
