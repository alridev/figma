package models

// EasingEasingFunctionCubicBezier A cubic bezier curve that defines the easing.
type EasingEasingFunctionCubicBezier struct {
	// The x component of the first control point.
	X1 float64 `json:"x1"`
	// The y component of the first control point.
	Y1 float64 `json:"y1"`
	// The x component of the second control point.
	X2 float64 `json:"x2"`
	// The y component of the second control point.
	Y2 float64 `json:"y2"`
}
