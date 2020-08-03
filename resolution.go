package backlog

import "context"

// Resolution : resolutions
type Resolution struct {
	ID   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// GetResolutions returns the list of resolutions
func (c *Client) GetResolutions() ([]*Resolution, error) {
	return c.GetResolutionsContext(context.Background())
}

// GetResolutionsContext returns the list of resolutions with context
func (c *Client) GetResolutionsContext(ctx context.Context) ([]*Resolution, error) {
	u := "/api/v2/resolutions"

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var resolutions []*Resolution
	if err := c.Do(ctx, req, &resolutions); err != nil {
		return nil, err
	}
	return resolutions, nil
}
