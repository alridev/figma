package models

// InlineObject struct for InlineObject
type InlineObject struct {
	// Status code
	Status float64 `json:"status"`
	// A string describing the error
	Err string `json:"err"`
}
