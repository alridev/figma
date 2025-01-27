package models

// OpenURLAction An action that opens a URL.
type OpenURLAction struct {
	Url  string `json:"url"`
	Type string `json:"type"`
}
