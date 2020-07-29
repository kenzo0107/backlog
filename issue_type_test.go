package backlog

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/pkg/errors"
)

const testJSONIssueType string = `{
	"id": 1,
	"projectId": 1,
	"name": "バグ",
	"color": "#990000",
	"displayOrder": 0,
	"templateSummary": "件名",
	"templateDescription": "詳細"
}`

func getTestIssueType() *IssueType {
	return &IssueType{
		ID:                  Int(1),
		ProjectID:           Int(1),
		Name:                String("バグ"),
		Color:               String("#990000"),
		DisplayOrder:        Int(0),
		TemplateSummary:     String("件名"),
		TemplateDescription: String("詳細"),
	}
}

func TestGetIssueTypes(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/issueTypes", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf("[%s]", testJSONIssueType)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	issueTypes, err := client.GetIssueTypes("SRE")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*IssueType{getTestIssueType()}
	if !reflect.DeepEqual(want, issueTypes) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, issueTypes)))
	}
}

func TestGetIssueTypesFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/issueTypes", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetIssueTypes("SRE")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetIssueTypesInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.GetIssueTypes("%%")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateIssueType(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/issueTypes", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONIssueType); err != nil {
			t.Fatal(err)
		}
	})

	input := &CreateIssueTypeInput{
		Name:                String("バグ"),
		Color:               String("#990000"),
		TemplateSummary:     String("件名"),
		TemplateDescription: String("詳細"),
	}
	issueType, err := client.CreateIssueType("SRE", input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestIssueType()
	if !reflect.DeepEqual(want, issueType) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, issueType)))
	}
}

func TestCreateIssueTypesFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/issueTypes", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &CreateIssueTypeInput{
		Name:                String("バグ"),
		Color:               String("#990000"),
		TemplateSummary:     String("件名"),
		TemplateDescription: String("詳細"),
	}
	_, err := client.CreateIssueType("SRE", input)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateIssueTypesInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	input := &CreateIssueTypeInput{
		Name:                String("バグ"),
		Color:               String("#990000"),
		TemplateSummary:     String("件名"),
		TemplateDescription: String("詳細"),
	}
	_, err := client.CreateIssueType("%%", input)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateIssueType(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/issueTypes/1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONIssueType); err != nil {
			t.Fatal(err)
		}
	})

	input := &UpdateIssueTypeInput{
		Name:                String("バグ"),
		Color:               String("#990000"),
		TemplateSummary:     String("件名"),
		TemplateDescription: String("詳細"),
	}
	issueType, err := client.UpdateIssueType("SRE", 1, input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestIssueType()
	if !reflect.DeepEqual(want, issueType) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, issueType)))
	}
}

func TestUpdateIssueTypesFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/issueTypes/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &UpdateIssueTypeInput{
		Name:                String("バグ"),
		Color:               String("#990000"),
		TemplateSummary:     String("件名"),
		TemplateDescription: String("詳細"),
	}
	_, err := client.UpdateIssueType("SRE", 1, input)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateIssueTypesInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	input := &UpdateIssueTypeInput{
		Name:                String("バグ"),
		Color:               String("#990000"),
		TemplateSummary:     String("件名"),
		TemplateDescription: String("詳細"),
	}
	_, err := client.UpdateIssueType("%%", 1, input)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteIssueType(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/issueTypes/1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONIssueType); err != nil {
			t.Fatal(err)
		}
	})

	input := &DeleteIssueTypeInput{
		SubstituteIssueTypeID: Int(1),
	}
	issueType, err := client.DeleteIssueType("SRE", 1, input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestIssueType()
	if !reflect.DeepEqual(want, issueType) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, issueType)))
	}
}

func TestDeleteIssueTypesFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/issueTypes/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &DeleteIssueTypeInput{
		SubstituteIssueTypeID: Int(1),
	}
	_, err := client.DeleteIssueType("SRE", 1, input)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteIssueTypesInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	input := &DeleteIssueTypeInput{
		SubstituteIssueTypeID: Int(1),
	}
	_, err := client.DeleteIssueType("%%", 1, input)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}
