package openapi

type ErrorResponse struct {
	ErrorCode string `json:"error_code,omitempty"`

	ErrorMessage string `json:"error_message,omitempty"`
}
