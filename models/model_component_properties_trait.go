package models

// ComponentPropertiesTrait struct for ComponentPropertiesTrait
type ComponentPropertiesTrait struct {
	// A mapping of name to `ComponentPropertyDefinition` for every component property on this component. Each property has a type, defaultValue, and other optional values.
	ComponentPropertyDefinitions map[string]ComponentPropertyDefinition `json:"componentPropertyDefinitions,omitempty"`
}
