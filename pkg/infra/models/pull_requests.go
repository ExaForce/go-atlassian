package models

// PullRequestsResponse represents the response for a list of pull requests
type PullRequestsResponse struct {
	Size     int                  `json:"size"`
	Page     int                  `json:"page"`
	Pagelen  int                  `json:"pagelen"`
	Next     string               `json:"next"`
	Previous string               `json:"previous"`
	Values   []*PullRequestScheme `json:"values"`
}

// PullRequestScheme represents a pull request
type PullRequestScheme struct {
	CommentCount      int                     `json:"comment_count"`
	TaskCount         int                     `json:"task_count"`
	Type              string                  `json:"type"`
	ID                int                     `json:"id"`
	Title             string                  `json:"title"`
	Description       string                  `json:"description"`
	State             string                  `json:"state"`
	MergeCommit       *CommitScheme           `json:"merge_commit"`
	CloseSourceBranch bool                    `json:"close_source_branch"`
	ClosedBy          *BitbucketAccountScheme `json:"closed_by"`
	Author            *BitbucketAccountScheme `json:"author"`
	Reason            string                  `json:"reason"`
	CreatedOn         string                  `json:"created_on"`
	UpdatedOn         string                  `json:"updated_on"`
	Source            *PullRequestRefScheme   `json:"source"`
	Destination       *PullRequestRefScheme   `json:"destination"`
	Links             *PullRequestLinksScheme `json:"links"`
	Summary           *SummaryScheme          `json:"summary"`
}

// PullRequestRefScheme represents a reference (source/destination) in a pull request
type PullRequestRefScheme struct {
	Branch     *PullRequestBranchScheme `json:"branch"`
	Commit     *CommitScheme            `json:"commit"`
	Repository *RepositoryScheme        `json:"repository"`
}

// CommitScheme represents a commit in a pull request
type CommitScheme struct {
	Hash    string                `json:"hash"`
	Type    string                `json:"type"`
	Message string                `json:"message"`
	Date    string                `json:"date"`
	Links   *WorkspaceLinksScheme `json:"links"`
}

// PullRequestLinksScheme represents the links available in a pull request
type PullRequestLinksScheme struct {
	Self           *BitbucketLinkScheme `json:"self"`
	HTML           *BitbucketLinkScheme `json:"html"`
	Commits        *BitbucketLinkScheme `json:"commits"`
	Approve        *BitbucketLinkScheme `json:"approve"`
	RequestChanges *BitbucketLinkScheme `json:"request-changes"`
	Diff           *BitbucketLinkScheme `json:"diff"`
	Diffstat       *BitbucketLinkScheme `json:"diffstat"`
	Comments       *BitbucketLinkScheme `json:"comments"`
	Activity       *BitbucketLinkScheme `json:"activity"`
	Merge          *BitbucketLinkScheme `json:"merge"`
	Decline        *BitbucketLinkScheme `json:"decline"`
	Statuses       *BitbucketLinkScheme `json:"statuses"`
}

// SummaryScheme represents the summary of a pull request
type SummaryScheme struct {
	Type   string `json:"type"`
	Raw    string `json:"raw"`
	Markup string `json:"markup"`
	HTML   string `json:"html"`
}
