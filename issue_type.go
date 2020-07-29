package backlog

import (
	"context"
	"fmt"
)

// IssueType : issue type
type IssueType struct {
	ID                  *int    `json:"id,omitempty"`
	ProjectID           *int    `json:"projectId,omitempty"`
	Name                *string `json:"name,omitempty"`
	Color               *string `json:"color,omitempty"`
	DisplayOrder        *int    `json:"displayOrder,omitempty"`
	TemplateSummary     *string `json:"templateSummary,omitempty"`
	TemplateDescription *string `json:"templateDescription,omitempty"`
}

// GetIssueTypes returns the list of categories
func (c *Client) GetIssueTypes(projectIDOrKey interface{}) ([]*IssueType, error) {
	return c.GetIssueTypesContext(context.Background(), projectIDOrKey)
}

// GetIssueTypesContext returns the list of categories with context
func (c *Client) GetIssueTypesContext(ctx context.Context, projectIDOrKey interface{}) ([]*IssueType, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/issueTypes", projectIDOrKey)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	issueTypes := []*IssueType{}
	if err := c.Do(ctx, req, &issueTypes); err != nil {
		return nil, err
	}
	return issueTypes, nil
}

// CreateIssueType creates an issue type
func (c *Client) CreateIssueType(projectIDOrKey interface{}, input *CreateIssueTypeInput) (*IssueType, error) {
	return c.CreateIssueTypeContext(context.Background(), projectIDOrKey, input)
}

// CreateIssueTypeContext creates an issue type with Context
func (c *Client) CreateIssueTypeContext(ctx context.Context, projectIDOrKey interface{}, input *CreateIssueTypeInput) (*IssueType, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/issueTypes", projectIDOrKey)

	req, err := c.NewRequest("POST", u, input)
	if err != nil {
		return nil, err
	}

	issueType := new(IssueType)
	if err := c.Do(ctx, req, &issueType); err != nil {
		return nil, err
	}
	return issueType, nil
}

// UpdateIssueType updates an issue type
func (c *Client) UpdateIssueType(projectIDOrKey interface{}, issueTypeID int, input *UpdateIssueTypeInput) (*IssueType, error) {
	return c.UpdateIssueTypeContext(context.Background(), projectIDOrKey, issueTypeID, input)
}

// UpdateIssueTypeContext updates an issue type with Context
func (c *Client) UpdateIssueTypeContext(ctx context.Context, projectIDOrKey interface{}, issueTypeID int, input *UpdateIssueTypeInput) (*IssueType, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/issueTypes/%v", projectIDOrKey, issueTypeID)

	req, err := c.NewRequest("PATCH", u, input)
	if err != nil {
		return nil, err
	}

	issueType := new(IssueType)
	if err := c.Do(ctx, req, &issueType); err != nil {
		return nil, err
	}
	return issueType, nil
}

// DeleteIssueType deletes an issue type
func (c *Client) DeleteIssueType(projectIDOrKey interface{}, issueTypeID int, input *DeleteIssueTypeInput) (*IssueType, error) {
	return c.DeleteIssueTypeContext(context.Background(), projectIDOrKey, issueTypeID, input)
}

// DeleteIssueTypeContext deletes an issue type with Context
func (c *Client) DeleteIssueTypeContext(ctx context.Context, projectIDOrKey interface{}, issueTypeID int, input *DeleteIssueTypeInput) (*IssueType, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/issueTypes/%v", projectIDOrKey, issueTypeID)

	req, err := c.NewRequest("DELETE", u, input)
	if err != nil {
		return nil, err
	}

	issueType := new(IssueType)
	if err := c.Do(ctx, req, &issueType); err != nil {
		return nil, err
	}
	return issueType, nil
}

// CreateIssueTypeInput specifies parameters to the CreateIssueType method.
type CreateIssueTypeInput struct {
	Name                *string `json:"name"`
	Color               *string `json:"color"`
	TemplateSummary     *string `json:"templateSummary,omitempty"`
	TemplateDescription *string `json:"templateDescription,omitempty"`
}

// UpdateIssueTypeInput specifies parameters to the UpdateIssueType method.
type UpdateIssueTypeInput struct {
	Name                *string `json:"name,omitempty"`
	Color               *string `json:"color,omitempty"`
	TemplateSummary     *string `json:"templateSummary,omitempty"`
	TemplateDescription *string `json:"templateDescription,omitempty"`
}

// DeleteIssueTypeInput specifies parameters to the DeleteIssueType method.
type DeleteIssueTypeInput struct {
	SubstituteIssueTypeID *int `json:"substituteIssueTypeId,omitempty"`
}
