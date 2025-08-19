package internal

import (
	"context"
	"errors"
	"net/http"
	"testing"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_internalOrgTokenImpl_Gets(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx   context.Context
		orgID string
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name: "when the parameters are correct",
			args: args{
				ctx:   context.Background(),
				orgID: "org-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"admin/api-access/v1/orgs/org-id-sample/api-tokens",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					mock.Anything).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the organization id is not provided",
			args: args{
				ctx:   context.Background(),
				orgID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminOrganization,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:   context.Background(),
				orgID: "org-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"admin/api-access/v1/orgs/org-id-sample/api-tokens",
					"",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newOrgTokenService := NewOrgTokenService(testCase.fields.c)

			gotResult, gotResponse, err := newOrgTokenService.Gets(testCase.args.ctx, testCase.args.orgID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}

func Test_internalOrgKeyImpl_Gets(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx   context.Context
		orgID string
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name: "when the parameters are correct",
			args: args{
				ctx:   context.Background(),
				orgID: "org-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"admin/api-access/v1/orgs/org-id-sample/api-keys",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					mock.Anything).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the organization id is not provided",
			args: args{
				ctx:   context.Background(),
				orgID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminOrganization,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:   context.Background(),
				orgID: "org-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"admin/api-access/v1/orgs/org-id-sample/api-keys",
					"",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newOrgKeyService := NewOrgKeyService(testCase.fields.c)

			gotResult, gotResponse, err := newOrgKeyService.Gets(testCase.args.ctx, testCase.args.orgID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}
