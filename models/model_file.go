package models

import (
	"time"
)

// File struct for File
type File struct {
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
	ThumbnailUrl string      `json:"thumbnailUrl,omitempty"`
	Document     DocumentNode `json:"document"`
	// A mapping from component IDs to component metadata.
	Components map[string]Component `json:"components"`
	// A mapping from component set IDs to component set metadata.
	ComponentSets map[string]ComponentSet `json:"componentSets"`
	// The version of the file schema that this file uses.
	SchemaVersion float64 `json:"schemaVersion"`
	// A mapping from style IDs to style metadata.
	Styles map[string]Style `json:"styles"`
	// The share permission level of the file link.
	LinkAccess string `json:"linkAccess,omitempty"`
	// The key of the main file for this file. If present, this file is a component or component set.
	MainFileKey string `json:"mainFileKey,omitempty"`
	// A list of branches for this file.
	Branches []FileBranchesInner `json:"branches,omitempty"`
}


