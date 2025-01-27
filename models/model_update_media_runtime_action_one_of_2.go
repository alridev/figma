package models

// UpdateMediaRuntimeActionOneOf2 An action that updates the runtime of a media node by skipping to a specific time.  The `destinationId` is the node ID of the media node to update. If `destinationId` is `null`, the action will update the media node that contains the action.  The `mediaAction` is the action to perform on the media node.  The `newTimestamp` is the new time to skip to in seconds.
type UpdateMediaRuntimeActionOneOf2 struct {
	Type          string  `json:"type"`
	DestinationId string  `json:"destinationId,omitempty"`
	MediaAction   string  `json:"mediaAction"`
	NewTimestamp  float64 `json:"newTimestamp"`
}
