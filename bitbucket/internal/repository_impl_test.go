package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryService_List(t *testing.T) {

	type args struct {
		ctx       context.Context
		workspace string
	}

	testCases := []struct {
		name    string
		args    args
		on      func(*testing.T) (*RepositoryService, *models.ResponseScheme)
		wantErr bool
		Err     error
	}{
		{
			name: "when the workspace is not provided",
			args: args{
				ctx:       context.Background(),
				workspace: "",
			},
			on: func(t *testing.T) (*RepositoryService, *models.ResponseScheme) {
				client := NewRepositoryService(nil)
				return client, nil
			},
			wantErr: true,
			Err:     models.ErrNoWorkspace,
		},
		{
			name: "when the http request cannot be created",
			args: args{
				ctx:       context.Background(),
				workspace: "workspace-uuid",
			},
			on: func(t *testing.T) (*RepositoryService, *models.ResponseScheme) {
				client := NewRepositoryService(testConnector{
					requestDoer: func(req *http.Request) (*http.Response, error) {
						return nil, nil
					},
					requestMaker: func(ctx context.Context, method, path, query string, body interface{}) (*http.Request, error) {
						return nil, fmt.Errorf("error creating request")
					},
				})
				return client, nil
			},
			wantErr: true,
			Err:     fmt.Errorf("error creating request"),
		},
		{
			name: "when the http request fails",
			args: args{
				ctx:       context.Background(),
				workspace: "workspace-uuid",
			},
			on: func(t *testing.T) (*RepositoryService, *models.ResponseScheme) {
				client := NewRepositoryService(testConnector{
					requestDoer: func(req *http.Request) (*http.Response, error) {
						return nil, fmt.Errorf("error making request")
					},
					requestMaker: func(ctx context.Context, method, path, query string, body interface{}) (*http.Request, error) {
						return &http.Request{}, nil
					},
				})
				return client, nil
			},
			wantErr: true,
			Err:     fmt.Errorf("error making request"),
		},
		{
			name: "when the request is successful",
			args: args{
				ctx:       context.Background(),
				workspace: "workspace-uuid",
			},
			on: func(t *testing.T) (*RepositoryService, *models.ResponseScheme) {
				client := NewRepositoryService(testConnector{
					requestDoer: func(req *http.Request) (*http.Response, error) {
						resp := &http.Response{
							StatusCode: http.StatusOK,
							Body:       nil,
						}
						return resp, nil
					},
					requestMaker: func(ctx context.Context, method, path, query string, body interface{}) (*http.Request, error) {
						assert.Equal(t, "GET", method)
						assert.Equal(t, "2.0/repositories/workspace-uuid", path)
						return &http.Request{}, nil
					},
				})

				return client, &models.ResponseScheme{
					Response: &http.Response{
						StatusCode: http.StatusOK,
					},
				}
			},
			wantErr: false,
			Err:     nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			client, _ := testCase.on(t)

			gotResult, gotResponse, err := client.List(testCase.args.ctx, testCase.args.workspace)

			if testCase.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, testCase.Err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}
		})
	}
}

func TestRepositoryService_ListBranchRestrictions(t *testing.T) {
	type args struct {
		ctx         context.Context
		workspace   string
		repoSlug    string
		pageOptions models.PageOptions
	}

	testCases := []struct {
		name    string
		args    args
		on      func(*testing.T) (*RepositoryService, *models.ResponseScheme)
		wantErr bool
		Err     error
	}{
		{
			name: "when the workspace is not provided",
			args: args{
				ctx:       context.Background(),
				workspace: "",
				repoSlug:  "repo-slug",
			},
			on: func(t *testing.T) (*RepositoryService, *models.ResponseScheme) {
				client := NewRepositoryService(nil)
				return client, nil
			},
			wantErr: true,
			Err:     models.ErrNoWorkspace,
		},
		{
			name: "when the repository slug is not provided",
			args: args{
				ctx:       context.Background(),
				workspace: "workspace-uuid",
				repoSlug:  "",
			},
			on: func(t *testing.T) (*RepositoryService, *models.ResponseScheme) {
				client := NewRepositoryService(nil)
				return client, nil
			},
			wantErr: true,
			Err:     models.ErrNoRepository,
		},
		{
			name: "when the http request cannot be created",
			args: args{
				ctx:       context.Background(),
				workspace: "workspace-uuid",
				repoSlug:  "repo-slug",
			},
			on: func(t *testing.T) (*RepositoryService, *models.ResponseScheme) {
				client := NewRepositoryService(testConnector{
					requestDoer: func(req *http.Request) (*http.Response, error) {
						return nil, nil
					},
					requestMaker: func(ctx context.Context, method, path, query string, body interface{}) (*http.Request, error) {
						return nil, fmt.Errorf("error creating request")
					},
				})
				return client, nil
			},
			wantErr: true,
			Err:     fmt.Errorf("error creating request"),
		},
		{
			name: "when the http request fails",
			args: args{
				ctx:       context.Background(),
				workspace: "workspace-uuid",
				repoSlug:  "repo-slug",
			},
			on: func(t *testing.T) (*RepositoryService, *models.ResponseScheme) {
				client := NewRepositoryService(testConnector{
					requestDoer: func(req *http.Request) (*http.Response, error) {
						return nil, fmt.Errorf("error making request")
					},
					requestMaker: func(ctx context.Context, method, path, query string, body interface{}) (*http.Request, error) {
						return &http.Request{}, nil
					},
				})
				return client, nil
			},
			wantErr: true,
			Err:     fmt.Errorf("error making request"),
		},
		{
			name: "when the request is successful",
			args: args{
				ctx:       context.Background(),
				workspace: "workspace-uuid",
				repoSlug:  "repo-slug",
			},
			on: func(t *testing.T) (*RepositoryService, *models.ResponseScheme) {
				client := NewRepositoryService(testConnector{
					requestDoer: func(req *http.Request) (*http.Response, error) {
						resp := &http.Response{
							StatusCode: http.StatusOK,
							Body:       nil,
						}
						return resp, nil
					},
					requestMaker: func(ctx context.Context, method, path, query string, body interface{}) (*http.Request, error) {
						assert.Equal(t, "GET", method)
						assert.Equal(t, "2.0/repositories/workspace-uuid/repo-slug/branch-restrictions", path)
						return &http.Request{}, nil
					},
				})

				return client, &models.ResponseScheme{
					Response: &http.Response{
						StatusCode: http.StatusOK,
					},
				}
			},
			wantErr: false,
			Err:     nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			client, _ := testCase.on(t)

			gotResult, gotResponse, err := client.ListBranchRestrictions(testCase.args.ctx, testCase.args.workspace, testCase.args.repoSlug, testCase.args.pageOptions)

			if testCase.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, testCase.Err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}
		})
	}
}

func TestRepositoryService_ListDefaultReviewers(t *testing.T) {
	type args struct {
		ctx       context.Context
		workspace string
		repoSlug  string
	}

	testCases := []struct {
		name    string
		args    args
		on      func(*testing.T) (*RepositoryService, *models.ResponseScheme)
		wantErr bool
		Err     error
	}{
		{
			name: "when the workspace is not provided",
			args: args{
				ctx:       context.Background(),
				workspace: "",
				repoSlug:  "repo-slug",
			},
			on: func(t *testing.T) (*RepositoryService, *models.ResponseScheme) {
				client := NewRepositoryService(nil)
				return client, nil
			},
			wantErr: true,
			Err:     models.ErrNoWorkspace,
		},
		{
			name: "when the repository slug is not provided",
			args: args{
				ctx:       context.Background(),
				workspace: "workspace-uuid",
				repoSlug:  "",
			},
			on: func(t *testing.T) (*RepositoryService, *models.ResponseScheme) {
				client := NewRepositoryService(nil)
				return client, nil
			},
			wantErr: true,
			Err:     models.ErrNoRepository,
		},
		{
			name: "when the http request cannot be created",
			args: args{
				ctx:       context.Background(),
				workspace: "workspace-uuid",
				repoSlug:  "repo-slug",
			},
			on: func(t *testing.T) (*RepositoryService, *models.ResponseScheme) {
				client := NewRepositoryService(testConnector{
					requestDoer: func(req *http.Request) (*http.Response, error) {
						return nil, nil
					},
					requestMaker: func(ctx context.Context, method, path, query string, body interface{}) (*http.Request, error) {
						return nil, fmt.Errorf("error creating request")
					},
				})
				return client, nil
			},
			wantErr: true,
			Err:     fmt.Errorf("error creating request"),
		},
		{
			name: "when the http request fails",
			args: args{
				ctx:       context.Background(),
				workspace: "workspace-uuid",
				repoSlug:  "repo-slug",
			},
			on: func(t *testing.T) (*RepositoryService, *models.ResponseScheme) {
				client := NewRepositoryService(testConnector{
					requestDoer: func(req *http.Request) (*http.Response, error) {
						return nil, fmt.Errorf("error making request")
					},
					requestMaker: func(ctx context.Context, method, path, query string, body interface{}) (*http.Request, error) {
						return &http.Request{}, nil
					},
				})
				return client, nil
			},
			wantErr: true,
			Err:     fmt.Errorf("error making request"),
		},
		{
			name: "when the request is successful",
			args: args{
				ctx:       context.Background(),
				workspace: "workspace-uuid",
				repoSlug:  "repo-slug",
			},
			on: func(t *testing.T) (*RepositoryService, *models.ResponseScheme) {
				client := NewRepositoryService(testConnector{
					requestDoer: func(req *http.Request) (*http.Response, error) {
						resp := &http.Response{
							StatusCode: http.StatusOK,
							Body:       nil,
						}
						return resp, nil
					},
					requestMaker: func(ctx context.Context, method, path, query string, body interface{}) (*http.Request, error) {
						assert.Equal(t, "GET", method)
						assert.Equal(t, "2.0/repositories/workspace-uuid/repo-slug/default-reviewers", path)
						return &http.Request{}, nil
					},
				})

				return client, &models.ResponseScheme{
					Response: &http.Response{
						StatusCode: http.StatusOK,
					},
				}
			},
			wantErr: false,
			Err:     nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			client, _ := testCase.on(t)

			gotResult, gotResponse, err := client.ListDefaultReviewers(testCase.args.ctx, testCase.args.workspace, testCase.args.repoSlug)

			if testCase.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, testCase.Err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}
		})
	}
}

func TestRepositoryService_ListPullRequests(t *testing.T) {
	type args struct {
		ctx       context.Context
		workspace string
		repoSlug  string
	}

	testCases := []struct {
		name    string
		args    args
		on      func(*testing.T) (*RepositoryService, *models.ResponseScheme)
		wantErr bool
		Err     error
	}{
		{
			name: "when the workspace is not provided",
			args: args{
				ctx:       context.Background(),
				workspace: "",
				repoSlug:  "repo-slug",
			},
			on: func(t *testing.T) (*RepositoryService, *models.ResponseScheme) {
				client := NewRepositoryService(nil)
				return client, nil
			},
			wantErr: true,
			Err:     models.ErrNoWorkspace,
		},
		{
			name: "when the repository slug is not provided",
			args: args{
				ctx:       context.Background(),
				workspace: "workspace-uuid",
				repoSlug:  "",
			},
			on: func(t *testing.T) (*RepositoryService, *models.ResponseScheme) {
				client := NewRepositoryService(nil)
				return client, nil
			},
			wantErr: true,
			Err:     models.ErrNoRepository,
		},
		{
			name: "when the http request cannot be created",
			args: args{
				ctx:       context.Background(),
				workspace: "workspace-uuid",
				repoSlug:  "repo-slug",
			},
			on: func(t *testing.T) (*RepositoryService, *models.ResponseScheme) {
				client := NewRepositoryService(testConnector{
					requestDoer: func(req *http.Request) (*http.Response, error) {
						return nil, nil
					},
					requestMaker: func(ctx context.Context, method, path, query string, body interface{}) (*http.Request, error) {
						return nil, fmt.Errorf("error creating request")
					},
				})
				return client, nil
			},
			wantErr: true,
			Err:     fmt.Errorf("error creating request"),
		},
		{
			name: "when the http request fails",
			args: args{
				ctx:       context.Background(),
				workspace: "workspace-uuid",
				repoSlug:  "repo-slug",
			},
			on: func(t *testing.T) (*RepositoryService, *models.ResponseScheme) {
				client := NewRepositoryService(testConnector{
					requestDoer: func(req *http.Request) (*http.Response, error) {
						return nil, fmt.Errorf("error making request")
					},
					requestMaker: func(ctx context.Context, method, path, query string, body interface{}) (*http.Request, error) {
						return &http.Request{}, nil
					},
				})
				return client, nil
			},
			wantErr: true,
			Err:     fmt.Errorf("error making request"),
		},
		{
			name: "when the request is successful",
			args: args{
				ctx:       context.Background(),
				workspace: "workspace-uuid",
				repoSlug:  "repo-slug",
			},
			on: func(t *testing.T) (*RepositoryService, *models.ResponseScheme) {
				client := NewRepositoryService(testConnector{
					requestDoer: func(req *http.Request) (*http.Response, error) {
						resp := &http.Response{
							StatusCode: http.StatusOK,
							Body:       nil,
						}
						return resp, nil
					},
					requestMaker: func(ctx context.Context, method, path, query string, body interface{}) (*http.Request, error) {
						assert.Equal(t, "GET", method)
						assert.Equal(t, "2.0/repositories/workspace-uuid/repo-slug/pullrequests?state=OPEN%2CMERGED%2CDECLINED%2CSUPERSEDED", path)

						// Verify query parameters
						u, err := url.Parse(query)
						assert.NoError(t, err)
						assert.Equal(t, "", u.Query().Get("state"))

						return &http.Request{}, nil
					},
				})

				return client, &models.ResponseScheme{
					Response: &http.Response{
						StatusCode: http.StatusOK,
					},
				}
			},
			wantErr: false,
			Err:     nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			client, _ := testCase.on(t)

			gotResult, gotResponse, err := client.ListPullRequests(testCase.args.ctx, testCase.args.workspace, testCase.args.repoSlug)

			if testCase.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, testCase.Err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}
		})
	}
}

// testConnector is a mock connector for testing
type testConnector struct {
	requestMaker func(ctx context.Context, method, path, query string, body interface{}) (*http.Request, error)
	requestDoer  func(req *http.Request) (*http.Response, error)
}

func (t testConnector) NewRequest(ctx context.Context, method, path, query string, body interface{}) (*http.Request, error) {
	return t.requestMaker(ctx, method, path, query, body)
}

func (t testConnector) Call(req *http.Request, v interface{}) (*models.ResponseScheme, error) {
	resp, err := t.requestDoer(req)
	if err != nil {
		return nil, err
	}
	return &models.ResponseScheme{
		Response: resp,
	}, nil
}
