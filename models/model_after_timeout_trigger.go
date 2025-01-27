package models

// AfterTimeoutTrigger struct for AfterTimeoutTrigger
type AfterTimeoutTrigger struct {
	Type    string  `json:"type"`
	Timeout float64 `json:"timeout"`
}
