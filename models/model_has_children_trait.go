package models

// HasChildrenTrait struct for HasChildrenTrait
type HasChildrenTrait struct {
	// An array of nodes that are direct children of this node
	Children []Node `json:"children"`
}
