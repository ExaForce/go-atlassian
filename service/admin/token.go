package admin

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// UserTokenConnector represents the cloud admin user token endpoints.
// Use it to search, get, create, delete, and change tokens.
type UserTokenConnector interface {

	// Gets gets the API tokens owned by the specified user.
	//
	// GET /users/{accountID}/manage/api-tokens
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/user/token#get-api-tokens
	Gets(ctx context.Context, accountID string) ([]*model.UserTokensScheme, *model.ResponseScheme, error)

	// Delete deletes a specified API token by ID.
	//
	// DELETE /users/{accountID}/manage/api-tokens/{tokenID}
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/user/token#delete-api-token
	Delete(ctx context.Context, accountID, tokenID string) (*model.ResponseScheme, error)
}

type OrgTokenConnector interface {

	// Gets gets the API tokens owned by the specified organization.
	//
	// GET /admin/api-access/v1/orgs/{orgID}/api-tokens
	//
	// https://developer.atlassian.com/cloud/admin/api-access/rest/api-group-api-key/#api-orgs-orgid-api-keys-get
	Gets(ctx context.Context, orgID string, params *model.OrgTokenQueryParams) (*model.OrgTokenPageScheme, *model.ResponseScheme, error)
}

type OrgKeyConnector interface {

	// Gets gets the API keys owned by the specified organization.
	//
	// GET /admin/api-access/v1/orgs/{orgID}/api-keys
	//
	// https://developer.atlassian.com/cloud/admin/api-access/rest/api-group-api-key/#api-orgs-orgid-api-keys-get
	Gets(ctx context.Context, orgID string, params *model.OrgTokenQueryParams) (*model.OrgKeyPageScheme, *model.ResponseScheme, error)
}
