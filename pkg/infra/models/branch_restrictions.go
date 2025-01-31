package models

// BranchRestrictionsPageScheme represents a paginated list of branch restrictions
type BranchRestrictionsPageScheme struct {
	Size     int                        `json:"size"`
	Page     int                        `json:"page"`
	Pagelen  int                        `json:"pagelen"`
	Next     string                     `json:"next"`
	Previous string                     `json:"previous"`
	Values   []*BranchRestrictionScheme `json:"values"`
}

// BranchRestrictionScheme represents a branch restriction rule
type BranchRestrictionScheme struct {
	Type            string                        `json:"type,omitempty"`
	Links           *BranchRestrictionLinksScheme `json:"links,omitempty"`
	ID              int                           `json:"id,omitempty"`
	Kind            string                        `json:"kind,omitempty"`
	BranchMatchKind string                        `json:"branch_match_kind,omitempty"`
	BranchType      string                        `json:"branch_type,omitempty"`
	Pattern         string                        `json:"pattern,omitempty"`
	Value           int                           `json:"value,omitempty"`
	Users           []*BitbucketAccountScheme     `json:"users,omitempty"`
	Groups          []*BitbucketGroupScheme       `json:"groups,omitempty"`
}

// BranchRestrictionLinksScheme represents the links related to a branch restriction
type BranchRestrictionLinksScheme struct {
	Self *BitbucketLinkScheme `json:"self,omitempty"`
}
