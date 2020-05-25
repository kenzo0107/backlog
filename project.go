package backlog

import (
	"context"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

// Project : project
type Project struct {
	ID                                int    `json:"id"`
	ProjectKey                        string `json:"projectKey"`
	Name                              string `json:"name"`
	ChartEnabled                      bool   `json:"chartEnabled"`
	SubtaskingEnabled                 bool   `json:"subtaskingEnabled"`
	ProjectLeaderCanEditProjectLeader bool   `json:"projectLeaderCanEditProjectLeader"`
	TextFormattingRule                string `json:"textFormattingRule"`
	Archived                          bool   `json:"archived"`
}

// ProjectStatus : the status of project
type ProjectStatus struct {
	ID           int    `json:"id"`
	ProjectID    int    `json:"projectId"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	DisplayOrder int    `json:"displayOrder"`
}

// GetProjects returns the list of projects
func (api *Client) GetProjects(archived bool, all bool) ([]Project, error) {
	return api.GetProjectsContext(context.Background(), archived, all)
}

// GetProjectsContext returns the list of projects
func (api *Client) GetProjectsContext(ctx context.Context, archived bool, all bool) ([]Project, error) {
	values := url.Values{}
	values.Add("archived", strconv.FormatBool(archived))
	values.Add("all", strconv.FormatBool(all))

	projects := []Project{}
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

	project := Project{}
	if err := api.getMethod(ctx, "/api/v2/projects/"+projIDOrKey, url.Values{}, &project); err != nil {
		return nil, err
	}
	return &project, nil
}

// GetProjectStatuses returns the statuses of a project
func (api *Client) GetProjectStatuses(projectIDOrKey interface{}) ([]ProjectStatus, error) {
	return api.GetProjectStatusesContext(context.Background(), projectIDOrKey)
}

// GetProjectStatusesContext returns the statuses of a project with context
func (api *Client) GetProjectStatusesContext(ctx context.Context, projectIDOrKey interface{}) ([]ProjectStatus, error) {
	projIDOrKey, err := projIDOrKey(projectIDOrKey)
	if err != nil {
		return nil, err
	}

	projectStatus := []ProjectStatus{}
	if err := api.getMethod(ctx, "/api/v2/projects/"+projIDOrKey+"/statuses", url.Values{}, &projectStatus); err != nil {
		return nil, err
	}
	return projectStatus, nil
}

// CreateProject creates a project
func (api *Client) CreateProject(input *CreateProjectInput) (*Project, error) {
	return api.CreateProjectContext(context.Background(), input)
}

// CreateProjectContext creates a project with Context
func (api *Client) CreateProjectContext(ctx context.Context, input *CreateProjectInput) (*Project, error) {
	if input.TextFormattingRule != "markdown" && input.TextFormattingRule != "backlog" {
		return nil, errors.New("textFormattingRule is invalid: textFormattingRule must be backlog or markdown")
	}

	values := url.Values{
		"name":                              {input.Name},
		"key":                               {input.Key},
		"chartEnabled":                      {strconv.FormatBool(input.ChartEnabled)},
		"projectLeaderCanEditProjectLeader": {strconv.FormatBool(input.ProjectLeaderCanEditProjectLeader)},
		"subtaskingEnabled":                 {strconv.FormatBool(input.SubtaskingEnabled)},
		"textFormattingRule":                {input.TextFormattingRule},
	}

	project := Project{}
	if err := api.postMethod(ctx, "/api/v2/projects", values, &project); err != nil {
		return nil, err
	}
	return &project, nil
}

// UpdateProject updates a project
func (api *Client) UpdateProject(input *UpdateProjectInput) (*Project, error) {
	return api.UpdateProjectContext(context.Background(), input)
}

// UpdateProjectContext updates a project with Context
func (api *Client) UpdateProjectContext(ctx context.Context, input *UpdateProjectInput) (*Project, error) {
	if input.TextFormattingRule != "markdown" && input.TextFormattingRule != "backlog" {
		return nil, errors.New("textFormattingRule is invalid: textFormattingRule must be backlog or markdown")
	}

	values := url.Values{
		"name":                              {input.Name},
		"key":                               {input.ProjectKey},
		"chartEnabled":                      {strconv.FormatBool(input.ChartEnabled)},
		"projectLeaderCanEditProjectLeader": {strconv.FormatBool(input.ProjectLeaderCanEditProjectLeader)},
		"subtaskingEnabled":                 {strconv.FormatBool(input.SubtaskingEnabled)},
		"textFormattingRule":                {input.TextFormattingRule},
		"archived":                          {strconv.FormatBool(input.Archived)},
	}

	project := Project{}
	if err := api.patchMethod(ctx, "/api/v2/projects/"+strconv.Itoa(input.ID), values, &project); err != nil {
		return nil, err
	}
	return &project, nil
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

	r := Project{}
	if err := api.deleteMethod(ctx, "/api/v2/projects/"+projIDOrKey, url.Values{}, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

// CreateProjectInput contains all the parameters necessary (including the optional ones) for a CreateProject() request.
type CreateProjectInput struct {
	Name                              string `required:"true"`
	Key                               string `required:"true"`
	ChartEnabled                      bool   `required:"true"`
	ProjectLeaderCanEditProjectLeader bool   `required:"false"`
	SubtaskingEnabled                 bool   `required:"true"`
	TextFormattingRule                string `required:"true"`
}

// UpdateProjectInput contains all the parameters necessary (including the optional ones) for a UpdateProject() request.
type UpdateProjectInput struct {
	ID                                int
	ProjectKey                        string
	Name                              string
	ChartEnabled                      bool
	SubtaskingEnabled                 bool
	ProjectLeaderCanEditProjectLeader bool
	TextFormattingRule                string
	Archived                          bool
}
