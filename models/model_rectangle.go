package models

// Rectangle A rectangle that expresses a bounding box in absolute coordinates.
type Rectangle struct {
	// X coordinate of top left corner of the rectangle.
	X float64 `json:"x"`
	// Y coordinate of top left corner of the rectangle.
	Y float64 `json:"y"`
	// Width of the rectangle.
	Width float64 `json:"width"`
	// Height of the rectangle.
	Height float64 `json:"height"`
}
