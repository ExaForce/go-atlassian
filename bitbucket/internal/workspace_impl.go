package internal

import (
	"context"
	"fmt"
	"net/http"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/utils"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/bitbucket"
)

// NewWorkspaceService handles communication with the workspace related methods of the Bitbucket API.
func NewWorkspaceService(client service.Connector, webhook *WorkspaceHookService, permission *WorkspacePermissionService, repository *RepositoryService, project *ProjectService) *WorkspaceService {

	return &WorkspaceService{
		internalClient: &internalWorkspaceServiceImpl{c: client},
		Hook:           webhook,
		Permission:     permission,
		Repository:     repository,
		Project:        project,
	}
}

// WorkspaceService handles communication with the workspace related methods of the Bitbucket API.
type WorkspaceService struct {
	internalClient bitbucket.WorkspaceConnector
	Hook           *WorkspaceHookService
	Permission     *WorkspacePermissionService
	Repository     *RepositoryService
	Project        *ProjectService
}

// Get returns the requested workspace with pagination support.
//
// GET /2.0/workspaces/{workspace}
//
// https://docs.go-atlassian.io/bitbucket-cloud/workspace#get-a-workspace
func (w *WorkspaceService) Get(ctx context.Context, workspace string, opts *model.PageOptions) (*model.WorkspaceScheme, *model.ResponseScheme, error) {
	return w.internalClient.Get(ctx, workspace, opts)
}

// Members returns all members of the requested workspace with pagination support.
//
// GET /2.0/workspaces/{workspace}/members
func (w *WorkspaceService) Members(ctx context.Context, workspace string, opts *model.PageOptions) (*model.WorkspaceMembershipPageScheme, *model.ResponseScheme, error) {
	return w.internalClient.Members(ctx, workspace, opts)
}

// Membership returns the workspace membership.
//
// which includes a User object for the member and a Workspace object for the requested workspace.
//
// GET /2.0/workspaces/{workspace}/members/{memberID}
//
// https://docs.go-atlassian.io/bitbucket-cloud/workspace#get-member-in-a-workspace
func (w *WorkspaceService) Membership(ctx context.Context, workspace, memberID string) (*model.WorkspaceMembershipScheme, *model.ResponseScheme, error) {
	return w.internalClient.Membership(ctx, workspace, memberID)
}

// Projects returns the list of projects in this workspace with pagination support.
//
// GET /2.0/workspaces/{workspace}/projects
func (w *WorkspaceService) Projects(ctx context.Context, workspace string, opts *model.PageOptions) (*model.BitbucketProjectPageScheme, *model.ResponseScheme, error) {
	return w.internalClient.Projects(ctx, workspace, opts)
}

type internalWorkspaceServiceImpl struct {
	c service.Connector
}

// Get returns the requested workspace with pagination support.
func (i *internalWorkspaceServiceImpl) Get(ctx context.Context, workspace string, opts *model.PageOptions) (*model.WorkspaceScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v", workspace)

	urlStr, err := utils.AddPaginationParams(endpoint, opts)
	if err != nil {
		return nil, nil, err
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, urlStr, "", nil)
	if err != nil {
		return nil, nil, err
	}

	ws := new(model.WorkspaceScheme)
	response, err := i.c.Call(request, ws)
	if err != nil {
		return nil, response, err
	}

	return ws, response, nil
}

// Members returns all members of the requested workspace with pagination support.
//
// GET /2.0/workspaces/{workspace}/members
func (i *internalWorkspaceServiceImpl) Members(ctx context.Context, workspace string, opts *model.PageOptions) (*model.WorkspaceMembershipPageScheme, *model.ResponseScheme, error) {
	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/members", workspace)

	urlStr, err := utils.AddPaginationParams(endpoint, opts)
	if err != nil {
		return nil, nil, err
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, urlStr, "", nil)
	if err != nil {
		return nil, nil, err
	}

	members := new(model.WorkspaceMembershipPageScheme)
	response, err := i.c.Call(request, members)
	if err != nil {
		return nil, response, err
	}

	return members, response, nil
}

// Membership returns the workspace membership.
func (i *internalWorkspaceServiceImpl) Membership(ctx context.Context, workspace, memberID string) (*model.WorkspaceMembershipScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	if memberID == "" {
		return nil, nil, model.ErrNoMemberID
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/members/%v", workspace, memberID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	member := new(model.WorkspaceMembershipScheme)
	response, err := i.c.Call(request, member)
	if err != nil {
		return nil, response, err
	}

	return member, response, nil
}

// Projects returns the list of projects in this workspace with pagination support.
func (i *internalWorkspaceServiceImpl) Projects(ctx context.Context, workspace string, opts *model.PageOptions) (*model.BitbucketProjectPageScheme, *model.ResponseScheme, error) {
	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/projects", workspace)

	urlStr, err := utils.AddPaginationParams(endpoint, opts)
	if err != nil {
		return nil, nil, err
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, urlStr, "", nil)
	if err != nil {
		return nil, nil, err
	}

	projects := new(model.BitbucketProjectPageScheme)
	response, err := i.c.Call(request, projects)
	if err != nil {
		return nil, response, err
	}

	return projects, response, nil
}
