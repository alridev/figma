package models

// EasingEasingFunctionSpring A spring function that defines the easing.
type EasingEasingFunctionSpring struct {
	Mass      float64 `json:"mass"`
	Stiffness float64 `json:"stiffness"`
	Damping   float64 `json:"damping"`
}
