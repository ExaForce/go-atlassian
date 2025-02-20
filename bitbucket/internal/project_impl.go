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

// NewProjectService creates a new instance of the ProjectService
func NewProjectService(client service.Connector) *ProjectService {
	return &ProjectService{
		internalClient: &internalProjectServiceImpl{c: client},
	}
}

type ProjectService struct {
	internalClient bitbucket.ProjectConnector
}

type internalProjectServiceImpl struct {
	c service.Connector
}

// ListExplicitProjectUserPermissions returns a paginated list of all explicit project user permissions
//
// GET /2.0/workspaces/{workspace}/projects/{project_key}/permissions-config/users
//
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-projects/#api-workspaces-workspace-projects-project-key-permissions-config-users-get
func (p *ProjectService) ListExplicitProjectUserPermissions(ctx context.Context, workspace string, projectKey string, opts *model.PageOptions) (*model.ProjectUserMembershipPageScheme, *model.ResponseScheme, error) {
	return p.internalClient.GetExplicitProjectUserPermissions(ctx, workspace, projectKey, opts)
}

// ListExplicitProjectGroupPermissions returns a paginated list of all explicit project group permissions
//
// GET /2.0/workspaces/{workspace}/projects/{project_key}/permissions-config/groups
//
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-projects/#api-workspaces-workspace-projects-project-key-permissions-config-groups-get
func (p *ProjectService) ListExplicitProjectGroupPermissions(ctx context.Context, workspace string, projectKey string, opts *model.PageOptions) (*model.ProjectGroupMembershipPageScheme, *model.ResponseScheme, error) {
	return p.internalClient.GetExplicitProjectGroupPermissions(ctx, workspace, projectKey, opts)
}

func (i *internalProjectServiceImpl) GetExplicitProjectUserPermissions(ctx context.Context, workspace string, projectKey string, opts *model.PageOptions) (*model.ProjectUserMembershipPageScheme, *model.ResponseScheme, error) {
	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	if projectKey == "" {
		return nil, nil, model.ErrNoProjectKey
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/projects/%v/permissions-config/users", workspace, projectKey)

	// Add pagination parameters
	urlStr, err := utils.AddPaginationParams(endpoint, opts)
	if err != nil {
		return nil, nil, err
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, urlStr, "", nil)
	if err != nil {
		return nil, nil, err
	}

	projectUserMembership := new(model.ProjectUserMembershipPageScheme)
	response, err := i.c.Call(request, projectUserMembership)
	if err != nil {
		return nil, response, err
	}

	return projectUserMembership, response, nil
}

func (i *internalProjectServiceImpl) GetExplicitProjectGroupPermissions(ctx context.Context, workspace string, projectKey string, opts *model.PageOptions) (*model.ProjectGroupMembershipPageScheme, *model.ResponseScheme, error) {
	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	if projectKey == "" {
		return nil, nil, model.ErrNoProjectKey
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/projects/%v/permissions-config/groups", workspace, projectKey)

	// Add pagination parameters
	urlStr, err := utils.AddPaginationParams(endpoint, opts)
	if err != nil {
		return nil, nil, err
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, urlStr, "", nil)
	if err != nil {
		return nil, nil, err
	}

	projectGroupMembership := new(model.ProjectGroupMembershipPageScheme)
	response, err := i.c.Call(request, projectGroupMembership)
	if err != nil {
		return nil, response, err
	}

	return projectGroupMembership, response, nil
}
