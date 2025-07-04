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

const testJSONPriority string = `{
	"id": 2,
	"name": "高"
}`

func getTestPriority() *Priority {
	return &Priority{
		ID:   Int(2),
		Name: String("高"),
	}
}

func TestGetPriorities(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/priorities", func(w http.ResponseWriter, _ *http.Request) {
		j := fmt.Sprintf("[%s]", testJSONPriority)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetPriorities()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*Priority{getTestPriority()}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetPrioritiesFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/priorities", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetPriorities(); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetPriorities_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://example.com/api/v2/")
	client.baseURL = invalidURL

	_, err := client.GetPriorities()
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}
