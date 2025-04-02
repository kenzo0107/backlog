package backlog

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/pkg/errors"
)

const testJSONTeam string = `{
	"id": 1,
	"name": "test",
	"members": [
		{
			"id": 2,
			"userId": "developer",
			"name": "developer",
			"roleType": 2,
			"lang": null,
			"mailAddress": "developer@nulab.example"
		}
	],
	"displayOrder": null,
	"createdUser": {
		"id": 1,
		"userId": "admin",
		"name": "admin",
		"roleType": 1,
		"lang": "ja",
		"mailAddress": "eguchi@nulab.example"
	},
	"created": "2006-01-02T15:04:05Z",
	"updatedUser": {
		"id": 1,
		"userId": "admin",
		"name": "admin",
		"roleType": 1,
		"lang": "ja",
		"mailAddress": "eguchi@nulab.example"
	},
	"updated": "2006-01-02T15:04:05Z"
}`

func getTestTeam() *Team {
	return &Team{
		ID:   Int(1),
		Name: String("test"),
		Members: []*User{
			{
				ID:          Int(2),
				UserID:      String("developer"),
				Name:        String("developer"),
				RoleType:    RoleType(2),
				Lang:        nil,
				MailAddress: String("developer@nulab.example"),
			},
		},
		DisplayOrder: nil,
		CreatedUser: &User{
			ID:          Int(1),
			UserID:      String("admin"),
			Name:        String("admin"),
			RoleType:    RoleType(1),
			Lang:        String("ja"),
			MailAddress: String("eguchi@nulab.example"),
		},
		Created: &Timestamp{referenceTime},
		UpdatedUser: &User{
			ID:          Int(1),
			UserID:      String("admin"),
			Name:        String("admin"),
			RoleType:    RoleType(1),
			Lang:        String("ja"),
			MailAddress: String("eguchi@nulab.example"),
		},
		Updated: &Timestamp{referenceTime},
	}
}

func TestGetTeams(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teams", func(w http.ResponseWriter, _ *http.Request) {
		j := fmt.Sprintf("[%s]", testJSONTeam)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetTeams(&GetTeamsOptions{})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*Team{getTestTeam()}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetTeamsFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teams", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetTeams(&GetTeamsOptions{}); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateTeam(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teams", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := fmt.Fprint(w, testJSONTeam); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.CreateTeam(&CreateTeamInput{
		Name:    String("test"),
		Members: []int{2},
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestTeam()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestCreateTeamFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teams", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.CreateTeam(&CreateTeamInput{
		Name:    String("test"),
		Members: []int{2},
	}); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetTeam(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teams/1", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := fmt.Fprint(w, testJSONTeam); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetTeam(1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestTeam()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetTeamFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teams/1", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetTeam(1); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateTeam(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teams/1", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := fmt.Fprint(w, testJSONTeam); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateTeam(1, &UpdateTeamInput{
		Name:    String("test"),
		Members: []int{2},
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestTeam()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateTeamFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teams/1", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.UpdateTeam(1, &UpdateTeamInput{
		Name:    String("test"),
		Members: []int{2},
	}); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteTeam(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teams/1", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := fmt.Fprint(w, testJSONTeam); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.DeleteTeam(1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestTeam()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestDeleteTeamFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teams/1", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.DeleteTeam(1); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetTeamIcon(_ *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	client.httpclient = &mockHTTPClient{}

	err := client.GetTeamIcon(1, &bytes.Buffer{})
	if err != nil {
		log.Fatalf("Unexpected error: %s in test", err)
	}
}

func TestGetTeamIconFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teams/1/icon", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if err := client.GetTeamIcon(1, &bytes.Buffer{}); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetProjectTeams(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/teams", func(w http.ResponseWriter, _ *http.Request) {
		j := fmt.Sprintf("[%s]", testJSONTeam)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetProjectTeams("SRE")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*Team{getTestTeam()}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetProjectTeamsFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/teams", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetProjectTeams("SRE"); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetProjectTeamsInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	if _, err := client.GetProjectTeams("%%"); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestAddProjectTeam(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/teams", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := fmt.Fprint(w, testJSONTeam); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.AddProjectTeam("SRE", &AddProjectTeamInput{
		TeamID: Int(1),
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestTeam()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestAddProjectTeamFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/teams", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.AddProjectTeam("SRE", &AddProjectTeamInput{
		TeamID: Int(1),
	}); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestAddProjectTeamInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	if _, err := client.AddProjectTeam("%", &AddProjectTeamInput{
		TeamID: Int(1),
	}); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteProjectTeam(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/teams", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := fmt.Fprint(w, testJSONTeam); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.DeleteProjectTeam("SRE", &DeleteProjectTeamInput{
		TeamID: Int(1),
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestTeam()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestDeleteProjectTeamFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/teams", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.DeleteProjectTeam("SRE", &DeleteProjectTeamInput{
		TeamID: Int(1),
	}); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteProjectTeamInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	if _, err := client.DeleteProjectTeam("%", &DeleteProjectTeamInput{
		TeamID: Int(1),
	}); err == nil {
		t.Fatal("expected an error but got none")
	}
}
