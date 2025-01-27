package models

import (
	"time"
)

// FileBranchesInner struct for FileBranchesInner
type FileBranchesInner struct {
	// The name of the branch.
	Name string `json:"name"`
	// The key of the branch.
	Key string `json:"key"`
	// A URL to a thumbnail image of the branch.
	ThumbnailUrl string `json:"thumbnail_url"`
	// The UTC ISO 8601 time at which the branch was last modified.
	LastModified time.Time `json:"last_modified"`
}
