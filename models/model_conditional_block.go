package models

// ConditionalBlock Either the if or else conditional blocks. The if block contains a condition to check. If that condition is met then it will run those list of actions, else it will run the actions in the else block.
type ConditionalBlock struct {
	Condition *VariableData `json:"condition,omitempty"`
	Actions   []Action      `json:"actions"`
}
