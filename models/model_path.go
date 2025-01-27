package models

// Path Defines a single path
type Path struct {
	// A series of path commands that encodes how to draw the path.
	Path string `json:"path"`
	// The winding rule for the path (same as in SVGs). This determines whether a given point in space is inside or outside the path.
	WindingRule string `json:"windingRule"`
	// If there is a per-region fill, this refers to an ID in the `fillOverrideTable`.
	OverrideID float64 `json:"overrideID,omitempty"`
}
