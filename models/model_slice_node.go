package models

// SliceNode struct for SliceNode
type SliceNode struct {
	// The name given to the node by the user in the tool.
	Name string `json:"name"`
	// The type of this node, represented by the string literal \"SLICE\"
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
}
