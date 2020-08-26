package backlog

import (
	"context"
	"io"
	"net/http"
)

// API provides an interface to enable mocking the
// backlog client's API operation.
// This makes unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Elastic Compute Cloud.
//    func myFunc(svc backlog.API) string {
//        // Make svc.GetUserMySelf() request
//    }
//
//    func main() {
//        svc := backlog.New()
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockBacklogClient struct {
//        backlog.API
//    }
//
//    func (m *mockBacklogClient) GetUserMySelf() (*backlog.User, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := new(mockBacklogClient)
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations.
// It's suggested to use the pattern above for testing or using
// tooling to generate mocks to satisfy the interfaces.
type API interface {
	GetUserActivities(id int, opts *GetUserActivitiesOptions) ([]*Activity, error)
	GetUserActivitiesContext(ctx context.Context, id int, opts *GetUserActivitiesOptions) ([]*Activity, error)
	GetProjectActivities(projectIDOrKey interface{}, opts *GetProjectActivitiesOptions) ([]*Activity, error)
	GetProjectActivitiesContext(ctx context.Context, projectIDOrKey interface{}, opts *GetProjectActivitiesOptions) ([]*Activity, error)
	Debugf(format string, v ...interface{})
	Debugln(v ...interface{})
	Debug() bool
	NewRequest(method, urlStr string, body interface{}) (*http.Request, error)
	UploadMultipartFile(ctx context.Context, method, urlStr, fpath, field string, v interface{}) (err error)
	Do(ctx context.Context, req *http.Request, v interface{}) error
	AddOptions(s string, opts interface{}) (string, error)
	GetCategories(projectIDOrKey interface{}) ([]*Category, error)
	GetCategoriesContext(ctx context.Context, projectIDOrKey interface{}) ([]*Category, error)
	CreateCategory(projectIDOrKey interface{}, input *CreateCategoryInput) (*Category, error)
	CreateCategoryContext(ctx context.Context, projectIDOrKey interface{}, input *CreateCategoryInput) (*Category, error)
	UpdateCategory(projectIDOrKey interface{}, categoryID int, input *UpdateCategoryInput) (*Category, error)
	UpdateCategoryContext(ctx context.Context, projectIDOrKey interface{}, categoryID int, input *UpdateCategoryInput) (*Category, error)
	DeleteCategory(projectIDOrKey interface{}, categoryID int) (*Category, error)
	DeleteCategoryContext(ctx context.Context, projectIDOrKey interface{}, categoryID int) (*Category, error)
	GetCustomFields(projectIDOrKey interface{}) ([]*CustomField, error)
	GetCustomFieldsContext(ctx context.Context, projectIDOrKey interface{}) ([]*CustomField, error)
	UploadFile(fpath string) (*FileUploadResponse, error)
	UploadFileContext(ctx context.Context, fpath string) (*FileUploadResponse, error)
	GetIssues(opts *GetIssuesOptions) ([]*Issue, error)
	GetIssuesContext(ctx context.Context, opts *GetIssuesOptions) ([]*Issue, error)
	GetUserMySelfRecentrlyViewedIssues(opts *GetUserMySelfRecentrlyViewedIssuesOptions) (Issues, error)
	GetUserMySelfRecentrlyViewedIssuesContext(ctx context.Context, opts *GetUserMySelfRecentrlyViewedIssuesOptions) (Issues, error)
	GetIssueCount(opts *GetIssuesCountOptions) (int, error)
	GetIssueCountContext(ctx context.Context, opts *GetIssuesCountOptions) (int, error)
	CreateIssue(input *CreateIssueInput) (*Issue, error)
	CreateIssueContext(ctx context.Context, input *CreateIssueInput) (*Issue, error)
	GetIssue(issueIDOrKey string) (*Issue, error)
	GetIssueContext(ctx context.Context, issueIDOrKey string) (*Issue, error)
	UpdateIssue(issueIDOrKey string, input *UpdateIssueInput) (*Issue, error)
	UpdateIssueContext(ctx context.Context, issueIDOrKey string, input *UpdateIssueInput) (*Issue, error)
	GetIssueComments(issueIDOrKey string, opts *GetIssueCommentsOptions) ([]*IssueComment, error)
	GetIssueCommentsContext(ctx context.Context, issueIDOrKey string, opts *GetIssueCommentsOptions) ([]*IssueComment, error)
	CreateIssueComment(issueIDOrKey string, input *CreateIssueCommentInput) (*IssueComment, error)
	CreateIssueCommentContext(ctx context.Context, issueIDOrKey string, input *CreateIssueCommentInput) (*IssueComment, error)
	GetIssueCommentsCount(issueIDOrKey string) (int, error)
	GetIssueCommentsCountContext(ctx context.Context, issueIDOrKey string) (int, error)
	GetIssueComment(issueIDOrKey string, commentID int) (*IssueComment, error)
	GetIssueCommentContext(ctx context.Context, issueIDOrKey string, commentID int) (*IssueComment, error)
	DeleteIssueComment(issueIDOrKey string, commentID int) (*IssueComment, error)
	DeleteIssueCommentContext(ctx context.Context, issueIDOrKey string, commentID int) (*IssueComment, error)
	UpdateIssueComment(issueIDOrKey string, commentID int, input *UpdateIssueCommentInput) (*IssueComment, error)
	UpdateIssueCommentContext(ctx context.Context, issueIDOrKey string, commentID int, input *UpdateIssueCommentInput) (*IssueComment, error)
	GetIssueCommentsNotifications(issueIDOrKey string, commentID int) ([]*Notification, error)
	GetIssueCommentsNotificationsContext(ctx context.Context, issueIDOrKey string, commentID int) ([]*Notification, error)
	CreateIssueCommentsNotification(issueIDOrKey string, commentID int, input *CreateIssueCommentsNotificationInput) (*IssueComment, error)
	CreateIssueCommentsNotificationContext(ctx context.Context, issueIDOrKey string, commentID int, input *CreateIssueCommentsNotificationInput) (*IssueComment, error)
	GetIssueAttachments(issueIDOrKey string) ([]*Attachment, error)
	GetIssueAttachmentsContext(ctx context.Context, issueIDOrKey string) ([]*Attachment, error)
	GetIssueAttachment(issueIDOrKey string, attachmentID int, writer io.Writer) error
	GetIssueAttachmentContext(ctx context.Context, issueIDOrKey string, attachmentID int, writer io.Writer) error
	DeleteIssueAttachment(issueIDOrKey string, attachmentID int) (*Attachment, error)
	DeleteIssueAttachmentContext(ctx context.Context, issueIDOrKey string, attachmentID int) (*Attachment, error)
	GetIssueParticipants(issueIDOrKey string) ([]*User, error)
	GetIssueParticipantsContext(ctx context.Context, issueIDOrKey string) ([]*User, error)
	GetIssueSharedFiles(issueIDOrKey string) ([]*SharedFile, error)
	GetIssueSharedFilesContext(ctx context.Context, issueIDOrKey string) ([]*SharedFile, error)
	CreateIssueSharedFiles(issueIDOrKey string, input *CreateIssueSharedFilesInput) ([]*SharedFile, error)
	CreateIssueSharedFilesContext(ctx context.Context, issueIDOrKey string, input *CreateIssueSharedFilesInput) ([]*SharedFile, error)
	DeleteIssueSharedFile(issueIDOrKey string, sharedFileID int) (*SharedFile, error)
	DeleteIssueSharedFileContext(ctx context.Context, issueIDOrKey string, sharedFileID int) (*SharedFile, error)
	GetIssueTypes(projectIDOrKey interface{}) ([]*IssueType, error)
	GetIssueTypesContext(ctx context.Context, projectIDOrKey interface{}) ([]*IssueType, error)
	CreateIssueType(projectIDOrKey interface{}, input *CreateIssueTypeInput) (*IssueType, error)
	CreateIssueTypeContext(ctx context.Context, projectIDOrKey interface{}, input *CreateIssueTypeInput) (*IssueType, error)
	UpdateIssueType(projectIDOrKey interface{}, issueTypeID int, input *UpdateIssueTypeInput) (*IssueType, error)
	UpdateIssueTypeContext(ctx context.Context, projectIDOrKey interface{}, issueTypeID int, input *UpdateIssueTypeInput) (*IssueType, error)
	DeleteIssueType(projectIDOrKey interface{}, issueTypeID int, input *DeleteIssueTypeInput) (*IssueType, error)
	DeleteIssueTypeContext(ctx context.Context, projectIDOrKey interface{}, issueTypeID int, input *DeleteIssueTypeInput) (*IssueType, error)
	GetPriorities() ([]*Priority, error)
	GetPrioritiesContext(ctx context.Context) ([]*Priority, error)
	GetMyRecentlyViewedProjects(opts *GetMyRecentlyViewedProjectsOptions) ([]*RecentlyViewedProject, error)
	GetMyRecentlyViewedProjectsContext(ctx context.Context, opts *GetMyRecentlyViewedProjectsOptions) ([]*RecentlyViewedProject, error)
	GetProjects(opts *GetProjectsOptions) ([]*Project, error)
	GetProjectsContext(ctx context.Context, opts *GetProjectsOptions) ([]*Project, error)
	GetProject(projectIDOrKey interface{}) (*Project, error)
	GetProjectContext(ctx context.Context, projectIDOrKey interface{}) (*Project, error)
	GetStatuses(projectIDOrKey interface{}) ([]*Status, error)
	GetStatusesContext(ctx context.Context, projectIDOrKey interface{}) ([]*Status, error)
	CreateProject(input *CreateProjectInput) (*Project, error)
	CreateProjectContext(ctx context.Context, input *CreateProjectInput) (*Project, error)
	UpdateProject(id int, input *UpdateProjectInput) (*Project, error)
	UpdateProjectContext(ctx context.Context, id int, input *UpdateProjectInput) (*Project, error)
	DeleteProject(projectIDOrKey interface{}) (*Project, error)
	DeleteProjectContext(ctx context.Context, projectIDOrKey interface{}) (*Project, error)
	GetProjectIcon(projectIDOrKey interface{}, writer io.Writer) error
	GetProjectIconContext(ctx context.Context, projectIDOrKey interface{}, writer io.Writer) error
	AddProjectUser(projectIDOrKey interface{}, input *AddProjectUserInput) (*User, error)
	AddProjectUserContext(ctx context.Context, projectIDOrKey interface{}, input *AddProjectUserInput) (*User, error)
	GetProjectUsers(projectIDOrKey interface{}, opts *GetProjectUsersOptions) ([]*User, error)
	GetProjectUsersContext(ctx context.Context, projectIDOrKey interface{}, opts *GetProjectUsersOptions) ([]*User, error)
	DeleteProjectUser(projectIDOrKey interface{}, input *DeleteProjectUserInput) (*User, error)
	DeleteProjectUserContext(ctx context.Context, projectIDOrKey interface{}, input *DeleteProjectUserInput) (*User, error)
	AddProjectAdministrator(projectIDOrKey interface{}, input *AddProjectAdministratorInput) (*User, error)
	AddProjectAdministratorContext(ctx context.Context, projectIDOrKey interface{}, input *AddProjectAdministratorInput) (*User, error)
	GetProjectAdministrators(projectIDOrKey interface{}) ([]*User, error)
	GetProjectAdministratorsContext(ctx context.Context, projectIDOrKey interface{}) ([]*User, error)
	DeleteProjectAdministrator(projectIDOrKey interface{}, input *DeleteProjectAdministratorInput) (*User, error)
	DeleteProjectAdministratorContext(ctx context.Context, projectIDOrKey interface{}, input *DeleteProjectAdministratorInput) (*User, error)
	CreateStatus(projectIDOrKey interface{}, input *CreateStatusInput) (*Status, error)
	CreateStatusContext(ctx context.Context, projectIDOrKey interface{}, input *CreateStatusInput) (*Status, error)
	UpdateStatus(projectIDOrKey interface{}, statusID int, input *UpdateStatusInput) (*Status, error)
	UpdateStatusContext(ctx context.Context, projectIDOrKey interface{}, statusID int, input *UpdateStatusInput) (*Status, error)
	DeleteStatus(projectIDOrKey interface{}, statusID int, input *DeleteStatusInput) (*Status, error)
	DeleteStatusContext(ctx context.Context, projectIDOrKey interface{}, statusID int, input *DeleteStatusInput) (*Status, error)
	SortStatuses(projectIDOrKey interface{}, input *SortStatusesInput) ([]*Status, error)
	SortStatusesContext(ctx context.Context, projectIDOrKey interface{}, input *SortStatusesInput) ([]*Status, error)
	GetProjectDiskUsage(projectIDOrKey interface{}) (*ProjectDiskUsage, error)
	GetProjectDiskUsageContext(ctx context.Context, projectIDOrKey interface{}) (*ProjectDiskUsage, error)
	GetResolutions() ([]*Resolution, error)
	GetResolutionsContext(ctx context.Context) ([]*Resolution, error)
	GetSpace() (*Space, error)
	GetSpaceContext(ctx context.Context) (*Space, error)
	GetSpaceIcon(writer io.Writer) error
	GetSpaceIconContext(ctx context.Context, writer io.Writer) error
	GetSpaceNotification() (*SpaceNotification, error)
	GetSpaceNotificationContext(ctx context.Context) (*SpaceNotification, error)
	UpdateSpaceNotification(input *UpdateSpaceNotificationInput) (*SpaceNotification, error)
	UpdateSpaceNotificationContext(ctx context.Context, input *UpdateSpaceNotificationInput) (*SpaceNotification, error)
	GetSpaceDiskUsage() (*SpaceDiskUsage, error)
	GetSpaceDiskUsageContext(ctx context.Context) (*SpaceDiskUsage, error)
	GetLicence() (*License, error)
	GetLicenceContext(ctx context.Context) (*License, error)
	GetTeams(opts *GetTeamsOptions) ([]*Team, error)
	GetTeamsContext(ctx context.Context, opts *GetTeamsOptions) ([]*Team, error)
	CreateTeam(input *CreateTeamInput) (*Team, error)
	CreateTeamContext(ctx context.Context, input *CreateTeamInput) (*Team, error)
	GetTeam(teamID int) (*Team, error)
	GetTeamContext(ctx context.Context, teamID int) (*Team, error)
	UpdateTeam(teamID int, input *UpdateTeamInput) (*Team, error)
	UpdateTeamContext(ctx context.Context, teamID int, input *UpdateTeamInput) (*Team, error)
	DeleteTeam(teamID int) (*Team, error)
	DeleteTeamContext(ctx context.Context, teamID int) (*Team, error)
	GetTeamIcon(teamID int, writer io.Writer) error
	GetTeamIconContext(ctx context.Context, teamID int, writer io.Writer) error
	GetProjectTeams(projectIDOrKey interface{}) ([]*Team, error)
	GetProjectTeamsContext(ctx context.Context, projectIDOrKey interface{}) ([]*Team, error)
	AddProjectTeam(projectIDOrKey interface{}, input *AddProjectTeamInput) (*Team, error)
	AddProjectTeamContext(ctx context.Context, projectIDOrKey interface{}, input *AddProjectTeamInput) (*Team, error)
	DeleteProjectTeam(projectIDOrKey interface{}, input *DeleteProjectTeamInput) (*Team, error)
	DeleteProjectTeamContext(ctx context.Context, projectIDOrKey interface{}, input *DeleteProjectTeamInput) (*Team, error)
	GetUserMySelf() (*User, error)
	GetUserMySelfContext(ctx context.Context) (*User, error)
	GetUser(id int) (*User, error)
	GetUserContext(ctx context.Context, id int) (*User, error)
	GetUsers() ([]*User, error)
	GetUsersContext(ctx context.Context) ([]*User, error)
	CreateUser(input *CreateUserInput) (*User, error)
	CreateUserContext(ctx context.Context, input *CreateUserInput) (*User, error)
	UpdateUser(id int, input *UpdateUserInput) (*User, error)
	UpdateUserContext(ctx context.Context, id int, input *UpdateUserInput) (*User, error)
	DeleteUser(id int) (*User, error)
	DeleteUserContext(ctx context.Context, id int) (*User, error)
	GetUserIcon(id int, writer io.Writer) error
	GetUserIconContext(ctx context.Context, id int, writer io.Writer) error
	GetUserStars(id int, opts *GetUserStarsOptions) ([]*Star, error)
	GetUserStarsContext(ctx context.Context, id int, opts *GetUserStarsOptions) ([]*Star, error)
	GetUserStarCount(id int, opts *GetUserStarCountOptions) (int, error)
	GetUserStarCountContext(ctx context.Context, id int, opts *GetUserStarCountOptions) (int, error)
	GetVersions(projectIDOrKey interface{}) ([]*Version, error)
	GetVersionsContext(ctx context.Context, projectIDOrKey interface{}) ([]*Version, error)
	CreateVersion(projectIDOrKey interface{}, input *CreateVersionInput) (*Version, error)
	CreateVersionContext(ctx context.Context, projectIDOrKey interface{}, input *CreateVersionInput) (*Version, error)
	UpdateVersion(projectIDOrKey interface{}, versionID int, input *UpdateVersionInput) (*Version, error)
	UpdateVersionContext(ctx context.Context, projectIDOrKey interface{}, versionID int, input *UpdateVersionInput) (*Version, error)
	DeleteVersion(projectIDOrKey interface{}, versionID int) (*Version, error)
	DeleteVersionContext(ctx context.Context, projectIDOrKey interface{}, versionID int) (*Version, error)
	GetUserWatchings(userID int) ([]*Watching, error)
	GetUserWatchingsContext(ctx context.Context, userID int) ([]*Watching, error)
	GetUserWatchingsCount(userID int, opts *GetUserWatchingsCountOptions) (int, error)
	GetUserWatchingsCountContext(ctx context.Context, userID int, opts *GetUserWatchingsCountOptions) (int, error)
	GetWatching(watchingID int) (*Watching, error)
	GetWatchingContext(ctx context.Context, watchingID int) (*Watching, error)
	CreateWatching(input *CreateWatchingInput) (*Watching, error)
	CreateWatchingContext(ctx context.Context, input *CreateWatchingInput) (*Watching, error)
	UpdateWatching(watchingID int, input *UpdateWatchingInput) (*Watching, error)
	UpdateWatchingContext(ctx context.Context, watchingID int, input *UpdateWatchingInput) (*Watching, error)
	DeleteWatching(watchingID int) (*Watching, error)
	DeleteWatchingContext(ctx context.Context, watchingID int) (*Watching, error)
	MarkAsReadWatching(watchingID int) error
	MarkAsReadWatchingContext(ctx context.Context, watchingID int) error
	GetWebhook(projectIDOrKey interface{}, webhookID int) (*Webhook, error)
	GetWebhookContext(ctx context.Context, projectIDOrKey interface{}, webhookID int) (*Webhook, error)
	GetWebhooks(projectIDOrKey interface{}) ([]*Webhook, error)
	GetWebhooksContext(ctx context.Context, projectIDOrKey interface{}) ([]*Webhook, error)
	CreateWebhook(projectIDOrKey interface{}, webhook *CreateWebhookInput) (*Webhook, error)
	CreateWebhookContext(ctx context.Context, projectIDOrKey interface{}, input *CreateWebhookInput) (*Webhook, error)
	UpdateWebhook(projectIDOrKey interface{}, webhookID int, input *UpdateWebhookInput) (*Webhook, error)
	UpdateWebhookContext(ctx context.Context, projectIDOrKey interface{}, webhookID int, input *UpdateWebhookInput) (*Webhook, error)
	DeleteWebhook(projectIDOrKey interface{}, webhookID int) (*Webhook, error)
	DeleteWebhookContext(ctx context.Context, projectIDOrKey interface{}, webhookID int) (*Webhook, error)
	GetMyRecentlyViewedWikis(opts *GetMyRecentlyViewedWikisOptions) ([]*RecentlyViewedWiki, error)
	GetMyRecentlyViewedWikisContext(ctx context.Context, opts *GetMyRecentlyViewedWikisOptions) ([]*RecentlyViewedWiki, error)
	GetWikis(opts *GetWikisOptions) ([]*Wiki, error)
	GetWikisContext(ctx context.Context, opts *GetWikisOptions) ([]*Wiki, error)
	GetWikiCount(opts *GetWikiCountOptions) (int, error)
	GetWikiCountContext(ctx context.Context, opts *GetWikiCountOptions) (int, error)
	GetWikiTags(opts *GetWikiTagsOptions) ([]*Tag, error)
	GetWikiTagsContext(ctx context.Context, opts *GetWikiTagsOptions) ([]*Tag, error)
	GetWiki(wikiID int) (*Wiki, error)
	GetWikiContext(ctx context.Context, wikiID int) (*Wiki, error)
	CreateWiki(input *CreateWikiInput) (*Wiki, error)
	CreateWikiContext(ctx context.Context, input *CreateWikiInput) (*Wiki, error)
	UpdateWiki(wikiID int, input *UpdateWikiInput) (*Wiki, error)
	UpdateWikiContext(ctx context.Context, wikiID int, input *UpdateWikiInput) (*Wiki, error)
	DeleteWiki(wikiID int) (*Wiki, error)
	DeleteWikiContext(ctx context.Context, wikiID int) (*Wiki, error)
	GetWikiAttachments(wikiID int) ([]*Attachment, error)
	GetWikiAttachmentsContext(ctx context.Context, wikiID int) ([]*Attachment, error)
	AddAttachmentToWiki(wikiID int, input *AddAttachmentToWikiInput) ([]*Attachment, error)
	AddAttachmentToWikiContext(ctx context.Context, wikiID int, input *AddAttachmentToWikiInput) ([]*Attachment, error)
	DeleteAttachmentInWiki(wikiID, attachmentID int) (*Attachment, error)
	DeleteAttachmentInWikiContext(ctx context.Context, wikiID, attachmentID int) (*Attachment, error)
}

var _ API = (*Client)(nil)
