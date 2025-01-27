package models

// ArcData Information about the arc properties of an ellipse. 0° is the x axis and increasing angles rotate clockwise.
type ArcData struct {
	// Start of the sweep in radians.
	StartingAngle float64 `json:"startingAngle"`
	// End of the sweep in radians.
	EndingAngle float64 `json:"endingAngle"`
	// Inner radius value between 0 and 1
	InnerRadius float64 `json:"innerRadius"`
}
