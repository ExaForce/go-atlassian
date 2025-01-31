package models

// BitbucketGroupScheme represents a Bitbucket group
type BitbucketGroupScheme struct {
	Type     string               `json:"type,omitempty"`
	Name     string               `json:"name,omitempty"`
	Slug     string               `json:"slug,omitempty"`
	FullSlug string               `json:"full_slug,omitempty"`
	Owner    *WorkspaceScheme     `json:"owner,omitempty"`
	Links    *BitbucketLinkScheme `json:"links,omitempty"`
}
