package models

// MeasurementStartEnd The node and side a measurement is pinned to
type MeasurementStartEnd struct {
	NodeId string `json:"nodeId"`
	Side   string `json:"side"`
}
