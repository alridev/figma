package models

// CanvasNode struct for CanvasNode
type CanvasNode struct {
	// The name given to the node by the user in the tool.
	Name string `json:"name"`
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
	// An array of export settings representing images to export from the node.
	ExportSettings []ExportSetting `json:"exportSettings,omitempty"`
	Children       []Node          `json:"children"`
	// Background color of the canvas.
	BackgroundColor RGBA `json:"backgroundColor"`
	// CanvasNode ID that corresponds to the start frame for prototypes. This is deprecated with the introduction of multiple flows. Please use the `flowStartingPoints` field.
	// Deprecated
	PrototypeStartNodeID string `json:"prototypeStartNodeID"`
	// An array of flow starting points sorted by its position in the prototype settings panel.
	FlowStartingPoints []FlowStartingPoint `json:"flowStartingPoints"`
	// The device used to view a prototype.
	PrototypeDevice PrototypeDevice `json:"prototypeDevice"`
	Measurements    []Measurement   `json:"measurements,omitempty"`
}
