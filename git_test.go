package backlog

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/kylelemons/godebug/pretty"
	"github.com/pkg/errors"
)

func TestGetGitRepositories(t *testing.T) {
	projectKey := "TEST"
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/projects/%s/git/repositories", projectKey),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			_, _ = fmt.Fprint(w, `[
				{
					"id": 1,
					"projectId": 1,
					"name": "test-repo",
					"description": "Test repository",
					"hookUrl": null,
					"httpUrl": "https://test.backlog.com/git/TEST/test-repo.git",
					"sshUrl": "git@test.backlog.com:TEST/test-repo.git",
					"displayOrder": 0,
					"pushedAt": "2015-05-21T05:36:00Z",
					"createdUser": {
						"id": 1,
						"userId": "admin",
						"name": "admin",
						"roleType": 1,
						"lang": "ja",
						"mailAddress": "eguchi@nulab.example"
					},
					"created": "2015-05-21T05:36:00Z",
					"updatedUser": {
						"id": 1,
						"userId": "admin",
						"name": "admin",
						"roleType": 1,
						"lang": "ja",
						"mailAddress": "eguchi@nulab.example"
					},
					"updated": "2015-05-21T05:36:00Z"
				}
			]`)
		})

	gitRepositories, err := client.GetGitRepositories(projectKey)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &ResponseGitRepositories{
		{
			ID:           Int(1),
			ProjectID:    Int(1),
			Name:         String("test-repo"),
			Description:  String("Test repository"),
			HookURL:      nil,
			HTTPURL:      String("https://test.backlog.com/git/TEST/test-repo.git"),
			SSHURL:       String("git@test.backlog.com:TEST/test-repo.git"),
			DisplayOrder: Int(0),
			PushedAt:     &Timestamp{time.Date(2015, 5, 21, 5, 36, 0, 0, time.UTC)},
			CreatedUser: &User{
				ID:          Int(1),
				UserID:      String("admin"),
				Name:        String("admin"),
				RoleType:    RoleType(1),
				Lang:        String("ja"),
				MailAddress: String("eguchi@nulab.example"),
			},
			Created: &Timestamp{time.Date(2015, 5, 21, 5, 36, 0, 0, time.UTC)},
			UpdatedUser: &User{
				ID:          Int(1),
				UserID:      String("admin"),
				Name:        String("admin"),
				RoleType:    RoleType(1),
				Lang:        String("ja"),
				MailAddress: String("eguchi@nulab.example"),
			},
			Updated: &Timestamp{time.Date(2015, 5, 21, 5, 36, 0, 0, time.UTC)},
		},
	}

	if !reflect.DeepEqual(gitRepositories, want) {
		t.Fatal(errors.Errorf("Response is incorrect: %s", pretty.Compare(gitRepositories, want)))
	}
}

func TestGetGitRepositoriesContext(t *testing.T) {
	projectKey := "TEST"
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/projects/%s/git/repositories", projectKey),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			_, _ = fmt.Fprint(w, `[]`)
		})

	ctx := context.Background()
	_, err := client.GetGitRepositoriesContext(ctx, projectKey)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestGetGitRepository(t *testing.T) {
	projectKey := "TEST"
	repoName := "test-repo"
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/projects/%s/git/repositories/%s", projectKey, repoName),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			_, _ = fmt.Fprint(w, `{
				"id": 1,
				"projectId": 1,
				"name": "test-repo",
				"description": "Test repository",
				"hookUrl": null,
				"httpUrl": "https://test.backlog.com/git/TEST/test-repo.git",
				"sshUrl": "git@test.backlog.com:TEST/test-repo.git",
				"displayOrder": 0,
				"pushedAt": "2015-05-21T05:36:00Z",
				"createdUser": {
					"id": 1,
					"userId": "admin",
					"name": "admin",
					"roleType": 1,
					"lang": "ja",
					"mailAddress": "eguchi@nulab.example"
				},
				"created": "2015-05-21T05:36:00Z",
				"updatedUser": {
					"id": 1,
					"userId": "admin",
					"name": "admin",
					"roleType": 1,
					"lang": "ja",
					"mailAddress": "eguchi@nulab.example"
				},
				"updated": "2015-05-21T05:36:00Z"
			}`)
		})

	gitRepository, err := client.GetGitRepository(projectKey, repoName)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &GitRepository{
		ID:           Int(1),
		ProjectID:    Int(1),
		Name:         String("test-repo"),
		Description:  String("Test repository"),
		HookURL:      nil,
		HTTPURL:      String("https://test.backlog.com/git/TEST/test-repo.git"),
		SSHURL:       String("git@test.backlog.com:TEST/test-repo.git"),
		DisplayOrder: Int(0),
		PushedAt:     &Timestamp{time.Date(2015, 5, 21, 5, 36, 0, 0, time.UTC)},
		CreatedUser: &User{
			ID:          Int(1),
			UserID:      String("admin"),
			Name:        String("admin"),
			RoleType:    RoleType(1),
			Lang:        String("ja"),
			MailAddress: String("eguchi@nulab.example"),
		},
		Created: &Timestamp{time.Date(2015, 5, 21, 5, 36, 0, 0, time.UTC)},
		UpdatedUser: &User{
			ID:          Int(1),
			UserID:      String("admin"),
			Name:        String("admin"),
			RoleType:    RoleType(1),
			Lang:        String("ja"),
			MailAddress: String("eguchi@nulab.example"),
		},
		Updated: &Timestamp{time.Date(2015, 5, 21, 5, 36, 0, 0, time.UTC)},
	}

	if !reflect.DeepEqual(gitRepository, want) {
		t.Fatal(errors.Errorf("Response is incorrect: %s", pretty.Compare(gitRepository, want)))
	}
}

func TestGetGitRepositoryContext(t *testing.T) {
	projectKey := "TEST"
	repoName := "test-repo"
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/projects/%s/git/repositories/%s", projectKey, repoName),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			_, _ = fmt.Fprint(w, `{}`)
		})

	ctx := context.Background()
	_, err := client.GetGitRepositoryContext(ctx, projectKey, repoName)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestGetGitRepositories_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://example.com/api/v2/")
	client.baseURL = invalidURL

	_, err := client.GetGitRepositories("TEST")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetGitRepository_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://example.com/api/v2/")
	client.baseURL = invalidURL

	_, err := client.GetGitRepository("TEST", "test-repo")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}
