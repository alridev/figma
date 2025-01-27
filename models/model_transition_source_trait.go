package models

// TransitionSourceTrait struct for TransitionSourceTrait
type TransitionSourceTrait struct {
	// Node ID of node to transition to in prototyping
	TransitionNodeID string `json:"transitionNodeID,omitempty"`
	// The duration of the prototyping transition on this node (in milliseconds). This will override the default transition duration on the prototype, for this node.
	TransitionDuration float64 `json:"transitionDuration,omitempty"`
	// The easing curve used in the prototyping transition on this node.
	TransitionEasing *EasingType   `json:"transitionEasing,omitempty"`
	Interactions     []Interaction `json:"interactions,omitempty"`
}
