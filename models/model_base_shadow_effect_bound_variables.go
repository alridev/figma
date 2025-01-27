package models

// BaseShadowEffectBoundVariables The variables bound to a particular field on this shadow effect
type BaseShadowEffectBoundVariables struct {
	Radius  *VariableAlias `json:"radius,omitempty"`
	Spread  *VariableAlias `json:"spread,omitempty"`
	Color   *VariableAlias `json:"color,omitempty"`
	OffsetX *VariableAlias `json:"offsetX,omitempty"`
	OffsetY *VariableAlias `json:"offsetY,omitempty"`
}
