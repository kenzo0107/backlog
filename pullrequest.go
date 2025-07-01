package backlog

import (
	"context"
	"fmt"
)

// PullRequest : pull request
type PullRequest struct {
	ID          *int    `json:"id,omitempty"`
	ProjectID   *int    `json:"projectId,omitempty"`
	RepositoryID *int   `json:"repositoryId,omitempty"`
	Number      *int    `json:"number,omitempty"`
	Summary     *string `json:"summary,omitempty"`
	Description *string `json:"description,omitempty"`
	Base        *string `json:"base,omitempty"`
	Branch      *string `json:"branch,omitempty"`
	Status      *Status `json:"status,omitempty"`
	Assignee    *User   `json:"assignee,omitempty"`
	Issue       *Issue  `json:"issue,omitempty"`
	BaseCommit  *string `json:"baseCommit,omitempty"`
	BranchCommit *string `json:"branchCommit,omitempty"`
	CloseAt     *Timestamp `json:"closeAt,omitempty"`
	MergeAt     *Timestamp `json:"mergeAt,omitempty"`
	CreatedUser *User      `json:"createdUser,omitempty"`
	Created     *Timestamp `json:"created,omitempty"`
	UpdatedUser *User      `json:"updatedUser,omitempty"`
	Updated     *Timestamp `json:"updated,omitempty"`
}

// PullRequestComment : pull request comment
type PullRequestComment struct {
	ID            *int    `json:"id,omitempty"`
	Content       *string `json:"content,omitempty"`
	ChangeLog     []*ChangeLog `json:"changeLog,omitempty"`
	CreatedUser   *User      `json:"createdUser,omitempty"`
	Created       *Timestamp `json:"created,omitempty"`
	UpdatedUser   *User      `json:"updatedUser,omitempty"`
	Updated       *Timestamp `json:"updated,omitempty"`
	Stars         []*Star `json:"stars,omitempty"`
	Notifications []*Notification `json:"notifications,omitempty"`
}

// GetPullRequestsOptions : options for GetPullRequests
type GetPullRequestsOptions struct {
	StatusID     []int `url:"statusId[],omitempty"`
	AssigneeID   []int `url:"assigneeId[],omitempty"`
	IssueID      []int `url:"issueId[],omitempty"`
	CreatedUserID []int `url:"createdUserId[],omitempty"`
	Offset       *int  `url:"offset,omitempty"`
	Count        *int  `url:"count,omitempty"`
}

// CreatePullRequestOptions : options for CreatePullRequest
type CreatePullRequestOptions struct {
	Summary     *string `json:"summary,omitempty"`
	Description *string `json:"description,omitempty"`
	Base        *string `json:"base,omitempty"`
	Branch      *string `json:"branch,omitempty"`
	IssueID     *int    `json:"issueId,omitempty"`
	AssigneeID  *int    `json:"assigneeId,omitempty"`
	NotifiedUserID []int `json:"notifiedUserId,omitempty"`
	AttachmentID []int `json:"attachmentId,omitempty"`
}

// UpdatePullRequestOptions : options for UpdatePullRequest
type UpdatePullRequestOptions struct {
	Summary     *string `json:"summary,omitempty"`
	Description *string `json:"description,omitempty"`
	IssueID     *int    `json:"issueId,omitempty"`
	AssigneeID  *int    `json:"assigneeId,omitempty"`
	NotifiedUserID []int `json:"notifiedUserId,omitempty"`
	Comment     *string `json:"comment,omitempty"`
}

// ResponsePullRequests : response for pull requests
type ResponsePullRequests []*PullRequest

// ResponsePullRequestComments : response for pull request comments
type ResponsePullRequestComments []*PullRequestComment

// ResponsePullRequestCount : response for pull request count
type ResponsePullRequestCount struct {
	Count *int `json:"count,omitempty"`
}

// GetPullRequests returns pull requests
func (c *Client) GetPullRequests(projectIDOrKey interface{}, repoIDOrName interface{}, options *GetPullRequestsOptions) (*ResponsePullRequests, error) {
	return c.GetPullRequestsContext(context.Background(), projectIDOrKey, repoIDOrName, options)
}

// GetPullRequestsContext returns pull requests
func (c *Client) GetPullRequestsContext(ctx context.Context, projectIDOrKey interface{}, repoIDOrName interface{}, options *GetPullRequestsOptions) (*ResponsePullRequests, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/git/repositories/%v/pullRequests", projectIDOrKey, repoIDOrName)

	if options != nil {
		s, err := c.AddOptions(u, options)
		if err != nil {
			return nil, err
		}
		u = s
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	responsePullRequests := new(ResponsePullRequests)
	if err := c.Do(ctx, req, responsePullRequests); err != nil {
		return nil, err
	}

	return responsePullRequests, nil
}

// GetPullRequestsCount returns pull requests count
func (c *Client) GetPullRequestsCount(projectIDOrKey interface{}, repoIDOrName interface{}, options *GetPullRequestsOptions) (*ResponsePullRequestCount, error) {
	return c.GetPullRequestsCountContext(context.Background(), projectIDOrKey, repoIDOrName, options)
}

// GetPullRequestsCountContext returns pull requests count
func (c *Client) GetPullRequestsCountContext(ctx context.Context, projectIDOrKey interface{}, repoIDOrName interface{}, options *GetPullRequestsOptions) (*ResponsePullRequestCount, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/git/repositories/%v/pullRequests/count", projectIDOrKey, repoIDOrName)

	if options != nil {
		s, err := c.AddOptions(u, options)
		if err != nil {
			return nil, err
		}
		u = s
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	responsePullRequestCount := new(ResponsePullRequestCount)
	if err := c.Do(ctx, req, responsePullRequestCount); err != nil {
		return nil, err
	}

	return responsePullRequestCount, nil
}

// GetPullRequest returns pull request
func (c *Client) GetPullRequest(projectIDOrKey interface{}, repoIDOrName interface{}, number int) (*PullRequest, error) {
	return c.GetPullRequestContext(context.Background(), projectIDOrKey, repoIDOrName, number)
}

// GetPullRequestContext returns pull request
func (c *Client) GetPullRequestContext(ctx context.Context, projectIDOrKey interface{}, repoIDOrName interface{}, number int) (*PullRequest, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/git/repositories/%v/pullRequests/%d", projectIDOrKey, repoIDOrName, number)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	pullRequest := new(PullRequest)
	if err := c.Do(ctx, req, pullRequest); err != nil {
		return nil, err
	}

	return pullRequest, nil
}

// CreatePullRequest creates pull request
func (c *Client) CreatePullRequest(projectIDOrKey interface{}, repoIDOrName interface{}, options *CreatePullRequestOptions) (*PullRequest, error) {
	return c.CreatePullRequestContext(context.Background(), projectIDOrKey, repoIDOrName, options)
}

// CreatePullRequestContext creates pull request
func (c *Client) CreatePullRequestContext(ctx context.Context, projectIDOrKey interface{}, repoIDOrName interface{}, options *CreatePullRequestOptions) (*PullRequest, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/git/repositories/%v/pullRequests", projectIDOrKey, repoIDOrName)

	req, err := c.NewRequest("POST", u, options)
	if err != nil {
		return nil, err
	}

	pullRequest := new(PullRequest)
	if err := c.Do(ctx, req, pullRequest); err != nil {
		return nil, err
	}

	return pullRequest, nil
}

// UpdatePullRequest updates pull request
func (c *Client) UpdatePullRequest(projectIDOrKey interface{}, repoIDOrName interface{}, number int, options *UpdatePullRequestOptions) (*PullRequest, error) {
	return c.UpdatePullRequestContext(context.Background(), projectIDOrKey, repoIDOrName, number, options)
}

// UpdatePullRequestContext updates pull request
func (c *Client) UpdatePullRequestContext(ctx context.Context, projectIDOrKey interface{}, repoIDOrName interface{}, number int, options *UpdatePullRequestOptions) (*PullRequest, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/git/repositories/%v/pullRequests/%d", projectIDOrKey, repoIDOrName, number)

	req, err := c.NewRequest("PATCH", u, options)
	if err != nil {
		return nil, err
	}

	pullRequest := new(PullRequest)
	if err := c.Do(ctx, req, pullRequest); err != nil {
		return nil, err
	}

	return pullRequest, nil
}

// GetPullRequestCommentsOptions : options for GetPullRequestComments
type GetPullRequestCommentsOptions struct {
	MinID *int `json:"minId,omitempty"`
	MaxID *int `json:"maxId,omitempty"`
	Count *int `json:"count,omitempty"`
	Order *string `json:"order,omitempty"`
}

// GetPullRequestComments returns pull request comments
func (c *Client) GetPullRequestComments(projectIDOrKey interface{}, repoIDOrName interface{}, number int, options *GetPullRequestCommentsOptions) (*ResponsePullRequestComments, error) {
	return c.GetPullRequestCommentsContext(context.Background(), projectIDOrKey, repoIDOrName, number, options)
}

// GetPullRequestCommentsContext returns pull request comments
func (c *Client) GetPullRequestCommentsContext(ctx context.Context, projectIDOrKey interface{}, repoIDOrName interface{}, number int, options *GetPullRequestCommentsOptions) (*ResponsePullRequestComments, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/git/repositories/%v/pullRequests/%d/comments", projectIDOrKey, repoIDOrName, number)

	if options != nil {
		s, err := c.AddOptions(u, options)
		if err != nil {
			return nil, err
		}
		u = s
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	responsePullRequestComments := new(ResponsePullRequestComments)
	if err := c.Do(ctx, req, responsePullRequestComments); err != nil {
		return nil, err
	}

	return responsePullRequestComments, nil
}

// ResponsePullRequestCommentsCount : response for pull request comments count
type ResponsePullRequestCommentsCount struct {
	Count *int `json:"count,omitempty"`
}

// GetPullRequestCommentsCount returns pull request comments count
func (c *Client) GetPullRequestCommentsCount(projectIDOrKey interface{}, repoIDOrName interface{}, number int) (*ResponsePullRequestCommentsCount, error) {
	return c.GetPullRequestCommentsCountContext(context.Background(), projectIDOrKey, repoIDOrName, number)
}

// GetPullRequestCommentsCountContext returns pull request comments count
func (c *Client) GetPullRequestCommentsCountContext(ctx context.Context, projectIDOrKey interface{}, repoIDOrName interface{}, number int) (*ResponsePullRequestCommentsCount, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/git/repositories/%v/pullRequests/%d/comments/count", projectIDOrKey, repoIDOrName, number)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	responseCount := new(ResponsePullRequestCommentsCount)
	if err := c.Do(ctx, req, responseCount); err != nil {
		return nil, err
	}

	return responseCount, nil
}
