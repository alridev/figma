package models

// ConditionalAction Checks if a condition is met before performing certain actions by using an if/else conditional statement.
type ConditionalAction struct {
	Type              string             `json:"type"`
	ConditionalBlocks []ConditionalBlock `json:"conditionalBlocks"`
}
