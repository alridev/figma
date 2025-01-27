package models

// VariableDataType Defines the types of data a VariableData object can hold
type VariableDataType string

// List of VariableDataType
const (
	VARIABLEDATATYPE_BOOLEAN        VariableDataType = "BOOLEAN"
	VARIABLEDATATYPE_FLOAT          VariableDataType = "FLOAT"
	VARIABLEDATATYPE_STRING         VariableDataType = "STRING"
	VARIABLEDATATYPE_COLOR          VariableDataType = "COLOR"
	VARIABLEDATATYPE_VARIABLE_ALIAS VariableDataType = "VARIABLE_ALIAS"
	VARIABLEDATATYPE_EXPRESSION     VariableDataType = "EXPRESSION"
)

// All allowed values of VariableDataType enum
var AllowedVariableDataTypeEnumValues = []VariableDataType{
	"BOOLEAN",
	"FLOAT",
	"STRING",
	"COLOR",
	"VARIABLE_ALIAS",
	"EXPRESSION",
}
