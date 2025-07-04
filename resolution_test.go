package backlog

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/pkg/errors"
)

func getTestResolution() *Resolution {
	return &Resolution{
		ID:   Int(0),
		Name: String("対応済み"),
	}
}

func TestGetResolutions(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/resolutions", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := fmt.Fprint(w, `[{"id": 0, "name": "対応済み"}]`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetResolutions()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*Resolution{getTestResolution()}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetResolutionsFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/resolutions", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetResolutions(); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetResolutions_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://example.com/api/v2/")
	client.baseURL = invalidURL

	_, err := client.GetResolutions()
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}
