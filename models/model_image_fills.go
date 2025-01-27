package models

// ImageFills struct for ImageFills
type ImageFills struct {
	// For successful requests, this value is always `false`.
	Error bool `json:"error"`
	// Status code
	Status float64        `json:"status"`
	Meta   ImageFillsMeta `json:"meta"`
}
