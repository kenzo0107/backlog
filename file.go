package backlog

import (
	"context"
	"net/url"
)

// FileUploadResponse : response of uploading file
type FileUploadResponse struct {
	ID   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Size *int    `json:"size,omitempty"`
}

// UploadFile uploads a file
func (c *Client) UploadFile(fpath string) (*FileUploadResponse, error) {
	return c.UploadFileContext(context.Background(), fpath)
}

// UploadFileContext uploads a file and setting a custom context
func (c *Client) UploadFileContext(ctx context.Context, fpath string) (*FileUploadResponse, error) {
	fileUploadResponse := new(FileUploadResponse)
	if err := postLocalWithMultipartResponse(ctx, c.httpclient, c.endpoint+"/api/v2/space/attachment?apiKey="+c.apiKey, fpath, "file", url.Values{}, &fileUploadResponse, c); err != nil {
		return nil, err
	}
	return fileUploadResponse, nil
}
