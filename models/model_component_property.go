package models

// ComponentProperty A property of a component.
type ComponentProperty struct {
	// Type of this component property.
	Type  ComponentPropertyType  `json:"type"`
	Value ComponentPropertyValue `json:"value"`
	// Preferred values for this property. Only applicable if type is `INSTANCE_SWAP`.
	PreferredValues []InstanceSwapPreferredValue     `json:"preferredValues,omitempty"`
	BoundVariables  *ComponentPropertyBoundVariables `json:"boundVariables,omitempty"`
}
