package models

type RepositoryPipelineVariablesPageScheme struct {
	Size     int                           `json:"size"`
	Page     int                           `json:"page"`
	Pagelen  int                           `json:"pagelen"`
	Next     string                        `json:"next"`
	Previous string                        `json:"previous"`
	Values   []*RepositoryPipelineVariable `json:"values"`
}

type RepositoryPipelineVariable struct {
	Type    string `json:"type"`
	UUID    string `json:"uuid"`
	Key     string `json:"key"`
	Secured bool   `json:"secured"`
	System  bool   `json:"system"`
	Scope   string `json:"scope"`
}
