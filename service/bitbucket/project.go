package bitbucket

import (
	"context"

	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type ProjectConnector interface {
	// Get explicit user permissions for a project
	//
	// GET /2.0/workspaces/{workspace}/projects/{project_key}/permissions-config/users
	//
	// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-projects/#api-workspaces-workspace-projects-project-key-permissions-config-users-get
	GetExplicitProjectUserPermissions(ctx context.Context, workspace string, projectKey string, opts *models.PageOptions) (*models.ProjectUserMembershipPageScheme, *models.ResponseScheme, error)

	// Get explicit groups permissions for a project
	//
	// GET /2.0/workspaces/{workspace}/projects/{project_key}/permissions-config/groups
	//
	// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-projects/#api-workspaces-workspace-projects-project-key-permissions-config-groups-get
	GetExplicitProjectGroupPermissions(ctx context.Context, workspace string, projectKey string, opts *models.PageOptions) (*models.ProjectGroupMembershipPageScheme, *models.ResponseScheme, error)
}
