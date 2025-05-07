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

type RepositoryPipelineRunsPageScheme struct {
	Size     int                      `json:"size"`
	Page     int                      `json:"page"`
	Pagelen  int                      `json:"pagelen"`
	Next     string                   `json:"next"`
	Previous string                   `json:"previous"`
	Values   []*RepositoryPipelineRun `json:"values"`
}

type RepositoryPipelineRun struct {
	Type            string              `json:"type"`
	UUID            string              `json:"uuid"`
	BuildNumber     int                 `json:"build_number"`
	Creator         *PipelineCreator    `json:"creator"`
	Repository      *PipelineRepository `json:"repository"`
	State           *PipelineState      `json:"state"`
	CreatedOn       string              `json:"created_on"`
	CompletedOn     string              `json:"completed_on"`
	Target          *PipelineTarget     `json:"target"`
	Trigger         *PipelineTrigger    `json:"trigger"`
	RunNumber       int                 `json:"run_number"`
	RunCreationDate string              `json:"run_creation_date"`
	DurationInSecs  int                 `json:"duration_in_seconds"`
	BuildSecsUsed   int                 `json:"build_seconds_used"`
	FirstSuccessful bool                `json:"first_successful"`
	Expired         bool                `json:"expired"`
	HasVariables    bool                `json:"has_variables"`
	Labels          map[string]string   `json:"labels"`
}

type RepositoryPipelineRunStepsPageScheme struct {
	Size     int                          `json:"size"`
	Page     int                          `json:"page"`
	Pagelen  int                          `json:"pagelen"`
	Next     string                       `json:"next"`
	Previous string                       `json:"previous"`
	Values   []*RepositoryPipelineRunStep `json:"values"`
}

type RepositoryPipelineRunStep struct {
	Type              string                 `json:"type"`
	UUID              string                 `json:"uuid"`
	StartedOn         string                 `json:"started_on"`
	CompletedOn       string                 `json:"completed_on"`
	State             *PipelineStepState     `json:"state"`
	Image             map[string]interface{} `json:"image"`
	SetupCommands     []*PipelineCommand     `json:"setup_commands"`
	ScriptCommands    []*PipelineCommand     `json:"script_commands"`
	TeardownCommands  []*PipelineCommand     `json:"teardown_commands"`
	MaxTime           int                    `json:"maxTime"`
	BuildSecondsUsed  int                    `json:"build_seconds_used"`
	Name              string                 `json:"name"`
	Trigger           *PipelineStepTrigger   `json:"trigger"`
	DurationInSeconds int                    `json:"duration_in_seconds"`
	RunNumber         int                    `json:"run_number"`
	Pipeline          *PipelineReference     `json:"pipeline"`
}

type PipelineCommand struct {
	CommandType string `json:"commandType"`
	Name        string `json:"name"`
	Command     string `json:"command"`
	Action      string `json:"action,omitempty"`
}

type PipelineStepTrigger struct {
	Type string `json:"type"`
}

type PipelineReference struct {
	Type string `json:"type"`
	UUID string `json:"uuid"`
}

type PipelineStepState struct {
	Name   string              `json:"name"`
	Type   string              `json:"type"`
	Result *PipelineStepResult `json:"result,omitempty"`
}

type PipelineStepResult struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type PipelineCreator struct {
	DisplayName string `json:"display_name"`
	Type        string `json:"type"`
	UUID        string `json:"uuid"`
	AccountID   string `json:"account_id"`
	Nickname    string `json:"nickname"`
}

type PipelineRepository struct {
	Type     string `json:"type"`
	FullName string `json:"full_name"`
	Name     string `json:"name"`
	UUID     string `json:"uuid"`
}

type PipelineTarget struct {
	Source            string               `json:"source"`
	Destination       string               `json:"destination"`
	DestinationCommit *PipelineCommit      `json:"destination_commit"`
	PullRequest       *PipelinePullRequest `json:"pullrequest"`
	Selector          *PipelineSelector    `json:"selector"`
	Commit            *PipelineCommit      `json:"commit"`
	Type              string               `json:"type"`
}

type PipelineCommit struct {
	Hash string `json:"hash"`
	Type string `json:"type"`
}

type PipelinePullRequest struct {
	Type  string `json:"type"`
	ID    int    `json:"id"`
	Title string `json:"title"`
	Draft bool   `json:"draft"`
}

type PipelineSelector struct {
	Type    string `json:"type"`
	Pattern string `json:"pattern"`
}

type PipelineTrigger struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type PipelineState struct {
	Name   string          `json:"name"`
	Type   string          `json:"type"`
	Result *PipelineResult `json:"result"`
}

type PipelineResult struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
