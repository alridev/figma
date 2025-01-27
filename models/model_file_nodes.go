package models

import (
	"time"
)

// FileNodes struct for FileNodes
type FileNodes struct {
	// The version number of the file. This number is incremented when a file is modified and can be used to check if the file has changed between requests.
	Version string `json:"version"`
	// The name of the file as it appears in the editor.
	Name string `json:"name"`
	// The role of the user making the API request in relation to the file.
	Role string `json:"role"`
	// The UTC ISO 8601 time at which the file was last modified.
	LastModified time.Time `json:"lastModified"`
	// The type of editor associated with this file.
	EditorType string `json:"editorType"`
	// A URL to a thumbnail image of the file.
	ThumbnailUrl string `json:"thumbnailUrl"`
	// A mapping from node IDs to node metadata.
	Nodes map[string]FileNodesNodesValue `json:"nodes"`
}
