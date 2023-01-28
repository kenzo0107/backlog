package backlog

import (
	"context"
	"time"
)

// LimitStatus : limit status
type LimitStatus struct {
	Limit     *int `json:"limit,omitempty"`
	Remaining *int `json:"remaining,omitempty"`
	Reset     *int `json:"reset,omitempty"`
}

func (ls *LimitStatus) ResetAsTime() time.Time {
	if ls.Reset == nil {
		return time.Time{}
	}
	return time.Unix(int64(*ls.Reset), 0)
}

// RateLimit : rate limit
type RateLimit struct {
	Read   *LimitStatus `json:"read,omitempty"`
	Update *LimitStatus `json:"update,omitempty"`
	Search *LimitStatus `json:"search,omitempty"`
	Icon   *LimitStatus `json:"icon,omitempty"`
}

// GetRateLimit returns the rate limit
func (c *Client) GetRateLimit() (*RateLimit, error) {
	return c.GetRateLimitContext(context.Background())
}

// GetWikisContext returns the rate limit
func (c *Client) GetRateLimitContext(ctx context.Context) (*RateLimit, error) {
	u := "/api/v2/rateLimit"

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	wrapper := new(struct {
		RateLimit *RateLimit `json:"rateLimit,omitempty"`
	})
	if err := c.Do(ctx, req, &wrapper); err != nil {
		return nil, err
	}
	return wrapper.RateLimit, nil
}
