package models

// MeasurementOffsetInner Measurement offset relative to the inside of the start node
type MeasurementOffsetInner struct {
	Type     string  `json:"type"`
	Relative float64 `json:"relative"`
}
