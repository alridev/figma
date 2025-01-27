package models

// DirectionalTransition Describes an animation used when navigating in a prototype.
type DirectionalTransition struct {
	Type      string `json:"type"`
	Direction string `json:"direction"`
	// The duration of the transition in milliseconds.
	Duration float64 `json:"duration"`
	// The easing curve of the transition.
	Easing Easing `json:"easing"`
	// When the transition `type` is `\"SMART_ANIMATE\"` or when `matchLayers` is `true`, then the transition will be performed using smart animate, which attempts to match corresponding layers an interpolate other properties during the animation.
	MatchLayers bool `json:"matchLayers,omitempty"`
}
