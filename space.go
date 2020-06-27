package backlog

import (
	"context"
	"io"
	"net/url"
)

// Space : backlog space
type Space struct {
	SpaceKey           *string    `json:"spaceKey,omitempty"`
	Name               *string    `json:"name,omitempty"`
	OwnerID            *int       `json:"ownerId,omitempty"`
	Lang               *string    `json:"lang,omitempty"`
	Timezone           *string    `json:"timezone,omitempty"`
	ReportSendTime     *string    `json:"reportSendTime,omitempty"`
	TextFormattingRule *string    `json:"textFormattingRule,omitempty"`
	Created            *Timestamp `json:"created,omitempty"`
	Updated            *Timestamp `json:"updated,omitempty"`
}

// SpaceNotification : backlog space notification
type SpaceNotification struct {
	Content *string    `json:"content,omitempty"`
	Updated *Timestamp `json:"updated,omitempty"`
}

// SpaceDiskUsage : disk usage of space
type SpaceDiskUsage struct {
	Capacity   *int                    `json:"capacity,omitempty"`
	Issue      *int                    `json:"issue,omitempty"`
	Wiki       *int                    `json:"wiki,omitempty"`
	File       *int                    `json:"file,omitempty"`
	Subversion *int                    `json:"subversion,omitempty"`
	Git        *int                    `json:"git,omitempty"`
	GitLFS     *int                    `json:"gitLFS,omitempty"`
	Details    []*SpaceDiskUsageDetail `json:"details,omitempty"`
}

// SpaceDiskUsageDetail : the detail of disk usage of a space
type SpaceDiskUsageDetail struct {
	ProjectID  *int `json:"projectId,omitempty"`
	Issue      *int `json:"issue,omitempty"`
	Wiki       *int `json:"wiki,omitempty"`
	File       *int `json:"file,omitempty"`
	Subversion *int `json:"subversion,omitempty"`
	Git        *int `json:"git,omitempty"`
	GitLFS     *int `json:"gitLFS,omitempty"`
}

// GetSpace returns backlog space
func (api *Client) GetSpace() (*Space, error) {
	return api.GetSpaceContext(context.Background())
}

// GetSpaceContext returns backlog space with context
func (api *Client) GetSpaceContext(ctx context.Context) (*Space, error) {
	space := new(Space)
	if err := api.getMethod(ctx, "/api/v2/space", url.Values{}, &space); err != nil {
		return nil, err
	}
	return space, nil
}

// GetSpaceIcon downloads space icon
func (api *Client) GetSpaceIcon(writer io.Writer) error {
	return api.GetSpaceIconContext(context.Background(), writer)
}

// GetSpaceIconContext downloads space icon with context
func (api *Client) GetSpaceIconContext(ctx context.Context, writer io.Writer) error {
	return downloadFile(ctx, api.httpclient, api.apiKey, api.endpoint+"/api/v2/space/image", writer, api)
}

// GetSpaceNotification returns a space notification
func (api *Client) GetSpaceNotification() (*SpaceNotification, error) {
	return api.GetSpaceNotificationContext(context.Background())
}

// GetSpaceNotificationContext returns a space notification with context
func (api *Client) GetSpaceNotificationContext(ctx context.Context) (*SpaceNotification, error) {
	spaceNotification := new(SpaceNotification)
	if err := api.getMethod(ctx, "/api/v2/space/notification", url.Values{}, &spaceNotification); err != nil {
		return nil, err
	}
	return spaceNotification, nil
}

// UpdateSpaceNotification updates a space notification
func (api *Client) UpdateSpaceNotification(content string) (*SpaceNotification, error) {
	return api.UpdateSpaceNotificationContext(context.Background(), content)
}

// UpdateSpaceNotificationContext updates a space notification with context
func (api *Client) UpdateSpaceNotificationContext(ctx context.Context, content string) (*SpaceNotification, error) {
	values := url.Values{
		"content": {content},
	}

	spaceNotification := new(SpaceNotification)
	if err := api.putMethod(ctx, "/api/v2/space/notification", values, &spaceNotification); err != nil {
		return nil, err
	}
	return spaceNotification, nil
}

// GetSpaceDiskUsage returns the disk usage of a space
func (api *Client) GetSpaceDiskUsage() (*SpaceDiskUsage, error) {
	return api.GetSpaceDiskUsageContext(context.Background())
}

// GetSpaceDiskUsageContext returns the disk usage of a space with context
func (api *Client) GetSpaceDiskUsageContext(ctx context.Context) (*SpaceDiskUsage, error) {
	diskUsage := new(SpaceDiskUsage)
	if err := api.getMethod(ctx, "/api/v2/space/diskUsage", url.Values{}, &diskUsage); err != nil {
		return nil, err
	}
	return diskUsage, nil
}
