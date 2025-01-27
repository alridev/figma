package models

// MinimalStrokesTrait struct for MinimalStrokesTrait
type MinimalStrokesTrait struct {
	// An array of stroke paints applied to the node.
	Strokes []Paint `json:"strokes,omitempty"`
	// The weight of strokes on the node.
	StrokeWeight float64 `json:"strokeWeight,omitempty"`
	// Position of stroke relative to vector outline, as a string enum  - `INSIDE`: stroke drawn inside the shape boundary - `OUTSIDE`: stroke drawn outside the shape boundary - `CENTER`: stroke drawn centered along the shape boundary
	StrokeAlign string `json:"strokeAlign,omitempty"`
	// A string enum with value of \"MITER\", \"BEVEL\", or \"ROUND\", describing how corners in vector paths are rendered.
	StrokeJoin string `json:"strokeJoin,omitempty"`
	// An array of floating point numbers describing the pattern of dash length and gap lengths that the vector stroke will use when drawn.  For example a value of [1, 2] indicates that the stroke will be drawn with a dash of length 1 followed by a gap of length 2, repeated.
	StrokeDashes []float64 `json:"strokeDashes,omitempty"`
}
