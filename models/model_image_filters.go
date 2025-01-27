package models

// ImageFilters Image filters to apply to the node.
type ImageFilters struct {
	Exposure    float64 `json:"exposure,omitempty"`
	Contrast    float64 `json:"contrast,omitempty"`
	Saturation  float64 `json:"saturation,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
	Tint        float64 `json:"tint,omitempty"`
	Highlights  float64 `json:"highlights,omitempty"`
	Shadows     float64 `json:"shadows,omitempty"`
}
