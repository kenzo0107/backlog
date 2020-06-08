package backlog

import (
	"context"
	"io"
	"net/url"
)

// Space : backlog space
type Space struct {
	SpaceKey           string   `json:"spaceKey"`
	Name               string   `json:"name"`
	OwnerID            int      `json:"ownerId"`
	Lang               string   `json:"lang"`
	Timezone           string   `json:"timezone"`
	ReportSendTime     string   `json:"reportSendTime"`
	TextFormattingRule string   `json:"textFormattingRule"`
	Created            JSONTime `json:"created"`
	Updated            JSONTime `json:"updated"`
}

// SpaceNotification : backlog space notification
type SpaceNotification struct {
	Content string   `json:"content"`
	Updated JSONTime `json:"updated"`
}

// SpaceDiskUsage : disk usage of space
type SpaceDiskUsage struct {
	Capacity   int                    `json:"capacity"`
	Issue      int                    `json:"issue"`
	Wiki       int                    `json:"wiki"`
	File       int                    `json:"file"`
	Subversion int                    `json:"subversion"`
	Git        int                    `json:"git"`
	GitLFS     int                    `json:"gitLFS"`
	Details    []SpaceDiskUsageDetail `json:"details"`
}

// SpaceDiskUsageDetail : the detail of disk usage of a space
type SpaceDiskUsageDetail struct {
	ProjectID  int `json:"projectId"`
	Issue      int `json:"issue"`
	Wiki       int `json:"wiki"`
	File       int `json:"file"`
	Subversion int `json:"subversion"`
	Git        int `json:"git"`
	GitLFS     int `json:"gitLFS"`
}

// GetSpace returns backlog space
func (api *Client) GetSpace() (*Space, error) {
	return api.GetSpaceContext(context.Background())
}

// GetSpaceContext returns backlog space with context
func (api *Client) GetSpaceContext(ctx context.Context) (*Space, error) {
	space := Space{}
	if err := api.getMethod(ctx, "/api/v2/space", url.Values{}, &space); err != nil {
		return nil, err
	}
	return &space, nil
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
	spaceNotification := SpaceNotification{}
	if err := api.getMethod(ctx, "/api/v2/space/notification", url.Values{}, &spaceNotification); err != nil {
		return nil, err
	}
	return &spaceNotification, nil
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
	spaceNotification := SpaceNotification{}
	if err := api.putMethod(ctx, "/api/v2/space/notification", values, &spaceNotification); err != nil {
		return nil, err
	}
	return &spaceNotification, nil
}

// GetSpaceDiskUsage returns the disk usage of a space
func (api *Client) GetSpaceDiskUsage() (*SpaceDiskUsage, error) {
	return api.GetSpaceDiskUsageContext(context.Background())
}

// GetSpaceDiskUsageContext returns the disk usage of a space with context
func (api *Client) GetSpaceDiskUsageContext(ctx context.Context) (*SpaceDiskUsage, error) {
	diskUsage := SpaceDiskUsage{}
	if err := api.getMethod(ctx, "/api/v2/space/diskUsage", url.Values{}, &diskUsage); err != nil {
		return nil, err
	}
	return &diskUsage, nil
}
