package models

// DocumentationLink Represents a link to documentation for a component or component set.
type DocumentationLink struct {
	// Should be a valid URI (e.g. https://www.figma.com).
	Uri string `json:"uri"`
}
