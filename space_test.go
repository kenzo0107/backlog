package backlog

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"testing"
)

const testJSONSpace string = `{
	"spaceKey": "nulab",
	"name": "Nulab Inc.",
	"ownerId": 1,
	"lang": "ja",
	"timezone": "Asia/Tokyo",
	"reportSendTime": "08:00:00",
	"textFormattingRule": "markdown",
	"created": "2006-01-02T15:04:05Z",
	"updated": "2006-01-02T15:04:05Z"
}`

func getTestSpace() *Space {
	return &Space{
		SpaceKey:           String("nulab"),
		Name:               String("Nulab Inc."),
		OwnerID:            Int(1),
		Lang:               String("ja"),
		Timezone:           String("Asia/Tokyo"),
		ReportSendTime:     String("08:00:00"),
		TextFormattingRule: String("markdown"),
		Created:            &Timestamp{referenceTime},
		Updated:            &Timestamp{referenceTime},
	}
}

func getTestSpaceNotification() *SpaceNotification {
	return &SpaceNotification{
		Content: String("Notification"),
		Updated: &Timestamp{referenceTime},
	}
}

func getTestSpaceDiskUsage() *SpaceDiskUsage {
	return &SpaceDiskUsage{
		Capacity:   Int(1073741824),
		Issue:      Int(119511),
		Wiki:       Int(48575),
		File:       Int(0),
		Subversion: Int(0),
		Git:        Int(0),
		GitLFS:     Int(0),
		Details: []*SpaceDiskUsageDetail{
			{
				ProjectID:  Int(1),
				Issue:      Int(11931),
				Wiki:       Int(0),
				File:       Int(0),
				Subversion: Int(0),
				Git:        Int(0),
				GitLFS:     Int(0),
			},
		},
	}
}

func TestGetSpace(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/space", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, testJSONSpace); err != nil {
			t.Fatal(err)
		}
	})

	space, err := client.GetSpace()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestSpace()
	if !reflect.DeepEqual(want, space) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetSpaceFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/space", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetSpace(); err == nil {
		t.Fatal("expected an error but got none")
	}
}

type mockHTTPClient struct{}

func (m *mockHTTPClient) Do(*http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBufferString(`OK`))}, nil
}

func TestGetSpaceIcon(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	client.httpclient = &mockHTTPClient{}

	err := client.GetSpaceIcon(&bytes.Buffer{})
	if err != nil {
		log.Fatalf("Unexpected error: %s in test", err)
	}
}

func TestGetSpaceNotification(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/space/notification", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{"content": "Notification", "updated": "2006-01-02T15:04:05Z"}`); err != nil {
			t.Fatal(err)
		}
	})

	spaceNotification, err := client.GetSpaceNotification()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestSpaceNotification()
	if !reflect.DeepEqual(want, spaceNotification) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetSpaceNotificationFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/space/notification", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetSpaceNotification(); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateSpaceNotification(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/space/notification", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{"content": "Notification", "updated": "2006-01-02T15:04:05Z"}`); err != nil {
			t.Fatal(err)
		}
	})

	spaceNotification, err := client.UpdateSpaceNotification("test")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestSpaceNotification()
	if !reflect.DeepEqual(want, spaceNotification) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestUpdateSpaceNotificationFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/space/notification", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.UpdateSpaceNotification("test"); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetSpaceDiskUsage(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/space/diskUsage", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `
			{
				"capacity": 1073741824,
				"issue": 119511,
				"wiki": 48575,
				"file": 0,
				"subversion": 0,
				"git": 0,
				"gitLFS": 0,
				"details":[
					{
						"projectId": 1,
						"issue": 11931,
						"wiki": 0,
						"file": 0,
						"subversion": 0,
						"git": 0,
						"gitLFS": 0
					}
				]
			}
		`); err != nil {
			t.Fatal(err)
		}
	})

	spaceNotification, err := client.GetSpaceDiskUsage()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestSpaceDiskUsage()
	if !reflect.DeepEqual(want, spaceNotification) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetSpaceDiskUsageFailed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/space/diskUsage", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := client.GetSpaceDiskUsage(); err == nil {
		t.Fatal("expected an error but got none")
	}
}
