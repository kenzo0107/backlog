package backlog

import (
	"context"
	"fmt"
)

// Category : category
type Category struct {
	ID           *int    `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
	DisplayOrder *int    `json:"displayOrder,omitempty"`
}

// GetCategories returns the list of categories
func (c *Client) GetCategories(projectIDOrKey interface{}) ([]*Category, error) {
	return c.GetCategoriesContext(context.Background(), projectIDOrKey)
}

// GetCategoriesContext returns the list of categories with context
func (c *Client) GetCategoriesContext(ctx context.Context, projectIDOrKey interface{}) ([]*Category, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/categories", projectIDOrKey)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	categories := []*Category{}
	if err := c.Do(ctx, req, &categories); err != nil {
		return nil, err
	}
	return categories, nil
}

// CreateCategory creates a category
func (c *Client) CreateCategory(projectIDOrKey interface{}, input *CreateCategoryInput) (*Category, error) {
	return c.CreateCategoryContext(context.Background(), projectIDOrKey, input)
}

// CreateCategoryContext creates a category with Context
func (c *Client) CreateCategoryContext(ctx context.Context, projectIDOrKey interface{}, input *CreateCategoryInput) (*Category, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/categories", projectIDOrKey)

	req, err := c.NewRequest("POST", u, input)
	if err != nil {
		return nil, err
	}

	category := new(Category)
	if err := c.Do(ctx, req, &category); err != nil {
		return nil, err
	}
	return category, nil
}

// UpdateCategory updates a category
func (c *Client) UpdateCategory(projectIDOrKey interface{}, categoryID int, input *UpdateCategoryInput) (*Category, error) {
	return c.UpdateCategoryContext(context.Background(), projectIDOrKey, categoryID, input)
}

// UpdateCategoryContext updates a category with Context
func (c *Client) UpdateCategoryContext(ctx context.Context, projectIDOrKey interface{}, categoryID int, input *UpdateCategoryInput) (*Category, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/categories/%v", projectIDOrKey, categoryID)

	req, err := c.NewRequest("PATCH", u, input)
	if err != nil {
		return nil, err
	}

	category := new(Category)
	if err := c.Do(ctx, req, &category); err != nil {
		return nil, err
	}
	return category, nil
}

// DeleteCategory deletes a category
func (c *Client) DeleteCategory(projectIDOrKey interface{}, categoryID int) (*Category, error) {
	return c.DeleteCategoryContext(context.Background(), projectIDOrKey, categoryID)
}

// DeleteCategoryContext deletes a category with Context
func (c *Client) DeleteCategoryContext(ctx context.Context, projectIDOrKey interface{}, categoryID int) (*Category, error) {
	u := fmt.Sprintf("/api/v2/projects/%v/categories/%v", projectIDOrKey, categoryID)

	req, err := c.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	category := new(Category)
	if err := c.Do(ctx, req, &category); err != nil {
		return nil, err
	}
	return category, nil
}

// CreateCategoryInput specifies parameters to the CreateCategory method.
type CreateCategoryInput struct {
	Name *string `json:"name"`
}

// UpdateCategoryInput specifies parameters to the UpdateCategory method.
type UpdateCategoryInput struct {
	Name *string `json:"name"`
}
