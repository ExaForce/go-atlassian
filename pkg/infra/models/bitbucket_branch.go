package models

// BranchScheme represents a branch in a repository.
type BranchScheme struct {
	MergeStrategies      []string `json:"merge_strategies"`       // The merge strategies available for the branch.
	DefaultMergeStrategy string   `json:"default_merge_strategy"` // The default merge strategy used for the branch.
}

// MainBranchScheme represents the main branch of a repository.
type MainBranchScheme struct {
	Name string `json:"name"` // The name of the main branch.
	Type string `json:"type"` // The type of the main branch.
}

type OverrideSettingsScheme struct {
	DefaultMergeStrategy bool `json:"default_merge_strategy"` // Whether to use the default merge strategy.
	BranchingModel       bool `json:"branching_model"`        // Whether to use the branching model.
}

type PullRequestBranchScheme struct {
	Name  string                `json:"name"`
	Links *WorkspaceLinksScheme `json:"links"`
}
