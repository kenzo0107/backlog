package backlog

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func getTestProject() Project {
	return getTestProjectWithID(1)
}

func getTestProjects() []Project {
	return []Project{
		getTestProjectWithID(1),
		getTestProjectWithID(2),
		getTestProjectWithID(3),
		getTestProjectWithID(4),
	}
}

func getTestProjectWithID(id int) Project {
	return Project{
		ID:                                id,
		ProjectKey:                        "TEST",
		Name:                              "test",
		ChartEnabled:                      false,
		SubtaskingEnabled:                 false,
		ProjectLeaderCanEditProjectLeader: false,
		TextFormattingRule:                "markdown",
		Archived:                          false,
	}
}

func getTestProjectStatuses() []ProjectStatus {
	return []ProjectStatus{
		ProjectStatus{
			ID:           1,
			ProjectID:    1,
			Name:         "未対応",
			Color:        "#ed8077",
			DisplayOrder: 1000,
		},
		ProjectStatus{
			ID:           2,
			ProjectID:    1,
			Name:         "処理中",
			Color:        "#4488c5",
			DisplayOrder: 2000,
		},
		ProjectStatus{
			ID:           3,
			ProjectID:    1,
			Name:         "処理済み",
			Color:        "#5eb5a6",
			DisplayOrder: 3000,
		},
		ProjectStatus{
			ID:           4,
			ProjectID:    1,
			Name:         "完了",
			Color:        "#b0be3c",
			DisplayOrder: 4000,
		},
	}
}

func getProjects(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(getTestProjects())
	if _, err := rw.Write(response); err != nil {
		fmt.Println(err)
	}
}

func getProject(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(getTestProject())
	if _, err := rw.Write(response); err != nil {
		fmt.Println(err)
	}
}

func getProjectStatuses(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(getTestProjectStatuses())
	if _, err := rw.Write(response); err != nil {
		fmt.Println(err)
	}
}

func TestGetProjects(t *testing.T) {
	http.HandleFunc("/api/v2/projects", getProjects)
	expected := getTestProjects()

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	projects, err := api.GetProjects(false, false)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	if !reflect.DeepEqual(expected, projects) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetProjectsFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/projects", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	_, err := api.GetProjects(false, false)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetProject(t *testing.T) {
	http.HandleFunc("/api/v2/projects/1", getProject)
	expected := getTestProject()

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	project, err := api.GetProject(1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	if !reflect.DeepEqual(expected, *project) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetProjectFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/projects/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	_, err := api.GetProject(1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetProjectWithInvalidProjectIDOrKeyFailed(t *testing.T) {
	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	_, err := api.GetProject(true)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetProjectStatuses(t *testing.T) {
	http.HandleFunc("/api/v2/projects/1/statuses", getProjectStatuses)
	expected := getTestProjectStatuses()

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	projectStatuses, err := api.GetProjectStatuses(1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	if !reflect.DeepEqual(expected, projectStatuses) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetProjectStatusesFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/projects/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	_, err := api.GetProjectStatuses(1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetProjectStatusesWithInvalidProjectIDOrKeyFailed(t *testing.T) {
	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	_, err := api.GetProjectStatuses(true)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateProject(t *testing.T) {
	http.HandleFunc("/api/v2/projects", getProject)
	expected := getTestProject()

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	input := &CreateProjectInput{
		Name:                              "test",
		Key:                               "TEST",
		ChartEnabled:                      false,
		SubtaskingEnabled:                 false,
		ProjectLeaderCanEditProjectLeader: false,
		TextFormattingRule:                "markdown",
	}
	project, err := api.CreateProject(input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	if !reflect.DeepEqual(expected, *project) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestCreateProjectFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/projects", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	input := &CreateProjectInput{
		Name:                              "test",
		Key:                               "TEST",
		ChartEnabled:                      false,
		SubtaskingEnabled:                 false,
		ProjectLeaderCanEditProjectLeader: false,
		TextFormattingRule:                "markdown",
	}
	_, err := api.CreateProject(input)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateProjectWithInvalidTextFormattingRuleFailed(t *testing.T) {
	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	input := &CreateProjectInput{
		Name:                              "test",
		Key:                               "TEST",
		ChartEnabled:                      false,
		SubtaskingEnabled:                 false,
		ProjectLeaderCanEditProjectLeader: false,
		TextFormattingRule:                "downtown",
	}

	_, err := api.CreateProject(input)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateProject(t *testing.T) {
	http.HandleFunc("/api/v2/projects/1", getProject)
	expected := getTestProject()

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	input := &UpdateProjectInput{
		ID:                                1,
		Name:                              "test",
		ProjectKey:                        "TEST",
		ChartEnabled:                      false,
		SubtaskingEnabled:                 false,
		ProjectLeaderCanEditProjectLeader: false,
		TextFormattingRule:                "markdown",
	}
	project, err := api.UpdateProject(input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	if !reflect.DeepEqual(expected, *project) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestUpdateProjectFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/projects/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	input := &UpdateProjectInput{
		Name:                              "test",
		ChartEnabled:                      false,
		SubtaskingEnabled:                 false,
		ProjectLeaderCanEditProjectLeader: false,
		TextFormattingRule:                "markdown",
	}
	_, err := api.UpdateProject(input)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateProjectWithInvalidTextFormattingRuleFailed(t *testing.T) {
	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	input := &UpdateProjectInput{
		Name:                              "test",
		ChartEnabled:                      false,
		SubtaskingEnabled:                 false,
		ProjectLeaderCanEditProjectLeader: false,
		TextFormattingRule:                "downtown",
	}

	_, err := api.UpdateProject(input)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteProject(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/projects/1", getProject)
	expected := getTestProject()

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	project, err := api.DeleteProject(1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	if !reflect.DeepEqual(expected, *project) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestDeleteProjectFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/projects/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	_, err := api.DeleteProject(1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteProjectWithInvalidProjectIDOrKeyFailed(t *testing.T) {
	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	_, err := api.DeleteProject(true)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}
