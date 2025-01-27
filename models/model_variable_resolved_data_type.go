package models

// VariableResolvedDataType Defines the types of data a VariableData object can eventually equal
type VariableResolvedDataType string

// List of VariableResolvedDataType
const (
	VARIABLERESOLVEDDATATYPE_BOOLEAN VariableResolvedDataType = "BOOLEAN"
	VARIABLERESOLVEDDATATYPE_FLOAT   VariableResolvedDataType = "FLOAT"
	VARIABLERESOLVEDDATATYPE_STRING  VariableResolvedDataType = "STRING"
	VARIABLERESOLVEDDATATYPE_COLOR   VariableResolvedDataType = "COLOR"
)

// All allowed values of VariableResolvedDataType enum
var AllowedVariableResolvedDataTypeEnumValues = []VariableResolvedDataType{
	"BOOLEAN",
	"FLOAT",
	"STRING",
	"COLOR",
}
