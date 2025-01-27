package models

// RGB An RGB color
type RGB struct {
	// Red channel value, between 0 and 1.
	R float64 `json:"r"`
	// Green channel value, between 0 and 1.
	G float64 `json:"g"`
	// Blue channel value, between 0 and 1.
	B float64 `json:"b"`
}
