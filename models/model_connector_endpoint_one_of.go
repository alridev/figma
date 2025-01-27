package models

// ConnectorEndpointOneOf struct for ConnectorEndpointOneOf
type ConnectorEndpointOneOf struct {
	// Node ID that this endpoint attaches to.
	EndpointNodeId string `json:"endpointNodeId,omitempty"`
	// The position of the endpoint relative to the node.
	Position *Vector `json:"position,omitempty"`
}
