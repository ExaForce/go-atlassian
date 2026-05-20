package bitbucket

import (
	"context"

	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// UserConnector exposes endpoints that operate on a Bitbucket user account.
type UserConnector interface {

	// SSHKeys returns a paginated list of SSH keys registered on the selected user account.
	//
	// selectedUser may be an account UUID (URL-encoded including braces), an Atlassian
	// account ID, or a username. The caller is responsible for encoding the identifier
	// appropriately; the implementation passes it through to the URL path as-is.
	//
	// GET /2.0/users/{selected_user}/ssh-keys
	//
	// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-ssh/#api-users-selected-user-ssh-keys-get
	SSHKeys(ctx context.Context, selectedUser string, opts *models.PageOptions) (*models.SSHKeyPageScheme, *models.ResponseScheme, error)
}
