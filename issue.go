package backlog

import (
	"context"
	"net/url"
	"strconv"
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
func (api *Client) GetIssues(input *GetIssuesInput) ([]*Issue, error) {
	return api.GetIssuesContext(context.Background(), input)
}

// GetIssuesContext returns the list of issues with context
func (api *Client) GetIssuesContext(ctx context.Context, input *GetIssuesInput) ([]*Issue, error) {

	values := url.Values{}

	if len(input.ProjectIDs) > 0 {
		for _, i := range input.ProjectIDs {
			values.Add("projectId[]", strconv.Itoa(i))
		}
	}

	if len(input.IssueTypeIDs) > 0 {
		for _, i := range input.IssueTypeIDs {
			values.Add("issueTypeId[]", strconv.Itoa(i))
		}
	}

	if len(input.CategoryIDs) > 0 {
		for _, i := range input.CategoryIDs {
			values.Add("categoryId[]", strconv.Itoa(i))
		}
	}

	if len(input.VersionIDs) > 0 {
		for _, i := range input.VersionIDs {
			values.Add("versionId[]", strconv.Itoa(i))
		}
	}

	if len(input.MilestoneIDs) > 0 {
		for _, i := range input.MilestoneIDs {
			values.Add("milestoneId[]", strconv.Itoa(i))
		}
	}

	if len(input.StatusIDs) > 0 {
		for _, i := range input.StatusIDs {
			values.Add("statusId[]", strconv.Itoa(i))
		}
	}

	if len(input.PriorityIDs) > 0 {
		for _, i := range input.PriorityIDs {
			values.Add("priorityId[]", strconv.Itoa(i))
		}
	}

	if len(input.AssigneeIDs) > 0 {
		for _, i := range input.AssigneeIDs {
			values.Add("assigneeId[]", strconv.Itoa(i))
		}
	}

	if len(input.CreatedUserIDs) > 0 {
		for _, i := range input.CreatedUserIDs {
			values.Add("createdUserId[]", strconv.Itoa(i))
		}
	}

	if len(input.ResolutionIDs) > 0 {
		for _, i := range input.ResolutionIDs {
			values.Add("resolutionId[]", strconv.Itoa(i))
		}
	}

	if input.ParentChild != nil && *input.ParentChild > 0 {
		values.Add("parentChild", strconv.Itoa(*input.ParentChild))
	}

	if input.Attachment != nil {
		values.Add("attachment", strconv.FormatBool(*input.Attachment))
	}

	if input.SharedFile != nil {
		values.Add("sharedFile", strconv.FormatBool(*input.SharedFile))
	}

	if input.Sort != "" {
		values.Add("sort", input.Sort.String())
	}

	if input.Order == "" {
		values.Add("order", OrderDesc.String())
	} else {
		values.Add("order", input.Order.String())
	}

	if input.Offset != nil {
		values.Add("offset", strconv.Itoa(*input.Offset))
	}

	if input.Count != nil {
		values.Add("count", strconv.Itoa(*input.Count))
	} else {
		values.Add("count", "20")
	}

	if input.CreatedSince != nil {
		values.Add("createdSince", *input.CreatedSince)
	}

	if input.CreatedUntil != nil {
		values.Add("createdUntil", *input.CreatedUntil)
	}

	if input.UpdatedSince != nil {
		values.Add("updatedSince", *input.UpdatedSince)
	}

	if input.UpdatedUntil != nil {
		values.Add("updatedUntil", *input.UpdatedUntil)
	}

	if input.StartDateSince != nil {
		values.Add("startDateSince", *input.StartDateSince)
	}

	if input.StartDateUntil != nil {
		values.Add("startDateUntil", *input.StartDateUntil)
	}

	if input.DueDateSince != nil {
		values.Add("dueDateSince", *input.DueDateSince)
	}

	if input.DueDateUntil != nil {
		values.Add("dueDateUntil", *input.DueDateUntil)
	}

	if len(input.IDs) > 0 {
		for _, i := range input.IDs {
			values.Add("id[]", strconv.Itoa(i))
		}
	}

	if len(input.ParentIssueIDs) > 0 {
		for _, i := range input.ParentIssueIDs {
			values.Add("parentIssueId[]", strconv.Itoa(i))
		}
	}

	if input.Keyword != nil {
		values.Add("keyword", *input.Keyword)
	}

	r := []*Issue{}
	if err := api.getMethod(ctx, "/api/v2/issues", values, &r); err != nil {
		return nil, err
	}
	return r, nil
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
func (api *Client) GetUserMySelfRecentrlyViewedIssues(input *GetUserMySelfRecentrlyViewedIssuesInput) (Issues, error) {
	return api.GetUserMySelfRecentrlyViewedIssuesContext(context.Background(), input)
}

// GetUserMySelfRecentrlyViewedIssuesContext returns the list of issues a user view recently with context
func (api *Client) GetUserMySelfRecentrlyViewedIssuesContext(ctx context.Context, input *GetUserMySelfRecentrlyViewedIssuesInput) (Issues, error) {

	values := url.Values{}

	if input.Order.String() != "" {
		values.Add("order", input.Order.String())
	} else {
		values.Add("order", OrderDesc.String())
	}

	if input.Offset != nil {
		values.Add("offset", strconv.Itoa(*input.Offset))
	}

	if input.Count != nil {
		values.Add("count", strconv.Itoa(*input.Count))
	}

	var issues Issues
	if err := api.getMethod(ctx, "/api/v2/users/myself/recentlyViewedIssues", values, &issues); err != nil {
		return nil, err
	}

	return issues, nil
}

// GetIssuesInput contains all the parameters necessary (including the optional ones) for a GetIssues() request.
type GetIssuesInput struct {
	ProjectIDs     []int   `required:"false"`
	IssueTypeIDs   []int   `required:"false"`
	CategoryIDs    []int   `required:"false"`
	VersionIDs     []int   `required:"false"`
	MilestoneIDs   []int   `required:"false"`
	StatusIDs      []int   `required:"false"`
	PriorityIDs    []int   `required:"false"`
	AssigneeIDs    []int   `required:"false"`
	CreatedUserIDs []int   `required:"false"`
	ResolutionIDs  []int   `required:"false"`
	ParentChild    *int    `required:"false"`
	Attachment     *bool   `required:"false"`
	SharedFile     *bool   `required:"false"`
	Sort           Sort    `required:"false"`
	Order          Order   `required:"false"`
	Offset         *int    `required:"false"`
	Count          *int    `required:"false"`
	CreatedSince   *string `required:"false"`
	CreatedUntil   *string `required:"false"`
	UpdatedSince   *string `required:"false"`
	UpdatedUntil   *string `required:"false"`
	StartDateSince *string `required:"false"`
	StartDateUntil *string `required:"false"`
	DueDateSince   *string `required:"false"`
	DueDateUntil   *string `required:"false"`
	IDs            []int   `required:"false"`
	ParentIssueIDs []int   `required:"false"`
	Keyword        *string `required:"false"`
}
