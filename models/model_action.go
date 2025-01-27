package models

// Action - An action that is performed when a trigger is activated.
type Action struct {
	ActionOneOf              *ActionOneOf
	ConditionalAction        *ConditionalAction
	NodeAction               *NodeAction
	OpenURLAction            *OpenURLAction
	SetVariableAction        *SetVariableAction
	SetVariableModeAction    *SetVariableModeAction
	UpdateMediaRuntimeAction *UpdateMediaRuntimeAction
}
