package models

// TableCellNode struct for TableCellNode
type TableCellNode struct {
	// The name given to the node by the user in the tool.
	Name string `json:"name"`
	// The type of this node, represented by the string literal \"TABLE_CELL\"
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
	// An array of fill paints applied to the node.
	Fills []Paint `json:"fills"`
	// A mapping of a StyleType to style ID (see Style) of styles present on this node. The style ID can be used to look up more information about the style in the top-level styles field.
	Styles               map[string]string `json:"styles,omitempty"`
	AbsoluteBoundingBox  Rectangle         `json:"absoluteBoundingBox"`
	AbsoluteRenderBounds Rectangle         `json:"absoluteRenderBounds"`
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
	// Text contained within a text box.
	Characters string `json:"characters"`
}
