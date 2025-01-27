package models

// TypeStyle Metadata for character formatting.
type TypeStyle struct {
	// Font family of text (standard name).
	FontFamily string `json:"fontFamily,omitempty"`
	// PostScript font name.
	FontPostScriptName string `json:"fontPostScriptName,omitempty"`
	// Describes visual weight or emphasis, such as Bold or Italic.
	FontStyle string `json:"fontStyle,omitempty"`
	// Space between paragraphs in px, 0 if not present.
	ParagraphSpacing float64 `json:"paragraphSpacing,omitempty"`
	// Paragraph indentation in px, 0 if not present.
	ParagraphIndent float64 `json:"paragraphIndent,omitempty"`
	// Space between list items in px, 0 if not present.
	ListSpacing float64 `json:"listSpacing,omitempty"`
	// Whether or not text is italicized.
	Italic bool `json:"italic,omitempty"`
	// Numeric font weight.
	FontWeight float64 `json:"fontWeight,omitempty"`
	// Font size in px.
	FontSize float64 `json:"fontSize,omitempty"`
	// Text casing applied to the node, default is the original casing.
	TextCase string `json:"textCase,omitempty"`
	// Text decoration applied to the node, default is none.
	TextDecoration string `json:"textDecoration,omitempty"`
	// Dimensions along which text will auto resize, default is that the text does not auto-resize. TRUNCATE means that the text will be shortened and trailing text will be replaced with \"…\" if the text contents is larger than the bounds. `TRUNCATE` as a return value is deprecated and will be removed in a future version. Read from `textTruncation` instead.
	TextAutoResize string `json:"textAutoResize,omitempty"`
	// Whether this text node will truncate with an ellipsis when the text contents is larger than the text node.
	TextTruncation string `json:"textTruncation,omitempty"`
	// When `textTruncation: \"ENDING\"` is set, `maxLines` determines how many lines a text node can grow to before it truncates.
	MaxLines float64 `json:"maxLines,omitempty"`
	// Horizontal text alignment as string enum.
	TextAlignHorizontal string `json:"textAlignHorizontal,omitempty"`
	// Vertical text alignment as string enum.
	TextAlignVertical string `json:"textAlignVertical,omitempty"`
	// Space between characters in px.
	LetterSpacing float64 `json:"letterSpacing,omitempty"`
	// An array of fill paints applied to the characters.
	Fills []Paint `json:"fills,omitempty"`
	// Link to a URL or frame.
	Hyperlink *Hyperlink `json:"hyperlink,omitempty"`
	// A map of OpenType feature flags to 1 or 0, 1 if it is enabled and 0 if it is disabled. Note that some flags aren't reflected here. For example, SMCP (small caps) is still represented by the `textCase` field.
	OpentypeFlags map[string]float64 `json:"opentypeFlags,omitempty"`
	// Line height in px.
	LineHeightPx float64 `json:"lineHeightPx,omitempty"`
	// Line height as a percentage of normal line height. This is deprecated; in a future version of the API only lineHeightPx and lineHeightPercentFontSize will be returned.
	LineHeightPercent float64 `json:"lineHeightPercent,omitempty"`
	// Line height as a percentage of the font size. Only returned when `lineHeightPercent` (deprecated) is not 100.
	LineHeightPercentFontSize float64 `json:"lineHeightPercentFontSize,omitempty"`
	// The unit of the line height value specified by the user.
	LineHeightUnit string                  `json:"lineHeightUnit,omitempty"`
	BoundVariables *TypeStyleBoundVariables `json:"boundVariables,omitempty"`
	//  Whether or not this style has overrides over a text style. The possible fields to override are semanticWeight, semanticItalic, hyperlink, and textDecoration. If this is true, then those fields are overrides if present.
	IsOverrideOverTextStyle bool `json:"isOverrideOverTextStyle,omitempty"`
	// Indicates how the font weight was overridden when there is a text style override.
	SemanticWeight string `json:"semanticWeight,omitempty"`
	// Indicates how the font style was overridden when there is a text style override.
	SemanticItalic string `json:"semanticItalic,omitempty"`
}
