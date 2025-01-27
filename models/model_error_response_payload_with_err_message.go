package models

// ErrorResponsePayloadWithErrMessage A response indicating an error occurred.
type ErrorResponsePayloadWithErrMessage struct {
	// Status code
	Status float64 `json:"status"`
	// A string describing the error
	Err string `json:"err"`
}
