package backlog

import (
	"context"
	"fmt"
)

// Project : project
type Project struct {
	ID                                *int    `json:"id,omitempty"`
	ProjectKey                        *string `json:"projectKey,omitempty"`
	Name                              *string `json:"name,omitempty"`
	ChartEnabled                      *bool   `json:"chartEnabled,omitempty"`
	SubtaskingEnabled                 *bool   `json:"subtaskingEnabled,omitempty"`
	ProjectLeaderCanEditProjectLeader *bool   `json:"projectLeaderCanEditProjectLeader,omitempty"`
	TextFormattingRule                *string `json:"textFormattingRule,omitempty"`
	Archived                          *bool   `json:"archived,omitempty"`
	DisplayOrder                      *int    `json:"displayOrder,omitempty"`
}

// Status : the status of project
type Status struct {
	ID           *int    `json:"id,omitempty"`
	ProjectID    *int    `json:"projectId,omitempty"`
	Name         *string `json:"name,omitempty"`
	Color        *string `json:"color,omitempty"`
	DisplayOrder *int    `json:"displayOrder,omitempty"`
}

// GetProjects returns the list of projects
func (c *Client) GetProjects(opts *GetProjectsOptions) ([]*Project, error) {
	return c.GetProjectsContext(context.Background(), opts)
}

// GetProjectsContext returns the list of projects
func (c *Client) GetProjectsContext(ctx context.Context, opts *GetProjectsOptions) ([]*Project, error) {
	u, err := c.AddOptions("/api/v2/projects", opts)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	projects := []*Project{}
	if err := c.Do(ctx, req, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}

// GetProject returns a project
func (c *Client) GetProject(projectIDOrKey interface{}) (*Project, error) {
	return c.GetProjectContext(context.Background(), projectIDOrKey)
}

// GetProjectContext returns a project with context
func (c *Client) GetProjectContext(ctx context.Context, projectIDOrKey interface{}) (*Project, error) {
	u := fmt.Sprintf("/api/v2/projects/%v", projectIDOrKey)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	project := new(Project)
	if err := c.Do(ctx, req, &project); err != nil {
		return nil, err
	}
	return project, nil
}

// GetProjectStatuses returns the statuses of a project
func (c *Client) GetProjectStatuses(projectIDOrKey interface{}) ([]*Status, error) {
	return c.GetProjectStatusesContext(context.Background(), projectIDOrKey)
}

// GetProjectStatusesContext returns the statuses of a project with context
func (c *Client) GetProjectStatusesContext(ctx context.Context, projectIDOrKey interface{}) ([]*Status, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/statuses", projectIDOrKey)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	statuses := []*Status{}
	if err := c.Do(ctx, req, &statuses); err != nil {
		return nil, err
	}
	return statuses, nil
}

// CreateProject creates a project
func (c *Client) CreateProject(input *CreateProjectInput) (*Project, error) {
	return c.CreateProjectContext(context.Background(), input)
}

// CreateProjectContext creates a project with Context
func (c *Client) CreateProjectContext(ctx context.Context, input *CreateProjectInput) (*Project, error) {
	u := "/api/v2/projects"

	req, err := c.NewRequest("POST", u, input)
	if err != nil {
		return nil, err
	}

	project := new(Project)
	if err := c.Do(ctx, req, &project); err != nil {
		return nil, err
	}
	return project, nil
}

// UpdateProject updates a project
func (c *Client) UpdateProject(id int, input *UpdateProjectInput) (*Project, error) {
	return c.UpdateProjectContext(context.Background(), id, input)
}

// UpdateProjectContext updates a project with Context
func (c *Client) UpdateProjectContext(ctx context.Context, id int, input *UpdateProjectInput) (*Project, error) {
	u := fmt.Sprintf("/api/v2/projects/%v", id)

	req, err := c.NewRequest("PATCH", u, input)
	if err != nil {
		return nil, err
	}

	project := new(Project)
	if err := c.Do(ctx, req, &project); err != nil {
		return nil, err
	}
	return project, nil
}

// DeleteProject deletes a project
func (c *Client) DeleteProject(projectIDOrKey interface{}) (*Project, error) {
	return c.DeleteProjectContext(context.Background(), projectIDOrKey)
}

// DeleteProjectContext deletes a project with Context
func (c *Client) DeleteProjectContext(ctx context.Context, projectIDOrKey interface{}) (*Project, error) {
	u := fmt.Sprintf("/api/v2/projects/%v", projectIDOrKey)

	req, err := c.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	project := new(Project)
	if err := c.Do(ctx, req, &project); err != nil {
		return nil, err
	}
	return project, nil
}

// GetProjectsOptions contains all the parameters necessary (including the optional ones) for a GetProjects() request.
type GetProjectsOptions struct {
	Archived *bool `url:"archived"`
	All      *bool `url:"all"`
}

// CreateProjectInput contains all the parameters necessary (including the optional ones) for a CreateProject() request.
type CreateProjectInput struct {
	Name                              *string `json:"name"`
	Key                               *string `json:"key"`
	ChartEnabled                      *bool   `json:"chartEnabled"`
	ProjectLeaderCanEditProjectLeader *bool   `json:"projectLeaderCanEditProjectLeader,omitempty"`
	SubtaskingEnabled                 *bool   `json:"subtaskingEnabled"`
	TextFormattingRule                *string `json:"textFormattingRule"`
}

// UpdateProjectInput contains all the parameters necessary (including the optional ones) for a UpdateProject() request.
type UpdateProjectInput struct {
	Name                              *string `json:"name,omitempty"`
	Key                               *string `json:"key,omitempty"`
	ChartEnabled                      *bool   `json:"chartEnabled,omitempty"`
	SubtaskingEnabled                 *bool   `json:"subtaskingEnabled,omitempty"`
	ProjectLeaderCanEditProjectLeader *bool   `json:"projectLeaderCanEditProjectLeader,omitempty"`
	TextFormattingRule                *string `json:"textFormattingRule,omitempty"`
	Archived                          *bool   `json:"archived,omitempty"`
}
