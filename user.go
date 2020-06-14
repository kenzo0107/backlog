package backlog

import (
	"context"
	"io"
	"net/url"
	"strconv"
)

// RoleType : role type
type RoleType int

// RoleType
const (
	RoleTypeAdministrator = RoleType(iota)
	RoleTypeGeneralUser
	RoleTypeReporter
	RoleTypeViewer
	RoleTypeGuestReporter
	RoleTypeGuestViewer
)

// Int converts const RoleType* into int
func (k RoleType) Int() int {
	switch k {
	case RoleTypeAdministrator:
		return 1
	case RoleTypeGeneralUser:
		return 2
	case RoleTypeReporter:
		return 3
	case RoleTypeViewer:
		return 4
	case RoleTypeGuestReporter:
		return 5
	case RoleTypeGuestViewer:
		return 6
	default:
		return 0
	}
}

// Order : asc or desc
type Order string

// Order by asc, desc
const (
	OrderAsc = Order(iota)
	OrderDesc
)

func (k Order) String() string {
	switch k {
	case OrderAsc:
		return "asc"
	case OrderDesc:
		return "desc"
	default:
		return ""
	}
}

// User : backlog user
type User struct {
	ID          int      `json:"id"`
	UserID      string   `json:"userId"`
	Name        string   `json:"name"`
	RoleType    RoleType `json:"roleType"`
	Lang        string   `json:"lang"`
	MailAddress string   `json:"mailAddress"`
}

// UserActivity : user's activity
type UserActivity struct {
	ID            int            `json:"id"` // User.ID
	Project       Project        `json:"project"`
	Type          int            `json:"type"`
	Content       Content        `json:"content"`
	Notifications []Notification `json:"notifications"`
	CreatedUser   User           `json:"createdUser"`
	Created       JSONTime       `json:"created"`
}

// Notification : -
type Notification struct {
	ID                  int  `json:"id"`
	AlreadyRead         bool `json:"alreadyRead"`
	Reason              int  `json:"reason"`
	User                User `json:"user"`
	ResourceAlreadyRead bool `json:"resourceAlreadyRead"`
}

// Content : -
type Content struct {
	ID          int      `json:"id"`
	KeyID       int      `json:"key_id"`
	Summary     string   `json:"summary"`
	Description string   `json:"description"`
	Comment     Comment  `json:"comment"`
	Changes     []Change `json:"changes"`
}

// Comment : -
type Comment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

// Change : -
type Change struct {
	Field    string `json:"field"`
	NewValue string `json:"new_value"`
	OldValue string `json:"old_value"`
	Type     string `json:"type"`
}

// ResponseIssue : response of issue api
type ResponseIssue struct {
	Issue   Issue    `json:"issue"`
	Updated JSONTime `json:"updated"`
}

// Priority : -
type Priority struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetUserMySelf returns get my user information
func (api *Client) GetUserMySelf() (*User, error) {
	return api.GetUserMySelfContext(context.Background())
}

// GetUserMySelfContext will retrieve the complete my user information by id with a custom context
func (api *Client) GetUserMySelfContext(ctx context.Context) (*User, error) {
	user := User{}
	if err := api.getMethod(ctx, "/api/v2/users/myself", url.Values{}, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByID returns a user by id
func (api *Client) GetUserByID(id int) (*User, error) {
	return api.GetUserByIDContext(context.Background(), id)
}

// GetUserByIDContext will retrieve the complete user information by id with a custom context
func (api *Client) GetUserByIDContext(ctx context.Context, id int) (*User, error) {
	var user User
	if err := api.getMethod(ctx, "/api/v2/users/"+strconv.Itoa(id), url.Values{}, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUsers returns the list of users
func (api *Client) GetUsers() ([]User, error) {
	return api.GetUsersContext(context.Background())
}

// GetUsersContext returns the list of users
func (api *Client) GetUsersContext(ctx context.Context) ([]User, error) {
	var users []User
	if err := api.getMethod(ctx, "/api/v2/users", url.Values{}, &users); err != nil {
		return nil, err
	}
	return users, nil
}

// CreateUser creates a user
func (api *Client) CreateUser(input *CreateUserInput) (*User, error) {
	return api.CreateUserContext(context.Background(), input)
}

// CreateUserContext creates a user with Context
func (api *Client) CreateUserContext(ctx context.Context, input *CreateUserInput) (*User, error) {
	values := url.Values{
		"userId":      {input.UserID},
		"password":    {input.Password},
		"name":        {input.Name},
		"mailAddress": {input.MailAddress},
		"roleType":    {strconv.Itoa(input.RoleType.Int())},
	}

	user := User{}
	if err := api.postMethod(ctx, "/api/v2/users", values, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates a user
func (api *Client) UpdateUser(input *UpdateUserInput) (*User, error) {
	return api.UpdateUserContext(context.Background(), input)
}

// UpdateUserContext updates a user with Context
func (api *Client) UpdateUserContext(ctx context.Context, input *UpdateUserInput) (*User, error) {
	values := url.Values{
		"password":    {input.Password},
		"name":        {input.Name},
		"mailAddress": {input.MailAddress},
		"roleType":    {strconv.Itoa(input.RoleType.Int())},
	}

	user := User{}
	if err := api.patchMethod(ctx, "/api/v2/users/"+strconv.Itoa(input.ID), values, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// DeleteUser deletes a user
func (api *Client) DeleteUser(id int) (*User, error) {
	return api.DeleteUserContext(context.Background(), id)
}

// DeleteUserContext deletes a user with Context
func (api *Client) DeleteUserContext(ctx context.Context, id int) (*User, error) {
	user := User{}
	if err := api.deleteMethod(ctx, "/api/v2/users/"+strconv.Itoa(id), url.Values{}, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserIcon downloads user icon
func (api *Client) GetUserIcon(id int, writer io.Writer) error {
	return api.GetUserIconContext(context.Background(), id, writer)
}

// GetUserIconContext downloads user icon with context
func (api *Client) GetUserIconContext(ctx context.Context, id int, writer io.Writer) error {
	return downloadFile(ctx, api.httpclient, api.apiKey, api.endpoint+"/api/v2/users/"+strconv.Itoa(id)+"/icon", writer, api)
}

// GetUserActivities returns the list of a user's activities
func (api *Client) GetUserActivities(input *GetUserActivityInput) ([]UserActivity, error) {
	return api.GetUserActivitiesContext(context.Background(), input)
}

// GetUserActivitiesContext returns the list of a user's activities with context
func (api *Client) GetUserActivitiesContext(ctx context.Context, input *GetUserActivityInput) ([]UserActivity, error) {
	values := url.Values{}

	if len(input.ActivityTypeIDs) > 0 {
		for _, i := range input.ActivityTypeIDs {
			values.Add("activityTypeId[]", strconv.Itoa(i))
		}
	}

	if input.MinID > 0 {
		values.Add("minId", strconv.Itoa(input.MinID))
	}

	if input.MaxID > 0 {
		values.Add("minId", strconv.Itoa(input.MaxID))
	}

	if input.Count > 0 {
		values.Add("count", strconv.Itoa(input.Count))
	}

	if input.Order.String() == "" {
		values.Add("order", OrderDesc.String())
	} else {
		values.Add("order", input.Order.String())
	}

	var userActivities []UserActivity
	if err := api.getMethod(ctx, "/api/v2/users/"+strconv.Itoa(input.ID)+"/activities", values, &userActivities); err != nil {
		return nil, err
	}
	return userActivities, nil
}

// GetUserStars returns the list of stared contents
func (api *Client) GetUserStars(input *GetUserStarsInput) ([]Star, error) {
	return api.GetUserStarsContext(context.Background(), input)
}

// GetUserStarsContext returns the list of a user's activities with context
func (api *Client) GetUserStarsContext(ctx context.Context, input *GetUserStarsInput) ([]Star, error) {

	values := url.Values{}

	if input.MinID > 0 {
		values.Add("minId", strconv.Itoa(input.MinID))
	}

	if input.MaxID > 0 {
		values.Add("minId", strconv.Itoa(input.MaxID))
	}

	if input.Count > 0 {
		values.Add("count", strconv.Itoa(input.Count))
	}

	if input.Order.String() == "" {
		values.Add("order", OrderDesc.String())
	} else {
		values.Add("order", input.Order.String())
	}

	var stars []Star
	if err := api.getMethod(ctx, "/api/v2/users/"+strconv.Itoa(input.ID)+"/stars", values, &stars); err != nil {
		return nil, err
	}
	return stars, nil
}

// GetUserStarCount returns the count of stars
func (api *Client) GetUserStarCount(input *GetUserStarCountInput) (int, error) {
	return api.GetUserStarCountContext(context.Background(), input)
}

// GetUserStarCountContext returns the count of stars with context
func (api *Client) GetUserStarCountContext(ctx context.Context, input *GetUserStarCountInput) (int, error) {

	values := url.Values{}

	if input.Since != "" {
		values.Add("since", input.Since)
	}

	if input.Until != "" {
		values.Add("until", input.Until)
	}

	type c struct {
		Count int
	}
	r := c{}

	if err := api.getMethod(ctx, "/api/v2/users/"+strconv.Itoa(input.ID)+"/stars/count", values, &r); err != nil {
		return 0, err
	}
	return r.Count, nil
}

// GetUserMySelfRecentrlyViewedIssues returns the list of issues a user view recently
func (api *Client) GetUserMySelfRecentrlyViewedIssues(input *GetUserMySelfRecentrlyViewedIssuesInput) ([]Issue, error) {
	return api.GetUserMySelfRecentrlyViewedIssuesContext(context.Background(), input)
}

// GetUserMySelfRecentrlyViewedIssuesContext returns the list of issues a user view recently with context
func (api *Client) GetUserMySelfRecentrlyViewedIssuesContext(ctx context.Context, input *GetUserMySelfRecentrlyViewedIssuesInput) ([]Issue, error) {

	values := url.Values{}

	if input.Order.String() == "" {
		values.Add("order", OrderDesc.String())
	} else {
		values.Add("order", input.Order.String())
	}

	if input.Offset > 0 {
		values.Add("offset", strconv.Itoa(input.Offset))
	}

	if input.Count > 0 {
		values.Add("count", strconv.Itoa(input.Count))
	}

	var issues []Issue
	if err := api.getMethod(ctx, "/api/v2/users/myself/recentlyViewedIssues", url.Values{}, &issues); err != nil {
		return nil, err
	}
	return issues, nil
}

// CreateUserInput contains all the parameters necessary (including the optional ones) for a CreateUser() request.
type CreateUserInput struct {
	UserID      string   `required:"true"`
	Password    string   `required:"true"`
	Name        string   `required:"true"`
	MailAddress string   `required:"true"`
	RoleType    RoleType `required:"true"`
}

// UpdateUserInput contains all the parameters necessary (including the optional ones) for a UpdateUser() request.
type UpdateUserInput struct {
	ID          int      `required:"true"`
	Password    string   `required:"true"`
	Name        string   `required:"true"`
	MailAddress string   `required:"true"`
	RoleType    RoleType `required:"true"`
}

// GetUserActivityInput contains all the parameters necessary (including the optional ones) for a GetUserActivities() request.
type GetUserActivityInput struct {
	ID              int   `required:"true"`
	ActivityTypeIDs []int `required:"false"`
	MinID           int   `required:"false"`
	MaxID           int   `required:"false"`
	Count           int   `required:"false"`
	Order           Order `required:"false"`
}

// GetUserStarsInput contains all the parameters necessary (including the optional ones) for a GetUserStars() request.
type GetUserStarsInput struct {
	ID    int   `required:"true"`
	MinID int   `required:"false"`
	MaxID int   `required:"false"`
	Count int   `required:"false"`
	Order Order `required:"false"`
}

// GetUserStarCountInput contains all the parameters necessary (including the optional ones) for a GetUserStarCount() request.
type GetUserStarCountInput struct {
	ID    int    `required:"true"`
	Since string `required:"false"`
	Until string `required:"false"`
}

// GetUserMySelfRecentrlyViewedIssuesInput contains all the parameters necessary (including the optional ones) for a GetUserMySelfRecentrlyViewedIssues() request.
type GetUserMySelfRecentrlyViewedIssuesInput struct {
	Order  Order `required:"false"`
	Offset int   `required:"false"`
	Count  int   `required:"false"`
}
