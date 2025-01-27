package models

// SimpleTransition Describes an animation used when navigating in a prototype.
type SimpleTransition struct {
	Type string `json:"type"`
	// The duration of the transition in milliseconds.
	Duration float64 `json:"duration"`
	// The easing curve of the transition.
	Easing Easing `json:"easing"`
}
