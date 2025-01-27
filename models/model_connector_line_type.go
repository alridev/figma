package models

// ConnectorLineType Connector line type.
type ConnectorLineType string

// List of ConnectorLineType
const (
	CONNECTORLINETYPE_STRAIGHT ConnectorLineType = "STRAIGHT"
	CONNECTORLINETYPE_ELBOWED  ConnectorLineType = "ELBOWED"
)

// All allowed values of ConnectorLineType enum
var AllowedConnectorLineTypeEnumValues = []ConnectorLineType{
	"STRAIGHT",
	"ELBOWED",
}
