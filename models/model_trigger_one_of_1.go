package models

// TriggerOneOf1 struct for TriggerOneOf1
type TriggerOneOf1 struct {
	Type  string  `json:"type"`
	Delay float64 `json:"delay"`
	// Whether this is a [deprecated version](https://help.figma.com/hc/en-us/articles/360040035834-Prototype-triggers#h_01HHN04REHJNP168R26P1CMP0A) of the trigger that was left unchanged for backwards compatibility. If not present, the trigger is the latest version.
	DeprecatedVersion bool `json:"deprecatedVersion,omitempty"`
}
