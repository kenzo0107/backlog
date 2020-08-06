package backlog

import (
	"context"
	"fmt"
)

// CustomField : custom field
type CustomField struct {
	ID                   *int    `json:"id,omitempty"`
	TypeID               *int    `json:"typeId,omitempty"`
	Name                 *string `json:"name,omitempty"`
	Description          *string `json:"description,omitempty"`
	Required             *bool   `json:"required,omitempty"`
	ApplicableIssueTypes []int   `json:"applicableIssueTypes,omitempty"`
	AllowAddItem         *bool   `json:"allowAddItem,omitempty"`
	Items                []*Item `json:"items,omitempty"`
}

// Item : item
type Item struct {
	ID           *int    `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
	DisplayOrder *int    `json:"displayOrder,omitempty"`
}

// GetCustomFields returns the list of custom fields
func (c *Client) GetCustomFields(projectIDOrKey interface{}) ([]*CustomField, error) {
	return c.GetCustomFieldsContext(context.Background(), projectIDOrKey)
}

// GetCustomFieldsContext returns the list of custom fields with context
func (c *Client) GetCustomFieldsContext(ctx context.Context, projectIDOrKey interface{}) ([]*CustomField, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/customFields", projectIDOrKey)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var customFields []*CustomField
	if err := c.Do(ctx, req, &customFields); err != nil {
		return nil, err
	}
	return customFields, nil
}
