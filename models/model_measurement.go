package models

// Measurement A pinned distance between two nodes in Dev Mode
type Measurement struct {
	Id     string              `json:"id"`
	Start  MeasurementStartEnd `json:"start"`
	End    MeasurementStartEnd `json:"end"`
	Offset MeasurementOffset   `json:"offset"`
	// When manually overridden, the displayed value of the measurement
	FreeText string `json:"freeText,omitempty"`
}
