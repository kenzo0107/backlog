package backlog

import (
	"context"
)

// Sort : sort
type Sort string

// IssueType is used in sort type
const (
	SortIssueType = Sort(iota)
	SortCategory
	SortVersion
	SortMilestone
	SortSummary
	SortStatus
	SortPriority
	SortAttachment
	SortSharedFile
	SortCreated
	SortCreatedUser
	SortUpdated
	SortUpdatedUser
	SortAssignee
	SortStartDate
	SortDueDate
	SortEstimatedHours
	SortActualHours
	SortChildIssue
)

func (k Sort) String() string {
	switch k {
	case SortIssueType:
		return "issueType"
	case SortCategory:
		return "category"
	case SortVersion:
		return "version"
	case SortMilestone:
		return "milestone"
	case SortSummary:
		return "summary"
	case SortStatus:
		return "status"
	case SortPriority:
		return "priority"
	case SortAttachment:
		return "attachment"
	case SortSharedFile:
		return "sharedFile"
	case SortCreated:
		return "created"
	case SortCreatedUser:
		return "createdUser"
	case SortUpdated:
		return "updated"
	case SortUpdatedUser:
		return "updatedUser"
	case SortAssignee:
		return "assignee"
	case SortStartDate:
		return "startDate"
	case SortDueDate:
		return "dueDate"
	case SortEstimatedHours:
		return "estimatedHours"
	case SortActualHours:
		return "actualHours"
	case SortChildIssue:
		return "childIssue"
	default:
		return ""
	}
}

// Issue : -
type Issue struct {
	ID             *int           `json:"id,omitempty"`
	ProjectID      *int           `json:"projectId,omitempty"`
	IssueKey       *string        `json:"issueKey,omitempty"`
	KeyID          *int           `json:"keyId,omitempty"`
	IssueType      *IssueType     `json:"issueType,omitempty"`
	Summary        *string        `json:"summary,omitempty"`
	Description    *string        `json:"description,omitempty"`
	Resolutions    *string        `json:"resolutions,omitempty"`
	Priority       *Priority      `json:"priority,omitempty"`
	Status         *Status        `json:"status,omitempty"`
	Assignee       *User          `json:"assignee,omitempty"`
	Category       []int          `json:"category,omitempty"`
	Versions       []int          `json:"versions,omitempty"`
	Milestone      []*Milestone   `json:"milestone,omitempty"`
	StartDate      *string        `json:"startDate,omitempty"`
	DueDate        *string        `json:"dueDate,omitempty"`
	EstimatedHours *int           `json:"estimatedHours,omitempty"`
	ActualHours    *int           `json:"actualHours,omitempty"`
	ParentIssueID  *int           `json:"parentIssueId,omitempty"`
	CreatedUser    *User          `json:"createdUser,omitempty"`
	Created        *Timestamp     `json:"created,omitempty"`
	UpdatedUser    *User          `json:"updatedUser,omitempty"`
	Updated        *Timestamp     `json:"updated,omitempty"`
	CustomFields   []*CustomField `json:"customFields,omitempty"`
	Attachments    []*Attachment  `json:"attachments,omitempty"`
	SharedFiles    []*SharedFile  `json:"sharedFiles,omitempty"`
	Stars          []*Star        `json:"stars,omitempty"`
}

// IssueType : issue type
type IssueType struct {
	ID           *int    `json:"id,omitempty"`
	ProjectID    *int    `json:"projectId,omitempty"`
	Name         *string `json:"name,omitempty"`
	Color        *string `json:"color,omitempty"`
	DisplayOrder *int    `json:"displayOrder,omitempty"`
}

// Milestone : -
type Milestone struct {
	ID             *int         `json:"id,omitempty"`
	ProjectID      *int         `json:"projectId,omitempty"`
	Name           *string      `json:"name,omitempty"`
	Description    *string      `json:"description,omitempty"`
	StartDate      *interface{} `json:"startDate,omitempty"`
	ReleaseDueDate *interface{} `json:"releaseDueDate,omitempty"`
	Archived       *bool        `json:"archived,omitempty"`
}

// GetIssues returns the list of issues
func (c *Client) GetIssues(opts *GetIssuesOptions) ([]*Issue, error) {
	return c.GetIssuesContext(context.Background(), opts)
}

// GetIssuesContext returns the list of issues with context
func (c *Client) GetIssuesContext(ctx context.Context, opts *GetIssuesOptions) ([]*Issue, error) {
	u, err := c.AddOptions("/api/v2/issues", opts)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	issues := []*Issue{}
	if err := c.Do(ctx, req, &issues); err != nil {
		return nil, err
	}
	return issues, nil
}

// Issues : list of issue
type Issues []*struct {
	Issue *Issue `json:"issue"`
}

// GetUserMySelfRecentrlyViewedIssues returns the list of issues a user view recently
// This api returns a json below:
// [
// 		{
//         "issue":{
// 				"id":1111111,
// 				...
// 			}
// 		},
// 		{
//         "issue":{
// 				"id":2222222,
// 				...
// 			}
// 		}
// 		...
// ]
func (c *Client) GetUserMySelfRecentrlyViewedIssues(opts *GetUserMySelfRecentrlyViewedIssuesOptions) (Issues, error) {
	return c.GetUserMySelfRecentrlyViewedIssuesContext(context.Background(), opts)
}

// GetUserMySelfRecentrlyViewedIssuesContext returns the list of issues a user view recently with context
func (c *Client) GetUserMySelfRecentrlyViewedIssuesContext(ctx context.Context, opts *GetUserMySelfRecentrlyViewedIssuesOptions) (Issues, error) {
	u := "/api/v2/users/myself/recentlyViewedIssues"

	u, err := c.AddOptions(u, opts)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var issues Issues
	if err := c.Do(ctx, req, &issues); err != nil {
		return nil, err
	}

	return issues, nil
}

// GetIssuesOptions specifies optional parameters to the GetIssues method.
type GetIssuesOptions struct {
	ProjectIDs     []int   `url:"projectId[],omitempty"`
	IssueTypeIDs   []int   `url:"issueTypeId[],omitempty"`
	CategoryIDs    []int   `url:"categoryId[],omitempty"`
	VersionIDs     []int   `url:"versionId[],omitempty"`
	MilestoneIDs   []int   `url:"milestoneId[],omitempty"`
	StatusIDs      []int   `url:"statusId[],omitempty"`
	PriorityIDs    []int   `url:"priorityId[],omitempty"`
	AssigneeIDs    []int   `url:"assigneeId[],omitempty"`
	CreatedUserIDs []int   `url:"createdUserId[],omitempty"`
	ResolutionIDs  []int   `url:"resolutionId[],omitempty"`
	ParentChild    *int    `url:"parentChild,omitempty"`
	Attachment     *bool   `url:"attachment,omitempty"`
	SharedFile     *bool   `url:"sharedFile,omitempty"`
	Sort           Sort    `url:"sort,omitempty"`
	Order          Order   `url:"order,omitempty"`
	Offset         *int    `url:"offset,omitempty"`
	Count          *int    `url:"count,omitempty"`
	CreatedSince   *string `url:"createdSince,omitempty"`
	CreatedUntil   *string `url:"createdUntil,omitempty"`
	UpdatedSince   *string `url:"updatedSince,omitempty"`
	UpdatedUntil   *string `url:"updatedUntil,omitempty"`
	StartDateSince *string `url:"startDateSince,omitempty"`
	StartDateUntil *string `url:"startDateUntil,omitempty"`
	DueDateSince   *string `url:"dueDateSince,omitempty"`
	DueDateUntil   *string `url:"dueDateUntil,omitempty"`
	IDs            []int   `url:"id[],omitempty"`
	ParentIssueIDs []int   `url:"parentIssueId[],omitempty"`
	Keyword        *string `url:"keyword,omitempty"`
}

// GetUserMySelfRecentrlyViewedIssuesOptions specifies optional parameters to the GetUserMySelfRecentrlyViewedIssues method.
type GetUserMySelfRecentrlyViewedIssuesOptions struct {
	Order  Order `url:"order,omitempty"`
	Offset *int  `url:"offset,omitempty"`
	Count  *int  `url:"count,omitempty"`
}
