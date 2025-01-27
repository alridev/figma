package models

// ExpressionFunction Defines the list of operators available to use in an Expression.
type ExpressionFunction string

// List of ExpressionFunction
const (
	EXPRESSIONFUNCTION_ADDITION              ExpressionFunction = "ADDITION"
	EXPRESSIONFUNCTION_SUBTRACTION           ExpressionFunction = "SUBTRACTION"
	EXPRESSIONFUNCTION_MULTIPLICATION        ExpressionFunction = "MULTIPLICATION"
	EXPRESSIONFUNCTION_DIVISION              ExpressionFunction = "DIVISION"
	EXPRESSIONFUNCTION_EQUALS                ExpressionFunction = "EQUALS"
	EXPRESSIONFUNCTION_NOT_EQUAL             ExpressionFunction = "NOT_EQUAL"
	EXPRESSIONFUNCTION_LESS_THAN             ExpressionFunction = "LESS_THAN"
	EXPRESSIONFUNCTION_LESS_THAN_OR_EQUAL    ExpressionFunction = "LESS_THAN_OR_EQUAL"
	EXPRESSIONFUNCTION_GREATER_THAN          ExpressionFunction = "GREATER_THAN"
	EXPRESSIONFUNCTION_GREATER_THAN_OR_EQUAL ExpressionFunction = "GREATER_THAN_OR_EQUAL"
	EXPRESSIONFUNCTION_AND                   ExpressionFunction = "AND"
	EXPRESSIONFUNCTION_OR                    ExpressionFunction = "OR"
	EXPRESSIONFUNCTION_VAR_MODE_LOOKUP       ExpressionFunction = "VAR_MODE_LOOKUP"
	EXPRESSIONFUNCTION_NEGATE                ExpressionFunction = "NEGATE"
	EXPRESSIONFUNCTION_NOT                   ExpressionFunction = "NOT"
)

// All allowed values of ExpressionFunction enum
var AllowedExpressionFunctionEnumValues = []ExpressionFunction{
	"ADDITION",
	"SUBTRACTION",
	"MULTIPLICATION",
	"DIVISION",
	"EQUALS",
	"NOT_EQUAL",
	"LESS_THAN",
	"LESS_THAN_OR_EQUAL",
	"GREATER_THAN",
	"GREATER_THAN_OR_EQUAL",
	"AND",
	"OR",
	"VAR_MODE_LOOKUP",
	"NEGATE",
	"NOT",
}
