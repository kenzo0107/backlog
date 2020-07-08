package backlog

import (
	"context"
	"fmt"
	"io"
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
	ID          *int     `json:"id,omitempty"`
	UserID      *string  `json:"userId,omitempty"`
	Name        *string  `json:"name,omitempty"`
	RoleType    RoleType `json:"roleType"`
	Lang        *string  `json:"lang,omitempty"`
	MailAddress *string  `json:"mailAddress,omitempty"`
}

// UserActivity : user's activity
type UserActivity struct {
	ID            *int            `json:"id,omitempty"` // User.ID
	Project       *Project        `json:"project,omitempty"`
	Type          *int            `json:"type,omitempty"`
	Content       *Content        `json:"content,omitempty"`
	Notifications []*Notification `json:"notifications,omitempty"`
	CreatedUser   *User           `json:"createdUser,omitempty"`
	Created       *Timestamp      `json:"created,omitempty"`
}

// Notification : -
type Notification struct {
	ID                  *int  `json:"id,omitempty"`
	AlreadyRead         *bool `json:"alreadyRead,omitempty"`
	Reason              *int  `json:"reason,omitempty"`
	User                *User `json:"user,omitempty"`
	ResourceAlreadyRead *bool `json:"resourceAlreadyRead,omitempty"`
}

// Content : -
type Content struct {
	ID          *int      `json:"id,omitempty"`
	KeyID       *int      `json:"key_id,omitempty"`
	Summary     *string   `json:"summary,omitempty"`
	Description *string   `json:"description,omitempty"`
	Comment     *Comment  `json:"comment,omitempty"`
	Changes     []*Change `json:"changes,omitempty"`
}

// Comment : -
type Comment struct {
	ID      *int    `json:"id,omitempty"`
	Content *string `json:"content,omitempty"`
}

// Change : -
type Change struct {
	Field    *string `json:"field,omitempty"`
	NewValue *string `json:"new_value,omitempty"`
	OldValue *string `json:"old_value,omitempty"`
	Type     *string `json:"type,omitempty"`
}

// ResponseIssue : response of issue api
type ResponseIssue struct {
	Issue   *Issue     `json:"issue,omitempty"`
	Updated *Timestamp `json:"updated,omitempty"`
}

// Priority : -
type Priority struct {
	ID   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// GetUserMySelf returns get my user information
func (c *Client) GetUserMySelf() (*User, error) {
	return c.GetUserMySelfContext(context.Background())
}

// GetUserMySelfContext will retrieve the complete my user information by id with a custom context
func (c *Client) GetUserMySelfContext(ctx context.Context) (*User, error) {
	u := "/api/v2/users/myself"

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	user := new(User)
	if err := c.Do(ctx, req, &user); err != nil {
		return nil, err
	}
	return user, nil
}

// GetUser returns a user by id
func (c *Client) GetUser(id int) (*User, error) {
	return c.GetUserContext(context.Background(), id)
}

// GetUserContext will retrieve the complete user information by id with a custom context
func (c *Client) GetUserContext(ctx context.Context, id int) (*User, error) {
	u := fmt.Sprintf("/api/v2/users/%v", id)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	user := new(User)
	if err := c.Do(ctx, req, &user); err != nil {
		return nil, err
	}
	return user, nil
}

// GetUsers returns the list of users
func (c *Client) GetUsers() ([]*User, error) {
	return c.GetUsersContext(context.Background())
}

// GetUsersContext returns the list of users
func (c *Client) GetUsersContext(ctx context.Context) ([]*User, error) {
	u := "/api/v2/users"

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var users []*User
	if err := c.Do(ctx, req, &users); err != nil {
		return nil, err
	}
	return users, nil
}

// CreateUser creates a user
func (c *Client) CreateUser(input *CreateUserInput) (*User, error) {
	return c.CreateUserContext(context.Background(), input)
}

// CreateUserContext creates a user with Context
func (c *Client) CreateUserContext(ctx context.Context, input *CreateUserInput) (*User, error) {
	u := "/api/v2/users"

	req, err := c.NewRequest("POST", u, input)
	if err != nil {
		return nil, err
	}

	user := new(User)
	if err := c.Do(ctx, req, &user); err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser updates a user
func (c *Client) UpdateUser(id int, input *UpdateUserInput) (*User, error) {
	return c.UpdateUserContext(context.Background(), id, input)
}

// UpdateUserContext updates a user with Context
func (c *Client) UpdateUserContext(ctx context.Context, id int, input *UpdateUserInput) (*User, error) {
	u := fmt.Sprintf("/api/v2/users/%v", id)

	req, err := c.NewRequest("PATCH", u, input)
	if err != nil {
		return nil, err
	}

	user := new(User)
	if err := c.Do(ctx, req, &user); err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUser deletes a user
func (c *Client) DeleteUser(id int) (*User, error) {
	return c.DeleteUserContext(context.Background(), id)
}

// DeleteUserContext deletes a user with Context
func (c *Client) DeleteUserContext(ctx context.Context, id int) (*User, error) {
	u := fmt.Sprintf("/api/v2/users/%v", id)

	req, err := c.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	user := new(User)
	if err := c.Do(ctx, req, &user); err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserIcon downloads user icon
func (c *Client) GetUserIcon(id int, writer io.Writer) error {
	return c.GetUserIconContext(context.Background(), id, writer)
}

// GetUserIconContext downloads user icon with context
func (c *Client) GetUserIconContext(ctx context.Context, id int, writer io.Writer) error {
	u := fmt.Sprintf("/api/v2/users/%v/icon", id)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, writer); err != nil {
		return err
	}

	return nil
}

// GetUserActivities returns the list of a user's activities
func (c *Client) GetUserActivities(id int, opts *GetUserActivityOptions) ([]*UserActivity, error) {
	return c.GetUserActivitiesContext(context.Background(), id, opts)
}

// GetUserActivitiesContext returns the list of a user's activities with context
func (c *Client) GetUserActivitiesContext(ctx context.Context, id int, opts *GetUserActivityOptions) ([]*UserActivity, error) {
	u := fmt.Sprintf("/api/v2/users/%v/activities", id)

	u, err := c.AddOptions(u, opts)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var userActivities []*UserActivity
	if err := c.Do(ctx, req, &userActivities); err != nil {
		return nil, err
	}
	return userActivities, nil
}

// GetUserStars returns the list of stared contents
func (c *Client) GetUserStars(id int, opts *GetUserStarsOptions) ([]*Star, error) {
	return c.GetUserStarsContext(context.Background(), id, opts)
}

// GetUserStarsContext returns the list of a user's activities with context
func (c *Client) GetUserStarsContext(ctx context.Context, id int, opts *GetUserStarsOptions) ([]*Star, error) {
	u := fmt.Sprintf("/api/v2/users/%v/stars", id)

	u, err := c.AddOptions(u, opts)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var stars []*Star
	if err := c.Do(ctx, req, &stars); err != nil {
		return nil, err
	}
	return stars, nil
}

// GetUserStarCount returns the count of stars
func (c *Client) GetUserStarCount(id int, opts *GetUserStarCountOptions) (int, error) {
	return c.GetUserStarCountContext(context.Background(), id, opts)
}

// GetUserStarCountContext returns the count of stars with context
func (c *Client) GetUserStarCountContext(ctx context.Context, id int, opts *GetUserStarCountOptions) (int, error) {
	u := fmt.Sprintf("/api/v2/users/%v/stars/count", id)

	u, err := c.AddOptions(u, opts)
	if err != nil {
		return 0, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return 0, err
	}

	type p struct {
		Count int
	}
	r := new(p)

	if err := c.Do(ctx, req, &r); err != nil {
		return 0, err
	}
	return r.Count, nil
}

// CreateUserInput contains all the parameters necessary (including the optional ones) for a CreateUser() request.
type CreateUserInput struct {
	UserID      *string
	Password    *string
	Name        *string
	MailAddress *string
	RoleType    RoleType
}

// UpdateUserInput contains all the parameters necessary (including the optional ones) for a UpdateUser() request.
type UpdateUserInput struct {
	Password    *string
	Name        *string
	MailAddress *string
	RoleType    RoleType
}

// GetUserActivityOptions specifies optional parameters to the GetUserActivities method.
type GetUserActivityOptions struct {
	ActivityTypeIDs []int `url:"activityTypeId[],omitempty"`
	MinID           *int  `url:"minId,omitempty"`
	MaxID           *int  `url:"maxId,omitempty"`
	Count           *int  `url:"count,omitempty"`
	Order           Order `url:"order,omitempty"`
}

// GetUserStarsOptions specifies optional parameters to the GetUserStars method.
type GetUserStarsOptions struct {
	MinID *int  `url:"minId,omitempty"`
	MaxID *int  `url:"maxId,omitempty"`
	Count *int  `url:"count,omitempty"`
	Order Order `url:"order,omitempty"`
}

// GetUserStarCountOptions specifies optional parameters to the GetUserStarCount method.
type GetUserStarCountOptions struct {
	Since *string `url:"since,omitempty"`
	Until *string `url:"until,omitempty"`
}
