package models

// Component A description of a main component. Helps you identify which component instances are attached to.
type Component struct {
	// Name of the component
	Name string `json:"name"`
	// The description of the component as entered in the editor
	Description string `json:"description"`
	// The key of the component
	Key string `json:"key"`
	// The ID of the component set if the component belongs to one
	ComponentSetId string `json:"componentSetId,omitempty"`
	// An array of documentation links attached to this component
	DocumentationLinks []DocumentationLink `json:"documentationLinks"`
	// Whether this component is a remote component that doesn't live in this file
	Remote bool `json:"remote"`
}
