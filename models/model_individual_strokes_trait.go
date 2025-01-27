package models

// IndividualStrokesTrait struct for IndividualStrokesTrait
type IndividualStrokesTrait struct {
	// An object including the top, bottom, left, and right stroke weights. Only returned if individual stroke weights are used.
	IndividualStrokeWeights *StrokeWeights `json:"individualStrokeWeights,omitempty"`
}
