package models

// InstanceSwapPreferredValue Instance swap preferred value.
type InstanceSwapPreferredValue struct {
	// Type of node for this preferred value.
	Type string `json:"type"`
	// Key of this component or component set.
	Key string `json:"key"`
}
