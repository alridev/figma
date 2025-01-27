package models

// Navigation The method of navigation. The possible values are: - `\"NAVIGATE\"`: Replaces the current screen with the destination, also closing all overlays. - `\"OVERLAY\"`: Opens the destination as an overlay on the current screen. - `\"SWAP\"`: On an overlay, replaces the current (topmost) overlay with the destination. On a top-level frame,   behaves the same as `\"NAVIGATE\"` except that no entry is added to the navigation history. - `\"SCROLL_TO\"`: Scrolls to the destination on the current screen. - `\"CHANGE_TO\"`: Changes the closest ancestor instance of source node to the specified variant.
type Navigation string

// List of Navigation
const (
	NAVIGATION_NAVIGATE  Navigation = "NAVIGATE"
	NAVIGATION_SWAP      Navigation = "SWAP"
	NAVIGATION_OVERLAY   Navigation = "OVERLAY"
	NAVIGATION_SCROLL_TO Navigation = "SCROLL_TO"
	NAVIGATION_CHANGE_TO Navigation = "CHANGE_TO"
)

// All allowed values of Navigation enum
var AllowedNavigationEnumValues = []Navigation{
	"NAVIGATE",
	"SWAP",
	"OVERLAY",
	"SCROLL_TO",
	"CHANGE_TO",
}
