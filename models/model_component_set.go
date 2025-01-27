package models

// ComponentSet A description of a component set, which is a node containing a set of variants of a component.
type ComponentSet struct {
	// Name of the component set
	Name string `json:"name"`
	// The description of the component set as entered in the editor
	Description string `json:"description"`
	// The key of the component set
	Key string `json:"key"`
	// An array of documentation links attached to this component set
	DocumentationLinks []DocumentationLink `json:"documentationLinks,omitempty"`
	// Whether this component set is a remote component set that doesn't live in this file
	Remote bool `json:"remote,omitempty"`
}
