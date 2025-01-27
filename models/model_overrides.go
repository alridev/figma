package models

// Overrides Fields directly overridden on an instance. Inherited overrides are not included.
type Overrides struct {
	// A unique ID for a node.
	Id string `json:"id"`
	// An array of properties.
	OverriddenFields []string `json:"overriddenFields"`
}
