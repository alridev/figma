package models

// GradientPaint struct for GradientPaint
type GradientPaint struct {
	// Is the paint enabled?
	Visible bool `json:"visible,omitempty"`
	// Overall opacity of paint (colors within the paint can also have opacity values which would blend with this)
	Opacity float64 `json:"opacity,omitempty"`
	// How this node blends with nodes behind it in the scene
	BlendMode BlendMode `json:"blendMode"`
	// The string literal representing the paint's type. Always check the `type` before reading other properties.
	Type string `json:"type"`
	// This field contains three vectors, each of which are a position in normalized object space (normalized object space is if the top left corner of the bounding box of the object is (0, 0) and the bottom right is (1,1)). The first position corresponds to the start of the gradient (value 0 for the purposes of calculating gradient stops), the second position is the end of the gradient (value 1), and the third handle position determines the width of the gradient.
	GradientHandlePositions []Vector `json:"gradientHandlePositions"`
	// Positions of key points along the gradient axis with the colors anchored there. Colors along the gradient are interpolated smoothly between neighboring gradient stops.
	GradientStops []ColorStop `json:"gradientStops"`
}
