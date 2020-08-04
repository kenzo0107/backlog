package backlog

import (
	"context"
	"fmt"
	"io"
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

// ProjectDiskUsage : disk usage of project
type ProjectDiskUsage struct {
	ProjectID  *int `json:"projectId,omitempty"`
	Issue      *int `json:"issue,omitempty"`
	Wiki       *int `json:"wiki,omitempty"`
	File       *int `json:"file,omitempty"`
	Subversion *int `json:"subversion,omitempty"`
	Git        *int `json:"git,omitempty"`
	GitLFS     *int `json:"gitLFS,omitempty"`
}

// RecentlyViewedProject : recently viewed project
type RecentlyViewedProject struct {
	Project *Project   `json:"project"`
	Updated *Timestamp `json:"updated"`
}

// GetMyRecentlyViewedProjects returns the list of projects I recently viewed
func (c *Client) GetMyRecentlyViewedProjects(opts *GetMyRecentlyViewedProjectsOptions) ([]*RecentlyViewedProject, error) {
	return c.GetMyRecentlyViewedProjectsContext(context.Background(), opts)
}

// GetMyRecentlyViewedProjectsContext returns the list of projects I recently viewed with context
func (c *Client) GetMyRecentlyViewedProjectsContext(ctx context.Context, opts *GetMyRecentlyViewedProjectsOptions) ([]*RecentlyViewedProject, error) {
	u, err := c.AddOptions("/api/v2/users/myself/recentlyViewedProjects", opts)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var recentlyViewedProjects []*RecentlyViewedProject
	if err := c.Do(ctx, req, &recentlyViewedProjects); err != nil {
		return nil, err
	}
	return recentlyViewedProjects, nil
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

// GetStatuses returns the statuses of a project
func (c *Client) GetStatuses(projectIDOrKey interface{}) ([]*Status, error) {
	return c.GetStatusesContext(context.Background(), projectIDOrKey)
}

// GetStatusesContext returns the statuses of a project with context
func (c *Client) GetStatusesContext(ctx context.Context, projectIDOrKey interface{}) ([]*Status, error) {
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

// GetProjectIcon downloads project icon
func (c *Client) GetProjectIcon(projectIDOrKey interface{}, writer io.Writer) error {
	return c.GetProjectIconContext(context.Background(), projectIDOrKey, writer)
}

// GetProjectIconContext downloads project icon with context
func (c *Client) GetProjectIconContext(ctx context.Context, projectIDOrKey interface{}, writer io.Writer) error {
	u := fmt.Sprintf("/api/v2/projects/%v/image", projectIDOrKey)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, writer); err != nil {
		return err
	}
	return nil
}

// AddProjectUser adds a user to a project
func (c *Client) AddProjectUser(projectIDOrKey interface{}, input *AddProjectUserInput) (*User, error) {
	return c.AddProjectUserContext(context.Background(), projectIDOrKey, input)
}

// AddProjectUserContext adds a user to a project with context
func (c *Client) AddProjectUserContext(ctx context.Context, projectIDOrKey interface{}, input *AddProjectUserInput) (*User, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/users", projectIDOrKey)

	req, err := c.NewRequest("POST", u, input)
	if err != nil {
		return nil, err
	}

	var user *User
	if err := c.Do(ctx, req, &user); err != nil {
		return nil, err
	}
	return user, nil
}

// GetProjectUsers returns the list of users in a project
func (c *Client) GetProjectUsers(projectIDOrKey interface{}, opts *GetProjectUsersOptions) ([]*User, error) {
	return c.GetProjectUsersContext(context.Background(), projectIDOrKey, opts)
}

// GetProjectUsersContext returns the list of users in a project with context
func (c *Client) GetProjectUsersContext(ctx context.Context, projectIDOrKey interface{}, opts *GetProjectUsersOptions) ([]*User, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/users", projectIDOrKey)

	u, err := c.AddOptions(u, opts)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var users []*User
	if err := c.Do(ctx, req, &users); err != nil {
		return nil, err
	}
	return users, nil
}

// DeleteProjectUser deletes a user in a project
func (c *Client) DeleteProjectUser(projectIDOrKey interface{}, input *DeleteProjectUserInput) (*User, error) {
	return c.DeleteProjectUserContext(context.Background(), projectIDOrKey, input)
}

// DeleteProjectUserContext deletes a user in a project with Context
func (c *Client) DeleteProjectUserContext(ctx context.Context, projectIDOrKey interface{}, input *DeleteProjectUserInput) (*User, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/users", projectIDOrKey)

	req, err := c.NewRequest("DELETE", u, input)
	if err != nil {
		return nil, err
	}

	user := new(User)
	if err := c.Do(ctx, req, &user); err != nil {
		return nil, err
	}
	return user, nil
}

// AddProjectAdministrator adds an administrator in a project
func (c *Client) AddProjectAdministrator(projectIDOrKey interface{}, input *AddProjectAdministratorInput) (*User, error) {
	return c.AddProjectAdministratorContext(context.Background(), projectIDOrKey, input)
}

// AddProjectAdministratorContext adds an administrator in a project with context
func (c *Client) AddProjectAdministratorContext(ctx context.Context, projectIDOrKey interface{}, input *AddProjectAdministratorInput) (*User, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/administrators", projectIDOrKey)

	req, err := c.NewRequest("POST", u, input)
	if err != nil {
		return nil, err
	}

	user := new(User)
	if err := c.Do(ctx, req, &user); err != nil {
		return nil, err
	}
	return user, nil
}

// GetProjectAdministrators returns the list of administrators in a project
func (c *Client) GetProjectAdministrators(projectIDOrKey interface{}) ([]*User, error) {
	return c.GetProjectAdministratorsContext(context.Background(), projectIDOrKey)
}

// GetProjectAdministratorsContext returns the list of administrators in a project with context
func (c *Client) GetProjectAdministratorsContext(ctx context.Context, projectIDOrKey interface{}) ([]*User, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/administrators", projectIDOrKey)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var users []*User
	if err := c.Do(ctx, req, &users); err != nil {
		return nil, err
	}
	return users, nil
}

// DeleteProjectAdministrator deletes a administrator in a project
func (c *Client) DeleteProjectAdministrator(projectIDOrKey interface{}, input *DeleteProjectAdministratorInput) (*User, error) {
	return c.DeleteProjectAdministratorContext(context.Background(), projectIDOrKey, input)
}

// DeleteProjectAdministratorContext deletes a administrator in a project with Context
func (c *Client) DeleteProjectAdministratorContext(ctx context.Context, projectIDOrKey interface{}, input *DeleteProjectAdministratorInput) (*User, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/administrators", projectIDOrKey)

	req, err := c.NewRequest("DELETE", u, input)
	if err != nil {
		return nil, err
	}

	user := new(User)
	if err := c.Do(ctx, req, &user); err != nil {
		return nil, err
	}
	return user, nil
}

// CreateStatus creates a status
func (c *Client) CreateStatus(projectIDOrKey interface{}, input *CreateStatusInput) (*Status, error) {
	return c.CreateStatusContext(context.Background(), projectIDOrKey, input)
}

// CreateStatusContext creates a status
func (c *Client) CreateStatusContext(ctx context.Context, projectIDOrKey interface{}, input *CreateStatusInput) (*Status, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/statuses", projectIDOrKey)

	req, err := c.NewRequest("POST", u, input)
	if err != nil {
		return nil, err
	}

	status := new(Status)
	if err := c.Do(ctx, req, &status); err != nil {
		return nil, err
	}
	return status, nil
}

// UpdateStatus updates a status
func (c *Client) UpdateStatus(projectIDOrKey interface{}, statusID int, input *UpdateStatusInput) (*Status, error) {
	return c.UpdateStatusContext(context.Background(), projectIDOrKey, statusID, input)
}

// UpdateStatusContext updates a status
func (c *Client) UpdateStatusContext(ctx context.Context, projectIDOrKey interface{}, statusID int, input *UpdateStatusInput) (*Status, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/statuses/%v", projectIDOrKey, statusID)

	req, err := c.NewRequest("PATCH", u, input)
	if err != nil {
		return nil, err
	}

	status := new(Status)
	if err := c.Do(ctx, req, &status); err != nil {
		return nil, err
	}
	return status, nil
}

// DeleteStatus deletes a status
func (c *Client) DeleteStatus(projectIDOrKey interface{}, statusID int, input *DeleteStatusInput) (*Status, error) {
	return c.DeleteStatusContext(context.Background(), projectIDOrKey, statusID, input)
}

// DeleteStatusContext deletes a status
func (c *Client) DeleteStatusContext(ctx context.Context, projectIDOrKey interface{}, statusID int, input *DeleteStatusInput) (*Status, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/statuses/%v", projectIDOrKey, statusID)

	req, err := c.NewRequest("DELETE", u, input)
	if err != nil {
		return nil, err
	}

	status := new(Status)
	if err := c.Do(ctx, req, &status); err != nil {
		return nil, err
	}
	return status, nil
}

// SortStatuses sorts the list of statuses
func (c *Client) SortStatuses(projectIDOrKey interface{}, input *SortStatusesInput) ([]*Status, error) {
	return c.SortStatusesContext(context.Background(), projectIDOrKey, input)
}

// SortStatusesContext sorts the list of statuses with context
func (c *Client) SortStatusesContext(ctx context.Context, projectIDOrKey interface{}, input *SortStatusesInput) ([]*Status, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/statuses/updateDisplayOrder", projectIDOrKey)

	req, err := c.NewRequest("PATCH", u, input)
	if err != nil {
		return nil, err
	}

	statuses := []*Status{}
	if err := c.Do(ctx, req, &statuses); err != nil {
		return nil, err
	}
	return statuses, nil
}

// GetProjectDiskUsage returns disk usage of a project
func (c *Client) GetProjectDiskUsage(projectIDOrKey interface{}) (*ProjectDiskUsage, error) {
	return c.GetProjectDiskUsageContext(context.Background(), projectIDOrKey)
}

// GetProjectDiskUsageContext returns the list of administrators in a project with context
func (c *Client) GetProjectDiskUsageContext(ctx context.Context, projectIDOrKey interface{}) (*ProjectDiskUsage, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/diskUsage", projectIDOrKey)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	projectDiskUsage := new(ProjectDiskUsage)
	if err := c.Do(ctx, req, &projectDiskUsage); err != nil {
		return nil, err
	}
	return projectDiskUsage, nil
}

// GetMyRecentlyViewedProjectsOptions specifies parameters to the GetMyRecentlyViewedProject method.
type GetMyRecentlyViewedProjectsOptions struct {
	Order  Order `url:"order,omitempty"`
	Offset *int  `url:"offset,omitempty"`
	Count  *int  `url:"count,omitempty"`
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

// AddProjectUserInput specifies parameters to the AddProjectUser method.
type AddProjectUserInput struct {
	UserID *int `json:"userId"`
}

// GetProjectUsersOptions specifies parameters to the GetProjectUsers method.
type GetProjectUsersOptions struct {
	ExcludeGroupMembers *bool `url:"excludeGroupMembers,omitempty"`
}

// DeleteProjectUserInput specifies parameters to the DeleteProjectUser method.
type DeleteProjectUserInput struct {
	UserID *int `json:"userId"`
}

// AddProjectAdministratorInput specifies parameters to the AddProjectAdministrator method.
type AddProjectAdministratorInput struct {
	UserID *int `json:"userId"`
}

// DeleteProjectAdministratorInput specifies parameters to the DeleteProjectAdministrator method.
type DeleteProjectAdministratorInput struct {
	UserID *int `json:"userId"`
}

// CreateStatusInput specifies parameters to the CreateStatus method.
type CreateStatusInput struct {
	Name  *string `json:"name"`
	Color *string `json:"color"`
}

// UpdateStatusInput specifies parameters to the UpdateStatus method.
type UpdateStatusInput struct {
	Name  *string `json:"name,omitempty"`
	Color *string `json:"color,omitempty"`
}

// DeleteStatusInput specifies parameters to the DeleteStatus method.
type DeleteStatusInput struct {
	SubstituteStatusID *int `json:"substituteStatusId"`
}

// SortStatusesInput specifies parameters to the SortStatuses method.
type SortStatusesInput struct {
	StatusIDs []int `json:"statusId,omitempty"`
}
