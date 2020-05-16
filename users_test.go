package backlog

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

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

func getUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(getTestUser())
	rw.Write(response)
}

func getUsers(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(getTestUsers())
	rw.Write(response)
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
