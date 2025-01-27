package models

// StyleType The type of style
type StyleType string

// List of StyleType
const (
	STYLETYPE_FILL   StyleType = "FILL"
	STYLETYPE_TEXT   StyleType = "TEXT"
	STYLETYPE_EFFECT StyleType = "EFFECT"
	STYLETYPE_GRID   StyleType = "GRID"
)

// All allowed values of StyleType enum
var AllowedStyleTypeEnumValues = []StyleType{
	"FILL",
	"TEXT",
	"EFFECT",
	"GRID",
}
