package models

// HasBlendModeAndOpacityTrait struct for HasBlendModeAndOpacityTrait
type HasBlendModeAndOpacityTrait struct {
	// How this node blends with nodes behind it in the scene (see blend mode section for more details)
	BlendMode BlendMode `json:"blendMode"`
	// Opacity of the node
	Opacity float64 `json:"opacity,omitempty"`
}
