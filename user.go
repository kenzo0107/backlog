package backlog

import (
	"context"
	"net/url"
	"strconv"
)

// User : backlog user
type User struct {
	ID          int    `json:"id"`
	UserID      string `json:"userId"`
	Name        string `json:"name"`
	RoleType    int    `json:"roleType"`
	Lang        string `json:"lang"`
	MailAddress string `json:"mailAddress"`
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
func (api *Client) GetUserByID(userID int) (*User, error) {
	return api.GetUserByIDContext(context.Background(), userID)
}

// GetUserByIDContext will retrieve the complete user information by id with a custom context
func (api *Client) GetUserByIDContext(ctx context.Context, userID int) (*User, error) {
	var user User
	if err := api.getMethod(ctx, "/api/v2/users/"+strconv.Itoa(userID), url.Values{}, &user); err != nil {
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
