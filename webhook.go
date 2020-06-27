package backlog

import (
	"context"
	"net/url"
	"strconv"
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

// GetWebhooks returns the list of webhooks
func (api *Client) GetWebhooks(projectIDOrKey interface{}) ([]*Webhook, error) {
	return api.GetWebhooksContext(context.Background(), projectIDOrKey)
}

// GetWebhooksContext returns the list of webhooks with context
func (api *Client) GetWebhooksContext(ctx context.Context, projectIDOrKey interface{}) ([]*Webhook, error) {
	projIDOrKey, err := projIDOrKey(projectIDOrKey)
	if err != nil {
		return nil, err
	}

	values := url.Values{}
	values.Add("projectIdOrKey", projIDOrKey)

	r := []*Webhook{}
	if err = api.getMethod(ctx, "/api/v2/projects/"+projIDOrKey+"/webhooks", values, &r); err != nil {
		return nil, err
	}
	return r, nil
}

// AddWebhook adds a webhook
func (api *Client) AddWebhook(input *AddWebhookInput) (*Webhook, error) {
	return api.AddWebhookContext(context.Background(), input)
}

// AddWebhookContext adds a webhook with context
func (api *Client) AddWebhookContext(ctx context.Context, input *AddWebhookInput) (*Webhook, error) {
	projIDOrKey, err := projIDOrKey(input.ProjectIDOrKey)
	if err != nil {
		return nil, err
	}

	values := url.Values{}

	if input.Name != nil {
		values.Add("name", *input.Name)
	}

	if input.Description != nil {
		values.Add("description", *input.Description)
	}

	if input.HookURL != nil {
		values.Add("hookUrl", *input.HookURL)
	}

	if input.AllEvent != nil {
		values.Add("allEvent", strconv.FormatBool(*input.AllEvent))
	}

	for _, i := range input.ActivityTypeIDs {
		values.Add("activityTypeIds[]", strconv.Itoa(i))
	}

	webhook := new(Webhook)
	if err = api.postMethod(ctx, "/api/v2/projects/"+projIDOrKey+"/webhooks", values, &webhook); err != nil {
		return nil, err
	}
	return webhook, nil
}

// GetWebhook returns the list of webhooks
func (api *Client) GetWebhook(projectIDOrKey interface{}, webhookID int) (*Webhook, error) {
	return api.GetWebhookContext(context.Background(), projectIDOrKey, webhookID)
}

// GetWebhookContext returns the list of webhooks with context
func (api *Client) GetWebhookContext(ctx context.Context, projectIDOrKey interface{}, webhookID int) (*Webhook, error) {
	projIDOrKey, err := projIDOrKey(projectIDOrKey)
	if err != nil {
		return nil, err
	}

	webhook := new(Webhook)
	if err = api.getMethod(ctx, "/api/v2/projects/"+projIDOrKey+"/webhooks/"+strconv.Itoa(webhookID), url.Values{}, &webhook); err != nil {
		return nil, err
	}
	return webhook, nil
}

// UpdateWebhook updates a webhook
func (api *Client) UpdateWebhook(input *UpdateWebhookInput) (*Webhook, error) {
	return api.UpdateWebhookContext(context.Background(), input)
}

// UpdateWebhookContext updates a webhook with context
func (api *Client) UpdateWebhookContext(ctx context.Context, input *UpdateWebhookInput) (*Webhook, error) {
	projIDOrKey, err := projIDOrKey(input.ProjectIDOrKey)
	if err != nil {
		return nil, err
	}

	values := url.Values{}

	if input.Name != nil {
		values.Add("name", *input.Name)
	}

	if input.Description != nil {
		values.Add("description", *input.Description)
	}

	if input.HookURL != nil {
		values.Add("hookUrl", *input.HookURL)
	}

	if input.AllEvent != nil {
		values.Add("allEvent", strconv.FormatBool(*input.AllEvent))
	}

	for _, i := range input.ActivityTypeIDs {
		values.Add("activityTypeIds[]", strconv.Itoa(i))
	}

	webhook := new(Webhook)
	if err = api.patchMethod(ctx, "/api/v2/projects/"+projIDOrKey+"/webhooks/"+strconv.Itoa(*input.WebhookID), values, &webhook); err != nil {
		return nil, err
	}
	return webhook, nil
}

// DeleteWebhook deletes a webhook
func (api *Client) DeleteWebhook(projectIDOrKey interface{}, webhookID int) (*Webhook, error) {
	return api.DeleteWebhookContext(context.Background(), projectIDOrKey, webhookID)
}

// DeleteWebhookContext updates a webhook with context
func (api *Client) DeleteWebhookContext(ctx context.Context, projectIDOrKey interface{}, webhookID int) (*Webhook, error) {
	projIDOrKey, err := projIDOrKey(projectIDOrKey)
	if err != nil {
		return nil, err
	}

	webhook := new(Webhook)
	if err = api.deleteMethod(ctx, "/api/v2/projects/"+projIDOrKey+"/webhooks/"+strconv.Itoa(webhookID), url.Values{}, &webhook); err != nil {
		return nil, err
	}
	return webhook, nil
}

// AddWebhookInput contains all the parameters necessary (including the optional ones) for a AddWebhook() request.
type AddWebhookInput struct {
	ProjectIDOrKey  interface{} `required:"true"`
	Name            *string     `required:"true"`
	Description     *string     `required:"false"`
	HookURL         *string     `required:"true"`
	AllEvent        *bool       `required:"false"`
	ActivityTypeIDs []int       `required:"false"`
}

// UpdateWebhookInput contains all the parameters necessary (including the optional ones) for a UpdateWebhook() request.
type UpdateWebhookInput struct {
	ProjectIDOrKey  interface{} `required:"true"`
	WebhookID       *int        `required:"true"`
	Name            *string     `required:"true"`
	Description     *string     `required:"false"`
	HookURL         *string     `required:"true"`
	AllEvent        *bool       `required:"false"`
	ActivityTypeIDs []int       `required:"false"`
}
