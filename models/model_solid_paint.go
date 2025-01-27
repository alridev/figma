package models

// SolidPaint struct for SolidPaint
type SolidPaint struct {
	// Is the paint enabled?
	Visible bool `json:"visible,omitempty"`
	// Overall opacity of paint (colors within the paint can also have opacity values which would blend with this)
	Opacity float64 `json:"opacity,omitempty"`
	// How this node blends with nodes behind it in the scene
	BlendMode BlendMode `json:"blendMode"`
	// The string literal \"SOLID\" representing the paint's type. Always check the `type` before reading other properties.
	Type string `json:"type"`
	// Solid color of the paint
	Color          RGBA                           `json:"color"`
	BoundVariables *SolidPaintAllOfBoundVariables `json:"boundVariables,omitempty"`
}
