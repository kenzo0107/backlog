package backlog

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
