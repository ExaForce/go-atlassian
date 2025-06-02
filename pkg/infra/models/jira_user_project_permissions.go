package models

type JiraUserProjectPermissionScheme struct {
	ProjectKey string `json:"projectKey,omitempty"`
	AccountID  string `json:"accountId,omitempty"`
	Permission string `json:"permission,omitempty"`
}
