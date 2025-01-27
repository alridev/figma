package models

// HasGeometryTrait struct for HasGeometryTrait
type HasGeometryTrait struct {
	// An array of fill paints applied to the node.
	Fills []Paint `json:"fills"`
	// A mapping of a StyleType to style ID (see Style) of styles present on this node. The style ID can be used to look up more information about the style in the top-level styles field.
	Styles map[string]string `json:"styles,omitempty"`
	// An array of stroke paints applied to the node.
	Strokes []Paint `json:"strokes,omitempty"`
	// The weight of strokes on the node.
	StrokeWeight float64 `json:"strokeWeight,omitempty"`
	// Position of stroke relative to vector outline, as a string enum  - `INSIDE`: stroke drawn inside the shape boundary - `OUTSIDE`: stroke drawn outside the shape boundary - `CENTER`: stroke drawn centered along the shape boundary
	StrokeAlign string `json:"strokeAlign,omitempty"`
	// A string enum with value of \"MITER\", \"BEVEL\", or \"ROUND\", describing how corners in vector paths are rendered.
	StrokeJoin string `json:"strokeJoin,omitempty"`
	// An array of floating point numbers describing the pattern of dash length and gap lengths that the vector stroke will use when drawn.  For example a value of [1, 2] indicates that the stroke will be drawn with a dash of length 1 followed by a gap of length 2, repeated.
	StrokeDashes []float64 `json:"strokeDashes,omitempty"`
	// Map from ID to PaintOverride for looking up fill overrides. To see which regions are overriden, you must use the `geometry=paths` option. Each path returned may have an `overrideID` which maps to this table.
	FillOverrideTable map[string]HasGeometryTraitAllOfFillOverrideTable `json:"fillOverrideTable,omitempty"`
	// Only specified if parameter `geometry=paths` is used. An array of paths representing the object fill.
	FillGeometry []Path `json:"fillGeometry,omitempty"`
	// Only specified if parameter `geometry=paths` is used. An array of paths representing the object stroke.
	StrokeGeometry []Path `json:"strokeGeometry,omitempty"`
	// A string enum describing the end caps of vector paths.
	StrokeCap string `json:"strokeCap,omitempty"`
	// Only valid if `strokeJoin` is \"MITER\". The corner angle, in degrees, below which `strokeJoin` will be set to \"BEVEL\" to avoid super sharp corners. By default this is 28.96 degrees.
	StrokeMiterAngle float64 `json:"strokeMiterAngle,omitempty"`
}
