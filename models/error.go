package models

// ErrorResponse is the server status response payload,
// gets sent back on GET /
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// NewErrorResponse creates a new response to send out with the
// given server ID
func NewErrorResponse(err string) ErrorResponse {
	return ErrorResponse{Status: "error", Message: err}
}
