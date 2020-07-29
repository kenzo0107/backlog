package backlog

import (
	"context"
	"fmt"
)

// Version : milestone
type Version struct {
	ID             *int    `json:"id,omitempty"`
	ProjectID      *int    `json:"projectId,omitempty"`
	Name           *string `json:"name,omitempty"`
	Description    *string `json:"description,omitempty"`
	StartDate      *string `json:"startDate,omitempty"`      // yyyy-MM-dd
	ReleaseDueDate *string `json:"releaseDueDate,omitempty"` // yyyy-MM-dd
	Archived       *bool   `json:"archived,omitempty"`
	DisplayOrder   *int    `json:"displayOrder,omitempty"`
}

// GetVersions returns the list of versions in a project
func (c *Client) GetVersions(projectIDOrKey interface{}) ([]*Version, error) {
	return c.GetVersionsContext(context.Background(), projectIDOrKey)
}

// GetVersionsContext returns a version of a project with context
func (c *Client) GetVersionsContext(ctx context.Context, projectIDOrKey interface{}) ([]*Version, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/versions", projectIDOrKey)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var versions []*Version
	if err := c.Do(ctx, req, &versions); err != nil {
		return nil, err
	}
	return versions, nil
}

// CreateVersion creates a versions (milestone) of a project
func (c *Client) CreateVersion(projectIDOrKey interface{}, input *CreateVersionInput) (*Version, error) {
	return c.CreateVersionContext(context.Background(), projectIDOrKey, input)
}

// CreateVersionContext creates a versions (milestone) of a project with Context
func (c *Client) CreateVersionContext(ctx context.Context, projectIDOrKey interface{}, input *CreateVersionInput) (*Version, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/versions", projectIDOrKey)

	req, err := c.NewRequest("POST", u, input)
	if err != nil {
		return nil, err
	}

	version := new(Version)
	if err := c.Do(ctx, req, &version); err != nil {
		return nil, err
	}
	return version, nil
}

// UpdateVersion updates a versions (milestone) of a project
func (c *Client) UpdateVersion(projectIDOrKey interface{}, versionID int, input *UpdateVersionInput) (*Version, error) {
	return c.UpdateVersionContext(context.Background(), projectIDOrKey, versionID, input)
}

// UpdateVersionContext updates a versions (milestone) of a project with Context
func (c *Client) UpdateVersionContext(ctx context.Context, projectIDOrKey interface{}, versionID int, input *UpdateVersionInput) (*Version, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/versions/%v", projectIDOrKey, versionID)

	req, err := c.NewRequest("PATCH", u, input)
	if err != nil {
		return nil, err
	}

	version := new(Version)
	if err := c.Do(ctx, req, &version); err != nil {
		return nil, err
	}
	return version, nil
}

// DeleteVersion deletes a versions (milestone) of a project
func (c *Client) DeleteVersion(projectIDOrKey interface{}, versionID int) (*Version, error) {
	return c.DeleteVersionContext(context.Background(), projectIDOrKey, versionID)
}

// DeleteVersionContext deletes a versions (milestone) of a project with Context
func (c *Client) DeleteVersionContext(ctx context.Context, projectIDOrKey interface{}, versionID int) (*Version, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/versions/%v", projectIDOrKey, versionID)

	req, err := c.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	version := new(Version)
	if err := c.Do(ctx, req, &version); err != nil {
		return nil, err
	}
	return version, nil
}

// CreateVersionInput specifies parameters to the CreateVersion method.
type CreateVersionInput struct {
	Name           *string `json:"name"`
	Description    *string `json:"description,omitempty"`
	StartDate      *string `json:"startDate,omitempty"`
	ReleaseDueDate *string `json:"releaseDueDate,omitempty"`
}

// UpdateVersionInput specifies parameters to the UpdateVersion method.
type UpdateVersionInput struct {
	Name           *string `json:"name"`
	Description    *string `json:"description,omitempty"`
	StartDate      *string `json:"startDate,omitempty"`
	ReleaseDueDate *string `json:"releaseDueDate,omitempty"`
	Archived       *bool   `json:"archived,omitempty"`
}
