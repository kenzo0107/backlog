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

func TestGetPullRequests(t *testing.T) {
	projectKey := "TEST"
	repoName := "test-repo"
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/projects/%s/git/repositories/%s/pullRequests", projectKey, repoName),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			_, _ = fmt.Fprint(w, `[
				{
					"id": 1,
					"projectId": 1,
					"repositoryId": 1,
					"number": 1,
					"summary": "Test pull request",
					"description": "Test description",
					"base": "master",
					"branch": "feature/test",
					"status": {
						"id": 1,
						"name": "Open"
					},
					"assignee": {
						"id": 1,
						"userId": "admin",
						"name": "admin",
						"roleType": 1,
						"lang": "ja",
						"mailAddress": "admin@example.com"
					},
					"issue": null,
					"baseCommit": "abc123",
					"branchCommit": "def456",
					"closeAt": null,
					"mergeAt": null,
					"createdUser": {
						"id": 1,
						"userId": "admin",
						"name": "admin",
						"roleType": 1,
						"lang": "ja",
						"mailAddress": "admin@example.com"
					},
					"created": "2015-05-21T05:36:00Z",
					"updatedUser": {
						"id": 1,
						"userId": "admin",
						"name": "admin",
						"roleType": 1,
						"lang": "ja",
						"mailAddress": "admin@example.com"
					},
					"updated": "2015-05-21T05:36:00Z"
				}
			]`)
		})

	pullRequests, err := client.GetPullRequests(projectKey, repoName, nil)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &ResponsePullRequests{
		{
			ID:           Int(1),
			ProjectID:    Int(1),
			RepositoryID: Int(1),
			Number:       Int(1),
			Summary:      String("Test pull request"),
			Description:  String("Test description"),
			Base:         String("master"),
			Branch:       String("feature/test"),
			Status: &Status{
				ID:   Int(1),
				Name: String("Open"),
			},
			Assignee: &User{
				ID:          Int(1),
				UserID:      String("admin"),
				Name:        String("admin"),
				RoleType:    RoleType(1),
				Lang:        String("ja"),
				MailAddress: String("admin@example.com"),
			},
			Issue:        nil,
			BaseCommit:   String("abc123"),
			BranchCommit: String("def456"),
			CloseAt:      nil,
			MergeAt:      nil,
			CreatedUser: &User{
				ID:          Int(1),
				UserID:      String("admin"),
				Name:        String("admin"),
				RoleType:    RoleType(1),
				Lang:        String("ja"),
				MailAddress: String("admin@example.com"),
			},
			Created: &Timestamp{time.Date(2015, 5, 21, 5, 36, 0, 0, time.UTC)},
			UpdatedUser: &User{
				ID:          Int(1),
				UserID:      String("admin"),
				Name:        String("admin"),
				RoleType:    RoleType(1),
				Lang:        String("ja"),
				MailAddress: String("admin@example.com"),
			},
			Updated: &Timestamp{time.Date(2015, 5, 21, 5, 36, 0, 0, time.UTC)},
		},
	}

	if !reflect.DeepEqual(pullRequests, want) {
		t.Fatal(errors.Errorf("Response is incorrect: %s", pretty.Compare(pullRequests, want)))
	}
}

func TestGetPullRequestsContext(t *testing.T) {
	projectKey := "TEST"
	repoName := "test-repo"
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/projects/%s/git/repositories/%s/pullRequests", projectKey, repoName),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			_, _ = fmt.Fprint(w, `[]`)
		})

	ctx := context.Background()
	_, err := client.GetPullRequestsContext(ctx, projectKey, repoName, nil)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestGetPullRequestsCount(t *testing.T) {
	projectKey := "TEST"
	repoName := "test-repo"
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/projects/%s/git/repositories/%s/pullRequests/count", projectKey, repoName),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			_, _ = fmt.Fprint(w, `{"count": 5}`)
		})

	count, err := client.GetPullRequestsCount(projectKey, repoName, nil)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &ResponsePullRequestCount{
		Count: Int(5),
	}

	if !reflect.DeepEqual(count, want) {
		t.Fatal(errors.Errorf("Response is incorrect: %s", pretty.Compare(count, want)))
	}
}

func TestGetPullRequest(t *testing.T) {
	projectKey := "TEST"
	repoName := "test-repo"
	number := 1
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/projects/%s/git/repositories/%s/pullRequests/%d", projectKey, repoName, number),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			_, _ = fmt.Fprint(w, `{
				"id": 1,
				"projectId": 1,
				"repositoryId": 1,
				"number": 1,
				"summary": "Test pull request",
				"description": "Test description",
				"base": "master",
				"branch": "feature/test",
				"status": {
					"id": 1,
					"name": "Open"
				},
				"assignee": null,
				"issue": null,
				"baseCommit": "abc123",
				"branchCommit": "def456",
				"closeAt": null,
				"mergeAt": null,
				"createdUser": {
					"id": 1,
					"userId": "admin",
					"name": "admin",
					"roleType": 1,
					"lang": "ja",
					"mailAddress": "admin@example.com"
				},
				"created": "2015-05-21T05:36:00Z",
				"updatedUser": {
					"id": 1,
					"userId": "admin",
					"name": "admin",
					"roleType": 1,
					"lang": "ja",
					"mailAddress": "admin@example.com"
				},
				"updated": "2015-05-21T05:36:00Z"
			}`)
		})

	pullRequest, err := client.GetPullRequest(projectKey, repoName, number)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &PullRequest{
		ID:           Int(1),
		ProjectID:    Int(1),
		RepositoryID: Int(1),
		Number:       Int(1),
		Summary:      String("Test pull request"),
		Description:  String("Test description"),
		Base:         String("master"),
		Branch:       String("feature/test"),
		Status: &Status{
			ID:   Int(1),
			Name: String("Open"),
		},
		Assignee:     nil,
		Issue:        nil,
		BaseCommit:   String("abc123"),
		BranchCommit: String("def456"),
		CloseAt:      nil,
		MergeAt:      nil,
		CreatedUser: &User{
			ID:          Int(1),
			UserID:      String("admin"),
			Name:        String("admin"),
			RoleType:    RoleType(1),
			Lang:        String("ja"),
			MailAddress: String("admin@example.com"),
		},
		Created: &Timestamp{time.Date(2015, 5, 21, 5, 36, 0, 0, time.UTC)},
		UpdatedUser: &User{
			ID:          Int(1),
			UserID:      String("admin"),
			Name:        String("admin"),
			RoleType:    RoleType(1),
			Lang:        String("ja"),
			MailAddress: String("admin@example.com"),
		},
		Updated: &Timestamp{time.Date(2015, 5, 21, 5, 36, 0, 0, time.UTC)},
	}

	if !reflect.DeepEqual(pullRequest, want) {
		t.Fatal(errors.Errorf("Response is incorrect: %s", pretty.Compare(pullRequest, want)))
	}
}

func TestCreatePullRequest(t *testing.T) {
	projectKey := "TEST"
	repoName := "test-repo"
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/projects/%s/git/repositories/%s/pullRequests", projectKey, repoName),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			_, _ = fmt.Fprint(w, `{
				"id": 1,
				"projectId": 1,
				"repositoryId": 1,
				"number": 1,
				"summary": "New pull request",
				"description": "New description",
				"base": "master",
				"branch": "feature/new",
				"status": {
					"id": 1,
					"name": "Open"
				}
			}`)
		})

	options := &CreatePullRequestOptions{
		Summary:     String("New pull request"),
		Description: String("New description"),
		Base:        String("master"),
		Branch:      String("feature/new"),
	}

	pullRequest, err := client.CreatePullRequest(projectKey, repoName, options)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &PullRequest{
		ID:           Int(1),
		ProjectID:    Int(1),
		RepositoryID: Int(1),
		Number:       Int(1),
		Summary:      String("New pull request"),
		Description:  String("New description"),
		Base:         String("master"),
		Branch:       String("feature/new"),
		Status: &Status{
			ID:   Int(1),
			Name: String("Open"),
		},
	}

	if !reflect.DeepEqual(pullRequest, want) {
		t.Fatal(errors.Errorf("Response is incorrect: %s", pretty.Compare(pullRequest, want)))
	}
}

func TestUpdatePullRequest(t *testing.T) {
	projectKey := "TEST"
	repoName := "test-repo"
	number := 1
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/projects/%s/git/repositories/%s/pullRequests/%d", projectKey, repoName, number),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "PATCH")
			_, _ = fmt.Fprint(w, `{
				"id": 1,
				"projectId": 1,
				"repositoryId": 1,
				"number": 1,
				"summary": "Updated pull request",
				"description": "Updated description"
			}`)
		})

	options := &UpdatePullRequestOptions{
		Summary:     String("Updated pull request"),
		Description: String("Updated description"),
	}

	pullRequest, err := client.UpdatePullRequest(projectKey, repoName, number, options)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &PullRequest{
		ID:           Int(1),
		ProjectID:    Int(1),
		RepositoryID: Int(1),
		Number:       Int(1),
		Summary:      String("Updated pull request"),
		Description:  String("Updated description"),
	}

	if !reflect.DeepEqual(pullRequest, want) {
		t.Fatal(errors.Errorf("Response is incorrect: %s", pretty.Compare(pullRequest, want)))
	}
}

func TestGetPullRequests_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://example.com/api/v2/")
	client.baseURL = invalidURL

	_, err := client.GetPullRequests("TEST", "test-repo", &GetPullRequestsOptions{})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}
