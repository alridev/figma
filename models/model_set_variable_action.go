package models

// SetVariableAction Sets a variable to a specific value.
type SetVariableAction struct {
	Type          string        `json:"type"`
	VariableId    string        `json:"variableId"`
	VariableValue *VariableData `json:"variableValue,omitempty"`
}
