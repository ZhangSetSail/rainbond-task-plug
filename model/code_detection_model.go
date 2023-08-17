package model

// CodeDetectionModel -
type CodeDetectionModel struct {
	ProjectName   string `json:"project_name"`
	RepositoryURL string `json:"source_address"`
	Branch        string `json:"branch"`
	User          string `json:"user"`
	Password      string `json:"password"`
}
