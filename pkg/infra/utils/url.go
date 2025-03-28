package utils

import (
	"net/url"
	"strconv"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// AddPaginationParams adds pagination parameters to the given endpoint URL
// and returns the complete URL string with encoded query parameters
func AddPaginationParams(endpoint string, opts *model.PageOptions) (string, error) {
	// Parse the endpoint URL
	u, err := url.Parse(endpoint)
	if err != nil {
		return "", err
	}

	// Get existing query parameters if any
	q := u.Query()

	// Add pagination parameters if options are provided
	if opts != nil {
		if opts.Page > 0 {
			q.Add("page", strconv.Itoa(opts.Page))
		}
		if opts.PageLen > 0 {
			q.Add("pagelen", strconv.Itoa(opts.PageLen))
		}
		if opts.Q != "" {
			q.Add("q", opts.Q)
		}
	}

	// Set the query parameters
	u.RawQuery = q.Encode()

	return u.String(), nil
}
