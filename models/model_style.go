package models

// Style A set of properties that can be applied to nodes and published. Styles for a property can be created in the corresponding property's panel while editing a file.
type Style struct {
	// Name of the style
	Name string `json:"name"`
	// Description of the style
	Description string `json:"description"`
	// The key of the style
	Key string `json:"key"`
	// Whether this style is a remote style that doesn't live in this file
	Remote    bool      `json:"remote"`
	StyleType StyleType `json:"styleType"`
}
