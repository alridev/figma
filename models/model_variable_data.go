package models

// VariableData A value to set a variable to during prototyping.
type VariableData struct {
	Type         *VariableDataType         `json:"type,omitempty"`
	ResolvedType *VariableResolvedDataType `json:"resolvedType,omitempty"`
	Value        *VariableDataValue        `json:"value,omitempty"`
}
