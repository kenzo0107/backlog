package backlog

import (
	"context"
	"io"
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

// License : license
type License struct {
	Active                            *bool      `json:"active,omitempty"`
	AttachmentLimit                   *int       `json:"attachmentLimit,omitempty"`
	AttachmentLimitPerFile            *int       `json:"attachmentLimitPerFile,omitempty"`
	AttachmentNumLimit                *int       `json:"attachmentNumLimit,omitempty"`
	Attribute                         *bool      `json:"attribute,omitempty"`
	AttributeLimit                    *int       `json:"attributeLimit,omitempty"`
	Burndown                          *bool      `json:"burndown,omitempty"`
	CommentLimit                      *int       `json:"commentLimit,omitempty"`
	ComponentLimit                    *int       `json:"componentLimit,omitempty"`
	FileSharing                       *bool      `json:"fileSharing,omitempty"`
	Gantt                             *bool      `json:"gantt,omitempty"`
	Git                               *bool      `json:"git,omitempty"`
	IssueLimit                        *int       `json:"issueLimit,omitempty"`
	LicenceTypeID                     *int       `json:"licenceTypeId,omitempty"`
	LimitDate                         *Timestamp `json:"limitDate,omitempty"`
	NulabAccount                      *bool      `json:"nulabAccount,omitempty"`
	ParentChildIssue                  *bool      `json:"parentChildIssue,omitempty"`
	PostIssueByMail                   *bool      `json:"postIssueByMail,omitempty"`
	ProjectGroup                      *bool      `json:"projectGroup,omitempty"`
	ProjectLimit                      *int       `json:"projectLimit,omitempty"`
	PullRequestAttachmentLimitPerFile *int       `json:"pullRequestAttachmentLimitPerFile,omitempty"`
	PullRequestAttachmentNumLimit     *int       `json:"pullRequestAttachmentNumLimit,omitempty"`
	RemoteAddress                     *bool      `json:"remoteAddress,omitempty"`
	RemoteAddressLimit                *int       `json:"remoteAddressLimit,omitempty"`
	StartedOn                         *Timestamp `json:"startedOn,omitempty"`
	StorageLimit                      *int64     `json:"storageLimit,omitempty"`
	Subversion                        *bool      `json:"subversion,omitempty"`
	SubversionExternal                *bool      `json:"subversionExternal,omitempty"`
	UserLimit                         *int       `json:"userLimit,omitempty"`
	VersionLimit                      *int       `json:"versionLimit,omitempty"`
	WikiAttachment                    *bool      `json:"wikiAttachment,omitempty"`
	WikiAttachmentLimitPerFile        *int       `json:"wikiAttachmentLimitPerFile,omitempty"`
	WikiAttachmentNumLimit            *int       `json:"wikiAttachmentNumLimit,omitempty"`
}

// GetSpace returns backlog space
func (c *Client) GetSpace() (*Space, error) {
	return c.GetSpaceContext(context.Background())
}

// GetSpaceContext returns backlog space with context
func (c *Client) GetSpaceContext(ctx context.Context) (*Space, error) {
	u := "/api/v2/space"

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	space := new(Space)
	if err := c.Do(ctx, req, &space); err != nil {
		return nil, err
	}
	return space, nil
}

// GetSpaceIcon downloads space icon
func (c *Client) GetSpaceIcon(writer io.Writer) error {
	return c.GetSpaceIconContext(context.Background(), writer)
}

// GetSpaceIconContext downloads space icon with context
func (c *Client) GetSpaceIconContext(ctx context.Context, writer io.Writer) error {
	u := "/api/v2/space/image"

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, writer); err != nil {
		return err
	}
	return nil
}

// GetSpaceNotification returns a space notification
func (c *Client) GetSpaceNotification() (*SpaceNotification, error) {
	return c.GetSpaceNotificationContext(context.Background())
}

// GetSpaceNotificationContext returns a space notification with context
func (c *Client) GetSpaceNotificationContext(ctx context.Context) (*SpaceNotification, error) {
	u := "/api/v2/space/notification"

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	spaceNotification := new(SpaceNotification)
	if err := c.Do(ctx, req, &spaceNotification); err != nil {
		return nil, err
	}
	return spaceNotification, nil
}

// UpdateSpaceNotification updates a space notification
func (c *Client) UpdateSpaceNotification(input *UpdateSpaceNotificationInput) (*SpaceNotification, error) {
	return c.UpdateSpaceNotificationContext(context.Background(), input)
}

// UpdateSpaceNotificationContext updates a space notification with context
func (c *Client) UpdateSpaceNotificationContext(ctx context.Context, input *UpdateSpaceNotificationInput) (*SpaceNotification, error) {
	u := "/api/v2/space/notification"

	req, err := c.NewRequest("PUT", u, input)
	if err != nil {
		return nil, err
	}

	spaceNotification := new(SpaceNotification)
	if err := c.Do(ctx, req, &spaceNotification); err != nil {
		return nil, err
	}
	return spaceNotification, nil
}

// GetSpaceDiskUsage returns the disk usage of a space
func (c *Client) GetSpaceDiskUsage() (*SpaceDiskUsage, error) {
	return c.GetSpaceDiskUsageContext(context.Background())
}

// GetSpaceDiskUsageContext returns the disk usage of a space with context
func (c *Client) GetSpaceDiskUsageContext(ctx context.Context) (*SpaceDiskUsage, error) {
	u := "/api/v2/space/diskUsage"

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	diskUsage := new(SpaceDiskUsage)
	if err := c.Do(ctx, req, &diskUsage); err != nil {
		return nil, err
	}
	return diskUsage, nil
}

// GetLicence returns the license information
func (c *Client) GetLicence() (*License, error) {
	return c.GetLicenceContext(context.Background())
}

// GetLicenceContext returns the license information with context
func (c *Client) GetLicenceContext(ctx context.Context) (*License, error) {
	u := "/api/v2/space/licence"

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	license := new(License)
	if err := c.Do(ctx, req, &license); err != nil {
		return nil, err
	}
	return license, nil
}

// UpdateSpaceNotificationInput contains all the parameters necessary (including the optional ones) for a UpdateSpaceNotification() request.
type UpdateSpaceNotificationInput struct {
	Content *string `json:"content"`
}
