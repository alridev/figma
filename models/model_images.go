package models

// Images struct for Images
type Images struct {
	// For successful requests, this value is always `null`.
	Err interface{} `json:"err"`
	// A map from node IDs to URLs of the rendered images.
	Images map[string]string `json:"images"`
}
