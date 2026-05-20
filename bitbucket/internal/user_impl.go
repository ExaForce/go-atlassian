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

// NewUserService creates a new UserService.
func NewUserService(client service.Connector) *UserService {
	return &UserService{
		internalClient: &internalUserServiceImpl{c: client},
	}
}

// UserService handles communication with the user-scoped methods of the Bitbucket API.
type UserService struct {
	internalClient bitbucket.UserConnector
}

// SSHKeys returns a paginated list of SSH keys registered on the selected user account.
//
// GET /2.0/users/{selected_user}/ssh-keys
//
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-ssh/#api-users-selected-user-ssh-keys-get
func (u *UserService) SSHKeys(ctx context.Context, selectedUser string, opts *model.PageOptions) (*model.SSHKeyPageScheme, *model.ResponseScheme, error) {
	return u.internalClient.SSHKeys(ctx, selectedUser, opts)
}

type internalUserServiceImpl struct {
	c service.Connector
}

func (i *internalUserServiceImpl) SSHKeys(ctx context.Context, selectedUser string, opts *model.PageOptions) (*model.SSHKeyPageScheme, *model.ResponseScheme, error) {

	if selectedUser == "" {
		return nil, nil, model.ErrNoSelectedUser
	}

	endpoint := fmt.Sprintf("2.0/users/%v/ssh-keys", selectedUser)

	urlStr, err := utils.AddPaginationParams(endpoint, opts)
	if err != nil {
		return nil, nil, err
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, urlStr, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.SSHKeyPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}
