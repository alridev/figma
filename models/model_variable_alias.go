package models

// VariableAlias Contains a variable alias
type VariableAlias struct {
	Type string `json:"type"`
	// The id of the variable that the current variable is aliased to. This variable can be a local or remote variable, and both can be retrieved via the GET /v1/files/:file_key/variables/local endpoint.
	Id string `json:"id"`
}
