package models

// UpdateMediaRuntimeActionOneOf1 An action that updates the runtime of a media node by skipping forward or backward.  The `destinationId` is the node ID of the media node to update. If `destinationId` is `null`, the action will update the media node that contains the action.  The `mediaAction` is the action to perform on the media node.  The `amountToSkip` is the amount of time to skip in seconds.
type UpdateMediaRuntimeActionOneOf1 struct {
	Type          string  `json:"type"`
	DestinationId string  `json:"destinationId,omitempty"`
	MediaAction   string  `json:"mediaAction"`
	AmountToSkip  float64 `json:"amountToSkip"`
}
