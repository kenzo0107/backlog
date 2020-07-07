package backlog

import (
	"context"
	"fmt"
)

// Wiki : wiki
type Wiki struct {
	ID          *int          `json:"id,omitempty"`
	ProjectID   *int          `json:"projectId,omitempty"`
	Name        *string       `json:"name,omitempty"`
	Content     *string       `json:"content,omitempty"`
	Tags        []*Tag        `json:"tags,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`
	SharedFiles []*SharedFile `json:"sharedFiles,omitempty"`
	Stars       []*Star       `json:"stars,omitempty"`
	CreatedUser *User         `json:"createdUser,omitempty"`
	Created     *Timestamp    `json:"created,omitempty"`
	UpdatedUser *User         `json:"updatedUser,omitempty"`
	Updated     *Timestamp    `json:"updated,omitempty"`
}

// Tag : tag
type Tag struct {
	ID   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Attachment : attachment of wiki
type Attachment struct {
	ID          *int       `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Size        *int       `json:"size,omitempty"`
	CreatedUser *User      `json:"createdUser,omitempty"`
	Created     *Timestamp `json:"created,omitempty"`
}

// SharedFile : shared file of wiki
type SharedFile struct {
	ID          *int       `json:"id,omitempty"`
	Type        *string    `json:"type,omitempty"`
	Dir         *string    `json:"dir,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Size        *int       `json:"size,omitempty"`
	CreatedUser *User      `json:"createdUser,omitempty"`
	Created     *Timestamp `json:"created,omitempty"`
	UpdatedUser *User      `json:"updatedUser,omitempty"`
	Updated     *Timestamp `json:"updated,omitempty"`
}

// Star : star of wiki
type Star struct {
	ID        *int       `json:"id,omitempty"`
	Comment   *string    `json:"comment,omitempty"`
	URL       *string    `json:"url,omitempty"`
	Title     *string    `json:"title,omitempty"`
	Presenter *User      `json:"presenter,omitempty"`
	Created   *Timestamp `json:"created,omitempty"`
}

// Page : wiki page information
type Page struct {
	Count *int `json:"count,omitempty"`
}

// GetWikis returns the list of wikis
func (c *Client) GetWikis(opts *GetWikisOptions) ([]*Wiki, error) {
	return c.GetWikisContext(context.Background(), opts)
}

// GetWikisContext returns the list of wikis
func (c *Client) GetWikisContext(ctx context.Context, opts *GetWikisOptions) ([]*Wiki, error) {
	u, err := c.AddOptions("/api/v2/wikis", opts)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	wikis := []*Wiki{}
	if err := c.Do(ctx, req, &wikis); err != nil {
		return nil, err
	}
	return wikis, nil
}

// GetWikiCount returns the number of wikis
func (c *Client) GetWikiCount(opts *GetWikiCountOptions) (int, error) {
	return c.GetWikiCountContext(context.Background(), opts)
}

// GetWikiCountContext returns the number of wikis
func (c *Client) GetWikiCountContext(ctx context.Context, opts *GetWikiCountOptions) (int, error) {
	u, err := c.AddOptions("/api/v2/wikis/count", opts)
	if err != nil {
		return 0, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return 0, err
	}

	page := new(Page)
	if err := c.Do(ctx, req, &page); err != nil {
		return 0, err
	}

	return *page.Count, nil
}

// GetWikiTags returns the tags of wikis
func (c *Client) GetWikiTags(opts *GetWikiTagsOptions) ([]*Tag, error) {
	return c.GetWikiTagsContext(context.Background(), opts)
}

// GetWikiTagsContext returns the tags of wikis
func (c *Client) GetWikiTagsContext(ctx context.Context, opts *GetWikiTagsOptions) ([]*Tag, error) {
	u, err := c.AddOptions("/api/v2/wikis/tags", opts)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	tags := []*Tag{}
	if err := c.Do(ctx, req, &tags); err != nil {
		return nil, err
	}
	return tags, nil
}

// GetWiki returns wiki by id
func (c *Client) GetWiki(wikiID int) (*Wiki, error) {
	return c.GetWikiContext(context.Background(), wikiID)
}

// GetWikiContext returns wiki by id
func (c *Client) GetWikiContext(ctx context.Context, wikiID int) (*Wiki, error) {
	u := fmt.Sprintf("/api/v2/wikis/%v", wikiID)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	wiki := new(Wiki)
	if err := c.Do(ctx, req, &wiki); err != nil {
		return nil, err
	}
	return wiki, nil
}

// CreateWiki creates a wiki
func (c *Client) CreateWiki(input *CreateWikiInput) (*Wiki, error) {
	return c.CreateWikiContext(context.Background(), input)
}

// CreateWikiContext creates a wiki with Context
func (c *Client) CreateWikiContext(ctx context.Context, input *CreateWikiInput) (*Wiki, error) {
	u := "/api/v2/wikis"

	req, err := c.NewRequest("POST", u, input)
	if err != nil {
		return nil, err
	}

	wiki := new(Wiki)
	if err := c.Do(ctx, req, &wiki); err != nil {
		return nil, err
	}
	return wiki, nil
}

// UpdateWiki updates a wiki
func (c *Client) UpdateWiki(wikiID int, input *UpdateWikiInput) (*Wiki, error) {
	return c.UpdateWikiContext(context.Background(), wikiID, input)
}

// UpdateWikiContext updates a wiki with Context
func (c *Client) UpdateWikiContext(ctx context.Context, wikiID int, input *UpdateWikiInput) (*Wiki, error) {
	u := fmt.Sprintf("/api/v2/wikis/%v", wikiID)

	req, err := c.NewRequest("PATCH", u, input)
	if err != nil {
		return nil, err
	}

	wiki := new(Wiki)
	if err := c.Do(ctx, req, &wiki); err != nil {
		return nil, err
	}
	return wiki, nil
}

// DeleteWiki deletes a wiki
func (c *Client) DeleteWiki(wikiID int) (*Wiki, error) {
	return c.DeleteWikiContext(context.Background(), wikiID)
}

// DeleteWikiContext deletes a wiki with Context
func (c *Client) DeleteWikiContext(ctx context.Context, wikiID int) (*Wiki, error) {
	u := fmt.Sprintf("/api/v2/wikis/%v", wikiID)

	req, err := c.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	wiki := new(Wiki)
	if err := c.Do(ctx, req, &wiki); err != nil {
		return nil, err
	}
	return wiki, nil
}

// GetWikiAttachments returns attachements of a wiki
func (c *Client) GetWikiAttachments(wikiID int) ([]*Attachment, error) {
	return c.GetWikiAttachmentsContext(context.Background(), wikiID)
}

// GetWikiAttachmentsContext returns attachements of a wiki with context
func (c *Client) GetWikiAttachmentsContext(ctx context.Context, wikiID int) ([]*Attachment, error) {
	u := fmt.Sprintf("/api/v2/wikis/%v/attachments", wikiID)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	attachments := []*Attachment{}
	if err := c.Do(ctx, req, &attachments); err != nil {
		return nil, err
	}
	return attachments, nil
}

// AddAttachmentToWiki adds attachments to a wiki
func (c *Client) AddAttachmentToWiki(wikiID int, input *AddAttachmentToWikiInput) ([]*Attachment, error) {
	return c.AddAttachmentToWikiContext(context.Background(), wikiID, input)
}

// AddAttachmentToWikiContext adds attachments to a wiki with context
func (c *Client) AddAttachmentToWikiContext(ctx context.Context, wikiID int, input *AddAttachmentToWikiInput) ([]*Attachment, error) {
	u := fmt.Sprintf("/api/v2/wikis/%v/attachments", wikiID)

	req, err := c.NewRequest("POST", u, input)
	if err != nil {
		return nil, err
	}

	attachments := []*Attachment{}
	if err := c.Do(ctx, req, &attachments); err != nil {
		return nil, err
	}
	return attachments, nil
}

// GetWikisOptions specifies optional parameters to the GetWikis method.
type GetWikisOptions struct {
	ProjectIDOrKey interface{} `url:"projectIdOrKey"`
	Keyword        *string     `url:"keyword,omitempty"`
}

// GetWikiCountOptions specifies optional parameters to the GetWikiCount method.
type GetWikiCountOptions struct {
	ProjectIDOrKey interface{} `url:"projectIdOrKey,omitempty"`
}

// GetWikiTagsOptions specifies optional parameters to the GetWikiTags method.
type GetWikiTagsOptions struct {
	ProjectIDOrKey interface{} `url:"projectIdOrKey,omitempty"`
}

// CreateWikiInput contains all the parameters necessary (including the optional ones) for a CreateWiki() request.
type CreateWikiInput struct {
	ProjectID  *int    `json:"projectId"`
	Name       *string `json:"name"`
	Content    *string `json:"content"`
	MailNotify *bool   `json:"mailNotify,omitempty"`
}

// UpdateWikiInput contains all the parameters necessary (including the optional ones) for a UpdateWiki() request.
type UpdateWikiInput struct {
	Name       *string `json:"name"`
	Content    *string `json:"content"`
	MailNotify *bool   `json:"mailNotify,omitempty"`
}

// AddAttachmentToWikiInput contains all the parameters necessary (including the optional ones) for a AddAttachmentToWiki() request.
type AddAttachmentToWikiInput struct {
	AttachmentIDs []int `json:"attachmentId"`
}
