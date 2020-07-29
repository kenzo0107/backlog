package backlog

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/pkg/errors"
)

const testJSONVersion string = `{
	"id": 3,
	"projectId": 1,
	"name": "いますぐ",
	"description": "",
	"startDate": null,
	"releaseDueDate": null,
	"archived": false,
	"displayOrder": 0
}`

func getTestVersion() *Version {
	return &Version{
		ID:             Int(3),
		ProjectID:      Int(1),
		Name:           String("いますぐ"),
		Description:    String(""),
		StartDate:      nil,
		ReleaseDueDate: nil,
		Archived:       Bool(false),
		DisplayOrder:   Int(0),
	}
}

func TestGetVersions(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/versions", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf("[%s]", testJSONVersion)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetVersions("SRE")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*Version{getTestVersion()}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetVersionsFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/versions", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetVersions("SRE")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetVersionsInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.GetVersions("%%")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateVersion(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/versions", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf("%s", testJSONVersion)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.CreateVersion("SRE", &CreateVersionInput{
		Name:        String("いますぐ"),
		Description: String(""),
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestVersion()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestCreateVersionsFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/versions", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.CreateVersion("SRE", &CreateVersionInput{
		Name:        String("いますぐ"),
		Description: String(""),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateVersionsInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.CreateVersion("%%", &CreateVersionInput{
		Name:        String("いますぐ"),
		Description: String(""),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateVersion(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/versions/1", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf("%s", testJSONVersion)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateVersion("SRE", 1, &UpdateVersionInput{
		Name:        String("いますぐ"),
		Description: String(""),
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestVersion()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateVersionsFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/versions/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateVersion("SRE", 1, &UpdateVersionInput{
		Name:        String("いますぐ"),
		Description: String(""),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateVersionsInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.UpdateVersion("%%", 1, &UpdateVersionInput{
		Name:        String("いますぐ"),
		Description: String(""),
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteVersion(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/versions/1", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf("%s", testJSONVersion)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.DeleteVersion("SRE", 1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestVersion()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestDeleteVersionsFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects/SRE/versions/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.DeleteVersion("SRE", 1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteVersionsInvalidProjectKey(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	_, err := client.DeleteVersion("%%", 1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}
