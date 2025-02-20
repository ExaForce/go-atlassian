package models

// ProjectMembershipPageScheme represents a paginated list of project memberships.
type ProjectUserMembershipPageScheme struct {
	Size     int                            `json:"size,omitempty"`     // The number of memberships in the current page.
	Page     int                            `json:"page,omitempty"`     // The current page number.
	Pagelen  int                            `json:"pagelen,omitempty"`  // The total number of pages.
	Next     string                         `json:"next,omitempty"`     // The URL to the next page.
	Previous string                         `json:"previous,omitempty"` // The URL to the previous page.
	Values   []*ProjectUserMembershipScheme `json:"values,omitempty"`   // The project memberships in the current page.
}

// ProjectMembershipPageScheme represents a paginated list of project memberships.
type ProjectGroupMembershipPageScheme struct {
	Size     int                             `json:"size,omitempty"`     // The number of memberships in the current page.
	Page     int                             `json:"page,omitempty"`     // The current page number.
	Pagelen  int                             `json:"pagelen,omitempty"`  // The total number of pages.
	Next     string                          `json:"next,omitempty"`     // The URL to the next page.
	Previous string                          `json:"previous,omitempty"` // The URL to the previous page.
	Values   []*ProjectGroupMembershipScheme `json:"values,omitempty"`   // The project memberships in the current page.
}

type ProjectGroupMembershipScheme struct {
	Type       string                  `json:"type,omitempty"`       // The type of the membership.
	Permission string                  `json:"permission,omitempty"` // The permission of the user in the project.
	Group      *BitbucketGroupScheme   `json:"group,omitempty"`      // The group who has the membership.
	Project    *BitbucketProjectScheme `json:"project,omitempty"`    // The project that the user has the membership in.
}

type ProjectUserMembershipScheme struct {
	Type       string                  `json:"type,omitempty"`       // The type of the membership.
	Permission string                  `json:"permission,omitempty"` // The permission of the user in the project.
	User       *BitbucketAccountScheme `json:"user,omitempty"`       // The user who has the membership.
	Project    *BitbucketProjectScheme `json:"project,omitempty"`    // The project that the user has the membership in.
}
