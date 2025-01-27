package models

// FrameNode struct for FrameNode
type FrameNode struct {
	// The name given to the node by the user in the tool.
	Name string `json:"name"`
	// The type of this node, represented by the string literal \"FRAME\"
	Type string `json:"type"`
	// A string uniquely identifying this node within the document.
	Id string `json:"id"`
	// Whether or not the node is visible on the canvas.
	Visible bool `json:"visible,omitempty"`
	// If true, layer is locked and cannot be edited
	Locked bool `json:"locked,omitempty"`
	// Whether the layer is fixed while the parent is scrolling
	// Deprecated
	IsFixed bool `json:"isFixed,omitempty"`
	// How layer should be treated when the frame is resized
	ScrollBehavior string `json:"scrollBehavior"`
	// The rotation of the node, if not 0.
	Rotation float64 `json:"rotation,omitempty"`
	// A mapping of a layer's property to component property name of component properties attached to this node. The component property name can be used to look up more information on the corresponding component's or component set's componentPropertyDefinitions.
	ComponentPropertyReferences map[string]string           `json:"componentPropertyReferences,omitempty"`
	PluginData                  interface{}                 `json:"pluginData,omitempty"`
	SharedPluginData            interface{}                 `json:"sharedPluginData,omitempty"`
	BoundVariables              *IsLayerTraitBoundVariables `json:"boundVariables,omitempty"`
	// A mapping of variable collection ID to mode ID representing the explicitly set modes for this node.
	ExplicitVariableModes map[string]string `json:"explicitVariableModes,omitempty"`
	// How this node blends with nodes behind it in the scene (see blend mode section for more details)
	BlendMode BlendMode `json:"blendMode"`
	// Opacity of the node
	Opacity float64 `json:"opacity,omitempty"`
	// An array of nodes that are direct children of this node
	Children             []Node `json:"children"`
	AbsoluteBoundingBox  Rectangle       `json:"absoluteBoundingBox"`
	AbsoluteRenderBounds Rectangle       `json:"absoluteRenderBounds"`
	// Keep height and width constrained to same ratio.
	PreserveRatio bool `json:"preserveRatio,omitempty"`
	// Horizontal and vertical layout constraints for node.
	Constraints *LayoutConstraint `json:"constraints,omitempty"`
	// A transformation matrix is standard way in computer graphics to represent translation and rotation. These are the top two rows of a 3x3 matrix. The bottom row of the matrix is assumed to be [0, 0, 1]. This is known as an affine transform and is enough to represent translation, rotation, and skew.  The identity transform is [[1, 0, 0], [0, 1, 0]].  A translation matrix will typically look like:  ``` [[1, 0, tx],   [0, 1, ty]] ```  and a rotation matrix will typically look like:  ``` [[cos(angle), sin(angle), 0],   [-sin(angle), cos(angle), 0]] ```  Another way to think about this transform is as three vectors:  - The x axis (t[0][0], t[1][0]) - The y axis (t[0][1], t[1][1]) - The translation offset (t[0][2], t[1][2])  The most common usage of the Transform matrix is the `relativeTransform property`. This particular usage of the matrix has a few additional restrictions. The translation offset can take on any value but we do enforce that the axis vectors are unit vectors (i.e. have length 1). The axes are not required to be at 90° angles to each other.
	RelativeTransform [][]float64 `json:"relativeTransform,omitempty"`
	// Width and height of element. This is different from the width and height of the bounding box in that the absolute bounding box represents the element after scaling and rotation. Only present if `geometry=paths` is passed.
	Size *Vector `json:"size,omitempty"`
	// Determines if the layer should stretch along the parent's counter axis. This property is only provided for direct children of auto-layout frames.  - `INHERIT` - `STRETCH`  In previous versions of auto layout, determined how the layer is aligned inside an auto-layout frame. This property is only provided for direct children of auto-layout frames.  - `MIN` - `CENTER` - `MAX` - `STRETCH`  In horizontal auto-layout frames, \"MIN\" and \"MAX\" correspond to \"TOP\" and \"BOTTOM\". In vertical auto-layout frames, \"MIN\" and \"MAX\" correspond to \"LEFT\" and \"RIGHT\".
	LayoutAlign string `json:"layoutAlign,omitempty"`
	// This property is applicable only for direct children of auto-layout frames, ignored otherwise. Determines whether a layer should stretch along the parent's primary axis. A `0` corresponds to a fixed size and `1` corresponds to stretch.
	LayoutGrow float64 `json:"layoutGrow,omitempty"`
	// Determines whether a layer's size and position should be determined by auto-layout settings or manually adjustable.
	LayoutPositioning string `json:"layoutPositioning,omitempty"`
	// The minimum width of the frame. This property is only applicable for auto-layout frames or direct children of auto-layout frames.
	MinWidth float64 `json:"minWidth,omitempty"`
	// The maximum width of the frame. This property is only applicable for auto-layout frames or direct children of auto-layout frames.
	MaxWidth float64 `json:"maxWidth,omitempty"`
	// The minimum height of the frame. This property is only applicable for auto-layout frames or direct children of auto-layout frames.
	MinHeight float64 `json:"minHeight,omitempty"`
	// The maximum height of the frame. This property is only applicable for auto-layout frames or direct children of auto-layout frames.
	MaxHeight float64 `json:"maxHeight,omitempty"`
	// The horizontal sizing setting on this auto-layout frame or frame child. - `FIXED` - `HUG`: only valid on auto-layout frames and text nodes - `FILL`: only valid on auto-layout frame children
	LayoutSizingHorizontal string `json:"layoutSizingHorizontal,omitempty"`
	// The vertical sizing setting on this auto-layout frame or frame child. - `FIXED` - `HUG`: only valid on auto-layout frames and text nodes - `FILL`: only valid on auto-layout frame children
	LayoutSizingVertical string `json:"layoutSizingVertical,omitempty"`
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
	// Radius of each corner if a single radius is set for all corners
	CornerRadius float64 `json:"cornerRadius,omitempty"`
	// A value that lets you control how \"smooth\" the corners are. Ranges from 0 to 1. 0 is the default and means that the corner is perfectly circular. A value of 0.6 means the corner matches the iOS 7 \"squircle\" icon shape. Other values produce various other curves.
	CornerSmoothing float64 `json:"cornerSmoothing,omitempty"`
	// Array of length 4 of the radius of each corner of the frame, starting in the top left and proceeding clockwise.  Values are given in the order top-left, top-right, bottom-right, bottom-left.
	RectangleCornerRadii []float64 `json:"rectangleCornerRadii,omitempty"`
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
	// An array of export settings representing images to export from the node.
	ExportSettings []ExportSetting `json:"exportSettings,omitempty"`
	// An array of effects attached to this node (see effects section for more details)
	Effects []Effect `json:"effects"`
	// Does this node mask sibling nodes in front of it?
	IsMask bool `json:"isMask,omitempty"`
	// If this layer is a mask, this property describes the operation used to mask the layer's siblings. The value may be one of the following:  - ALPHA: the mask node's alpha channel will be used to determine the opacity of each pixel in the masked result. - VECTOR: if the mask node has visible fill paints, every pixel inside the node's fill regions will be fully visible in the masked result. If the mask has visible stroke paints, every pixel inside the node's stroke regions will be fully visible in the masked result. - LUMINANCE: the luminance value of each pixel of the mask node will be used to determine the opacity of that pixel in the masked result.
	MaskType string `json:"maskType,omitempty"`
	// True if maskType is VECTOR. This field is deprecated; use maskType instead.
	// Deprecated
	IsMaskOutline bool `json:"isMaskOutline,omitempty"`
	// Node ID of node to transition to in prototyping
	TransitionNodeID string `json:"transitionNodeID,omitempty"`
	// The duration of the prototyping transition on this node (in milliseconds). This will override the default transition duration on the prototype, for this node.
	TransitionDuration float64 `json:"transitionDuration,omitempty"`
	// The easing curve used in the prototyping transition on this node.
	TransitionEasing *EasingType   `json:"transitionEasing,omitempty"`
	Interactions     []Interaction `json:"interactions,omitempty"`
	// An object including the top, bottom, left, and right stroke weights. Only returned if individual stroke weights are used.
	IndividualStrokeWeights *StrokeWeights           `json:"individualStrokeWeights,omitempty"`
	DevStatus               *DevStatusTraitDevStatus `json:"devStatus,omitempty"`
}
