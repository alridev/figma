package models

// BlendMode This type is a string enum with the following possible values  Normal blends: - `PASS_THROUGH` (only applicable to objects with children) - `NORMAL`  Darken: - `DARKEN` - `MULTIPLY` - `LINEAR_BURN` - `COLOR_BURN`  Lighten: - `LIGHTEN` - `SCREEN` - `LINEAR_DODGE` - `COLOR_DODGE`  Contrast: - `OVERLAY` - `SOFT_LIGHT` - `HARD_LIGHT`  Inversion: - `DIFFERENCE` - `EXCLUSION`  Component: - `HUE` - `SATURATION` - `COLOR` - `LUMINOSITY`
type BlendMode string

// List of BlendMode
const (
	BLENDMODE_PASS_THROUGH BlendMode = "PASS_THROUGH"
	BLENDMODE_NORMAL       BlendMode = "NORMAL"
	BLENDMODE_DARKEN       BlendMode = "DARKEN"
	BLENDMODE_MULTIPLY     BlendMode = "MULTIPLY"
	BLENDMODE_LINEAR_BURN  BlendMode = "LINEAR_BURN"
	BLENDMODE_COLOR_BURN   BlendMode = "COLOR_BURN"
	BLENDMODE_LIGHTEN      BlendMode = "LIGHTEN"
	BLENDMODE_SCREEN       BlendMode = "SCREEN"
	BLENDMODE_LINEAR_DODGE BlendMode = "LINEAR_DODGE"
	BLENDMODE_COLOR_DODGE  BlendMode = "COLOR_DODGE"
	BLENDMODE_OVERLAY      BlendMode = "OVERLAY"
	BLENDMODE_SOFT_LIGHT   BlendMode = "SOFT_LIGHT"
	BLENDMODE_HARD_LIGHT   BlendMode = "HARD_LIGHT"
	BLENDMODE_DIFFERENCE   BlendMode = "DIFFERENCE"
	BLENDMODE_EXCLUSION    BlendMode = "EXCLUSION"
	BLENDMODE_HUE          BlendMode = "HUE"
	BLENDMODE_SATURATION   BlendMode = "SATURATION"
	BLENDMODE_COLOR        BlendMode = "COLOR"
	BLENDMODE_LUMINOSITY   BlendMode = "LUMINOSITY"
)

// All allowed values of BlendMode enum
var AllowedBlendModeEnumValues = []BlendMode{
	"PASS_THROUGH",
	"NORMAL",
	"DARKEN",
	"MULTIPLY",
	"LINEAR_BURN",
	"COLOR_BURN",
	"LIGHTEN",
	"SCREEN",
	"LINEAR_DODGE",
	"COLOR_DODGE",
	"OVERLAY",
	"SOFT_LIGHT",
	"HARD_LIGHT",
	"DIFFERENCE",
	"EXCLUSION",
	"HUE",
	"SATURATION",
	"COLOR",
	"LUMINOSITY",
}
