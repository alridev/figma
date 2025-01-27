package models

// BasePaint struct for BasePaint
type BasePaint struct {
	// Is the paint enabled?
	Visible bool `json:"visible,omitempty"`
	// Overall opacity of paint (colors within the paint can also have opacity values which would blend with this)
	Opacity float64 `json:"opacity,omitempty"`
	// How this node blends with nodes behind it in the scene
	BlendMode BlendMode `json:"blendMode"`
}
