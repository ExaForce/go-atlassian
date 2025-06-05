package models

type JiraUserProjectPermissionScheme struct {
	ProjectID  string `json:"projectId,omitempty"`
	AccountID  string `json:"accountId,omitempty"`
	Permission string `json:"permission,omitempty"`
}
