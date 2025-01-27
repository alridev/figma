package models

// SetVariableModeAction Sets a variable to a specific mode.
type SetVariableModeAction struct {
	Type                 string `json:"type"`
	VariableCollectionId string `json:"variableCollectionId,omitempty"`
	VariableModeId       string `json:"variableModeId,omitempty"`
}
