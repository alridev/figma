package models

// ComponentPropertyDefinition A property of a component.
type ComponentPropertyDefinition struct {
	// Type of this component property.
	Type         ComponentPropertyType `json:"type"`
	DefaultValue interface{}           `json:"defaultValue"`
	// All possible values for this property. Only exists on VARIANT properties.
	VariantOptions []string `json:"variantOptions,omitempty"`
	// Preferred values for this property. Only applicable if type is `INSTANCE_SWAP`.
	PreferredValues []InstanceSwapPreferredValue `json:"preferredValues,omitempty"`
}
