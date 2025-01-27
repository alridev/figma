package models

// Constraint Sizing constraint for exports.
type Constraint struct {
	// Type of constraint to apply:  - `SCALE`: Scale by `value`. - `WIDTH`: Scale proportionally and set width to `value`. - `HEIGHT`: Scale proportionally and set height to `value`.
	Type string `json:"type"`
	// See type property for effect of this field.
	Value float64 `json:"value"`
}
