package models

// Expression Defines the [Expression](https://help.figma.com/hc/en-us/articles/15253194385943) object, which contains a list of `VariableData` objects strung together by operators (`ExpressionFunction`).
type Expression struct {
	ExpressionFunction  ExpressionFunction `json:"expressionFunction"`
	ExpressionArguments []VariableData     `json:"expressionArguments"`
}
