package backlog

import (
	"fmt"
	"net/http"
	"reflect"
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

	mux.HandleFunc("/priorities", func(w http.ResponseWriter, r *http.Request) {
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

	mux.HandleFunc("/priorities", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetPriorities(); err == nil {
		t.Fatal("expected an error but got none")
	}
}
