package models

// UpdateMediaRuntimeActionOneOf An action that updates the runtime of a media node by playing, pausing, toggling play/pause, muting, unmuting, or toggling mute/unmute.  The `destinationId` is the node ID of the media node to update. If `destinationId` is `null`, the action will update the media node that contains the action.  The `mediaAction` is the action to perform on the media node.
type UpdateMediaRuntimeActionOneOf struct {
	Type          string `json:"type"`
	DestinationId string `json:"destinationId"`
	MediaAction   string `json:"mediaAction"`
}
