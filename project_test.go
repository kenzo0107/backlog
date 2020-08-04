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

const testJSONProject string = `{
	"id": 1,
	"projectKey": "TEST",
	"name": "test",
	"chartEnabled": false,
	"subtaskingEnabled": false,
	"projectLeaderCanEditProjectLeader": false,
	"textFormattingRule": "markdown",
	"archived":false
}`

const testJSONProjectStatus string = `{
	"id": 1,
	"projectId": 1,
	"name": "未対応",
	"color": "#ed8077",
	"displayOrder": 1000
}`

func getTestProject() *Project {
	return &Project{
		ID:                                Int(1),
		ProjectKey:                        String("TEST"),
		Name:                              String("test"),
		ChartEnabled:                      Bool(false),
		SubtaskingEnabled:                 Bool(false),
		ProjectLeaderCanEditProjectLeader: Bool(false),
		TextFormattingRule:                String("markdown"),
		Archived:                          Bool(false),
	}
}

func getTestProjectStatus() *Status {
	return &Status{
		ID:           Int(1),
		ProjectID:    Int(1),
		Name:         String("未対応"),
		Color:        String("#ed8077"),
		DisplayOrder: Int(1000),
	}
}

func getRecentlyViewedProject() *RecentlyViewedProject {
	return &RecentlyViewedProject{
		Project: getTestProject(),
		Updated: &Timestamp{referenceTime},
	}
}

func TestGetMyRecentlyViewedProjects(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/myself/recentlyViewedProjects", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf(`[{
			"project": %v,
			"updated": "2006-01-02T15:04:05Z"
		}]`, testJSONProject)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetMyRecentlyViewedProjects(&GetMyRecentlyViewedProjectsOptions{})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*RecentlyViewedProject{getRecentlyViewedProject()}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetMyRecentlyViewedProjectsFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/myself/recentlyViewedProjects", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetMyRecentlyViewedProjects(&GetMyRecentlyViewedProjectsOptions{
		Order: Order("asc"),
		Count: Int(30),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetProjects(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf("[%s]", testJSONProject)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	opts := &GetProjectsOptions{
		All:      Bool(true),
		Archived: Bool(false),
	}
	projects, err := client.GetProjects(opts)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*Project{getTestProject()}
	if !reflect.DeepEqual(want, projects) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetProjectsFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	opts := &GetProjectsOptions{}
	_, err := client.GetProjects(opts)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetProject(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONProject); err != nil {
			t.Fatal(err)
		}
	})

	project, err := client.GetProject(1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestProject()
	if !reflect.DeepEqual(want, project) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetProjectFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetProject(1); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetProjectWithInvalidProjectIDOrKeyFailed(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.GetProject(true)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetStatuses(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/1/statuses", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf("[%s]", testJSONProjectStatus)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	projectStatuses, err := client.GetStatuses(1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*Status{getTestProjectStatus()}
	if !reflect.DeepEqual(want, projectStatuses) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetProjectStatusesFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/1/statuses", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetStatuses(1); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetProjectStatusesWithInvalidProjectIDOrKeyFailed(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	if _, err := client.GetStatuses(true); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateProject(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONProject); err != nil {
			t.Fatal(err)
		}
	})

	input := &CreateProjectInput{
		Name:                              String("test"),
		Key:                               String("TEST"),
		ChartEnabled:                      Bool(false),
		SubtaskingEnabled:                 Bool(false),
		ProjectLeaderCanEditProjectLeader: Bool(false),
		TextFormattingRule:                String("markdown"),
	}
	project, err := client.CreateProject(input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestProject()
	if !reflect.DeepEqual(want, project) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestCreateProjectFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &CreateProjectInput{
		Name:                              String("test"),
		Key:                               String("TEST"),
		ChartEnabled:                      Bool(false),
		SubtaskingEnabled:                 Bool(false),
		ProjectLeaderCanEditProjectLeader: Bool(false),
		TextFormattingRule:                String("markdown"),
	}
	_, err := client.CreateProject(input)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateProjectWithInvalidTextFormattingRuleFailed(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	input := &CreateProjectInput{
		Name:                              String("test"),
		Key:                               String("TEST"),
		ChartEnabled:                      Bool(false),
		SubtaskingEnabled:                 Bool(false),
		ProjectLeaderCanEditProjectLeader: Bool(false),
		TextFormattingRule:                String("downtown"),
	}

	_, err := client.CreateProject(input)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateProject(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONProject); err != nil {
			t.Fatal(err)
		}
	})

	input := &UpdateProjectInput{
		Name:                              String("test"),
		Key:                               String("TEST"),
		ChartEnabled:                      Bool(false),
		SubtaskingEnabled:                 Bool(false),
		ProjectLeaderCanEditProjectLeader: Bool(false),
		TextFormattingRule:                String("markdown"),
		Archived:                          Bool(false),
	}
	project, err := client.UpdateProject(1, input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestProject()
	if !reflect.DeepEqual(want, project) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestUpdateProjectFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &UpdateProjectInput{
		Name:                              String("test"),
		ChartEnabled:                      Bool(false),
		SubtaskingEnabled:                 Bool(false),
		ProjectLeaderCanEditProjectLeader: Bool(false),
		TextFormattingRule:                String("markdown"),
	}
	if _, err := client.UpdateProject(1, input); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateProjectWithoutIDFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &UpdateProjectInput{
		Name:                              String("test"),
		ChartEnabled:                      Bool(false),
		SubtaskingEnabled:                 Bool(false),
		ProjectLeaderCanEditProjectLeader: Bool(false),
		TextFormattingRule:                String("markdown"),
	}
	if _, err := client.UpdateProject(1, input); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateProjectWithInvalidTextFormattingRuleFailed(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	input := &UpdateProjectInput{
		Name:                              String("test"),
		ChartEnabled:                      Bool(false),
		SubtaskingEnabled:                 Bool(false),
		ProjectLeaderCanEditProjectLeader: Bool(false),
		TextFormattingRule:                String("downtown"),
	}

	if _, err := client.UpdateProject(1, input); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteProject(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONProject); err != nil {
			t.Fatal(err)
		}
	})

	project, err := client.DeleteProject(1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestProject()
	if !reflect.DeepEqual(want, project) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestDeleteProjectFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.DeleteProject(1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteProjectWithInvalidProjectIDOrKeyFailed(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	if _, err := client.DeleteProject(true); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetProjectIcon(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	client.httpclient = &mockHTTPClient{}

	err := client.GetProjectIcon("SRE", &bytes.Buffer{})
	if err != nil {
		log.Fatalf("Unexpected error: %s in test", err)
	}
}

func TestGetProjectIconFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("//projects/SRE/image", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if err := client.GetProjectIcon("SRE", &bytes.Buffer{}); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetProjectIconInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	if err := client.GetProjectIcon("%", &bytes.Buffer{}); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestAddProjectUser(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/users", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONUser); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.AddProjectUser("SRE", &AddProjectUserInput{
		UserID: Int(1),
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestUser()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestAddProjectUserFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/users", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.AddProjectUser("SRE", &AddProjectUserInput{
		UserID: Int(1),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestAddProjectUserInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.AddProjectUser("%", &AddProjectUserInput{
		UserID: Int(1),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetProjectUsers(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/users", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf("[%s]", testJSONUser)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetProjectUsers("SRE", &GetProjectUsersOptions{
		ExcludeGroupMembers: Bool(false),
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*User{getTestUser()}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetProjectUsersFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/users", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetProjectUsers("SRE", &GetProjectUsersOptions{
		ExcludeGroupMembers: Bool(false),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetProjectUsersInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.GetProjectUsers("%", &GetProjectUsersOptions{
		ExcludeGroupMembers: Bool(false),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteProjectUser(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/users", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf("%s", testJSONUser)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.DeleteProjectUser("SRE", &DeleteProjectUserInput{
		UserID: Int(1),
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestUser()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestDeleteProjectUserFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/users", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.DeleteProjectUser("SRE", &DeleteProjectUserInput{
		UserID: Int(1),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteProjectUserInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.DeleteProjectUser("%", &DeleteProjectUserInput{
		UserID: Int(1),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestAddProjectAdministrator(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/administrators", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONUser); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.AddProjectAdministrator("SRE", &AddProjectAdministratorInput{
		UserID: Int(1),
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestUser()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestAddProjectAdministratorFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/administrators", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.AddProjectAdministrator("SRE", &AddProjectAdministratorInput{
		UserID: Int(1),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestAddProjectAdministratorInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.AddProjectAdministrator("%", &AddProjectAdministratorInput{
		UserID: Int(1),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetProjectAdministrators(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/administrators", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf(`[%s]`, testJSONUser)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetProjectAdministrators("SRE")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*User{getTestUser()}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetProjectAdministratorsFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/administrators", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetProjectAdministrators("SRE")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetProjectAdministratorsInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.GetProjectAdministrators("%")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteProjectAdministrator(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/administrators", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONUser); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.DeleteProjectAdministrator("SRE", &DeleteProjectAdministratorInput{
		UserID: Int(1),
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestUser()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestDeleteProjectAdministratorFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/administrators", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.DeleteProjectAdministrator("SRE", &DeleteProjectAdministratorInput{
		UserID: Int(1),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteProjectAdministratorInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.DeleteProjectAdministrator("%", &DeleteProjectAdministratorInput{
		UserID: Int(1),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateStatus(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/statuses", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONProjectStatus); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.CreateStatus("SRE", &CreateStatusInput{
		Name:  String("未対応"),
		Color: String("#ed8077"),
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestProjectStatus()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestCreateStatusFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/statuses", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.CreateStatus("SRE", &CreateStatusInput{
		Name:  String("未対応"),
		Color: String("#ed8077"),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateStatusInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.CreateStatus("%", &CreateStatusInput{
		Name:  String("未対応"),
		Color: String("#ed8077"),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateStatus(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/statuses/1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONProjectStatus); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateStatus("SRE", 1, &UpdateStatusInput{
		Name:  String("未対応"),
		Color: String("#ed8077"),
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestProjectStatus()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateStatusFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/statuses", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateStatus("SRE", 1, &UpdateStatusInput{
		Name:  String("未対応"),
		Color: String("#ed8077"),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateStatusInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.UpdateStatus("%", 1, &UpdateStatusInput{
		Name:  String("未対応"),
		Color: String("#ed8077"),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteStatus(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/statuses/1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONProjectStatus); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.DeleteStatus("SRE", 1, &DeleteStatusInput{
		SubstituteStatusID: Int(1),
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestProjectStatus()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestDeleteStatusFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/statuses", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.DeleteStatus("SRE", 1, &DeleteStatusInput{
		SubstituteStatusID: Int(1),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteStatusInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.DeleteStatus("%", 1, &DeleteStatusInput{
		SubstituteStatusID: Int(1),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestSortStatuses(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/statuses/updateDisplayOrder", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf(`[%s]`, testJSONProjectStatus)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.SortStatuses("SRE", &SortStatusesInput{
		StatusIDs: []int{1},
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*Status{getTestProjectStatus()}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestSortStatusesFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/statuses", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.SortStatuses("SRE", &SortStatusesInput{
		StatusIDs: []int{1},
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestSortStatusesInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.SortStatuses("%", &SortStatusesInput{
		StatusIDs: []int{1},
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetProjectDiskUsage(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/diskUsage", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"projectId": 1,
			"issue": 11931,
			"wiki": 0,
			"file": 0,
			"subversion": 0,
			"git": 0,
			"gitLFS": 0
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetProjectDiskUsage("SRE")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &ProjectDiskUsage{
		ProjectID:  Int(1),
		Issue:      Int(11931),
		Wiki:       Int(0),
		File:       Int(0),
		Subversion: Int(0),
		Git:        Int(0),
		GitLFS:     Int(0),
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetProjectDiskUsageFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/diskUsage", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetProjectDiskUsage("SRE")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetProjectDiskUsageInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.GetProjectDiskUsage("%")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}
