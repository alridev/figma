package models

// InnerShadowEffect struct for InnerShadowEffect
type InnerShadowEffect struct {
	// The color of the shadow
	Color RGBA `json:"color"`
	// Blend mode of the shadow
	BlendMode BlendMode `json:"blendMode"`
	// How far the shadow is projected in the x and y directions
	Offset Vector `json:"offset"`
	// Radius of the blur effect (applies to shadows as well)
	Radius float64 `json:"radius"`
	// The distance by which to expand (or contract) the shadow.  For drop shadows, a positive `spread` value creates a shadow larger than the node, whereas a negative value creates a shadow smaller than the node.  For inner shadows, a positive `spread` value contracts the shadow. Spread values are only accepted on rectangles and ellipses, or on frames, components, and instances with visible fill paints and `clipsContent` enabled. When left unspecified, the default value is 0.
	Spread float64 `json:"spread,omitempty"`
	// Whether this shadow is visible.
	Visible        bool                            `json:"visible"`
	BoundVariables *BaseShadowEffectBoundVariables `json:"boundVariables,omitempty"`
	// A string literal representing the effect's type. Always check the type before reading other properties.
	Type string `json:"type,omitempty"`
}
