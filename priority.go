package backlog

import "context"

// Priority : priority
type Priority struct {
	ID   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// GetPriorities returns the list of priorities
func (c *Client) GetPriorities() ([]*Priority, error) {
	return c.GetPrioritiesContext(context.Background())
}

// GetPrioritiesContext returns the list of priorities with context
func (c *Client) GetPrioritiesContext(ctx context.Context) ([]*Priority, error) {
	u := "/api/v2/priorities"

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var priorities []*Priority
	if err := c.Do(ctx, req, &priorities); err != nil {
		return nil, err
	}
	return priorities, nil
}
