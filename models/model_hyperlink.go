package models

// Hyperlink A link to either a URL or another frame (node) in the document.
type Hyperlink struct {
	// The URL that the hyperlink points to, if `type` is `URL`.
	Url string `json:"url,omitempty"`
	// The type of hyperlink. Can be either `URL` or `NODE`.
	Type string `json:"type"`
	// The ID of the node that the hyperlink points to, if `type` is `NODE`.
	NodeID string `json:"nodeID,omitempty"`
}
