package models

type RepositoryGroupPermissionsPageScheme struct {
	Size     int                                 `json:"size"`
	Page     int                                 `json:"page"`
	Pagelen  int                                 `json:"pagelen"`
	Next     string                              `json:"next"`
	Previous string                              `json:"previous"`
	Values   []*RepositoryGroupPermissionsScheme `json:"values"`
}

type RepositoryGroupPermissionsScheme struct {
	Type       string                `json:"type"`
	Permission string                `json:"permission"`
	Group      *BitbucketGroupScheme `json:"group"`
	Repository *RepositoryScheme     `json:"repository"`
}
