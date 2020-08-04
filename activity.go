package backlog

import (
	"context"
	"fmt"
)

// Activity : activity
type Activity struct {
	ID            *int            `json:"id,omitempty"` // User.ID
	Project       *Project        `json:"project,omitempty"`
	Type          *int            `json:"type,omitempty"`
	Content       *Content        `json:"content,omitempty"`
	Notifications []*Notification `json:"notifications,omitempty"`
	CreatedUser   *User           `json:"createdUser,omitempty"`
	Created       *Timestamp      `json:"created,omitempty"`
}

// GetUserActivities returns the list of a user's activities
func (c *Client) GetUserActivities(id int, opts *GetUserActivitiesOptions) ([]*Activity, error) {
	return c.GetUserActivitiesContext(context.Background(), id, opts)
}

// GetUserActivitiesContext returns the list of a user's activities with context
func (c *Client) GetUserActivitiesContext(ctx context.Context, id int, opts *GetUserActivitiesOptions) ([]*Activity, error) {
	u := fmt.Sprintf("/api/v2/users/%v/activities", id)

	u, err := c.AddOptions(u, opts)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var activities []*Activity
	if err := c.Do(ctx, req, &activities); err != nil {
		return nil, err
	}
	return activities, nil
}

// GetProjectActivities returns the list of a project's activities
func (c *Client) GetProjectActivities(projectIDOrKey interface{}, opts *GetProjectActivitiesOptions) ([]*Activity, error) {
	return c.GetProjectActivitiesContext(context.Background(), projectIDOrKey, opts)
}

// GetProjectActivitiesContext returns the list of a project's activities with context
func (c *Client) GetProjectActivitiesContext(ctx context.Context, projectIDOrKey interface{}, opts *GetProjectActivitiesOptions) ([]*Activity, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/activities", projectIDOrKey)

	u, err := c.AddOptions(u, opts)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var activities []*Activity
	if err := c.Do(ctx, req, &activities); err != nil {
		return nil, err
	}
	return activities, nil
}

// GetUserActivitiesOptions specifies parameters to the GetUserActivities method.
type GetUserActivitiesOptions struct {
	ActivityTypeIDs []int `url:"activityTypeId[],omitempty"`
	MinID           *int  `url:"minId,omitempty"`
	MaxID           *int  `url:"maxId,omitempty"`
	Count           *int  `url:"count,omitempty"`
	Order           Order `url:"order,omitempty"`
}

// GetProjectActivitiesOptions specifies parameters to the GetProjectActivities method.
type GetProjectActivitiesOptions struct {
	ActivityTypeIDs []int `url:"activityTypeId[],omitempty"`
	MinID           *int  `url:"minId,omitempty"`
	MaxID           *int  `url:"maxId,omitempty"`
	Count           *int  `url:"count,omitempty"`
	Order           Order `url:"order,omitempty"`
}
