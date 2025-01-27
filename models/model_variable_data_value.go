package models

// VariableDataValue - struct for VariableDataValue
type VariableDataValue struct {
	Expression    *Expression
	RGB           *RGB
	RGBA          *RGBA
	VariableAlias *VariableAlias
	Bool          bool
	Float32       float64
	String        string
}
