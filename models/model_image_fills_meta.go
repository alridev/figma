package models

// ImageFillsMeta struct for ImageFillsMeta
type ImageFillsMeta struct {
	// A map of image references to URLs of the image fills.
	Images map[string]string `json:"images"`
}
