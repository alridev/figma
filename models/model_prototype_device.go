package models

// PrototypeDevice The device used to view a prototype.
type PrototypeDevice struct {
	Type             string  `json:"type"`
	Size             *Size   `json:"size,omitempty"`
	PresetIdentifier string `json:"presetIdentifier,omitempty"`
	Rotation         string  `json:"rotation"`
}
