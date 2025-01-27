package models

// DevStatusTraitDevStatus Represents whether or not a node has a particular handoff (or dev) status applied to it.
type DevStatusTraitDevStatus struct {
	Type string `json:"type"`
	// An optional field where the designer can add more information about the design and what has changed.
	Description string `json:"description,omitempty"`
}
