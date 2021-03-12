package backlog

import (
	"context"
	"fmt"
	"io"
	"net/url"
)

// Sort : sort
type Sort string

// IssueType is used in sort type
const (
	SortIssueType = Sort(rune(iota))
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
	ID             *int                `json:"id,omitempty"`
	ProjectID      *int                `json:"projectId,omitempty"`
	IssueKey       *string             `json:"issueKey,omitempty"`
	KeyID          *int                `json:"keyId,omitempty"`
	IssueType      *IssueType          `json:"issueType,omitempty"`
	Summary        *string             `json:"summary,omitempty"`
	Description    *string             `json:"description,omitempty"`
	Resolution     *Resolution         `json:"resolution,omitempty"`
	Priority       *Priority           `json:"priority,omitempty"`
	Status         *Status             `json:"status,omitempty"`
	Assignee       *User               `json:"assignee,omitempty"`
	Category       []*Category         `json:"category,omitempty"`
	Versions       []*Version          `json:"versions,omitempty"`
	Milestone      []*Milestone        `json:"milestone,omitempty"`
	StartDate      *string             `json:"startDate,omitempty"`
	DueDate        *string             `json:"dueDate,omitempty"`
	EstimatedHours *float64            `json:"estimatedHours,omitempty"`
	ActualHours    *float64            `json:"actualHours,omitempty"`
	ParentIssueID  *int                `json:"parentIssueId,omitempty"`
	CreatedUser    *User               `json:"createdUser,omitempty"`
	Created        *Timestamp          `json:"created,omitempty"`
	UpdatedUser    *User               `json:"updatedUser,omitempty"`
	Updated        *Timestamp          `json:"updated,omitempty"`
	CustomFields   []*IssueCustomField `json:"customFields,omitempty"`
	Attachments    []*Attachment       `json:"attachments,omitempty"`
	SharedFiles    []*SharedFile       `json:"sharedFiles,omitempty"`
	Stars          []*Star             `json:"stars,omitempty"`
}

// Milestone : milestone
type Milestone struct {
	ID             *int    `json:"id,omitempty"`
	ProjectID      *int    `json:"projectId,omitempty"`
	Name           *string `json:"name,omitempty"`
	Description    *string `json:"description,omitempty"`
	StartDate      *string `json:"startDate,omitempty"`
	ReleaseDueDate *string `json:"releaseDueDate,omitempty"`
	Archived       *bool   `json:"archived,omitempty"`
}

// IssueComment : issue comment
type IssueComment struct {
	ID            *int            `json:"id,omitempty"`
	Content       *string         `json:"content,omitempty"`
	ChangeLog     []*ChangeLog    `json:"changeLog,omitempty"`
	CreatedUser   *User           `json:"createdUser,omitempty"`
	Created       *Timestamp      `json:"created,omitempty"`
	Updated       *Timestamp      `json:"updated,omitempty"`
	Stars         []*Star         `json:"stars,omitempty"`
	Notifications []*Notification `json:"notifications,omitempty"`
}

// ChangeLog : change log
type ChangeLog struct {
	AttachmentInfo   *AttachmentInfo   `json:"attachmentInfo,omitempty"`
	AttributeInfo    *AttributeInfo    `json:"attributeInfo,omitempty"`
	Field            *string           `json:"field,omitempty"`
	NewValue         *string           `json:"newValue,omitempty"`
	NotificationInfo *NotificationInfo `json:"notificationInfo,omitempty"`
	OriginalValue    *string           `json:"originalValue,omitempty"`
}

// AttachmentInfo : attachment information
type AttachmentInfo struct {
	ID   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// AttributeInfo : attribute information
type AttributeInfo struct {
	ID     *int `json:"id,omitempty"`
	TypeID *int `json:"typeId,omitempty"`
}

// NotificationInfo : notification information
type NotificationInfo struct {
	Type *string `json:"type,omitempty"`
}

// IssueCustomField : custom field in issue
type IssueCustomField struct {
	ID          *int        `json:"id,omitempty"`
	FieldTypeID *int        `json:"fieldTypeId,omitempty"`
	Name        *string     `json:"name,omitempty"`
	Value       interface{} `json:"value,omitempty"`
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

// GetIssueCount returns the count of issues
func (c *Client) GetIssueCount(opts *GetIssuesCountOptions) (int, error) {
	return c.GetIssueCountContext(context.Background(), opts)
}

// GetIssueCountContext returns the count of issues with context
func (c *Client) GetIssueCountContext(ctx context.Context, opts *GetIssuesCountOptions) (int, error) {
	u := "/api/v2/issues/count"

	u, err := c.AddOptions(u, opts)
	if err != nil {
		return 0, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return 0, err
	}

	r := new(p)
	if err := c.Do(ctx, req, &r); err != nil {
		return 0, err
	}

	return r.Count, nil
}

func createQueryStringsFromIssueCustomFileds(icf []*IssueCustomField) string {
	if len(icf) == 0 {
		return ""
	}
	q := url.Values{}
	for _, cf := range icf {
		var vals []interface{}
		if items, ok := cf.Value.([]*Item); ok {
			if len(items) == 0 {
				vals = append(vals, "")
			}
			for _, item := range items {
				vals = append(vals, *item.ID)
			}
		} else {
			vals = append(vals, cf.Value)
		}

		for _, val := range vals {
			f := fmt.Sprintf("customField_%v", *cf.ID)
			s := fmt.Sprint(val)
			q.Add(f, s)
		}
	}
	return q.Encode()
}

// CreateIssue creates a issue
func (c *Client) CreateIssue(input *CreateIssueInput) (*Issue, error) {
	return c.CreateIssueContext(context.Background(), input)
}

// CreateIssueContext creates a issue with context
func (c *Client) CreateIssueContext(ctx context.Context, input *CreateIssueInput) (*Issue, error) {
	u := "/api/v2/issues"

	if q := createQueryStringsFromIssueCustomFileds(input.CustomFields); q != "" {
		u += "?" + q
	}

	req, err := c.NewRequest("POST", u, input)
	if err != nil {
		return nil, err
	}

	issue := new(Issue)
	if err := c.Do(ctx, req, &issue); err != nil {
		return nil, err
	}
	return issue, nil
}

// GetIssue gets a issue
func (c *Client) GetIssue(issueIDOrKey string) (*Issue, error) {
	return c.GetIssueContext(context.Background(), issueIDOrKey)
}

// GetIssueContext gets a issue with context
func (c *Client) GetIssueContext(ctx context.Context, issueIDOrKey string) (*Issue, error) {
	u := fmt.Sprintf("/api/v2/issues/%v", issueIDOrKey)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	issue := new(Issue)
	if err := c.Do(ctx, req, &issue); err != nil {
		return nil, err
	}
	return issue, nil
}

// UpdateIssue updates a issue
func (c *Client) UpdateIssue(issueIDOrKey string, input *UpdateIssueInput) (*Issue, error) {
	return c.UpdateIssueContext(context.Background(), issueIDOrKey, input)
}

// UpdateIssueContext updates a issue with context
func (c *Client) UpdateIssueContext(ctx context.Context, issueIDOrKey string, input *UpdateIssueInput) (*Issue, error) {
	u := fmt.Sprintf("/api/v2/issues/%v", issueIDOrKey)

	if q := createQueryStringsFromIssueCustomFileds(input.CustomFields); q != "" {
		u += "?" + q
	}

	req, err := c.NewRequest("PATCH", u, input)
	if err != nil {
		return nil, err
	}

	issue := new(Issue)
	if err := c.Do(ctx, req, &issue); err != nil {
		return nil, err
	}
	return issue, nil
}

// GetIssueComments get list of the issue comments
func (c *Client) GetIssueComments(issueIDOrKey string, opts *GetIssueCommentsOptions) ([]*IssueComment, error) {
	return c.GetIssueCommentsContext(context.Background(), issueIDOrKey, opts)
}

// GetIssueCommentsContext gets list of the issue comments with context
func (c *Client) GetIssueCommentsContext(ctx context.Context, issueIDOrKey string, opts *GetIssueCommentsOptions) ([]*IssueComment, error) {
	u := fmt.Sprintf("/api/v2/issues/%v/comments", issueIDOrKey)

	req, err := c.NewRequest("GET", u, opts)
	if err != nil {
		return nil, err
	}

	issueComment := []*IssueComment{}
	if err := c.Do(ctx, req, &issueComment); err != nil {
		return nil, err
	}
	return issueComment, nil
}

// CreateIssueComment creates a issue comments
func (c *Client) CreateIssueComment(issueIDOrKey string, input *CreateIssueCommentInput) (*IssueComment, error) {
	return c.CreateIssueCommentContext(context.Background(), issueIDOrKey, input)
}

// CreateIssueCommentContext creates a issue comments with context
func (c *Client) CreateIssueCommentContext(ctx context.Context, issueIDOrKey string, input *CreateIssueCommentInput) (*IssueComment, error) {
	u := fmt.Sprintf("/api/v2/issues/%v/comments", issueIDOrKey)

	req, err := c.NewRequest("POST", u, input)
	if err != nil {
		return nil, err
	}

	issueComment := new(IssueComment)
	if err := c.Do(ctx, req, &issueComment); err != nil {
		return nil, err
	}
	return issueComment, nil
}

// GetIssueCommentsCount gets count of issue comments
func (c *Client) GetIssueCommentsCount(issueIDOrKey string) (int, error) {
	return c.GetIssueCommentsCountContext(context.Background(), issueIDOrKey)
}

// GetIssueCommentsCountContext gets count of issue comments with context
func (c *Client) GetIssueCommentsCountContext(ctx context.Context, issueIDOrKey string) (int, error) {
	u := fmt.Sprintf("/api/v2/issues/%v/comments/count", issueIDOrKey)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return 0, err
	}

	r := new(p)
	if err := c.Do(ctx, req, &r); err != nil {
		return 0, err
	}
	return r.Count, nil
}

// GetIssueComment gets a issue comment
func (c *Client) GetIssueComment(issueIDOrKey string, commentID int) (*IssueComment, error) {
	return c.GetIssueCommentContext(context.Background(), issueIDOrKey, commentID)
}

// GetIssueCommentContext gets a issue comment with context
func (c *Client) GetIssueCommentContext(ctx context.Context, issueIDOrKey string, commentID int) (*IssueComment, error) {
	u := fmt.Sprintf("/api/v2/issues/%v/comments/%v", issueIDOrKey, commentID)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	issueComment := new(IssueComment)
	if err := c.Do(ctx, req, &issueComment); err != nil {
		return nil, err
	}
	return issueComment, nil
}

// DeleteIssueComment deletes a issue comment
func (c *Client) DeleteIssueComment(issueIDOrKey string, commentID int) (*IssueComment, error) {
	return c.DeleteIssueCommentContext(context.Background(), issueIDOrKey, commentID)
}

// DeleteIssueCommentContext deletes a issue comment with context
func (c *Client) DeleteIssueCommentContext(ctx context.Context, issueIDOrKey string, commentID int) (*IssueComment, error) {
	u := fmt.Sprintf("/api/v2/issues/%v/comments/%v", issueIDOrKey, commentID)

	req, err := c.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	issueComment := new(IssueComment)
	if err := c.Do(ctx, req, &issueComment); err != nil {
		return nil, err
	}
	return issueComment, nil
}

// UpdateIssueComment updates a issue comment
func (c *Client) UpdateIssueComment(issueIDOrKey string, commentID int, input *UpdateIssueCommentInput) (*IssueComment, error) {
	return c.UpdateIssueCommentContext(context.Background(), issueIDOrKey, commentID, input)
}

// UpdateIssueCommentContext updates a issue comment with context
func (c *Client) UpdateIssueCommentContext(ctx context.Context, issueIDOrKey string, commentID int, input *UpdateIssueCommentInput) (*IssueComment, error) {
	u := fmt.Sprintf("/api/v2/issues/%v/comments/%v", issueIDOrKey, commentID)

	req, err := c.NewRequest("PATCH", u, input)
	if err != nil {
		return nil, err
	}

	issueComment := new(IssueComment)
	if err := c.Do(ctx, req, &issueComment); err != nil {
		return nil, err
	}
	return issueComment, nil
}

// GetIssueCommentsNotifications gets notifications in issue comments
func (c *Client) GetIssueCommentsNotifications(issueIDOrKey string, commentID int) ([]*Notification, error) {
	return c.GetIssueCommentsNotificationsContext(context.Background(), issueIDOrKey, commentID)
}

// GetIssueCommentsNotificationsContext gets a issue comment with context
func (c *Client) GetIssueCommentsNotificationsContext(ctx context.Context, issueIDOrKey string, commentID int) ([]*Notification, error) {
	u := fmt.Sprintf("/api/v2/issues/%v/comments/%v/notifications", issueIDOrKey, commentID)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	notifications := []*Notification{}
	if err := c.Do(ctx, req, &notifications); err != nil {
		return nil, err
	}
	return notifications, nil
}

// CreateIssueCommentsNotification creates a notification
func (c *Client) CreateIssueCommentsNotification(issueIDOrKey string, commentID int, input *CreateIssueCommentsNotificationInput) (*IssueComment, error) {
	return c.CreateIssueCommentsNotificationContext(context.Background(), issueIDOrKey, commentID, input)
}

// CreateIssueCommentsNotificationContext creates a notification with context
func (c *Client) CreateIssueCommentsNotificationContext(ctx context.Context, issueIDOrKey string, commentID int, input *CreateIssueCommentsNotificationInput) (*IssueComment, error) {
	u := fmt.Sprintf("/api/v2/issues/%v/comments/%v/notifications", issueIDOrKey, commentID)

	req, err := c.NewRequest("POST", u, input)
	if err != nil {
		return nil, err
	}

	issueComment := new(IssueComment)
	if err := c.Do(ctx, req, &issueComment); err != nil {
		return nil, err
	}
	return issueComment, nil
}

// GetIssueAttachments gets issue attachments
func (c *Client) GetIssueAttachments(issueIDOrKey string) ([]*Attachment, error) {
	return c.GetIssueAttachmentsContext(context.Background(), issueIDOrKey)
}

// GetIssueAttachmentsContext gets issue attachments with context
func (c *Client) GetIssueAttachmentsContext(ctx context.Context, issueIDOrKey string) ([]*Attachment, error) {
	u := fmt.Sprintf("/api/v2/issues/%v/attachments", issueIDOrKey)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	attachments := []*Attachment{}
	if err := c.Do(ctx, req, &attachments); err != nil {
		return nil, err
	}
	return attachments, nil
}

// GetIssueAttachment downloads an issue attachment
func (c *Client) GetIssueAttachment(issueIDOrKey string, attachmentID int, writer io.Writer) error {
	return c.GetIssueAttachmentContext(context.Background(), issueIDOrKey, attachmentID, writer)
}

// GetIssueAttachmentContext downloads an issue attachment with context
func (c *Client) GetIssueAttachmentContext(ctx context.Context, issueIDOrKey string, attachmentID int, writer io.Writer) error {
	u := fmt.Sprintf("/api/v2/issues/%v/attachments/%v", issueIDOrKey, attachmentID)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, writer); err != nil {
		return err
	}
	return nil
}

// DeleteIssueAttachment deletes an issue attachment
func (c *Client) DeleteIssueAttachment(issueIDOrKey string, attachmentID int) (*Attachment, error) {
	return c.DeleteIssueAttachmentContext(context.Background(), issueIDOrKey, attachmentID)
}

// DeleteIssueAttachmentContext deletes an issue attachments with context
func (c *Client) DeleteIssueAttachmentContext(ctx context.Context, issueIDOrKey string, attachmentID int) (*Attachment, error) {
	u := fmt.Sprintf("/api/v2/issues/%v/attachments/%v", issueIDOrKey, attachmentID)

	req, err := c.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	attachment := new(Attachment)
	if err := c.Do(ctx, req, &attachment); err != nil {
		return nil, err
	}
	return attachment, nil
}

// GetIssueParticipants gets participants of a issue
func (c *Client) GetIssueParticipants(issueIDOrKey string) ([]*User, error) {
	return c.GetIssueParticipantsContext(context.Background(), issueIDOrKey)
}

// GetIssueParticipantsContext gets participants of a issue with context
func (c *Client) GetIssueParticipantsContext(ctx context.Context, issueIDOrKey string) ([]*User, error) {
	u := fmt.Sprintf("/api/v2/issues/%v/participants", issueIDOrKey)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	users := []*User{}
	if err := c.Do(ctx, req, &users); err != nil {
		return nil, err
	}
	return users, nil
}

// GetIssueSharedFiles gets shared files of a issue
func (c *Client) GetIssueSharedFiles(issueIDOrKey string) ([]*SharedFile, error) {
	return c.GetIssueSharedFilesContext(context.Background(), issueIDOrKey)
}

// GetIssueSharedFilesContext gets shared files of a issue with context
func (c *Client) GetIssueSharedFilesContext(ctx context.Context, issueIDOrKey string) ([]*SharedFile, error) {
	u := fmt.Sprintf("/api/v2/issues/%v/sharedFiles", issueIDOrKey)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	sharedFiles := []*SharedFile{}
	if err := c.Do(ctx, req, &sharedFiles); err != nil {
		return nil, err
	}
	return sharedFiles, nil
}

// CreateIssueSharedFiles link a shared files to a issue
func (c *Client) CreateIssueSharedFiles(issueIDOrKey string, input *CreateIssueSharedFilesInput) ([]*SharedFile, error) {
	return c.CreateIssueSharedFilesContext(context.Background(), issueIDOrKey, input)
}

// CreateIssueSharedFilesContext link a shared files to a issue with context
func (c *Client) CreateIssueSharedFilesContext(ctx context.Context, issueIDOrKey string, input *CreateIssueSharedFilesInput) ([]*SharedFile, error) {
	u := fmt.Sprintf("/api/v2/issues/%v/sharedFiles", issueIDOrKey)

	req, err := c.NewRequest("POST", u, input)
	if err != nil {
		return nil, err
	}

	sharedFiles := []*SharedFile{}
	if err := c.Do(ctx, req, &sharedFiles); err != nil {
		return nil, err
	}
	return sharedFiles, nil
}

// DeleteIssueSharedFile link a shared files to a issue
func (c *Client) DeleteIssueSharedFile(issueIDOrKey string, sharedFileID int) (*SharedFile, error) {
	return c.DeleteIssueSharedFileContext(context.Background(), issueIDOrKey, sharedFileID)
}

// DeleteIssueSharedFileContext link a shared files to a issue with context
func (c *Client) DeleteIssueSharedFileContext(ctx context.Context, issueIDOrKey string, sharedFileID int) (*SharedFile, error) {
	u := fmt.Sprintf("/api/v2/issues/%v/sharedFiles/%v", issueIDOrKey, sharedFileID)

	req, err := c.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	sharedFile := new(SharedFile)
	if err := c.Do(ctx, req, &sharedFile); err != nil {
		return nil, err
	}
	return sharedFile, nil
}

// GetIssuesOptions specifies parameters to the GetIssues method.
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

// GetUserMySelfRecentrlyViewedIssuesOptions specifies parameters to the GetUserMySelfRecentrlyViewedIssues method.
type GetUserMySelfRecentrlyViewedIssuesOptions struct {
	Order  Order `url:"order,omitempty"`
	Offset *int  `url:"offset,omitempty"`
	Count  *int  `url:"count,omitempty"`
}

// GetIssuesCountOptions specifies parameters to the GetIssueCount method.
type GetIssuesCountOptions struct {
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

// CreateIssueInput specifies parameters to the CreateIssue method.
type CreateIssueInput struct {
	ProjectID       *int                `json:"projectId"`
	Summary         *string             `json:"summary"`
	ParentIssueID   *int                `json:"parentIssueId,omitempty"`
	Description     *string             `json:"description,omitempty"`
	StartDate       *string             `json:"startDate,omitempty"`
	DueDate         *string             `json:"dueDate,omitempty"`
	EstimatedHours  *float64            `json:"estimatedHours,omitempty"`
	ActualHours     *float64            `json:"actualHours,omitempty"`
	IssueTypeID     *int                `json:"issueTypeId"`
	CategoryIDs     []int               `json:"categoryId,omitempty"`
	VersionIDs      []int               `json:"versionId,omitempty"`
	MilestoneIDs    []int               `json:"milestoneId,omitempty"`
	PriorityID      *int                `json:"priorityId"`
	AssigneeID      *int                `json:"assigneeId,omitempty"`
	NotifiedUserIDs []int               `json:"notifiedUserId,omitempty"`
	AttachmentIDs   []int               `json:"attachmentId,omitempty"`
	CustomFields    []*IssueCustomField `json:"-"`
}

// UpdateIssueInput specifies parameters to the UpdateIssue method.
//
// The following values ​​are updated as integer or an empty string:
// - resolutionId
// - estimatedHours
// - actualHours
// - assigneeId
//
// ex. update `estimatedHours` to 10
// - UpdateIssue("EX-1", &UpdateIssueInput{ EstimatedHours: Int(10) })
//
// ex. update `estimatedHours` to an empty string
// - UpdateIssue("EX-1", &UpdateIssueInput{ EstimatedHours: String("") })
type UpdateIssueInput struct {
	Summary         *string             `json:"summary,omitempty"`
	ParentIssueID   *int                `json:"parentIssueId,omitempty"`
	Description     *string             `json:"description,omitempty"`
	StatusID        *int                `json:"statusId,omitempty"`
	ResolutionID    interface{}         `json:"resolutionId,omitempty"`
	StartDate       *string             `json:"startDate,omitempty"`
	DueDate         *string             `json:"dueDate,omitempty"`
	EstimatedHours  interface{}         `json:"estimatedHours,omitempty"`
	ActualHours     interface{}         `json:"actualHours,omitempty"`
	IssueTypeID     *int                `json:"issueTypeId,omitempty"`
	CategoryIDs     []int               `json:"categoryId,omitempty"`
	VersionIDs      []int               `json:"versionId,omitempty"`
	MilestoneIDs    []int               `json:"milestoneId,omitempty"`
	PriorityID      *int                `json:"priorityId,omitempty"`
	AssigneeID      interface{}         `json:"assigneeId,omitempty"`
	NotifiedUserIDs []int               `json:"notifiedUserId,omitempty"`
	AttachmentIDs   []int               `json:"attachmentId,omitempty"`
	Comment         *string             `json:"comment,omitempty"`
	CustomFields    []*IssueCustomField `json:"-"`
}

// GetIssueCommentsOptions specifies parameters to the GetIssueComments method.
type GetIssueCommentsOptions struct {
	MinID *int  `json:"minId,omitempty"`
	MaxID *int  `json:"maxId,omitempty"`
	Count *int  `json:"count,omitempty"`
	Order Order `json:"order,omitempty"`
}

// CreateIssueCommentInput specifies parameters to the CreateIssueComment method.
type CreateIssueCommentInput struct {
	Content         *string `json:"content"`
	NotifiedUserIDs []int   `json:"notifiedUserId,omitempty"`
	AttachmentIDs   []int   `json:"attachmentId,omitempty"`
}

// UpdateIssueCommentInput specifies parameters to the UpdateIssueComment method.
type UpdateIssueCommentInput struct {
	Content *string `json:"content,omitempty"`
}

// CreateIssueCommentsNotificationInput specifies parameters to the CreateIssueCommentsNotification method.
type CreateIssueCommentsNotificationInput struct {
	NotifiedUserIDs []int `json:"notifiedUserId,omitempty"`
}

// CreateIssueSharedFilesInput specifies parameters to the CreateIssueSharedFiles method.
type CreateIssueSharedFilesInput struct {
	FileIDs []int `json:"fileId,omitempty"`
}
