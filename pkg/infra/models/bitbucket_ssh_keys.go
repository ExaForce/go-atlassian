package models

// SSHKeyPageScheme represents a paginated list of SSH keys for a Bitbucket user.
type SSHKeyPageScheme struct {
	Size     int             `json:"size,omitempty"`     // The number of SSH keys in the current page.
	Page     int             `json:"page,omitempty"`     // The current page number.
	Pagelen  int             `json:"pagelen,omitempty"`  // The total number of pages.
	Next     string          `json:"next,omitempty"`     // The URL to the next page.
	Previous string          `json:"previous,omitempty"` // The URL to the previous page.
	Values   []*SSHKeyScheme `json:"values,omitempty"`   // The SSH keys in the current page.
}

// SSHKeyScheme represents an SSH key registered on a Bitbucket user account.
type SSHKeyScheme struct {
	Type      string                  `json:"type,omitempty"`       // The resource type.
	UUID      string                  `json:"uuid,omitempty"`       // The unique identifier of the SSH key.
	Key       string                  `json:"key,omitempty"`        // The public SSH key material.
	Comment   string                  `json:"comment,omitempty"`    // The comment parsed from the SSH key, if any.
	Label     string                  `json:"label,omitempty"`      // A user-supplied label for the SSH key.
	CreatedOn string                  `json:"created_on,omitempty"` // The time the SSH key was registered.
	LastUsed  string                  `json:"last_used,omitempty"`  // The last time the SSH key was used.
	Owner     *BitbucketAccountScheme `json:"owner,omitempty"`      // The owner of the SSH key.
	Links     *SSHKeyLinksScheme      `json:"links,omitempty"`      // The links related to the SSH key.
}

// SSHKeyLinksScheme represents a collection of links related to an SSH key.
type SSHKeyLinksScheme struct {
	Self *BitbucketLinkScheme `json:"self,omitempty"` // The link to the SSH key itself.
}
