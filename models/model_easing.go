package models

// Easing Describes an easing curve.
type Easing struct {
	// The type of easing curve.
	Type                      EasingType                       `json:"type"`
	EasingFunctionCubicBezier *EasingEasingFunctionCubicBezier `json:"easingFunctionCubicBezier,omitempty"`
	EasingFunctionSpring      *EasingEasingFunctionSpring      `json:"easingFunctionSpring,omitempty"`
}
