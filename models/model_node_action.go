package models

// NodeAction An action that navigates to a specific node in the Figma viewer.
type NodeAction struct {
	Type          string     `json:"type"`
	DestinationId string     `json:"destinationId"`
	Navigation    Navigation `json:"navigation"`
	Transition    Transition `json:"transition"`
	// Whether the scroll offsets of any scrollable elements in the current screen or overlay are preserved when navigating to the destination. This is applicable only if the layout of both the current frame and its destination are the same.
	PreserveScrollPosition bool `json:"preserveScrollPosition,omitempty"`
	// Applicable only when `navigation` is `\"OVERLAY\"` and the destination is a frame with `overlayPosition` equal to `\"MANUAL\"`. This value represents the offset by which the overlay is opened relative to this node.
	OverlayRelativePosition *Vector `json:"overlayRelativePosition,omitempty"`
	// When true, all videos within the destination frame will reset their memorized playback position to 00:00 before starting to play.
	ResetVideoPosition bool `json:"resetVideoPosition,omitempty"`
	// Whether the scroll offsets of any scrollable elements in the current screen or overlay reset when navigating to the destination. This is applicable only if the layout of both the current frame and its destination are the same.
	ResetScrollPosition bool `json:"resetScrollPosition,omitempty"`
	// Whether the state of any interactive components in the current screen or overlay reset when navigating to the destination. This is applicable if there are interactive components in the destination frame.
	ResetInteractiveComponents bool `json:"resetInteractiveComponents,omitempty"`
}
