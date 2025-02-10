package models

type DeployKeysPageScheme struct {
	Page     int                `json:"page"`
	Pagelen  int                `json:"pagelen"`
	Next     string             `json:"next"`
	Previous string             `json:"previous"`
	Values   []*DeployKeyScheme `json:"values"`
}

type DeployKeyScheme struct {
	Id         int                     `json:"id"`
	Key        string                  `json:"key"`
	Repository *RepositoryScheme       `json:"repository"`
	Comment    string                  `json:"comment"`
	Label      string                  `json:"label"`
	Type       string                  `json:"type"`
	CreatedOn  string                  `json:"created_on"`
	LastUsed   string                  `json:"last_used"`
	Owner      *BitbucketAccountScheme `json:"owner"`
}
