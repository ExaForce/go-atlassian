package internal

import (
	"context"
	"fmt"
	"net/http"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/admin"
)

type OrgTokenService struct {
	internalClient admin.OrgTokenConnector
}

type OrgKeyService struct {
	internalClient admin.OrgKeyConnector
}

func NewOrgTokenService(client service.Connector) *OrgTokenService {
	return &OrgTokenService{internalClient: &internalOrgTokenImpl{c: client}}
}

func NewOrgKeyService(client service.Connector) *OrgKeyService {
	return &OrgKeyService{internalClient: &internalOrgKeyImpl{c: client}}
}

// Gets gets the API tokens owned by the specified organization.
//
// GET /admin/api-access/v1/orgs/{orgID}/api-tokens
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/org/token#get-api-tokens
func (o *OrgTokenService) Gets(ctx context.Context, orgID string, params *model.OrgTokenQueryParams) (*model.OrgTokenPageScheme, *model.ResponseScheme, error) {
	return o.internalClient.Gets(ctx, orgID, params)
}

type internalOrgTokenImpl struct {
	c service.Connector
}

// Gets gets the API keys owned by the specified organization.
//
// GET /admin/api-access/v1/orgs/{orgID}/api-keys
//
// https://docs.go-atlassian.io/atlassian-admin-cloud/org/key#get-api-keys
func (o *OrgKeyService) Gets(ctx context.Context, orgID string, params *model.OrgTokenQueryParams) (*model.OrgKeyPageScheme, *model.ResponseScheme, error) {
	return o.internalClient.Gets(ctx, orgID, params)
}

type internalOrgKeyImpl struct {
	c service.Connector
}

func (i *internalOrgKeyImpl) Gets(ctx context.Context, orgID string, params *model.OrgTokenQueryParams) (*model.OrgKeyPageScheme, *model.ResponseScheme, error) {

	if orgID == "" {
		return nil, nil, model.ErrNoAdminOrganization
	}

	endpoint := fmt.Sprintf("admin/api-access/v1/orgs/%v/api-keys", orgID)
	if params != nil && (params.PageSize > 0 || params.Cursor != "") {
		endpoint = fmt.Sprintf("%v?pageSize=%v&cursor=%v", endpoint, params.PageSize, params.Cursor)
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var keys *model.OrgKeyPageScheme
	response, err := i.c.Call(request, &keys)
	if err != nil {
		return nil, response, err
	}

	return keys, response, nil
}

func (i *internalOrgTokenImpl) Gets(ctx context.Context, orgID string, params *model.OrgTokenQueryParams) (*model.OrgTokenPageScheme, *model.ResponseScheme, error) {

	if orgID == "" {
		return nil, nil, model.ErrNoAdminOrganization
	}

	endpoint := fmt.Sprintf("admin/api-access/v1/orgs/%v/api-tokens", orgID)

	if params != nil && (params.PageSize > 0 || params.Cursor != "") {
		endpoint = fmt.Sprintf("%v?pageSize=%v&cursor=%v", endpoint, params.PageSize, params.Cursor)
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var tokens *model.OrgTokenPageScheme
	response, err := i.c.Call(request, &tokens)
	if err != nil {
		return nil, response, err
	}

	return tokens, response, nil
}
