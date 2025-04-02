package backlog

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/kylelemons/godebug/pretty"
	"github.com/pkg/errors"
)

func TestResetAsTime(t *testing.T) {
	ls := &LimitStatus{
		Reset: Int(1603881873),
	}
	expected := ls.ResetAsTime().UTC()
	want := time.Date(2020, time.October, 28, 10, 44, 33, 0, time.UTC)
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(errors.New(pretty.Compare(want, expected)))
	}
}

func TestResetAsTimeWithResetNull(t *testing.T) {
	ls := &LimitStatus{}
	expected := ls.ResetAsTime()
	want := time.Time{}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(errors.New(pretty.Compare(want, expected)))
	}
}

const testJSONRateLimit string = `{
	"rateLimit": {
	  "read": {
		"limit": 600,
		"remaining": 600,
		"reset": 1603881873
	  },
	  "update": {
		"limit": 150,
		"remaining": 150,
		"reset": 1603881873
	  },
	  "search": {
		"limit": 150,
		"remaining": 150,
		"reset": 1603881873
	  },
	  "icon": {
		"limit": 60,
		"remaining": 60,
		"reset": 1603881873
	  }
	}
  }`

func getTestRateLimit() *ResponseRateLimit {
	return &ResponseRateLimit{
		RateLimit: &RateLimit{
			Read: &LimitStatus{
				Limit:     Int(600),
				Remaining: Int(600),
				Reset:     Int(1603881873),
			},
			Update: &LimitStatus{
				Limit:     Int(150),
				Remaining: Int(150),
				Reset:     Int(1603881873),
			},
			Search: &LimitStatus{
				Limit:     Int(150),
				Remaining: Int(150),
				Reset:     Int(1603881873),
			},
			Icon: &LimitStatus{
				Limit:     Int(60),
				Remaining: Int(60),
				Reset:     Int(1603881873),
			},
		},
	}
}

func TestGetRateLimit(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/rateLimit", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := fmt.Fprint(w, testJSONRateLimit); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetRateLimit()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	r := getTestRateLimit()
	want := r.RateLimit
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetRateLimitFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/rateLimit", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetRateLimit(); err == nil {
		t.Fatal("expected an error but got none")
	}
}
