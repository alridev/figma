package models

// ColorStop A single color stop with its position along the gradient axis, color, and bound variables if any
type ColorStop struct {
	// Value between 0 and 1 representing position along gradient axis.
	Position float64 `json:"position"`
	// Color attached to corresponding position.
	Color          RGBA                     `json:"color"`
	BoundVariables *ColorStopBoundVariables `json:"boundVariables,omitempty"`
}
