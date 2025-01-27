package models

// ConnectorEndpointOneOf1 struct for ConnectorEndpointOneOf1
type ConnectorEndpointOneOf1 struct {
	// Node ID that this endpoint attaches to.
	EndpointNodeId string `json:"endpointNodeId,omitempty"`
	// The magnet type is a string enum.
	Magnet string `json:"magnet,omitempty"`
}
