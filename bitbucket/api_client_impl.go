package bitbucket

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ctreminiom/go-atlassian/v2/bitbucket/internal"
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service/common"
)

// DefaultBitbucketSite is the default Bitbucket API site.
const DefaultBitbucketSite = "https://api.bitbucket.org"

// New creates a new Bitbucket API client.
func New(httpClient common.HTTPClient, site string, config *models.ClientConfig) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if site == "" {
		site = DefaultBitbucketSite
	}

	if !strings.HasSuffix(site, "/") {
		site += "/"
	}

	u, err := url.Parse(site)
	if err != nil {
		return nil, err
	}

	// Use default config if not provided
	if config == nil {
		config = &models.ClientConfig{
			MaxRetries:        5,
			InitialRetryDelay: time.Duration(1) * time.Minute,
			MaxRetryDelay:     time.Duration(10) * time.Minute,
		}
	}

	client := &Client{
		HTTP:              httpClient,
		Site:              u,
		MaxRetries:        config.MaxRetries,
		InitialRetryDelay: config.InitialRetryDelay,
		MaxRetryDelay:     config.MaxRetryDelay,
	}

	client.Auth = internal.NewAuthenticationService(client)

	client.Workspace = internal.NewWorkspaceService(client,
		internal.NewWorkspaceHookService(client),
		internal.NewWorkspacePermissionService(client),
		internal.NewRepositoryService(client),
		internal.NewProjectService(client),
	)

	return client, nil
}

// Client is a Bitbucket API client.
type Client struct {
	HTTP              common.HTTPClient
	Site              *url.URL
	MaxRetries        int
	InitialRetryDelay time.Duration
	MaxRetryDelay     time.Duration
	Auth              common.Authentication
	Workspace         *internal.WorkspaceService
}

// NewRequest creates an API request.
func (c *Client) NewRequest(ctx context.Context, method, urlStr, typ string, body interface{}) (*http.Request, error) {

	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.Site.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		if err = json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}

	// If the body interface is a *bytes.Buffer type
	// it means the NewRequest() requires to handle the RFC 1867 ISO
	if attachBuffer, ok := body.(*bytes.Buffer); ok {
		buf = attachBuffer
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if typ != "" {
		// When the typ is provided, it means the request needs to be created to handle files
		req.Header.Set("Content-Type", typ)
		req.Header.Set("X-Atlassian-Token", "no-check")
	}

	if c.Auth.HasBasicAuth() {
		req.SetBasicAuth(c.Auth.GetBasicAuth())
	}

	if c.Auth.HasUserAgent() {
		req.Header.Set("User-Agent", c.Auth.GetUserAgent())
	}

	if c.Auth.GetBearerToken() != "" && !c.Auth.HasBasicAuth() {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", c.Auth.GetBearerToken()))
	}

	return req, nil
}

// Call executes an API request and returns the response.
func (c *Client) Call(request *http.Request, structure interface{}) (*models.ResponseScheme, error) {
	retryCount := 0
	ctx := request.Context()

	for {
		response, err := c.HTTP.Do(request)
		if err != nil {
			return nil, err
		}

		// If rate limit exceeded, sleep with exponential backoff
		if response.StatusCode == http.StatusTooManyRequests {
			delay := c.InitialRetryDelay
			// Use bit shifting for exponential backoff (1 << retryCount)
			delay = delay * (1 << uint(retryCount))
			if delay > c.MaxRetryDelay {
				delay = c.MaxRetryDelay
			}
			log.Printf("Rate limit exceeded, sleeping for %v request %v", delay, request.URL.String())

			// Get timer
			timer := time.NewTimer(delay)
			defer timer.Stop()

			// Wait for either context cancellation or timer
			select {
			case <-ctx.Done():
				timer.Stop()
				return nil, ctx.Err()
			case <-timer.C:
				// Timer completed successfully
			}

			retryCount++
			if retryCount > c.MaxRetries {
				return c.processResponse(response, structure)
			}
			continue
		}

		return c.processResponse(response, structure)
	}
}

func (c *Client) processResponse(response *http.Response, structure interface{}) (*models.ResponseScheme, error) {

	defer response.Body.Close()

	res := &models.ResponseScheme{
		Response: response,
		Code:     response.StatusCode,
		Endpoint: response.Request.URL.String(),
		Method:   response.Request.Method,
	}

	responseAsBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return res, err
	}

	res.Bytes.Write(responseAsBytes)

	wasSuccess := response.StatusCode >= 200 && response.StatusCode < 300

	if !wasSuccess {

		switch response.StatusCode {

		case http.StatusNotFound:
			return res, models.ErrNotFound

		case http.StatusUnauthorized:
			return res, models.ErrUnauthorized

		case http.StatusInternalServerError:
			return res, models.ErrInternal

		case http.StatusBadRequest:
			return res, models.ErrBadRequest

		case http.StatusTooManyRequests:
			return res, models.ErrRateLimited

		default:
			return res, models.ErrInvalidStatusCode
		}
	}

	if structure != nil {
		if err = json.Unmarshal(responseAsBytes, &structure); err != nil {
			return res, err
		}
	}

	return res, nil
}
