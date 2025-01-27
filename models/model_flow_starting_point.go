package models

// FlowStartingPoint A flow starting point used when launching a prototype to enter Presentation view.
type FlowStartingPoint struct {
	// Name of flow.
	Name string `json:"name"`
	// Unique identifier specifying the frame.
	NodeId string `json:"nodeId"`
}
