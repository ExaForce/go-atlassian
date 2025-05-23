package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/utils"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/bitbucket"
)

// NewWorkspacePermissionService creates a new WorkspacePermissionService instance.
// It takes a service.Connector as input and returns a pointer to WorkspacePermissionService.
func NewWorkspacePermissionService(client service.Connector) *WorkspacePermissionService {

	return &WorkspacePermissionService{
		internalClient: &internalWorkspacePermissionServiceImpl{c: client},
	}
}

// WorkspacePermissionService provides methods to interact with workspace permissions in Bitbucket.
type WorkspacePermissionService struct {
	internalClient bitbucket.WorkspacePermissionConnector
}

// Members returns the list of members in a workspace and their permission levels.
//
// GET /2.0/workspaces/{workspace}/permissions
//
// https://docs.go-atlassian.io/bitbucket-cloud/workspace/permissions#get-user-permissions-in-a-workspace
func (w *WorkspacePermissionService) Members(ctx context.Context, workspace, query string, opts *model.PageOptions) (*model.WorkspaceMembershipPageScheme, *model.ResponseScheme, error) {
	return w.internalClient.Members(ctx, workspace, query, opts)
}

// Repositories returns an object for each repository permission for all of a workspaces repositories.
//
// Permissions returned are effective permissions: the highest level of permission the user has.
//
// NOTE: Only users with admin permission for the team may access this resource.
//
// GET /2.0/workspaces/{workspace}/permissions/repositories
//
// https://docs.go-atlassian.io/bitbucket-cloud/workspace/permissions#gets-all-repository-permissions-in-a-workspace
func (w *WorkspacePermissionService) Repositories(ctx context.Context, workspace, query, sort string, opts *model.PageOptions) (*model.RepositoryPermissionPageScheme, *model.ResponseScheme, error) {
	return w.internalClient.Repositories(ctx, workspace, query, sort, opts)
}

// Repository returns an object for the repository permission of each user in the requested repository.
//
// Permissions returned are effective permissions: the highest level of permission the user has.
//
// GET /2.0/workspaces/{workspace}/permissions/repositories/{repo_slug}
//
// https://docs.go-atlassian.io/bitbucket-cloud/workspace/permissions#get-repository-permission-in-a-workspace
func (w *WorkspacePermissionService) Repository(ctx context.Context, workspace, repository, sort string, opts *model.PageOptions) (*model.RepositoryPermissionPageScheme, *model.ResponseScheme, error) {
	return w.internalClient.Repository(ctx, workspace, repository, sort, opts)
}

type internalWorkspacePermissionServiceImpl struct {
	c service.Connector
}

func (i *internalWorkspacePermissionServiceImpl) Members(ctx context.Context, workspace, query string, opts *model.PageOptions) (*model.WorkspaceMembershipPageScheme, *model.ResponseScheme, error) {
	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/permissions", workspace)

	// Add pagination parameters first
	urlStr, err := utils.AddPaginationParams(endpoint, opts)
	if err != nil {
		return nil, nil, err
	}

	// Parse the URL to add additional parameters
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, nil, err
	}

	// Add query parameter
	q := u.Query()
	if query != "" {
		q.Add("q", query)
	}
	u.RawQuery = q.Encode()

	request, err := i.c.NewRequest(ctx, http.MethodGet, u.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.WorkspaceMembershipPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalWorkspacePermissionServiceImpl) Repositories(ctx context.Context, workspace, query, sort string, opts *model.PageOptions) (*model.RepositoryPermissionPageScheme, *model.ResponseScheme, error) {
	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/permissions/repositories", workspace)

	// Add pagination parameters first
	urlStr, err := utils.AddPaginationParams(endpoint, opts)
	if err != nil {
		return nil, nil, err
	}

	// Parse the URL to add additional parameters
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, nil, err
	}

	// Add query and sort parameters
	q := u.Query()
	if query != "" {
		q.Add("q", query)
	}
	if sort != "" {
		q.Add("sort", sort)
	}
	u.RawQuery = q.Encode()

	request, err := i.c.NewRequest(ctx, http.MethodGet, u.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.RepositoryPermissionPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalWorkspacePermissionServiceImpl) Repository(ctx context.Context, workspace, repository, sort string, opts *model.PageOptions) (*model.RepositoryPermissionPageScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	if repository == "" {
		return nil, nil, model.ErrNoRepository
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/permissions/repositories/%v", workspace, repository)

	// Add pagination parameters first
	urlStr, err := utils.AddPaginationParams(endpoint, opts)
	if err != nil {
		return nil, nil, err
	}
	// Parse the URL to add additional parameters
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, nil, err
	}

	// Add query and sort parameters
	q := u.Query()
	if sort != "" {
		q.Add("sort", sort)
	}
	u.RawQuery = q.Encode()

	request, err := i.c.NewRequest(ctx, http.MethodGet, u.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.RepositoryPermissionPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}
