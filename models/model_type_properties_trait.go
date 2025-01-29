package models

// TypePropertiesTrait struct for TypePropertiesTrait
type TypePropertiesTrait struct {
	// The raw characters in the text node.
	Characters string `json:"characters"`
	// Style of text including font family and weight.
	Style TypeStyle `json:"style"`
	// The array corresponds to characters in the text box, where each element references the 'styleOverrideTable' to apply specific styles to each character. The array's length can be less than or equal to the number of characters due to the removal of trailing zeros. Elements with a value of 0 indicate characters that use the default type style. If the array is shorter than the total number of characters, the characters beyond the array's length also use the default style.
	CharacterStyleOverrides []int `json:"characterStyleOverrides"`
	// Internal property, preserved for backward compatibility. Avoid using this value.
	LayoutVersion float64 `json:"layoutVersion,omitempty"`
	// Map from ID to TypeStyle for looking up style overrides.
	StyleOverrideTable map[string]TypeStyle `json:"styleOverrideTable"`
	// An array with the same number of elements as lines in the text node, where lines are delimited by newline or paragraph separator characters. Each element in the array corresponds to the list type of a specific line. List types are represented as string enums with one of these possible values:  - `NONE`: Not a list item. - `ORDERED`: Text is an ordered list (numbered). - `UNORDERED`: Text is an unordered list (bulleted).
	LineTypes []string `json:"lineTypes"`
	// An array with the same number of elements as lines in the text node, where lines are delimited by newline or paragraph separator characters. Each element in the array corresponds to the indentation level of a specific line.
	LineIndentations []float64 `json:"lineIndentations"`
}
