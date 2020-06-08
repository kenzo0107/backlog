package backlog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"testing"
)

func getTestSpace() *Space {
	return &Space{
		SpaceKey:           "nulab",
		Name:               "Nulab Inc.",
		OwnerID:            1,
		Lang:               "ja",
		Timezone:           "Asia/Tokyo",
		ReportSendTime:     "08:00:00",
		TextFormattingRule: "markdown",
		Created:            "2008-07-06T15:00:00Z",
		Updated:            "2013-06-18T07:55:37Z",
	}
}

func getTestSpaceNotification(content string) *SpaceNotification {
	return &SpaceNotification{
		Content: content,
		Updated: JSONTime("2013-06-18T07:55:37Z"),
	}
}

func getTestSpaceDiskUsage() *SpaceDiskUsage {
	return &SpaceDiskUsage{
		Capacity:   1073741824,
		Issue:      119511,
		Wiki:       48575,
		File:       0,
		Subversion: 0,
		Git:        0,
		GitLFS:     0,
		Details: []SpaceDiskUsageDetail{
			{
				ProjectID:  1,
				Issue:      11931,
				Wiki:       0,
				File:       0,
				Subversion: 0,
				Git:        0,
				GitLFS:     0,
			},
		},
	}
}

func errorResponse() ErrorResponse {
	return ErrorResponse{
		Errors: []Error{
			{
				Message:  "No space.",
				Code:     6,
				MoreInfo: "",
			},
		},
	}
}

func getErrorResponse(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusInternalServerError)
	response, _ := json.Marshal(errorResponse())
	if _, err := rw.Write(response); err != nil {
		fmt.Println(err)
	}
}

func getSpace(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(getTestSpace())
	if _, err := rw.Write(response); err != nil {
		fmt.Println(err)
	}
}

func getSpaceNotification(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(getTestSpaceNotification("test"))
	if _, err := rw.Write(response); err != nil {
		fmt.Println(err)
	}
}

func getSpaceDiskUsage(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(getTestSpaceDiskUsage())
	if _, err := rw.Write(response); err != nil {
		fmt.Println(err)
	}
}

func TestGetSpace(t *testing.T) {
	http.HandleFunc("/api/v2/space", getSpace)
	expected := getTestSpace()

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	space, err := api.GetSpace()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	if !reflect.DeepEqual(expected, space) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetSpaceFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/space", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	if _, err := api.GetSpace(); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetSpaceErrorResponse(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/space", getErrorResponse)

	once.Do(startServer)

	logger := log.New(os.Stdout, "backlog: ", log.Lshortfile|log.LstdFlags)
	api := New(
		"testing-token",
		"http://"+serverAddr+"/",
		OptionDebug(true),
		OptionHTTPClient(&http.Client{}),
		OptionLog(logger),
	)

	api.Debugf("%s", "test")

	if _, err := api.GetSpace(); err == nil {
		log.Println("err", err)
		t.Fatal("expected an error but got none")
	}
}

type mockHTTPClient struct{}

func (m *mockHTTPClient) Do(*http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBufferString(`OK`))}, nil
}

func TestGetSpaceIcon(t *testing.T) {
	api := &Client{
		endpoint:   "http://" + serverAddr + "/",
		apiKey:     "testing-token",
		httpclient: &mockHTTPClient{},
	}

	err := api.GetSpaceIcon(&bytes.Buffer{})
	if err != nil {
		log.Fatalf("Unexpected error: %s in test", err)
	}
}

func TestGetSpaceNotification(t *testing.T) {
	http.HandleFunc("/api/v2/space/notification", getSpaceNotification)
	expected := getTestSpaceNotification("test")

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	spaceNotification, err := api.GetSpaceNotification()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	if !reflect.DeepEqual(expected, spaceNotification) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetSpaceNotificationFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/space/notification", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	if _, err := api.GetSpaceNotification(); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateSpaceNotification(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/space/notification", getSpaceNotification)
	expected := getTestSpaceNotification("test")

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	spaceNotification, err := api.UpdateSpaceNotification("test")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	if !reflect.DeepEqual(expected, spaceNotification) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestUpdateSpaceNotificationFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/space/notification", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	if _, err := api.UpdateSpaceNotification("test"); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetSpaceDiskUsage(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/space/diskUsage", getSpaceDiskUsage)
	expected := getTestSpaceDiskUsage()

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	spaceNotification, err := api.GetSpaceDiskUsage()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	if !reflect.DeepEqual(expected, spaceNotification) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetSpaceDiskUsageFailed(t *testing.T) {
	http.DefaultServeMux = new(http.ServeMux)
	http.HandleFunc("/api/v2/space/diskUsage", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	once.Do(startServer)
	api := New("testing-token", "http://"+serverAddr+"/")

	if _, err := api.GetSpaceDiskUsage(); err == nil {
		t.Fatal("expected an error but got none")
	}
}
