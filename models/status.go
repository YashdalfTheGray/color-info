package models

// Status is the server status response payload, gets sent back on GET /
type Status struct {
	Status   string `json:"status"`
	ServerID string `json:"serverId"`
}
