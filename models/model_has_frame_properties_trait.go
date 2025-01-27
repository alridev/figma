package models

// HasFramePropertiesTrait struct for HasFramePropertiesTrait
type HasFramePropertiesTrait struct {
	// Whether or not this node clip content outside of its bounds
	ClipsContent bool `json:"clipsContent"`
	// Background of the node. This is deprecated, as backgrounds for frames are now in the `fills` field.
	Background []Paint `json:"background,omitempty"`
	// Background color of the node. This is deprecated, as frames now support more than a solid color as a background. Please use the `fills` field instead.
	// Deprecated
	BackgroundColor *RGBA `json:"backgroundColor,omitempty"`
	// An array of layout grids attached to this node (see layout grids section for more details). GROUP nodes do not have this attribute
	LayoutGrids []LayoutGrid `json:"layoutGrids,omitempty"`
	// Whether a node has primary axis scrolling, horizontal or vertical.
	OverflowDirection string `json:"overflowDirection,omitempty"`
	// Whether this layer uses auto-layout to position its children.
	LayoutMode string `json:"layoutMode,omitempty"`
	// Whether the primary axis has a fixed length (determined by the user) or an automatic length (determined by the layout engine). This property is only applicable for auto-layout frames.
	PrimaryAxisSizingMode string `json:"primaryAxisSizingMode,omitempty"`
	// Whether the counter axis has a fixed length (determined by the user) or an automatic length (determined by the layout engine). This property is only applicable for auto-layout frames.
	CounterAxisSizingMode string `json:"counterAxisSizingMode,omitempty"`
	// Determines how the auto-layout frame's children should be aligned in the primary axis direction. This property is only applicable for auto-layout frames.
	PrimaryAxisAlignItems string `json:"primaryAxisAlignItems,omitempty"`
	// Determines how the auto-layout frame's children should be aligned in the counter axis direction. This property is only applicable for auto-layout frames.
	CounterAxisAlignItems string `json:"counterAxisAlignItems,omitempty"`
	// The padding between the left border of the frame and its children. This property is only applicable for auto-layout frames.
	PaddingLeft float64 `json:"paddingLeft,omitempty"`
	// The padding between the right border of the frame and its children. This property is only applicable for auto-layout frames.
	PaddingRight float64 `json:"paddingRight,omitempty"`
	// The padding between the top border of the frame and its children. This property is only applicable for auto-layout frames.
	PaddingTop float64 `json:"paddingTop,omitempty"`
	// The padding between the bottom border of the frame and its children. This property is only applicable for auto-layout frames.
	PaddingBottom float64 `json:"paddingBottom,omitempty"`
	// The distance between children of the frame. Can be negative. This property is only applicable for auto-layout frames.
	ItemSpacing float64 `json:"itemSpacing,omitempty"`
	// Determines the canvas stacking order of layers in this frame. When true, the first layer will be draw on top. This property is only applicable for auto-layout frames.
	ItemReverseZIndex bool `json:"itemReverseZIndex,omitempty"`
	// Determines whether strokes are included in layout calculations. When true, auto-layout frames behave like css \"box-sizing: border-box\". This property is only applicable for auto-layout frames.
	StrokesIncludedInLayout bool `json:"strokesIncludedInLayout,omitempty"`
	// Whether this auto-layout frame has wrapping enabled.
	LayoutWrap string `json:"layoutWrap,omitempty"`
	// The distance between wrapped tracks of an auto-layout frame. This property is only applicable for auto-layout frames with `layoutWrap: \"WRAP\"`
	CounterAxisSpacing float64 `json:"counterAxisSpacing,omitempty"`
	// Determines how the auto-layout frame’s wrapped tracks should be aligned in the counter axis direction. This property is only applicable for auto-layout frames with `layoutWrap: \"WRAP\"`.
	CounterAxisAlignContent string `json:"counterAxisAlignContent,omitempty"`
}
