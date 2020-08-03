package backlog

import (
	"context"
	"fmt"
	"io"
)

// Team : team
type Team struct {
	ID           *int       `json:"id,omitempty"`
	Name         *string    `json:"name,omitempty"`
	Members      []*User    `json:"members,omitempty"`
	DisplayOrder *int       `json:"displayOrder,omitempty"`
	CreatedUser  *User      `json:"createdUser,omitempty"`
	Created      *Timestamp `json:"created,omitempty"`
	UpdatedUser  *User      `json:"updatedUser,omitempty"`
	Updated      *Timestamp `json:"updated,omitempty"`
}

// GetTeams returns the list of teams
func (c *Client) GetTeams(opts *GetTeamsOptions) ([]*Team, error) {
	return c.GetTeamsContext(context.Background(), opts)
}

// GetTeamsContext returns the list of teams with context
func (c *Client) GetTeamsContext(ctx context.Context, opts *GetTeamsOptions) ([]*Team, error) {
	u, err := c.AddOptions("/api/v2/teams", opts)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var teams []*Team
	if err := c.Do(ctx, req, &teams); err != nil {
		return nil, err
	}
	return teams, nil
}

// CreateTeam creates a team
// a space backlog.com cannot use this API
func (c *Client) CreateTeam(input *CreateTeamInput) (*Team, error) {
	return c.CreateTeamContext(context.Background(), input)
}

// CreateTeamContext creates a team with Context
// a space backlog.com cannot use this API
func (c *Client) CreateTeamContext(ctx context.Context, input *CreateTeamInput) (*Team, error) {
	u := "/api/v2/teams"

	req, err := c.NewRequest("POST", u, input)
	if err != nil {
		return nil, err
	}

	team := new(Team)
	if err := c.Do(ctx, req, &team); err != nil {
		return nil, err
	}
	return team, nil
}

// GetTeam returns a teams
func (c *Client) GetTeam(teamID int) (*Team, error) {
	return c.GetTeamContext(context.Background(), teamID)
}

// GetTeamContext returns a team with context
func (c *Client) GetTeamContext(ctx context.Context, teamID int) (*Team, error) {
	u := fmt.Sprintf("/api/v2/teams/%v", teamID)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var team *Team
	if err := c.Do(ctx, req, &team); err != nil {
		return nil, err
	}
	return team, nil
}

// UpdateTeam updates a team
func (c *Client) UpdateTeam(teamID int, input *UpdateTeamInput) (*Team, error) {
	return c.UpdateTeamContext(context.Background(), teamID, input)
}

// UpdateTeamContext updates a team with Context
func (c *Client) UpdateTeamContext(ctx context.Context, teamID int, input *UpdateTeamInput) (*Team, error) {
	u := fmt.Sprintf("/api/v2/teams/%v", teamID)

	req, err := c.NewRequest("PATCH", u, input)
	if err != nil {
		return nil, err
	}

	team := new(Team)
	if err := c.Do(ctx, req, &team); err != nil {
		return nil, err
	}
	return team, nil
}

// DeleteTeam deletes a team
func (c *Client) DeleteTeam(teamID int) (*Team, error) {
	return c.DeleteTeamContext(context.Background(), teamID)
}

// DeleteTeamContext deletes a team with Context
func (c *Client) DeleteTeamContext(ctx context.Context, teamID int) (*Team, error) {
	u := fmt.Sprintf("/api/v2/teams/%v", teamID)

	req, err := c.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	team := new(Team)
	if err := c.Do(ctx, req, &team); err != nil {
		return nil, err
	}
	return team, nil
}

// GetTeamIcon downloads team icon
func (c *Client) GetTeamIcon(teamID int, writer io.Writer) error {
	return c.GetTeamIconContext(context.Background(), teamID, writer)
}

// GetTeamIconContext downloads team icon with context
func (c *Client) GetTeamIconContext(ctx context.Context, teamID int, writer io.Writer) error {
	u := fmt.Sprintf("/api/v2/teams/%v/icon", teamID)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, writer); err != nil {
		return err
	}
	return nil
}

// GetProjectTeams returns the list of teams in a project
func (c *Client) GetProjectTeams(projectIDOrKey interface{}) ([]*Team, error) {
	return c.GetProjectTeamsContext(context.Background(), projectIDOrKey)
}

// GetProjectTeamsContext returns the list of teams in a project with context
func (c *Client) GetProjectTeamsContext(ctx context.Context, projectIDOrKey interface{}) ([]*Team, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/teams", projectIDOrKey)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var teams []*Team
	if err := c.Do(ctx, req, &teams); err != nil {
		return nil, err
	}
	return teams, nil
}

// AddProjectTeam adds a team to a project
func (c *Client) AddProjectTeam(projectIDOrKey interface{}, input *AddProjectTeamInput) (*Team, error) {
	return c.AddProjectTeamContext(context.Background(), projectIDOrKey, input)
}

// AddProjectTeamContext adds a team to a project with context
func (c *Client) AddProjectTeamContext(ctx context.Context, projectIDOrKey interface{}, input *AddProjectTeamInput) (*Team, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/teams", projectIDOrKey)

	req, err := c.NewRequest("POST", u, input)
	if err != nil {
		return nil, err
	}

	team := new(Team)
	if err := c.Do(ctx, req, &team); err != nil {
		return nil, err
	}
	return team, nil
}

// DeleteProjectTeam deletes a team to a project
func (c *Client) DeleteProjectTeam(projectIDOrKey interface{}, input *DeleteProjectTeamInput) (*Team, error) {
	return c.DeleteProjectTeamContext(context.Background(), projectIDOrKey, input)
}

// DeleteProjectTeamContext deletes a team to a project with context
func (c *Client) DeleteProjectTeamContext(ctx context.Context, projectIDOrKey interface{}, input *DeleteProjectTeamInput) (*Team, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/teams", projectIDOrKey)

	req, err := c.NewRequest("DELETE", u, input)
	if err != nil {
		return nil, err
	}

	team := new(Team)
	if err := c.Do(ctx, req, &team); err != nil {
		return nil, err
	}
	return team, nil
}

// GetTeamsOptions specifies parameters to the GetTeams method.
type GetTeamsOptions struct {
	Order  Order `url:"order,omitempty"`
	Offset *int  `url:"offset,omitempty"`
	Count  *int  `url:"count,omitempty"`
}

// CreateTeamInput specifies parameters to the CreateTeam method.
type CreateTeamInput struct {
	Name    *string `json:"name"`
	Members []int   `json:"members,omitempty"`
}

// UpdateTeamInput specifies parameters to the UpdateTeam method.
type UpdateTeamInput struct {
	Name    *string `json:"name"`
	Members []int   `json:"members,omitempty"`
}

// AddProjectTeamInput specifies parameters to the AddProjectTeam method.
type AddProjectTeamInput struct {
	TeamID *int `json:"teamId"`
}

// DeleteProjectTeamInput specifies parameters to the DeleteProjectTeam method.
type DeleteProjectTeamInput struct {
	TeamID *int `json:"teamId"`
}
