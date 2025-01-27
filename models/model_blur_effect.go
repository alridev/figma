package models

// BlurEffect A blur effect
type BlurEffect struct {
	// A string literal representing the effect's type. Always check the type before reading other properties.
	Type string `json:"type"`
	// Whether this blur is active.
	Visible bool `json:"visible"`
	// Radius of the blur effect
	Radius         float64                   `json:"radius"`
	BoundVariables *BlurEffectBoundVariables `json:"boundVariables,omitempty"`
}
