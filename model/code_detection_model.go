package model

// CodeInspectionModel -
type CodeInspectionModel struct {
	ProjectName   string `json:"project_name"`
	RepositoryURL string `json:"source_address"`
	Branch        string `json:"branch"`
	User          string `json:"user"`
	Password      string `json:"password"`
}

type NormativeInspectionModel struct {
	ExtendMethod string `json:"extend_method"`
	ComponentID  string `json:"component_id"`
}

type CodeIssues struct {
	Key       string `json:"key"`
	Project   string `json:"project"`
	Severity  string `json:"severity"`
	Message   string `json:"message"`
	Component string `json:"component"`
}
