package models

// MinimalFillsTrait struct for MinimalFillsTrait
type MinimalFillsTrait struct {
	// An array of fill paints applied to the node.
	Fills []Paint `json:"fills"`
	// A mapping of a StyleType to style ID (see Style) of styles present on this node. The style ID can be used to look up more information about the style in the top-level styles field.
	Styles map[string]string `json:"styles,omitempty"`
}
