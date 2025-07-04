package backlog

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

const testJSONUser string = `{
	"id": 1,
	"userId": "admin",
	"name": "admin",
	"roleType": 1,
	"lang": "ja",
	"mailAddress": "eguchi@nulab.example"
}`

const testJSONStar string = `{
	"id": 1,
	"comment": "",
	"url": "https://xx.backlog.jp/view/BLG-1",
	"title": "[BLG-1] first issue | 課題の表示 - Backlog",
	"presenter":{
		"id":1,
		"userId": "admin",
		"name":"admin",
		"roleType":1,
		"lang":"ja",
		"mailAddress":"eguchi@nulab.example"
	},
	"created": "2006-01-02T15:04:05Z"
}`

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

func getTestUser() *User {
	return &User{
		ID:          Int(1),
		UserID:      String("admin"),
		Name:        String("admin"),
		RoleType:    RoleType(1),
		Lang:        String("ja"),
		MailAddress: String("eguchi@nulab.example"),
	}
}

func getTestActivity() *Activity {
	return &Activity{
		ID: Int(1),
		Project: &Project{
			ID:                                Int(92),
			ProjectKey:                        String("SUB"),
			Name:                              String("サブタスク"),
			ChartEnabled:                      Bool(true),
			SubtaskingEnabled:                 Bool(true),
			ProjectLeaderCanEditProjectLeader: Bool(false),
			TextFormattingRule:                String(""),
			Archived:                          Bool(false),
			DisplayOrder:                      Int(0),
		},
		Type: Int(2),
		Content: &Content{
			ID:          Int(4809),
			KeyID:       Int(121),
			Summary:     String("コメント"),
			Description: String(""),
			Comment: &Comment{
				ID:      Int(7237),
				Content: String(""),
			},
			Changes: []*Change{
				{
					Field:    String("milestone"),
					NewValue: String(" R2014-07-23"),
					OldValue: String(""),
					Type:     String("standard"),
				},
				{
					Field:    String("status"),
					NewValue: String("4"),
					OldValue: String("1"),
					Type:     String("standard"),
				},
			},
		},
		Notifications: []*Notification{
			{
				ID:          Int(25),
				AlreadyRead: Bool(false),
				Reason:      Int(2),
				User: &User{
					ID:          Int(5686),
					UserID:      String("takada"),
					Name:        String("takada"),
					RoleType:    RoleType(2),
					Lang:        String("ja"),
					MailAddress: String("takada@nulab.example"),
				},
				ResourceAlreadyRead: Bool(false),
			},
		},
		CreatedUser: &User{
			ID:          Int(1),
			UserID:      String("admin"),
			Name:        String("admin"),
			RoleType:    RoleType(1),
			Lang:        String("ja"),
			MailAddress: String("eguchi@nulab.example"),
		},
		Created: &Timestamp{referenceTime},
	}
}

func getTestUserStarsWitID(id int) *Star {
	return &Star{
		ID:      Int(id),
		Comment: String(""),
		URL:     String("https://xx.backlog.jp/view/BLG-1"),
		Title:   String("[BLG-1] first issue | 課題の表示 - Backlog"),
		Presenter: &User{
			ID:          Int(1),
			UserID:      String("admin"),
			Name:        String("admin"),
			RoleType:    RoleType(1),
			Lang:        String("ja"),
			MailAddress: String("eguchi@nulab.example"),
		},
		Created: &Timestamp{referenceTime},
	}
}

func TestGetUserMySelf(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/myself", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, testJSONUser); err != nil {
			t.Fatal(err)
		}
	})

	user, err := client.GetUserMySelf()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestUser()
	if !reflect.DeepEqual(want, user) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetUserMySelfFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/myself", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetUserMySelf(); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetUserByID(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, testJSONUser); err != nil {
			t.Fatal(err)
		}
	})

	user, err := client.GetUser(1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestUser()
	if !reflect.DeepEqual(want, user) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetUserByIDFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/1", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetUser(1); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetUsers(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users", func(w http.ResponseWriter, _ *http.Request) {
		j := fmt.Sprintf("[%s]", testJSONUser)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	users, err := client.GetUsers()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	want := []*User{getTestUser()}
	if !reflect.DeepEqual(want, users) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetUsersFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetUsers(); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateUser(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := fmt.Fprint(w, testJSONUser); err != nil {
			t.Fatal(err)
		}
	})

	input := &CreateUserInput{
		UserID:      String("admin"),
		Password:    String("password"),
		Name:        String("admin"),
		MailAddress: String("eguchi@nulab.example"),
		RoleType:    RoleType(RoleTypeAdministrator),
	}
	user, err := client.CreateUser(input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestUser()
	if !reflect.DeepEqual(want, user) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestCreateUserFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &CreateUserInput{
		UserID:      String("admin"),
		Password:    String("password"),
		Name:        String("admin"),
		MailAddress: String("eguchi@nulab.example"),
		RoleType:    RoleTypeAdministrator,
	}
	if _, err := client.CreateUser(input); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateUser(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/1", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := fmt.Fprint(w, testJSONUser); err != nil {
			t.Fatal(err)
		}
	})

	input := &UpdateUserInput{
		Password:    String("password"),
		Name:        String("admin"),
		MailAddress: String("eguchi@nulab.example"),
		RoleType:    RoleTypeAdministrator,
	}
	user, err := client.UpdateUser(1, input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestUser()
	if !reflect.DeepEqual(want, user) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestUpdateUserFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/1", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &UpdateUserInput{
		Password:    String("password"),
		Name:        String("admin"),
		MailAddress: String("eguchi@nulab.example"),
		RoleType:    RoleTypeAdministrator,
	}
	if _, err := client.UpdateUser(1, input); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteUser(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/1", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := fmt.Fprint(w, testJSONUser); err != nil {
			t.Fatal(err)
		}
	})

	user, err := client.DeleteUser(1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestUser()
	if !reflect.DeepEqual(want, user) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestDeleteUserFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/1", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.DeleteUser(1); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetUserIcon(_ *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	client.httpclient = &mockHTTPClient{}

	err := client.GetUserIcon(1, &bytes.Buffer{})
	if err != nil {
		log.Fatalf("Unexpected error: %s in test", err)
	}
}

func TestGetUserStars(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/1/stars", func(w http.ResponseWriter, _ *http.Request) {
		j := fmt.Sprintf("[%s]", testJSONStar)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	input := &GetUserStarsOptions{
		MinID: Int(1),
		MaxID: Int(10),
		Count: Int(20),
		Order: OrderAsc,
	}
	stars, err := client.GetUserStars(1, input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*Star{getTestUserStarsWitID(1)}
	if !reflect.DeepEqual(want, stars) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetUserStarsFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/1/stars", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &GetUserStarsOptions{}
	if _, err := client.GetUserStars(1, input); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetUserStarCount(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/1/stars/count", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := fmt.Fprint(w, `{"count":54 }`); err != nil {
			t.Fatal(err)
		}
	})

	input := &GetUserStarCountOptions{
		Since: String("2019-01-07"),
		Until: String("2020-01-07"),
	}
	count, err := client.GetUserStarCount(1, input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	if !reflect.DeepEqual(54, count) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetUserStarCountFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/1/stars/count", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &GetUserStarCountOptions{}
	if _, err := client.GetUserStarCount(1, input); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetUserMySelf_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://example.com/api/v2/")
	client.baseURL = invalidURL

	_, err := client.GetUserMySelf()
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetUsers_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://example.com/api/v2/")
	client.baseURL = invalidURL

	_, err := client.GetUsers()
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestCreateUser_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://example.com/api/v2/")
	client.baseURL = invalidURL

	_, err := client.CreateUser(&CreateUserInput{UserID: String("test"), Password: String("pass"), Name: String("Test User"), MailAddress: String("test@example.com"), RoleType: RoleTypeAdministrator})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}
