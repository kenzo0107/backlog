package backlog

import (
	"context"
	"net/url"
	"strconv"
)

// Wiki : wiki
type Wiki struct {
	ID          *int          `json:"id"`
	ProjectID   *int          `json:"projectId"`
	Name        *string       `json:"name"`
	Content     *string       `json:"content"`
	Tags        []*Tag        `json:"tags"`
	Attachments []*Attachment `json:"attachments"`
	SharedFiles []*SharedFile `json:"sharedFiles"`
	Stars       []*Star       `json:"stars"`
	CreatedUser *User         `json:"createdUser"`
	Created     *Timestamp    `json:"created"`
	UpdatedUser *User         `json:"updatedUser"`
	Updated     *Timestamp    `json:"updated"`
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
func (api *Client) GetWikis(projectIDOrKey interface{}, keyword string) ([]*Wiki, error) {
	return api.GetWikisContext(context.Background(), projectIDOrKey, keyword)
}

// GetWikisContext returns the list of wikis
func (api *Client) GetWikisContext(ctx context.Context, projectIDOrKey interface{}, keyword string) ([]*Wiki, error) {
	projIDOrKey, err := projIDOrKey(projectIDOrKey)
	if err != nil {
		return nil, err
	}

	values := url.Values{}
	values.Add("projectIdOrKey", projIDOrKey)
	if keyword != "" {
		values.Add("keyword", keyword)
	}

	r := []*Wiki{}
	if err = api.getMethod(ctx, "/api/v2/wikis", values, &r); err != nil {
		return nil, err
	}
	return r, nil
}

// GetWikiCount returns the number of wikis
func (api *Client) GetWikiCount(projectIDOrKey interface{}) (int, error) {
	return api.GetWikiCountContext(context.Background(), projectIDOrKey)
}

// GetWikiCountContext returns the number of wikis
func (api *Client) GetWikiCountContext(ctx context.Context, projectIDOrKey interface{}) (int, error) {
	projIDOrKey, err := projIDOrKey(projectIDOrKey)
	if err != nil {
		return 0, err
	}

	values := url.Values{}
	values.Add("projectIdOrKey", projIDOrKey)

	page := &Page{}
	if err := api.getMethod(ctx, "/api/v2/wikis/count", values, &page); err != nil {
		return 0, err
	}

	return *page.Count, nil
}

// GetWikiTags returns the tags of wikis
func (api *Client) GetWikiTags(projectIDOrKey interface{}) ([]*Tag, error) {
	return api.GetWikiTagsContext(context.Background(), projectIDOrKey)
}

// GetWikiTagsContext returns the tags of wikis
func (api *Client) GetWikiTagsContext(ctx context.Context, projectIDOrKey interface{}) ([]*Tag, error) {
	projIDOrKey, err := projIDOrKey(projectIDOrKey)
	if err != nil {
		return nil, err
	}

	values := url.Values{}
	values.Add("projectIdOrKey", projIDOrKey)

	r := []*Tag{}
	if err := api.getMethod(ctx, "/api/v2/wikis/tags", values, &r); err != nil {
		return nil, err
	}
	return r, nil
}

// GetWiki returns wiki by id
func (api *Client) GetWiki(wikiID int) (*Wiki, error) {
	return api.GetWikiContext(context.Background(), wikiID)
}

// GetWikiContext returns wiki by id
func (api *Client) GetWikiContext(ctx context.Context, wikiID int) (*Wiki, error) {
	wiki := new(Wiki)
	if err := api.getMethod(ctx, "/api/v2/wikis/"+strconv.Itoa(wikiID), url.Values{}, &wiki); err != nil {
		return nil, err
	}
	return wiki, nil
}

// CreateWiki creates a wiki
func (api *Client) CreateWiki(input *CreateWikiInput) (*Wiki, error) {
	return api.CreateWikiContext(context.Background(), input)
}

// CreateWikiContext creates a wiki with Context
func (api *Client) CreateWikiContext(ctx context.Context, input *CreateWikiInput) (*Wiki, error) {
	values := url.Values{
		"projectId": {strconv.Itoa(*input.ProjectID)},
		"name":      {*input.Name},
		"content":   {*input.Content},
	}

	if input.MailNotify != nil {
		values.Add("mailNotify", strconv.FormatBool(*input.MailNotify))
	}

	wiki := new(Wiki)
	if err := api.postMethod(ctx, "/api/v2/wikis", values, &wiki); err != nil {
		return nil, err
	}
	return wiki, nil
}

// UpdateWiki updates a wiki
func (api *Client) UpdateWiki(input *UpdateWikiInput) (*Wiki, error) {
	return api.UpdateWikiContext(context.Background(), input)
}

// UpdateWikiContext updates a wiki with Context
func (api *Client) UpdateWikiContext(ctx context.Context, input *UpdateWikiInput) (*Wiki, error) {
	values := url.Values{
		"name":    {*input.Name},
		"content": {*input.Content},
	}

	if input.MailNotify != nil {
		values.Add("mailNotify", strconv.FormatBool(*input.MailNotify))
	}

	wiki := new(Wiki)
	if err := api.patchMethod(ctx, "/api/v2/wikis/"+strconv.Itoa(*input.WikiID), values, &wiki); err != nil {
		return nil, err
	}
	return wiki, nil
}

// DeleteWiki deletes a wiki
func (api *Client) DeleteWiki(wikiID int) (*Wiki, error) {
	return api.DeleteWikiContext(context.Background(), wikiID)
}

// DeleteWikiContext deletes a wiki with Context
func (api *Client) DeleteWikiContext(ctx context.Context, wikiID int) (*Wiki, error) {
	wiki := new(Wiki)
	if err := api.deleteMethod(ctx, "/api/v2/wikis/"+strconv.Itoa(wikiID), url.Values{}, &wiki); err != nil {
		return nil, err
	}
	return wiki, nil
}

// GetWikiAttachments returns attachements of a wiki
func (api *Client) GetWikiAttachments(wikiID int) ([]*Attachment, error) {
	return api.GetWikiAttachmentsContext(context.Background(), wikiID)
}

// GetWikiAttachmentsContext returns attachements of a wiki with context
func (api *Client) GetWikiAttachmentsContext(ctx context.Context, wikiID int) ([]*Attachment, error) {
	attachments := []*Attachment{}
	if err := api.getMethod(ctx, "/api/v2/wikis/"+strconv.Itoa(wikiID)+"/attachments", url.Values{}, &attachments); err != nil {
		return nil, err
	}
	return attachments, nil
}

// AddAttachmentToWiki adds attachments to a wiki
func (api *Client) AddAttachmentToWiki(input *AddAttachmentToWikiInput) ([]*Attachment, error) {
	return api.AddAttachmentToWikiContext(context.Background(), input)
}

// AddAttachmentToWikiContext adds attachments to a wiki with context
func (api *Client) AddAttachmentToWikiContext(ctx context.Context, input *AddAttachmentToWikiInput) ([]*Attachment, error) {
	values := url.Values{}
	for _, attachmentID := range input.AttachmentIDs {
		values.Add("attachmentId[]", strconv.Itoa(attachmentID))
	}

	attachements := []*Attachment{}
	if err := api.postMethod(ctx, "/api/v2/wikis/"+strconv.Itoa(*input.WikiID)+"/attachments", values, &attachements); err != nil {
		return nil, err
	}
	return attachements, nil
}

// CreateWikiInput contains all the parameters necessary (including the optional ones) for a CreateWiki() request.
type CreateWikiInput struct {
	ProjectID  *int    `required:"true"`
	Name       *string `required:"true"`
	Content    *string `required:"true"`
	MailNotify *bool   `required:"false"`
}

// UpdateWikiInput contains all the parameters necessary (including the optional ones) for a UpdateWiki() request.
type UpdateWikiInput struct {
	WikiID     *int    `required:"true"`
	Name       *string `required:"true"`
	Content    *string `required:"true"`
	MailNotify *bool   `required:"false"`
}

// AddAttachmentToWikiInput contains all the parameters necessary (including the optional ones) for a AddAttachmentToWiki() request.
type AddAttachmentToWikiInput struct {
	WikiID        *int  `required:"true"`
	AttachmentIDs []int `required:"true"`
}
