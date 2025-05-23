package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/utils"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/bitbucket"
)

// NewRepositoryService creates a new instance of the repository service
func NewRepositoryService(client service.Connector) *RepositoryService {
	return &RepositoryService{
		internalClient: &internalRepositoryServiceImpl{c: client},
	}
}

// RepositoryService handles communication with the repository related methods
type RepositoryService struct {
	internalClient bitbucket.RepositoryConnector
}

// List returns a paginated list of all repositories owned by the specified workspace
//
// GET /2.0/repositories/{workspace}
//
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-repositories/#api-repositories-workspace-get
func (r *RepositoryService) List(ctx context.Context, workspace string, opts *model.PageOptions) (*model.RepositoryPageScheme, *model.ResponseScheme, error) {
	return r.internalClient.List(ctx, workspace, opts)
}

type internalRepositoryServiceImpl struct {
	c service.Connector
}

// List returns a paginated list of all repositories owned by the specified workspace
func (i *internalRepositoryServiceImpl) List(ctx context.Context, workspace string, opts *model.PageOptions) (*model.RepositoryPageScheme, *model.ResponseScheme, error) {
	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	endpoint := fmt.Sprintf("2.0/repositories/%v", workspace)

	// Add pagination parameters
	urlStr, err := utils.AddPaginationParams(endpoint, opts)
	if err != nil {
		return nil, nil, err
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, urlStr, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.RepositoryPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

// Create creates a new repository in the specified workspace
//
// POST /2.0/repositories/{workspace}/{repo_slug}
func (i *internalRepositoryServiceImpl) Create(ctx context.Context, workspace string, repository string, payload *model.RepositoryScheme) (*model.RepositoryScheme, *model.ResponseScheme, error) {
	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	if repository == "" {
		return nil, nil, model.ErrNoRepository
	}

	endpoint := fmt.Sprintf("2.0/repositories/%v/%v", workspace, repository)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	repo := new(model.RepositoryScheme)
	response, err := i.c.Call(request, repo)
	if err != nil {
		return nil, response, err
	}

	return repo, response, nil
}

// Also add the method to the RepositoryService
func (r *RepositoryService) Create(ctx context.Context, workspace string, repository string, payload *model.RepositoryScheme) (*model.RepositoryScheme, *model.ResponseScheme, error) {
	return r.internalClient.Create(ctx, workspace, repository, payload)
}

// ListBranchRestrictions returns a paginated list of all branch restrictions on the repository
//
// GET /2.0/repositories/{workspace}/{repo_slug}/branch-restrictions
//
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-branch-restrictions/#api-repositories-workspace-repo-slug-branch-restrictions-get
func (r *RepositoryService) ListBranchRestrictions(ctx context.Context, workspace, repoSlug string, opts *model.PageOptions) (*model.BranchRestrictionsPageScheme, *model.ResponseScheme, error) {
	return r.internalClient.ListBranchRestrictions(ctx, workspace, repoSlug, opts)
}

func (i *internalRepositoryServiceImpl) ListBranchRestrictions(ctx context.Context, workspace, repoSlug string, opts *model.PageOptions) (*model.BranchRestrictionsPageScheme, *model.ResponseScheme, error) {
	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	if repoSlug == "" {
		return nil, nil, model.ErrNoRepository
	}

	endpoint := fmt.Sprintf("2.0/repositories/%v/%v/branch-restrictions", workspace, repoSlug)

	// Add pagination parameters
	urlStr, err := utils.AddPaginationParams(endpoint, opts)
	if err != nil {
		return nil, nil, err
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, urlStr, "", nil)
	if err != nil {
		return nil, nil, err
	}

	branchRestrictions := new(model.BranchRestrictionsPageScheme)
	response, err := i.c.Call(request, branchRestrictions)
	if err != nil {
		return nil, response, err
	}

	return branchRestrictions, response, nil
}

// ListDefaultReviewers returns a paginated list of default reviewers for the repository
//
// GET /2.0/repositories/{workspace}/{repo_slug}/default-reviewers
//
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-default-reviewers-get
func (r *RepositoryService) ListDefaultReviewers(ctx context.Context, workspace, repoSlug string, opts *model.PageOptions) (*model.DefaultReviewersPageScheme, *model.ResponseScheme, error) {
	return r.internalClient.ListDefaultReviewers(ctx, workspace, repoSlug, opts)
}

func (i *internalRepositoryServiceImpl) ListDefaultReviewers(ctx context.Context, workspace, repoSlug string, opts *model.PageOptions) (*model.DefaultReviewersPageScheme, *model.ResponseScheme, error) {
	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	if repoSlug == "" {
		return nil, nil, model.ErrNoRepository
	}

	endpoint := fmt.Sprintf("2.0/repositories/%v/%v/default-reviewers", workspace, repoSlug)

	// Add pagination parameters
	urlStr, err := utils.AddPaginationParams(endpoint, opts)
	if err != nil {
		return nil, nil, err
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, urlStr, "", nil)
	if err != nil {
		return nil, nil, err
	}

	defaultReviewers := new(model.DefaultReviewersPageScheme)
	response, err := i.c.Call(request, defaultReviewers)
	if err != nil {
		return nil, response, err
	}

	return defaultReviewers, response, nil
}

// ListPullRequests returns all pull requests on the specified repository
//
// GET /2.0/repositories/{workspace}/{repo_slug}/pullrequests
//
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-get
func (r *RepositoryService) ListPullRequests(ctx context.Context, workspace, repoSlug string, opts *model.PageOptions) (*models.PullRequestsResponse, *models.ResponseScheme, error) {
	return r.internalClient.ListPullRequests(ctx, workspace, repoSlug, opts)
}

func (i *internalRepositoryServiceImpl) ListPullRequests(ctx context.Context, workspace, repoSlug string, opts *model.PageOptions) (*models.PullRequestsResponse, *models.ResponseScheme, error) {
	if workspace == "" {
		return nil, nil, models.ErrNoWorkspace
	}

	if repoSlug == "" {
		return nil, nil, models.ErrNoRepository
	}

	endpoint := fmt.Sprintf("2.0/repositories/%v/%v/pullrequests", workspace, repoSlug)

	// Add pagination parameters first
	urlStr, err := utils.AddPaginationParams(endpoint, opts)
	if err != nil {
		return nil, nil, err
	}

	// Parse URL to add additional parameters
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, nil, err
	}

	// Add state parameters
	q := u.Query()
	q.Add("state", "OPEN,MERGED,DECLINED,SUPERSEDED")
	u.RawQuery = q.Encode()

	request, err := i.c.NewRequest(ctx, http.MethodGet, u.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	pullRequests := new(models.PullRequestsResponse)
	response, err := i.c.Call(request, pullRequests)
	if err != nil {
		return nil, response, err
	}

	return pullRequests, response, nil
}

// ListDeployKeys returns a paginated list of all deploy keys for the specified repository
//
// GET /2.0/repositories/{workspace}/{repo_slug}/deploy-keys
//
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-deployments/#api-repositories-workspace-repo-slug-deploy-keys-get
func (r *RepositoryService) ListDeployKeys(ctx context.Context, workspace, repoSlug string, opts *model.PageOptions) (*model.DeployKeysPageScheme, *model.ResponseScheme, error) {
	return r.internalClient.ListDeployKeys(ctx, workspace, repoSlug, opts)
}

func (i *internalRepositoryServiceImpl) ListDeployKeys(ctx context.Context, workspace, repoSlug string, opts *model.PageOptions) (*model.DeployKeysPageScheme, *model.ResponseScheme, error) {
	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	if repoSlug == "" {
		return nil, nil, model.ErrNoRepository
	}

	endpoint := fmt.Sprintf("2.0/repositories/%v/%v/deploy-keys", workspace, repoSlug)

	// Add pagination parameters
	urlStr, err := utils.AddPaginationParams(endpoint, opts)
	if err != nil {
		return nil, nil, err
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, urlStr, "", nil)
	if err != nil {
		return nil, nil, err
	}

	deployKeys := new(model.DeployKeysPageScheme)
	response, err := i.c.Call(request, deployKeys)
	if err != nil {
		return nil, response, err
	}

	return deployKeys, response, nil
}

// ListRepositoryExplicitGroupPermissions returns a paginated list of all explicit group permissions for the specified repository
//
// GET /2.0/repositories/{workspace}/{repo_slug}/permissions-config/groups
//
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-repositories/#api-repositories-workspace-repo-slug-permissions-config-groups-get
func (r *RepositoryService) ListRepositoryExplicitGroupPermissions(ctx context.Context, workspace, repoSlug string, opts *model.PageOptions) (*model.RepositoryGroupPermissionsPageScheme, *model.ResponseScheme, error) {
	return r.internalClient.ListRepositoryExplicitGroupPermissions(ctx, workspace, repoSlug, opts)
}

func (i *internalRepositoryServiceImpl) ListRepositoryExplicitGroupPermissions(ctx context.Context, workspace, repoSlug string, opts *model.PageOptions) (*model.RepositoryGroupPermissionsPageScheme, *model.ResponseScheme, error) {
	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	if repoSlug == "" {
		return nil, nil, model.ErrNoRepository
	}

	endpoint := fmt.Sprintf("2.0/repositories/%v/%v/permissions-config/groups", workspace, repoSlug)

	// Add pagination parameters
	urlStr, err := utils.AddPaginationParams(endpoint, opts)
	if err != nil {
		return nil, nil, err
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, urlStr, "", nil)
	if err != nil {
		return nil, nil, err
	}

	repositoryGroupPermissions := new(model.RepositoryGroupPermissionsPageScheme)
	response, err := i.c.Call(request, repositoryGroupPermissions)
	if err != nil {
		return nil, response, err
	}

	return repositoryGroupPermissions, response, nil
}

// ListRepositoryPipelineVariables returns a paginated list of all pipeline variables for the specified repository
//
// GET /2.0/repositories/{workspace}/{repo_slug}/pipelines_config/variables
//
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-repositories/#api-repositories-workspace-repo-slug-pipelines-config-variables-get
func (r *RepositoryService) ListRepositoryPipelineVariables(ctx context.Context, workspace, repoSlug string, opts *model.PageOptions) (*model.RepositoryPipelineVariablesPageScheme, *model.ResponseScheme, error) {
	return r.internalClient.ListRepositoryPipelineVariables(ctx, workspace, repoSlug, opts)
}

func (i *internalRepositoryServiceImpl) ListRepositoryPipelineVariables(ctx context.Context, workspace, repoSlug string, opts *model.PageOptions) (*model.RepositoryPipelineVariablesPageScheme, *model.ResponseScheme, error) {
	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	if repoSlug == "" {
		return nil, nil, model.ErrNoRepository
	}

	endpoint := fmt.Sprintf("2.0/repositories/%v/%v/pipelines_config/variables", workspace, repoSlug)

	// Add pagination parameters
	urlStr, err := utils.AddPaginationParams(endpoint, opts)
	if err != nil {
		return nil, nil, err
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, urlStr, "", nil)
	if err != nil {
		return nil, nil, err
	}

	pipelineVariables := new(model.RepositoryPipelineVariablesPageScheme)
	response, err := i.c.Call(request, pipelineVariables)
	if err != nil {
		return nil, response, err
	}

	return pipelineVariables, response, nil
}

// ListRepositoryPipelineRuns returns a paginated list of all pipeline runs for the specified repository
//
// GET /2.0/repositories/{workspace}/{repo_slug}/pipelines
//
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-repositories/#api-repositories-workspace-repo-slug-pipelines-get
func (r *RepositoryService) ListRepositoryPipelineRuns(ctx context.Context, workspace, repoSlug string, opts *model.PageOptions) (*model.RepositoryPipelineRunsPageScheme, *model.ResponseScheme, error) {
	return r.internalClient.ListRepositoryPipelineRuns(ctx, workspace, repoSlug, opts)
}

func (i *internalRepositoryServiceImpl) ListRepositoryPipelineRuns(ctx context.Context, workspace, repoSlug string, opts *model.PageOptions) (*model.RepositoryPipelineRunsPageScheme, *model.ResponseScheme, error) {
	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	if repoSlug == "" {
		return nil, nil, model.ErrNoRepository
	}

	endpoint := fmt.Sprintf("2.0/repositories/%v/%v/pipelines", workspace, repoSlug)

	// Add pagination parameters
	urlStr, err := utils.AddPaginationParams(endpoint, opts)
	if err != nil {
		return nil, nil, err
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, urlStr, "", nil)
	if err != nil {
		return nil, nil, err
	}

	pipelineRuns := new(model.RepositoryPipelineRunsPageScheme)
	response, err := i.c.Call(request, pipelineRuns)
	if err != nil {
		return nil, response, err
	}

	return pipelineRuns, response, nil
}

// ListRepositoryPipelineRunSteps returns a paginated list of all pipeline run steps for the specified pipeline run
//
// GET /2.0/repositories/{workspace}/{repo_slug}/pipelines/{pipeline_uuid}/steps
//
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pipelines/#api-repositories-workspace-repo-slug-pipelines-pipeline-uuid-steps-get
func (r *RepositoryService) ListRepositoryPipelineRunSteps(ctx context.Context, workspace, repoSlug string, pipelineUUID string, opts *model.PageOptions) (*model.RepositoryPipelineRunStepsPageScheme, *model.ResponseScheme, error) {
	return r.internalClient.ListRepositoryPipelineRunSteps(ctx, workspace, repoSlug, pipelineUUID, opts)
}

func (i *internalRepositoryServiceImpl) ListRepositoryPipelineRunSteps(ctx context.Context, workspace, repoSlug string, pipelineUUID string, opts *model.PageOptions) (*model.RepositoryPipelineRunStepsPageScheme, *model.ResponseScheme, error) {
	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	if repoSlug == "" {
		return nil, nil, model.ErrNoRepository
	}

	if pipelineUUID == "" {
		return nil, nil, model.ErrNoPipelineUUID
	}

	endpoint := fmt.Sprintf("2.0/repositories/%v/%v/pipelines/%v/steps", workspace, repoSlug, pipelineUUID)

	// Add pagination parameters
	urlStr, err := utils.AddPaginationParams(endpoint, opts)
	if err != nil {
		return nil, nil, err
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, urlStr, "", nil)
	if err != nil {
		return nil, nil, err
	}

	pipelineRunSteps := new(model.RepositoryPipelineRunStepsPageScheme)
	response, err := i.c.Call(request, pipelineRunSteps)
	if err != nil {
		return nil, response, err
	}

	return pipelineRunSteps, response, nil
}
