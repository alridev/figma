package models

// ExportSetting An export setting.
type ExportSetting struct {
	Suffix     string     `json:"suffix"`
	Format     string     `json:"format"`
	Constraint Constraint `json:"constraint"`
}
