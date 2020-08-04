package backlog

import (
	"context"
	"fmt"
)

// Watching : -
type Watching struct {
	ID                  *int       `json:"id,omitempty"`
	ResourceAlreadyRead *bool      `json:"resourceAlreadyRead,omitempty"`
	Note                *string    `json:"note,omitempty"`
	Type                *string    `json:"type,omitempty"`
	Issue               *Issue     `json:"issue,omitempty"`
	LastContentUpdated  *Timestamp `json:"lastContentUpdated,omitempty"`
	Created             *Timestamp `json:"created,omitempty"`
	Updated             *Timestamp `json:"updated,omitempty"`
}

// GetUserWatchings returns the list of user's watchings
func (c *Client) GetUserWatchings(userID int) ([]*Watching, error) {
	return c.GetUserWatchingsContext(context.Background(), userID)
}

// GetUserWatchingsContext returns the list of user's watchings with context
func (c *Client) GetUserWatchingsContext(ctx context.Context, userID int) ([]*Watching, error) {
	u := fmt.Sprintf("/api/v2/users/%v/watchings", userID)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var watchings []*Watching
	if err := c.Do(ctx, req, &watchings); err != nil {
		return nil, err
	}
	return watchings, nil
}

// GetUserWatchingsCount returns the count of user's watchings
func (c *Client) GetUserWatchingsCount(userID int, opts *GetUserWatchingsCountOptions) (int, error) {
	return c.GetUserWatchingsCountContext(context.Background(), userID, opts)
}

// GetUserWatchingsCountContext returns the count of user's watchings with context
func (c *Client) GetUserWatchingsCountContext(ctx context.Context, userID int, opts *GetUserWatchingsCountOptions) (int, error) {
	u := fmt.Sprintf("/api/v2/users/%v/watchings/count", userID)

	u, err := c.AddOptions(u, opts)
	if err != nil {
		return 0, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return 0, err
	}

	r := new(p)
	if err := c.Do(ctx, req, &r); err != nil {
		return 0, err
	}
	return r.Count, nil
}

// GetWatching returns a watching
func (c *Client) GetWatching(watchingID int) (*Watching, error) {
	return c.GetWatchingContext(context.Background(), watchingID)
}

// GetWatchingContext returns a watching with context
func (c *Client) GetWatchingContext(ctx context.Context, watchingID int) (*Watching, error) {
	u := fmt.Sprintf("/api/v2/watchings/%v", watchingID)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	watching := new(Watching)
	if err := c.Do(ctx, req, &watching); err != nil {
		return nil, err
	}
	return watching, nil
}

// CreateWatching creates a watching
func (c *Client) CreateWatching(input *CreateWatchingInput) (*Watching, error) {
	return c.CreateWatchingContext(context.Background(), input)
}

// CreateWatchingContext creates a watching with Context
func (c *Client) CreateWatchingContext(ctx context.Context, input *CreateWatchingInput) (*Watching, error) {
	u := "/api/v2/watchings"

	req, err := c.NewRequest("POST", u, input)
	if err != nil {
		return nil, err
	}

	watching := new(Watching)
	if err := c.Do(ctx, req, &watching); err != nil {
		return nil, err
	}
	return watching, nil
}

// UpdateWatching updates a watching
func (c *Client) UpdateWatching(watchingID int, input *UpdateWatchingInput) (*Watching, error) {
	return c.UpdateWatchingContext(context.Background(), watchingID, input)
}

// UpdateWatchingContext updates a watching with Context
func (c *Client) UpdateWatchingContext(ctx context.Context, watchingID int, input *UpdateWatchingInput) (*Watching, error) {
	u := fmt.Sprintf("/api/v2/watchings/%v", watchingID)

	req, err := c.NewRequest("PATCH", u, input)
	if err != nil {
		return nil, err
	}

	watching := new(Watching)
	if err := c.Do(ctx, req, &watching); err != nil {
		return nil, err
	}
	return watching, nil
}

// DeleteWatching deletes a watching
func (c *Client) DeleteWatching(watchingID int) (*Watching, error) {
	return c.DeleteWatchingContext(context.Background(), watchingID)
}

// DeleteWatchingContext deletes a watching with Context
func (c *Client) DeleteWatchingContext(ctx context.Context, watchingID int) (*Watching, error) {
	u := fmt.Sprintf("/api/v2/watchings/%v", watchingID)

	req, err := c.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	watching := new(Watching)
	if err := c.Do(ctx, req, &watching); err != nil {
		return nil, err
	}
	return watching, nil
}

// MarkAsReadWatching marks a watching as read
func (c *Client) MarkAsReadWatching(watchingID int) error {
	return c.MarkAsReadWatchingContext(context.Background(), watchingID)
}

// MarkAsReadWatchingContext marks a watching as read with Context
func (c *Client) MarkAsReadWatchingContext(ctx context.Context, watchingID int) error {
	u := fmt.Sprintf("/api/v2/watchings/%v/markAsRead", watchingID)

	req, err := c.NewRequest("POST", u, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}
	return nil
}

// GetUserWatchingsOptions specifies parameters to the GetUserWatchings method.
type GetUserWatchingsOptions struct {
	Order               *Order  `url:"order,omitempty"`
	Sort                *string `url:"sort,omitempty"`
	Count               *int    `url:"count,omitempty"`
	Offset              *int    `url:"offset,omitempty"`
	ResourceAlreadyRead *bool   `url:"resourceAlreadyRead,omitempty"`
	IssueIDs            []int   `url:"issueId[],omitempty"`
}

// GetUserWatchingsCountOptions specifies parameters to the GetUserWatchingsCount method.
type GetUserWatchingsCountOptions struct {
	ResourceAlreadyRead *bool `url:"resourceAlreadyRead,omitempty"`
	AlreadyRead         *bool `url:"alreadyRead,omitempty"`
}

// CreateWatchingInput specifies parameters to the CreateWatching method.
type CreateWatchingInput struct {
	IssueIDOrKey *string `json:"issueIdOrKey"`
	Note         *string `json:"note,omitempty"`
}

// UpdateWatchingInput specifies parameters to the UpdateWatching method.
type UpdateWatchingInput struct {
	Note *string `json:"note,omitempty"`
}
