package models

// TypeStyleBoundVariables The variables bound to a particular field on this style
type TypeStyleBoundVariables struct {
	FontFamily       *VariableAlias `json:"fontFamily,omitempty"`
	FontSize         *VariableAlias `json:"fontSize,omitempty"`
	FontStyle        *VariableAlias `json:"fontStyle,omitempty"`
	FontWeight       *VariableAlias `json:"fontWeight,omitempty"`
	LetterSpacing    *VariableAlias `json:"letterSpacing,omitempty"`
	LineHeight       *VariableAlias `json:"lineHeight,omitempty"`
	ParagraphSpacing *VariableAlias `json:"paragraphSpacing,omitempty"`
	ParagraphIndent  *VariableAlias `json:"paragraphIndent,omitempty"`
}
