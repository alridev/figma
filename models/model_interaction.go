package models

// Interaction An interaction in the Figma viewer, containing a trigger and one or more actions.
type Interaction struct {
	Trigger Trigger `json:"trigger"`
	// The actions that are performed when the trigger is activated.
	Actions []Action `json:"actions,omitempty"`
}
