package models

// CornerTrait struct for CornerTrait
type CornerTrait struct {
	// Radius of each corner if a single radius is set for all corners
	CornerRadius float64 `json:"cornerRadius,omitempty"`
	// A value that lets you control how \"smooth\" the corners are. Ranges from 0 to 1. 0 is the default and means that the corner is perfectly circular. A value of 0.6 means the corner matches the iOS 7 \"squircle\" icon shape. Other values produce various other curves.
	CornerSmoothing float64 `json:"cornerSmoothing,omitempty"`
	// Array of length 4 of the radius of each corner of the frame, starting in the top left and proceeding clockwise.  Values are given in the order top-left, top-right, bottom-right, bottom-left.
	RectangleCornerRadii []float64 `json:"rectangleCornerRadii,omitempty"`
}
