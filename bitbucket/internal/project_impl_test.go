package internal

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/mocks"
)

func Test_internalProjectServiceImpl_ListExplicitProjectUserPermissions(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx        context.Context
		workspace  string
		projectKey string
		opts       *model.PageOptions
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
				ctx:        context.Background(),
				workspace:  "workspace",
				projectKey: "projectKey",
				opts:       &model.PageOptions{Page: 1, PageLen: 10},
			},
			on: func(fields *fields) {
				client := mocks.NewConnector(t)
				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"2.0/workspaces/workspace/projects/projectKey/permissions-config/users?page=1&pagelen=10",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ProjectUserMembershipPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the workspace is not provided",
			args: args{
				ctx:        context.Background(),
				workspace:  "",
				projectKey: "projectKey",
				opts:       &model.PageOptions{Page: 1, PageLen: 10},
			},
			wantErr: true,
			Err:     model.ErrNoWorkspace,
		},

		{
			name: "when the project key is not provided",
			args: args{
				ctx:        context.Background(),
				workspace:  "workspace",
				projectKey: "",
				opts:       &model.PageOptions{Page: 1, PageLen: 10},
			},
			wantErr: true,
			Err:     model.ErrNoProjectKey,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewProjectService(testCase.fields.c)

			got, got1, err := newService.ListExplicitProjectUserPermissions(testCase.args.ctx, testCase.args.workspace, testCase.args.projectKey, testCase.args.opts)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, got, nil)
				assert.NotEqual(t, got1, nil)
			}

		})
	}
}

func Test_internalProjectServiceImpl_ListExplicitProjectGroupPermissions(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx        context.Context
		workspace  string
		projectKey string
		opts       *model.PageOptions
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
				ctx:        context.Background(),
				workspace:  "workspace",
				projectKey: "projectKey",
				opts:       &model.PageOptions{Page: 1, PageLen: 10},
			},
			on: func(fields *fields) {
				client := mocks.NewConnector(t)
				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"2.0/workspaces/workspace/projects/projectKey/permissions-config/groups?page=1&pagelen=10",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ProjectGroupMembershipPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the workspace is not provided",
			args: args{
				ctx:        context.Background(),
				workspace:  "",
				projectKey: "projectKey",
				opts:       &model.PageOptions{Page: 1, PageLen: 10},
			},
			wantErr: true,
			Err:     model.ErrNoWorkspace,
		},

		{
			name: "when the project key is not provided",
			args: args{
				ctx:        context.Background(),
				workspace:  "workspace",
				projectKey: "",
				opts:       &model.PageOptions{Page: 1, PageLen: 10},
			},
			wantErr: true,
			Err:     model.ErrNoProjectKey,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewProjectService(testCase.fields.c)

			got, got1, err := newService.ListExplicitProjectGroupPermissions(testCase.args.ctx, testCase.args.workspace, testCase.args.projectKey, testCase.args.opts)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, got, nil)
				assert.NotEqual(t, got1, nil)
			}

		})
	}
}
