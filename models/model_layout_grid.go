package models

// LayoutGrid Guides to align and place objects within a frames.
type LayoutGrid struct {
	// Orientation of the grid as a string enum  - `COLUMNS`: Vertical grid - `ROWS`: Horizontal grid - `GRID`: Square grid
	Pattern string `json:"pattern"`
	// Width of column grid or height of row grid or square grid spacing.
	SectionSize float64 `json:"sectionSize"`
	// Is the grid currently visible?
	Visible bool `json:"visible"`
	// Color of the grid
	Color RGBA `json:"color"`
	// Positioning of grid as a string enum  - `MIN`: Grid starts at the left or top of the frame - `MAX`: Grid starts at the right or bottom of the frame - `STRETCH`: Grid is stretched to fit the frame - `CENTER`: Grid is center aligned
	Alignment string `json:"alignment"`
	// Spacing in between columns and rows
	GutterSize float64 `json:"gutterSize"`
	// Spacing before the first column or row
	Offset float64 `json:"offset"`
	// Number of columns or rows
	Count          float64                   `json:"count"`
	BoundVariables *LayoutGridBoundVariables `json:"boundVariables,omitempty"`
}
