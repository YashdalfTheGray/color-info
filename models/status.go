package models

// StatusResponse is the server status response payload,
// gets sent back on GET /
type StatusResponse struct {
	Status   string `json:"status"`
	ServerID string `json:"serverId"`
}
