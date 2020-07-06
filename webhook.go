package backlog

import (
	"context"
	"fmt"
)

// Webhook : -
type Webhook struct {
	ID              *int       `json:"id,omitempty"`
	Name            *string    `json:"name,omitempty"`
	Description     *string    `json:"description,omitempty"`
	HookURL         *string    `json:"hookUrl,omitempty"`
	AllEvent        *bool      `json:"allEvent,omitempty"`
	ActivityTypeIds []int      `json:"activityTypeIds,omitempty"`
	CreatedUser     *User      `json:"createdUser,omitempty"`
	Created         *Timestamp `json:"created,omitempty"`
	UpdatedUser     *User      `json:"updatedUser,omitempty"`
	Updated         *Timestamp `json:"updated,omitempty"`
}

// GetWebhook returns the list of webhooks
func (c *Client) GetWebhook(projectIDOrKey interface{}, webhookID int) (*Webhook, error) {
	return c.GetWebhookContext(context.Background(), projectIDOrKey, webhookID)
}

// GetWebhookContext returns the list of webhooks with context
func (c *Client) GetWebhookContext(ctx context.Context, projectIDOrKey interface{}, webhookID int) (*Webhook, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/webhooks/%v", projectIDOrKey, webhookID)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	webhook := new(Webhook)
	if err := c.Do(ctx, req, &webhook); err != nil {
		return nil, err
	}
	return webhook, nil
}

// GetWebhooks returns the list of webhooks
func (c *Client) GetWebhooks(projectIDOrKey interface{}) ([]*Webhook, error) {
	return c.GetWebhooksContext(context.Background(), projectIDOrKey)
}

// GetWebhooksContext returns the list of webhooks with context
func (c *Client) GetWebhooksContext(ctx context.Context, projectIDOrKey interface{}) ([]*Webhook, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/webhooks", projectIDOrKey)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	webhooks := []*Webhook{}
	if err := c.Do(ctx, req, &webhooks); err != nil {
		return nil, err
	}
	return webhooks, nil
}

// CreateWebhook adds a webhook
func (c *Client) CreateWebhook(projectIDOrKey interface{}, webhook *CreateWebhookInput) (*Webhook, error) {
	return c.CreateWebhookContext(context.Background(), projectIDOrKey, webhook)
}

// CreateWebhookContext adds a webhook with context
func (c *Client) CreateWebhookContext(ctx context.Context, projectIDOrKey interface{}, input *CreateWebhookInput) (*Webhook, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/webhooks", projectIDOrKey)

	req, err := c.NewRequest("POST", u, input)
	if err != nil {
		return nil, err
	}

	webhook := new(Webhook)
	if err := c.Do(ctx, req, &webhook); err != nil {
		return nil, err
	}
	return webhook, nil
}

// UpdateWebhook updates a webhook
func (c *Client) UpdateWebhook(projectIDOrKey interface{}, webhookID int, input *UpdateWebhookInput) (*Webhook, error) {
	return c.UpdateWebhookContext(context.Background(), projectIDOrKey, webhookID, input)
}

// UpdateWebhookContext updates a webhook with context
func (c *Client) UpdateWebhookContext(ctx context.Context, projectIDOrKey interface{}, webhookID int, input *UpdateWebhookInput) (*Webhook, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/webhooks/%v", projectIDOrKey, webhookID)

	req, err := c.NewRequest("PATCH", u, input)
	if err != nil {
		return nil, err
	}

	webhook := new(Webhook)
	if err := c.Do(ctx, req, &webhook); err != nil {
		return nil, err
	}
	return webhook, nil
}

// DeleteWebhook deletes a webhook
func (c *Client) DeleteWebhook(projectIDOrKey interface{}, webhookID int) (*Webhook, error) {
	return c.DeleteWebhookContext(context.Background(), projectIDOrKey, webhookID)
}

// DeleteWebhookContext updates a webhook with context
func (c *Client) DeleteWebhookContext(ctx context.Context, projectIDOrKey interface{}, webhookID int) (*Webhook, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/webhooks/%v", projectIDOrKey, webhookID)

	req, err := c.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	webhook := new(Webhook)
	if err := c.Do(ctx, req, &webhook); err != nil {
		return nil, err
	}
	return webhook, nil
}

// CreateWebhookInput contains all the parameters necessary (including the optional ones) for a CreateWebhook() request.
type CreateWebhookInput struct {
	Name            *string `json:"name"`
	Description     *string `json:"description,omitempty"`
	HookURL         *string `json:"hookUrl"`
	AllEvent        *bool   `json:"allEvent,omitempty"`
	ActivityTypeIDs []int   `json:"activityTypeIds,omitempty"`
}

// UpdateWebhookInput contains all the parameters necessary (including the optional ones) for a UpdateWebhook() request.
type UpdateWebhookInput struct {
	Name            *string `json:"name,omitempty"`
	Description     *string `json:"description,omitempty"`
	HookURL         *string `json:"hookUrl,omitempty"`
	AllEvent        *bool   `json:"allEvent,omitempty"`
	ActivityTypeIDs []int   `json:"activityTypeIds,omitempty"`
}
