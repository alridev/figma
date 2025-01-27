package models

// OnKeyDownTrigger struct for OnKeyDownTrigger
type OnKeyDownTrigger struct {
	Type     string    `json:"type"`
	Device   string    `json:"device"`
	KeyCodes []float64 `json:"keyCodes"`
}
