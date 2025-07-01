package backlog

import (
	"context"
	"fmt"
)

// GitRepository : Git repository
type GitRepository struct {
	ID           *int       `json:"id,omitempty"`
	ProjectID    *int       `json:"projectId,omitempty"`
	Name         *string    `json:"name,omitempty"`
	Description  *string    `json:"description,omitempty"`
	HookURL      *string    `json:"hookUrl,omitempty"`
	HTTPURL      *string    `json:"httpUrl,omitempty"`
	SSHURL       *string    `json:"sshUrl,omitempty"`
	DisplayOrder *int       `json:"displayOrder,omitempty"`
	PushedAt     *Timestamp `json:"pushedAt,omitempty"`
	CreatedUser  *User      `json:"createdUser,omitempty"`
	Created      *Timestamp `json:"created,omitempty"`
	UpdatedUser  *User      `json:"updatedUser,omitempty"`
	Updated      *Timestamp `json:"updated,omitempty"`
}

// ResponseGitRepositories : response for git repositories
type ResponseGitRepositories []*GitRepository

// GetGitRepositories returns git repositories
func (c *Client) GetGitRepositories(projectIDOrKey interface{}) (*ResponseGitRepositories, error) {
	return c.GetGitRepositoriesContext(context.Background(), projectIDOrKey)
}

// GetGitRepositoriesContext returns git repositories
func (c *Client) GetGitRepositoriesContext(ctx context.Context, projectIDOrKey interface{}) (*ResponseGitRepositories, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/git/repositories", projectIDOrKey)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	responseGitRepositories := new(ResponseGitRepositories)
	if err := c.Do(ctx, req, responseGitRepositories); err != nil {
		return nil, err
	}

	return responseGitRepositories, nil
}

// GetGitRepository returns git repository
func (c *Client) GetGitRepository(projectIDOrKey interface{}, repoIDOrName interface{}) (*GitRepository, error) {
	return c.GetGitRepositoryContext(context.Background(), projectIDOrKey, repoIDOrName)
}

// GetGitRepositoryContext returns git repository
func (c *Client) GetGitRepositoryContext(ctx context.Context, projectIDOrKey interface{}, repoIDOrName interface{}) (*GitRepository, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/git/repositories/%v", projectIDOrKey, repoIDOrName)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	gitRepository := new(GitRepository)
	if err := c.Do(ctx, req, gitRepository); err != nil {
		return nil, err
	}

	return gitRepository, nil
}
