package backlog

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/pkg/errors"
)

var testJSONUserWatching string = fmt.Sprintf(`{
	"id": 1,
	"resourceAlreadyRead": true,
	"note": "This is a note for the watching issue.",
	"type": "issue",
	"issue": %v,
	"lastContentUpdated":"2006-01-02T15:04:05Z",
	"created": "2006-01-02T15:04:05Z",
	"updated": "2006-01-02T15:04:05Z"
}`, testJSONIssue)

func getTestWatching() *Watching {
	return &Watching{
		ID:                  Int(1),
		ResourceAlreadyRead: Bool(true),
		Note:                String("This is a note for the watching issue."),
		Type:                String("issue"),
		Issue:               getTestIssue(),
		LastContentUpdated:  &Timestamp{referenceTime},
		Created:             &Timestamp{referenceTime},
		Updated:             &Timestamp{referenceTime},
	}
}

func TestGetUserWatchings(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/1/watchings", func(w http.ResponseWriter, r *http.Request) {
		j := fmt.Sprintf("[%s]", testJSONUserWatching)
		if _, err := fmt.Fprint(w, j); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetUserWatchings(1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*Watching{getTestWatching()}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetUserWatchingsFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/1/watchings", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetUserWatchings(1); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetUserWatchingsCount(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/1/watchings/count", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"count": 138
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetUserWatchingsCount(1, &GetUserWatchingsCountOptions{})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := 138
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetUserWatchingsCountFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/users/1/watchings/count", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetUserWatchingsCount(1, &GetUserWatchingsCountOptions{}); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetWatching(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/watchings/1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONUserWatching); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetWatching(1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestWatching()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetWatchingFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/watchings/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetWatching(1); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateWatching(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/watchings", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONUserWatching); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.CreateWatching(&CreateWatchingInput{
		IssueIDOrKey: String("BLG-1"),
		Note:         String("This is a note for the watching issue."),
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestWatching()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestCreateWatchingFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/watchings", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.CreateWatching(&CreateWatchingInput{
		IssueIDOrKey: String("BLG-1"),
		Note:         String("This is a note for the watching issue."),
	}); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateWatching(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/watchings/1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONUserWatching); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateWatching(1, &UpdateWatchingInput{
		Note: String("This is a note for the watching issue."),
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestWatching()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateWatchingFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/watchings/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.UpdateWatching(1, &UpdateWatchingInput{
		Note: String("This is a note for the watching issue."),
	}); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteWatching(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/watchings/1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONUserWatching); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.DeleteWatching(1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestWatching()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestDeleteWatchingFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/watchings/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.DeleteWatching(1); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestMarkAsReadWatching(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/watchings/1/markAsRead", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	if err := client.MarkAsReadWatching(1); err != nil {
		log.Fatalf("Unexpected error: %s in test", err)
	}
}

func TestMarkAsReadWatchingFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("watchings/1/markAsRead", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if err := client.MarkAsReadWatching(1); err == nil {
		t.Fatal("expected an error but got none")
	}
}
