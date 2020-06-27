package backlog

import (
	"context"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
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
func (api *Client) GetProjects(input *GetProjectsInput) ([]*Project, error) {
	return api.GetProjectsContext(context.Background(), input)
}

// GetProjectsContext returns the list of projects
func (api *Client) GetProjectsContext(ctx context.Context, input *GetProjectsInput) ([]*Project, error) {
	values := url.Values{}

	if input.All != nil {
		values.Add("all", strconv.FormatBool(*input.All))
	}

	if input.Archived != nil {
		values.Add("archived", strconv.FormatBool(*input.Archived))
	}

	projects := []*Project{}
	if err := api.getMethod(ctx, "/api/v2/projects", values, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}

// GetProject returns a project
func (api *Client) GetProject(projectIDOrKey interface{}) (*Project, error) {
	return api.GetProjectContext(context.Background(), projectIDOrKey)
}

// GetProjectContext returns a project with context
func (api *Client) GetProjectContext(ctx context.Context, projectIDOrKey interface{}) (*Project, error) {
	projIDOrKey, err := projIDOrKey(projectIDOrKey)
	if err != nil {
		return nil, err
	}

	project := new(Project)
	if err := api.getMethod(ctx, "/api/v2/projects/"+projIDOrKey, url.Values{}, &project); err != nil {
		return nil, err
	}
	return project, nil
}

// GetProjectStatuses returns the statuses of a project
func (api *Client) GetProjectStatuses(projectIDOrKey interface{}) ([]*Status, error) {
	return api.GetProjectStatusesContext(context.Background(), projectIDOrKey)
}

// GetProjectStatusesContext returns the statuses of a project with context
func (api *Client) GetProjectStatusesContext(ctx context.Context, projectIDOrKey interface{}) ([]*Status, error) {
	projIDOrKey, err := projIDOrKey(projectIDOrKey)
	if err != nil {
		return nil, err
	}

	statuses := []*Status{}
	if err := api.getMethod(ctx, "/api/v2/projects/"+projIDOrKey+"/statuses", url.Values{}, &statuses); err != nil {
		return nil, err
	}
	return statuses, nil
}

// CreateProject creates a project
func (api *Client) CreateProject(input *CreateProjectInput) (*Project, error) {
	return api.CreateProjectContext(context.Background(), input)
}

// CreateProjectContext creates a project with Context
func (api *Client) CreateProjectContext(ctx context.Context, input *CreateProjectInput) (*Project, error) {
	if *input.TextFormattingRule != "markdown" && *input.TextFormattingRule != "backlog" {
		return nil, errors.New("textFormattingRule is invalid: textFormattingRule must be backlog or markdown")
	}

	values := url.Values{}

	if input.Name != nil {
		values.Add("name", *input.Name)
	}

	if input.Key != nil {
		values.Add("key", *input.Key)
	}

	if input.ChartEnabled != nil {
		values.Add("chartEnabled", strconv.FormatBool(*input.ChartEnabled))
	}

	if input.ProjectLeaderCanEditProjectLeader != nil {
		values.Add("projectLeaderCanEditProjectLeader", strconv.FormatBool(*input.ProjectLeaderCanEditProjectLeader))
	}

	if input.SubtaskingEnabled != nil {
		values.Add("subtaskingEnabled", strconv.FormatBool(*input.SubtaskingEnabled))
	}

	if input.TextFormattingRule != nil {
		values.Add("textFormattingRule", *input.TextFormattingRule)
	}

	project := new(Project)
	if err := api.postMethod(ctx, "/api/v2/projects", values, &project); err != nil {
		return nil, err
	}
	return project, nil
}

// UpdateProject updates a project
func (api *Client) UpdateProject(input *UpdateProjectInput) (*Project, error) {
	return api.UpdateProjectContext(context.Background(), input)
}

// UpdateProjectContext updates a project with Context
func (api *Client) UpdateProjectContext(ctx context.Context, input *UpdateProjectInput) (*Project, error) {
	if input.TextFormattingRule != nil && *input.TextFormattingRule != "markdown" && *input.TextFormattingRule != "backlog" {
		return nil, errors.New("textFormattingRule is invalid: textFormattingRule must be backlog or markdown")
	}

	if input.ID == nil {
		return nil, errors.New("id is empty")
	}

	values := url.Values{}

	if input.Name != nil {
		values.Add("name", *input.Name)
	}

	if input.ProjectKey != nil {
		values.Add("key", *input.ProjectKey)
	}

	if input.ChartEnabled != nil {
		values.Add("chartEnabled", strconv.FormatBool(*input.ChartEnabled))
	}

	if input.ProjectLeaderCanEditProjectLeader != nil {
		values.Add("projectLeaderCanEditProjectLeader", strconv.FormatBool(*input.ProjectLeaderCanEditProjectLeader))
	}

	if input.SubtaskingEnabled != nil {
		values.Add("subtaskingEnabled", strconv.FormatBool(*input.SubtaskingEnabled))
	}

	if input.TextFormattingRule != nil {
		values.Add("textFormattingRule", *input.TextFormattingRule)
	}

	if input.Archived != nil {
		values.Add("archived", strconv.FormatBool(*input.Archived))
	}

	project := new(Project)
	if err := api.patchMethod(ctx, "/api/v2/projects/"+strconv.Itoa(*input.ID), values, &project); err != nil {
		return nil, err
	}
	return project, nil
}

// DeleteProject deletes a project
func (api *Client) DeleteProject(projectIDOrKey interface{}) (*Project, error) {
	return api.DeleteProjectContext(context.Background(), projectIDOrKey)
}

// DeleteProjectContext deletes a project with Context
func (api *Client) DeleteProjectContext(ctx context.Context, projectIDOrKey interface{}) (*Project, error) {
	projIDOrKey, err := projIDOrKey(projectIDOrKey)
	if err != nil {
		return nil, err
	}

	project := new(Project)
	if err := api.deleteMethod(ctx, "/api/v2/projects/"+projIDOrKey, url.Values{}, &project); err != nil {
		return nil, err
	}
	return project, nil
}

// GetProjectsInput contains all the parameters necessary (including the optional ones) for a GetProjects() request.
type GetProjectsInput struct {
	Archived *bool `required:"false"`
	All      *bool `required:"false"`
}

// CreateProjectInput contains all the parameters necessary (including the optional ones) for a CreateProject() request.
type CreateProjectInput struct {
	Name                              *string `required:"true"`
	Key                               *string `required:"true"`
	ChartEnabled                      *bool   `required:"true"`
	ProjectLeaderCanEditProjectLeader *bool   `required:"false"`
	SubtaskingEnabled                 *bool   `required:"true"`
	TextFormattingRule                *string `required:"true"`
}

// UpdateProjectInput contains all the parameters necessary (including the optional ones) for a UpdateProject() request.
type UpdateProjectInput struct {
	ID                                *int
	ProjectKey                        *string
	Name                              *string
	ChartEnabled                      *bool
	SubtaskingEnabled                 *bool
	ProjectLeaderCanEditProjectLeader *bool
	TextFormattingRule                *string
	Archived                          *bool
}
