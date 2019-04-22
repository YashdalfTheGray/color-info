package models

// StatusResponse is the server status response payload,
// gets sent back on GET /
type StatusResponse struct {
	Status   string `json:"status"`
	ServerID string `json:"serverId"`
}

// NewStatusResponse creates a new response to send out with the
// given server ID
func NewStatusResponse(serverID string) StatusResponse {
	return StatusResponse{Status: "ok", ServerID: serverID}
}
