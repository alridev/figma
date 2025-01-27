package models

// ColorStopBoundVariables The variables bound to a particular gradient stop
type ColorStopBoundVariables struct {
	Color *VariableAlias `json:"color,omitempty"`
}
