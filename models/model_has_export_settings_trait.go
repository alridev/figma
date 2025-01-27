package models

// HasExportSettingsTrait struct for HasExportSettingsTrait
type HasExportSettingsTrait struct {
	// An array of export settings representing images to export from the node.
	ExportSettings []ExportSetting `json:"exportSettings,omitempty"`
}
