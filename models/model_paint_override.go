package models

// PaintOverride Paint metadata to override default paints.
type PaintOverride struct {
	// Paints applied to characters.
	Fills []Paint `json:"fills,omitempty"`
	// ID of style node, if any, that this inherits fill data from.
	InheritFillStyleId string `json:"inheritFillStyleId,omitempty"`
}
