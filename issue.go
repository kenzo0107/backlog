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
	ID             int           `json:"id"`
	ProjectID      int           `json:"projectId"`
	IssueKey       string        `json:"issueKey"`
	KeyID          int           `json:"keyId"`
	IssueType      IssueType     `json:"issueType"`
	Summary        string        `json:"summary"`
	Description    string        `json:"description"`
	Resolutions    interface{}   `json:"resolutions"`
	Priority       Priority      `json:"priority"`
	Status         Status        `json:"status"`
	Assignee       User          `json:"assignee"`
	Category       []interface{} `json:"category"`
	Versions       []interface{} `json:"versions"`
	Milestone      []Milestone   `json:"milestone"`
	StartDate      interface{}   `json:"startDate"`
	DueDate        interface{}   `json:"dueDate"`
	EstimatedHours interface{}   `json:"estimatedHours"`
	ActualHours    interface{}   `json:"actualHours"`
	ParentIssueID  interface{}   `json:"parentIssueId"`
	CreatedUser    User          `json:"createdUser"`
	Created        JSONTime      `json:"created"`
	UpdatedUser    User          `json:"updatedUser"`
	Updated        JSONTime      `json:"updated"`
	CustomFields   []interface{} `json:"customFields"`
	Attachments    []Attachment  `json:"attachments"`
	SharedFiles    []SharedFile  `json:"sharedFiles"`
	Stars          []Star        `json:"stars"`
}

// IssueType : issue type
type IssueType struct {
	ID           int    `json:"id"`
	ProjectID    int    `json:"projectId"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	DisplayOrder int    `json:"displayOrder"`
}

// Milestone : -
type Milestone struct {
	ID             int         `json:"id"`
	ProjectID      int         `json:"projectId"`
	Name           string      `json:"name"`
	Description    string      `json:"description"`
	StartDate      interface{} `json:"startDate"`
	ReleaseDueDate interface{} `json:"releaseDueDate"`
	Archived       bool        `json:"archived"`
}

// GetIssues returns the list of issues
func (api *Client) GetIssues(input *GetIssuesInput) ([]Issue, error) {
	return api.GetIssuesContext(context.Background(), input)
}

// GetIssuesContext returns the list of issues with context
func (api *Client) GetIssuesContext(ctx context.Context, input *GetIssuesInput) ([]Issue, error) {

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

	if input.ParentChild > 0 {
		values.Add("parentChild", strconv.Itoa(input.ParentChild))
	}

	values.Add("attachment", strconv.FormatBool(input.Attachment))

	values.Add("sharedFile", strconv.FormatBool(input.SharedFile))

	if input.Sort != "" {
		values.Add("sort", input.Sort.String())
	}

	if input.Order == "" {
		values.Add("order", OrderDesc.String())
	} else {
		values.Add("order", input.Order.String())
	}

	if input.Offset > 0 {
		values.Add("offset", strconv.Itoa(input.Offset))
	}

	if input.Count == 0 {
		values.Add("count", "20")
	} else {
		values.Add("count", strconv.Itoa(input.Count))
	}

	if input.CreatedSince != "" {
		values.Add("createdSince", input.CreatedSince)
	}

	if input.CreatedUntil != "" {
		values.Add("createdUntil", input.CreatedUntil)
	}

	if input.UpdatedSince != "" {
		values.Add("updatedSince", input.UpdatedSince)
	}

	if input.UpdatedUntil != "" {
		values.Add("updatedUntil", input.UpdatedUntil)
	}

	if input.StartDateSince != "" {
		values.Add("startDateSince", input.StartDateSince)
	}

	if input.StartDateUntil != "" {
		values.Add("startDateUntil", input.StartDateUntil)
	}

	if input.DueDateSince != "" {
		values.Add("dueDateSince", input.DueDateSince)
	}

	if input.DueDateUntil != "" {
		values.Add("dueDateUntil", input.DueDateUntil)
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

	if input.Keyword != "" {
		values.Add("keyword", input.Keyword)
	}

	r := []Issue{}
	if err := api.getMethod(ctx, "/api/v2/issues", values, &r); err != nil {
		return nil, err
	}
	return r, nil
}

// GetIssuesInput contains all the parameters necessary (including the optional ones) for a GetIssues() request.
type GetIssuesInput struct {
	ProjectIDs     []int
	IssueTypeIDs   []int
	CategoryIDs    []int
	VersionIDs     []int
	MilestoneIDs   []int
	StatusIDs      []int
	PriorityIDs    []int
	AssigneeIDs    []int
	CreatedUserIDs []int
	ResolutionIDs  []int
	ParentChild    int
	Attachment     bool
	SharedFile     bool
	Sort           Sort
	Order          Order
	Offset         int
	Count          int
	CreatedSince   string
	CreatedUntil   string
	UpdatedSince   string
	UpdatedUntil   string
	StartDateSince string
	StartDateUntil string
	DueDateSince   string
	DueDateUntil   string
	IDs            []int
	ParentIssueIDs []int
	Keyword        string
}
