package backlog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"
)

func get4xxError() http.Response {
	return http.Response{
		Status:     "400 Bad Request",
		StatusCode: 400,
		Proto:      "HTTP/2.0",
	}
}

func get4xxErrorResponse(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusBadRequest)
	response, _ := json.Marshal(get4xxError())
	if _, err := rw.Write(response); err != nil {
		fmt.Println(err)
	}
}

func TestOrder_String(t *testing.T) {
	tests := []struct {
		name string
		k    Order
		want string
	}{
		{
			name: "asc",
			k:    OrderAsc,
			want: "asc",
		},
		{
			name: "desc",
			k:    OrderDesc,
			want: "desc",
		},
		{
			name: "empty",
			k:    Order(""),
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.k.String(); got != tt.want {
				t.Errorf("Order.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoleType_Int(t *testing.T) {
	tests := []struct {
		name string
		k    RoleType
		want int
	}{
		{
			name: "administrator",
			k:    RoleTypeAdministrator,
			want: 1,
		},
		{
			name: "generalUser",
			k:    RoleTypeGeneralUser,
			want: 2,
		},
		{
			name: "reporter",
			k:    RoleTypeReporter,
			want: 3,
		},
		{
			name: "viewer",
			k:    RoleTypeViewer,
			want: 4,
		},
		{
			name: "guestReporter",
			k:    RoleTypeGuestReporter,
			want: 5,
		},
		{
			name: "guestViewer",
			k:    RoleTypeGuestViewer,
			want: 6,
		},
		{
			name: "empty",
			k:    RoleType(9999),
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.k.Int(); got != tt.want {
				t.Errorf("RoleType.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getTestUser() User {
	return getTestUserWithID(1)
}

func getTestUsers() []User {
	return []User{
		getTestUserWithID(1),
		getTestUserWithID(2),
		getTestUserWithID(3),
		getTestUserWithID(4),
	}
}

func getTestUserWithID(id int) User {
	return User{
		ID:          id,
		UserID:      "admin",
		Name:        "admin",
		RoleType:    1,
		Lang:        "ja",
		MailAddress: "eguchi@nulab.example",
	}
}

func getTestUserActivities() []UserActivity {
	return []UserActivity{
		getTestUserActivityWithID(1),
		getTestUserActivityWithID(2),
		getTestUserActivityWithID(3),
	}
}

func getTestUserActivityWithID(id int) UserActivity {
	return UserActivity{
		ID: id,
		Project: Project{
			ID:                                92,
			ProjectKey:                        "SUB",
			Name:                              "サブタスク",
			ChartEnabled:                      true,
			SubtaskingEnabled:                 true,
			ProjectLeaderCanEditProjectLeader: false,
			TextFormattingRule:                "",
			Archived:                          false,
			DisplayOrder:                      0,
		},
		Type: 2,
		Content: Content{
			ID:          4809,
			KeyID:       121,
			Summary:     "コメント",
			Description: "",
			Comment: Comment{
				ID:      7237,
				Content: "",
			},
			Changes: []Change{
				{
					Field:    "milestone",
					NewValue: " R2014-07-23",
					OldValue: "",
					Type:     "standard",
				},
				{
					Field:    "status",
					NewValue: "4",
					OldValue: "1",
					Type:     "standard",
				},
			},
		},
		Notifications: []Notification{
			{
				ID:          25,
				AlreadyRead: false,
				Reason:      2,
				User: User{
					ID:          5686,
					UserID:      "takada",
					Name:        "takada",
					RoleType:    RoleType(2),
					Lang:        "ja",
					MailAddress: "takada@nulab.example",
				},
				ResourceAlreadyRead: false,
			},
		},
		CreatedUser: User{
			ID:          1,
			UserID:      "admin",
			Name:        "admin",
			RoleType:    RoleType(1),
			Lang:        "ja",
			MailAddress: "eguchi@nulab.example",
		},
		Created: JSONTime("2013-12-27T07:50:44Z"),
	}
}

func getTestUserStars() []Star {
	return []Star{
		getTestUserStarsWitID(1),
		getTestUserStarsWitID(2),
		getTestUserStarsWitID(3),
	}
}

func getTestUserStarsWitID(id int) Star {
	return Star{
		ID:      id,
		Comment: "",
		URL:     "https://xx.backlog.jp/view/BLG-1",
		Title:   "[BLG-1] first issue | 課題の表示 - Backlog",
		Presenter: User{
			ID:          1,
			UserID:      "admin",
			Name:        "admin",
			RoleType:    1,
			Lang:        "ja",
			MailAddress: "eguchi@nulab.example",
		},
		Created: JSONTime("2014-01-23T10:55:19Z"),
	}
}

func getUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(getTestUser())
	if _, err := rw.Write(response); err != nil {
		log.Fatal(err)
	}
}

func getUsers(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(getTestUsers())
	if _, err := rw.Write(response); err != nil {
		log.Fatal(err)
	}
}

func getUserActivities(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(getTestUserActivities())
	if _, err := rw.Write(response); err != nil {
		log.Fatal(err)
	}
}

func getUserStars(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(getTestUserStars())
	if _, err := rw.Write(response); err != nil {
		log.Fatal(err)
	}
}

func getUserStarCount(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	type c struct {
		Count int
	}
	t := c{Count: 1}
	response, _ := json.Marshal(t)
	if _, err := rw.Write(response); err != nil {
		log.Fatal(err)
	}
}

func TestGetUserMySelf(t *testing.T) {
	http.HandleFunc("/api/v2/users/myself", getUser)
	expectedUser := getTestUser()

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	user, err := api.GetUserMySelf()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if !reflect.DeepEqual(expectedUser, *user) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetUserMySelfFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/users/myself", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	if _, err := api.GetUserMySelf(); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetUserByID(t *testing.T) {
	http.HandleFunc("/api/v2/users/1", getUser)
	expectedUser := getTestUser()

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	user, err := api.GetUserByID(1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if !reflect.DeepEqual(expectedUser, *user) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetUserByIDFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/users/myself", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	if _, err := api.GetUserByID(1); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetUsers(t *testing.T) {
	http.HandleFunc("/api/v2/users", getUsers)
	expectedUsers := getTestUsers()

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	users, err := api.GetUsers()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if !reflect.DeepEqual(expectedUsers, users) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetUsersFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/users/myself", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	if _, err := api.GetUsers(); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateUser(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/users", getUser)
	expected := getTestUser()

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	input := &CreateUserInput{
		UserID:      "admin",
		Password:    "password",
		Name:        "admin",
		MailAddress: "eguchi@nulab.example",
		RoleType:    RoleTypeAdministrator,
	}
	user, err := api.CreateUser(input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	if !reflect.DeepEqual(expected, *user) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestCreateUserFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/users", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	input := &CreateUserInput{
		UserID:      "admin",
		Password:    "password",
		Name:        "admin",
		MailAddress: "eguchi@nulab.example",
		RoleType:    RoleTypeAdministrator,
	}
	if _, err := api.CreateUser(input); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateUser(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/users/1", getUser)
	expected := getTestUser()

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	input := &UpdateUserInput{
		ID:          1,
		Password:    "password",
		Name:        "admin",
		MailAddress: "eguchi@nulab.example",
		RoleType:    RoleTypeAdministrator,
	}
	user, err := api.UpdateUser(input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	if !reflect.DeepEqual(expected, *user) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestUpdateUserFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/users/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	input := &UpdateUserInput{
		ID:          1,
		Password:    "password",
		Name:        "admin",
		MailAddress: "eguchi@nulab.example",
		RoleType:    RoleTypeAdministrator,
	}
	if _, err := api.UpdateUser(input); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteUser(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/users/1", getUser)
	expected := getTestUser()

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	user, err := api.DeleteUser(1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	if !reflect.DeepEqual(expected, *user) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestDeleteUserFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/users/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	if _, err := api.DeleteUser(1); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetUserIcon(t *testing.T) {
	api := &Client{
		endpoint:   "http://" + serverAddr + "/",
		apiKey:     "testing-token",
		httpclient: &mockHTTPClient{},
	}

	err := api.GetUserIcon(1, &bytes.Buffer{})
	if err != nil {
		log.Fatalf("Unexpected error: %s in test", err)
	}
}

func TestGetUserActivities(t *testing.T) {
	http.HandleFunc("/api/v2/users/1/activities", getUserActivities)
	expected := getTestUserActivities()

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	input := &GetUserActivityInput{
		ID:              1,
		ActivityTypeIDs: []int{1, 2, 3},
		MinID:           1,
		MaxID:           10,
		Count:           20,
		Order:           OrderAsc,
	}
	activities, err := api.GetUserActivities(input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if !reflect.DeepEqual(expected, activities) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetUserActivitiesFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/users/myself", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	input := &GetUserActivityInput{}
	if _, err := api.GetUserActivities(input); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetUserStars(t *testing.T) {
	http.HandleFunc("/api/v2/users/1/stars", getUserStars)
	expected := getTestUserStars()

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	input := &GetUserStarsInput{
		ID:    1,
		MinID: 1,
		MaxID: 10,
		Count: 20,
		Order: OrderAsc,
	}
	stars, err := api.GetUserStars(input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if !reflect.DeepEqual(expected, stars) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetUserStarsFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/users/1/stars", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	input := &GetUserStarsInput{}
	if _, err := api.GetUserStars(input); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetUserStarCount(t *testing.T) {
	http.HandleFunc("/api/v2/users/1/stars/count", getUserStarCount)
	expected := 1

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	input := &GetUserStarCountInput{
		ID:    1,
		Since: "2019-01-07",
		Until: "2020-01-07",
	}
	count, err := api.GetUserStarCount(input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if !reflect.DeepEqual(expected, count) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetUserStarCountFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/users/1/stars", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	input := &GetUserStarCountInput{}
	if _, err := api.GetUserStarCount(input); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetUserMySelfRecentrlyViewedIssues(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/users/myself/recentlyViewedIssues", getIssues)
	expected := getTestIssues()

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	input := &GetUserMySelfRecentrlyViewedIssuesInput{
		Order:  OrderAsc,
		Offset: 1,
		Count:  100,
	}
	issues, err := api.GetUserMySelfRecentrlyViewedIssues(input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if !reflect.DeepEqual(expected, issues) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetUserMySelfRecentrlyViewedIssuesFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/users/myself/recentlyViewedIssues", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	input := &GetUserMySelfRecentrlyViewedIssuesInput{}
	_, err := api.GetUserMySelfRecentrlyViewedIssues(input)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetUserMySelfRecentrlyViewedIssues4xxErrorFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/users/myself/recentlyViewedIssues", get4xxErrorResponse)

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	input := &GetUserMySelfRecentrlyViewedIssuesInput{}

	if _, err := api.GetUserMySelfRecentrlyViewedIssues(input); err == nil {
		t.Fatal("expected an error but got none")
	}
}
