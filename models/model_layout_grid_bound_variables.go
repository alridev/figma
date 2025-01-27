package models

// LayoutGridBoundVariables The variables bound to a particular field on this layout grid
type LayoutGridBoundVariables struct {
	GutterSize  *VariableAlias `json:"gutterSize,omitempty"`
	NumSections *VariableAlias `json:"numSections,omitempty"`
	SectionSize *VariableAlias `json:"sectionSize,omitempty"`
	Offset      *VariableAlias `json:"offset,omitempty"`
}
