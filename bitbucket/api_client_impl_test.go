package bitbucket

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/ctreminiom/go-atlassian/v2/bitbucket/internal"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service/common"
	"github.com/ctreminiom/go-atlassian/v2/service/mocks"
)

func TestClient_Call(t *testing.T) {

	expectedResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader("Hello, world!")),
		Request: &http.Request{
			Method: http.MethodGet,
			URL:    &url.URL{},
		},
	}

	badRequestResponse := &http.Response{
		StatusCode: http.StatusBadRequest,
		Body:       io.NopCloser(strings.NewReader("Hello, world!")),
		Request: &http.Request{
			Method: http.MethodGet,
			URL:    &url.URL{},
		},
	}

	internalServerResponse := &http.Response{
		StatusCode: http.StatusInternalServerError,
		Body:       io.NopCloser(strings.NewReader("Hello, world!")),
		Request: &http.Request{
			Method: http.MethodGet,
			URL:    &url.URL{},
		},
	}

	unauthorizedResponse := &http.Response{
		StatusCode: http.StatusUnauthorized,
		Body:       io.NopCloser(strings.NewReader("Hello, world!")),
		Request: &http.Request{
			Method: http.MethodGet,
			URL:    &url.URL{},
		},
	}

	notFoundResponse := &http.Response{
		StatusCode: http.StatusNotFound,
		Body:       io.NopCloser(strings.NewReader("Hello, world!")),
		Request: &http.Request{
			Method: http.MethodGet,
			URL:    &url.URL{},
		},
	}

	// Create test requests
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "http://test.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 0)
	defer cancel()
	reqWithTimeout, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://test.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	type fields struct {
		HTTP common.HTTPClient
		Site *url.URL
		Auth common.Authentication
	}
	type args struct {
		request   *http.Request
		structure interface{}
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		want    *model.ResponseScheme
		wantErr bool
		Err     error
	}{
		{
			name: "when the parameters are correct",
			on: func(fields *fields) {
				client := mocks.NewHTTPClient(t)

				client.On("Do", mock.AnythingOfType("*http.Request")).
					Return(expectedResponse, nil)

				fields.HTTP = client
			},
			args: args{
				request:   req,
				structure: nil,
			},
			want: &model.ResponseScheme{
				Response: expectedResponse,
				Code:     http.StatusOK,
				Method:   http.MethodGet,
				Bytes:    *bytes.NewBufferString("Hello, world!"),
			},
			wantErr: false,
		},

		{
			name: "when the response status is a bad request",
			on: func(fields *fields) {
				client := mocks.NewHTTPClient(t)

				client.On("Do", mock.AnythingOfType("*http.Request")).
					Return(badRequestResponse, nil)

				fields.HTTP = client
			},
			args: args{
				request:   req,
				structure: nil,
			},
			want: &model.ResponseScheme{
				Response: badRequestResponse,
				Code:     http.StatusBadRequest,
				Method:   http.MethodGet,
				Bytes:    *bytes.NewBufferString("Hello, world!"),
			},
			wantErr: true,
			Err:     model.ErrBadRequest,
		},

		{
			name: "when the response status is an internal service error",
			on: func(fields *fields) {
				client := mocks.NewHTTPClient(t)

				client.On("Do", mock.AnythingOfType("*http.Request")).
					Return(internalServerResponse, nil)

				fields.HTTP = client
			},
			args: args{
				request:   req,
				structure: nil,
			},
			want: &model.ResponseScheme{
				Response: internalServerResponse,
				Code:     http.StatusInternalServerError,
				Method:   http.MethodGet,
				Bytes:    *bytes.NewBufferString("Hello, world!"),
			},
			wantErr: true,
			Err:     model.ErrInternal,
		},

		{
			name: "when the response status is a not found",
			on: func(fields *fields) {
				client := mocks.NewHTTPClient(t)

				client.On("Do", mock.AnythingOfType("*http.Request")).
					Return(notFoundResponse, nil)

				fields.HTTP = client
			},
			args: args{
				request:   req,
				structure: nil,
			},
			want: &model.ResponseScheme{
				Response: notFoundResponse,
				Code:     http.StatusNotFound,
				Method:   http.MethodGet,
				Bytes:    *bytes.NewBufferString("Hello, world!"),
			},
			wantErr: true,
			Err:     model.ErrNotFound,
		},

		{
			name: "when the response status is unauthorized",
			on: func(fields *fields) {
				client := mocks.NewHTTPClient(t)

				client.On("Do", mock.AnythingOfType("*http.Request")).
					Return(unauthorizedResponse, nil)

				fields.HTTP = client
			},
			args: args{
				request:   req,
				structure: nil,
			},
			want: &model.ResponseScheme{
				Response: unauthorizedResponse,
				Code:     http.StatusUnauthorized,
				Method:   http.MethodGet,
				Bytes:    *bytes.NewBufferString("Hello, world!"),
			},
			wantErr: true,
			Err:     model.ErrUnauthorized,
		},
		{
			name: "when rate limit is hit and retry succeeds",
			on: func(fields *fields) {
				client := mocks.NewHTTPClient(t)

				// First call returns rate limit
				rateLimitResponse := &http.Response{
					StatusCode: http.StatusTooManyRequests,
					Body:       io.NopCloser(strings.NewReader("Rate limit exceeded")),
					Request: &http.Request{
						Method: http.MethodGet,
						URL:    &url.URL{},
					},
				}

				// Second call succeeds
				client.On("Do", mock.AnythingOfType("*http.Request")).
					Return(rateLimitResponse, nil).
					Times(1)

				client.On("Do", mock.AnythingOfType("*http.Request")).
					Return(expectedResponse, nil).
					Times(1)

				fields.HTTP = client
			},
			args: args{
				request:   req,
				structure: nil,
			},
			want: &model.ResponseScheme{
				Response: expectedResponse,
				Code:     http.StatusOK,
				Method:   http.MethodGet,
				Bytes:    *bytes.NewBufferString("Hello, world!"),
			},
			wantErr: false,
		},

		{
			name: "when rate limit is hit and max retries exceeded",
			on: func(fields *fields) {
				client := mocks.NewHTTPClient(t)

				// Return rate limit response 6 times (max retries + 1)
				rateLimitResponse := &http.Response{
					StatusCode: http.StatusTooManyRequests,
					Body:       io.NopCloser(strings.NewReader("Rate limit exceeded")),
					Request: &http.Request{
						Method: http.MethodGet,
						URL:    &url.URL{},
					},
				}

				// Set up expectations for all retries
				client.On("Do", mock.AnythingOfType("*http.Request")).
					Return(rateLimitResponse, nil).
					Times(6)

				fields.HTTP = client
			},
			args: args{
				request:   req,
				structure: nil,
			},
			want: &model.ResponseScheme{
				Response: &http.Response{
					StatusCode: http.StatusTooManyRequests,
					Body:       io.NopCloser(strings.NewReader("Rate limit exceeded")),
					Request: &http.Request{
						Method: http.MethodGet,
						URL:    &url.URL{},
					},
				},
				Code:   http.StatusTooManyRequests,
				Method: http.MethodGet,
				Bytes:  *bytes.NewBufferString("Rate limit exceeded"),
			},
			wantErr: true,
			Err:     model.ErrInvalidStatusCode,
		},

		{
			name: "when context is cancelled during rate limit retry",
			on: func(fields *fields) {
				client := mocks.NewHTTPClient(t)

				// Return rate limit response
				rateLimitResponse := &http.Response{
					StatusCode: http.StatusTooManyRequests,
					Body:       io.NopCloser(strings.NewReader("Rate limit exceeded")),
					Request: &http.Request{
						Method: http.MethodGet,
						URL:    &url.URL{},
					},
				}

				client.On("Do", mock.AnythingOfType("*http.Request")).
					Return(rateLimitResponse, nil).
					Times(1)

				fields.HTTP = client
			},
			args: args{
				request:   reqWithTimeout,
				structure: nil,
			},
			want:    nil,
			wantErr: true,
			Err:     context.Canceled,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			config := &ClientConfig{
				MaxRetries:        5,
				InitialRetryDelay: 1000,
				MaxRetryDelay:     10000,
			}

			c := &Client{
				HTTP:              testCase.fields.HTTP,
				Site:              testCase.fields.Site,
				MaxRetries:        config.MaxRetries,
				InitialRetryDelay: config.InitialRetryDelay,
				MaxRetryDelay:     config.MaxRetryDelay,
			}

			got, err := c.Call(testCase.args.request, testCase.args.structure)

			if testCase.wantErr {
				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.EqualError(t, err, testCase.Err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
				assert.Equal(t, testCase.want.Code, got.Code)
				assert.Equal(t, testCase.want.Method, got.Method)
				assert.Equal(t, testCase.want.Bytes.String(), got.Bytes.String())
			}

		})
	}
}

func TestClient_NewRequest(t *testing.T) {

	authMocked := internal.NewAuthenticationService(nil)
	authMocked.SetBasicAuth("mail", "token")
	authMocked.SetUserAgent("firefox")
	authMocked.SetBearerToken("token_sample")

	siteAsURL, err := url.Parse("https://ctreminiom.atlassian.net")
	if err != nil {
		t.Fatal(err)
	}

	requestMocked, err := http.NewRequestWithContext(context.TODO(),
		http.MethodGet,
		"https://ctreminiom.atlassian.net/rest/2/issue/attachment",
		bytes.NewReader([]byte("Hello World")),
	)

	if err != nil {
		t.Fatal(err)
	}

	requestMocked.Header.Set("Accept", "application/json")
	requestMocked.Header.Set("Content-Type", "application/json")

	type fields struct {
		HTTP common.HTTPClient
		Auth common.Authentication
		Site *url.URL
	}

	type args struct {
		ctx         context.Context
		method      string
		urlStr      string
		contentType string
		body        interface{}
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Request
		wantErr bool
	}{
		{
			name: "when the parameters are correct",
			fields: fields{
				HTTP: http.DefaultClient,
				Auth: authMocked,
				Site: siteAsURL,
			},
			args: args{
				ctx:         context.Background(),
				method:      http.MethodGet,
				urlStr:      "rest/2/issue/attachment",
				contentType: "",
				body:        bytes.NewReader([]byte("Hello World")),
			},
			want:    requestMocked,
			wantErr: false,
		},

		{
			name: "when the url cannot be parsed",
			fields: fields{
				HTTP: http.DefaultClient,
				Auth: internal.NewAuthenticationService(nil),
				Site: siteAsURL,
			},
			args: args{
				ctx:    context.Background(),
				method: http.MethodGet,
				urlStr: " https://zhidao.baidu.com/special/view?id=49105a24626975510000&preview=1",
				body:   bytes.NewReader([]byte("Hello World")),
			},
			want:    nil,
			wantErr: true,
		},

		{
			name: "when the content type is provided",
			fields: fields{
				HTTP: http.DefaultClient,
				Auth: authMocked,
				Site: siteAsURL,
			},
			args: args{
				ctx:         context.Background(),
				method:      http.MethodGet,
				urlStr:      "rest/2/issue/attachment",
				contentType: "type_sample",
				body:        bytes.NewReader([]byte("Hello World")),
			},
			want:    requestMocked,
			wantErr: false,
		},

		{
			name: "when the request cannot be created",
			fields: fields{
				HTTP: http.DefaultClient,
				Auth: internal.NewAuthenticationService(nil),
				Site: siteAsURL,
			},
			args: args{
				ctx:    nil,
				method: http.MethodGet,
				urlStr: "rest/2/issue/attachment",
				body:   bytes.NewReader([]byte("Hello World")),
			},
			want:    requestMocked,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			config := &ClientConfig{
				MaxRetries:        5,
				InitialRetryDelay: 1000,
				MaxRetryDelay:     10000,
			}

			c := &Client{
				HTTP:              testCase.fields.HTTP,
				Auth:              testCase.fields.Auth,
				Site:              testCase.fields.Site,
				MaxRetries:        config.MaxRetries,
				InitialRetryDelay: config.InitialRetryDelay,
				MaxRetryDelay:     config.MaxRetryDelay,
			}

			got, err := c.NewRequest(
				testCase.args.ctx,
				testCase.args.method,
				testCase.args.urlStr,
				testCase.args.contentType,
				testCase.args.body,
			)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEqual(t, got, nil)
			}

		})
	}
}

func TestClient_processResponse(t *testing.T) {

	expectedJSONResponse := `
	{
	  "id": 4,
	  "self": "https://ctreminiom.atlassian.net/rest/agile/1.0/board/4",
	  "name": "KP - Scrum",
	  "type": "scrum"
	}`

	expectedResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(expectedJSONResponse)),
		Request: &http.Request{
			Method: http.MethodGet,
			URL:    &url.URL{},
		},
	}

	type fields struct {
		HTTP           common.HTTPClient
		Site           *url.URL
		Authentication common.Authentication
	}
	type args struct {
		response  *http.Response
		structure interface{}
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		want    *model.ResponseScheme
		wantErr bool
		Err     error
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				response:  expectedResponse,
				structure: model.BoardScheme{},
			},
			want: &model.ResponseScheme{
				Response: expectedResponse,
				Code:     http.StatusOK,
				Method:   http.MethodGet,
				Bytes:    *bytes.NewBufferString(expectedJSONResponse),
			},
			wantErr: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			config := &ClientConfig{
				MaxRetries:        5,
				InitialRetryDelay: 1000,
				MaxRetryDelay:     10000,
			}

			c := &Client{
				HTTP:              testCase.fields.HTTP,
				Site:              testCase.fields.Site,
				Auth:              testCase.fields.Authentication,
				MaxRetries:        config.MaxRetries,
				InitialRetryDelay: config.InitialRetryDelay,
				MaxRetryDelay:     config.MaxRetryDelay,
			}

			got, err := c.processResponse(testCase.args.response, testCase.args.structure)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
				assert.EqualError(t, err, testCase.Err.Error())

			} else {
				assert.NoError(t, err)
				assert.NotEqual(t, got, nil)
			}

		})
	}
}

func TestNew(t *testing.T) {
	mockClient, err := New(http.DefaultClient, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	mockClient.Auth.SetBasicAuth("test", "test")
	mockClient.Auth.SetUserAgent("aaa")

	// Create a client with custom config
	customConfig := &ClientConfig{
		MaxRetries:        10,
		InitialRetryDelay: 2000,
		MaxRetryDelay:     20000,
	}
	customClient, err := New(http.DefaultClient, "", customConfig)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		httpClient common.HTTPClient
		site       string
		config     *ClientConfig
	}

	testCases := []struct {
		name    string
		args    args
		want    *Client
		wantErr bool
		Err     error
	}{
		{
			name: "when using default config",
			args: args{
				httpClient: http.DefaultClient,
				site:       "",
				config:     nil,
			},
			want:    mockClient,
			wantErr: false,
		},
		{
			name: "when using custom config",
			args: args{
				httpClient: http.DefaultClient,
				site:       "",
				config:     customConfig,
			},
			want:    customClient,
			wantErr: false,
		},
		{
			name: "when the site url is not valid",
			args: args{
				httpClient: http.DefaultClient,
				site:       " https://zhidao.baidu.com/special/view?id=sd&preview=1",
				config:     nil,
			},
			want:    nil,
			wantErr: true,
			Err:     errors.New("parse \" https://zhidao.baidu.com/special/view?id=sd&preview=1/\": first path segment in URL cannot contain colon"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			gotClient, err := New(testCase.args.httpClient, testCase.args.site, testCase.args.config)

			if testCase.wantErr {
				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)
				assert.EqualError(t, err, testCase.Err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, gotClient)
				assert.Equal(t, testCase.want.MaxRetries, gotClient.MaxRetries)
				assert.Equal(t, testCase.want.InitialRetryDelay, gotClient.InitialRetryDelay)
				assert.Equal(t, testCase.want.MaxRetryDelay, gotClient.MaxRetryDelay)
			}
		})
	}
}
