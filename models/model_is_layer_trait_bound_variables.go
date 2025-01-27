package models

// IsLayerTraitBoundVariables A mapping of field to the variables applied to this field. Most fields will only map to a single `VariableAlias`. However, for properties like `fills`, `strokes`, `size`, `componentProperties`, and `textRangeFills`, it is possible to have multiple variables bound to the field.
type IsLayerTraitBoundVariables struct {
	Size                    *IsLayerTraitBoundVariablesSize                    `json:"size,omitempty"`
	IndividualStrokeWeights *IsLayerTraitBoundVariablesIndividualStrokeWeights `json:"individualStrokeWeights,omitempty"`
	Characters              *VariableAlias                                     `json:"characters,omitempty"`
	ItemSpacing             *VariableAlias                                     `json:"itemSpacing,omitempty"`
	PaddingLeft             *VariableAlias                                     `json:"paddingLeft,omitempty"`
	PaddingRight            *VariableAlias                                     `json:"paddingRight,omitempty"`
	PaddingTop              *VariableAlias                                     `json:"paddingTop,omitempty"`
	PaddingBottom           *VariableAlias                                     `json:"paddingBottom,omitempty"`
	Visible                 *VariableAlias                                     `json:"visible,omitempty"`
	TopLeftRadius           *VariableAlias                                     `json:"topLeftRadius,omitempty"`
	TopRightRadius          *VariableAlias                                     `json:"topRightRadius,omitempty"`
	BottomLeftRadius        *VariableAlias                                     `json:"bottomLeftRadius,omitempty"`
	BottomRightRadius       *VariableAlias                                     `json:"bottomRightRadius,omitempty"`
	MinWidth                *VariableAlias                                     `json:"minWidth,omitempty"`
	MaxWidth                *VariableAlias                                     `json:"maxWidth,omitempty"`
	MinHeight               *VariableAlias                                     `json:"minHeight,omitempty"`
	MaxHeight               *VariableAlias                                     `json:"maxHeight,omitempty"`
	CounterAxisSpacing      *VariableAlias                                     `json:"counterAxisSpacing,omitempty"`
	Opacity                 *VariableAlias                                     `json:"opacity,omitempty"`
	FontFamily              []VariableAlias                                    `json:"fontFamily,omitempty"`
	FontSize                []VariableAlias                                    `json:"fontSize,omitempty"`
	FontStyle               []VariableAlias                                    `json:"fontStyle,omitempty"`
	FontWeight              []VariableAlias                                    `json:"fontWeight,omitempty"`
	LetterSpacing           []VariableAlias                                    `json:"letterSpacing,omitempty"`
	LineHeight              []VariableAlias                                    `json:"lineHeight,omitempty"`
	ParagraphSpacing        []VariableAlias                                    `json:"paragraphSpacing,omitempty"`
	ParagraphIndent         []VariableAlias                                    `json:"paragraphIndent,omitempty"`
	Fills                   []VariableAlias                                    `json:"fills,omitempty"`
	Strokes                 []VariableAlias                                    `json:"strokes,omitempty"`
	ComponentProperties     map[string]VariableAlias                           `json:"componentProperties,omitempty"`
	TextRangeFills          []VariableAlias                                    `json:"textRangeFills,omitempty"`
	Effects                 []VariableAlias                                    `json:"effects,omitempty"`
	LayoutGrids             []VariableAlias                                    `json:"layoutGrids,omitempty"`
}
