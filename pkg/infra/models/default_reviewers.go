package models

// DefaultReviewersPageScheme represents a paginated list of default reviewers
type DefaultReviewersPageScheme struct {
	Size     int                       `json:"size"`
	Page     int                       `json:"page"`
	Pagelen  int                       `json:"pagelen"`
	Next     string                    `json:"next"`
	Previous string                    `json:"previous"`
	Values   []*BitbucketAccountScheme `json:"values"`
}
