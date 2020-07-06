package backlog

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
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

func TestGetProjects(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf("[%s]", testJSONProject)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	input := &GetProjectsOptions{
		All:      Bool(true),
		Archived: Bool(false),
	}
	projects, err := client.GetProjects(input)
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

	input := &GetProjectsOptions{}
	_, err := client.GetProjects(input)
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

func TestGetProjectStatuses(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/1/statuses", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf("[%s]", testJSONProjectStatus)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	projectStatuses, err := client.GetProjectStatuses(1)
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

	if _, err := client.GetProjectStatuses(1); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetProjectStatusesWithInvalidProjectIDOrKeyFailed(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	if _, err := client.GetProjectStatuses(true); err == nil {
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
