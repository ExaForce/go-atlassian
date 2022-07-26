package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type CommentRichTextService struct {
	internalClient jira.CommentRichTextConnector
}

func (i *CommentRichTextService) Delete(ctx context.Context, issueKeyOrId, commentId string) (*model.ResponseScheme, error) {
	return i.internalClient.Delete(ctx, issueKeyOrId, commentId)
}

func (i *CommentRichTextService) Gets(ctx context.Context, issueKeyOrId, orderBy string, expand []string, startAt, maxResults int) (*model.IssueCommentPageSchemeV2, *model.ResponseScheme, error) {
	return i.internalClient.Gets(ctx, issueKeyOrId, orderBy, expand, startAt, maxResults)
}

func (i *CommentRichTextService) Get(ctx context.Context, issueKeyOrId, commentId string) (*model.IssueCommentSchemeV2, *model.ResponseScheme, error) {
	return i.internalClient.Get(ctx, issueKeyOrId, commentId)
}

func (i *CommentRichTextService) Add(ctx context.Context, issueKeyOrId string, payload *model.CommentPayloadSchemeV2, expand []string) (*model.IssueCommentSchemeV2, *model.ResponseScheme, error) {
	return i.internalClient.Add(ctx, issueKeyOrId, payload, expand)
}

type internalRichTextCommentImpl struct {
	c       service.Client
	version string
}

func (i *internalRichTextCommentImpl) Delete(ctx context.Context, issueKeyOrId, commentId string) (*model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, model.ErrNoIssueKeyOrIDError
	}

	if commentId == "" {
		return nil, model.ErrNoCommentIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/comment/%v", i.version, issueKeyOrId, commentId)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalRichTextCommentImpl) Gets(ctx context.Context, issueKeyOrId, orderBy string, expand []string, startAt, maxResults int) (*model.IssueCommentPageSchemeV2, *model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	if orderBy != "" {
		params.Add("orderBy", orderBy)
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/comment?%v", i.version, issueKeyOrId, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	comments := new(model.IssueCommentPageSchemeV2)
	response, err := i.c.Call(request, comments)
	if err != nil {
		return nil, response, err
	}

	return comments, response, nil
}

func (i *internalRichTextCommentImpl) Get(ctx context.Context, issueKeyOrId, commentId string) (*model.IssueCommentSchemeV2, *model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	if commentId == "" {
		return nil, nil, model.ErrNoCommentIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/comment/%v", i.version, issueKeyOrId, commentId)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	comment := new(model.IssueCommentSchemeV2)
	response, err := i.c.Call(request, comment)
	if err != nil {
		return nil, response, err
	}

	return comment, response, nil
}

func (i *internalRichTextCommentImpl) Add(ctx context.Context, issueKeyOrId string, payload *model.CommentPayloadSchemeV2, expand []string) (*model.IssueCommentSchemeV2, *model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	params := url.Values{}
	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/issue/%v/comment", i.version, issueKeyOrId))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint.String(), reader)
	if err != nil {
		return nil, nil, err
	}

	comment := new(model.IssueCommentSchemeV2)
	response, err := i.c.Call(request, comment)
	if err != nil {
		return nil, response, err
	}

	return comment, response, nil
}
