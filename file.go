package backlog

import (
	"context"
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
	u := "/api/v2/space/attachment"

	fileUploadResponse := new(FileUploadResponse)
	if err := c.UploadMultipartFile(ctx, "POST", u, fpath, "file", &fileUploadResponse); err != nil {
		return nil, err
	}
	return fileUploadResponse, nil
}
