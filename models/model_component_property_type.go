package models

// ComponentPropertyType Component property type.
type ComponentPropertyType string

// List of ComponentPropertyType
const (
	COMPONENTPROPERTYTYPE_BOOLEAN       ComponentPropertyType = "BOOLEAN"
	COMPONENTPROPERTYTYPE_INSTANCE_SWAP ComponentPropertyType = "INSTANCE_SWAP"
	COMPONENTPROPERTYTYPE_TEXT          ComponentPropertyType = "TEXT"
	COMPONENTPROPERTYTYPE_VARIANT       ComponentPropertyType = "VARIANT"
)

// All allowed values of ComponentPropertyType enum
var AllowedComponentPropertyTypeEnumValues = []ComponentPropertyType{
	"BOOLEAN",
	"INSTANCE_SWAP",
	"TEXT",
	"VARIANT",
}
