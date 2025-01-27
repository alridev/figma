package models

// MeasurementOffsetOuter Measurement offset relative to the outside of the start node
type MeasurementOffsetOuter struct {
	Type  string  `json:"type"`
	Fixed float64 `json:"fixed"`
}
