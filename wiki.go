package backlog

import (
	"context"
	"net/url"
	"strconv"
)

// Wiki : wiki
type Wiki struct {
	ID          int          `json:"id"`
	ProjectID   int          `json:"projectId"`
	Name        string       `json:"name"`
	Content     string       `json:"content"`
	Tags        []Tag        `json:"tags"`
	Attachments []Attachment `json:"attachments"`
	SharedFiles []SharedFile `json:"sharedFiles"`
	Stars       []Star       `json:"stars"`
	CreatedUser User         `json:"createdUser"`
	Created     JSONTime     `json:"created"`
	UpdatedUser User         `json:"updatedUser"`
	Updated     JSONTime     `json:"updated"`
}

// Tag : tag
type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Attachment : attachment of wiki
type Attachment struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Size        int      `json:"size"`
	CreatedUser User     `json:"createdUser"`
	Created     JSONTime `json:"created"`
}

// SharedFile : shared file of wiki
type SharedFile struct {
	ID          int      `json:"id"`
	Type        string   `json:"type"`
	Dir         string   `json:"dir"`
	Name        string   `json:"name"`
	Size        int      `json:"size"`
	CreatedUser User     `json:"createdUser"`
	Created     JSONTime `json:"created"`
	UpdatedUser User     `json:"updatedUser"`
	Updated     JSONTime `json:"updated"`
}

// Star : star of wiki
type Star struct {
	ID        int         `json:"id"`
	Comment   interface{} `json:"comment"`
	URL       string      `json:"url"`
	Title     string      `json:"title"`
	Presenter User        `json:"presenter"`
	Created   JSONTime    `json:"created"`
}

// Count : count
type Count struct {
	Count int `json:"count,omitempty"`
}

// GetWikis returns the list of wikis
func (api *Client) GetWikis(projectIDOrKey interface{}, keyword string) ([]Wiki, error) {
	return api.GetWikisContext(context.Background(), projectIDOrKey, keyword)
}

// GetWikisContext returns the list of wikis
func (api *Client) GetWikisContext(ctx context.Context, projectIDOrKey interface{}, keyword string) (wikis []Wiki, err error) {
	projIDOrKey, err := projIDOrKey(projectIDOrKey)
	if err != nil {
		return nil, err
	}

	values := url.Values{}
	values.Add("projectIdOrKey", projIDOrKey)
	values.Add("keyword", keyword)
	if err = api.getMethod(ctx, "/api/v2/wikis", values, &wikis); err != nil {
		return nil, err
	}
	return wikis, nil
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

	c := new(Count)
	if err := api.getMethod(ctx, "/api/v2/wikis/count", values, &c); err != nil {
		return 0, err
	}
	return c.Count, nil
}

// GetWikiTags returns the tags of wikis
func (api *Client) GetWikiTags(projectIDOrKey interface{}) ([]Tag, error) {
	return api.GetWikiTagsContext(context.Background(), projectIDOrKey)
}

// GetWikiTagsContext returns the tags of wikis
func (api *Client) GetWikiTagsContext(ctx context.Context, projectIDOrKey interface{}) (tags []Tag, err error) {
	projIDOrKey, err := projIDOrKey(projectIDOrKey)
	if err != nil {
		return tags, err
	}

	values := url.Values{}
	values.Add("projectIdOrKey", projIDOrKey)
	if err := api.getMethod(ctx, "/api/v2/wikis/tags", values, &tags); err != nil {
		return nil, err
	}
	return tags, nil
}

// GetWikiByID returns wiki by id
func (api *Client) GetWikiByID(id int) (Wiki, error) {
	return api.GetWikiByIDContext(context.Background(), id)
}

// GetWikiByIDContext returns wiki by id
func (api *Client) GetWikiByIDContext(ctx context.Context, id int) (wiki Wiki, err error) {
	if err = api.getMethod(ctx, "/api/v2/wikis/"+strconv.Itoa(id), url.Values{}, &wiki); err != nil {
		return wiki, err
	}
	return wiki, nil
}

// func (api *Client) AddWiki(projectID int, name, content string, mailNotify bool) {

// }
