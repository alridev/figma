package models

// StrokeWeights Individual stroke weights
type StrokeWeights struct {
	// The top stroke weight.
	Top float64 `json:"top"`
	// The right stroke weight.
	Right float64 `json:"right"`
	// The bottom stroke weight.
	Bottom float64 `json:"bottom"`
	// The left stroke weight.
	Left float64 `json:"left"`
}
